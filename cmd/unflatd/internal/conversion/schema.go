package conversion

import (
	"errors"
	"log/slog"
	"sort"
	"strconv"

	"github.com/arisu-archive/arona-unflatd/cmd/unflatd/internal/utils"
	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
)

var (
	ErrNoTablesOrEnumsFound = errors.New("no tables or enums found")
	ErrNoFieldsFound        = errors.New("no fields found")
)

type SchemaConverter struct {
	logger *slog.Logger
}

func NewSchemaConverter(logger *slog.Logger) *SchemaConverter {
	return &SchemaConverter{logger: logger}
}

func (sc *SchemaConverter) Convert(info *ast.FileInfo) (*fbs.Schema, error) {
	schema := &fbs.Schema{
		Namespace: info.Namespace,
	}

	sc.processStructs(info, schema)
	sc.processEnums(info, schema)
	if len(schema.Tables) == 0 && len(schema.Enums) == 0 {
		return nil, ErrNoTablesOrEnumsFound
	}
	return schema, nil
}

func (sc *SchemaConverter) processStructs(info *ast.FileInfo, schema *fbs.Schema) {
	for _, structInfo := range info.Structs {
		table := sc.createTable(structInfo)
		if len(table.Fields) == 0 {
			sc.logger.Debug("no fields found in table", "table", structInfo.Name)
			continue
		}
		if !utils.Contains(structInfo.BaseList, []string{"IFlatbufferObject"}) {
			sc.logger.Debug(
				"struct is not a flatbuffer type",
				"struct",
				structInfo.Name,
				"baseList",
				structInfo.BaseList,
			)
			continue
		}
		if structInfo.HasMethod("Finish" + structInfo.Name + "Buffer") {
			schema.RootTypes = append(schema.RootTypes, structInfo.Name)
		}
		schema.Tables = append(schema.Tables, table)
	}
}

func sortByRVA(fields []*ast.FieldInfo) func(i, j int) bool {
	return func(i, j int) bool {
		if len(fields[i].Accessors) == 0 {
			return false
		}
		if len(fields[j].Accessors) == 0 {
			return true
		}
		accessorI := fields[i].Accessors["Address"]
		accessorJ := fields[j].Accessors["Address"]
		// Convert the RVA to an integer
		rvaI, err := strconv.ParseInt(accessorI["RVA"], 16, 64)
		if err != nil {
			return false
		}
		rvaJ, err := strconv.ParseInt(accessorJ["RVA"], 16, 64)
		if err != nil {
			return false
		}
		return rvaI < rvaJ
	}
}

func (sc *SchemaConverter) createTable(structInfo *ast.StructInfo) fbs.Table {
	if len(structInfo.Fields) == 0 {
		return fbs.Table{}
	}

	table := fbs.Table{
		Name: structInfo.Name,
	}

	sort.Slice(structInfo.Fields, sortByRVA(structInfo.Fields))
	for _, fieldData := range structInfo.Fields {
		// We need to filter out the fields that are not public and are overridden by properties
		sc.logger.Debug("Field", "name", fieldData.Name, "modifiers", fieldData.Modifiers, "type", fieldData.Type)
		if !structInfo.HasMethod("Add"+fieldData.Name) && !structInfo.IsVector(fieldData.Name, fieldData.Type) {
			sc.logger.Debug("Field is not an add method", "name", fieldData.Name, "modifiers", fieldData.Modifiers, "type", fieldData.Type)
			continue
		}
		field := sc.createField(structInfo, fieldData.Name, fieldData.Type)
		if table.FieldExists(field.Name) {
			field.Name = field.Name + "_" + utils.Checksum(field.Name+field.Type)
		}
		table.Fields = append(table.Fields, field)
	}
	return table
}

func (sc *SchemaConverter) createField(structInfo *ast.StructInfo, fieldName, fieldType string) fbs.Field {
	if structInfo.IsVector(fieldName, fieldType) {
		vectorFieldName := structInfo.ToVectorFieldName(fieldName)
		vectorFieldType, err := structInfo.GetVectorFieldType(vectorFieldName)
		if err != nil {
			// Since this is called from within a loop, we'll log the error and continue
			sc.logger.Error("failed to get vector field type", "error", err)
			return fbs.Field{Name: utils.ToSnakeCase(vectorFieldName), Type: fieldType}
		}
		return fbs.Field{
			Name:    utils.ToSnakeCase(vectorFieldName),
			Type:    ConvertFieldType(vectorFieldType),
			IsArray: true,
		}
	}
	return fbs.Field{
		Name: utils.ToSnakeCase(fieldName),
		Type: ConvertFieldType(fieldType),
	}
}

func (sc *SchemaConverter) processEnums(info *ast.FileInfo, schema *fbs.Schema) {
	for _, enumInfo := range info.Enums {
		if info.Namespace != "FlatData" {
			sc.logger.Debug("enum is not in FlatData namespace", "enum", enumInfo.Name)
			continue
		}
		dataType := ConvertFieldType(enumInfo.BaseType)
		schema.Enums = append(schema.Enums, fbs.Enum{
			Name:     enumInfo.Name,
			DataType: dataType,
			Values:   ConvertEnumValues(dataType, utils.Zip(enumInfo.Keys, enumInfo.Values)),
		})
	}
}

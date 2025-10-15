package conversion

import (
	"errors"
	"log/slog"

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
		table := sc.createTable(structInfo)
		if len(table.Fields) == 0 {
			sc.logger.Debug("no fields found in table", "table", structInfo.Name)
			continue
		}
		schema.Tables = append(schema.Tables, table)
	}
}

func (sc *SchemaConverter) createTable(structInfo *ast.StructInfo) fbs.Table {
	if len(structInfo.Fields) == 0 {
		return fbs.Table{}
	}

	table := fbs.Table{
		Name: structInfo.Name,
	}

	// Filter the invalid fields first
	validFields := make([]*ast.FieldInfo, 0, len(structInfo.Fields))
	for _, field := range structInfo.Fields {
		// We need to filter out the fields that are not public and are overridden by properties
		sc.logger.Debug("Field", "name", field.Name, "modifiers", field.Modifiers, "type", field.Type)
		if !structInfo.HasMethod("Add"+field.Name) && !structInfo.IsVector(field.Name, field.Type) {
			sc.logger.Debug("Field is not an add method", "name", field.Name, "modifiers", field.Modifiers, "type", field.Type)
			continue
		}
		validFields = append(validFields, field)
	}

	for _, field := range validFields {
		field := sc.createField(structInfo, field.Name, field.Type)
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

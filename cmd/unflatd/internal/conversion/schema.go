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

	// FIXME: This is a temporary fix to make sure the namespace is set to FlatData
	if schema.Namespace == "" {
		schema.Namespace = "FlatData"
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
			sc.logger.Warn("no fields found in table", "table", structInfo.Name)
			continue
		}
		if !utils.Contains(structInfo.BaseList, []string{"IFlatbufferObject"}) {
			sc.logger.Warn(
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

func (sc *SchemaConverter) createTable(structInfo *ast.StructInfo) fbs.Table {
	if len(structInfo.Fields) == 0 {
		return fbs.Table{}
	}

	table := fbs.Table{
		Name: structInfo.Name,
	}

	for _, fieldData := range structInfo.Fields {
		// We need to filter out the fields that are not public and are overridden by properties
		if !utils.Contains(fieldData.Modifiers, []string{"public", "static"}) {
			continue
		}
		if utils.Contains(fieldData.Modifiers, []string{"override"}) {
			continue
		}
		sc.logger.Debug("Field", "name", fieldData.Name, "type", fieldData.Type)
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
			sc.logger.Warn("enum is not in FlatData namespace", "enum", enumInfo.Name)
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

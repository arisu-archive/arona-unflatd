package fbs

import (
	"fmt"
	"strings"
)

type Visitor interface {
	VisitSchema(*Schema) string
	VisitEnum(*Enum)
	VisitTable(*Table)
	VisitField(*Field)
}

type SchemaVisitor struct {
	printer *Printer
}

func NewSchemaVisitor() *SchemaVisitor {
	return &SchemaVisitor{printer: NewPrinter()}
}

func (v *SchemaVisitor) VisitSchema(s *Schema) string {
	if len(s.Imports) > 0 {
		for _, imp := range s.Imports {
			v.printer.Println("include \"%s.fbs\";", imp)
		}
		v.printer.Newline()
	}

	if s.Namespace != "" {
		v.printer.Println("namespace %s;", s.Namespace)
		v.printer.Newline()
	}

	for _, enum := range s.Enums {
		v.VisitEnum(&enum)
		v.printer.Newline()
	}

	for _, table := range s.Tables {
		v.VisitTable(&table)
		v.printer.Newline()
	}

	for _, rt := range s.RootTypes {
		v.printer.Println("root_type %s;", rt)
		v.printer.Newline()
	}

	return strings.TrimSpace(v.printer.Flush())
}

func (v *SchemaVisitor) VisitEnum(e *Enum) {
	v.printer.Println("enum %s: %s {", e.Name, e.DataType)
	v.printer.Indent()
	for _, value := range e.Values {
		v.printer.Println("%s = %v,", value.Name, value.Value)
	}
	v.printer.Unindent()
	v.printer.Println("}")
}

func (v *SchemaVisitor) VisitTable(t *Table) {
	v.printer.Println("table %s {", t.Name)
	v.printer.Indent()
	for _, field := range t.Fields {
		dataType := field.Type
		if field.IsArray {
			dataType = fmt.Sprintf("[%s]", dataType)
		}
		v.printer.Println("%s: %s;", field.Name, dataType)
	}
	v.printer.Unindent()
	v.printer.Println("}")
}

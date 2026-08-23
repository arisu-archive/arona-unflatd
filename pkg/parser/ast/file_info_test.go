package ast_test

import (
	"testing"

	"github.com/arisu-archive/arona-unflatd/pkg/parser/ast"
)

func TestStructInfoIsVector(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fieldName string
		fieldType string
		methods   map[string]*ast.MethodInfo
		want      bool
	}{
		{
			name:      "generated vector pattern",
			fieldName: "ItemsLength",
			fieldType: "int",
			methods: vectorMethods(
				"Items",
			),
			want: true,
		},
		{
			name:      "length property has wrong type",
			fieldName: "ItemsLength",
			fieldType: "float",
			methods: vectorMethods(
				"Items",
			),
			want: false,
		},
		{
			name:      "int property lacks length suffix",
			fieldName: "ItemsCount",
			fieldType: "int",
			methods: vectorMethods(
				"ItemsCount",
			),
			want: false,
		},
		{
			name:      "start method is missing",
			fieldName: "ItemsLength",
			fieldType: "int",
			methods: map[string]*ast.MethodInfo{
				"CreateItemsVector": {Name: "CreateItemsVector"},
			},
			want: false,
		},
		{
			name:      "create method is missing",
			fieldName: "ItemsLength",
			fieldType: "int",
			methods: map[string]*ast.MethodInfo{
				"StartItemsVector": {Name: "StartItemsVector"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			info := &ast.StructInfo{Methods: tt.methods}
			got := info.IsVector(tt.fieldName, tt.fieldType)
			if got != tt.want {
				t.Errorf("StructInfo.IsVector(%q, %q) = %t, want %t", tt.fieldName, tt.fieldType, got, tt.want)
			}
		})
	}
}

func vectorMethods(fieldName string) map[string]*ast.MethodInfo {
	return map[string]*ast.MethodInfo{
		"Start" + fieldName + "Vector":  {Name: "Start" + fieldName + "Vector"},
		"Create" + fieldName + "Vector": {Name: "Create" + fieldName + "Vector"},
	}
}

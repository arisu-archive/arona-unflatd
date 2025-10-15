package ast

import (
	"fmt"
	"strings"
)

type FileInfo struct {
	FileName  string
	Namespace string
	Enums     map[string]*EnumInfo
	Structs   map[string]*StructInfo
}

type EnumInfo struct {
	Name     string
	BaseType string
	Keys     []string
	Values   []string
}

func (e *EnumInfo) Merge(other *EnumInfo) {
	e.Keys = append(e.Keys, other.Keys...)
	e.Values = append(e.Values, other.Values...)
}

type StructInfo struct {
	Name     string
	BaseList []string
	Fields   []*FieldInfo
	Methods  map[string]*MethodInfo
}

type FieldInfo struct {
	Modifiers []string
	Name      string
	Type      string
}

type MethodInfo struct {
	Name           string
	Modifiers      []string
	ReturnType     string
	ParameterNames []string
	ParameterTypes []string
}

func (m *MethodInfo) Merge(other *MethodInfo) {
	m.ParameterNames = append(m.ParameterNames, other.ParameterNames...)
	m.ParameterTypes = append(m.ParameterTypes, other.ParameterTypes...)
}

type ParameterInfo struct {
	Name string
	Type string
}

func (s *StructInfo) GetMethod(methodName string) (*MethodInfo, error) {
	for _, method := range s.Methods {
		if method.Name == methodName {
			return method, nil
		}
	}
	return nil, fmt.Errorf("method %s not found", methodName)
}

func (s *StructInfo) HasMethod(methodName string) bool {
	_, err := s.GetMethod(methodName)
	return err == nil
}

func (s *StructInfo) IsVector(fieldName, fieldType string) bool {
	// First condition: field name is end with "Length" and field type is int
	if !strings.HasSuffix(fieldName, "Length") && fieldType != "int" {
		return false
	}

	// Remove the "Length" suffix
	vectorFieldName := s.ToVectorFieldName(fieldName)
	// Second condition: method "Start" + field name + "Vector" exists
	if _, err := s.GetMethod("Start" + vectorFieldName + "Vector"); err != nil {
		return false
	}

	// Third condition: method "Create" + field name + "Vector" exists
	if _, err := s.GetMethod("Create" + vectorFieldName + "Vector"); err != nil {
		return false
	}

	return true
}

func (*StructInfo) ToVectorFieldName(fieldName string) string {
	return strings.TrimSuffix(fieldName, "Length")
}

func (s *StructInfo) GetVectorFieldType(fieldName string) (string, error) {
	method, err := s.GetMethod(s.ToVectorFieldName(fieldName))
	if err != nil {
		return "", err
	}
	return method.ReturnType, nil
}

package conversion

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/arisu-archive/arona-unflatd/pkg/fbs"
)

func ExtractNamespace(fieldType string) string {
	if idx := strings.LastIndex(fieldType, "."); idx != -1 {
		return fieldType[:idx]
	}
	return ""
}

func ConvertFieldType(fieldType string) string {
	// Handle Nullable<T>
	if strings.HasPrefix(fieldType, "Nullable<") {
		return strings.TrimPrefix(strings.TrimSuffix(fieldType, ">"), "Nullable<")
	}
	// Handle Nullable types like "int?", "long?", etc.
	if strings.HasSuffix(fieldType, "?") {
		return strings.TrimSuffix(fieldType, "?")
	}

	if fieldType == "sbyte" {
		return "byte"
	}

	return fieldType
}

func ConvertEnumValues(dataType string, values map[string]string) ([]fbs.EnumValue, error) {
	// sbyte,byte,short,ushort,int,uint,long,ulong.
	converted := make([]fbs.EnumValue, 0)
	switch dataType {
	case "int", "short", "sbyte", "long":
		for k, v := range values {
			if val, err := strconv.ParseInt(v, 10, 64); err == nil {
				converted = append(converted, fbs.EnumValue{
					Name:  k,
					Value: val,
				})
			}
		}
	case "uint", "ushort", "byte", "ulong":
		for k, v := range values {
			if val, err := strconv.ParseUint(v, 10, 64); err == nil {
				converted = append(converted, fbs.EnumValue{
					Name:  k,
					Value: val,
				})
			}
		}
	default:
		return nil, fmt.Errorf("unsupported enum type %q", dataType)
	}
	slices.SortFunc(converted, compareEnumValues)
	return converted, nil
}

func compareEnumValues(left, right fbs.EnumValue) int {
	switch leftValue := left.Value.(type) {
	case int64:
		rightValue, ok := right.Value.(int64)
		if !ok {
			return 0
		}
		return cmp.Compare(leftValue, rightValue)
	case uint64:
		rightValue, ok := right.Value.(uint64)
		if !ok {
			return 0
		}
		return cmp.Compare(leftValue, rightValue)
	default:
		return 0
	}
}

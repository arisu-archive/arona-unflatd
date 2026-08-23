package conversion

import (
	"fmt"
	"sort"
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
	// Sort the values by value. As the value is any, we need to convert it to the correct type.
	// We need to use reflection to get the type of the value.
	sort.Slice(converted, func(i, j int) bool {
		switch v := converted[i].Value.(type) {
		case int64:
			return v < converted[j].Value.(int64)
		case uint64:
			return v < converted[j].Value.(uint64)
		default:
			return false
		}
	})
	return converted, nil
}

package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func Zip[T comparable](a, b []T) map[T]T {
	if len(a) != len(b) {
		panic(fmt.Sprintf("slice lengths do not match: %d != %d", len(a), len(b)))
	}
	result := make(map[T]T, len(a))
	for i := range a {
		result[a[i]] = b[i]
	}
	return result
}

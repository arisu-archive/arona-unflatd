package utils

import "slices"

func Contains(slice, items []string) bool {
	for _, item := range items {
		if slices.Contains(slice, item) {
			return true
		}
	}
	return false
}

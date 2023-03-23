package util

import "strings"

func AnyEmpty(params ...string) bool {
	for _, item := range params {
		if item == "" || strings.Trim(item, " ") == "" {
			return true
		}
	}

	return false
}

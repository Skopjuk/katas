package main

import "strings"

func ToCamelCase(s string) string {
	var list []string
	if strings.Contains(s, "_") {
		list = strings.Split(s, "_")
	} else {
		list = strings.Split(s, "-")
	}

	for i := 1; i < len(list); i++ {
		list[i] = strings.Title(list[i])
	}
	newStr := strings.Join(list, "")
	return newStr
}

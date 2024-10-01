package main

import "strings"

func reverseString(str string) string {
	newStr := ""
	for _, i := range str {
		newStr = string(i) + newStr
	}
	return newStr
}

func SpinWords(str string) string {
	list := strings.Split(str, " ")
	var newStr []string

	for _, item := range list {
		if len(item) >= 5 {
			newStr = append(newStr, reverseString(item))
		} else {
			newStr = append(newStr, item)
		}
	}
	return strings.Join(newStr, " ")
}

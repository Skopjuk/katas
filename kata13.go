package main

import (
	"fmt"
	"strconv"
)

func ValidISBN10(isbn string) bool {

	var list []int
	for i, n := range isbn {
		if (string(n) == "X" || string(n) == "x") && i == 9 {
			list = append(list, (i+1)*10)
		} else if n > 57 || n < 48 {
			return false
		} else {
			k, _ := strconv.Atoi(string(n))
			list = append(list, (i+1)*k)
		}
	}

	sumOfList := 0
	for _, n := range list {
		sumOfList += n
	}
	fmt.Println(sumOfList)
	if sumOfList%11 == 0 {
		return true
	} else {
		return false
	}
}

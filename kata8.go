package main

import "fmt"

func reverse2(ls []int) []int {
	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}
	return ls
}

func Decode(roman string) int {
	arabic := 0
	list := []int{}
	for _, n := range roman {
		switch string(n) {
		case "I":
			list = append(list, 1)
		case "V":
			list = append(list, 5)
		case "X":
			list = append(list, 10)
		case "L":
			list = append(list, 50)
		case "C":
			list = append(list, 100)
		case "D":
			list = append(list, 500)
		case "M":
			list = append(list, 1000)
		}
	}

	if len(list) == 1 {
		arabic += list[0]
		return arabic
	}
	list = reverse2(list)
	if list[0] >= list[1] {
		arabic += list[0]
	} else {
		arabic += list[0]
	}

	fmt.Println(list)
	for i := 1; i < len(list); i++ {
		if list[i] >= list[i-1] {
			arabic += list[i]
		} else {
			arabic -= list[i]
		}
	}
	return arabic
}

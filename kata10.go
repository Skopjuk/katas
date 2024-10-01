package main

//https://www.codewars.com/kata/5659c6d896bc135c4c00021e/train/go
import (
	"sort"
	"strconv"
	"strings"
)

func NextSmaller(n int) int {
	strInt := strconv.Itoa(n)
	var listOfInt []int
	for _, i := range strInt {
		newi, _ := strconv.Atoi(string(i))
		listOfInt = append(listOfInt, newi)
	}
	sort.Ints(listOfInt)

	var newOfStr []string

	for _, i := range listOfInt {
		newOfStr = append(newOfStr, strconv.Itoa(i))
	}

	joinedStr := strings.Join(newOfStr, "")

	newInt, _ := strconv.Atoi(joinedStr)

	if n == newInt {
		return -1
	} else {
		return newInt
	}
}

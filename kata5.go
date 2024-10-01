package main

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	phnum := "("
	for n, i := range numbers {
		if n == 2 {
			phnum += strconv.Itoa(int(i)) + ") "
		} else if n == 5 {
			phnum += strconv.Itoa(int(i)) + "-"
		} else {
			phnum += strconv.Itoa(int(i))
		}
	}
	return phnum
}

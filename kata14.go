package main

import "fmt"

func Beeramid(bonus int, price float64) int {

	fmt.Println(bonus, price)
	numOfLayers := 0

	numOfCans := float64(bonus) / price

	a := 1

	for numOfCans > 0 {
		numOfCans -= float64(a * a)
		a += 1
		numOfLayers += 1
	}
	if numOfCans == 0 || numOfCans == 1 {
		return numOfLayers
	} else if bonus < 0 {
		return 0
	} else {
		return numOfLayers - 1
	}
}

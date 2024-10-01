package main

func MoveZeros(arr []int) []int {
	var list1 []int
	var list2 []int

	for _, i := range arr {
		if i != 0 {
			list1 = append(list1, i)
		} else {
			list2 = append(list2, i)
		}
	}
	return append(list1, list2...)
}

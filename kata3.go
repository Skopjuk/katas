package main

func Solution(str string) []string {
	var list []string
	var newStr string

	if len(str)%2 != 0 {
		newStr += str + "_"
	} else {
		newStr = str
	}

	for n, _ := range newStr {
		if n%2 == 0 {
			list = append(list, newStr[n:n+2])
		}
	}
	return list
}

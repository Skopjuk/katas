package main

func Rot13(msg string) string {
	newStr := ""
	for _, i := range msg {
		if string(i) == "m" {
			newStr += "z"
		} else if (i > 96 && i < 109) || (i > 64 && i < 78) {
			newStr += string(i + 13)
		} else if (i > 77 && i < 91) || (i > 109 && i < 123) {
			newStr += string(i - 13)
		} else {
			newStr += string(i)
		}
	}
	return newStr
}

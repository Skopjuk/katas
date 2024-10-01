package main

func Going(n int) float64 {

	//result := 0
	//for a := 1; a < n; n++{
	//	result += (math.f)
	//}
	return 0
}

func factorial(n int) float32 {
	final := n
	var result float32
	if final == 0 {
		return result
	} else {
		final -= 1
		result += factorial(n-1) * factorial(n-1)
	}

	return result

}

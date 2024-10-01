package map_array

type Apply interface {
	Apply(int) int
}

func MapArray(xs []int, operation Apply) []int {
	var new_array []int
	for _, n := range xs {
		new_array = append(new_array, operation.Apply(n))
	}

	return new_array
}

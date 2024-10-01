package main

func reverse(ls []uint64) []uint64 {
	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}
	return ls
}

func PartsSums(ls []uint64) []uint64 {
	var new_list []uint64
	var iter uint64
	iter = 0
	new_list = append(new_list, iter)

	ls = reverse(ls)

	for _, i := range ls {
		iter += i
		new_list = append(new_list, iter)
	}

	return reverse(new_list)
}

//func sum(ls []uint64) uint64 {
//	var res uint64
//	for _, i := range ls {
//		res += i
//	}
//	return res
//}

//func PartsSums(ls []uint64) []uint64 {
//	var new_list []uint64
//	for i := 0; i < len(ls); i++ {
//		new_list = append(new_list, sum(ls[i:len(ls)]))
//	}
//
//	new_list = append(new_list, 0)
//
//	return new_list
//}

package main

func FindUniq(arr []float32) float32 {
	var list1 []float32
	var list2 []float32
	for _, i := range arr {

		if len(list1) == 0 {
			list1 = append(list1, i)
		} else if len(list2) == 0 && i != list1[0] {
			list2 = append(list2, i)
		} else {
			if list1[0] == i {
				list1 = append(list1, i)
			} else {
				list2 = append(list2, i)
			}
		}
	}

	if len(list1) > len(list2) {
		return list2[0]
	} else {
		return list1[0]
	}
}

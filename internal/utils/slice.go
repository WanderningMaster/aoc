package utils

func Map(vs []string, f func(string) interface{}) []interface{} {
	vsm := make([]interface{}, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func MaxSlice(slice []int) (int, int) {
	max := 0
	maxIdx := 0

	for idx, x := range slice {
		if x > max {
			max = x
			maxIdx = idx
		}
	}

	return max, maxIdx
}


func LastMaxSlice(slice []int) (int, int) {
	max := 0
	maxIdx := 0

	for idx, x := range slice {
		if x >= max {
			max = x
			maxIdx = idx
		}
	}

	return max, maxIdx
}

func RemoveIntSlice(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

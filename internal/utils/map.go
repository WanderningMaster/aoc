package utils

func MapKeys(_map map[string]int) []string {
	keys := make([]string, len(_map))

	i := 0
	for k := range _map {
		keys[i] = k
		i++
	}

	return keys
}

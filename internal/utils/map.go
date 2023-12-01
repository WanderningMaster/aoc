package utils

func MapKeys[K comparable, V any](_map map[K]V) []K {
	keys := make([]K, len(_map))

	i := 0
	for k := range _map {
		keys[i] = k
		i++
	}

	return keys
}

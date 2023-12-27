package utils

func Remove[T comparable](list []T, index int) []T {
	list[index] = list[len(list)-1]
	return list[:len(list)-1]
}

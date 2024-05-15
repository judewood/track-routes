package utils

func ReverseCollection[T any](collection *[]T) *[]T {
	for i, j := 0, len(*collection)-1; i < j; i, j = i+1, j-1 {
		(*collection)[i], (*collection)[j] = (*collection)[j], (*collection)[i]
	}
	return collection
}
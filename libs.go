package main

func removeElementFromSlice[T comparable](slice *[]T, elementToRemove T) {
	for i, v := range *slice {
		if v == elementToRemove {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
		}
	}
}

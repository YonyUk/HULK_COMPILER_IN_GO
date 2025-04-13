package tools

import "errors"

// Counts the elementes that checks the given condition
func Count[T any](array []T, condition func(T) bool) int {
	counter := 0
	for _, item := range array {
		if condition(item) {
			counter++
		}
	}
	return counter
}

// Returns the elements that checks the given condition
func FiltBy[T any](array []T, filter func(T) bool) []T {
	result := []T{}
	for _, item := range array {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

// Returns the elements modifed by the given rule
func Map[T any, K any](array []T, modifier func(T) K) []K {
	result := []K{}
	for _, item := range array {
		result = append(result, modifier(item))
	}
	return result
}

// Returns the index of the element in case of be found
func IndexOf[T any](array []T, condition func(T) bool) (int, error) {
	for index, element := range array {
		if condition(element) {
			return index, nil
		}
	}
	return -1, errors.New("The element with that features is not in this secuence")
}

// Compare two arrays
func CompareArrays[T comparable](array1 []T, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := 0; i < len(array1); i++ {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

// Sort an array, the comparer function most return 1 if a > b, -1 if a < b and 0 if a == b
func Sort[T any](array []T, comparer func(a T, b T) int) []T {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if comparer(array[j], array[i]) == -1 {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
	return array
}

// Sort an array, the comparer function most return 1 if a > b, -1 if a < b and 0 if a == b
func MergeSort[T any](array []T, comparer func(a T, b T) int) []T {
	if len(array) == 1 {
		return array
	}
	result := merge(MergeSort(array[:int(len(array)/2)], comparer), MergeSort(array[int(len(array)/2):], comparer), comparer)
	return result
}

func merge[T any](array1 []T, array2 []T, comparer func(a T, b T) int) []T {
	result := make([]T, len(array1)+len(array2))
	index1 := 0
	index2 := 0
	r_index := 0
	for index1 < len(array1) || index2 < len(array2) {
		if index1 == len(array1) {
			result[r_index] = array2[index2]
			r_index++
			index2++
			continue
		}
		if index2 == len(array2) {
			result[r_index] = array1[index1]
			r_index++
			index1++
			continue
		}
		if comparer(array1[index1], array2[index2]) == -1 {
			result[r_index] = array1[index1]
			index1++
		} else {
			result[r_index] = array2[index2]
			index2++
		}
		r_index++
	}
	return result
}

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

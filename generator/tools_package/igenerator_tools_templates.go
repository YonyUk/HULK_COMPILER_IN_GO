package toolspackage

func IGeneratorCode() string {
	return `package tools

type IGenerator[T any] interface {
	// Return true if there is a next element
	Next() bool
	// Return the current element
	Current() T
}
`
}

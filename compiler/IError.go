package compiler

// Error interface
type IError interface {
	Line() int
	Column() int
	Message() string
	Type() ErrorType
}

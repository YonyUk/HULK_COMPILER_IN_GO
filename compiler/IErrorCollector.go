package compiler

// Error collector interface
type IErrorCollector interface {
	Errors() []IError
	AddError(error_ IError)
}

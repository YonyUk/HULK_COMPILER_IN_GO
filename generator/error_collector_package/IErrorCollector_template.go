package errorcollectorpackage

func IErrorCollectorCode() string {
	return `package compiler

// Error collector interface
type IErrorCollector interface {
	Errors() []IError
	AddError(error_ IError)
}
`
}

package errorcollectorpackage

func ErrorCollectorCode() string {
	return `package compiler

type ErrorCollector struct {
	errors_ []IError
}

func NewErrorCollector() *ErrorCollector {
	return &ErrorCollector{
		errors_: []IError{},
	}
}

func (ec *ErrorCollector) Errors() []IError {
	result := make([]IError, len(ec.errors_))
	copy(result, ec.errors_)
	return result
}

func (ec *ErrorCollector) AddError(error_ IError) {
	ec.errors_ = append(ec.errors_, error_)
}
`
}

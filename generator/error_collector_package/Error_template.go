package errorcollectorpackage

func ErrorCode() string {
	return `package compiler

type Error struct {
	msg    string
	line   int
	column int
	type_  ErrorType
}

func NewError(msg string, line int, column int, type_ ErrorType) *Error {
	return &Error{
		msg:    msg,
		line:   line,
		column: column,
		type_:  type_,
	}
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Line() int {
	return e.line
}

func (e *Error) Column() int {
	return e.column
}

func (e *Error) Type() ErrorType {
	return e.type_
}
`
}

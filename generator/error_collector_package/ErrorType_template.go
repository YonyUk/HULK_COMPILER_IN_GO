package errorcollectorpackage

func ErrorTypeCode() string {
	return `package compiler

type ErrorType int

const (
	Lexical ErrorType = iota
	Gramatical
	Semantic
)
`
}

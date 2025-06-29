package interpreterpackage

func IInterpreterCode() string {
	return `package interpreter

// Interpreter interface
type IInterpreter interface {
	Execute(code string)
}
`
}

package astpackage

func BaseASTCode() string {
	return `package ast

// AST base implementation
type BaseAST struct {
	Line   int
	Column int
	Symbol string
}
`
}

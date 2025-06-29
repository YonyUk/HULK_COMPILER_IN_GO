package astpackage

func IASTCode() string {
	return `package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

// AST interface
type IAST interface {
	Line() int
	Column() int
	Eval(context IContext, collector IErrorCollector) any
	Symbol() string
	UpdateSymbol(string)
}
`
}

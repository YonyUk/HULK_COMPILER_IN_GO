package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

// AST interface
type IAST interface {
	Line() int
	Column() int
	Eval(context IContext, collector IErrorCollector) interface{}
	Symbol() string
	UpdateSymbol(string)
}

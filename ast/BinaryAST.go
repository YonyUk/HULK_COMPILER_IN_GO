package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

// Ast binary implementation
type BinaryAST struct {
	BaseAST
	Left     IAST
	Right    IAST
	operator func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{}
}

func NewBinaryAST(symbol string, line int, column int, operator func(left IAST, right IAST, context IContext, collector IErrorCollector) interface{}) *BinaryAST {
	return &BinaryAST{
		Left:     nil,
		Right:    nil,
		operator: operator,
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
	}
}

func (a *BinaryAST) Line() int {
	return a.BaseAST.Line
}

func (a *BinaryAST) Column() int {
	return a.BaseAST.Column
}

func (a *BinaryAST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *BinaryAST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.operator(a.Left, a.Right, context, collector)
}

func (a *BinaryAST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

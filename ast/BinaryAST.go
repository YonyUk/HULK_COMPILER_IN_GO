package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

// Ast binary implementation
type BinaryAST struct {
	BaseAST
	left     IAST
	right    IAST
	operator func(left IAST, right IAST) interface{}
}

func NewBinaryAST(symbol string, line int, column int, operator func(left IAST, right IAST) interface{}) *BinaryAST {
	return &BinaryAST{
		left:     nil,
		right:    nil,
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
	return a.operator(a.left, a.right)
}

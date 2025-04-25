package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type InFrontOperatorAST struct {
	BaseAST
	Target   IAST
	operator func(target IAST, context IContext, collector IErrorCollector) interface{}
}

func NewInFrontOperatorAST(symbol string, line int, column int, operator func(target IAST, context IContext, collector IErrorCollector) interface{}) *InFrontOperatorAST {
	return &InFrontOperatorAST{
		Target:   nil,
		operator: operator,
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
	}
}

func (a *InFrontOperatorAST) Line() int {
	return a.BaseAST.Line
}

func (a *InFrontOperatorAST) Column() int {
	return a.BaseAST.Column
}

func (a *InFrontOperatorAST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *InFrontOperatorAST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.operator(a.Target, context, collector)
}

func (a *InFrontOperatorAST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type GarbageAST struct {
	BaseAST
	value interface{}
}

func NewGarbageAST(symbol string, line int, column int, value interface{}) *GarbageAST {
	return &GarbageAST{
		value: value,
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
	}
}

func (a *GarbageAST) Line() int {
	return a.BaseAST.Line
}

func (a *GarbageAST) Column() int {
	return a.BaseAST.Column
}

func (a *GarbageAST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.value
}

func (a *GarbageAST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *GarbageAST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

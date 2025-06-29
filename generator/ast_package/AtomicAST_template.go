package astpackage

func AtomicASTCode() string {
	return `package ast

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

// Ast atomic implementation
type AtomicAST struct {
	BaseAST
	value any
}

func NewAtomicAST(symbol string, line int, column int, value any) *AtomicAST {
	return &AtomicAST{
		value: value,
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
	}
}

func (a *AtomicAST) Line() int {
	return a.BaseAST.Line
}

func (a *AtomicAST) Column() int {
	return a.BaseAST.Column
}

func (a *AtomicAST) Eval(context IContext, collector IErrorCollector) any {
	return a.value
}

func (a *AtomicAST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *AtomicAST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}
`
}

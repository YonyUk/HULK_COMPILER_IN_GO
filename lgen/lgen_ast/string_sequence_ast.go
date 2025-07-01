package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type STRING_SEQUENCE_AST struct {
	BaseAST
	Items []string
}

func NewStringSequenceAST(symbol string, line int, column int) *STRING_SEQUENCE_AST {
	return &STRING_SEQUENCE_AST{
		BaseAST: BaseAST{
			Symbol: symbol,
			Line:   line,
			Column: column,
		},
		Items: []string{},
	}
}

func (seq *STRING_SEQUENCE_AST) Line() int {
	return seq.BaseAST.Line
}

func (seq *STRING_SEQUENCE_AST) Column() int {
	return seq.BaseAST.Column
}

func (seq *STRING_SEQUENCE_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return seq.Items
}

func (seq *STRING_SEQUENCE_AST) Symbol() string {
	return seq.BaseAST.Symbol
}

func (seq *STRING_SEQUENCE_AST) UpdateSymbol(symbol string) {
	seq.BaseAST.Symbol = symbol
}

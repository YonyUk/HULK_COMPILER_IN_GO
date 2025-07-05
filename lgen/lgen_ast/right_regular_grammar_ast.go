package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/grammar"
)

type RIGHT_REGULAR_GRAMMAR_AST struct {
	BaseAST
	RegularGrammar IGrammar
}

func NewRightRegularGrammar(symbol string, line int, column int) *RIGHT_REGULAR_GRAMMAR_AST {
	return &RIGHT_REGULAR_GRAMMAR_AST{
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
		RegularGrammar: nil,
	}
}

func (a *RIGHT_REGULAR_GRAMMAR_AST) Line() int {
	return a.BaseAST.Line
}

func (a *RIGHT_REGULAR_GRAMMAR_AST) Column() int {
	return a.BaseAST.Column
}

func (a *RIGHT_REGULAR_GRAMMAR_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.RegularGrammar
}

func (a *RIGHT_REGULAR_GRAMMAR_AST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *RIGHT_REGULAR_GRAMMAR_AST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

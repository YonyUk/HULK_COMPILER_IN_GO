package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/grammar"
)

type GrammarProduction struct {
	Head       IGrammarSymbol
	Production []IGrammarSymbol
}

// Ast atomic implementation
type RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST struct {
	BaseAST
	Productions []GrammarProduction
}

func NewRightRegularGrammarProductionSequenceAST(symbol string, line int, column int) *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST {
	return &RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST{
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
		Productions: []GrammarProduction{},
	}
}

func (a *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST) Line() int {
	return a.BaseAST.Line
}

func (a *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST) Column() int {
	return a.BaseAST.Column
}

func (a *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.Productions
}

func (a *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE_AST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

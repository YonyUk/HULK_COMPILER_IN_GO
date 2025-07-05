package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/grammar"
)

// Ast atomic implementation
type TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST struct {
	BaseAST
	TokenGrammar IGrammar
}

func NewTokenDeclarationDefinedByGrammarAST(symbol string, line int, column int, value interface{}) *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST {
	return &TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST{
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
		TokenGrammar: nil,
	}
}

func (a *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST) Line() int {
	return a.BaseAST.Line
}

func (a *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST) Column() int {
	return a.BaseAST.Column
}

func (a *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.TokenGrammar
}

func (a *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *TOKEN_DECLARATION_DEFINED_BY_GRAMMAR_AST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

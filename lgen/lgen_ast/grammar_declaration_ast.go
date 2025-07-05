package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
	. "hulk.com/app/grammar"
)

// Ast atomic implementation
type GRAMMAR_DECLARATION_AST struct {
	BaseAST
	GrammarValue IGrammar
	GrammarName  string
}

func NewGrammarDeclarationAST(symbol string, line int, column int) *GRAMMAR_DECLARATION_AST {
	return &GRAMMAR_DECLARATION_AST{
		BaseAST: BaseAST{
			Line:   line,
			Column: column,
			Symbol: symbol,
		},
		GrammarValue: nil,
		GrammarName:  "",
	}
}

func (a *GRAMMAR_DECLARATION_AST) Line() int {
	return a.BaseAST.Line
}

func (a *GRAMMAR_DECLARATION_AST) Column() int {
	return a.BaseAST.Column
}

func (a *GRAMMAR_DECLARATION_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return a.GrammarValue
}

func (a *GRAMMAR_DECLARATION_AST) Symbol() string {
	return a.BaseAST.Symbol
}

func (a *GRAMMAR_DECLARATION_AST) UpdateSymbol(symbol string) {
	a.BaseAST.Symbol = symbol
}

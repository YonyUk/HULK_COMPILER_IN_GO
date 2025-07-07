package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type TOKEN_DECLARATION_DEFINED_BY_LIST_AST struct {
	BaseAST
	Tokens    []string
	TokenName string
}

func NewTokenDeclarationDefinedByListAST(symbol string, line int, column int) *TOKEN_DECLARATION_DEFINED_BY_LIST_AST {
	return &TOKEN_DECLARATION_DEFINED_BY_LIST_AST{
		BaseAST: BaseAST{
			Symbol: symbol,
			Line:   line,
			Column: column,
		},
		Tokens:    []string{},
		TokenName: "",
	}
}

func (td *TOKEN_DECLARATION_DEFINED_BY_LIST_AST) Line() int {
	return td.BaseAST.Line
}

func (td *TOKEN_DECLARATION_DEFINED_BY_LIST_AST) Column() int {
	return td.BaseAST.Column
}

func (td *TOKEN_DECLARATION_DEFINED_BY_LIST_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return td.Tokens
}

func (td *TOKEN_DECLARATION_DEFINED_BY_LIST_AST) Symbol() string {
	return td.BaseAST.Symbol
}

func (td *TOKEN_DECLARATION_DEFINED_BY_LIST_AST) UpdateSymbol(symbol string) {
	td.BaseAST.Symbol = symbol
}

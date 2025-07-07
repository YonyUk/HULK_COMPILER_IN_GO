package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type TOKEN_DECLARATION_AST struct {
	BaseAST
	TokenName string
}

func NewTokenDeclarationAST(symbol string, line int, column int) *TOKEN_DECLARATION_AST {
	return &TOKEN_DECLARATION_AST{
		BaseAST: BaseAST{
			Symbol: symbol,
			Line:   line,
			Column: column,
		},
		TokenName: "",
	}
}

func (td *TOKEN_DECLARATION_AST) Line() int {
	return td.BaseAST.Line
}

func (td *TOKEN_DECLARATION_AST) Column() int {
	return td.BaseAST.Column
}

func (td *TOKEN_DECLARATION_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return td.TokenName
}

func (td *TOKEN_DECLARATION_AST) Symbol() string {
	return td.BaseAST.Symbol
}

func (td *TOKEN_DECLARATION_AST) UpdateSymbol(symbol string) {
	td.BaseAST.Symbol = symbol
}

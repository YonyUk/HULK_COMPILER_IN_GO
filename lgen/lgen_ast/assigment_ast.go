package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type ASSIGMENT_AST struct {
	BaseAST
	Left  IAST
	Right IAST
}

func NewAssigmentAst(symbol string, line int, column int) *ASSIGMENT_AST {
	return &ASSIGMENT_AST{
		BaseAST: BaseAST{
			Symbol: symbol,
			Line:   line,
			Column: column,
		},
		Left:  nil,
		Right: nil,
	}
}

func (assigment *ASSIGMENT_AST) Line() int {
	return assigment.BaseAST.Line
}

func (assigment *ASSIGMENT_AST) Column() int {
	return assigment.BaseAST.Column
}

func (assigment *ASSIGMENT_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	// TODO: realizar la assignacion
	return nil
}

func (assigment *ASSIGMENT_AST) Symbol() string {
	return assigment.BaseAST.Symbol
}

func (assigment *ASSIGMENT_AST) UpdateSymbol(symbol string) {
	assigment.BaseAST.Symbol = symbol
}

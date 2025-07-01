package lgenast

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/context"
)

type VARIABLE_AST struct {
	BaseAST
	VariableName string
	Value        interface{}
}

func NewVariableAST(symbol string, line int, column int, name string) *VARIABLE_AST {
	return &VARIABLE_AST{
		BaseAST: BaseAST{
			Symbol: symbol,
			Line:   line,
			Column: column,
		},
		VariableName: name,
		Value:        nil,
	}
}

func (variable *VARIABLE_AST) Line() int {
	return variable.BaseAST.Line
}

func (variable *VARIABLE_AST) Column() int {
	return variable.BaseAST.Column
}

func (variable *VARIABLE_AST) Eval(context IContext, collector IErrorCollector) interface{} {
	return variable.Value
}

func (variable *VARIABLE_AST) Symbol() string {
	return variable.BaseAST.Symbol
}

func (variable *VARIABLE_AST) UpdateSymbol(symbol string) {
	variable.BaseAST.Symbol = symbol
}

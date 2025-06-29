package parser

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/tokens"
)

// Parser interface
type IParser interface {
	// Parse the current sequence into the parser with the new token
	Parse(token IToken, collector IErrorCollector)
	// Return the ast result from parse the sequence
	AST() IAST
	// Sets the reductor method for one production
	SetReduction(reduction_id string, reductor func(asts []IAST, new_symbol string) IAST)
	// The symbol to mark the end of the stack
	EndMarker() string
	// Gets the action table
	ActionTable() map[string]map[string]ActionStruct
	// Gets the Reduce table
	ReduceTable() map[string]map[string]ReduceStruct
}

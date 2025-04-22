package parser

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/automaton"
	. "hulk.com/app/compiler"
	. "hulk.com/app/tokens"
)

// Parser interface
type IParser interface {
	// The parser is an automaton too
	IAutomaton[IToken]
	// Parse the current sequence into the parser with the new token
	Parse(token IToken, collector IErrorCollector)
	// Return the ast result from parse the sequence
	AST() IAST
}

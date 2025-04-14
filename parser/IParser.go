package parser

import (
	. "hulk.com/app/ast"
	. "hulk.com/app/compiler"
	. "hulk.com/app/tokens"
)

// Parser interface
type IParser interface {
	// Push a new token into the sequence to parse
	PushToken(token IToken)
	// Parse the current sequence into the parser
	Parse(collector IErrorCollector)
	// Return the ast result from parse the sequence
	AST() (IAST, error)
}

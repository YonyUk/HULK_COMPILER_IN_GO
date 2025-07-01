package lgenast

import (
	. "hulk.com/app/ast"
	// . "hulk.com/app/compiler"
	// . "hulk.com/app/context"
	. "hulk.com/app/grammar"
)

type RIGHT_REGULAR_GRAMMAR_AST struct {
	BaseAST
	RegularGrammar IGrammar
}

// func NewRightRegularGrammar()

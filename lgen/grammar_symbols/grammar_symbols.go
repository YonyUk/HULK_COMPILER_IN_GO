package grammarsymbols

import (
	. "hulk.com/app/grammar"
)

// Non terminals
var TOKEN IGrammarSymbol
var STRING_SEQUENCE IGrammarSymbol
var STRING_LIST IGrammarSymbol
var TOKEN_DECLARATION IGrammarSymbol

// Terminals
var TokenKeyword IGrammarSymbol
var Text IGrammarSymbol
var Comma IGrammarSymbol
var DotComma IGrammarSymbol
var OpenCorchet IGrammarSymbol
var ClosedCorchet IGrammarSymbol
var Equal IGrammarSymbol
var Variable IGrammarSymbol
var EndSymbol IGrammarSymbol

func init() {
	TOKEN = NewGrammarSymbol("TOKEN", NonTerminal, false)
	TOKEN_DECLARATION = NewGrammarSymbol("TOKEN_DECLARATION", NonTerminal, false)
	STRING_SEQUENCE = NewGrammarSymbol("STRING_SEQUENCE", NonTerminal, false)
	STRING_LIST = NewGrammarSymbol("STRING_LIST", NonTerminal, false)

	Variable = NewGrammarSymbol("variable", Terminal, false)
	Equal = NewGrammarSymbol("=", Terminal, false)
	TokenKeyword = NewGrammarSymbol("token", Terminal, false)
	Text = NewGrammarSymbol("Text", Terminal, false)
	Comma = NewGrammarSymbol(",", Terminal, false)
	DotComma = NewGrammarSymbol(";", Terminal, false)
	OpenCorchet = NewGrammarSymbol("[", Terminal, false)
	ClosedCorchet = NewGrammarSymbol("]", Terminal, false)
	EndSymbol = NewGrammarSymbol("$", Terminal, false)
}

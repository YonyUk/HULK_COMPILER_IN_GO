package grammarsymbols

import (
	. "hulk.com/app/grammar"
)

// Non terminals
var TOKEN IGrammarSymbol
var STRING_SEQUENCE IGrammarSymbol
var STRING_LIST IGrammarSymbol
var TOKEN_DECLARATION IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR IGrammarSymbol
var GRAMMAR_PRODUCTION IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR_PRODUCTION IGrammarSymbol
var GRAMMAR_PRODUCTION_SEQUENCE IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE IGrammarSymbol

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
var GrammarNonTerminal IGrammarSymbol
var GrammarTerminal IGrammarSymbol
var GrammarProductionArrow IGrammarSymbol
var GrammarProductionConcatenation IGrammarSymbol
var LessThan IGrammarSymbol
var GreaterThan IGrammarSymbol

func init() {
	TOKEN = NewGrammarSymbol("TOKEN", NonTerminal, false)
	TOKEN_DECLARATION = NewGrammarSymbol("TOKEN_DECLARATION", NonTerminal, false)
	STRING_SEQUENCE = NewGrammarSymbol("STRING_SEQUENCE", NonTerminal, false)
	STRING_LIST = NewGrammarSymbol("STRING_LIST", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR", NonTerminal, false)
	GRAMMAR_PRODUCTION = NewGrammarSymbol("GRAMMAR_PRODUCTION", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR_PRODUCTION = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR_PRODUCTION", NonTerminal, false)
	GRAMMAR_PRODUCTION_SEQUENCE = NewGrammarSymbol("GRAMMAR_PRODUCTION_SEQUENCE", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR_PRODUCTION", NonTerminal, false)

	Variable = NewGrammarSymbol("variable", Terminal, false)
	Equal = NewGrammarSymbol("=", Terminal, false)
	TokenKeyword = NewGrammarSymbol("token", Terminal, false)
	Text = NewGrammarSymbol("Text", Terminal, false)
	Comma = NewGrammarSymbol(",", Terminal, false)
	DotComma = NewGrammarSymbol(";", Terminal, false)
	OpenCorchet = NewGrammarSymbol("[", Terminal, false)
	ClosedCorchet = NewGrammarSymbol("]", Terminal, false)
	EndSymbol = NewGrammarSymbol("$", Terminal, false)
	GrammarNonTerminal = NewGrammarSymbol("NonTerminal", Terminal, false)
	GrammarTerminal = NewGrammarSymbol("Terminal", Terminal, false)
	GrammarProductionArrow = NewGrammarSymbol("--->", Terminal, false)
	GrammarProductionConcatenation = NewGrammarSymbol("|", Terminal, false)
	LessThan = NewGrammarSymbol("<", Terminal, false)
	GreaterThan = NewGrammarSymbol(">", Terminal, false)
}

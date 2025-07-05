package grammarsymbols

import (
	. "hulk.com/app/grammar"
)

// Non terminals
var TOKEN IGrammarSymbol
var TOKEN_DEFINED_BY_LIST IGrammarSymbol
var TOKEN_DEFINED_BY_GRAMMAR IGrammarSymbol
var STRING_SEQUENCE IGrammarSymbol
var STRING_LIST IGrammarSymbol
var TOKEN_DECLARATION IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR IGrammarSymbol
var GRAMMAR_PRODUCTION IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR_PRODUCTION IGrammarSymbol
var GRAMMAR_PRODUCTION_SEQUENCE IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE IGrammarSymbol
var RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE IGrammarSymbol
var GRAMMAR_DECLARATION IGrammarSymbol

// Terminals
var TokenKeyword IGrammarSymbol
var GrammarKeyword IGrammarSymbol
var Text IGrammarSymbol
var Comma IGrammarSymbol
var DotComma IGrammarSymbol
var OpenCorchet IGrammarSymbol
var OpenParent IGrammarSymbol
var ClosedParent IGrammarSymbol
var ClosedCorchet IGrammarSymbol
var Equal IGrammarSymbol
var Variable IGrammarSymbol
var EndSymbol IGrammarSymbol
var GrammarProductionArrow IGrammarSymbol
var Or IGrammarSymbol
var LessThan IGrammarSymbol
var GreaterThan IGrammarSymbol
var Epsilon IGrammarSymbol

func init() {
	TOKEN = NewGrammarSymbol("TOKEN", NonTerminal, false)
	TOKEN_DEFINED_BY_LIST = NewGrammarSymbol("TOKEN_DEFINED_BY_LIST", NonTerminal, false)
	TOKEN_DEFINED_BY_GRAMMAR = NewGrammarSymbol("TOKEN_DEFINED_BY_GRAMMAR", NonTerminal, false)
	TOKEN_DECLARATION = NewGrammarSymbol("TOKEN_DECLARATION", NonTerminal, false)
	STRING_SEQUENCE = NewGrammarSymbol("STRING_SEQUENCE", NonTerminal, false)
	STRING_LIST = NewGrammarSymbol("STRING_LIST", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR", NonTerminal, false)
	GRAMMAR_PRODUCTION = NewGrammarSymbol("GRAMMAR_PRODUCTION", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR_PRODUCTION = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR_PRODUCTION", NonTerminal, false)
	GRAMMAR_PRODUCTION_SEQUENCE = NewGrammarSymbol("GRAMMAR_PRODUCTION_SEQUENCE", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR_PRODUCTION_SEQUENCE = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR_PRODUCTION", NonTerminal, false)
	GRAMMAR_DECLARATION = NewGrammarSymbol("GRAMMAR_DECLARATION", NonTerminal, false)
	RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE = NewGrammarSymbol("RIGHT_REGULAR_GRAMMAR_DERIVATION_SEQUENCE", NonTerminal, false)

	Variable = NewGrammarSymbol("variable", Terminal, false)
	Equal = NewGrammarSymbol("=", Terminal, false)
	TokenKeyword = NewGrammarSymbol("token", Terminal, false)
	GrammarKeyword = NewGrammarSymbol("grammar", Terminal, false)
	Text = NewGrammarSymbol("Text", Terminal, false)
	Comma = NewGrammarSymbol(",", Terminal, false)
	DotComma = NewGrammarSymbol(";", Terminal, false)
	OpenCorchet = NewGrammarSymbol("[", Terminal, false)
	ClosedCorchet = NewGrammarSymbol("]", Terminal, false)
	EndSymbol = NewGrammarSymbol("$", Terminal, false)
	GrammarProductionArrow = NewGrammarSymbol("--->", Terminal, false)
	Or = NewGrammarSymbol("|", Terminal, false)
	LessThan = NewGrammarSymbol("<", Terminal, false)
	GreaterThan = NewGrammarSymbol(">", Terminal, false)
	OpenParent = NewGrammarSymbol("(", Terminal, false)
	ClosedParent = NewGrammarSymbol(")", Terminal, false)
	Epsilon = NewGrammarSymbol("epsilon", Terminal, false)
}

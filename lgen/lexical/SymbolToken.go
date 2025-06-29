package lexical

import (
	. "hulk.com/app/grammar"
)

var SymbolTokenGrammar IGrammar

func init() {
	symbols := []string{
		"[",
		"]",
		";",
		",",
	}
	SymbolTokenGrammar = GetWordsGrammar(symbols)
}

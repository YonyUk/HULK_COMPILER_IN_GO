package hulk

import (
	. "hulk.com/app/grammar"
)

var SymbolGrammar IGrammar

func init() {
	symbols := []string{
		"self",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"=>",
		";",
		".",
	}
	SymbolGrammar = GetWordsGrammar(symbols)
}

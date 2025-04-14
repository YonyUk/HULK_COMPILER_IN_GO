package hulk

import (
	. "hulk.com/app/grammar"
)

var StringGrammar IGrammar

func init() {
	start_symbol := NewGrammarSymbol("start_symbol", NonTerminal, false)
	tail := NewGrammarSymbol("tail", NonTerminal, false)
	double_quote_symbol := NewGrammarSymbol("\"", Terminal, false)
	symbols := []IGrammarSymbol{}
	for r := rune(0); r <= rune(256); r++ {
		symbols = append(symbols, NewGrammarSymbol(string(r), Terminal, false))
	}
	StringGrammar = NewGrammar(start_symbol)
	StringGrammar.AddProduction(start_symbol, []IGrammarSymbol{double_quote_symbol, tail})

	for _, symbol := range symbols {
		StringGrammar.AddProduction(tail, []IGrammarSymbol{symbol, tail})
	}

	StringGrammar.AddProduction(tail, []IGrammarSymbol{double_quote_symbol})
}

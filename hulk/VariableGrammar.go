package hulk

import (
	. "hulk.com/app/grammar"
)

var VariableGrammar IGrammar

func init() {
	start_symbol := NewGrammarSymbol("start_symbol", NonTerminal, false)
	tail := NewGrammarSymbol("tail", NonTerminal, false)
	epsilon := NewGrammarSymbol("e", Terminal, true)
	down_line := NewGrammarSymbol("_", Terminal, false)
	digits := []IGrammarSymbol{}
	letters := []IGrammarSymbol{}

	VariableGrammar = NewGrammar(start_symbol)

	for r := rune(65); r < rune(90); r++ {
		letters = append(letters, NewGrammarSymbol(string(r), Terminal, false))
	}

	for r := rune(97); r < rune(122); r++ {
		letters = append(letters, NewGrammarSymbol(string(r), Terminal, false))
	}

	for r := rune(48); r < rune(58); r++ {
		digits = append(digits, NewGrammarSymbol(string(r), Terminal, false))
	}

	VariableGrammar.AddProduction(start_symbol, []IGrammarSymbol{down_line, tail})
	for _, symbol := range letters {
		VariableGrammar.AddProduction(start_symbol, []IGrammarSymbol{symbol, tail})
	}

	VariableGrammar.AddProduction(tail, []IGrammarSymbol{down_line, tail})
	for _, symbol := range letters {
		VariableGrammar.AddProduction(tail, []IGrammarSymbol{symbol, tail})
	}
	for _, symbol := range digits {
		VariableGrammar.AddProduction(tail, []IGrammarSymbol{symbol, tail})
	}
	VariableGrammar.AddProduction(tail, []IGrammarSymbol{epsilon})
}

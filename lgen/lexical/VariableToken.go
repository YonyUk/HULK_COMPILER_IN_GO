package lexical

import (
	. "hulk.com/app/grammar"
)

var VariableTokenGrammar IGrammar

func init() {
	letters := []string{}
	for r := rune(65); r < rune(91); r++ {
		letters = append(letters, string(r))
	}
	for r := rune(97); r < rune(123); r++ {
		letters = append(letters, string(r))
	}
	for r := rune(48); r < rune(58); r++ {
		letters = append(letters, string(r))
	}

	Variable := NewGrammarSymbol("String", NonTerminal, false)
	Text := NewGrammarSymbol("Text", NonTerminal, false)
	epsilon := NewGrammarSymbol("epsilon", Terminal, true)

	VariableTokenGrammar = NewGrammar(Variable)
	for _, letter := range letters {
		VariableTokenGrammar.AddProduction(Variable, []IGrammarSymbol{NewGrammarSymbol(letter, Terminal, false), Text})
		VariableTokenGrammar.AddProduction(Text, []IGrammarSymbol{NewGrammarSymbol(letter, Terminal, false), Text})
	}
	VariableTokenGrammar.AddProduction(Variable, []IGrammarSymbol{epsilon})
}

package lexical

import (
	. "hulk.com/app/grammar"
)

var LiteralStringTokenGrammar IGrammar

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
	letters = append(letters, " ")
	String := NewGrammarSymbol("String", NonTerminal, false)
	Text := NewGrammarSymbol("Text", NonTerminal, false)
	double_quote := NewGrammarSymbol(`"`, Terminal, false)

	LiteralStringTokenGrammar = NewGrammar(String)
	LiteralStringTokenGrammar.AddProduction(String, []IGrammarSymbol{double_quote, Text})
	for _, letter := range letters {
		LiteralStringTokenGrammar.AddProduction(Text, []IGrammarSymbol{NewGrammarSymbol(letter, Terminal, false), Text})
	}
	LiteralStringTokenGrammar.AddProduction(Text, []IGrammarSymbol{double_quote})
}

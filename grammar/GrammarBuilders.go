package grammar

import (
	"strconv"
)

// Adds a new word to the given grammar
func AddWordToGrammar(g IGrammar, word string) IGrammar {
	result := g
	var head IGrammarSymbol = g.StartSymbol()
	var tail IGrammarSymbol
	for i := 0; i < len(word); i++ {
		tail = NewGrammarSymbol(word+"_next_"+string(word[i])+"_"+strconv.Itoa(i), NonTerminal, false)
		if i < len(word)-1 {
			g.AddProduction(head, []IGrammarSymbol{NewGrammarSymbol(string(word[i]), Terminal, false), tail})
		} else {
			g.AddProduction(head, []IGrammarSymbol{NewGrammarSymbol(string(word[i]), Terminal, false)})
		}
		head = tail
	}
	return result
}

// Gets the grammar for the given set of words
func GetWordsGrammar(words []string) IGrammar {
	start_symbol := NewGrammarSymbol("start_symbol", NonTerminal, false)
	var result IGrammar = NewGrammar(start_symbol)
	for _, word := range words {
		result = AddWordToGrammar(result, word)
	}
	return result
}

package main

import (
	"fmt"

	. "hulk.com/app/grammar"
	. "hulk.com/app/regex"
)

func main() {

	Number := NewGrammarSymbol("Number", NonTerminal, false)
	tail := NewGrammarSymbol("tail", NonTerminal, false)
	d0 := NewGrammarSymbol("0", Terminal, false)
	d1 := NewGrammarSymbol("1", Terminal, false)
	d2 := NewGrammarSymbol("2", Terminal, false)
	d3 := NewGrammarSymbol("3", Terminal, false)
	d4 := NewGrammarSymbol("4", Terminal, false)
	d5 := NewGrammarSymbol("5", Terminal, false)
	d6 := NewGrammarSymbol("6", Terminal, false)
	d7 := NewGrammarSymbol("7", Terminal, false)
	d8 := NewGrammarSymbol("8", Terminal, false)
	d9 := NewGrammarSymbol("9", Terminal, false)
	epsilon := NewGrammarSymbol("epsilon", Terminal, true)

	NumberGrammar := NewGrammar(Number)
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d1, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d2, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d3, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d4, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d5, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d6, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d7, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d8, tail})
	NumberGrammar.AddProduction(Number, []IGrammarSymbol{d9, tail})

	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d0, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d1, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d2, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d3, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d4, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d5, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d6, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d7, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d8, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{d9, tail})
	NumberGrammar.AddProduction(tail, []IGrammarSymbol{epsilon})

	NumberRegex := NewRegexEngine()
	aut, err := NumberRegex.Regex(NumberGrammar)
	if err != nil {
		fmt.Println(err)
	}

	chain := "680129832"

	fmt.Println(chain[1:len(chain)])

	for _, char := range chain {
		fmt.Println(aut.Walk(char))
	}

	fmt.Println()
	fmt.Println(aut.CurrentState().IsAccepting())

}

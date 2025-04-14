package main

import (
	"fmt"

	. "hulk.com/app/filesystem"
	. "hulk.com/app/grammar"
	. "hulk.com/app/hulk"
	. "hulk.com/app/tools"
)

func main() {
	interpreter := NewHulkInterpreter()
	reader, _ := NewFileReader("code.hulk")
	code, _ := reader.ReadFile()
	interpreter.Execute(code)

	E := NewGrammarSymbol("E", NonTerminal, false)
	T := NewGrammarSymbol("T", NonTerminal, false)
	F := NewGrammarSymbol("F", NonTerminal, false)
	X := NewGrammarSymbol("X", NonTerminal, false)
	Y := NewGrammarSymbol("Y", NonTerminal, false)

	plus := NewGrammarSymbol("+", Terminal, false)
	mul := NewGrammarSymbol("*", Terminal, false)
	epsilon := NewGrammarSymbol("epsilon", Terminal, true)
	RP := NewGrammarSymbol(")", Terminal, false)
	LP := NewGrammarSymbol("(", Terminal, false)
	id := NewGrammarSymbol("id", Terminal, false)

	G := NewGrammar(E)
	G.AddProduction(E, []IGrammarSymbol{T, X})
	G.AddProduction(X, []IGrammarSymbol{plus, T, X})
	G.AddProduction(X, []IGrammarSymbol{epsilon})
	G.AddProduction(T, []IGrammarSymbol{F, Y})
	G.AddProduction(Y, []IGrammarSymbol{mul, F, Y})
	G.AddProduction(Y, []IGrammarSymbol{epsilon})
	G.AddProduction(F, []IGrammarSymbol{LP, E, RP})
	G.AddProduction(F, []IGrammarSymbol{id})

	G.MakeFirstsAndFollows(NewGrammarSymbol("$", Terminal, false))

	for _, nt := range G.NonTerminals() {
		fmt.Println(nt.Symbol(), "--->", Map(G.FOLLOW(nt), func(symbol IGrammarSymbol) string { return symbol.Symbol() }))
	}
}

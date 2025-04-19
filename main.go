package main

import (
	"fmt"

	. "hulk.com/app/filesystem"
	. "hulk.com/app/grammar"
	. "hulk.com/app/hulk"
	. "hulk.com/app/hulk/sintax"
	. "hulk.com/app/parser"
	. "hulk.com/app/tools"
)

func main() {
	interpreter := NewHulkInterpreter()
	reader, _ := NewFileReader("code.hulk")
	code, _ := reader.ReadFile()
	interpreter.Execute(code)

	G := AugmentGrammar(HulkGrammar)

	G.MakeFirstsAndFollows(NewGrammarSymbol("$", Terminal, false))

	itemsLR0collections := make(map[string][]IItemLR0)
	for _, nt := range G.NonTerminals() {
		itemsLR0collections[nt.Symbol()] = []IItemLR0{}
		productions := G.GetProductions(nt)
		for _, production := range productions {
			collection := GetLR0Collection(nt, production)
			for _, item := range collection {
				itemsLR0collections[nt.Symbol()] = append(itemsLR0collections[nt.Symbol()], item)
			}
		}
	}
	clousures := make(map[string][]IItemLR0)
	for key, items := range itemsLR0collections {
		clousures[key] = ItemLR0Clousure(items, G)
	}
	for _, clousure := range clousures {
		for _, item := range clousure {
			left := Map(item.LeftTail(), func(s IGrammarSymbol) string { return s.Symbol() })
			right := Map(item.RightTail(), func(s IGrammarSymbol) string { return s.Symbol() })
			fmt.Println(item.Head().Symbol(), "---->", left, ".", right)
		}
		fmt.Println()
		fmt.Println()
	}
}

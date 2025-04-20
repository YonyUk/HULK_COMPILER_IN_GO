package main

import (
	. "hulk.com/app/filesystem"
	. "hulk.com/app/grammar"
	. "hulk.com/app/hulk"

	. "hulk.com/app/hulk/sintax"
	. "hulk.com/app/parser"
	// . "hulk.com/app/tools"
)

func main() {
	interpreter := NewHulkInterpreter()
	reader, _ := NewFileReader("code.hulk")
	code, _ := reader.ReadFile()
	interpreter.Execute(code)

	G := AugmentGrammar(HulkGrammar)

	G.MakeFirstsAndFollows(NewGrammarSymbol("$", Terminal, false))

	canonicalCollection := GetCanonicalLR0Collection(G)
	for _, collection := range canonicalCollection {
		ShowCollection(collection)
	}
}

package main

import (
	"fmt"

	. "hulk.com/app/hulk"
	. "hulk.com/app/lexer"
	. "hulk.com/app/tokens"
	// . "hulk.com/app/tools"
)

func main() {
	chain := "-1e-583 0 0.000001 12.3 -12.3 1 -1 +1 +1e+500 +1e-500 -1e-500 +12.3 -0 +0"

	lexer := NewLexer()
	lexer.AddTokenExpression(NumberToken, 0, NumberGrammar)
	lexer.LoadCode(chain)
	for lexer.Next() {
		fmt.Println(lexer.Current())
	}
}

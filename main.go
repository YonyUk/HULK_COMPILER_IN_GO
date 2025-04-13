package main

import (
	"fmt"

	. "hulk.com/app/hulk"
	. "hulk.com/app/lexer"
	. "hulk.com/app/tokens"
	// . "hulk.com/app/tools"
)

func main() {

	// for r := rune(0); r < rune(256); r++ {
	// 	fmt.Println(string(r), r)
	// }

	chain := "\"hello world\" \"here we go\" \n 34e+200 nada_nuevo_01"

	fmt.Println(chain)

	lexer := NewLexer()
	lexer.AddTokenExpression(NumberToken, 0, NumberGrammar)
	lexer.AddTokenExpression(StringToken, 1, StringGrammar)
	lexer.AddTokenExpression(VariableToken, 2, VariableGrammar)
	lexer.LoadCode(chain)
	for lexer.Next() {
		fmt.Println(lexer.Current())
	}
}

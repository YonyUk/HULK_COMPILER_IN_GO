package main

import (
	"fmt"

	. "hulk.com/app/hulk"
	. "hulk.com/app/lexer"
	. "hulk.com/app/tokens"
	// . "hulk.com/app/grammar"
	// . "hulk.com/app/regex"
	// . "hulk.com/app/tools"
)

func main() {

	chain := "\"hello world\" \"here we go\" \n34e+200 nada_nuevo_01 function"

	lexer := NewLexer()
	lexer.AddTokenExpression(KeywordToken, 0, KeywordsGrammar)
	lexer.AddTokenExpression(NumberToken, 1, NumberGrammar)
	lexer.AddTokenExpression(StringToken, 2, StringGrammar)
	lexer.AddTokenExpression(VariableToken, 3, VariableGrammar)
	lexer.LoadCode(chain)
	for lexer.Next() {
		fmt.Println(lexer.Current())
	}

}

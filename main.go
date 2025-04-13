package main

import (
	// . "hulk.com/app/grammar"

	"fmt"

	. "hulk.com/app/hulk"

	. "hulk.com/app/lexer"
	// . "hulk.com/app/regex"
	. "hulk.com/app/tokens"
	// . "hulk.com/app/tools"
)

func main() {

	chain := "1 + 0.5 - 9e+300 * -1e-3 / 12"

	map_type_to_str := make(map[TokenType]string)

	map_type_to_str[KeywordToken] = "Keyword"
	map_type_to_str[BooleanToken] = "Boolean"
	map_type_to_str[VariableToken] = "Variable"
	map_type_to_str[NumberToken] = "Number"
	map_type_to_str[StringToken] = "String"
	map_type_to_str[SymbolToken] = "Symbol"
	map_type_to_str[OperatorToken] = "Operator"

	// KeywordsGrammar := GetWordsGrammar([]string{"print", "function", "rand"})

	lexer := NewLexer()
	lexer.AddTokenExpression(KeywordToken, 0, KeywordsGrammar)
	lexer.AddTokenExpression(BooleanToken, 1, BooleanGrammar)
	lexer.AddTokenExpression(VariableToken, 2, VariableGrammar)
	lexer.AddTokenExpression(OperatorToken, 3, OperatorGrammar)
	lexer.AddTokenExpression(SymbolToken, 4, SymbolGrammar)
	lexer.AddTokenExpression(NumberToken, 5, NumberGrammar)
	lexer.AddTokenExpression(StringToken, 6, StringGrammar)
	lexer.LoadCode(chain)
	for lexer.Next() {
		fmt.Println(lexer.Current(), map_type_to_str[lexer.Current().Type()])
	}
}

package main

import (
	"fmt"

	. "hulk.com/app/compiler"
	. "hulk.com/app/filesystem"
	. "hulk.com/app/lgen"
	. "hulk.com/app/tokens"
)

func main() {
	lgen_lexer := LGENLexer
	lgen_parser := LGENParser
	filereader, _ := NewFileReader("code.lgn")
	code, _ := filereader.ReadFile()

	collector := NewErrorCollector()

	lgen_lexer.LoadCode(code)
	for lgen_lexer.Next() {
		lgen_parser.Parse(lgen_lexer.Current(), collector)
		// fmt.Println(lgen_lexer.Current(), Type(lgen_lexer.Current().Type()))
	}
	lgen_parser.Parse(NewToken(lgen_lexer.Current().Line(), lgen_lexer.Current().Column()+1, lgen_parser.EndMarker(), EndToken), collector)

	for _, err := range collector.Errors() {
		fmt.Println("Error ", err.Type(), ": ", err.Message(), ", at: line ", err.Line(), ", column ", err.Column())
	}

}

func Type(t TokenType) string {
	switch t {
	case KeywordToken:
		return "Keyword"

	case VariableToken:
		return "Variable"

	case LiteralStringToken:
		return "string"

	case OperatorToken:
		return "Operator"

	case SymbolToken:
		return "Symbol"

	case GarbageToken:
		return "Garbage"

	default:
		panic("Unknown")
	}
}

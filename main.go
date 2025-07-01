package main

import (
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
		lgen_parser.Parse(NewToken(lgen_lexer.Current().Line(), lgen_lexer.Current().Column(), lgen_parser.EndMarker(), EndToken), collector)
		// fmt.Println(lgen_lexer.Current(), Type(lgen_lexer.Current().Type()))
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

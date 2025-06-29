package lgen

import (
	// . "hulk.com/app/ast"
	// . "hulk.com/app/grammar"
	. "hulk.com/app/interpreter"
	. "hulk.com/app/lexer"
	. "hulk.com/app/lexical_analisys"
	. "hulk.com/app/lgen/lexical"

	// . "hulk.com/app/lgen/sintax"
	// . "hulk.com/app/parser"
	. "hulk.com/app/tokens"
	. "hulk.com/app/tools"
)

var LGENInterpreter IInterpreter

func init() {
	LGENLexer := NewLexer()
	LGENLexer.AddTokenExpression(KeywordToken, 0, KeywordTokenGrammar)
	LGENLexer.AddTokenExpression(SymbolToken, 1, SymbolTokenGrammar)
	LGENLexer.AddTokenExpression(OperatorToken, 2, OperatorTokenGrammar)
	LGENLexer.AddTokenExpression(LiteralStringToken, 3, LiteralStringTokenGrammar)
	LGENLexer.AddTokenExpression(VariableToken, 4, VariableTokenGrammar)

	letters := []rune{}
	for r := rune(65); r < rune(91); r++ {
		letters = append(letters, r)
	}
	for r := rune(97); r < rune(123); r++ {
		letters = append(letters, r)
	}
	for r := rune(48); r < rune(58); r++ {
		letters = append(letters, r)
	}

	analizer := NewLexicalAnalizer()
	analizer.AddRule(KeywordToken, NewLexicalRule("The keywords only can have letters non-whitespace and non-numeric", func(token IToken) bool {
		for _, char := range token.Text() {
			_, err := IndexOf(letters, func(c rune) bool { return c == char })
			if err != nil {
				return false
			}
		}
		return true
	}))
	analizer.AddRule(VariableToken, NewLexicalRule("Variables must start with a letter", func(token IToken) bool {
		_, err := IndexOf(letters, func(c rune) bool { return c == rune(token.Text()[0]) })
		return err == nil
	}))

	// LGENParser := NewParserSLRFromAttributedGrammar(TokenGrammar,NewGrammarSymbol("$",Terminal,false),)

	// LGENInterpreter = NewInterpreter(LGENLexer, analizer)
}

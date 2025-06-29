package lgen

import (
	. "hulk.com/app/lexer"
	. "hulk.com/app/lgen/lexical"
	. "hulk.com/app/tokens"
)

var LGENInterpreter ILexer

func init() {
	LGENInterpreter = NewLexer()
	LGENInterpreter.AddTokenExpression(KeywordToken, 0, KeywordTokenGrammar)
	LGENInterpreter.AddTokenExpression(SymbolToken, 1, SymbolTokenGrammar)
	LGENInterpreter.AddTokenExpression(OperatorToken, 2, OperatorTokenGrammar)
	LGENInterpreter.AddTokenExpression(LiteralStringToken, 3, LiteralStringTokenGrammar)
	LGENInterpreter.AddTokenExpression(VariableToken, 4, VariableTokenGrammar)
}

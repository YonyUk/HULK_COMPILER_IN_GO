package lexerpackage

func ILexerCode() string {
	return `package lexer

import (
	. "hulk.com/app/grammar"
	. "hulk.com/app/tokens"
	. "hulk.com/app/tools"
)

// Lexer interface
type ILexer interface {
	IGenerator[IToken]
	// Load the whole code to tokenize
	LoadCode(code string)
	// Adds a new token to the lexer given the grammar that define that token
	AddTokenExpression(token_type TokenType, priority int, re IGrammar)
}
`
}

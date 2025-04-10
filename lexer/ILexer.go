package lexer

import (
	. "hulk.com/app/grammar"
	. "hulk.com/app/tokens"
)

// Lexer interface
type ILexer interface {
	// Load the whole code to tokenize
	LoadCode(code string)
	// Returns all the tokens
	Tokenize() (chan<- IToken, error)
	// Adds a new token to the lexer given the grammar that define that token
	AddToken(token_type TokenType, re IGrammar)
}

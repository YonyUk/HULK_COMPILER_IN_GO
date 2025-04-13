package lexicalanalisys

import "hulk.com/app/tokens"

// Lexical rule interface
type ILexicalRule interface {
	// Error that will be displayed when this rule won't be acomplished
	ErrorMessage() string
	// Rule predicate
	Rule() func(token tokens.IToken) bool
}

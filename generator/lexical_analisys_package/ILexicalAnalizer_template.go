package lexicalanalisyspackage

func ILexicalAnalizerCode() string {
	return `package lexicalanalisys

import (
	"hulk.com/app/compiler"
	"hulk.com/app/tokens"
)

// Lexical analizer interface
type ILexicalAnalizer interface {
	// Checks the rules for the given token
	Check(token tokens.IToken) compiler.IError
	// Adds a rule for the given token
	AddRule(token_type tokens.TokenType, rule ILexicalRule)
}
`
}

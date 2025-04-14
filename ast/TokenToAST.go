package ast

import (
	"strconv"

	. "hulk.com/app/tokens"
)

var symbolsByToken map[TokenType]string

func TokenToAST(token IToken) IAST {
	switch token.Type() {
	case BooleanToken:
		if token.Text() == "false" {
			return NewAtomicAST(symbolsByToken[BooleanToken], token.Line(), token.Column(), false)
		}
		return NewAtomicAST(symbolsByToken[BooleanToken], token.Line(), token.Column(), true)
	case StringToken:
		return NewAtomicAST(symbolsByToken[StringToken], token.Line(), token.Column(), token.Text())
	case NumberToken:
		value, _ := strconv.ParseFloat(token.Text(), 64)
		return NewAtomicAST(symbolsByToken[BooleanToken], token.Line(), token.Column(), value)
	default:
		return nil
	}
}

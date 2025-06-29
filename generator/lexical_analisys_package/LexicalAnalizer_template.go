package lexicalanalisyspackage

func LexicalAnalizerCode() string {
	return `package lexicalanalisys

import (
	. "hulk.com/app/compiler"
	. "hulk.com/app/tokens"
)

type LexicalAnalizer struct {
	rules map[TokenType][]ILexicalRule
}

func NewLexicalAnalizer() *LexicalAnalizer {
	return &LexicalAnalizer{
		rules: make(map[TokenType][]ILexicalRule),
	}
}

func (lexical *LexicalAnalizer) Check(token IToken) IError {
	if token.Type() == GarbageToken {
		return NewError("Undefined token "+token.Text(), token.Line(), token.Column(), Lexical)
	}
	if rules, ok := lexical.rules[token.Type()]; ok {
		for _, rule := range rules {
			if !rule.Rule()(token) {
				return NewError(rule.ErrorMessage(), token.Line(), token.Column(), Lexical)
			}
		}
	}
	return nil
}

func (lexical *LexicalAnalizer) AddRule(type_ TokenType, rule ILexicalRule) {
	if _, ok := lexical.rules[type_]; ok {
		lexical.rules[type_] = append(lexical.rules[type_], rule)
	} else {
		lexical.rules[type_] = []ILexicalRule{rule}
	}
}
`
}

package lexicalanalisys

import "hulk.com/app/tokens"

type LexicalRule struct {
	msg  string
	rule func(token tokens.IToken) bool
}

func NewLexicalRule(msg string, rule func(token tokens.IToken) bool) *LexicalRule {
	return &LexicalRule{
		msg:  msg,
		rule: rule,
	}
}

func (lr *LexicalRule) ErrorMessage() string {
	return lr.msg
}

func (lr *LexicalRule) Rule() func(token tokens.IToken) bool {
	return lr.rule
}

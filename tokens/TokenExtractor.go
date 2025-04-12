package tokens

import (
	. "hulk.com/app/tools"
)

// Token extractor implementation
type TokenExtractor struct {
	priorities map[int]TokenType
}

func NewTokenExtractor(priorities map[int]TokenType) *TokenExtractor {
	return &TokenExtractor{
		priorities: priorities,
	}
}

func (extractor *TokenExtractor) GetToken(token_types []TokenType, line int, column int, text string) IToken {
	for _, value := range extractor.priorities {
		_, err := IndexOf(token_types, func(t TokenType) bool { return t == value })
		if err == nil {
			return NewToken(line, column, text, value)
		}
	}
	return NewToken(line, column, text, GarbageToken)
}

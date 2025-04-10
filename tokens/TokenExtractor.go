package tokens

import (
	"errors"

	. "hulk.com/app/tools"
)

// Token extractor implementation
type TokenExtractor struct {
	priorities map[int]TokenType
}

func NewTokenExtractor(priorities map[int]TokenType) *TokenExtractor {
	return &TokenExtractor{
		priorities: make(map[int]TokenType),
	}
}

func (extractor *TokenExtractor) GetToken(token_types []TokenType, line int, column int, text string) (IToken, error) {
	for _, value := range extractor.priorities {
		_, err := IndexOf(token_types, func(t TokenType) bool { return t == value })
		if err == nil {
			return NewToken(line, column, text, value), nil
		}
	}
	return nil, errors.New("token type not found")
}

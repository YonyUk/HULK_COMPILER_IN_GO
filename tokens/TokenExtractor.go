package tokens

import (
	. "hulk.com/app/tools"
)

// Token extractor implementation
type TokenExtractor struct {
	priorities       map[int]TokenType
	priorities_order []int
}

func NewTokenExtractor(priorities map[int]TokenType) *TokenExtractor {
	priorities_order := []int{}
	for key, _ := range priorities {
		priorities_order = append(priorities_order, key)
	}
	return &TokenExtractor{
		priorities: priorities,
		priorities_order: MergeSort(priorities_order, func(a int, b int) int {
			if a < b {
				return -1
			}
			if a > b {
				return 1
			}
			return 0
		}),
	}
}

func (extractor *TokenExtractor) GetToken(token_types []TokenType, line int, column int, text string) IToken {
	for _, priority := range extractor.priorities_order {
		_, err := IndexOf(token_types, func(t TokenType) bool { return t == extractor.priorities[priority] })
		if err == nil {
			return NewToken(line, column, text, extractor.priorities[priority])
		}
	}
	return NewToken(line, column, text, GarbageToken)
}

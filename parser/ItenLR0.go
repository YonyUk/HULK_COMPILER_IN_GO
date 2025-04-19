package parser

import (
	. "hulk.com/app/grammar"
)

type ItemLR0 struct {
	head  IGrammarSymbol
	left  []IGrammarSymbol
	right []IGrammarSymbol
}

func NewItemLR0(head IGrammarSymbol, left []IGrammarSymbol, right []IGrammarSymbol) *ItemLR0 {
	return &ItemLR0{
		head:  head,
		left:  left,
		right: right,
	}
}

func (item *ItemLR0) Head() IGrammarSymbol {
	return item.head
}

func (item *ItemLR0) LeftTail() []IGrammarSymbol {
	return item.left
}

func (item *ItemLR0) RightTail() []IGrammarSymbol {
	return item.right
}

package parser

import (
	. "hulk.com/app/grammar"
)

type ItemLR0 struct {
	head  IGrammarSymbol
	left  []IGrammarSymbol
	right []IGrammarSymbol
	id    string
}

func NewItemLR0(head IGrammarSymbol, left []IGrammarSymbol, right []IGrammarSymbol) *ItemLR0 {
	id := head.Symbol() + " --->"
	for _, symbol := range left {
		id += " " + symbol.Symbol()
	}
	id += " ."
	for _, symbol := range right {
		id += " " + symbol.Symbol()
	}
	return &ItemLR0{
		head:  head,
		left:  left,
		right: right,
		id:    id,
	}
}

func (item *ItemLR0) ID() string {
	return item.id
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

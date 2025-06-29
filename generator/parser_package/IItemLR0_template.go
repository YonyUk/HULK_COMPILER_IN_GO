package parserpackage

func IItemLR0Code() string {
	return `package parser

import (
	. "hulk.com/app/grammar"
)

// Item LR(0) interface
type IItemLR0 interface {
	// ID of this item
	ID() string
	// A in the production A ---> A.B where A and B are sentence form
	Head() IGrammarSymbol
	// Left part of the item
	LeftTail() []IGrammarSymbol
	// Right part of the item
	RightTail() []IGrammarSymbol
}
`
}

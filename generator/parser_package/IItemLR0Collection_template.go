package parserpackage

func IItemLR0CollectionCode() string {
	return `package parser

type IItemLR0Collection interface {
	// Id for this collection
	ID() string
	// Items
	Items() []IItemLR0
}
`
}

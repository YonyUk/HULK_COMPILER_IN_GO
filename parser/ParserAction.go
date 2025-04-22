package parser

type ParserAction int

const (
	SHIFT ParserAction = iota
	REDUCE
	ACCEPT
)

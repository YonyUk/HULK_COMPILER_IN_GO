package parserpackage

func ParserActionCode() string {
	return `package parser

type ParserAction int

const (
	SHIFT ParserAction = iota
	REDUCE
	ACCEPT
)
`
}

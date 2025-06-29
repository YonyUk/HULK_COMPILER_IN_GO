package parserpackage

func ActionStructCode() string {
	return `package parser

type ActionStruct struct {
	Action    ParserAction
	NextState string
}
`
}

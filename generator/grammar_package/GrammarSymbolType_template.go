package grammarpackage

func GrammarSymbolTypeCode() string {
	return `package grammar

type GrammarSymbolType int

const (
	Terminal GrammarSymbolType = iota
	NonTerminal
)
`
}

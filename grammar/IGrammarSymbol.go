package grammar

// Grammar symbol interface
type IGrammarSymbol interface {
	// String representation for this symbol
	Symbol() string
	// Return the type of the symbol in the grammar
	Type() GrammarSymbolType
	// True if this symbol is the epsilon symbol
	Epsilon() bool
}

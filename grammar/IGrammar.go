package grammar

// Grammar interface
type IGrammar interface {
	// Terminals of this grammar
	Terminals() []IGrammarSymbol
	// Nonterminals for this grammar
	NonTerminals() []IGrammarSymbol
	// The start symbol of this grammar
	StartSymbol() IGrammarSymbol
	// Adds a new production to this grammar
	AddProduction(symbol IGrammarSymbol, symbols []IGrammarSymbol) error
	// Gets the productions of one symbol
	GetProductions(symbol IGrammarSymbol) [][]IGrammarSymbol
	// Gets the FIRST of the given symbol
	FIRST(symbols []IGrammarSymbol) []IGrammarSymbol
	// Gets the FOLLOW of the given symbol
	FOLLOW(symbol IGrammarSymbol) []IGrammarSymbol
	// Computes the firsts and follows for the grammar
	MakeFirstsAndFollows(endmarker IGrammarSymbol)
}

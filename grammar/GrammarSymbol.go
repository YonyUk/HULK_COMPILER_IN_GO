package grammar

// Grammar symbol implementation
type GrammarSymbol struct {
	symbol  string
	_type   GrammarSymbolType
	epsilon bool
}

func NewGrammarSymbol(symbol string, symbol_type GrammarSymbolType, epsilon bool) *GrammarSymbol {
	return &GrammarSymbol{
		symbol:  symbol,
		_type:   symbol_type,
		epsilon: epsilon,
	}
}

func (symbol *GrammarSymbol) Symbol() string {
	return symbol.symbol
}

func (symbol *GrammarSymbol) Type() GrammarSymbolType {
	return symbol._type
}

func (symbol *GrammarSymbol) Epsilon() bool {
	return symbol.epsilon
}

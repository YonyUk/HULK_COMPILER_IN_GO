package grammar

// Grammar symbol implementation
type GrammarSymbol struct {
	symbol string
	_type  GrammarSymbolType
}

func NewGrammarSymbol(symbol string, symbol_type GrammarSymbolType) *GrammarSymbol {
	return &GrammarSymbol{
		symbol: symbol,
		_type:  symbol_type,
	}
}

func (symbol *GrammarSymbol) Symbol() string {
	return symbol.symbol
}

func (symbol *GrammarSymbol) Type() GrammarSymbolType {
	return symbol._type
}

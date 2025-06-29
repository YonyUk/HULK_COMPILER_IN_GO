package tokens

type TokenType int

const (
	GarbageToken TokenType = iota
	EndToken
	KeywordToken
	SymbolToken
	OperatorToken
	LiteralStringToken
	VariableToken
)

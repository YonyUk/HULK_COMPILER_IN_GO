package tokens

type TokenType int

const (
	GarbageToken TokenType = iota
	KeywordToken
	SymbolToken
	OperatorToken
	VariableToken
	NumberToken
	BooleanToken
	StringToken
)

package tokens

type TokenType int

const (
	InvalidToken TokenType = iota
	KeywordToken
	SymbolToken
	OperatorToken
	VariableToken
	NumberToken
	BooleanToken
	StringToken
)

package tokenspackage

func TokensCode() string {
	return `package tokens

// Tokens definition
type Token struct {
	line       int
	column     int
	text       string
	token_type TokenType
}

func NewToken(line int, column int, text string, token_type TokenType) *Token {
	return &Token{
		line:       line,
		column:     column,
		text:       text,
		token_type: token_type,
	}
}

func (token *Token) Line() int {
	return token.line
}

func (token *Token) Column() int {
	return token.column
}

func (token *Token) Text() string {
	return token.text
}

func (token *Token) Type() TokenType {
	return token.token_type
}
`
}

package tokens

// Token recognizer interface
type ITokenExtractor interface {
	GetToken(token_types []TokenType, line int, column int, text string) IToken
}

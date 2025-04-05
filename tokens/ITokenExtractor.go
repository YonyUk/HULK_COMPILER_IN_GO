package tokens

// Token recognizer interface
type ITokenExtractor interface {
	GetToken(token_type TokenType) Token
}

package tokens

// Tokens interface
type IToken interface {
	Line() int
	Column() int
	Text() string
	Type() TokenType
}

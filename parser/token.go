package parser

type Token int

const (
	TokenEof Token = iota
	TokenIdent
	TokenString
	TokenNumber
	TokenComma
	TokenEqual
	TokenOpenBracket
	TokenCloseBracket
	TokenOpenSquareBracket
	TokenCloseSquareBracket
)

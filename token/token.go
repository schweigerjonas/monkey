package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

// TokenTypes
const (
	ILLEGAL = "ILLEGAL" // for unknown TokenTypes 
	EOF = "EOF" // "end of file"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 123456
	
	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// token we don't know
	ILLEGAL = "ILLEGAL"

	// "end of file"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Keywords are used to identify the type of token.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Checks if the identifier is a keyword. If it is, it returns the token type for that keyword.
// If it is not a keyword, it returns the IDENT token type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

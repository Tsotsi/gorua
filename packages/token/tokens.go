package token

//go:generate stringer -type Token
type Token int
type Pos int // rune position
var NoPos Pos = 0

const (
	ILLEGAL Token = iota
	EOF
	COMMENT

	// ADD operator
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND // &
	OR
	XOR

	LPAREN // (
	RPAREN
	LBRACK // [
	RBRACK
	LBRACE // {
	RBRACE

	SEMICOLON // ;
	COLON     // :

	// literal
	IDENT
	INT
	FLOAT
	STRING

	// keywords

	DEF
	END
	IF
	ELSE
	RETURN
)

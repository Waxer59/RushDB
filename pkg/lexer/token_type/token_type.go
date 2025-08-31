package token_type

type TokenType int

type Token struct {
	Value string
	Type  TokenType
}

const (
	// Types
	Number TokenType = iota
	Identifier
	Null
	String
	Boolean

	// Keywords
	Create
	Table
	Select
	From
	Drop
	Insert
	Update
	Delete
	Alter

	// Grouping
	LeftParen  // (
	RightParen // )
	SemiColon  // ;

	// Comparison operators
	EqualEqual   // ==
	Greater      // >
	GreaterEqual // >=
	Less         // <
	LessEqual    // <=
	Bang         // !
	NotEqual     // !=
	Or           // ||
	And          // &&

	// End Of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"create": Create,
	"table":  Table,
	"select": Select,
	"drop":   Drop,
	"insert": Insert,
	"update": Update,
	"delete": Delete,
	"alter":  Alter,
	"from":   From,
	"null":   Null,
	"true":   Boolean,
	"false":  Boolean,
	"bool":   Boolean,
	"number": Number,
	"string": String,
}

var SkippableChars = []rune{' ', '\t', '\n', '\r'}

package ast_types

type NodeType string

const (
	// Statements
	Program              NodeType = "Program"
	SelectStatement      NodeType = "SelectStatement"
	CreateTableStatement NodeType = "CreateTableStatement"
	DropTableStatement   NodeType = "DropTableStatement"
	InsertStatement      NodeType = "InsertStatement"
	UpdateStatement      NodeType = "UpdateStatement"
	DeleteStatement      NodeType = "DeleteStatement"
	AlterTableStatement  NodeType = "AlterTableStatement"

	// Expressions
	Identifier NodeType = "Identifier"
	BinaryExpr NodeType = "BinaryExpr"

	// Clauses
	FromClause  NodeType = "FromClause"
	WhereClause NodeType = "WhereClause"
)

var (
	// BOOLEAN EXPR
	ComparisonExpr = []string{"<", "<=", ">", ">="}
	EqualityExpr   = []string{"==", "!="}
	BoolExpr       = []string{"<", "<=", ">", ">=", "==", "!="}
)

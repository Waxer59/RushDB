package ast

import "github.com/Waxer59/RushDB/pkg/ast/ast_types"

type Expr interface {
	Stmt
}

type BinaryExpr struct {
	Kind     ast_types.NodeType
	Left     Expr
	Right    Expr
	Operator string
}

func (b BinaryExpr) GetKind() ast_types.NodeType {
	return b.Kind
}

type Identifier struct {
	Kind   ast_types.NodeType
	Symbol string
}

func (i Identifier) GetKind() ast_types.NodeType {
	return i.Kind
}

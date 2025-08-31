package ast

import "github.com/Waxer59/RushDB/pkg/ast/ast_types"

type Clause interface {
	Stmt
}

type FromClause struct {
	Kind  ast_types.NodeType
	Exprs Expr
}

func (f FromClause) GetKind() ast_types.NodeType {
	return f.Kind
}

type WhereClause struct {
	Kind ast_types.NodeType
	Expr Expr
}

func (w WhereClause) GetKind() ast_types.NodeType {
	return w.Kind
}

package ast

import "github.com/Waxer59/RushDB/pkg/ast/ast_types"

type Stmt interface {
	GetKind() ast_types.NodeType
}

type Program struct {
	Kind ast_types.NodeType
	Body []Stmt
}

func (p Program) GetKind() ast_types.NodeType {
	return p.Kind
}

type CreateTableStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (c CreateTableStatement) GetKind() ast_types.NodeType {
	return c.Kind
}

type DropTableStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (d DropTableStatement) GetKind() ast_types.NodeType {
	return d.Kind
}

type InsertStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (i InsertStatement) GetKind() ast_types.NodeType {
	return i.Kind
}

type UpdateStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (u UpdateStatement) GetKind() ast_types.NodeType {
	return u.Kind
}

type DeleteStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (d DeleteStatement) GetKind() ast_types.NodeType {
	return d.Kind
}

type AlterTableStatement struct {
	Kind    ast_types.NodeType
	Clauses []Clause
}

func (a AlterTableStatement) GetKind() ast_types.NodeType {
	return a.Kind
}

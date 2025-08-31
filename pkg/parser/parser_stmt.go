package parser

import (
	"errors"

	"github.com/Waxer59/RushDB/pkg/ast"
	"github.com/Waxer59/RushDB/pkg/lexer/token_type"
)

func (p *Parser) ParseStmt() (ast.Stmt, error) {
	switch p.at().Type {
	case token_type.Select:
		return p.ParseSelectStmt()
	case token_type.Create:
		return p.ParseCreateTableStmt()
	case token_type.Drop:
		return p.ParseDropTableStmt()
	case token_type.Insert:
		return p.ParseInsertStmt()
	case token_type.Update:
		return p.ParseUpdateStmt()
	case token_type.Delete:
		return p.ParseDeleteStmt()
	case token_type.Alter:
		return p.ParseAlterTableStmt()
	}

	return nil, errors.New("Invalid statement")
}

func (p *Parser) ParseSelectStmt(ast.Stmt, error) {}

func (p *Parser) ParseCreateTableStmt(ast.Stmt, error) {}

func (p *Parser) ParseDropTableStmt(ast.Stmt, error) {}

func (p *Parser) ParseInsertStmt(ast.Stmt, error) {}

func (p *Parser) ParseUpdateStmt(ast.Stmt, error) {}

func (p *Parser) ParseDeleteStmt(ast.Stmt, error) {}

func (p *Parser) ParseAlterTableStmt(ast.Stmt, error) {}

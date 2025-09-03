package parser

import (
	"errors"

	"github.com/Waxer59/RushDB/pkg/lexer/token_type"
)

func (p *Parser) at() token_type.Token {
	return p.tokens[0]
}

func (p *Parser) atNext() token_type.Token {
	if len(p.tokens) > 1 {
		return p.tokens[1]
	}
	return token_type.Token{}
}

func (p *Parser) subtract(params ...int) token_type.Token {
	n := 1
	if len(params) > 0 {
		n = params[0]
	}

	prev := p.at()
	p.tokens = p.tokens[n:]
	return prev
}

func (p *Parser) expect(typeExpected token_type.TokenType, errMsg string) (token_type.Token, error) {
	prev := p.subtract()

	if (prev == token_type.Token{} || prev.Type != typeExpected) {
		return prev, errors.New(errMsg)
	}

	return prev, nil
}

func (p *Parser) notEOF() bool {
	return p.at().Type != token_type.EOF
}

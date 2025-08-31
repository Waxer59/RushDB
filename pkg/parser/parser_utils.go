package parser

import "github.com/Waxer59/RushDB/pkg/lexer/token_type"

func (p *Parser) at() token_type.Token {
	return p.tokens[0]
}

func (p *Parser) notEOF() bool {
	return p.at().Type != token_type.EOF
}

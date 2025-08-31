package lexer

import (
	"github.com/Waxer59/RushDB/pkg/lexer/internal/utils"
	"github.com/Waxer59/RushDB/pkg/lexer/token_type"
)

func Tokenize(input string) ([]token_type.Token, error) {
	var tokens []token_type.Token
	src := []rune(input)
	subtract := func(i int) rune {
		if len(src) <= 0 {
			return 0
		}
		char := src[0]
		src = src[i:]
		return char
	}

	nextChar := func() rune {
		if len(src) <= 1 {
			return 0
		}
		return src[1]
	}

	for len(src) > 0 {
		tokenChar := src[0]

		// Check if token is skippable
		if utils.IsSkippable(tokenChar) {
			subtract(1)
			continue
		}

		// Check for number
		if utils.IsInt(tokenChar) {
			num, rest := utils.ExtractNum(src)
			tokens = append(tokens, token_type.Token{Type: token_type.Number, Value: num})
			src = rest
			continue
		}

		// Check for alpha
		if utils.IsAlpha(tokenChar) {
			alpha, rest := utils.ExtractIdentifier(src)
			alphaType := token_type.Identifier

			if keyword, ok := utils.IsKeyword(alpha); ok {
				alphaType = keyword
			}

			tokens = append(tokens, token_type.Token{Type: alphaType, Value: alpha})

			src = rest
			continue
		}

		// Check for operators
		switch tokenChar {
		case '=':
			switch nextChar() {
			case '=':
				subtract(2) // consume '=='
				tokens = append(tokens, token_type.Token{Type: token_type.EqualEqual, Value: "=="})
				continue
			}
		case '!':
			if nextChar() == '=' {
				subtract(2) // consume '!='
				tokens = append(tokens, token_type.Token{Type: token_type.NotEqual, Value: "!="})
				continue
			}
		case '>':
			if nextChar() == '=' {
				subtract(2) // consume '>='
				tokens = append(tokens, token_type.Token{Type: token_type.GreaterEqual, Value: ">="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Greater, Value: string(tokenChar)})
		case '<':
			if nextChar() == '=' {
				subtract(2) // consume '<='
				tokens = append(tokens, token_type.Token{Type: token_type.LessEqual, Value: "<="})
				continue
			}
			tokens = append(tokens, token_type.Token{Type: token_type.Less, Value: string(tokenChar)})
		case '(':
			tokens = append(tokens, token_type.Token{Type: token_type.LeftParen, Value: string(tokenChar)})
		case ')':
			tokens = append(tokens, token_type.Token{Type: token_type.RightParen, Value: string(tokenChar)})
		default:
			tokens = append(tokens, token_type.Token{Type: token_type.Identifier, Value: string(tokenChar)})
		}
		subtract(1)
	}
	tokens = append(tokens, token_type.Token{Type: token_type.EOF, Value: "EndOfFile"})
	return tokens, nil
}

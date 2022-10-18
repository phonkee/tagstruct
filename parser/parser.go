package parser

import (
	"fmt"
	"io"
)

// Parse parses given input and returns parsed properties
func Parse(input io.Reader) ([]Property, error) {
	p := parser{
		lex: newLexer(input),
	}
	return p.Parse()
}

type parser struct {
	lex *lexer
}

func (p *parser) Parse() ([]Property, error) {
	prop, err := p.parseObject(true)
	if err != nil {
		return nil, err
	}
	return prop.Object, nil
}

// parseIdent parses identifier (can start with "=" or "[" or "(")
func (p *parser) parseIdent() (prop Property, _ error) {
	for {
		pos, token, _ := p.lex.Lex()
		prop.Position = pos - 1

		switch token {
		case TokenEof:
			return prop, nil
		case TokenEqual:
			value, err := p.parseValue()
			if err != nil {
				return prop, err
			}
			prop.Value = &value
			return prop, nil
		case TokenOpenBracket:
			value, err := p.parseObject(false)
			if err != nil {
				return prop, err
			}
			prop.Object = value.Object
			return prop, nil
		case TokenOpenSquareBracket:
			arr, err := p.parseArray()
			if err != nil {
				return prop, err
			}
			prop.Array = arr.Array
			return prop, nil
		}
		panic("nope")
	}
}

// parseArray parses array values, it's a bit tricky
func (p *parser) parseArray() (result Property, _ error) {
	hasValue := false
	for {
		pos, ident, value := p.lex.Lex()
		switch ident {
		case TokenComma:
			if !hasValue {
				return result, fmt.Errorf("%w: unexpected comma at [%v]", ErrParser, pos)
			}
			continue
		case TokenEof:
			return result, fmt.Errorf("unexpected end of file [array]: %w at [%v]", ErrParser, pos)
		case TokenNumber, TokenString:
			switch ident {
			case TokenNumber:
				result.Array = append(result.Array, Property{
					Position: pos - 1,
					Value:    &Value{Position: pos - 1, Number: &value},
				})
			case TokenString:
				result.Array = append(result.Array, Property{
					Position: pos - 1,
					Value:    &Value{Position: pos - 1, String: &value},
				})
			}
			hasValue = true
		case TokenCloseSquareBracket:
			return result, nil
		case TokenIdent:
			prop, err := p.parseIdent()
			if err != nil {
				return result, err
			}
			prop.Position = pos - 1
			prop.Name = value
			result.Array = append(result.Array, prop)
			hasValue = true
		default:
			return result, fmt.Errorf("unexpected token [array]: %w at [%v]", ErrParser, pos)
		}
	}
}

func (p *parser) parseObject(allowMissingClose bool) (result Property, _ error) {
	result.Object = make([]Property, 0)

outer:
	for {
		pos, token, value := p.lex.Lex()

		switch token {
		case TokenIdent:
			prop, err := p.parseIdent()
			if err != nil {
				return prop, err
			}
			prop.Position = pos - 1
			prop.Name = value
			result.Object = append(result.Object, prop)

			_, token, _ = p.lex.Lex()
			switch token {
			case TokenEof:
				if allowMissingClose {
					break outer
				}
				return result, fmt.Errorf("%w: unexpected end of stream", ErrParser)
			case TokenCloseBracket:
				break outer
			case TokenComma:
				continue outer
			}
		case TokenEof:
			break outer
		}
	}

	return result, nil
}

func (p *parser) parseValue() (result Value, _ error) {
	pos, tok, str := p.lex.Lex()
	result.Position = pos - 1

	switch tok {
	case TokenString, TokenIdent:
		result.String = &str
		return result, nil
	case TokenNumber:
		result.Number = &str
		return result, nil
	}

	return result, fmt.Errorf("%w: unexpected token %v", ErrParser, str)
}

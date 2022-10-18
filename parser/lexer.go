package parser

import (
	"bufio"
	"io"
	"unicode"
)

func newLexer(reader io.Reader) *lexer {
	return &lexer{
		reader: bufio.NewReader(reader),
	}
}

type lexer struct {
	pos    int
	reader *bufio.Reader
}

func (l *lexer) Lex() (int, Token, string) {
	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return l.pos, TokenEof, ""
			}
			// this should not happen
			panic(err)
		}

		switch r {
		case '(':
			return l.pos, TokenOpenBracket, ")"
		case ')':
			return l.pos, TokenCloseBracket, ")"
		case '[':
			return l.pos, TokenOpenSquareBracket, "]"
		case ']':
			return l.pos, TokenCloseSquareBracket, "]"
		case '=':
			return l.pos, TokenEqual, "="
		case ',':
			return l.pos, TokenComma, ","
		case '\'':
			pos := l.pos
			str := ""
			for {
				r, err := l.read()
				if err != nil {
					if err == io.EOF {
						return l.pos, TokenEof, ""
					}
					// this should not happen
					panic(err)
				}

				if r == '\'' {
					return pos, TokenString, str
				}

				str += string(r)
			}
		default:
			if unicode.IsSpace(r) {
				continue // nothing to do here, just move on
			}
			if unicode.IsDigit(r) {
				pos := l.pos
				str := string(r)
				for {
					r, err := l.read()
					if err != nil {
						if err == io.EOF {
							break
						}
						panic(err)
					}

					if unicode.IsNumber(r) {
						str += string(r)
					} else if r == '.' {
						str += string(r)
					} else {
						l.unread()
						break
					}
				}
				return pos, TokenNumber, str
			}
			if unicode.IsLetter(r) {
				pos := l.pos
				str := string(r)

				for {
					r, err := l.read()
					if err != nil {
						if err == io.EOF {
							break
						}
						panic(err)
					}

					if unicode.IsNumber(r) || unicode.IsLetter(r) || r == '_' {
						str += string(r)
					} else {
						l.unread()
						break
					}
				}
				return pos, TokenIdent, str

			}
		}
	}
}

func (l *lexer) peek() (rune, error) {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return 0, err
	}
	_ = l.reader.UnreadRune()
	return r, err
}

func (l *lexer) read() (rune, error) {
	r, _, err := l.reader.ReadRune()
	l.pos++
	return r, err
}

func (l *lexer) unread() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos--
}

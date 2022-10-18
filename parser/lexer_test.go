package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer_Lex(t *testing.T) {
	t.Run("test basic values", func(t *testing.T) {
		data := []struct {
			input    string
			token    Token
			value    string
			position int
		}{
			{"", TokenEof, "", 1},
			{"(", TokenOpenBracket, "(", 1},
			{")", TokenCloseBracket, ")", 1},
			{"[", TokenOpenSquareBracket, "[", 1},
			{"]", TokenCloseSquareBracket, "]", 1},
			{"'hello world'", TokenString, "hello world", 1},
			{"1234", TokenNumber, "1234", 1},
			{"1234xxx", TokenNumber, "1234", 1},
			{"1234.566", TokenNumber, "1234.566", 1},
			{"ident", TokenIdent, "ident", 1},
			{"ident_1", TokenIdent, "ident_1", 1},
		}

		for _, item := range data {
			pos, token, _ := NewLexer(strings.NewReader(item.input)).Lex()
			assert.Equal(t, item.position, pos)
			assert.Equal(t, item.token, token)
		}
	})
}

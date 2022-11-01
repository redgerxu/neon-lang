package lexer

import (
	"github.com/narutopig/neon-lang/lang/token"
)

// Tokenize parses the next few characters and returns a token and the length of the token
func (l *Lexer) Tokenize(str string) {
	b := []byte(str)
	ws := token.WhiteSpace.FindIndex(b)

	if ws != nil {
		wse := ws[1]
		l.cursor += wse
		return
	}

	for key, value := range token.Spec {
		res := key.FindIndex(b)

		if res == nil {
			continue
		}

		end := res[1]

		l.cursor += end
		l.tokens = append(l.tokens, token.New(value, str[0:end]))
		return
	}

	panic("Unexpected token")
}

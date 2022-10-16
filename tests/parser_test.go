package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/lib"
)

func TestParse(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []lib.Token
		wantErr bool
	}{
		{"Simple hello world", args{"print(\"Hello World!\")"}, []lib.Token{
			lib.NewToken(lib.IDENTIFIER, "print"),
			lib.NewToken(lib.LEFTPAREN, ""),
			lib.NewToken(lib.STRINGVALUE, "Hello World!"),
			lib.NewToken(lib.RIGHTPAREN, ""),
		}, false},
		{
			"Main function",
			args{"int main() {\nprint(\"Hello World!\")\n}"},
			[]lib.Token{
				lib.NewToken(lib.INTTYPE, ""),
				lib.NewToken(lib.IDENTIFIER, "main"),
				lib.NewToken(lib.LEFTPAREN, ""),
				lib.NewToken(lib.RIGHTPAREN, ""),
				lib.NewToken(lib.LEFTCURLY, ""),
				lib.NewToken(lib.NEWLINE, ""),
				lib.NewToken(lib.IDENTIFIER, "print"),
				lib.NewToken(lib.LEFTPAREN, ""),
				lib.NewToken(lib.STRINGVALUE, "Hello World!"),
				lib.NewToken(lib.RIGHTPAREN, ""),
				lib.NewToken(lib.NEWLINE, ""),
				lib.NewToken(lib.RIGHTCURLY, ""),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lib.Parse(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

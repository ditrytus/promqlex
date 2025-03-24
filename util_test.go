package promqlex

import (
	_ "embed"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"testing"
)

//go:embed testdata/promql/promql_examples.yaml
var promqlExamplesYaml string

type Examples struct {
	Source  string
	Queries []string
}

type FailTestErrorListener struct {
	antlr.DefaultErrorListener
	t *testing.T
}

func (f FailTestErrorListener) SyntaxError(antlr.Recognizer, interface{}, int, int, string, antlr.RecognitionException) {
	f.t.Error("Syntax Error")
}

func NewFailTestErrorListener(t *testing.T) *FailTestErrorListener {
	return &FailTestErrorListener{
		t: t,
	}
}

func debugPrintLexer(lexer antlr.Lexer) {
	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Println(tokenStringWithName(token, lexer))
	}
}

package promqlex

import (
	_ "embed"
	"github.com/antlr4-go/antlr/v4"
	"testing"
)

//go:embed test/promql_examples.yaml
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

func (f FailTestErrorListener) ReportContextSensitivity(antlr.Parser, *antlr.DFA, int, int, int, *antlr.ATNConfigSet) {
	f.t.Error("Context Sensitivity")
}

func NewFailTestErrorListener(t *testing.T) *FailTestErrorListener {
	return &FailTestErrorListener{
		t: t,
	}
}

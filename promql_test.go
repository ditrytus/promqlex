package promqlex

import (
	_ "embed"
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promql"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

//go:embed test/promql_examples.yaml
var promqlExamplesYaml string

type Examples struct {
	Source  string
	Queries []string
}

type FailTestErrorListener struct {
	t *testing.T
}

func (f FailTestErrorListener) SyntaxError(antlr.Recognizer, interface{}, int, int, string, antlr.RecognitionException) {
	f.t.Error("Syntax Error")
}

func (f FailTestErrorListener) ReportAmbiguity(antlr.Parser, *antlr.DFA, int, int, bool, *antlr.BitSet, *antlr.ATNConfigSet) {
	//f.t.Error("Ambiguity")
}

func (f FailTestErrorListener) ReportAttemptingFullContext(antlr.Parser, *antlr.DFA, int, int, *antlr.BitSet, *antlr.ATNConfigSet) {
	//f.t.Error("Attempting Full Context")
}

func (f FailTestErrorListener) ReportContextSensitivity(antlr.Parser, *antlr.DFA, int, int, int, *antlr.ATNConfigSet) {
	f.t.Error("Context Sensitivity")
}

func NewFailTestErrorListener(t *testing.T) *FailTestErrorListener {
	return &FailTestErrorListener{
		t: t,
	}
}

func TestPromQLParser_Parse(t *testing.T) {
	yamlDecoder := yaml.NewDecoder(strings.NewReader(promqlExamplesYaml))
	for {
		var example Examples
		err := yamlDecoder.Decode(&example)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Run(example.Source, func(t *testing.T) {
			for _, query := range example.Queries {
				t.Run(query, func(t *testing.T) {
					input := antlr.NewInputStream(query)
					lexer := promql.NewPromQLLexer(input)
					tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
					p := promql.NewPromQLParser(tokens)
					p.RemoveErrorListeners()
					p.AddErrorListener(NewFailTestErrorListener(t))
					p.AddErrorListener(antlr.NewConsoleErrorListener())
					_ = p.Expression()
				})
			}
		})
	}
}

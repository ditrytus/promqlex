package promqlex

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

func TestMacroEvaluator(t *testing.T) {
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
					_, tree, _ := parsePromQLEx(t, query)
					macroEval := NewSymbolsDefiner()
					walker := antlr.NewParseTreeWalker()
					walker.Walk(macroEval, tree)
				})
			}
		})
	}
}

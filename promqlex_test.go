package promqlex

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlextension"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

func TestPromQLExtensionParser_ValidPromQLIsValidPromQLEx(t *testing.T) {
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
					functionSet := NewPrometheusFunctionSet(
						promqlextension.PromQLExtensionLexerFUNCTION,
						promqlextension.PromQLExtensionLexerAGGREGATION_OPERATOR,
					)
					functionSet.AddAggregationOperators(PrometheusExperimentalAggregationOperators)
					providerInput := NewFunctionProviderInputStream(input, functionSet)
					lexer := promqlextension.NewPromQLExtensionLexer(providerInput)
					//readLexer(lexer)
					tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
					p := promqlextension.NewPromQLExtensionParser(tokens)
					p.RemoveErrorListeners()
					p.AddErrorListener(NewFailTestErrorListener(t))
					p.AddErrorListener(antlr.NewConsoleErrorListener())
					_ = p.Promqlx()
				})
			}
		})
	}
}

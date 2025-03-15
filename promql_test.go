package promqlex

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promql"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

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
					functionSet := NewPrometheusFunctionSet(
						promql.PromQLLexerFUNCTION,
						promql.PromQLLexerAGGREGATION_OPERATOR,
					)
					functionSet.AddAggregationOperators(PrometheusExperimentalAggregationOperators)
					providerInput := NewFunctionProviderInputStream(input, functionSet)
					lexer := promql.NewPromQLLexer(providerInput)
					lexer.GetTokenNames()
					tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
					p := promql.NewPromQLParser(tokens)
					p.RemoveErrorListeners()
					p.AddErrorListener(NewFailTestErrorListener(t))
					p.AddErrorListener(antlr.NewConsoleErrorListener())
					tree := p.Expression()
					var builder strings.Builder
					NewAsciiAstPrinterVisitor(&builder, p, lexer).Visit(tree)
					fmt.Println(builder.String())
				})
			}
		})
	}
}

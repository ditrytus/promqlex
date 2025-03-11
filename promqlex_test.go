package promqlex

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/ditrytus/promqlex/parsers/promqlex"
	"gopkg.in/yaml.v3"
	"io"
	"strings"
	"testing"
)

func TestPromQLExParser_ValidPromQLIsValidPromQLEx(t *testing.T) {
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
						parser.PromQLExLexerFUNCTION,
						parser.PromQLExLexerAGGREGATION_OPERATOR,
					)
					functionSet.AddAggregationOperators(PrometheusExperimentalAggregationOperators)
					providerInput := NewFunctionProviderInputStream(input, functionSet)
					lexer := parser.NewPromQLExLexer(providerInput)
					tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
					p := parser.NewPromQLExParser(tokens)
					p.RemoveErrorListeners()
					p.AddErrorListener(NewFailTestErrorListener(t))
					p.AddErrorListener(antlr.NewConsoleErrorListener())
					tree := p.Promqlx()
					if query == "rate(http_requests_total[5m])[30m:1m]" {
						fmt.Print()
					}
					var builder strings.Builder
					NewAsciiAstPrinterVisitor(&builder, p).Visit(tree)
					fmt.Println(builder.String())
				})
			}
		})
	}
}

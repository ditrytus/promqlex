package promqlex

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "github.com/ditrytus/promqlex/parsers/promqlex"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	input := antlr.NewInputStream("http_requests_count @ 2024-01-01T00:00:00")
	functionSet := NewPrometheusFunctionSet(
		parser.PromQLExLexerFUNCTION,
		parser.PromQLExLexerAGGREGATION_OPERATOR,
	)
	functionSet.AddAggregationOperators(PrometheusExperimentalAggregationOperators)
	providerInput := NewInputStream(input, functionSet)
	lexer := parser.NewPromQLExLexer(providerInput)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPromQLExParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(NewFailTestErrorListener(t))
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	tree := p.Promqlx()
	var builder strings.Builder
	NewAsciiAstPrinterVisitor(&builder, p, lexer).Visit(tree)
	fmt.Println(builder.String())
	rewriter := antlr.NewTokenStreamRewriter(tokens)
	transpiler := NewTranspiler(rewriter)
	walker := antlr.NewParseTreeWalker()
	walker.Walk(transpiler, tree)
	fmt.Println(rewriter.GetTextDefault())
}

package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promql"
	parser "github.com/ditrytus/promqlex/parsers/promqlex"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type TranspilerTestCase struct {
	Name     string
	PromQL   string
	PromQLEx string
	Sub      []TranspilerTestCase
}

func TestTranspiler(t *testing.T) {
	if err := filepath.WalkDir("testdata/promqlex", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		t.Run(path, func(t *testing.T) {
			file, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			var testCases []TranspilerTestCase
			if err := yaml.Unmarshal(file, &testCases); err != nil {
				t.Fatal(err)
			}

			runTestCases(t, testCases)
		})
		return nil
	}); err != nil {
		t.Error(err)
	}
}

func runTestCases(t *testing.T, testCases []TranspilerTestCase) {
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.PromQLEx != "" {
				transpiledPromQL := transpilePromQLEx(t, testCase.PromQLEx)
				transpiledPromQLASCIITree := parsePromQL(t, transpiledPromQL)
				expectedPromQLASCIITree := parsePromQL(t, testCase.PromQL)
				assert.Equal(t, expectedPromQLASCIITree, transpiledPromQLASCIITree)
			}
			runTestCases(t, testCase.Sub)
		})
	}
}

func transpilePromQLEx(t *testing.T, promQLEx string) string {
	promQLExInputStream := newInputStream(
		parser.PromQLExLexerFUNCTION,
		parser.PromQLExLexerAGGREGATION_OPERATOR,
		promQLEx,
	)
	promQLExLexer := parser.NewPromQLExLexer(promQLExInputStream)
	promQLExTokens := antlr.NewCommonTokenStream(promQLExLexer, antlr.TokenDefaultChannel)
	promQLExParser := parser.NewPromQLExParser(promQLExTokens)
	setErrorListeners(t, promQLExParser)
	promQLExTree := promQLExParser.Promqlx()
	var promQLExBuilder strings.Builder
	if promQLExParser.HasError() {
		NewAsciiAstPrinterVisitor(&promQLExBuilder, promQLExParser, promQLExLexer).Visit(promQLExTree)
		t.Log("PromQLEx text:\n")
		t.Log(promQLEx + "\n")
		t.Log("PromQLEx AST:\n")
		t.Log(promQLExBuilder.String())
	}

	rewriter := antlr.NewTokenStreamRewriter(promQLExTokens)
	transpiler := NewTranspiler(rewriter)
	walker := antlr.NewParseTreeWalker()

	walker.Walk(transpiler, promQLExTree)
	result := rewriter.GetTextDefault()
	return result
}

func parsePromQL(t *testing.T, result string) string {
	promQLInputStream := newInputStream(
		promql.PromQLLexerFUNCTION,
		promql.PromQLLexerAGGREGATION_OPERATOR,
		result,
	)
	promQLLexer := promql.NewPromQLLexer(promQLInputStream)
	promQLTokens := antlr.NewCommonTokenStream(promQLLexer, antlr.TokenDefaultChannel)
	promQLParser := promql.NewPromQLParser(promQLTokens)
	setErrorListeners(t, promQLParser)
	promQLTree := promQLParser.Expression()
	var promQLBuilder strings.Builder
	NewAsciiAstPrinterVisitor(&promQLBuilder, promQLParser, promQLLexer).Visit(promQLTree)
	str := promQLBuilder.String()
	if promQLParser.HasError() {
		t.Log("PromQL text:\n")
		t.Log(result + "\n")
		t.Log("PromQL AST:\n")
		t.Log(str)
	}
	return str
}

func setErrorListeners(t *testing.T, p antlr.Recognizer) {
	p.RemoveErrorListeners()
	p.AddErrorListener(NewFailTestErrorListener(t))
	p.AddErrorListener(antlr.NewConsoleErrorListener())
}

func newInputStream(functionTokenType int, aggregationOperatorTokenType int, promQLEx string) *InputStream {
	input := antlr.NewInputStream(promQLEx)
	functionSet := NewPrometheusFunctionSet(
		functionTokenType,
		aggregationOperatorTokenType,
	)
	functionSet.AddAggregationOperators(PrometheusExperimentalAggregationOperators)
	inputStream := NewInputStream(input, functionSet)
	return inputStream
}

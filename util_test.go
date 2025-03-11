package promqlex

import (
	_ "embed"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
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

func debugPrintLexer(lexer antlr.Lexer) {
	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Println(token.String())
	}
}

type AsciiAstPrinter struct {
	io.Writer
	prefix string
	antlr.Recognizer
}

func (a *AsciiAstPrinter) Visit(tree antlr.ParseTree) interface{} {
	tree.Accept(a)
	return nil
}

func NewAsciiAstPrinterVisitor(w io.Writer, recognizer antlr.Recognizer) *AsciiAstPrinter {
	return &AsciiAstPrinter{
		Writer:     w,
		Recognizer: recognizer,
	}
}

func (a *AsciiAstPrinter) VisitChildren(node antlr.RuleNode) interface{} {
	ruleIndex := node.GetRuleContext().GetRuleIndex()
	ruleName := a.Recognizer.GetRuleNames()[ruleIndex]
	line := fmt.Sprintf("%s%s\n", a.prefix, ruleName)
	_, _ = a.Write([]byte(line))
	const (
		TREE_DOWN_RIGHT = "├ "
		TREE_DOWN       = "│ "
		TREE_RIGHT      = "└ "
		SPACE           = "  "
	)
	var replaceMap = map[string]string{
		TREE_DOWN_RIGHT: TREE_DOWN,
		TREE_RIGHT:      SPACE,
	}
	var replacedSuffix string
	for from, to := range replaceMap {
		if strings.HasSuffix(a.prefix, from) {
			a.prefix = strings.TrimSuffix(a.prefix, from) + to
			replacedSuffix = from
			break
		}
	}
	for i, child := range node.GetChildren() {
		var appendedSuffix string
		if i == node.GetChildCount()-1 {
			appendedSuffix = TREE_RIGHT
		} else {
			appendedSuffix = TREE_DOWN_RIGHT
		}
		a.prefix = a.prefix + appendedSuffix
		if parseTree, ok := child.(antlr.ParseTree); ok {
			parseTree.Accept(a)
		}
		a.prefix = strings.TrimSuffix(a.prefix, appendedSuffix)
	}
	if replacedSuffix != "" {
		a.prefix = strings.TrimSuffix(a.prefix, replaceMap[replacedSuffix]) + replacedSuffix
	}
	return nil
}

func (a *AsciiAstPrinter) VisitTerminal(node antlr.TerminalNode) interface{} {
	line := fmt.Sprintf("%s%s\n", a.prefix, node.GetSymbol().String())
	_, _ = a.Write([]byte(line))
	return nil
}

func (a *AsciiAstPrinter) VisitErrorNode(node antlr.ErrorNode) interface{} {
	line := fmt.Sprintf("%s[ERROR] %s\n", a.prefix, node.GetSymbol().String())
	_, _ = a.Write([]byte(line))
	return nil
}

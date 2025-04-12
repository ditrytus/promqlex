package promqlex

import (
	_ "embed"
	"encoding/json"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"slices"
	"testing"
)

type ExpectedSymbol struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Ary  int    `yaml:"ary,omitempty"`
}

func (s ExpectedSymbol) MarshalJSON() ([]byte, error) {
	fields := map[string]any{
		"name": s.Name,
		"type": s.Type,
	}
	if s.Type == "macro" {
		fields["ary"] = s.Ary
	}
	return json.Marshal(fields)
}

type ExpectedScope struct {
	Symbols  map[string]ExpectedSymbol `yaml:"symbols,omitempty" json:"symbols,omitempty"`
	Children []ExpectedScope           `yaml:"children,omitempty" json:"children,omitempty"`
}

type SymbolDefinerTestCase struct {
	Name     string
	PromQLEx string
	Errors   []string
	Scope    ExpectedScope
}

//go:embed testdata/symbols.yaml
var symbolsYaml string

func TestSymbolDefiner(t *testing.T) {
	var testCases []SymbolDefinerTestCase
	err := yaml.Unmarshal([]byte(symbolsYaml), &testCases)
	if !assert.NoError(t, err) {
		assert.FailNow(t, err.Error())
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			_, tree, asciiTree := parsePromQLEx(t, testCase.PromQLEx)
			if t.Failed() {
				t.Log(asciiTree)
				return
			}
			symbolsDefiner := NewSymbolsDefiner()
			walker := antlr.NewParseTreeWalker()
			walker.Walk(symbolsDefiner, tree)
			if len(testCase.Errors) > 0 {
				assert.True(t, symbolsDefiner.HasErrors())
				assert.Len(t, symbolsDefiner.Errors(), len(testCase.Errors))
				for _, expectedError := range testCase.Errors {
					assert.True(
						t,
						slices.ContainsFunc(
							symbolsDefiner.Errors(),
							func(err *ParseTreeError) bool {
								return err.Error() == expectedError
							},
						),
					)
				}
				return
			}
			assert.Equal(t, symbolsDefiner.GlobalScope(), symbolsDefiner.Scopes()[tree.Statement_block()])
			for parseTree, scope := range symbolsDefiner.Scopes() {
				var pathFromScopes []Scope
				for currentScope := scope; currentScope != nil; currentScope = currentScope.Parent() {
					pathFromScopes = append(pathFromScopes, currentScope)
				}
				var pathFromAST []Scope
				for currentNode := parseTree; currentNode != nil; currentNode = castOrDefault[antlr.ParseTree](currentNode.GetParent()) {
					if parentScope, ok := symbolsDefiner.Scopes()[currentNode]; ok {
						pathFromAST = append(pathFromAST, parentScope)
					}
				}
				assert.Equal(t, pathFromScopes, pathFromAST)
			}
			actualSymbols, err := json.MarshalIndent(symbolsDefiner.GlobalScope(), "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			expectedSymbols, err := json.MarshalIndent(testCase.Scope, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			if !assert.JSONEq(t, string(expectedSymbols), string(actualSymbols)) {
				t.Log(string(expectedSymbols))
				t.Log(string(actualSymbols))
			}
		})
	}
}

func castOrDefault[T any](v interface{}) T {
	c, _ := v.(T)
	return c
}

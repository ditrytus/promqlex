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

type SymbolDefinerTestCase struct {
	Name     string
	PromQLEx string
	Count    int
	Errors   []string
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
			marshal, err := json.Marshal(symbolsDefiner.GlobalScope())
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(marshal))
		})
	}
}

func castOrDefault[T any](v interface{}) T {
	c, _ := v.(T)
	return c
}

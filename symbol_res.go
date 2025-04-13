package promqlex

import (
	"fmt"
	"github.com/ditrytus/promqlex/parsers/promqlex"
	"gonum.org/v1/gonum/graph/simple"
)

type ErrorSymbolUndefined struct {
	name string
}

func NewErrorSymbolUndefined(name string) *ErrorSymbolUndefined {
	return &ErrorSymbolUndefined{
		name: name,
	}
}

func (e *ErrorSymbolUndefined) Error() string {
	return fmt.Sprintf("symbol '%s' is undefined", e.name)
}

type DependenyGraph interface {
}

type gonumDeps struct {
	graph simple.DirectedGraph
}

type SymbolsResolver struct {
	promqlex.BasePromQLExParserListener
	ParseTreeErrorCollector
	scopes ParseTreeProperty[Scope]
	subs   ParseTreeProperty[Symbol]
	deps   DependenyGraph
}

var _ promqlex.PromQLExParserListener = &SymbolsResolver{}

func NewSymbolsResolver(scopes ParseTreeProperty[Scope]) *SymbolsResolver {
	return &SymbolsResolver{
		scopes: scopes,
	}
}

func (s *SymbolsResolver) EnterSubstitute(ctx *promqlex.SubstituteContext) {
	scope, _ := s.scopes.Get(ctx)
	if scope == nil {
		panic("no scope found, probably walking wrong tree")
	}
	name := ctx.ID().GetText()
	symbol := scope.Resolve(name)
	if symbol == nil {
		s.addError(NewErrorSymbolUndefined(name), ctx.ID())
		return
	}
	s.subs[ctx] = symbol

	parameterList := ctx.ParameterList()
	switch sym := symbol.(type) {
	case *MacroSymbol:
		if parameterList == nil {
			s.addError(
				fmt.Errorf("macro %s is missing parameter list", sym.Name()), ctx)
			return
		}
		params := parameterList.AllParameter()
		if len(params) != sym.Ary() {
			s.addError(
				fmt.Errorf(
					"macro %s expects %d parameters, but %d were provided",
					sym.Name(),
					sym.Ary(),
					len(params),
				), parameterList)
			return
		}
	case *AliasSymbol:
		if parameterList != nil {
			s.addError(
				fmt.Errorf("alias %s should not have parameters", sym.Name()), parameterList)
			return
		}
	}
}

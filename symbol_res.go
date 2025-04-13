package promqlex

import (
	"fmt"
	"github.com/ditrytus/promqlex/parsers/promqlex"
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

type SymbolsResolver struct {
	promqlex.BasePromQLExParserListener
	ParseTreeErrorCollector
	scopes ParseTreeProperty[Scope]
}

var _ promqlex.PromQLExParserListener = &SymbolsResolver{}

func NewSymbolsResolver(scopes ParseTreeProperty[Scope]) *SymbolsResolver {
	return &SymbolsResolver{
		scopes: scopes,
	}
}

func (s *SymbolsResolver) EnterSubstitute(ctx *promqlex.SubstituteContext) {
	scope := s.scopes.Get(ctx)
	if scope == nil {
		panic("no scope found, probably walking wrong tree")
	}
	name := ctx.ID().GetText()
	symbol := scope.Resolve(name)
	if symbol == nil {
		s.addError(NewErrorSymbolUndefined(name), ctx.ID())
	}
}

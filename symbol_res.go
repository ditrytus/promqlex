package promqlex

import "github.com/ditrytus/promqlex/parsers/promqlex"

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

package promqlex

import "fmt"

type ErrorSymbolAlreadyDefined struct {
	otherDef Symbol
}

func (e *ErrorSymbolAlreadyDefined) Error() string {
	return fmt.Sprintf("symbol '%s' is already defined in this scope", e.otherDef.Name())
}

type Scope struct {
	parent  *Scope
	symbols map[string]Symbol
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent:  parent,
		symbols: make(map[string]Symbol),
	}
}

func NewErrorSymbolAlreadyDefined(otherDef Symbol) *ErrorSymbolAlreadyDefined {
	return &ErrorSymbolAlreadyDefined{
		otherDef: otherDef,
	}
}

func (s *Scope) Define(symbol Symbol) error {
	if otherDef, ok := s.symbols[symbol.Name()]; ok {
		return NewErrorSymbolAlreadyDefined(otherDef)
	}
	s.symbols[symbol.Name()] = symbol
	return nil
}

func (s *Scope) Resolve(name string) Symbol {
	if s == nil {
		return nil
	}
	if symbol, ok := s.symbols[name]; ok {
		return symbol
	}
	return s.Parent().Resolve(name)
}

func (s *Scope) Parent() *Scope {
	if s == nil {
		return nil
	}
	return s.parent
}

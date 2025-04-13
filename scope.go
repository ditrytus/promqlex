package promqlex

import (
	"encoding/json"
	"fmt"
)

type ErrorSymbolAlreadyDefined struct {
	otherDef Symbol
}

func NewErrorSymbolAlreadyDefined(otherDef Symbol) *ErrorSymbolAlreadyDefined {
	return &ErrorSymbolAlreadyDefined{
		otherDef: otherDef,
	}
}

func (e *ErrorSymbolAlreadyDefined) Error() string {
	return fmt.Sprintf("symbol '%s' is already defined in this scope", e.otherDef.Name())
}

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

type Scope interface {
	json.Marshaler
	Define(symbol Symbol) error
	Resolve(name string) Symbol
	Parent() Scope
	Children() []Scope
}

type scopeImpl struct {
	parent   Scope
	symbols  map[string]Symbol
	children []Scope
}

func NewScope(parent Scope) Scope {
	newScope := &scopeImpl{
		parent:   parent,
		symbols:  make(map[string]Symbol),
		children: nil,
	}
	if impl, ok := parent.(*scopeImpl); ok {
		impl.addChild(newScope)
	}
	return newScope
}

func (s *scopeImpl) Define(symbol Symbol) error {
	if otherDef, ok := s.symbols[symbol.Name()]; ok {
		return NewErrorSymbolAlreadyDefined(otherDef)
	}
	s.symbols[symbol.Name()] = symbol
	return nil
}

func (s *scopeImpl) Resolve(name string) Symbol {
	if s == nil {
		return nil
	}
	if symbol, ok := s.symbols[name]; ok {
		return symbol
	}
	return s.Parent().Resolve(name)
}

func (s *scopeImpl) Parent() Scope {
	if s == nil {
		return nil
	}
	return s.parent
}

func (s *scopeImpl) Children() []Scope {
	return s.children
}

func (s *scopeImpl) addChild(scope *scopeImpl) {
	if s == nil {
		return
	}
	s.children = append(s.children, scope)
}

func (s *scopeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Symbols  map[string]Symbol `json:"symbols,omitempty"`
		Children []Scope           `json:"children,omitempty"`
	}{
		Symbols:  s.symbols,
		Children: s.children,
	})
}

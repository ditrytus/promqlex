package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
)

type ParseTreeErrorCollector struct {
	errors []*ParseTreeError
}

func (c *ParseTreeErrorCollector) addError(err error, site antlr.ParseTree) {
	c.errors = append(c.errors, NewParseTreeError(err, site))
}

type SymbolsDefiner struct {
	promqlex.BasePromQLExParserListener
	ParseTreeErrorCollector
	scopes       ParseTreeProperty[Scope]
	globalScope  Scope
	currentScope Scope
}

func NewSymbolsDefiner() *SymbolsDefiner {
	return &SymbolsDefiner{
		scopes: make(map[antlr.ParseTree]Scope),
	}
}

var _ promqlex.PromQLExParserListener = &SymbolsDefiner{}

func (m *SymbolsDefiner) EnterAlias_def(c *promqlex.Alias_defContext) {
	m.defineAliasSymbol(c.ID())
}

func (m *SymbolsDefiner) EnterMacro_def(c *promqlex.Macro_defContext) {
	var ary int
	if c.Args_decl() != nil {
		ary = len(c.Args_decl().AllArg_name())
	}
	err := m.currentScope.Define(NewMacroSymbol(c.ID().GetText(), ary))
	if err != nil {
		m.addError(err, c.ID())
	}
	m.enterNewScope(c)
}

func (m *SymbolsDefiner) EnterArg_name(c *promqlex.Arg_nameContext) {
	m.defineAliasSymbol(c.ID())
}

func (m *SymbolsDefiner) EnterStatement_block(c *promqlex.Statement_blockContext) {
	m.enterNewScope(c)
}

func (m *SymbolsDefiner) ExitMacro_def(c *promqlex.Macro_defContext) {
	m.exitScope()
}

func (m *SymbolsDefiner) ExitStatement_block(c *promqlex.Statement_blockContext) {
	m.exitScope()
}

func (m *SymbolsDefiner) HasErrors() bool {
	return len(m.errors) > 0
}

func (m *SymbolsDefiner) Errors() []*ParseTreeError {
	return m.errors
}

func (m *SymbolsDefiner) Scopes() ParseTreeProperty[Scope] {
	return m.scopes
}

func (m *SymbolsDefiner) GlobalScope() Scope {
	return m.globalScope
}

func (m *SymbolsDefiner) enterNewScope(c antlr.ParseTree) {
	newScope := NewScope(m.currentScope)
	m.currentScope = newScope
	if m.globalScope == nil {
		m.globalScope = newScope
	}
	m.scopes[c] = newScope
}

func (m *SymbolsDefiner) defineAliasSymbol(idTerminal antlr.TerminalNode) {
	err := m.currentScope.Define(NewAliasSymbol(idTerminal.GetText()))
	if err != nil {
		m.addError(err, idTerminal)
	}
}

func (m *SymbolsDefiner) exitScope() {
	m.currentScope = m.currentScope.Parent()
}

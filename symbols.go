package promqlex

import (
	"encoding/json"
	"github.com/ditrytus/promqlex/parsers/promqlex"
)

type Symbol interface {
	json.Marshaler
	Name() string
}

type AliasSymbol struct {
	name string
	def  promqlex.IAlias_defContext
}

func (v *AliasSymbol) Name() string {
	return v.name
}

func (v *AliasSymbol) Def() promqlex.IAlias_defContext {
	return v.def
}

func NewAliasSymbol(name string, def promqlex.IAlias_defContext) *AliasSymbol {
	return &AliasSymbol{
		name: name,
		def:  def,
	}
}

func (v *AliasSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name": v.name,
		"type": "alias",
	})
}

type ArgSymbol struct {
	name string
	def  promqlex.IArg_nameContext
}

func (v *ArgSymbol) Name() string {
	return v.name
}

func (v *ArgSymbol) Def() promqlex.IArg_nameContext {
	return v.def
}

func NewArgSymbol(name string, def promqlex.IArg_nameContext) *ArgSymbol {
	return &ArgSymbol{
		name: name,
		def:  def,
	}
}

func (v *ArgSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name": v.name,
		"type": "arg",
	})
}

type MacroSymbol struct {
	name string
	ary  int
	def  promqlex.IMacro_defContext
}

func NewMacroSymbol(name string, ary int, def promqlex.IMacro_defContext) *MacroSymbol {
	return &MacroSymbol{
		name: name,
		ary:  ary,
		def:  def,
	}
}

func (v *MacroSymbol) Name() string {
	return v.name
}

func (v *MacroSymbol) Ary() int {
	return v.ary
}

func (v *MacroSymbol) Def() promqlex.IMacro_defContext {
	return v.def
}

func (v *MacroSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name": v.name,
		"type": "macro",
		"ary":  v.ary,
	})
}

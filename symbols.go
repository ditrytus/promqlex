package promqlex

import "encoding/json"

type Symbol interface {
	json.Marshaler
	Name() string
}

type AliasSymbol struct {
	name string
}

func (v *AliasSymbol) Name() string {
	return v.name
}

func NewAliasSymbol(name string) *AliasSymbol {
	return &AliasSymbol{name: name}
}

func (v *AliasSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name": v.name,
		"type": "alias",
	})
}

type MacroSymbol struct {
	name string
	ary  int
}

func (v *MacroSymbol) Name() string {
	return v.name
}

func (v *MacroSymbol) Ary() int {
	return v.ary
}

func NewMacroSymbol(name string, ary int) *MacroSymbol {
	return &MacroSymbol{name: name, ary: ary}
}

func (v *MacroSymbol) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"name": v.name,
		"type": "macro",
		"ary":  v.ary,
	})
}

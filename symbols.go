package promqlex

type Symbol interface {
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

type MacroSymbol struct {
	name string
	ary  int
}

func (v *MacroSymbol) Name() string {
	return v.name
}

func NewMacroSymbol(name string, ary int) *MacroSymbol {
	return &MacroSymbol{name: name, ary: ary}
}

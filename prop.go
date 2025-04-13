package promqlex

import "github.com/antlr4-go/antlr/v4"

type ParseTreeProperty[T any] map[antlr.ParseTree]T

func (p ParseTreeProperty[T]) Get(tree antlr.ParseTree) (T, antlr.ParseTree) {
	for {
		prop, ok := p[tree]
		if ok {
			return prop, tree
		}
		tree = tree.GetParent().(antlr.ParseTree)
		if tree == nil {
			var empty T
			return empty, nil
		}
	}
}

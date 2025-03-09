package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
)

type FunctionProviderInputStream struct {
	antlr.CharStream
	FunctionsSet
}

func NewFunctionProviderInputStream(input antlr.CharStream, functionsSet FunctionsSet) *FunctionProviderInputStream {
	return &FunctionProviderInputStream{
		CharStream:   input,
		FunctionsSet: functionsSet,
	}
}

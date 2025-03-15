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

type InputStream struct {
	FunctionProviderInputStream
	bracketCounter int
}

func NewInputStream(input antlr.CharStream, functionsSet FunctionsSet) *InputStream {
	return &InputStream{
		FunctionProviderInputStream: FunctionProviderInputStream{
			CharStream:   input,
			FunctionsSet: functionsSet,
		},
		bracketCounter: 0,
	}
}

func (i *InputStream) BracketCount() int {
	return i.bracketCounter
}

func (i *InputStream) SetBracketCount(val int) {
	i.bracketCounter = val
}

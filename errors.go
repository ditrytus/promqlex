package promqlex

import "github.com/antlr4-go/antlr/v4"

type ParseTreeError struct {
	innerErr error
	location antlr.ParseTree
}

func NewParseTreeError(innerErr error, location antlr.ParseTree) *ParseTreeError {
	return &ParseTreeError{
		innerErr: innerErr,
		location: location,
	}
}

func (f *ParseTreeError) Error() string {
	return f.innerErr.Error()
}

func (f *ParseTreeError) Location() antlr.ParseTree {
	return f.location
}

func (f *ParseTreeError) InnerError() error {
	return f.innerErr
}

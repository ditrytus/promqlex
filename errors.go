package promqlex

import "github.com/antlr4-go/antlr/v4"

type ParseTreeError struct {
	innerErr error
	site     antlr.ParseTree
}

func NewParseTreeError(innerErr error, site antlr.ParseTree) *ParseTreeError {
	return &ParseTreeError{
		innerErr: innerErr,
		site:     site,
	}
}

func (f *ParseTreeError) Error() string {
	return f.innerErr.Error()
}

func (f *ParseTreeError) Location() antlr.ParseTree {
	return f.site
}

func (f *ParseTreeError) InnerError() error {
	return f.innerErr
}

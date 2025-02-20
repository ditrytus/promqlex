ANTLR4 = antlr4

.PHONY: all
all: parsers/promql_parser.go parsers/promqlextension_parser.go

define build_lexer_and_parser

$(eval GRAMMAR_NAME := $(1))
$(eval LOWERCASE_GRAMMAR_NAME := $(shell echo $(GRAMMAR_NAME) | tr '[:upper:]' '[:lower:]'))

$(eval LEXER_OUTPUT := grammars/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.interp $(LOWERCASE_GRAMMAR_NAME)_lexer.go)
$(LEXER_OUTPUT): grammars/$(GRAMMAR_NAME)Lexer.g4
	cd grammars && $(ANTLR4) -Dlanguage=Go -o ../parsers $(GRAMMAR_NAME)Lexer.g4
	mv parsers/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.tokens
	mv parsers/$(GRAMMAR_NAME)Lexer.interp grammars/$(GRAMMAR_NAME)Lexer.interp
	rm -frd gen
	rm -frd grammars/gen

$(eval PARSER_OUTPUT := parsers/$(LOWERCASE_GRAMMAR_NAME)_parser.go parsers/$(LOWERCASE_GRAMMAR_NAME)parser_listener.go parsers/$(LOWERCASE_GRAMMAR_NAME)parser_base_listener.go)
$(PARSER_OUTPUT): grammars/$(GRAMMAR_NAME)Parser.g4 grammars/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.interp
	cd grammars && $(ANTLR4) -Dlanguage=Go -o ../parsers $(GRAMMAR_NAME)Parser.g4
	rm -f parsers/$(GRAMMAR_NAME)Parser.tokens
	rm -f parsers/$(GRAMMAR_NAME)Parser.interp
endef

$(eval $(call build_lexer_and_parser,PromQL))
$(eval $(call build_lexer_and_parser,PromQLExtension))
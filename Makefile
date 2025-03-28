MAKEFLAGS += --no-builtin-rules

ANTLR4 = antlr4
ANTLR4_OUTPUT_DIR = parsers/$(LOWERCASE_GRAMMAR_NAME)
ANTLR4_ARGS = -Dlanguage=Go -o ../$(ANTLR4_OUTPUT_DIR) -package $(LOWERCASE_GRAMMAR_NAME)
BUILD_DIR = bin
OUT_LIB_NAME = promqlex
OUT_FULL_PATH = $(BUILD_DIR)/$(OUT_LIB_NAME)

.PHONY: grammars
grammars: $(wildcard parsers/promql/*.go) $(wildcard parsers/promqlex/*.go)

define build_lexer_and_parser

$(eval GRAMMAR_NAME := $(1))
$(eval LOWERCASE_GRAMMAR_NAME := $(shell echo $(GRAMMAR_NAME) | tr '[:upper:]' '[:lower:]'))

$(eval LEXER_OUTPUT := grammars/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.interp $(ANTLR4_OUTPUT_DIR)/$(LOWERCASE_GRAMMAR_NAME)_lexer.go)
GENERATED_PARSERS += $(LEXER_OUTPUT)
$(LEXER_OUTPUT): grammars/$(GRAMMAR_NAME)Lexer.g4
	cd grammars && $(ANTLR4) $(ANTLR4_ARGS) $(GRAMMAR_NAME)Lexer.g4
	mv $(ANTLR4_OUTPUT_DIR)/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.tokens
	mv $(ANTLR4_OUTPUT_DIR)/$(GRAMMAR_NAME)Lexer.interp grammars/$(GRAMMAR_NAME)Lexer.interp
	rm -frd gen
	rm -frd grammars/gen

$(eval PARSER_OUTPUT := $(ANTLR4_OUTPUT_DIR)/$(LOWERCASE_GRAMMAR_NAME)_parser.go $(ANTLR4_OUTPUT_DIR)/$(LOWERCASE_GRAMMAR_NAME)parser_listener.go $(ANTLR4_OUTPUT_DIR)/$(LOWERCASE_GRAMMAR_NAME)parser_base_listener.go)
GENERATED_PARSERS += $(PARSER_OUTPUT)
$(PARSER_OUTPUT): grammars/$(GRAMMAR_NAME)Parser.g4 grammars/$(GRAMMAR_NAME)Lexer.tokens grammars/$(GRAMMAR_NAME)Lexer.interp
	cd grammars && $(ANTLR4) $(ANTLR4_ARGS) $(GRAMMAR_NAME)Parser.g4
	rm -f $(ANTLR4_OUTPUT_DIR)/$(GRAMMAR_NAME)Parser.tokens
	rm -f $(ANTLR4_OUTPUT_DIR)/$(GRAMMAR_NAME)Parser.interp
endef

$(eval $(call build_lexer_and_parser,PromQL))
$(eval $(call build_lexer_and_parser,PromQLEx))

.PHONY: test
test:
	go test ./...

$(OUT_FULL_PATH): $(GENERATED_PARSERS) $(wildcard **/*.go)
	go build -o $@ .

.PHONY: build
build: $(OUT_FULL_PATH)
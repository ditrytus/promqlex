ANTLR4 = antlr4

LEXER_OUTPUT = grammar/PromQLLexer.tokens grammar/PromQLLexer.interp promql_lexer.go
$(LEXER_OUTPUT): grammar/PromQLLexer.g4
	cd grammar && $(ANTLR4) -Dlanguage=Go -o ../parser PromQLLexer.g4
	mv parser/PromQLLexer.tokens grammar/PromQLLexer.tokens
	mv parser/PromQLLexer.interp grammar/PromQLLexer.interp

PARSER_OUTPUT = parser/promql_parser.go parser/promqlparser_listener.go parser/promqlparser_base_listener.go
$(PARSER_OUTPUT): grammar/PromQLParser.g4 grammar/PromQLLexer.tokens grammar/PromQLLexer.interp
	cd grammar && $(ANTLR4) -Dlanguage=Go -o ../parser PromQLParser.g4

.PHONY: parser
parser: $(PARSER_OUTPUT)
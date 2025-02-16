ANTLR4 = antlr4

all: $(PARSER_OUTPUT) $(LEXER_OUTPUT)

LEXER_OUTPUT = grammars/PromQLLexer.tokens grammars/PromQLLexer.interp promql_lexer.go
$(LEXER_OUTPUT): grammars/PromQLLexer.g4
	cd grammars && $(ANTLR4) -Dlanguage=Go -o ../parser PromQLLexer.g4
	mv parser/PromQLLexer.tokens grammars/PromQLLexer.tokens
	mv parser/PromQLLexer.interp grammars/PromQLLexer.interp

PARSER_OUTPUT = parser/promql_parser.go parser/promqlparser_listener.go parser/promqlparser_base_listener.go
$(PARSER_OUTPUT): grammars/PromQLParser.g4 grammars/PromQLLexer.tokens grammars/PromQLLexer.interp
	cd grammars && $(ANTLR4) -Dlanguage=Go -o ../parser PromQLParser.g4
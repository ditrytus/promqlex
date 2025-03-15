// Code generated from PromQLExLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

import (
	"slices"
	"strings"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PromQLExLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PromQLExLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promqlexlexerLexerInit() {
	staticData := &PromQLExLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "WHITESPACE", "COMMENTS",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "ID_MODE", "RAW_STRING_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'if'", "'true'", "'false'", "'T'", "':'", "';'", "'.'",
		"", "", "", "'metric'", "'label'", "'def'", "'$'", "'['", "']'", "",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "METRIC_NAME", "IF", "TRUE",
		"FALSE", "T", "COLON", "SEMICOLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS",
		"DIGITS", "METRIC_KEYWORD", "LABEL_KEYWORD", "DEF", "CALL_SIGN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD",
		"POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE",
		"RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT",
		"OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN",
		"COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "ID", "RAW_STRING",
		"BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"SUBQUERY_RANGE", "TIME_RANGE", "METRIC_NAME", "IF", "TRUE", "FALSE",
		"T", "COLON", "SEMICOLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS",
		"DIGITS", "METRIC_KEYWORD", "LABEL_KEYWORD", "DEF", "CALL_SIGN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "NUMERAL", "SCIENTIFIC_NUMBER", "NUMBER", "STRING",
		"BACKTICK_OPEN", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR",
		"UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE", "BY",
		"WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET",
		"BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN", "COMMA",
		"AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "ID", "RAW_STRING_CONTENT",
		"RAW_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 60, 450, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
		2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8,
		2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2,
		14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19,
		7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7,
		24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29,
		2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2,
		35, 7, 35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40,
		7, 40, 2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7,
		45, 2, 46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50,
		2, 51, 7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2,
		56, 7, 56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61,
		7, 61, 2, 62, 7, 62, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 137,
		8, 2, 10, 2, 12, 2, 140, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 4, 10, 169, 8, 10, 11,
		10, 12, 10, 170, 1, 10, 5, 10, 174, 8, 10, 10, 10, 12, 10, 177, 9, 10,
		3, 10, 179, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 185, 8, 12, 11, 12,
		12, 12, 186, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 4, 19, 219, 8, 19, 11, 19, 12, 19, 220, 1, 19, 1, 19, 4, 19, 225, 8,
		19, 11, 19, 12, 19, 226, 3, 19, 229, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20,
		234, 8, 20, 1, 20, 3, 20, 237, 8, 20, 1, 21, 1, 21, 3, 21, 241, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 247, 8, 22, 10, 22, 12, 22, 250, 9,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 257, 8, 22, 10, 22, 12, 22,
		260, 9, 22, 1, 22, 3, 22, 263, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 4, 56, 391, 8, 56, 11, 56, 12, 56,
		392, 1, 56, 1, 56, 1, 56, 3, 56, 398, 8, 56, 4, 56, 400, 8, 56, 11, 56,
		12, 56, 401, 1, 57, 1, 57, 5, 57, 406, 8, 57, 10, 57, 12, 57, 409, 9, 57,
		1, 58, 4, 58, 412, 8, 58, 11, 58, 12, 58, 413, 1, 58, 1, 58, 1, 59, 1,
		59, 5, 59, 420, 8, 59, 10, 59, 12, 59, 423, 9, 59, 1, 59, 1, 59, 1, 59,
		1, 59, 1, 60, 1, 60, 5, 60, 431, 8, 60, 10, 60, 12, 60, 434, 9, 60, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 61, 4, 61, 441, 8, 61, 11, 61, 12, 61, 442,
		1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 421, 0, 63, 3, 0, 5, 0, 7,
		3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27,
		13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 0, 43, 0, 45, 20,
		47, 21, 49, 60, 51, 22, 53, 23, 55, 24, 57, 25, 59, 26, 61, 27, 63, 28,
		65, 29, 67, 30, 69, 31, 71, 32, 73, 33, 75, 34, 77, 35, 79, 36, 81, 37,
		83, 38, 85, 39, 87, 40, 89, 41, 91, 42, 93, 43, 95, 44, 97, 45, 99, 46,
		101, 47, 103, 48, 105, 49, 107, 50, 109, 51, 111, 52, 113, 53, 115, 54,
		117, 55, 119, 56, 121, 57, 123, 58, 125, 0, 127, 59, 3, 0, 1, 2, 33, 4,
		0, 58, 58, 65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90, 95, 95, 97, 122,
		2, 0, 73, 73, 105, 105, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84, 116, 116,
		2, 0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 69, 69, 101, 101,
		2, 0, 65, 65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 83, 83, 115, 115, 1,
		0, 48, 57, 1, 0, 49, 57, 2, 0, 77, 77, 109, 109, 2, 0, 67, 67, 99, 99,
		2, 0, 66, 66, 98, 98, 2, 0, 68, 68, 100, 100, 2, 0, 43, 43, 45, 45, 2,
		0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 2, 0, 78, 78, 110, 110, 2, 0,
		79, 79, 111, 111, 2, 0, 89, 89, 121, 121, 2, 0, 87, 87, 119, 119, 2, 0,
		72, 72, 104, 104, 2, 0, 71, 71, 103, 103, 2, 0, 80, 80, 112, 112, 12, 0,
		68, 68, 72, 72, 77, 77, 83, 83, 87, 87, 89, 89, 100, 100, 104, 104, 109,
		109, 115, 115, 119, 119, 121, 121, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65,
		90, 97, 122, 2, 0, 92, 92, 96, 96, 468, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83,
		1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0,
		91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0,
		0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0,
		0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113,
		1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0,
		0, 121, 1, 0, 0, 0, 1, 123, 1, 0, 0, 0, 2, 125, 1, 0, 0, 0, 2, 127, 1,
		0, 0, 0, 3, 129, 1, 0, 0, 0, 5, 131, 1, 0, 0, 0, 7, 133, 1, 0, 0, 0, 9,
		144, 1, 0, 0, 0, 11, 147, 1, 0, 0, 0, 13, 152, 1, 0, 0, 0, 15, 158, 1,
		0, 0, 0, 17, 160, 1, 0, 0, 0, 19, 162, 1, 0, 0, 0, 21, 164, 1, 0, 0, 0,
		23, 178, 1, 0, 0, 0, 25, 180, 1, 0, 0, 0, 27, 184, 1, 0, 0, 0, 29, 188,
		1, 0, 0, 0, 31, 195, 1, 0, 0, 0, 33, 201, 1, 0, 0, 0, 35, 207, 1, 0, 0,
		0, 37, 211, 1, 0, 0, 0, 39, 214, 1, 0, 0, 0, 41, 218, 1, 0, 0, 0, 43, 230,
		1, 0, 0, 0, 45, 240, 1, 0, 0, 0, 47, 262, 1, 0, 0, 0, 49, 264, 1, 0, 0,
		0, 51, 269, 1, 0, 0, 0, 53, 271, 1, 0, 0, 0, 55, 273, 1, 0, 0, 0, 57, 275,
		1, 0, 0, 0, 59, 277, 1, 0, 0, 0, 61, 279, 1, 0, 0, 0, 63, 281, 1, 0, 0,
		0, 65, 285, 1, 0, 0, 0, 67, 288, 1, 0, 0, 0, 69, 295, 1, 0, 0, 0, 71, 297,
		1, 0, 0, 0, 73, 300, 1, 0, 0, 0, 75, 303, 1, 0, 0, 0, 77, 305, 1, 0, 0,
		0, 79, 307, 1, 0, 0, 0, 81, 310, 1, 0, 0, 0, 83, 313, 1, 0, 0, 0, 85, 316,
		1, 0, 0, 0, 87, 319, 1, 0, 0, 0, 89, 322, 1, 0, 0, 0, 91, 330, 1, 0, 0,
		0, 93, 333, 1, 0, 0, 0, 95, 342, 1, 0, 0, 0, 97, 353, 1, 0, 0, 0, 99, 365,
		1, 0, 0, 0, 101, 372, 1, 0, 0, 0, 103, 377, 1, 0, 0, 0, 105, 379, 1, 0,
		0, 0, 107, 381, 1, 0, 0, 0, 109, 383, 1, 0, 0, 0, 111, 385, 1, 0, 0, 0,
		113, 387, 1, 0, 0, 0, 115, 399, 1, 0, 0, 0, 117, 403, 1, 0, 0, 0, 119,
		411, 1, 0, 0, 0, 121, 417, 1, 0, 0, 0, 123, 428, 1, 0, 0, 0, 125, 440,
		1, 0, 0, 0, 127, 446, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 4, 1, 0, 0,
		0, 131, 132, 1, 0, 0, 0, 132, 6, 1, 0, 0, 0, 133, 134, 4, 2, 0, 0, 134,
		138, 7, 0, 0, 0, 135, 137, 7, 1, 0, 0, 136, 135, 1, 0, 0, 0, 137, 140,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0,
		0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 4, 2, 1, 0, 142, 143, 6, 2, 0, 0,
		143, 8, 1, 0, 0, 0, 144, 145, 7, 2, 0, 0, 145, 146, 7, 3, 0, 0, 146, 10,
		1, 0, 0, 0, 147, 148, 7, 4, 0, 0, 148, 149, 7, 5, 0, 0, 149, 150, 7, 6,
		0, 0, 150, 151, 7, 7, 0, 0, 151, 12, 1, 0, 0, 0, 152, 153, 7, 3, 0, 0,
		153, 154, 7, 8, 0, 0, 154, 155, 7, 9, 0, 0, 155, 156, 7, 10, 0, 0, 156,
		157, 7, 7, 0, 0, 157, 14, 1, 0, 0, 0, 158, 159, 7, 4, 0, 0, 159, 16, 1,
		0, 0, 0, 160, 161, 5, 58, 0, 0, 161, 18, 1, 0, 0, 0, 162, 163, 5, 59, 0,
		0, 163, 20, 1, 0, 0, 0, 164, 165, 5, 46, 0, 0, 165, 22, 1, 0, 0, 0, 166,
		179, 7, 11, 0, 0, 167, 169, 7, 12, 0, 0, 168, 167, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 175, 1, 0,
		0, 0, 172, 174, 7, 11, 0, 0, 173, 172, 1, 0, 0, 0, 174, 177, 1, 0, 0, 0,
		175, 173, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177,
		175, 1, 0, 0, 0, 178, 166, 1, 0, 0, 0, 178, 168, 1, 0, 0, 0, 179, 24, 1,
		0, 0, 0, 180, 181, 7, 11, 0, 0, 181, 182, 7, 11, 0, 0, 182, 26, 1, 0, 0,
		0, 183, 185, 7, 11, 0, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 28, 1, 0, 0, 0, 188, 189, 7,
		13, 0, 0, 189, 190, 7, 7, 0, 0, 190, 191, 7, 4, 0, 0, 191, 192, 7, 5, 0,
		0, 192, 193, 7, 2, 0, 0, 193, 194, 7, 14, 0, 0, 194, 30, 1, 0, 0, 0, 195,
		196, 7, 9, 0, 0, 196, 197, 7, 8, 0, 0, 197, 198, 7, 15, 0, 0, 198, 199,
		7, 7, 0, 0, 199, 200, 7, 9, 0, 0, 200, 32, 1, 0, 0, 0, 201, 202, 7, 16,
		0, 0, 202, 203, 7, 7, 0, 0, 203, 204, 7, 3, 0, 0, 204, 205, 1, 0, 0, 0,
		205, 206, 6, 15, 1, 0, 206, 34, 1, 0, 0, 0, 207, 208, 5, 36, 0, 0, 208,
		209, 1, 0, 0, 0, 209, 210, 6, 16, 1, 0, 210, 36, 1, 0, 0, 0, 211, 212,
		5, 91, 0, 0, 212, 213, 6, 17, 2, 0, 213, 38, 1, 0, 0, 0, 214, 215, 5, 93,
		0, 0, 215, 216, 6, 18, 3, 0, 216, 40, 1, 0, 0, 0, 217, 219, 7, 11, 0, 0,
		218, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 220,
		221, 1, 0, 0, 0, 221, 228, 1, 0, 0, 0, 222, 224, 5, 46, 0, 0, 223, 225,
		7, 11, 0, 0, 224, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0,
		0, 0, 226, 227, 1, 0, 0, 0, 227, 229, 1, 0, 0, 0, 228, 222, 1, 0, 0, 0,
		228, 229, 1, 0, 0, 0, 229, 42, 1, 0, 0, 0, 230, 236, 3, 41, 19, 0, 231,
		233, 7, 7, 0, 0, 232, 234, 7, 17, 0, 0, 233, 232, 1, 0, 0, 0, 233, 234,
		1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 3, 41, 19, 0, 236, 231, 1,
		0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 44, 1, 0, 0, 0, 238, 241, 3, 41, 19,
		0, 239, 241, 3, 43, 20, 0, 240, 238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0,
		241, 46, 1, 0, 0, 0, 242, 248, 5, 39, 0, 0, 243, 247, 8, 18, 0, 0, 244,
		245, 5, 92, 0, 0, 245, 247, 9, 0, 0, 0, 246, 243, 1, 0, 0, 0, 246, 244,
		1, 0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0,
		0, 0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 263, 5, 39, 0, 0,
		252, 258, 5, 34, 0, 0, 253, 257, 8, 19, 0, 0, 254, 255, 5, 92, 0, 0, 255,
		257, 9, 0, 0, 0, 256, 253, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 260,
		1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261, 1, 0,
		0, 0, 260, 258, 1, 0, 0, 0, 261, 263, 5, 34, 0, 0, 262, 242, 1, 0, 0, 0,
		262, 252, 1, 0, 0, 0, 263, 48, 1, 0, 0, 0, 264, 265, 5, 96, 0, 0, 265,
		266, 1, 0, 0, 0, 266, 267, 6, 23, 4, 0, 267, 268, 6, 23, 5, 0, 268, 50,
		1, 0, 0, 0, 269, 270, 5, 43, 0, 0, 270, 52, 1, 0, 0, 0, 271, 272, 5, 45,
		0, 0, 272, 54, 1, 0, 0, 0, 273, 274, 5, 42, 0, 0, 274, 56, 1, 0, 0, 0,
		275, 276, 5, 47, 0, 0, 276, 58, 1, 0, 0, 0, 277, 278, 5, 37, 0, 0, 278,
		60, 1, 0, 0, 0, 279, 280, 5, 94, 0, 0, 280, 62, 1, 0, 0, 0, 281, 282, 7,
		8, 0, 0, 282, 283, 7, 20, 0, 0, 283, 284, 7, 16, 0, 0, 284, 64, 1, 0, 0,
		0, 285, 286, 7, 21, 0, 0, 286, 287, 7, 5, 0, 0, 287, 66, 1, 0, 0, 0, 288,
		289, 7, 6, 0, 0, 289, 290, 7, 20, 0, 0, 290, 291, 7, 9, 0, 0, 291, 292,
		7, 7, 0, 0, 292, 293, 7, 10, 0, 0, 293, 294, 7, 10, 0, 0, 294, 68, 1, 0,
		0, 0, 295, 296, 5, 61, 0, 0, 296, 70, 1, 0, 0, 0, 297, 298, 5, 61, 0, 0,
		298, 299, 5, 61, 0, 0, 299, 72, 1, 0, 0, 0, 300, 301, 5, 33, 0, 0, 301,
		302, 5, 61, 0, 0, 302, 74, 1, 0, 0, 0, 303, 304, 5, 62, 0, 0, 304, 76,
		1, 0, 0, 0, 305, 306, 5, 60, 0, 0, 306, 78, 1, 0, 0, 0, 307, 308, 5, 62,
		0, 0, 308, 309, 5, 61, 0, 0, 309, 80, 1, 0, 0, 0, 310, 311, 5, 60, 0, 0,
		311, 312, 5, 61, 0, 0, 312, 82, 1, 0, 0, 0, 313, 314, 5, 61, 0, 0, 314,
		315, 5, 126, 0, 0, 315, 84, 1, 0, 0, 0, 316, 317, 5, 33, 0, 0, 317, 318,
		5, 126, 0, 0, 318, 86, 1, 0, 0, 0, 319, 320, 7, 15, 0, 0, 320, 321, 7,
		22, 0, 0, 321, 88, 1, 0, 0, 0, 322, 323, 7, 23, 0, 0, 323, 324, 7, 2, 0,
		0, 324, 325, 7, 4, 0, 0, 325, 326, 7, 24, 0, 0, 326, 327, 7, 21, 0, 0,
		327, 328, 7, 6, 0, 0, 328, 329, 7, 4, 0, 0, 329, 90, 1, 0, 0, 0, 330, 331,
		7, 21, 0, 0, 331, 332, 7, 20, 0, 0, 332, 92, 1, 0, 0, 0, 333, 334, 7, 2,
		0, 0, 334, 335, 7, 25, 0, 0, 335, 336, 7, 20, 0, 0, 336, 337, 7, 21, 0,
		0, 337, 338, 7, 5, 0, 0, 338, 339, 7, 2, 0, 0, 339, 340, 7, 20, 0, 0, 340,
		341, 7, 25, 0, 0, 341, 94, 1, 0, 0, 0, 342, 343, 7, 25, 0, 0, 343, 344,
		7, 5, 0, 0, 344, 345, 7, 21, 0, 0, 345, 346, 7, 6, 0, 0, 346, 347, 7, 26,
		0, 0, 347, 348, 5, 95, 0, 0, 348, 349, 7, 9, 0, 0, 349, 350, 7, 7, 0, 0,
		350, 351, 7, 3, 0, 0, 351, 352, 7, 4, 0, 0, 352, 96, 1, 0, 0, 0, 353, 354,
		7, 25, 0, 0, 354, 355, 7, 5, 0, 0, 355, 356, 7, 21, 0, 0, 356, 357, 7,
		6, 0, 0, 357, 358, 7, 26, 0, 0, 358, 359, 5, 95, 0, 0, 359, 360, 7, 5,
		0, 0, 360, 361, 7, 2, 0, 0, 361, 362, 7, 25, 0, 0, 362, 363, 7, 24, 0,
		0, 363, 364, 7, 4, 0, 0, 364, 98, 1, 0, 0, 0, 365, 366, 7, 21, 0, 0, 366,
		367, 7, 3, 0, 0, 367, 368, 7, 3, 0, 0, 368, 369, 7, 10, 0, 0, 369, 370,
		7, 7, 0, 0, 370, 371, 7, 4, 0, 0, 371, 100, 1, 0, 0, 0, 372, 373, 7, 15,
		0, 0, 373, 374, 7, 21, 0, 0, 374, 375, 7, 21, 0, 0, 375, 376, 7, 9, 0,
		0, 376, 102, 1, 0, 0, 0, 377, 378, 5, 123, 0, 0, 378, 104, 1, 0, 0, 0,
		379, 380, 5, 125, 0, 0, 380, 106, 1, 0, 0, 0, 381, 382, 5, 40, 0, 0, 382,
		108, 1, 0, 0, 0, 383, 384, 5, 41, 0, 0, 384, 110, 1, 0, 0, 0, 385, 386,
		5, 44, 0, 0, 386, 112, 1, 0, 0, 0, 387, 388, 5, 64, 0, 0, 388, 114, 1,
		0, 0, 0, 389, 391, 7, 11, 0, 0, 390, 389, 1, 0, 0, 0, 391, 392, 1, 0, 0,
		0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393, 397, 1, 0, 0, 0, 394,
		395, 7, 13, 0, 0, 395, 398, 7, 10, 0, 0, 396, 398, 7, 27, 0, 0, 397, 394,
		1, 0, 0, 0, 397, 396, 1, 0, 0, 0, 398, 400, 1, 0, 0, 0, 399, 390, 1, 0,
		0, 0, 400, 401, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0,
		402, 116, 1, 0, 0, 0, 403, 407, 7, 28, 0, 0, 404, 406, 7, 29, 0, 0, 405,
		404, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408,
		1, 0, 0, 0, 408, 118, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 412, 7, 30,
		0, 0, 411, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0,
		413, 414, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416, 6, 58, 6, 0, 416,
		120, 1, 0, 0, 0, 417, 421, 5, 35, 0, 0, 418, 420, 9, 0, 0, 0, 419, 418,
		1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 421, 419, 1, 0,
		0, 0, 422, 424, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 424, 425, 5, 10, 0, 0,
		425, 426, 1, 0, 0, 0, 426, 427, 6, 59, 7, 0, 427, 122, 1, 0, 0, 0, 428,
		432, 7, 31, 0, 0, 429, 431, 7, 29, 0, 0, 430, 429, 1, 0, 0, 0, 431, 434,
		1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 435, 1, 0,
		0, 0, 434, 432, 1, 0, 0, 0, 435, 436, 6, 60, 8, 0, 436, 124, 1, 0, 0, 0,
		437, 441, 8, 32, 0, 0, 438, 439, 5, 92, 0, 0, 439, 441, 5, 96, 0, 0, 440,
		437, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 440,
		1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 445, 6, 61,
		4, 0, 445, 126, 1, 0, 0, 0, 446, 447, 5, 96, 0, 0, 447, 448, 1, 0, 0, 0,
		448, 449, 6, 62, 8, 0, 449, 128, 1, 0, 0, 0, 28, 0, 1, 2, 138, 170, 175,
		178, 186, 220, 226, 228, 233, 236, 240, 246, 248, 256, 258, 262, 392, 397,
		401, 407, 413, 421, 432, 440, 442, 9, 1, 2, 0, 5, 1, 0, 1, 17, 1, 1, 18,
		2, 3, 0, 0, 5, 2, 0, 0, 2, 0, 0, 3, 0, 4, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PromQLExLexerInit initializes any static state used to implement PromQLExLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPromQLExLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLExLexerInit() {
	staticData := &PromQLExLexerLexerStaticData
	staticData.once.Do(promqlexlexerLexerInit)
}

// NewPromQLExLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPromQLExLexer(input antlr.CharStream) *PromQLExLexer {
	PromQLExLexerInit()
	l := new(PromQLExLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PromQLExLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PromQLExLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PromQLExLexer tokens.
const (
	PromQLExLexerAGGREGATION_OPERATOR = 1
	PromQLExLexerFUNCTION             = 2
	PromQLExLexerMETRIC_NAME          = 3
	PromQLExLexerIF                   = 4
	PromQLExLexerTRUE                 = 5
	PromQLExLexerFALSE                = 6
	PromQLExLexerT                    = 7
	PromQLExLexerCOLON                = 8
	PromQLExLexerSEMICOLON            = 9
	PromQLExLexerDOT                  = 10
	PromQLExLexerPOSITIVE_INTEGER     = 11
	PromQLExLexerTWO_DIGITS           = 12
	PromQLExLexerDIGITS               = 13
	PromQLExLexerMETRIC_KEYWORD       = 14
	PromQLExLexerLABEL_KEYWORD        = 15
	PromQLExLexerDEF                  = 16
	PromQLExLexerCALL_SIGN            = 17
	PromQLExLexerLEFT_BRACKET         = 18
	PromQLExLexerRIGHT_BRACKET        = 19
	PromQLExLexerNUMBER               = 20
	PromQLExLexerSTRING               = 21
	PromQLExLexerADD                  = 22
	PromQLExLexerSUB                  = 23
	PromQLExLexerMULT                 = 24
	PromQLExLexerDIV                  = 25
	PromQLExLexerMOD                  = 26
	PromQLExLexerPOW                  = 27
	PromQLExLexerAND                  = 28
	PromQLExLexerOR                   = 29
	PromQLExLexerUNLESS               = 30
	PromQLExLexerEQ                   = 31
	PromQLExLexerDEQ                  = 32
	PromQLExLexerNE                   = 33
	PromQLExLexerGT                   = 34
	PromQLExLexerLT                   = 35
	PromQLExLexerGE                   = 36
	PromQLExLexerLE                   = 37
	PromQLExLexerRE                   = 38
	PromQLExLexerNRE                  = 39
	PromQLExLexerBY                   = 40
	PromQLExLexerWITHOUT              = 41
	PromQLExLexerON                   = 42
	PromQLExLexerIGNORING             = 43
	PromQLExLexerGROUP_LEFT           = 44
	PromQLExLexerGROUP_RIGHT          = 45
	PromQLExLexerOFFSET               = 46
	PromQLExLexerBOOL                 = 47
	PromQLExLexerLEFT_BRACE           = 48
	PromQLExLexerRIGHT_BRACE          = 49
	PromQLExLexerLEFT_PAREN           = 50
	PromQLExLexerRIGHT_PAREN          = 51
	PromQLExLexerCOMMA                = 52
	PromQLExLexerAT                   = 53
	PromQLExLexerDURATION             = 54
	PromQLExLexerLABEL_NAME           = 55
	PromQLExLexerWS                   = 56
	PromQLExLexerSL_COMMENT           = 57
	PromQLExLexerID                   = 58
	PromQLExLexerRAW_STRING           = 59
	PromQLExLexerBACKTICK_OPEN        = 60
)

// PromQLExLexer escapedChannels.
const (
	PromQLExLexerWHITESPACE = 2
	PromQLExLexerCOMMENTS   = 3
)

// PromQLExLexer modes.
const (
	PromQLExLexerID_MODE = iota + 1
	PromQLExLexerRAW_STRING_MODE
)

type BracketCounter interface {
	BracketCount() int
	SetBracketCount(val int)
}

func (l *PromQLExLexer) IsLiteralName(text string) bool {
	return slices.ContainsFunc(l.LiteralNames, func(literlName string) bool {
		if strings.ToLower(strings.Trim(literlName, "'")) == strings.ToLower(text) {
			return true
		}
		return false
	})
}

type FunctionsProvider interface {
	GetTokenType(text string) (int, bool)
}

func (l *PromQLExLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 2:
		l.METRIC_NAME_Action(localctx, actionIndex)

	case 17:
		l.LEFT_BRACKET_Action(localctx, actionIndex)

	case 18:
		l.RIGHT_BRACKET_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PromQLExLexer) METRIC_NAME_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:

		if prov, ok := l.GetInputStream().(FunctionsProvider); ok {
			if tokenType, ok := prov.GetTokenType(l.GetText()); ok {
				l.SetType(tokenType)
			}
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *PromQLExLexer) LEFT_BRACKET_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:

		if cnt, ok := l.GetInputStream().(BracketCounter); ok {
			cnt.SetBracketCount(cnt.BracketCount() + 1)
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *PromQLExLexer) RIGHT_BRACKET_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:

		if cnt, ok := l.GetInputStream().(BracketCounter); ok {
			cnt.SetBracketCount(cnt.BracketCount() - 1)
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

func (l *PromQLExLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		return l.METRIC_NAME_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLExLexer) METRIC_NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return func() bool {
			cnt, ok := p.GetInputStream().(BracketCounter)
			return (!ok || (ok && cnt.BracketCount() == 0))
		}()

	case 1:
		return !p.IsLiteralName(p.GetText())

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

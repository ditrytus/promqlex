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
		"DEFAULT_MODE", "ID_MODE", "RAW_STRING_MODE", "AT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'if'", "'true'", "'false'", "'T'", "':'", "';'", "'.'",
		"", "", "", "'metric'", "'label'", "'def'", "'$'", "'['", "']'", "",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'", "",
		"", "", "", "", "", "", "'start'", "'end'",
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
		"AT_NUMBER", "START", "END", "AT_WS", "AT_OTHER", "BACKTICK_OPEN",
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
		"RAW_STRING", "AT_NUMBER", "START", "END", "AT_WS", "AT_OTHER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 65, 493, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2,
		7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8,
		7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13,
		2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2,
		19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24,
		7, 24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7,
		29, 2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34,
		2, 35, 7, 35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2,
		40, 7, 40, 2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45,
		7, 45, 2, 46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7,
		50, 2, 51, 7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55,
		2, 56, 7, 56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2,
		61, 7, 61, 2, 62, 7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66,
		7, 66, 2, 67, 7, 67, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 148,
		8, 2, 10, 2, 12, 2, 151, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 4, 10, 180, 8, 10, 11,
		10, 12, 10, 181, 1, 10, 5, 10, 185, 8, 10, 10, 10, 12, 10, 188, 9, 10,
		3, 10, 190, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 196, 8, 12, 11, 12,
		12, 12, 197, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 4, 19, 230, 8, 19, 11, 19, 12, 19, 231, 1, 19, 1, 19, 4, 19, 236, 8,
		19, 11, 19, 12, 19, 237, 3, 19, 240, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20,
		245, 8, 20, 1, 20, 3, 20, 248, 8, 20, 1, 21, 1, 21, 3, 21, 252, 8, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 258, 8, 22, 10, 22, 12, 22, 261, 9,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 268, 8, 22, 10, 22, 12, 22,
		271, 9, 22, 1, 22, 3, 22, 274, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
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
		53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 4, 56, 404, 8, 56,
		11, 56, 12, 56, 405, 1, 56, 1, 56, 1, 56, 3, 56, 411, 8, 56, 4, 56, 413,
		8, 56, 11, 56, 12, 56, 414, 1, 57, 1, 57, 5, 57, 419, 8, 57, 10, 57, 12,
		57, 422, 9, 57, 1, 58, 4, 58, 425, 8, 58, 11, 58, 12, 58, 426, 1, 58, 1,
		58, 1, 59, 1, 59, 5, 59, 433, 8, 59, 10, 59, 12, 59, 436, 9, 59, 1, 59,
		1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 5, 60, 444, 8, 60, 10, 60, 12, 60, 447,
		9, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 4, 61, 454, 8, 61, 11, 61, 12,
		61, 455, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64,
		1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1,
		65, 1, 65, 1, 65, 1, 66, 4, 66, 481, 8, 66, 11, 66, 12, 66, 482, 1, 66,
		1, 66, 1, 67, 4, 67, 488, 8, 67, 11, 67, 12, 67, 489, 1, 67, 1, 67, 2,
		434, 489, 0, 68, 4, 0, 6, 0, 8, 3, 10, 4, 12, 5, 14, 6, 16, 7, 18, 8, 20,
		9, 22, 10, 24, 11, 26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 17, 38,
		18, 40, 19, 42, 0, 44, 0, 46, 20, 48, 21, 50, 65, 52, 22, 54, 23, 56, 24,
		58, 25, 60, 26, 62, 27, 64, 28, 66, 29, 68, 30, 70, 31, 72, 32, 74, 33,
		76, 34, 78, 35, 80, 36, 82, 37, 84, 38, 86, 39, 88, 40, 90, 41, 92, 42,
		94, 43, 96, 44, 98, 45, 100, 46, 102, 47, 104, 48, 106, 49, 108, 50, 110,
		51, 112, 52, 114, 53, 116, 54, 118, 55, 120, 56, 122, 57, 124, 58, 126,
		0, 128, 59, 130, 60, 132, 61, 134, 62, 136, 63, 138, 64, 4, 0, 1, 2, 3,
		33, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90, 95, 95,
		97, 122, 2, 0, 73, 73, 105, 105, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84,
		116, 116, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 69, 69,
		101, 101, 2, 0, 65, 65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 83, 83, 115,
		115, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 77, 77, 109, 109, 2, 0, 67, 67,
		99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 68, 68, 100, 100, 2, 0, 43, 43, 45,
		45, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 2, 0, 78, 78, 110, 110,
		2, 0, 79, 79, 111, 111, 2, 0, 89, 89, 121, 121, 2, 0, 87, 87, 119, 119,
		2, 0, 72, 72, 104, 104, 2, 0, 71, 71, 103, 103, 2, 0, 80, 80, 112, 112,
		12, 0, 68, 68, 72, 72, 77, 77, 83, 83, 87, 87, 89, 89, 100, 100, 104, 104,
		109, 109, 115, 115, 119, 119, 121, 121, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2,
		0, 65, 90, 97, 122, 2, 0, 92, 92, 96, 96, 512, 0, 8, 1, 0, 0, 0, 0, 10,
		1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0,
		18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0,
		0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0,
		0, 0, 34, 1, 0, 0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0, 0, 0, 40, 1, 0,
		0, 0, 0, 46, 1, 0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1,
		0, 0, 0, 0, 54, 1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60,
		1, 0, 0, 0, 0, 62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0,
		68, 1, 0, 0, 0, 0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0,
		0, 76, 1, 0, 0, 0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0,
		0, 0, 84, 1, 0, 0, 0, 0, 86, 1, 0, 0, 0, 0, 88, 1, 0, 0, 0, 0, 90, 1, 0,
		0, 0, 0, 92, 1, 0, 0, 0, 0, 94, 1, 0, 0, 0, 0, 96, 1, 0, 0, 0, 0, 98, 1,
		0, 0, 0, 0, 100, 1, 0, 0, 0, 0, 102, 1, 0, 0, 0, 0, 104, 1, 0, 0, 0, 0,
		106, 1, 0, 0, 0, 0, 108, 1, 0, 0, 0, 0, 110, 1, 0, 0, 0, 0, 112, 1, 0,
		0, 0, 0, 114, 1, 0, 0, 0, 0, 116, 1, 0, 0, 0, 0, 118, 1, 0, 0, 0, 0, 120,
		1, 0, 0, 0, 0, 122, 1, 0, 0, 0, 1, 124, 1, 0, 0, 0, 2, 126, 1, 0, 0, 0,
		2, 128, 1, 0, 0, 0, 3, 130, 1, 0, 0, 0, 3, 132, 1, 0, 0, 0, 3, 134, 1,
		0, 0, 0, 3, 136, 1, 0, 0, 0, 3, 138, 1, 0, 0, 0, 4, 140, 1, 0, 0, 0, 6,
		142, 1, 0, 0, 0, 8, 144, 1, 0, 0, 0, 10, 155, 1, 0, 0, 0, 12, 158, 1, 0,
		0, 0, 14, 163, 1, 0, 0, 0, 16, 169, 1, 0, 0, 0, 18, 171, 1, 0, 0, 0, 20,
		173, 1, 0, 0, 0, 22, 175, 1, 0, 0, 0, 24, 189, 1, 0, 0, 0, 26, 191, 1,
		0, 0, 0, 28, 195, 1, 0, 0, 0, 30, 199, 1, 0, 0, 0, 32, 206, 1, 0, 0, 0,
		34, 212, 1, 0, 0, 0, 36, 218, 1, 0, 0, 0, 38, 222, 1, 0, 0, 0, 40, 225,
		1, 0, 0, 0, 42, 229, 1, 0, 0, 0, 44, 241, 1, 0, 0, 0, 46, 251, 1, 0, 0,
		0, 48, 273, 1, 0, 0, 0, 50, 275, 1, 0, 0, 0, 52, 280, 1, 0, 0, 0, 54, 282,
		1, 0, 0, 0, 56, 284, 1, 0, 0, 0, 58, 286, 1, 0, 0, 0, 60, 288, 1, 0, 0,
		0, 62, 290, 1, 0, 0, 0, 64, 292, 1, 0, 0, 0, 66, 296, 1, 0, 0, 0, 68, 299,
		1, 0, 0, 0, 70, 306, 1, 0, 0, 0, 72, 308, 1, 0, 0, 0, 74, 311, 1, 0, 0,
		0, 76, 314, 1, 0, 0, 0, 78, 316, 1, 0, 0, 0, 80, 318, 1, 0, 0, 0, 82, 321,
		1, 0, 0, 0, 84, 324, 1, 0, 0, 0, 86, 327, 1, 0, 0, 0, 88, 330, 1, 0, 0,
		0, 90, 333, 1, 0, 0, 0, 92, 341, 1, 0, 0, 0, 94, 344, 1, 0, 0, 0, 96, 353,
		1, 0, 0, 0, 98, 364, 1, 0, 0, 0, 100, 376, 1, 0, 0, 0, 102, 383, 1, 0,
		0, 0, 104, 388, 1, 0, 0, 0, 106, 390, 1, 0, 0, 0, 108, 392, 1, 0, 0, 0,
		110, 394, 1, 0, 0, 0, 112, 396, 1, 0, 0, 0, 114, 398, 1, 0, 0, 0, 116,
		412, 1, 0, 0, 0, 118, 416, 1, 0, 0, 0, 120, 424, 1, 0, 0, 0, 122, 430,
		1, 0, 0, 0, 124, 441, 1, 0, 0, 0, 126, 453, 1, 0, 0, 0, 128, 459, 1, 0,
		0, 0, 130, 463, 1, 0, 0, 0, 132, 465, 1, 0, 0, 0, 134, 473, 1, 0, 0, 0,
		136, 480, 1, 0, 0, 0, 138, 487, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		5, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 7, 1, 0, 0, 0, 144, 145, 4, 2,
		0, 0, 145, 149, 7, 0, 0, 0, 146, 148, 7, 1, 0, 0, 147, 146, 1, 0, 0, 0,
		148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150,
		152, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 153, 4, 2, 1, 0, 153, 154,
		6, 2, 0, 0, 154, 9, 1, 0, 0, 0, 155, 156, 7, 2, 0, 0, 156, 157, 7, 3, 0,
		0, 157, 11, 1, 0, 0, 0, 158, 159, 7, 4, 0, 0, 159, 160, 7, 5, 0, 0, 160,
		161, 7, 6, 0, 0, 161, 162, 7, 7, 0, 0, 162, 13, 1, 0, 0, 0, 163, 164, 7,
		3, 0, 0, 164, 165, 7, 8, 0, 0, 165, 166, 7, 9, 0, 0, 166, 167, 7, 10, 0,
		0, 167, 168, 7, 7, 0, 0, 168, 15, 1, 0, 0, 0, 169, 170, 7, 4, 0, 0, 170,
		17, 1, 0, 0, 0, 171, 172, 5, 58, 0, 0, 172, 19, 1, 0, 0, 0, 173, 174, 5,
		59, 0, 0, 174, 21, 1, 0, 0, 0, 175, 176, 5, 46, 0, 0, 176, 23, 1, 0, 0,
		0, 177, 190, 7, 11, 0, 0, 178, 180, 7, 12, 0, 0, 179, 178, 1, 0, 0, 0,
		180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182,
		186, 1, 0, 0, 0, 183, 185, 7, 11, 0, 0, 184, 183, 1, 0, 0, 0, 185, 188,
		1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 190, 1, 0,
		0, 0, 188, 186, 1, 0, 0, 0, 189, 177, 1, 0, 0, 0, 189, 179, 1, 0, 0, 0,
		190, 25, 1, 0, 0, 0, 191, 192, 7, 11, 0, 0, 192, 193, 7, 11, 0, 0, 193,
		27, 1, 0, 0, 0, 194, 196, 7, 11, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197,
		1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 29, 1, 0,
		0, 0, 199, 200, 7, 13, 0, 0, 200, 201, 7, 7, 0, 0, 201, 202, 7, 4, 0, 0,
		202, 203, 7, 5, 0, 0, 203, 204, 7, 2, 0, 0, 204, 205, 7, 14, 0, 0, 205,
		31, 1, 0, 0, 0, 206, 207, 7, 9, 0, 0, 207, 208, 7, 8, 0, 0, 208, 209, 7,
		15, 0, 0, 209, 210, 7, 7, 0, 0, 210, 211, 7, 9, 0, 0, 211, 33, 1, 0, 0,
		0, 212, 213, 7, 16, 0, 0, 213, 214, 7, 7, 0, 0, 214, 215, 7, 3, 0, 0, 215,
		216, 1, 0, 0, 0, 216, 217, 6, 15, 1, 0, 217, 35, 1, 0, 0, 0, 218, 219,
		5, 36, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 6, 16, 1, 0, 221, 37, 1, 0,
		0, 0, 222, 223, 5, 91, 0, 0, 223, 224, 6, 17, 2, 0, 224, 39, 1, 0, 0, 0,
		225, 226, 5, 93, 0, 0, 226, 227, 6, 18, 3, 0, 227, 41, 1, 0, 0, 0, 228,
		230, 7, 11, 0, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 239, 1, 0, 0, 0, 233, 235, 5, 46,
		0, 0, 234, 236, 7, 11, 0, 0, 235, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0,
		237, 235, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239,
		233, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 43, 1, 0, 0, 0, 241, 247, 3,
		42, 19, 0, 242, 244, 7, 7, 0, 0, 243, 245, 7, 17, 0, 0, 244, 243, 1, 0,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248, 3, 42, 19,
		0, 247, 242, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 45, 1, 0, 0, 0, 249,
		252, 3, 42, 19, 0, 250, 252, 3, 44, 20, 0, 251, 249, 1, 0, 0, 0, 251, 250,
		1, 0, 0, 0, 252, 47, 1, 0, 0, 0, 253, 259, 5, 39, 0, 0, 254, 258, 8, 18,
		0, 0, 255, 256, 5, 92, 0, 0, 256, 258, 9, 0, 0, 0, 257, 254, 1, 0, 0, 0,
		257, 255, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 262, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 274,
		5, 39, 0, 0, 263, 269, 5, 34, 0, 0, 264, 268, 8, 19, 0, 0, 265, 266, 5,
		92, 0, 0, 266, 268, 9, 0, 0, 0, 267, 264, 1, 0, 0, 0, 267, 265, 1, 0, 0,
		0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0, 0, 270,
		272, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 274, 5, 34, 0, 0, 273, 253,
		1, 0, 0, 0, 273, 263, 1, 0, 0, 0, 274, 49, 1, 0, 0, 0, 275, 276, 5, 96,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 6, 23, 4, 0, 278, 279, 6, 23, 5,
		0, 279, 51, 1, 0, 0, 0, 280, 281, 5, 43, 0, 0, 281, 53, 1, 0, 0, 0, 282,
		283, 5, 45, 0, 0, 283, 55, 1, 0, 0, 0, 284, 285, 5, 42, 0, 0, 285, 57,
		1, 0, 0, 0, 286, 287, 5, 47, 0, 0, 287, 59, 1, 0, 0, 0, 288, 289, 5, 37,
		0, 0, 289, 61, 1, 0, 0, 0, 290, 291, 5, 94, 0, 0, 291, 63, 1, 0, 0, 0,
		292, 293, 7, 8, 0, 0, 293, 294, 7, 20, 0, 0, 294, 295, 7, 16, 0, 0, 295,
		65, 1, 0, 0, 0, 296, 297, 7, 21, 0, 0, 297, 298, 7, 5, 0, 0, 298, 67, 1,
		0, 0, 0, 299, 300, 7, 6, 0, 0, 300, 301, 7, 20, 0, 0, 301, 302, 7, 9, 0,
		0, 302, 303, 7, 7, 0, 0, 303, 304, 7, 10, 0, 0, 304, 305, 7, 10, 0, 0,
		305, 69, 1, 0, 0, 0, 306, 307, 5, 61, 0, 0, 307, 71, 1, 0, 0, 0, 308, 309,
		5, 61, 0, 0, 309, 310, 5, 61, 0, 0, 310, 73, 1, 0, 0, 0, 311, 312, 5, 33,
		0, 0, 312, 313, 5, 61, 0, 0, 313, 75, 1, 0, 0, 0, 314, 315, 5, 62, 0, 0,
		315, 77, 1, 0, 0, 0, 316, 317, 5, 60, 0, 0, 317, 79, 1, 0, 0, 0, 318, 319,
		5, 62, 0, 0, 319, 320, 5, 61, 0, 0, 320, 81, 1, 0, 0, 0, 321, 322, 5, 60,
		0, 0, 322, 323, 5, 61, 0, 0, 323, 83, 1, 0, 0, 0, 324, 325, 5, 61, 0, 0,
		325, 326, 5, 126, 0, 0, 326, 85, 1, 0, 0, 0, 327, 328, 5, 33, 0, 0, 328,
		329, 5, 126, 0, 0, 329, 87, 1, 0, 0, 0, 330, 331, 7, 15, 0, 0, 331, 332,
		7, 22, 0, 0, 332, 89, 1, 0, 0, 0, 333, 334, 7, 23, 0, 0, 334, 335, 7, 2,
		0, 0, 335, 336, 7, 4, 0, 0, 336, 337, 7, 24, 0, 0, 337, 338, 7, 21, 0,
		0, 338, 339, 7, 6, 0, 0, 339, 340, 7, 4, 0, 0, 340, 91, 1, 0, 0, 0, 341,
		342, 7, 21, 0, 0, 342, 343, 7, 20, 0, 0, 343, 93, 1, 0, 0, 0, 344, 345,
		7, 2, 0, 0, 345, 346, 7, 25, 0, 0, 346, 347, 7, 20, 0, 0, 347, 348, 7,
		21, 0, 0, 348, 349, 7, 5, 0, 0, 349, 350, 7, 2, 0, 0, 350, 351, 7, 20,
		0, 0, 351, 352, 7, 25, 0, 0, 352, 95, 1, 0, 0, 0, 353, 354, 7, 25, 0, 0,
		354, 355, 7, 5, 0, 0, 355, 356, 7, 21, 0, 0, 356, 357, 7, 6, 0, 0, 357,
		358, 7, 26, 0, 0, 358, 359, 5, 95, 0, 0, 359, 360, 7, 9, 0, 0, 360, 361,
		7, 7, 0, 0, 361, 362, 7, 3, 0, 0, 362, 363, 7, 4, 0, 0, 363, 97, 1, 0,
		0, 0, 364, 365, 7, 25, 0, 0, 365, 366, 7, 5, 0, 0, 366, 367, 7, 21, 0,
		0, 367, 368, 7, 6, 0, 0, 368, 369, 7, 26, 0, 0, 369, 370, 5, 95, 0, 0,
		370, 371, 7, 5, 0, 0, 371, 372, 7, 2, 0, 0, 372, 373, 7, 25, 0, 0, 373,
		374, 7, 24, 0, 0, 374, 375, 7, 4, 0, 0, 375, 99, 1, 0, 0, 0, 376, 377,
		7, 21, 0, 0, 377, 378, 7, 3, 0, 0, 378, 379, 7, 3, 0, 0, 379, 380, 7, 10,
		0, 0, 380, 381, 7, 7, 0, 0, 381, 382, 7, 4, 0, 0, 382, 101, 1, 0, 0, 0,
		383, 384, 7, 15, 0, 0, 384, 385, 7, 21, 0, 0, 385, 386, 7, 21, 0, 0, 386,
		387, 7, 9, 0, 0, 387, 103, 1, 0, 0, 0, 388, 389, 5, 123, 0, 0, 389, 105,
		1, 0, 0, 0, 390, 391, 5, 125, 0, 0, 391, 107, 1, 0, 0, 0, 392, 393, 5,
		40, 0, 0, 393, 109, 1, 0, 0, 0, 394, 395, 5, 41, 0, 0, 395, 111, 1, 0,
		0, 0, 396, 397, 5, 44, 0, 0, 397, 113, 1, 0, 0, 0, 398, 399, 5, 64, 0,
		0, 399, 400, 1, 0, 0, 0, 400, 401, 6, 55, 6, 0, 401, 115, 1, 0, 0, 0, 402,
		404, 7, 11, 0, 0, 403, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 403,
		1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 410, 1, 0, 0, 0, 407, 408, 7, 13,
		0, 0, 408, 411, 7, 10, 0, 0, 409, 411, 7, 27, 0, 0, 410, 407, 1, 0, 0,
		0, 410, 409, 1, 0, 0, 0, 411, 413, 1, 0, 0, 0, 412, 403, 1, 0, 0, 0, 413,
		414, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 117,
		1, 0, 0, 0, 416, 420, 7, 28, 0, 0, 417, 419, 7, 29, 0, 0, 418, 417, 1,
		0, 0, 0, 419, 422, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421, 1, 0, 0,
		0, 421, 119, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 423, 425, 7, 30, 0, 0, 424,
		423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427,
		1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 6, 58, 7, 0, 429, 121, 1, 0,
		0, 0, 430, 434, 5, 35, 0, 0, 431, 433, 9, 0, 0, 0, 432, 431, 1, 0, 0, 0,
		433, 436, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 435,
		437, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 438, 5, 10, 0, 0, 438, 439,
		1, 0, 0, 0, 439, 440, 6, 59, 8, 0, 440, 123, 1, 0, 0, 0, 441, 445, 7, 31,
		0, 0, 442, 444, 7, 29, 0, 0, 443, 442, 1, 0, 0, 0, 444, 447, 1, 0, 0, 0,
		445, 443, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 448, 1, 0, 0, 0, 447,
		445, 1, 0, 0, 0, 448, 449, 6, 60, 9, 0, 449, 125, 1, 0, 0, 0, 450, 454,
		8, 32, 0, 0, 451, 452, 5, 92, 0, 0, 452, 454, 5, 96, 0, 0, 453, 450, 1,
		0, 0, 0, 453, 451, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0, 455, 453, 1, 0, 0,
		0, 455, 456, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 458, 6, 61, 4, 0, 458,
		127, 1, 0, 0, 0, 459, 460, 5, 96, 0, 0, 460, 461, 1, 0, 0, 0, 461, 462,
		6, 62, 9, 0, 462, 129, 1, 0, 0, 0, 463, 464, 3, 46, 21, 0, 464, 131, 1,
		0, 0, 0, 465, 466, 7, 10, 0, 0, 466, 467, 7, 4, 0, 0, 467, 468, 7, 8, 0,
		0, 468, 469, 7, 5, 0, 0, 469, 470, 7, 4, 0, 0, 470, 471, 1, 0, 0, 0, 471,
		472, 6, 64, 9, 0, 472, 133, 1, 0, 0, 0, 473, 474, 7, 7, 0, 0, 474, 475,
		7, 20, 0, 0, 475, 476, 7, 16, 0, 0, 476, 477, 1, 0, 0, 0, 477, 478, 6,
		65, 9, 0, 478, 135, 1, 0, 0, 0, 479, 481, 7, 30, 0, 0, 480, 479, 1, 0,
		0, 0, 481, 482, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0,
		483, 484, 1, 0, 0, 0, 484, 485, 6, 66, 7, 0, 485, 137, 1, 0, 0, 0, 486,
		488, 9, 0, 0, 0, 487, 486, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 490,
		1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492, 6, 67,
		9, 0, 492, 139, 1, 0, 0, 0, 31, 0, 1, 2, 3, 149, 181, 186, 189, 197, 231,
		237, 239, 244, 247, 251, 257, 259, 267, 269, 273, 405, 410, 414, 420, 426,
		434, 445, 453, 455, 482, 489, 10, 1, 2, 0, 5, 1, 0, 1, 17, 1, 1, 18, 2,
		3, 0, 0, 5, 2, 0, 5, 3, 0, 0, 2, 0, 0, 3, 0, 4, 0, 0,
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
	PromQLExLexerAT_NUMBER            = 60
	PromQLExLexerSTART                = 61
	PromQLExLexerEND                  = 62
	PromQLExLexerAT_WS                = 63
	PromQLExLexerAT_OTHER             = 64
	PromQLExLexerBACKTICK_OPEN        = 65
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
	PromQLExLexerAT_MODE
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

// Code generated from PromQLExLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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
		"DEFAULT_MODE", "RAW_STRING_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'if'", "'true'", "'false'", "'T'", "':'", "'.'",
		"", "", "", "'metric'", "'label'", "'def'", "'$'", "", "'['", "']'",
		"", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "ID", "METRIC_NAME", "IF", "TRUE",
		"FALSE", "T", "COLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS", "DIGITS",
		"METRIC_KEYWORD", "LABEL_KEYWORD", "DEF", "CALL_SIGN", "NL", "LEFT_BRACKET",
		"RIGHT_BRACKET", "NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD",
		"POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE",
		"RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT",
		"OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN",
		"COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "RAW_STRING",
		"BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"SUBQUERY_RANGE", "TIME_RANGE", "ID", "METRIC_NAME", "IF", "TRUE", "FALSE",
		"T", "COLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS", "DIGITS", "METRIC_KEYWORD",
		"LABEL_KEYWORD", "DEF", "CALL_SIGN", "NL", "LEFT_BRACKET", "RIGHT_BRACKET",
		"NUMERAL", "SCIENTIFIC_NUMBER", "NUMBER", "STRING", "BACKTICK_OPEN",
		"ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR", "UNLESS", "EQ",
		"DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON",
		"IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE",
		"RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN", "COMMA", "AT", "DURATION",
		"LABEL_NAME", "WS", "SL_COMMENT", "RAW_STRING_CONTENT", "RAW_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 60, 445, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2,
		25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30,
		7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7,
		35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40,
		2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2,
		46, 7, 46, 2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51,
		7, 51, 2, 52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7,
		56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61,
		2, 62, 7, 62, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 135, 8, 2, 10,
		2, 12, 2, 138, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 145, 8, 3, 10,
		3, 12, 3, 148, 9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 4, 10, 172, 8, 10, 11, 10, 12, 10, 173, 1, 10, 5, 10, 177, 8,
		10, 10, 10, 12, 10, 180, 9, 10, 3, 10, 182, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 4, 12, 188, 8, 12, 11, 12, 12, 12, 189, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 214, 8,
		17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 4, 20, 223, 8, 20,
		11, 20, 12, 20, 224, 1, 20, 1, 20, 4, 20, 229, 8, 20, 11, 20, 12, 20, 230,
		3, 20, 233, 8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 238, 8, 21, 1, 21, 3, 21,
		241, 8, 21, 1, 22, 1, 22, 3, 22, 245, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		5, 23, 251, 8, 23, 10, 23, 12, 23, 254, 9, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 5, 23, 261, 8, 23, 10, 23, 12, 23, 264, 9, 23, 1, 23, 3, 23,
		267, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56,
		1, 56, 1, 57, 4, 57, 395, 8, 57, 11, 57, 12, 57, 396, 1, 57, 1, 57, 1,
		57, 3, 57, 402, 8, 57, 4, 57, 404, 8, 57, 11, 57, 12, 57, 405, 1, 58, 1,
		58, 5, 58, 410, 8, 58, 10, 58, 12, 58, 413, 9, 58, 1, 59, 4, 59, 416, 8,
		59, 11, 59, 12, 59, 417, 1, 59, 1, 59, 1, 60, 1, 60, 5, 60, 424, 8, 60,
		10, 60, 12, 60, 427, 9, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1,
		61, 4, 61, 436, 8, 61, 11, 61, 12, 61, 437, 1, 61, 1, 61, 1, 62, 1, 62,
		1, 62, 1, 62, 1, 425, 0, 63, 2, 0, 4, 0, 6, 3, 8, 4, 10, 5, 12, 6, 14,
		7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28, 14, 30, 15, 32, 16,
		34, 17, 36, 18, 38, 19, 40, 20, 42, 0, 44, 0, 46, 21, 48, 22, 50, 60, 52,
		23, 54, 24, 56, 25, 58, 26, 60, 27, 62, 28, 64, 29, 66, 30, 68, 31, 70,
		32, 72, 33, 74, 34, 76, 35, 78, 36, 80, 37, 82, 38, 84, 39, 86, 40, 88,
		41, 90, 42, 92, 43, 94, 44, 96, 45, 98, 46, 100, 47, 102, 48, 104, 49,
		106, 50, 108, 51, 110, 52, 112, 53, 114, 54, 116, 55, 118, 56, 120, 57,
		122, 58, 124, 0, 126, 59, 2, 0, 1, 33, 2, 0, 65, 90, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 58, 65, 90, 95, 95, 97, 122, 2, 0, 73, 73, 105, 105, 2, 0, 70, 70,
		102, 102, 2, 0, 84, 84, 116, 116, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85,
		117, 117, 2, 0, 69, 69, 101, 101, 2, 0, 65, 65, 97, 97, 2, 0, 76, 76, 108,
		108, 2, 0, 83, 83, 115, 115, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 77, 77,
		109, 109, 2, 0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 68, 68, 100,
		100, 2, 0, 43, 43, 45, 45, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92,
		2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 89, 89, 121, 121,
		2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 71, 71, 103, 103,
		2, 0, 80, 80, 112, 112, 12, 0, 68, 68, 72, 72, 77, 77, 83, 83, 87, 87,
		89, 89, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119, 121, 121, 3,
		0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 92, 92,
		96, 96, 465, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0,
		12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0,
		0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0,
		0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0,
		0, 0, 0, 36, 1, 0, 0, 0, 0, 38, 1, 0, 0, 0, 0, 40, 1, 0, 0, 0, 0, 46, 1,
		0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50, 1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54,
		1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0, 58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0,
		62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0, 0, 66, 1, 0, 0, 0, 0, 68, 1, 0, 0, 0,
		0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0, 0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0,
		0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0, 0, 0, 0, 82, 1, 0, 0, 0, 0, 84, 1, 0,
		0, 0, 0, 86, 1, 0, 0, 0, 0, 88, 1, 0, 0, 0, 0, 90, 1, 0, 0, 0, 0, 92, 1,
		0, 0, 0, 0, 94, 1, 0, 0, 0, 0, 96, 1, 0, 0, 0, 0, 98, 1, 0, 0, 0, 0, 100,
		1, 0, 0, 0, 0, 102, 1, 0, 0, 0, 0, 104, 1, 0, 0, 0, 0, 106, 1, 0, 0, 0,
		0, 108, 1, 0, 0, 0, 0, 110, 1, 0, 0, 0, 0, 112, 1, 0, 0, 0, 0, 114, 1,
		0, 0, 0, 0, 116, 1, 0, 0, 0, 0, 118, 1, 0, 0, 0, 0, 120, 1, 0, 0, 0, 0,
		122, 1, 0, 0, 0, 1, 124, 1, 0, 0, 0, 1, 126, 1, 0, 0, 0, 2, 128, 1, 0,
		0, 0, 4, 130, 1, 0, 0, 0, 6, 132, 1, 0, 0, 0, 8, 141, 1, 0, 0, 0, 10, 149,
		1, 0, 0, 0, 12, 152, 1, 0, 0, 0, 14, 157, 1, 0, 0, 0, 16, 163, 1, 0, 0,
		0, 18, 165, 1, 0, 0, 0, 20, 167, 1, 0, 0, 0, 22, 181, 1, 0, 0, 0, 24, 183,
		1, 0, 0, 0, 26, 187, 1, 0, 0, 0, 28, 191, 1, 0, 0, 0, 30, 198, 1, 0, 0,
		0, 32, 204, 1, 0, 0, 0, 34, 208, 1, 0, 0, 0, 36, 213, 1, 0, 0, 0, 38, 215,
		1, 0, 0, 0, 40, 218, 1, 0, 0, 0, 42, 222, 1, 0, 0, 0, 44, 234, 1, 0, 0,
		0, 46, 244, 1, 0, 0, 0, 48, 266, 1, 0, 0, 0, 50, 268, 1, 0, 0, 0, 52, 273,
		1, 0, 0, 0, 54, 275, 1, 0, 0, 0, 56, 277, 1, 0, 0, 0, 58, 279, 1, 0, 0,
		0, 60, 281, 1, 0, 0, 0, 62, 283, 1, 0, 0, 0, 64, 285, 1, 0, 0, 0, 66, 289,
		1, 0, 0, 0, 68, 292, 1, 0, 0, 0, 70, 299, 1, 0, 0, 0, 72, 301, 1, 0, 0,
		0, 74, 304, 1, 0, 0, 0, 76, 307, 1, 0, 0, 0, 78, 309, 1, 0, 0, 0, 80, 311,
		1, 0, 0, 0, 82, 314, 1, 0, 0, 0, 84, 317, 1, 0, 0, 0, 86, 320, 1, 0, 0,
		0, 88, 323, 1, 0, 0, 0, 90, 326, 1, 0, 0, 0, 92, 334, 1, 0, 0, 0, 94, 337,
		1, 0, 0, 0, 96, 346, 1, 0, 0, 0, 98, 357, 1, 0, 0, 0, 100, 369, 1, 0, 0,
		0, 102, 376, 1, 0, 0, 0, 104, 381, 1, 0, 0, 0, 106, 383, 1, 0, 0, 0, 108,
		385, 1, 0, 0, 0, 110, 387, 1, 0, 0, 0, 112, 389, 1, 0, 0, 0, 114, 391,
		1, 0, 0, 0, 116, 403, 1, 0, 0, 0, 118, 407, 1, 0, 0, 0, 120, 415, 1, 0,
		0, 0, 122, 421, 1, 0, 0, 0, 124, 435, 1, 0, 0, 0, 126, 441, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 3, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 5,
		1, 0, 0, 0, 132, 136, 7, 0, 0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0,
		0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 6, 2, 0, 0, 140,
		7, 1, 0, 0, 0, 141, 142, 4, 3, 0, 0, 142, 146, 7, 2, 0, 0, 143, 145, 7,
		3, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 9, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149,
		150, 7, 4, 0, 0, 150, 151, 7, 5, 0, 0, 151, 11, 1, 0, 0, 0, 152, 153, 7,
		6, 0, 0, 153, 154, 7, 7, 0, 0, 154, 155, 7, 8, 0, 0, 155, 156, 7, 9, 0,
		0, 156, 13, 1, 0, 0, 0, 157, 158, 7, 5, 0, 0, 158, 159, 7, 10, 0, 0, 159,
		160, 7, 11, 0, 0, 160, 161, 7, 12, 0, 0, 161, 162, 7, 9, 0, 0, 162, 15,
		1, 0, 0, 0, 163, 164, 7, 6, 0, 0, 164, 17, 1, 0, 0, 0, 165, 166, 5, 58,
		0, 0, 166, 19, 1, 0, 0, 0, 167, 168, 5, 46, 0, 0, 168, 21, 1, 0, 0, 0,
		169, 182, 7, 13, 0, 0, 170, 172, 7, 14, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		173, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 178,
		1, 0, 0, 0, 175, 177, 7, 13, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0,
		0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0,
		180, 178, 1, 0, 0, 0, 181, 169, 1, 0, 0, 0, 181, 171, 1, 0, 0, 0, 182,
		23, 1, 0, 0, 0, 183, 184, 7, 13, 0, 0, 184, 185, 7, 13, 0, 0, 185, 25,
		1, 0, 0, 0, 186, 188, 7, 13, 0, 0, 187, 186, 1, 0, 0, 0, 188, 189, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 27, 1, 0, 0, 0,
		191, 192, 7, 15, 0, 0, 192, 193, 7, 9, 0, 0, 193, 194, 7, 6, 0, 0, 194,
		195, 7, 7, 0, 0, 195, 196, 7, 4, 0, 0, 196, 197, 7, 16, 0, 0, 197, 29,
		1, 0, 0, 0, 198, 199, 7, 11, 0, 0, 199, 200, 7, 10, 0, 0, 200, 201, 7,
		17, 0, 0, 201, 202, 7, 9, 0, 0, 202, 203, 7, 11, 0, 0, 203, 31, 1, 0, 0,
		0, 204, 205, 7, 18, 0, 0, 205, 206, 7, 9, 0, 0, 206, 207, 7, 5, 0, 0, 207,
		33, 1, 0, 0, 0, 208, 209, 5, 36, 0, 0, 209, 35, 1, 0, 0, 0, 210, 214, 5,
		10, 0, 0, 211, 212, 5, 13, 0, 0, 212, 214, 5, 10, 0, 0, 213, 210, 1, 0,
		0, 0, 213, 211, 1, 0, 0, 0, 214, 37, 1, 0, 0, 0, 215, 216, 5, 91, 0, 0,
		216, 217, 6, 18, 1, 0, 217, 39, 1, 0, 0, 0, 218, 219, 5, 93, 0, 0, 219,
		220, 6, 19, 2, 0, 220, 41, 1, 0, 0, 0, 221, 223, 7, 13, 0, 0, 222, 221,
		1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0,
		0, 0, 225, 232, 1, 0, 0, 0, 226, 228, 5, 46, 0, 0, 227, 229, 7, 13, 0,
		0, 228, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230,
		231, 1, 0, 0, 0, 231, 233, 1, 0, 0, 0, 232, 226, 1, 0, 0, 0, 232, 233,
		1, 0, 0, 0, 233, 43, 1, 0, 0, 0, 234, 240, 3, 42, 20, 0, 235, 237, 7, 9,
		0, 0, 236, 238, 7, 19, 0, 0, 237, 236, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0,
		238, 239, 1, 0, 0, 0, 239, 241, 3, 42, 20, 0, 240, 235, 1, 0, 0, 0, 240,
		241, 1, 0, 0, 0, 241, 45, 1, 0, 0, 0, 242, 245, 3, 42, 20, 0, 243, 245,
		3, 44, 21, 0, 244, 242, 1, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 47, 1, 0,
		0, 0, 246, 252, 5, 39, 0, 0, 247, 251, 8, 20, 0, 0, 248, 249, 5, 92, 0,
		0, 249, 251, 9, 0, 0, 0, 250, 247, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251,
		254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255,
		1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 267, 5, 39, 0, 0, 256, 262, 5, 34,
		0, 0, 257, 261, 8, 21, 0, 0, 258, 259, 5, 92, 0, 0, 259, 261, 9, 0, 0,
		0, 260, 257, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 261, 264, 1, 0, 0, 0, 262,
		260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0, 0, 0, 264, 262,
		1, 0, 0, 0, 265, 267, 5, 34, 0, 0, 266, 246, 1, 0, 0, 0, 266, 256, 1, 0,
		0, 0, 267, 49, 1, 0, 0, 0, 268, 269, 5, 96, 0, 0, 269, 270, 1, 0, 0, 0,
		270, 271, 6, 24, 3, 0, 271, 272, 6, 24, 4, 0, 272, 51, 1, 0, 0, 0, 273,
		274, 5, 43, 0, 0, 274, 53, 1, 0, 0, 0, 275, 276, 5, 45, 0, 0, 276, 55,
		1, 0, 0, 0, 277, 278, 5, 42, 0, 0, 278, 57, 1, 0, 0, 0, 279, 280, 5, 47,
		0, 0, 280, 59, 1, 0, 0, 0, 281, 282, 5, 37, 0, 0, 282, 61, 1, 0, 0, 0,
		283, 284, 5, 94, 0, 0, 284, 63, 1, 0, 0, 0, 285, 286, 7, 10, 0, 0, 286,
		287, 7, 22, 0, 0, 287, 288, 7, 18, 0, 0, 288, 65, 1, 0, 0, 0, 289, 290,
		7, 23, 0, 0, 290, 291, 7, 7, 0, 0, 291, 67, 1, 0, 0, 0, 292, 293, 7, 8,
		0, 0, 293, 294, 7, 22, 0, 0, 294, 295, 7, 11, 0, 0, 295, 296, 7, 9, 0,
		0, 296, 297, 7, 12, 0, 0, 297, 298, 7, 12, 0, 0, 298, 69, 1, 0, 0, 0, 299,
		300, 5, 61, 0, 0, 300, 71, 1, 0, 0, 0, 301, 302, 5, 61, 0, 0, 302, 303,
		5, 61, 0, 0, 303, 73, 1, 0, 0, 0, 304, 305, 5, 33, 0, 0, 305, 306, 5, 61,
		0, 0, 306, 75, 1, 0, 0, 0, 307, 308, 5, 62, 0, 0, 308, 77, 1, 0, 0, 0,
		309, 310, 5, 60, 0, 0, 310, 79, 1, 0, 0, 0, 311, 312, 5, 62, 0, 0, 312,
		313, 5, 61, 0, 0, 313, 81, 1, 0, 0, 0, 314, 315, 5, 60, 0, 0, 315, 316,
		5, 61, 0, 0, 316, 83, 1, 0, 0, 0, 317, 318, 5, 61, 0, 0, 318, 319, 5, 126,
		0, 0, 319, 85, 1, 0, 0, 0, 320, 321, 5, 33, 0, 0, 321, 322, 5, 126, 0,
		0, 322, 87, 1, 0, 0, 0, 323, 324, 7, 17, 0, 0, 324, 325, 7, 24, 0, 0, 325,
		89, 1, 0, 0, 0, 326, 327, 7, 25, 0, 0, 327, 328, 7, 4, 0, 0, 328, 329,
		7, 6, 0, 0, 329, 330, 7, 26, 0, 0, 330, 331, 7, 23, 0, 0, 331, 332, 7,
		8, 0, 0, 332, 333, 7, 6, 0, 0, 333, 91, 1, 0, 0, 0, 334, 335, 7, 23, 0,
		0, 335, 336, 7, 22, 0, 0, 336, 93, 1, 0, 0, 0, 337, 338, 7, 4, 0, 0, 338,
		339, 7, 27, 0, 0, 339, 340, 7, 22, 0, 0, 340, 341, 7, 23, 0, 0, 341, 342,
		7, 7, 0, 0, 342, 343, 7, 4, 0, 0, 343, 344, 7, 22, 0, 0, 344, 345, 7, 27,
		0, 0, 345, 95, 1, 0, 0, 0, 346, 347, 7, 27, 0, 0, 347, 348, 7, 7, 0, 0,
		348, 349, 7, 23, 0, 0, 349, 350, 7, 8, 0, 0, 350, 351, 7, 28, 0, 0, 351,
		352, 5, 95, 0, 0, 352, 353, 7, 11, 0, 0, 353, 354, 7, 9, 0, 0, 354, 355,
		7, 5, 0, 0, 355, 356, 7, 6, 0, 0, 356, 97, 1, 0, 0, 0, 357, 358, 7, 27,
		0, 0, 358, 359, 7, 7, 0, 0, 359, 360, 7, 23, 0, 0, 360, 361, 7, 8, 0, 0,
		361, 362, 7, 28, 0, 0, 362, 363, 5, 95, 0, 0, 363, 364, 7, 7, 0, 0, 364,
		365, 7, 4, 0, 0, 365, 366, 7, 27, 0, 0, 366, 367, 7, 26, 0, 0, 367, 368,
		7, 6, 0, 0, 368, 99, 1, 0, 0, 0, 369, 370, 7, 23, 0, 0, 370, 371, 7, 5,
		0, 0, 371, 372, 7, 5, 0, 0, 372, 373, 7, 12, 0, 0, 373, 374, 7, 9, 0, 0,
		374, 375, 7, 6, 0, 0, 375, 101, 1, 0, 0, 0, 376, 377, 7, 17, 0, 0, 377,
		378, 7, 23, 0, 0, 378, 379, 7, 23, 0, 0, 379, 380, 7, 11, 0, 0, 380, 103,
		1, 0, 0, 0, 381, 382, 5, 123, 0, 0, 382, 105, 1, 0, 0, 0, 383, 384, 5,
		125, 0, 0, 384, 107, 1, 0, 0, 0, 385, 386, 5, 40, 0, 0, 386, 109, 1, 0,
		0, 0, 387, 388, 5, 41, 0, 0, 388, 111, 1, 0, 0, 0, 389, 390, 5, 44, 0,
		0, 390, 113, 1, 0, 0, 0, 391, 392, 5, 64, 0, 0, 392, 115, 1, 0, 0, 0, 393,
		395, 7, 13, 0, 0, 394, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 394,
		1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 401, 1, 0, 0, 0, 398, 399, 7, 15,
		0, 0, 399, 402, 7, 12, 0, 0, 400, 402, 7, 29, 0, 0, 401, 398, 1, 0, 0,
		0, 401, 400, 1, 0, 0, 0, 402, 404, 1, 0, 0, 0, 403, 394, 1, 0, 0, 0, 404,
		405, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 117,
		1, 0, 0, 0, 407, 411, 7, 30, 0, 0, 408, 410, 7, 1, 0, 0, 409, 408, 1, 0,
		0, 0, 410, 413, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0,
		412, 119, 1, 0, 0, 0, 413, 411, 1, 0, 0, 0, 414, 416, 7, 31, 0, 0, 415,
		414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 417, 418,
		1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 6, 59, 5, 0, 420, 121, 1, 0,
		0, 0, 421, 425, 5, 35, 0, 0, 422, 424, 9, 0, 0, 0, 423, 422, 1, 0, 0, 0,
		424, 427, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426,
		428, 1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 428, 429, 5, 10, 0, 0, 429, 430,
		1, 0, 0, 0, 430, 431, 6, 60, 6, 0, 431, 123, 1, 0, 0, 0, 432, 436, 8, 32,
		0, 0, 433, 434, 5, 92, 0, 0, 434, 436, 5, 96, 0, 0, 435, 432, 1, 0, 0,
		0, 435, 433, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437,
		438, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 440, 6, 61, 3, 0, 440, 125,
		1, 0, 0, 0, 441, 442, 5, 96, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444, 6, 62,
		7, 0, 444, 127, 1, 0, 0, 0, 28, 0, 1, 136, 146, 173, 178, 181, 189, 213,
		224, 230, 232, 237, 240, 244, 250, 252, 260, 262, 266, 396, 401, 405, 411,
		417, 425, 435, 437, 8, 1, 2, 0, 1, 18, 1, 1, 19, 2, 3, 0, 0, 5, 1, 0, 0,
		2, 0, 0, 3, 0, 4, 0, 0,
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
	PromQLExLexerID                   = 3
	PromQLExLexerMETRIC_NAME          = 4
	PromQLExLexerIF                   = 5
	PromQLExLexerTRUE                 = 6
	PromQLExLexerFALSE                = 7
	PromQLExLexerT                    = 8
	PromQLExLexerCOLON                = 9
	PromQLExLexerDOT                  = 10
	PromQLExLexerPOSITIVE_INTEGER     = 11
	PromQLExLexerTWO_DIGITS           = 12
	PromQLExLexerDIGITS               = 13
	PromQLExLexerMETRIC_KEYWORD       = 14
	PromQLExLexerLABEL_KEYWORD        = 15
	PromQLExLexerDEF                  = 16
	PromQLExLexerCALL_SIGN            = 17
	PromQLExLexerNL                   = 18
	PromQLExLexerLEFT_BRACKET         = 19
	PromQLExLexerRIGHT_BRACKET        = 20
	PromQLExLexerNUMBER               = 21
	PromQLExLexerSTRING               = 22
	PromQLExLexerADD                  = 23
	PromQLExLexerSUB                  = 24
	PromQLExLexerMULT                 = 25
	PromQLExLexerDIV                  = 26
	PromQLExLexerMOD                  = 27
	PromQLExLexerPOW                  = 28
	PromQLExLexerAND                  = 29
	PromQLExLexerOR                   = 30
	PromQLExLexerUNLESS               = 31
	PromQLExLexerEQ                   = 32
	PromQLExLexerDEQ                  = 33
	PromQLExLexerNE                   = 34
	PromQLExLexerGT                   = 35
	PromQLExLexerLT                   = 36
	PromQLExLexerGE                   = 37
	PromQLExLexerLE                   = 38
	PromQLExLexerRE                   = 39
	PromQLExLexerNRE                  = 40
	PromQLExLexerBY                   = 41
	PromQLExLexerWITHOUT              = 42
	PromQLExLexerON                   = 43
	PromQLExLexerIGNORING             = 44
	PromQLExLexerGROUP_LEFT           = 45
	PromQLExLexerGROUP_RIGHT          = 46
	PromQLExLexerOFFSET               = 47
	PromQLExLexerBOOL                 = 48
	PromQLExLexerLEFT_BRACE           = 49
	PromQLExLexerRIGHT_BRACE          = 50
	PromQLExLexerLEFT_PAREN           = 51
	PromQLExLexerRIGHT_PAREN          = 52
	PromQLExLexerCOMMA                = 53
	PromQLExLexerAT                   = 54
	PromQLExLexerDURATION             = 55
	PromQLExLexerLABEL_NAME           = 56
	PromQLExLexerWS                   = 57
	PromQLExLexerSL_COMMENT           = 58
	PromQLExLexerRAW_STRING           = 59
	PromQLExLexerBACKTICK_OPEN        = 60
)

// PromQLExLexer escapedChannels.
const (
	PromQLExLexerWHITESPACE = 2
	PromQLExLexerCOMMENTS   = 3
)

// PromQLExLexerRAW_STRING_MODE is the PromQLExLexer mode.
const PromQLExLexerRAW_STRING_MODE = 1

type BracketCounter interface {
	BracketCount() int
	SetBracketCount(val int)
}

type FunctionsProvider interface {
	GetTokenType(text string) (int, bool)
}

func (l *PromQLExLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 2:
		l.ID_Action(localctx, actionIndex)

	case 18:
		l.LEFT_BRACKET_Action(localctx, actionIndex)

	case 19:
		l.RIGHT_BRACKET_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PromQLExLexer) ID_Action(localctx antlr.RuleContext, actionIndex int) {
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
	case 3:
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
			return !ok || (ok && cnt.BracketCount() == 0)
		}()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

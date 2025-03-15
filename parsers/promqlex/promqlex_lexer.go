// Code generated from PromQLExLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
	"sync"
	"unicode"
)

import "slices"

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
		4, 0, 60, 446, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
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
		2, 12, 2, 138, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 146, 8,
		3, 10, 3, 12, 3, 149, 9, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 4, 10, 173, 8, 10, 11, 10, 12, 10, 174, 1, 10, 5, 10,
		178, 8, 10, 10, 10, 12, 10, 181, 9, 10, 3, 10, 183, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 12, 4, 12, 189, 8, 12, 11, 12, 12, 12, 190, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 215,
		8, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 4, 20, 224, 8,
		20, 11, 20, 12, 20, 225, 1, 20, 1, 20, 4, 20, 230, 8, 20, 11, 20, 12, 20,
		231, 3, 20, 234, 8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 239, 8, 21, 1, 21,
		3, 21, 242, 8, 21, 1, 22, 1, 22, 3, 22, 246, 8, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 5, 23, 252, 8, 23, 10, 23, 12, 23, 255, 9, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 5, 23, 262, 8, 23, 10, 23, 12, 23, 265, 9, 23, 1, 23,
		3, 23, 268, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55,
		1, 56, 1, 56, 1, 57, 4, 57, 396, 8, 57, 11, 57, 12, 57, 397, 1, 57, 1,
		57, 1, 57, 3, 57, 403, 8, 57, 4, 57, 405, 8, 57, 11, 57, 12, 57, 406, 1,
		58, 1, 58, 5, 58, 411, 8, 58, 10, 58, 12, 58, 414, 9, 58, 1, 59, 4, 59,
		417, 8, 59, 11, 59, 12, 59, 418, 1, 59, 1, 59, 1, 60, 1, 60, 5, 60, 425,
		8, 60, 10, 60, 12, 60, 428, 9, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1,
		61, 1, 61, 4, 61, 437, 8, 61, 11, 61, 12, 61, 438, 1, 61, 1, 61, 1, 62,
		1, 62, 1, 62, 1, 62, 1, 426, 0, 63, 2, 0, 4, 0, 6, 3, 8, 4, 10, 5, 12,
		6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28, 14, 30, 15,
		32, 16, 34, 17, 36, 18, 38, 19, 40, 20, 42, 0, 44, 0, 46, 21, 48, 22, 50,
		60, 52, 23, 54, 24, 56, 25, 58, 26, 60, 27, 62, 28, 64, 29, 66, 30, 68,
		31, 70, 32, 72, 33, 74, 34, 76, 35, 78, 36, 80, 37, 82, 38, 84, 39, 86,
		40, 88, 41, 90, 42, 92, 43, 94, 44, 96, 45, 98, 46, 100, 47, 102, 48, 104,
		49, 106, 50, 108, 51, 110, 52, 112, 53, 114, 54, 116, 55, 118, 56, 120,
		57, 122, 58, 124, 0, 126, 59, 2, 0, 1, 33, 2, 0, 65, 90, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 58, 65, 90, 95, 95, 97, 122, 2, 0, 73, 73, 105, 105, 2, 0, 70,
		70, 102, 102, 2, 0, 84, 84, 116, 116, 2, 0, 82, 82, 114, 114, 2, 0, 85,
		85, 117, 117, 2, 0, 69, 69, 101, 101, 2, 0, 65, 65, 97, 97, 2, 0, 76, 76,
		108, 108, 2, 0, 83, 83, 115, 115, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 77,
		77, 109, 109, 2, 0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 68, 68,
		100, 100, 2, 0, 43, 43, 45, 45, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92,
		92, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 89, 89, 121,
		121, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 71, 71, 103,
		103, 2, 0, 80, 80, 112, 112, 12, 0, 68, 68, 72, 72, 77, 77, 83, 83, 87,
		87, 89, 89, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119, 121, 121,
		3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 92, 92,
		96, 96, 466, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0,
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
		0, 0, 4, 130, 1, 0, 0, 0, 6, 132, 1, 0, 0, 0, 8, 142, 1, 0, 0, 0, 10, 150,
		1, 0, 0, 0, 12, 153, 1, 0, 0, 0, 14, 158, 1, 0, 0, 0, 16, 164, 1, 0, 0,
		0, 18, 166, 1, 0, 0, 0, 20, 168, 1, 0, 0, 0, 22, 182, 1, 0, 0, 0, 24, 184,
		1, 0, 0, 0, 26, 188, 1, 0, 0, 0, 28, 192, 1, 0, 0, 0, 30, 199, 1, 0, 0,
		0, 32, 205, 1, 0, 0, 0, 34, 209, 1, 0, 0, 0, 36, 214, 1, 0, 0, 0, 38, 216,
		1, 0, 0, 0, 40, 219, 1, 0, 0, 0, 42, 223, 1, 0, 0, 0, 44, 235, 1, 0, 0,
		0, 46, 245, 1, 0, 0, 0, 48, 267, 1, 0, 0, 0, 50, 269, 1, 0, 0, 0, 52, 274,
		1, 0, 0, 0, 54, 276, 1, 0, 0, 0, 56, 278, 1, 0, 0, 0, 58, 280, 1, 0, 0,
		0, 60, 282, 1, 0, 0, 0, 62, 284, 1, 0, 0, 0, 64, 286, 1, 0, 0, 0, 66, 290,
		1, 0, 0, 0, 68, 293, 1, 0, 0, 0, 70, 300, 1, 0, 0, 0, 72, 302, 1, 0, 0,
		0, 74, 305, 1, 0, 0, 0, 76, 308, 1, 0, 0, 0, 78, 310, 1, 0, 0, 0, 80, 312,
		1, 0, 0, 0, 82, 315, 1, 0, 0, 0, 84, 318, 1, 0, 0, 0, 86, 321, 1, 0, 0,
		0, 88, 324, 1, 0, 0, 0, 90, 327, 1, 0, 0, 0, 92, 335, 1, 0, 0, 0, 94, 338,
		1, 0, 0, 0, 96, 347, 1, 0, 0, 0, 98, 358, 1, 0, 0, 0, 100, 370, 1, 0, 0,
		0, 102, 377, 1, 0, 0, 0, 104, 382, 1, 0, 0, 0, 106, 384, 1, 0, 0, 0, 108,
		386, 1, 0, 0, 0, 110, 388, 1, 0, 0, 0, 112, 390, 1, 0, 0, 0, 114, 392,
		1, 0, 0, 0, 116, 404, 1, 0, 0, 0, 118, 408, 1, 0, 0, 0, 120, 416, 1, 0,
		0, 0, 122, 422, 1, 0, 0, 0, 124, 436, 1, 0, 0, 0, 126, 442, 1, 0, 0, 0,
		128, 129, 1, 0, 0, 0, 129, 3, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 5,
		1, 0, 0, 0, 132, 136, 7, 0, 0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0,
		0, 0, 135, 138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0,
		137, 139, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 4, 2, 0, 0, 140,
		141, 6, 2, 0, 0, 141, 7, 1, 0, 0, 0, 142, 143, 4, 3, 1, 0, 143, 147, 7,
		2, 0, 0, 144, 146, 7, 3, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0,
		0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 9, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 150, 151, 7, 4, 0, 0, 151, 152, 7, 5, 0, 0, 152, 11, 1,
		0, 0, 0, 153, 154, 7, 6, 0, 0, 154, 155, 7, 7, 0, 0, 155, 156, 7, 8, 0,
		0, 156, 157, 7, 9, 0, 0, 157, 13, 1, 0, 0, 0, 158, 159, 7, 5, 0, 0, 159,
		160, 7, 10, 0, 0, 160, 161, 7, 11, 0, 0, 161, 162, 7, 12, 0, 0, 162, 163,
		7, 9, 0, 0, 163, 15, 1, 0, 0, 0, 164, 165, 7, 6, 0, 0, 165, 17, 1, 0, 0,
		0, 166, 167, 5, 58, 0, 0, 167, 19, 1, 0, 0, 0, 168, 169, 5, 46, 0, 0, 169,
		21, 1, 0, 0, 0, 170, 183, 7, 13, 0, 0, 171, 173, 7, 14, 0, 0, 172, 171,
		1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0,
		0, 0, 175, 179, 1, 0, 0, 0, 176, 178, 7, 13, 0, 0, 177, 176, 1, 0, 0, 0,
		178, 181, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 170, 1, 0, 0, 0, 182, 172,
		1, 0, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 7, 13, 0, 0, 185, 186, 7, 13,
		0, 0, 186, 25, 1, 0, 0, 0, 187, 189, 7, 13, 0, 0, 188, 187, 1, 0, 0, 0,
		189, 190, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191,
		27, 1, 0, 0, 0, 192, 193, 7, 15, 0, 0, 193, 194, 7, 9, 0, 0, 194, 195,
		7, 6, 0, 0, 195, 196, 7, 7, 0, 0, 196, 197, 7, 4, 0, 0, 197, 198, 7, 16,
		0, 0, 198, 29, 1, 0, 0, 0, 199, 200, 7, 11, 0, 0, 200, 201, 7, 10, 0, 0,
		201, 202, 7, 17, 0, 0, 202, 203, 7, 9, 0, 0, 203, 204, 7, 11, 0, 0, 204,
		31, 1, 0, 0, 0, 205, 206, 7, 18, 0, 0, 206, 207, 7, 9, 0, 0, 207, 208,
		7, 5, 0, 0, 208, 33, 1, 0, 0, 0, 209, 210, 5, 36, 0, 0, 210, 35, 1, 0,
		0, 0, 211, 215, 5, 10, 0, 0, 212, 213, 5, 13, 0, 0, 213, 215, 5, 10, 0,
		0, 214, 211, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 215, 37, 1, 0, 0, 0, 216,
		217, 5, 91, 0, 0, 217, 218, 6, 18, 1, 0, 218, 39, 1, 0, 0, 0, 219, 220,
		5, 93, 0, 0, 220, 221, 6, 19, 2, 0, 221, 41, 1, 0, 0, 0, 222, 224, 7, 13,
		0, 0, 223, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0,
		225, 226, 1, 0, 0, 0, 226, 233, 1, 0, 0, 0, 227, 229, 5, 46, 0, 0, 228,
		230, 7, 13, 0, 0, 229, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 229,
		1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233, 227, 1, 0,
		0, 0, 233, 234, 1, 0, 0, 0, 234, 43, 1, 0, 0, 0, 235, 241, 3, 42, 20, 0,
		236, 238, 7, 9, 0, 0, 237, 239, 7, 19, 0, 0, 238, 237, 1, 0, 0, 0, 238,
		239, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 242, 3, 42, 20, 0, 241, 236,
		1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 45, 1, 0, 0, 0, 243, 246, 3, 42,
		20, 0, 244, 246, 3, 44, 21, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0, 0,
		0, 246, 47, 1, 0, 0, 0, 247, 253, 5, 39, 0, 0, 248, 252, 8, 20, 0, 0, 249,
		250, 5, 92, 0, 0, 250, 252, 9, 0, 0, 0, 251, 248, 1, 0, 0, 0, 251, 249,
		1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0,
		0, 0, 254, 256, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 268, 5, 39, 0, 0,
		257, 263, 5, 34, 0, 0, 258, 262, 8, 21, 0, 0, 259, 260, 5, 92, 0, 0, 260,
		262, 9, 0, 0, 0, 261, 258, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 265,
		1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 266, 1, 0,
		0, 0, 265, 263, 1, 0, 0, 0, 266, 268, 5, 34, 0, 0, 267, 247, 1, 0, 0, 0,
		267, 257, 1, 0, 0, 0, 268, 49, 1, 0, 0, 0, 269, 270, 5, 96, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 272, 6, 24, 3, 0, 272, 273, 6, 24, 4, 0, 273, 51,
		1, 0, 0, 0, 274, 275, 5, 43, 0, 0, 275, 53, 1, 0, 0, 0, 276, 277, 5, 45,
		0, 0, 277, 55, 1, 0, 0, 0, 278, 279, 5, 42, 0, 0, 279, 57, 1, 0, 0, 0,
		280, 281, 5, 47, 0, 0, 281, 59, 1, 0, 0, 0, 282, 283, 5, 37, 0, 0, 283,
		61, 1, 0, 0, 0, 284, 285, 5, 94, 0, 0, 285, 63, 1, 0, 0, 0, 286, 287, 7,
		10, 0, 0, 287, 288, 7, 22, 0, 0, 288, 289, 7, 18, 0, 0, 289, 65, 1, 0,
		0, 0, 290, 291, 7, 23, 0, 0, 291, 292, 7, 7, 0, 0, 292, 67, 1, 0, 0, 0,
		293, 294, 7, 8, 0, 0, 294, 295, 7, 22, 0, 0, 295, 296, 7, 11, 0, 0, 296,
		297, 7, 9, 0, 0, 297, 298, 7, 12, 0, 0, 298, 299, 7, 12, 0, 0, 299, 69,
		1, 0, 0, 0, 300, 301, 5, 61, 0, 0, 301, 71, 1, 0, 0, 0, 302, 303, 5, 61,
		0, 0, 303, 304, 5, 61, 0, 0, 304, 73, 1, 0, 0, 0, 305, 306, 5, 33, 0, 0,
		306, 307, 5, 61, 0, 0, 307, 75, 1, 0, 0, 0, 308, 309, 5, 62, 0, 0, 309,
		77, 1, 0, 0, 0, 310, 311, 5, 60, 0, 0, 311, 79, 1, 0, 0, 0, 312, 313, 5,
		62, 0, 0, 313, 314, 5, 61, 0, 0, 314, 81, 1, 0, 0, 0, 315, 316, 5, 60,
		0, 0, 316, 317, 5, 61, 0, 0, 317, 83, 1, 0, 0, 0, 318, 319, 5, 61, 0, 0,
		319, 320, 5, 126, 0, 0, 320, 85, 1, 0, 0, 0, 321, 322, 5, 33, 0, 0, 322,
		323, 5, 126, 0, 0, 323, 87, 1, 0, 0, 0, 324, 325, 7, 17, 0, 0, 325, 326,
		7, 24, 0, 0, 326, 89, 1, 0, 0, 0, 327, 328, 7, 25, 0, 0, 328, 329, 7, 4,
		0, 0, 329, 330, 7, 6, 0, 0, 330, 331, 7, 26, 0, 0, 331, 332, 7, 23, 0,
		0, 332, 333, 7, 8, 0, 0, 333, 334, 7, 6, 0, 0, 334, 91, 1, 0, 0, 0, 335,
		336, 7, 23, 0, 0, 336, 337, 7, 22, 0, 0, 337, 93, 1, 0, 0, 0, 338, 339,
		7, 4, 0, 0, 339, 340, 7, 27, 0, 0, 340, 341, 7, 22, 0, 0, 341, 342, 7,
		23, 0, 0, 342, 343, 7, 7, 0, 0, 343, 344, 7, 4, 0, 0, 344, 345, 7, 22,
		0, 0, 345, 346, 7, 27, 0, 0, 346, 95, 1, 0, 0, 0, 347, 348, 7, 27, 0, 0,
		348, 349, 7, 7, 0, 0, 349, 350, 7, 23, 0, 0, 350, 351, 7, 8, 0, 0, 351,
		352, 7, 28, 0, 0, 352, 353, 5, 95, 0, 0, 353, 354, 7, 11, 0, 0, 354, 355,
		7, 9, 0, 0, 355, 356, 7, 5, 0, 0, 356, 357, 7, 6, 0, 0, 357, 97, 1, 0,
		0, 0, 358, 359, 7, 27, 0, 0, 359, 360, 7, 7, 0, 0, 360, 361, 7, 23, 0,
		0, 361, 362, 7, 8, 0, 0, 362, 363, 7, 28, 0, 0, 363, 364, 5, 95, 0, 0,
		364, 365, 7, 7, 0, 0, 365, 366, 7, 4, 0, 0, 366, 367, 7, 27, 0, 0, 367,
		368, 7, 26, 0, 0, 368, 369, 7, 6, 0, 0, 369, 99, 1, 0, 0, 0, 370, 371,
		7, 23, 0, 0, 371, 372, 7, 5, 0, 0, 372, 373, 7, 5, 0, 0, 373, 374, 7, 12,
		0, 0, 374, 375, 7, 9, 0, 0, 375, 376, 7, 6, 0, 0, 376, 101, 1, 0, 0, 0,
		377, 378, 7, 17, 0, 0, 378, 379, 7, 23, 0, 0, 379, 380, 7, 23, 0, 0, 380,
		381, 7, 11, 0, 0, 381, 103, 1, 0, 0, 0, 382, 383, 5, 123, 0, 0, 383, 105,
		1, 0, 0, 0, 384, 385, 5, 125, 0, 0, 385, 107, 1, 0, 0, 0, 386, 387, 5,
		40, 0, 0, 387, 109, 1, 0, 0, 0, 388, 389, 5, 41, 0, 0, 389, 111, 1, 0,
		0, 0, 390, 391, 5, 44, 0, 0, 391, 113, 1, 0, 0, 0, 392, 393, 5, 64, 0,
		0, 393, 115, 1, 0, 0, 0, 394, 396, 7, 13, 0, 0, 395, 394, 1, 0, 0, 0, 396,
		397, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 402,
		1, 0, 0, 0, 399, 400, 7, 15, 0, 0, 400, 403, 7, 12, 0, 0, 401, 403, 7,
		29, 0, 0, 402, 399, 1, 0, 0, 0, 402, 401, 1, 0, 0, 0, 403, 405, 1, 0, 0,
		0, 404, 395, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406,
		407, 1, 0, 0, 0, 407, 117, 1, 0, 0, 0, 408, 412, 7, 30, 0, 0, 409, 411,
		7, 1, 0, 0, 410, 409, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0,
		0, 0, 412, 413, 1, 0, 0, 0, 413, 119, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0,
		415, 417, 7, 31, 0, 0, 416, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418,
		416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 421,
		6, 59, 5, 0, 421, 121, 1, 0, 0, 0, 422, 426, 5, 35, 0, 0, 423, 425, 9,
		0, 0, 0, 424, 423, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 427, 1, 0, 0,
		0, 426, 424, 1, 0, 0, 0, 427, 429, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 429,
		430, 5, 10, 0, 0, 430, 431, 1, 0, 0, 0, 431, 432, 6, 60, 6, 0, 432, 123,
		1, 0, 0, 0, 433, 437, 8, 32, 0, 0, 434, 435, 5, 92, 0, 0, 435, 437, 5,
		96, 0, 0, 436, 433, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 438, 1, 0, 0,
		0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 440, 1, 0, 0, 0, 440,
		441, 6, 61, 3, 0, 441, 125, 1, 0, 0, 0, 442, 443, 5, 96, 0, 0, 443, 444,
		1, 0, 0, 0, 444, 445, 6, 62, 7, 0, 445, 127, 1, 0, 0, 0, 28, 0, 1, 136,
		147, 174, 179, 182, 190, 214, 225, 231, 233, 238, 241, 245, 251, 253, 261,
		263, 267, 397, 402, 406, 412, 418, 426, 436, 438, 8, 1, 2, 0, 1, 18, 1,
		1, 19, 2, 3, 0, 0, 5, 1, 0, 0, 2, 0, 0, 3, 0, 4, 0, 0,
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
	case 2:
		return l.ID_Sempred(localctx, predIndex)

	case 3:
		return l.METRIC_NAME_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLExLexer) ID_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !p.IsLiteralName(p.GetText())

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PromQLExLexer) METRIC_NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return func() bool {
			cnt, ok := p.GetInputStream().(BracketCounter)
			return !ok || (ok && cnt.BracketCount() == 0)
		}()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

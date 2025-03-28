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
		"", "", "", "", "':'", "", "'if'", "'true'", "'false'", "';'", "'metric'",
		"'label'", "'def'", "'$'", "'['", "']'", "", "", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'^'", "'and'", "'or'", "'unless'", "'start'", "'end'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "RFC_3339_TIMESTAMP", "COLON",
		"METRIC_NAME", "IF", "TRUE", "FALSE", "SEMICOLON", "METRIC_KEYWORD",
		"LABEL_KEYWORD", "DEF", "CALL_SIGN", "LEFT_BRACKET", "RIGHT_BRACKET",
		"NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND",
		"OR", "UNLESS", "START", "END", "EQ", "DEQ", "NE", "GT", "LT", "GE",
		"LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT",
		"GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN",
		"RIGHT_PAREN", "COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT",
		"ID", "ID_WS", "RAW_STRING", "BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"SUBQUERY_RANGE", "TIME_RANGE", "RFC_3339_TIMESTAMP", "DIGIT", "T",
		"Z", "COLON", "DOT", "METRIC_NAME", "IF", "TRUE", "FALSE", "SEMICOLON",
		"METRIC_KEYWORD", "LABEL_KEYWORD", "DEF", "CALL_SIGN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "NUMERAL", "SCIENTIFIC_NUMBER", "NUMBER", "STRING",
		"BACKTICK_OPEN", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR",
		"UNLESS", "START", "END", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE",
		"RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT",
		"OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN",
		"COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "ID", "ID_WS",
		"RAW_STRING_CONTENT", "RAW_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 59, 499, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
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
		7, 61, 2, 62, 7, 62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 4,
		2, 161, 8, 2, 11, 2, 12, 2, 162, 3, 2, 165, 8, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 170, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 178, 8, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		5, 8, 193, 8, 8, 10, 8, 12, 8, 196, 9, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 4, 19, 247, 8, 19, 11, 19, 12, 19, 248, 1, 19, 1, 19, 4,
		19, 253, 8, 19, 11, 19, 12, 19, 254, 3, 19, 257, 8, 19, 1, 20, 1, 20, 1,
		20, 3, 20, 262, 8, 20, 1, 20, 3, 20, 265, 8, 20, 1, 21, 1, 21, 3, 21, 269,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 275, 8, 22, 10, 22, 12, 22, 278,
		9, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 285, 8, 22, 10, 22, 12,
		22, 288, 9, 22, 1, 22, 3, 22, 291, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46,
		1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 3, 58, 430, 8, 58, 1, 58,
		4, 58, 433, 8, 58, 11, 58, 12, 58, 434, 1, 58, 1, 58, 1, 58, 3, 58, 440,
		8, 58, 4, 58, 442, 8, 58, 11, 58, 12, 58, 443, 1, 59, 1, 59, 5, 59, 448,
		8, 59, 10, 59, 12, 59, 451, 9, 59, 1, 60, 4, 60, 454, 8, 60, 11, 60, 12,
		60, 455, 1, 60, 1, 60, 1, 61, 1, 61, 5, 61, 462, 8, 61, 10, 61, 12, 61,
		465, 9, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 5, 62, 473, 8, 62,
		10, 62, 12, 62, 476, 9, 62, 1, 62, 1, 62, 1, 63, 4, 63, 481, 8, 63, 11,
		63, 12, 63, 482, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 4, 64, 490, 8, 64,
		11, 64, 12, 64, 491, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 463,
		0, 66, 3, 0, 5, 0, 7, 3, 9, 0, 11, 0, 13, 0, 15, 4, 17, 0, 19, 5, 21, 6,
		23, 7, 25, 8, 27, 9, 29, 10, 31, 11, 33, 12, 35, 13, 37, 14, 39, 15, 41,
		0, 43, 0, 45, 16, 47, 17, 49, 59, 51, 18, 53, 19, 55, 20, 57, 21, 59, 22,
		61, 23, 63, 24, 65, 25, 67, 26, 69, 27, 71, 28, 73, 29, 75, 30, 77, 31,
		79, 32, 81, 33, 83, 34, 85, 35, 87, 36, 89, 37, 91, 38, 93, 39, 95, 40,
		97, 41, 99, 42, 101, 43, 103, 44, 105, 45, 107, 46, 109, 47, 111, 48, 113,
		49, 115, 50, 117, 51, 119, 52, 121, 53, 123, 54, 125, 55, 127, 56, 129,
		57, 131, 0, 133, 58, 3, 0, 1, 2, 33, 1, 0, 48, 57, 2, 0, 84, 84, 116, 116,
		2, 0, 90, 90, 122, 122, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		58, 65, 90, 95, 95, 97, 122, 2, 0, 73, 73, 105, 105, 2, 0, 70, 70, 102,
		102, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 69, 69, 101,
		101, 2, 0, 65, 65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 83, 83, 115, 115,
		2, 0, 77, 77, 109, 109, 2, 0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2,
		0, 68, 68, 100, 100, 2, 0, 43, 43, 45, 45, 2, 0, 39, 39, 92, 92, 2, 0,
		34, 34, 92, 92, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 89,
		89, 121, 121, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 71,
		71, 103, 103, 2, 0, 80, 80, 112, 112, 12, 0, 68, 68, 72, 72, 77, 77, 83,
		83, 87, 87, 89, 89, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119, 121,
		121, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 2, 0, 92, 92, 96, 96,
		516, 0, 7, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0,
		87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0,
		0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0,
		0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109,
		1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0,
		0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1,
		0, 0, 0, 0, 125, 1, 0, 0, 0, 1, 127, 1, 0, 0, 0, 1, 129, 1, 0, 0, 0, 2,
		131, 1, 0, 0, 0, 2, 133, 1, 0, 0, 0, 3, 135, 1, 0, 0, 0, 5, 137, 1, 0,
		0, 0, 7, 139, 1, 0, 0, 0, 9, 179, 1, 0, 0, 0, 11, 181, 1, 0, 0, 0, 13,
		183, 1, 0, 0, 0, 15, 185, 1, 0, 0, 0, 17, 187, 1, 0, 0, 0, 19, 189, 1,
		0, 0, 0, 21, 200, 1, 0, 0, 0, 23, 203, 1, 0, 0, 0, 25, 208, 1, 0, 0, 0,
		27, 214, 1, 0, 0, 0, 29, 216, 1, 0, 0, 0, 31, 223, 1, 0, 0, 0, 33, 229,
		1, 0, 0, 0, 35, 235, 1, 0, 0, 0, 37, 239, 1, 0, 0, 0, 39, 242, 1, 0, 0,
		0, 41, 246, 1, 0, 0, 0, 43, 258, 1, 0, 0, 0, 45, 268, 1, 0, 0, 0, 47, 290,
		1, 0, 0, 0, 49, 292, 1, 0, 0, 0, 51, 297, 1, 0, 0, 0, 53, 299, 1, 0, 0,
		0, 55, 301, 1, 0, 0, 0, 57, 303, 1, 0, 0, 0, 59, 305, 1, 0, 0, 0, 61, 307,
		1, 0, 0, 0, 63, 309, 1, 0, 0, 0, 65, 313, 1, 0, 0, 0, 67, 316, 1, 0, 0,
		0, 69, 323, 1, 0, 0, 0, 71, 329, 1, 0, 0, 0, 73, 333, 1, 0, 0, 0, 75, 335,
		1, 0, 0, 0, 77, 338, 1, 0, 0, 0, 79, 341, 1, 0, 0, 0, 81, 343, 1, 0, 0,
		0, 83, 345, 1, 0, 0, 0, 85, 348, 1, 0, 0, 0, 87, 351, 1, 0, 0, 0, 89, 354,
		1, 0, 0, 0, 91, 357, 1, 0, 0, 0, 93, 360, 1, 0, 0, 0, 95, 368, 1, 0, 0,
		0, 97, 371, 1, 0, 0, 0, 99, 380, 1, 0, 0, 0, 101, 391, 1, 0, 0, 0, 103,
		403, 1, 0, 0, 0, 105, 410, 1, 0, 0, 0, 107, 415, 1, 0, 0, 0, 109, 417,
		1, 0, 0, 0, 111, 419, 1, 0, 0, 0, 113, 421, 1, 0, 0, 0, 115, 423, 1, 0,
		0, 0, 117, 425, 1, 0, 0, 0, 119, 429, 1, 0, 0, 0, 121, 445, 1, 0, 0, 0,
		123, 453, 1, 0, 0, 0, 125, 459, 1, 0, 0, 0, 127, 470, 1, 0, 0, 0, 129,
		480, 1, 0, 0, 0, 131, 489, 1, 0, 0, 0, 133, 495, 1, 0, 0, 0, 135, 136,
		1, 0, 0, 0, 136, 4, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 6, 1, 0, 0,
		0, 139, 140, 3, 9, 3, 0, 140, 141, 3, 9, 3, 0, 141, 142, 3, 9, 3, 0, 142,
		143, 3, 9, 3, 0, 143, 144, 3, 53, 25, 0, 144, 145, 3, 9, 3, 0, 145, 146,
		3, 9, 3, 0, 146, 147, 3, 53, 25, 0, 147, 148, 3, 9, 3, 0, 148, 149, 3,
		9, 3, 0, 149, 150, 3, 11, 4, 0, 150, 151, 3, 9, 3, 0, 151, 152, 3, 9, 3,
		0, 152, 153, 3, 15, 6, 0, 153, 154, 3, 9, 3, 0, 154, 155, 3, 9, 3, 0, 155,
		156, 3, 15, 6, 0, 156, 157, 3, 9, 3, 0, 157, 164, 3, 9, 3, 0, 158, 160,
		3, 17, 7, 0, 159, 161, 3, 9, 3, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 1, 0, 0, 0,
		164, 158, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 177, 1, 0, 0, 0, 166,
		178, 3, 13, 5, 0, 167, 170, 3, 51, 24, 0, 168, 170, 3, 53, 25, 0, 169,
		167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172,
		3, 9, 3, 0, 172, 173, 3, 9, 3, 0, 173, 174, 3, 15, 6, 0, 174, 175, 3, 9,
		3, 0, 175, 176, 3, 9, 3, 0, 176, 178, 1, 0, 0, 0, 177, 166, 1, 0, 0, 0,
		177, 169, 1, 0, 0, 0, 178, 8, 1, 0, 0, 0, 179, 180, 7, 0, 0, 0, 180, 10,
		1, 0, 0, 0, 181, 182, 7, 1, 0, 0, 182, 12, 1, 0, 0, 0, 183, 184, 7, 2,
		0, 0, 184, 14, 1, 0, 0, 0, 185, 186, 5, 58, 0, 0, 186, 16, 1, 0, 0, 0,
		187, 188, 5, 46, 0, 0, 188, 18, 1, 0, 0, 0, 189, 190, 4, 8, 0, 0, 190,
		194, 7, 3, 0, 0, 191, 193, 7, 4, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196,
		1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 197, 198, 4, 8, 1, 0, 198, 199, 6, 8, 0, 0,
		199, 20, 1, 0, 0, 0, 200, 201, 7, 5, 0, 0, 201, 202, 7, 6, 0, 0, 202, 22,
		1, 0, 0, 0, 203, 204, 7, 1, 0, 0, 204, 205, 7, 7, 0, 0, 205, 206, 7, 8,
		0, 0, 206, 207, 7, 9, 0, 0, 207, 24, 1, 0, 0, 0, 208, 209, 7, 6, 0, 0,
		209, 210, 7, 10, 0, 0, 210, 211, 7, 11, 0, 0, 211, 212, 7, 12, 0, 0, 212,
		213, 7, 9, 0, 0, 213, 26, 1, 0, 0, 0, 214, 215, 5, 59, 0, 0, 215, 28, 1,
		0, 0, 0, 216, 217, 7, 13, 0, 0, 217, 218, 7, 9, 0, 0, 218, 219, 7, 1, 0,
		0, 219, 220, 7, 7, 0, 0, 220, 221, 7, 5, 0, 0, 221, 222, 7, 14, 0, 0, 222,
		30, 1, 0, 0, 0, 223, 224, 7, 11, 0, 0, 224, 225, 7, 10, 0, 0, 225, 226,
		7, 15, 0, 0, 226, 227, 7, 9, 0, 0, 227, 228, 7, 11, 0, 0, 228, 32, 1, 0,
		0, 0, 229, 230, 7, 16, 0, 0, 230, 231, 7, 9, 0, 0, 231, 232, 7, 6, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 234, 6, 15, 1, 0, 234, 34, 1, 0, 0, 0, 235,
		236, 5, 36, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 6, 16, 1, 0, 238, 36,
		1, 0, 0, 0, 239, 240, 5, 91, 0, 0, 240, 241, 6, 17, 2, 0, 241, 38, 1, 0,
		0, 0, 242, 243, 5, 93, 0, 0, 243, 244, 6, 18, 3, 0, 244, 40, 1, 0, 0, 0,
		245, 247, 7, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 256, 1, 0, 0, 0, 250, 252,
		5, 46, 0, 0, 251, 253, 7, 0, 0, 0, 252, 251, 1, 0, 0, 0, 253, 254, 1, 0,
		0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0,
		256, 250, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 42, 1, 0, 0, 0, 258, 264,
		3, 41, 19, 0, 259, 261, 7, 9, 0, 0, 260, 262, 7, 17, 0, 0, 261, 260, 1,
		0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 3, 41, 19,
		0, 264, 259, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 44, 1, 0, 0, 0, 266,
		269, 3, 41, 19, 0, 267, 269, 3, 43, 20, 0, 268, 266, 1, 0, 0, 0, 268, 267,
		1, 0, 0, 0, 269, 46, 1, 0, 0, 0, 270, 276, 5, 39, 0, 0, 271, 275, 8, 18,
		0, 0, 272, 273, 5, 92, 0, 0, 273, 275, 9, 0, 0, 0, 274, 271, 1, 0, 0, 0,
		274, 272, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276,
		277, 1, 0, 0, 0, 277, 279, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 291,
		5, 39, 0, 0, 280, 286, 5, 34, 0, 0, 281, 285, 8, 19, 0, 0, 282, 283, 5,
		92, 0, 0, 283, 285, 9, 0, 0, 0, 284, 281, 1, 0, 0, 0, 284, 282, 1, 0, 0,
		0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287,
		289, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289, 291, 5, 34, 0, 0, 290, 270,
		1, 0, 0, 0, 290, 280, 1, 0, 0, 0, 291, 48, 1, 0, 0, 0, 292, 293, 5, 96,
		0, 0, 293, 294, 1, 0, 0, 0, 294, 295, 6, 23, 4, 0, 295, 296, 6, 23, 5,
		0, 296, 50, 1, 0, 0, 0, 297, 298, 5, 43, 0, 0, 298, 52, 1, 0, 0, 0, 299,
		300, 5, 45, 0, 0, 300, 54, 1, 0, 0, 0, 301, 302, 5, 42, 0, 0, 302, 56,
		1, 0, 0, 0, 303, 304, 5, 47, 0, 0, 304, 58, 1, 0, 0, 0, 305, 306, 5, 37,
		0, 0, 306, 60, 1, 0, 0, 0, 307, 308, 5, 94, 0, 0, 308, 62, 1, 0, 0, 0,
		309, 310, 7, 10, 0, 0, 310, 311, 7, 20, 0, 0, 311, 312, 7, 16, 0, 0, 312,
		64, 1, 0, 0, 0, 313, 314, 7, 21, 0, 0, 314, 315, 7, 7, 0, 0, 315, 66, 1,
		0, 0, 0, 316, 317, 7, 8, 0, 0, 317, 318, 7, 20, 0, 0, 318, 319, 7, 11,
		0, 0, 319, 320, 7, 9, 0, 0, 320, 321, 7, 12, 0, 0, 321, 322, 7, 12, 0,
		0, 322, 68, 1, 0, 0, 0, 323, 324, 7, 12, 0, 0, 324, 325, 7, 1, 0, 0, 325,
		326, 7, 10, 0, 0, 326, 327, 7, 7, 0, 0, 327, 328, 7, 1, 0, 0, 328, 70,
		1, 0, 0, 0, 329, 330, 7, 9, 0, 0, 330, 331, 7, 20, 0, 0, 331, 332, 7, 16,
		0, 0, 332, 72, 1, 0, 0, 0, 333, 334, 5, 61, 0, 0, 334, 74, 1, 0, 0, 0,
		335, 336, 5, 61, 0, 0, 336, 337, 5, 61, 0, 0, 337, 76, 1, 0, 0, 0, 338,
		339, 5, 33, 0, 0, 339, 340, 5, 61, 0, 0, 340, 78, 1, 0, 0, 0, 341, 342,
		5, 62, 0, 0, 342, 80, 1, 0, 0, 0, 343, 344, 5, 60, 0, 0, 344, 82, 1, 0,
		0, 0, 345, 346, 5, 62, 0, 0, 346, 347, 5, 61, 0, 0, 347, 84, 1, 0, 0, 0,
		348, 349, 5, 60, 0, 0, 349, 350, 5, 61, 0, 0, 350, 86, 1, 0, 0, 0, 351,
		352, 5, 61, 0, 0, 352, 353, 5, 126, 0, 0, 353, 88, 1, 0, 0, 0, 354, 355,
		5, 33, 0, 0, 355, 356, 5, 126, 0, 0, 356, 90, 1, 0, 0, 0, 357, 358, 7,
		15, 0, 0, 358, 359, 7, 22, 0, 0, 359, 92, 1, 0, 0, 0, 360, 361, 7, 23,
		0, 0, 361, 362, 7, 5, 0, 0, 362, 363, 7, 1, 0, 0, 363, 364, 7, 24, 0, 0,
		364, 365, 7, 21, 0, 0, 365, 366, 7, 8, 0, 0, 366, 367, 7, 1, 0, 0, 367,
		94, 1, 0, 0, 0, 368, 369, 7, 21, 0, 0, 369, 370, 7, 20, 0, 0, 370, 96,
		1, 0, 0, 0, 371, 372, 7, 5, 0, 0, 372, 373, 7, 25, 0, 0, 373, 374, 7, 20,
		0, 0, 374, 375, 7, 21, 0, 0, 375, 376, 7, 7, 0, 0, 376, 377, 7, 5, 0, 0,
		377, 378, 7, 20, 0, 0, 378, 379, 7, 25, 0, 0, 379, 98, 1, 0, 0, 0, 380,
		381, 7, 25, 0, 0, 381, 382, 7, 7, 0, 0, 382, 383, 7, 21, 0, 0, 383, 384,
		7, 8, 0, 0, 384, 385, 7, 26, 0, 0, 385, 386, 5, 95, 0, 0, 386, 387, 7,
		11, 0, 0, 387, 388, 7, 9, 0, 0, 388, 389, 7, 6, 0, 0, 389, 390, 7, 1, 0,
		0, 390, 100, 1, 0, 0, 0, 391, 392, 7, 25, 0, 0, 392, 393, 7, 7, 0, 0, 393,
		394, 7, 21, 0, 0, 394, 395, 7, 8, 0, 0, 395, 396, 7, 26, 0, 0, 396, 397,
		5, 95, 0, 0, 397, 398, 7, 7, 0, 0, 398, 399, 7, 5, 0, 0, 399, 400, 7, 25,
		0, 0, 400, 401, 7, 24, 0, 0, 401, 402, 7, 1, 0, 0, 402, 102, 1, 0, 0, 0,
		403, 404, 7, 21, 0, 0, 404, 405, 7, 6, 0, 0, 405, 406, 7, 6, 0, 0, 406,
		407, 7, 12, 0, 0, 407, 408, 7, 9, 0, 0, 408, 409, 7, 1, 0, 0, 409, 104,
		1, 0, 0, 0, 410, 411, 7, 15, 0, 0, 411, 412, 7, 21, 0, 0, 412, 413, 7,
		21, 0, 0, 413, 414, 7, 11, 0, 0, 414, 106, 1, 0, 0, 0, 415, 416, 5, 123,
		0, 0, 416, 108, 1, 0, 0, 0, 417, 418, 5, 125, 0, 0, 418, 110, 1, 0, 0,
		0, 419, 420, 5, 40, 0, 0, 420, 112, 1, 0, 0, 0, 421, 422, 5, 41, 0, 0,
		422, 114, 1, 0, 0, 0, 423, 424, 5, 44, 0, 0, 424, 116, 1, 0, 0, 0, 425,
		426, 5, 64, 0, 0, 426, 118, 1, 0, 0, 0, 427, 430, 3, 51, 24, 0, 428, 430,
		3, 53, 25, 0, 429, 427, 1, 0, 0, 0, 429, 428, 1, 0, 0, 0, 429, 430, 1,
		0, 0, 0, 430, 441, 1, 0, 0, 0, 431, 433, 7, 0, 0, 0, 432, 431, 1, 0, 0,
		0, 433, 434, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435,
		439, 1, 0, 0, 0, 436, 437, 7, 13, 0, 0, 437, 440, 7, 12, 0, 0, 438, 440,
		7, 27, 0, 0, 439, 436, 1, 0, 0, 0, 439, 438, 1, 0, 0, 0, 440, 442, 1, 0,
		0, 0, 441, 432, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0,
		443, 444, 1, 0, 0, 0, 444, 120, 1, 0, 0, 0, 445, 449, 7, 28, 0, 0, 446,
		448, 7, 29, 0, 0, 447, 446, 1, 0, 0, 0, 448, 451, 1, 0, 0, 0, 449, 447,
		1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 122, 1, 0, 0, 0, 451, 449, 1, 0,
		0, 0, 452, 454, 7, 30, 0, 0, 453, 452, 1, 0, 0, 0, 454, 455, 1, 0, 0, 0,
		455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457,
		458, 6, 60, 6, 0, 458, 124, 1, 0, 0, 0, 459, 463, 5, 35, 0, 0, 460, 462,
		9, 0, 0, 0, 461, 460, 1, 0, 0, 0, 462, 465, 1, 0, 0, 0, 463, 464, 1, 0,
		0, 0, 463, 461, 1, 0, 0, 0, 464, 466, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0,
		466, 467, 5, 10, 0, 0, 467, 468, 1, 0, 0, 0, 468, 469, 6, 61, 7, 0, 469,
		126, 1, 0, 0, 0, 470, 474, 7, 31, 0, 0, 471, 473, 7, 29, 0, 0, 472, 471,
		1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0, 474, 475, 1, 0,
		0, 0, 475, 477, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 477, 478, 6, 62, 8, 0,
		478, 128, 1, 0, 0, 0, 479, 481, 7, 30, 0, 0, 480, 479, 1, 0, 0, 0, 481,
		482, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484,
		1, 0, 0, 0, 484, 485, 6, 63, 6, 0, 485, 130, 1, 0, 0, 0, 486, 490, 8, 32,
		0, 0, 487, 488, 5, 92, 0, 0, 488, 490, 5, 96, 0, 0, 489, 486, 1, 0, 0,
		0, 489, 487, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 491,
		492, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 494, 6, 64, 4, 0, 494, 132,
		1, 0, 0, 0, 495, 496, 5, 96, 0, 0, 496, 497, 1, 0, 0, 0, 497, 498, 6, 65,
		8, 0, 498, 134, 1, 0, 0, 0, 30, 0, 1, 2, 162, 164, 169, 177, 194, 248,
		254, 256, 261, 264, 268, 274, 276, 284, 286, 290, 429, 434, 439, 443, 449,
		455, 463, 474, 482, 489, 491, 9, 1, 8, 0, 5, 1, 0, 1, 17, 1, 1, 18, 2,
		3, 0, 0, 5, 2, 0, 0, 2, 0, 0, 3, 0, 4, 0, 0,
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
	PromQLExLexerRFC_3339_TIMESTAMP   = 3
	PromQLExLexerCOLON                = 4
	PromQLExLexerMETRIC_NAME          = 5
	PromQLExLexerIF                   = 6
	PromQLExLexerTRUE                 = 7
	PromQLExLexerFALSE                = 8
	PromQLExLexerSEMICOLON            = 9
	PromQLExLexerMETRIC_KEYWORD       = 10
	PromQLExLexerLABEL_KEYWORD        = 11
	PromQLExLexerDEF                  = 12
	PromQLExLexerCALL_SIGN            = 13
	PromQLExLexerLEFT_BRACKET         = 14
	PromQLExLexerRIGHT_BRACKET        = 15
	PromQLExLexerNUMBER               = 16
	PromQLExLexerSTRING               = 17
	PromQLExLexerADD                  = 18
	PromQLExLexerSUB                  = 19
	PromQLExLexerMULT                 = 20
	PromQLExLexerDIV                  = 21
	PromQLExLexerMOD                  = 22
	PromQLExLexerPOW                  = 23
	PromQLExLexerAND                  = 24
	PromQLExLexerOR                   = 25
	PromQLExLexerUNLESS               = 26
	PromQLExLexerSTART                = 27
	PromQLExLexerEND                  = 28
	PromQLExLexerEQ                   = 29
	PromQLExLexerDEQ                  = 30
	PromQLExLexerNE                   = 31
	PromQLExLexerGT                   = 32
	PromQLExLexerLT                   = 33
	PromQLExLexerGE                   = 34
	PromQLExLexerLE                   = 35
	PromQLExLexerRE                   = 36
	PromQLExLexerNRE                  = 37
	PromQLExLexerBY                   = 38
	PromQLExLexerWITHOUT              = 39
	PromQLExLexerON                   = 40
	PromQLExLexerIGNORING             = 41
	PromQLExLexerGROUP_LEFT           = 42
	PromQLExLexerGROUP_RIGHT          = 43
	PromQLExLexerOFFSET               = 44
	PromQLExLexerBOOL                 = 45
	PromQLExLexerLEFT_BRACE           = 46
	PromQLExLexerRIGHT_BRACE          = 47
	PromQLExLexerLEFT_PAREN           = 48
	PromQLExLexerRIGHT_PAREN          = 49
	PromQLExLexerCOMMA                = 50
	PromQLExLexerAT                   = 51
	PromQLExLexerDURATION             = 52
	PromQLExLexerLABEL_NAME           = 53
	PromQLExLexerWS                   = 54
	PromQLExLexerSL_COMMENT           = 55
	PromQLExLexerID                   = 56
	PromQLExLexerID_WS                = 57
	PromQLExLexerRAW_STRING           = 58
	PromQLExLexerBACKTICK_OPEN        = 59
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
	case 8:
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
	case 8:
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

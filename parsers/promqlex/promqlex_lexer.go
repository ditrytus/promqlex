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
		"'metric'", "'label'", "'def'", "'$'", "'['", "']'", "", "", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'", "'start'",
		"'end'", "'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'",
		"'!~'", "'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "METRIC_NAME", "IF", "TRUE",
		"FALSE", "T", "COLON", "SEMICOLON", "DOT", "METRIC_KEYWORD", "LABEL_KEYWORD",
		"DEF", "CALL_SIGN", "LEFT_BRACKET", "RIGHT_BRACKET", "NUMBER", "STRING",
		"ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR", "UNLESS", "START",
		"END", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE", "BY",
		"WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET",
		"BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN", "COMMA",
		"AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "ID", "RAW_STRING",
		"BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"SUBQUERY_RANGE", "TIME_RANGE", "METRIC_NAME", "IF", "TRUE", "FALSE",
		"T", "COLON", "SEMICOLON", "DOT", "METRIC_KEYWORD", "LABEL_KEYWORD",
		"DEF", "CALL_SIGN", "LEFT_BRACKET", "RIGHT_BRACKET", "NUMERAL", "SCIENTIFIC_NUMBER",
		"NUMBER", "STRING", "BACKTICK_OPEN", "ADD", "SUB", "MULT", "DIV", "MOD",
		"POW", "AND", "OR", "UNLESS", "START", "END", "EQ", "DEQ", "NE", "GT",
		"LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT",
		"GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN",
		"RIGHT_PAREN", "COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT",
		"ID", "RAW_STRING_CONTENT", "RAW_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 59, 436, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
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
		7, 61, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 135, 8, 2, 10, 2,
		12, 2, 138, 9, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 16, 4, 16, 195, 8, 16, 11, 16, 12, 16, 196, 1, 16, 1, 16, 4,
		16, 201, 8, 16, 11, 16, 12, 16, 202, 3, 16, 205, 8, 16, 1, 17, 1, 17, 1,
		17, 3, 17, 210, 8, 17, 1, 17, 3, 17, 213, 8, 17, 1, 18, 1, 18, 3, 18, 217,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 223, 8, 19, 10, 19, 12, 19, 226,
		9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 233, 8, 19, 10, 19, 12,
		19, 236, 9, 19, 1, 19, 3, 19, 239, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 4, 55, 377, 8, 55, 11, 55, 12, 55,
		378, 1, 55, 1, 55, 1, 55, 3, 55, 384, 8, 55, 4, 55, 386, 8, 55, 11, 55,
		12, 55, 387, 1, 56, 1, 56, 5, 56, 392, 8, 56, 10, 56, 12, 56, 395, 9, 56,
		1, 57, 4, 57, 398, 8, 57, 11, 57, 12, 57, 399, 1, 57, 1, 57, 1, 58, 1,
		58, 5, 58, 406, 8, 58, 10, 58, 12, 58, 409, 9, 58, 1, 58, 1, 58, 1, 58,
		1, 58, 1, 59, 1, 59, 5, 59, 417, 8, 59, 10, 59, 12, 59, 420, 9, 59, 1,
		59, 1, 59, 1, 60, 1, 60, 1, 60, 4, 60, 427, 8, 60, 11, 60, 12, 60, 428,
		1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 407, 0, 62, 3, 0, 5, 0, 7,
		3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27,
		13, 29, 14, 31, 15, 33, 16, 35, 0, 37, 0, 39, 17, 41, 18, 43, 59, 45, 19,
		47, 20, 49, 21, 51, 22, 53, 23, 55, 24, 57, 25, 59, 26, 61, 27, 63, 28,
		65, 29, 67, 30, 69, 31, 71, 32, 73, 33, 75, 34, 77, 35, 79, 36, 81, 37,
		83, 38, 85, 39, 87, 40, 89, 41, 91, 42, 93, 43, 95, 44, 97, 45, 99, 46,
		101, 47, 103, 48, 105, 49, 107, 50, 109, 51, 111, 52, 113, 53, 115, 54,
		117, 55, 119, 56, 121, 57, 123, 0, 125, 58, 3, 0, 1, 2, 32, 4, 0, 58, 58,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90, 95, 95, 97, 122, 2, 0, 73,
		73, 105, 105, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84, 116, 116, 2, 0, 82,
		82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 69, 69, 101, 101, 2, 0, 65,
		65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 83, 83, 115, 115, 2, 0, 77, 77,
		109, 109, 2, 0, 67, 67, 99, 99, 2, 0, 66, 66, 98, 98, 2, 0, 68, 68, 100,
		100, 1, 0, 48, 57, 2, 0, 43, 43, 45, 45, 2, 0, 39, 39, 92, 92, 2, 0, 34,
		34, 92, 92, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 89, 89,
		121, 121, 2, 0, 87, 87, 119, 119, 2, 0, 72, 72, 104, 104, 2, 0, 71, 71,
		103, 103, 2, 0, 80, 80, 112, 112, 12, 0, 68, 68, 72, 72, 77, 77, 83, 83,
		87, 87, 89, 89, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119, 121,
		121, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 2, 0, 92, 92, 96, 96,
		450, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0,
		87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0,
		0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0,
		0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109,
		1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0,
		0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 1, 121, 1, 0, 0, 0, 2, 123, 1,
		0, 0, 0, 2, 125, 1, 0, 0, 0, 3, 127, 1, 0, 0, 0, 5, 129, 1, 0, 0, 0, 7,
		131, 1, 0, 0, 0, 9, 142, 1, 0, 0, 0, 11, 145, 1, 0, 0, 0, 13, 150, 1, 0,
		0, 0, 15, 156, 1, 0, 0, 0, 17, 158, 1, 0, 0, 0, 19, 160, 1, 0, 0, 0, 21,
		162, 1, 0, 0, 0, 23, 164, 1, 0, 0, 0, 25, 171, 1, 0, 0, 0, 27, 177, 1,
		0, 0, 0, 29, 183, 1, 0, 0, 0, 31, 187, 1, 0, 0, 0, 33, 190, 1, 0, 0, 0,
		35, 194, 1, 0, 0, 0, 37, 206, 1, 0, 0, 0, 39, 216, 1, 0, 0, 0, 41, 238,
		1, 0, 0, 0, 43, 240, 1, 0, 0, 0, 45, 245, 1, 0, 0, 0, 47, 247, 1, 0, 0,
		0, 49, 249, 1, 0, 0, 0, 51, 251, 1, 0, 0, 0, 53, 253, 1, 0, 0, 0, 55, 255,
		1, 0, 0, 0, 57, 257, 1, 0, 0, 0, 59, 261, 1, 0, 0, 0, 61, 264, 1, 0, 0,
		0, 63, 271, 1, 0, 0, 0, 65, 277, 1, 0, 0, 0, 67, 281, 1, 0, 0, 0, 69, 283,
		1, 0, 0, 0, 71, 286, 1, 0, 0, 0, 73, 289, 1, 0, 0, 0, 75, 291, 1, 0, 0,
		0, 77, 293, 1, 0, 0, 0, 79, 296, 1, 0, 0, 0, 81, 299, 1, 0, 0, 0, 83, 302,
		1, 0, 0, 0, 85, 305, 1, 0, 0, 0, 87, 308, 1, 0, 0, 0, 89, 316, 1, 0, 0,
		0, 91, 319, 1, 0, 0, 0, 93, 328, 1, 0, 0, 0, 95, 339, 1, 0, 0, 0, 97, 351,
		1, 0, 0, 0, 99, 358, 1, 0, 0, 0, 101, 363, 1, 0, 0, 0, 103, 365, 1, 0,
		0, 0, 105, 367, 1, 0, 0, 0, 107, 369, 1, 0, 0, 0, 109, 371, 1, 0, 0, 0,
		111, 373, 1, 0, 0, 0, 113, 385, 1, 0, 0, 0, 115, 389, 1, 0, 0, 0, 117,
		397, 1, 0, 0, 0, 119, 403, 1, 0, 0, 0, 121, 414, 1, 0, 0, 0, 123, 426,
		1, 0, 0, 0, 125, 432, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 4, 1, 0, 0,
		0, 129, 130, 1, 0, 0, 0, 130, 6, 1, 0, 0, 0, 131, 132, 4, 2, 0, 0, 132,
		136, 7, 0, 0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138,
		1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 4, 2, 1, 0, 140, 141, 6, 2, 0, 0,
		141, 8, 1, 0, 0, 0, 142, 143, 7, 2, 0, 0, 143, 144, 7, 3, 0, 0, 144, 10,
		1, 0, 0, 0, 145, 146, 7, 4, 0, 0, 146, 147, 7, 5, 0, 0, 147, 148, 7, 6,
		0, 0, 148, 149, 7, 7, 0, 0, 149, 12, 1, 0, 0, 0, 150, 151, 7, 3, 0, 0,
		151, 152, 7, 8, 0, 0, 152, 153, 7, 9, 0, 0, 153, 154, 7, 10, 0, 0, 154,
		155, 7, 7, 0, 0, 155, 14, 1, 0, 0, 0, 156, 157, 7, 4, 0, 0, 157, 16, 1,
		0, 0, 0, 158, 159, 5, 58, 0, 0, 159, 18, 1, 0, 0, 0, 160, 161, 5, 59, 0,
		0, 161, 20, 1, 0, 0, 0, 162, 163, 5, 46, 0, 0, 163, 22, 1, 0, 0, 0, 164,
		165, 7, 11, 0, 0, 165, 166, 7, 7, 0, 0, 166, 167, 7, 4, 0, 0, 167, 168,
		7, 5, 0, 0, 168, 169, 7, 2, 0, 0, 169, 170, 7, 12, 0, 0, 170, 24, 1, 0,
		0, 0, 171, 172, 7, 9, 0, 0, 172, 173, 7, 8, 0, 0, 173, 174, 7, 13, 0, 0,
		174, 175, 7, 7, 0, 0, 175, 176, 7, 9, 0, 0, 176, 26, 1, 0, 0, 0, 177, 178,
		7, 14, 0, 0, 178, 179, 7, 7, 0, 0, 179, 180, 7, 3, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 182, 6, 12, 1, 0, 182, 28, 1, 0, 0, 0, 183, 184, 5, 36, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 186, 6, 13, 1, 0, 186, 30, 1, 0, 0, 0, 187,
		188, 5, 91, 0, 0, 188, 189, 6, 14, 2, 0, 189, 32, 1, 0, 0, 0, 190, 191,
		5, 93, 0, 0, 191, 192, 6, 15, 3, 0, 192, 34, 1, 0, 0, 0, 193, 195, 7, 15,
		0, 0, 194, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0,
		196, 197, 1, 0, 0, 0, 197, 204, 1, 0, 0, 0, 198, 200, 5, 46, 0, 0, 199,
		201, 7, 15, 0, 0, 200, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 200,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 198, 1, 0,
		0, 0, 204, 205, 1, 0, 0, 0, 205, 36, 1, 0, 0, 0, 206, 212, 3, 35, 16, 0,
		207, 209, 7, 7, 0, 0, 208, 210, 7, 16, 0, 0, 209, 208, 1, 0, 0, 0, 209,
		210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 213, 3, 35, 16, 0, 212, 207,
		1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 38, 1, 0, 0, 0, 214, 217, 3, 35,
		16, 0, 215, 217, 3, 37, 17, 0, 216, 214, 1, 0, 0, 0, 216, 215, 1, 0, 0,
		0, 217, 40, 1, 0, 0, 0, 218, 224, 5, 39, 0, 0, 219, 223, 8, 17, 0, 0, 220,
		221, 5, 92, 0, 0, 221, 223, 9, 0, 0, 0, 222, 219, 1, 0, 0, 0, 222, 220,
		1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0,
		0, 0, 225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 239, 5, 39, 0, 0,
		228, 234, 5, 34, 0, 0, 229, 233, 8, 18, 0, 0, 230, 231, 5, 92, 0, 0, 231,
		233, 9, 0, 0, 0, 232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 236,
		1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1, 0,
		0, 0, 236, 234, 1, 0, 0, 0, 237, 239, 5, 34, 0, 0, 238, 218, 1, 0, 0, 0,
		238, 228, 1, 0, 0, 0, 239, 42, 1, 0, 0, 0, 240, 241, 5, 96, 0, 0, 241,
		242, 1, 0, 0, 0, 242, 243, 6, 20, 4, 0, 243, 244, 6, 20, 5, 0, 244, 44,
		1, 0, 0, 0, 245, 246, 5, 43, 0, 0, 246, 46, 1, 0, 0, 0, 247, 248, 5, 45,
		0, 0, 248, 48, 1, 0, 0, 0, 249, 250, 5, 42, 0, 0, 250, 50, 1, 0, 0, 0,
		251, 252, 5, 47, 0, 0, 252, 52, 1, 0, 0, 0, 253, 254, 5, 37, 0, 0, 254,
		54, 1, 0, 0, 0, 255, 256, 5, 94, 0, 0, 256, 56, 1, 0, 0, 0, 257, 258, 7,
		8, 0, 0, 258, 259, 7, 19, 0, 0, 259, 260, 7, 14, 0, 0, 260, 58, 1, 0, 0,
		0, 261, 262, 7, 20, 0, 0, 262, 263, 7, 5, 0, 0, 263, 60, 1, 0, 0, 0, 264,
		265, 7, 6, 0, 0, 265, 266, 7, 19, 0, 0, 266, 267, 7, 9, 0, 0, 267, 268,
		7, 7, 0, 0, 268, 269, 7, 10, 0, 0, 269, 270, 7, 10, 0, 0, 270, 62, 1, 0,
		0, 0, 271, 272, 7, 10, 0, 0, 272, 273, 7, 4, 0, 0, 273, 274, 7, 8, 0, 0,
		274, 275, 7, 5, 0, 0, 275, 276, 7, 4, 0, 0, 276, 64, 1, 0, 0, 0, 277, 278,
		7, 7, 0, 0, 278, 279, 7, 19, 0, 0, 279, 280, 7, 14, 0, 0, 280, 66, 1, 0,
		0, 0, 281, 282, 5, 61, 0, 0, 282, 68, 1, 0, 0, 0, 283, 284, 5, 61, 0, 0,
		284, 285, 5, 61, 0, 0, 285, 70, 1, 0, 0, 0, 286, 287, 5, 33, 0, 0, 287,
		288, 5, 61, 0, 0, 288, 72, 1, 0, 0, 0, 289, 290, 5, 62, 0, 0, 290, 74,
		1, 0, 0, 0, 291, 292, 5, 60, 0, 0, 292, 76, 1, 0, 0, 0, 293, 294, 5, 62,
		0, 0, 294, 295, 5, 61, 0, 0, 295, 78, 1, 0, 0, 0, 296, 297, 5, 60, 0, 0,
		297, 298, 5, 61, 0, 0, 298, 80, 1, 0, 0, 0, 299, 300, 5, 61, 0, 0, 300,
		301, 5, 126, 0, 0, 301, 82, 1, 0, 0, 0, 302, 303, 5, 33, 0, 0, 303, 304,
		5, 126, 0, 0, 304, 84, 1, 0, 0, 0, 305, 306, 7, 13, 0, 0, 306, 307, 7,
		21, 0, 0, 307, 86, 1, 0, 0, 0, 308, 309, 7, 22, 0, 0, 309, 310, 7, 2, 0,
		0, 310, 311, 7, 4, 0, 0, 311, 312, 7, 23, 0, 0, 312, 313, 7, 20, 0, 0,
		313, 314, 7, 6, 0, 0, 314, 315, 7, 4, 0, 0, 315, 88, 1, 0, 0, 0, 316, 317,
		7, 20, 0, 0, 317, 318, 7, 19, 0, 0, 318, 90, 1, 0, 0, 0, 319, 320, 7, 2,
		0, 0, 320, 321, 7, 24, 0, 0, 321, 322, 7, 19, 0, 0, 322, 323, 7, 20, 0,
		0, 323, 324, 7, 5, 0, 0, 324, 325, 7, 2, 0, 0, 325, 326, 7, 19, 0, 0, 326,
		327, 7, 24, 0, 0, 327, 92, 1, 0, 0, 0, 328, 329, 7, 24, 0, 0, 329, 330,
		7, 5, 0, 0, 330, 331, 7, 20, 0, 0, 331, 332, 7, 6, 0, 0, 332, 333, 7, 25,
		0, 0, 333, 334, 5, 95, 0, 0, 334, 335, 7, 9, 0, 0, 335, 336, 7, 7, 0, 0,
		336, 337, 7, 3, 0, 0, 337, 338, 7, 4, 0, 0, 338, 94, 1, 0, 0, 0, 339, 340,
		7, 24, 0, 0, 340, 341, 7, 5, 0, 0, 341, 342, 7, 20, 0, 0, 342, 343, 7,
		6, 0, 0, 343, 344, 7, 25, 0, 0, 344, 345, 5, 95, 0, 0, 345, 346, 7, 5,
		0, 0, 346, 347, 7, 2, 0, 0, 347, 348, 7, 24, 0, 0, 348, 349, 7, 23, 0,
		0, 349, 350, 7, 4, 0, 0, 350, 96, 1, 0, 0, 0, 351, 352, 7, 20, 0, 0, 352,
		353, 7, 3, 0, 0, 353, 354, 7, 3, 0, 0, 354, 355, 7, 10, 0, 0, 355, 356,
		7, 7, 0, 0, 356, 357, 7, 4, 0, 0, 357, 98, 1, 0, 0, 0, 358, 359, 7, 13,
		0, 0, 359, 360, 7, 20, 0, 0, 360, 361, 7, 20, 0, 0, 361, 362, 7, 9, 0,
		0, 362, 100, 1, 0, 0, 0, 363, 364, 5, 123, 0, 0, 364, 102, 1, 0, 0, 0,
		365, 366, 5, 125, 0, 0, 366, 104, 1, 0, 0, 0, 367, 368, 5, 40, 0, 0, 368,
		106, 1, 0, 0, 0, 369, 370, 5, 41, 0, 0, 370, 108, 1, 0, 0, 0, 371, 372,
		5, 44, 0, 0, 372, 110, 1, 0, 0, 0, 373, 374, 5, 64, 0, 0, 374, 112, 1,
		0, 0, 0, 375, 377, 7, 15, 0, 0, 376, 375, 1, 0, 0, 0, 377, 378, 1, 0, 0,
		0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 383, 1, 0, 0, 0, 380,
		381, 7, 11, 0, 0, 381, 384, 7, 10, 0, 0, 382, 384, 7, 26, 0, 0, 383, 380,
		1, 0, 0, 0, 383, 382, 1, 0, 0, 0, 384, 386, 1, 0, 0, 0, 385, 376, 1, 0,
		0, 0, 386, 387, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0,
		388, 114, 1, 0, 0, 0, 389, 393, 7, 27, 0, 0, 390, 392, 7, 28, 0, 0, 391,
		390, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 393, 394,
		1, 0, 0, 0, 394, 116, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 396, 398, 7, 29,
		0, 0, 397, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0,
		399, 400, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 6, 57, 6, 0, 402,
		118, 1, 0, 0, 0, 403, 407, 5, 35, 0, 0, 404, 406, 9, 0, 0, 0, 405, 404,
		1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 407, 405, 1, 0,
		0, 0, 408, 410, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 411, 5, 10, 0, 0,
		411, 412, 1, 0, 0, 0, 412, 413, 6, 58, 7, 0, 413, 120, 1, 0, 0, 0, 414,
		418, 7, 30, 0, 0, 415, 417, 7, 28, 0, 0, 416, 415, 1, 0, 0, 0, 417, 420,
		1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 421, 1, 0,
		0, 0, 420, 418, 1, 0, 0, 0, 421, 422, 6, 59, 8, 0, 422, 122, 1, 0, 0, 0,
		423, 427, 8, 31, 0, 0, 424, 425, 5, 92, 0, 0, 425, 427, 5, 96, 0, 0, 426,
		423, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 426,
		1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 431, 6, 60,
		4, 0, 431, 124, 1, 0, 0, 0, 432, 433, 5, 96, 0, 0, 433, 434, 1, 0, 0, 0,
		434, 435, 6, 61, 8, 0, 435, 126, 1, 0, 0, 0, 24, 0, 1, 2, 136, 196, 202,
		204, 209, 212, 216, 222, 224, 232, 234, 238, 378, 383, 387, 393, 399, 407,
		418, 426, 428, 9, 1, 2, 0, 5, 1, 0, 1, 14, 1, 1, 15, 2, 3, 0, 0, 5, 2,
		0, 0, 2, 0, 0, 3, 0, 4, 0, 0,
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
	PromQLExLexerMETRIC_KEYWORD       = 11
	PromQLExLexerLABEL_KEYWORD        = 12
	PromQLExLexerDEF                  = 13
	PromQLExLexerCALL_SIGN            = 14
	PromQLExLexerLEFT_BRACKET         = 15
	PromQLExLexerRIGHT_BRACKET        = 16
	PromQLExLexerNUMBER               = 17
	PromQLExLexerSTRING               = 18
	PromQLExLexerADD                  = 19
	PromQLExLexerSUB                  = 20
	PromQLExLexerMULT                 = 21
	PromQLExLexerDIV                  = 22
	PromQLExLexerMOD                  = 23
	PromQLExLexerPOW                  = 24
	PromQLExLexerAND                  = 25
	PromQLExLexerOR                   = 26
	PromQLExLexerUNLESS               = 27
	PromQLExLexerSTART                = 28
	PromQLExLexerEND                  = 29
	PromQLExLexerEQ                   = 30
	PromQLExLexerDEQ                  = 31
	PromQLExLexerNE                   = 32
	PromQLExLexerGT                   = 33
	PromQLExLexerLT                   = 34
	PromQLExLexerGE                   = 35
	PromQLExLexerLE                   = 36
	PromQLExLexerRE                   = 37
	PromQLExLexerNRE                  = 38
	PromQLExLexerBY                   = 39
	PromQLExLexerWITHOUT              = 40
	PromQLExLexerON                   = 41
	PromQLExLexerIGNORING             = 42
	PromQLExLexerGROUP_LEFT           = 43
	PromQLExLexerGROUP_RIGHT          = 44
	PromQLExLexerOFFSET               = 45
	PromQLExLexerBOOL                 = 46
	PromQLExLexerLEFT_BRACE           = 47
	PromQLExLexerRIGHT_BRACE          = 48
	PromQLExLexerLEFT_PAREN           = 49
	PromQLExLexerRIGHT_PAREN          = 50
	PromQLExLexerCOMMA                = 51
	PromQLExLexerAT                   = 52
	PromQLExLexerDURATION             = 53
	PromQLExLexerLABEL_NAME           = 54
	PromQLExLexerWS                   = 55
	PromQLExLexerSL_COMMENT           = 56
	PromQLExLexerID                   = 57
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
	case 2:
		l.METRIC_NAME_Action(localctx, actionIndex)

	case 14:
		l.LEFT_BRACKET_Action(localctx, actionIndex)

	case 15:
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

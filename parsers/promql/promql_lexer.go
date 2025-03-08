// Code generated from PromQLLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promql

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

import (
	"slices"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PromQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PromQLLexerLexerStaticData struct {
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

func promqllexerLexerInit() {
	staticData := &PromQLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "WHITESPACE", "COMMENTS",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'",
		"'or'", "'unless'", "'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
		"'=~'", "'!~'", "'by'", "'without'", "'on'", "'ignoring'", "'group_left'",
		"'group_right'", "'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "'['",
		"']'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "NUMBER", "STRING", "ADD", "SUB",
		"MULT", "DIV", "MOD", "POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE",
		"GT", "LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING",
		"GROUP_LEFT", "GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE",
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA",
		"AT", "SUBQUERY_RANGE", "TIME_RANGE", "DURATION", "METRIC_NAME", "LABEL_NAME",
		"WS", "SL_COMMENT",
	}
	staticData.RuleNames = []string{
		"NUMERAL", "SCIENTIFIC_NUMBER", "NUMBER", "STRING", "ADD", "SUB", "MULT",
		"DIV", "MOD", "POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT",
		"LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT",
		"GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN",
		"RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA", "AT", "SUBQUERY_RANGE",
		"TIME_RANGE", "DURATION", "METRIC_NAME", "LABEL_NAME", "WS", "SL_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 45, 322, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 4, 0, 93, 8, 0, 11,
		0, 12, 0, 94, 1, 0, 1, 0, 4, 0, 99, 8, 0, 11, 0, 12, 0, 100, 3, 0, 103,
		8, 0, 1, 1, 1, 1, 1, 1, 3, 1, 108, 8, 1, 1, 1, 3, 1, 111, 8, 1, 1, 2, 1,
		2, 3, 2, 115, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 121, 8, 3, 10, 3, 12,
		3, 124, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 131, 8, 3, 10, 3, 12,
		3, 134, 9, 3, 1, 3, 3, 3, 137, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 267, 8, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 4, 40, 276, 8, 40, 11, 40, 12, 40,
		277, 1, 40, 1, 40, 1, 40, 3, 40, 283, 8, 40, 4, 40, 285, 8, 40, 11, 40,
		12, 40, 286, 1, 41, 1, 41, 5, 41, 291, 8, 41, 10, 41, 12, 41, 294, 9, 41,
		1, 41, 1, 41, 1, 42, 1, 42, 5, 42, 300, 8, 42, 10, 42, 12, 42, 303, 9,
		42, 1, 43, 4, 43, 306, 8, 43, 11, 43, 12, 43, 307, 1, 43, 1, 43, 1, 44,
		1, 44, 5, 44, 314, 8, 44, 10, 44, 12, 44, 317, 9, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 315, 0, 45, 1, 0, 3, 0, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 1, 0, 29, 1, 0, 48, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45,
		45, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 2, 0, 65, 65, 97, 97, 2,
		0, 78, 78, 110, 110, 2, 0, 68, 68, 100, 100, 2, 0, 79, 79, 111, 111, 2,
		0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 76, 76, 108, 108, 2,
		0, 83, 83, 115, 115, 2, 0, 66, 66, 98, 98, 2, 0, 89, 89, 121, 121, 2, 0,
		87, 87, 119, 119, 2, 0, 73, 73, 105, 105, 2, 0, 84, 84, 116, 116, 2, 0,
		72, 72, 104, 104, 2, 0, 71, 71, 103, 103, 2, 0, 80, 80, 112, 112, 2, 0,
		70, 70, 102, 102, 2, 0, 77, 77, 109, 109, 12, 0, 68, 68, 72, 72, 77, 77,
		83, 83, 87, 87, 89, 89, 100, 100, 104, 104, 109, 109, 115, 115, 119, 119,
		121, 121, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90,
		95, 95, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 338, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 92, 1,
		0, 0, 0, 3, 104, 1, 0, 0, 0, 5, 114, 1, 0, 0, 0, 7, 136, 1, 0, 0, 0, 9,
		138, 1, 0, 0, 0, 11, 140, 1, 0, 0, 0, 13, 142, 1, 0, 0, 0, 15, 144, 1,
		0, 0, 0, 17, 146, 1, 0, 0, 0, 19, 148, 1, 0, 0, 0, 21, 150, 1, 0, 0, 0,
		23, 154, 1, 0, 0, 0, 25, 157, 1, 0, 0, 0, 27, 164, 1, 0, 0, 0, 29, 166,
		1, 0, 0, 0, 31, 169, 1, 0, 0, 0, 33, 172, 1, 0, 0, 0, 35, 174, 1, 0, 0,
		0, 37, 176, 1, 0, 0, 0, 39, 179, 1, 0, 0, 0, 41, 182, 1, 0, 0, 0, 43, 185,
		1, 0, 0, 0, 45, 188, 1, 0, 0, 0, 47, 191, 1, 0, 0, 0, 49, 199, 1, 0, 0,
		0, 51, 202, 1, 0, 0, 0, 53, 211, 1, 0, 0, 0, 55, 222, 1, 0, 0, 0, 57, 234,
		1, 0, 0, 0, 59, 241, 1, 0, 0, 0, 61, 246, 1, 0, 0, 0, 63, 248, 1, 0, 0,
		0, 65, 250, 1, 0, 0, 0, 67, 252, 1, 0, 0, 0, 69, 254, 1, 0, 0, 0, 71, 256,
		1, 0, 0, 0, 73, 258, 1, 0, 0, 0, 75, 260, 1, 0, 0, 0, 77, 262, 1, 0, 0,
		0, 79, 270, 1, 0, 0, 0, 81, 284, 1, 0, 0, 0, 83, 288, 1, 0, 0, 0, 85, 297,
		1, 0, 0, 0, 87, 305, 1, 0, 0, 0, 89, 311, 1, 0, 0, 0, 91, 93, 7, 0, 0,
		0, 92, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 102, 1, 0, 0, 0, 96, 98, 5, 46, 0, 0, 97, 99, 7, 0, 0,
		0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 103, 1, 0,
		0, 0, 103, 2, 1, 0, 0, 0, 104, 110, 3, 1, 0, 0, 105, 107, 7, 1, 0, 0, 106,
		108, 7, 2, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 111, 3, 1, 0, 0, 110, 105, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 4, 1, 0, 0, 0, 112, 115, 3, 1, 0, 0, 113, 115, 3, 3, 1, 0, 114,
		112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 6, 1, 0, 0, 0, 116, 122, 5,
		39, 0, 0, 117, 121, 8, 3, 0, 0, 118, 119, 5, 92, 0, 0, 119, 121, 9, 0,
		0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0,
		122, 120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 125, 137, 5, 39, 0, 0, 126, 132, 5, 34, 0, 0, 127, 131,
		8, 4, 0, 0, 128, 129, 5, 92, 0, 0, 129, 131, 9, 0, 0, 0, 130, 127, 1, 0,
		0, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135,
		137, 5, 34, 0, 0, 136, 116, 1, 0, 0, 0, 136, 126, 1, 0, 0, 0, 137, 8, 1,
		0, 0, 0, 138, 139, 5, 43, 0, 0, 139, 10, 1, 0, 0, 0, 140, 141, 5, 45, 0,
		0, 141, 12, 1, 0, 0, 0, 142, 143, 5, 42, 0, 0, 143, 14, 1, 0, 0, 0, 144,
		145, 5, 47, 0, 0, 145, 16, 1, 0, 0, 0, 146, 147, 5, 37, 0, 0, 147, 18,
		1, 0, 0, 0, 148, 149, 5, 94, 0, 0, 149, 20, 1, 0, 0, 0, 150, 151, 7, 5,
		0, 0, 151, 152, 7, 6, 0, 0, 152, 153, 7, 7, 0, 0, 153, 22, 1, 0, 0, 0,
		154, 155, 7, 8, 0, 0, 155, 156, 7, 9, 0, 0, 156, 24, 1, 0, 0, 0, 157, 158,
		7, 10, 0, 0, 158, 159, 7, 6, 0, 0, 159, 160, 7, 11, 0, 0, 160, 161, 7,
		1, 0, 0, 161, 162, 7, 12, 0, 0, 162, 163, 7, 12, 0, 0, 163, 26, 1, 0, 0,
		0, 164, 165, 5, 61, 0, 0, 165, 28, 1, 0, 0, 0, 166, 167, 5, 61, 0, 0, 167,
		168, 5, 61, 0, 0, 168, 30, 1, 0, 0, 0, 169, 170, 5, 33, 0, 0, 170, 171,
		5, 61, 0, 0, 171, 32, 1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 34, 1, 0,
		0, 0, 174, 175, 5, 60, 0, 0, 175, 36, 1, 0, 0, 0, 176, 177, 5, 62, 0, 0,
		177, 178, 5, 61, 0, 0, 178, 38, 1, 0, 0, 0, 179, 180, 5, 60, 0, 0, 180,
		181, 5, 61, 0, 0, 181, 40, 1, 0, 0, 0, 182, 183, 5, 61, 0, 0, 183, 184,
		5, 126, 0, 0, 184, 42, 1, 0, 0, 0, 185, 186, 5, 33, 0, 0, 186, 187, 5,
		126, 0, 0, 187, 44, 1, 0, 0, 0, 188, 189, 7, 13, 0, 0, 189, 190, 7, 14,
		0, 0, 190, 46, 1, 0, 0, 0, 191, 192, 7, 15, 0, 0, 192, 193, 7, 16, 0, 0,
		193, 194, 7, 17, 0, 0, 194, 195, 7, 18, 0, 0, 195, 196, 7, 8, 0, 0, 196,
		197, 7, 10, 0, 0, 197, 198, 7, 17, 0, 0, 198, 48, 1, 0, 0, 0, 199, 200,
		7, 8, 0, 0, 200, 201, 7, 6, 0, 0, 201, 50, 1, 0, 0, 0, 202, 203, 7, 16,
		0, 0, 203, 204, 7, 19, 0, 0, 204, 205, 7, 6, 0, 0, 205, 206, 7, 8, 0, 0,
		206, 207, 7, 9, 0, 0, 207, 208, 7, 16, 0, 0, 208, 209, 7, 6, 0, 0, 209,
		210, 7, 19, 0, 0, 210, 52, 1, 0, 0, 0, 211, 212, 7, 19, 0, 0, 212, 213,
		7, 9, 0, 0, 213, 214, 7, 8, 0, 0, 214, 215, 7, 10, 0, 0, 215, 216, 7, 20,
		0, 0, 216, 217, 5, 95, 0, 0, 217, 218, 7, 11, 0, 0, 218, 219, 7, 1, 0,
		0, 219, 220, 7, 21, 0, 0, 220, 221, 7, 17, 0, 0, 221, 54, 1, 0, 0, 0, 222,
		223, 7, 19, 0, 0, 223, 224, 7, 9, 0, 0, 224, 225, 7, 8, 0, 0, 225, 226,
		7, 10, 0, 0, 226, 227, 7, 20, 0, 0, 227, 228, 5, 95, 0, 0, 228, 229, 7,
		9, 0, 0, 229, 230, 7, 16, 0, 0, 230, 231, 7, 19, 0, 0, 231, 232, 7, 18,
		0, 0, 232, 233, 7, 17, 0, 0, 233, 56, 1, 0, 0, 0, 234, 235, 7, 8, 0, 0,
		235, 236, 7, 21, 0, 0, 236, 237, 7, 21, 0, 0, 237, 238, 7, 12, 0, 0, 238,
		239, 7, 1, 0, 0, 239, 240, 7, 17, 0, 0, 240, 58, 1, 0, 0, 0, 241, 242,
		7, 13, 0, 0, 242, 243, 7, 8, 0, 0, 243, 244, 7, 8, 0, 0, 244, 245, 7, 11,
		0, 0, 245, 60, 1, 0, 0, 0, 246, 247, 5, 123, 0, 0, 247, 62, 1, 0, 0, 0,
		248, 249, 5, 125, 0, 0, 249, 64, 1, 0, 0, 0, 250, 251, 5, 40, 0, 0, 251,
		66, 1, 0, 0, 0, 252, 253, 5, 41, 0, 0, 253, 68, 1, 0, 0, 0, 254, 255, 5,
		91, 0, 0, 255, 70, 1, 0, 0, 0, 256, 257, 5, 93, 0, 0, 257, 72, 1, 0, 0,
		0, 258, 259, 5, 44, 0, 0, 259, 74, 1, 0, 0, 0, 260, 261, 5, 64, 0, 0, 261,
		76, 1, 0, 0, 0, 262, 263, 3, 69, 34, 0, 263, 264, 3, 81, 40, 0, 264, 266,
		5, 58, 0, 0, 265, 267, 3, 81, 40, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1,
		0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 3, 71, 35, 0, 269, 78, 1, 0, 0,
		0, 270, 271, 3, 69, 34, 0, 271, 272, 3, 81, 40, 0, 272, 273, 3, 71, 35,
		0, 273, 80, 1, 0, 0, 0, 274, 276, 7, 0, 0, 0, 275, 274, 1, 0, 0, 0, 276,
		277, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 282,
		1, 0, 0, 0, 279, 280, 7, 22, 0, 0, 280, 283, 7, 12, 0, 0, 281, 283, 7,
		23, 0, 0, 282, 279, 1, 0, 0, 0, 282, 281, 1, 0, 0, 0, 283, 285, 1, 0, 0,
		0, 284, 275, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286,
		287, 1, 0, 0, 0, 287, 82, 1, 0, 0, 0, 288, 292, 7, 24, 0, 0, 289, 291,
		7, 25, 0, 0, 290, 289, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290, 1, 0,
		0, 0, 292, 293, 1, 0, 0, 0, 293, 295, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0,
		295, 296, 6, 41, 0, 0, 296, 84, 1, 0, 0, 0, 297, 301, 7, 26, 0, 0, 298,
		300, 7, 27, 0, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299,
		1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 86, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 304, 306, 7, 28, 0, 0, 305, 304, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0,
		307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309,
		310, 6, 43, 1, 0, 310, 88, 1, 0, 0, 0, 311, 315, 5, 35, 0, 0, 312, 314,
		9, 0, 0, 0, 313, 312, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 316, 1, 0,
		0, 0, 315, 313, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0,
		318, 319, 5, 10, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 6, 44, 2, 0, 321,
		90, 1, 0, 0, 0, 20, 0, 94, 100, 102, 107, 110, 114, 120, 122, 130, 132,
		136, 266, 277, 282, 286, 292, 301, 307, 315, 3, 1, 41, 0, 0, 2, 0, 0, 3,
		0,
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

// PromQLLexerInit initializes any static state used to implement PromQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPromQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLLexerInit() {
	staticData := &PromQLLexerLexerStaticData
	staticData.once.Do(promqllexerLexerInit)
}

// NewPromQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPromQLLexer(input antlr.CharStream) *PromQLLexer {
	PromQLLexerInit()
	l := new(PromQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PromQLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PromQLLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PromQLLexer tokens.
const (
	PromQLLexerAGGREGATION_OPERATOR = 1
	PromQLLexerFUNCTION             = 2
	PromQLLexerNUMBER               = 3
	PromQLLexerSTRING               = 4
	PromQLLexerADD                  = 5
	PromQLLexerSUB                  = 6
	PromQLLexerMULT                 = 7
	PromQLLexerDIV                  = 8
	PromQLLexerMOD                  = 9
	PromQLLexerPOW                  = 10
	PromQLLexerAND                  = 11
	PromQLLexerOR                   = 12
	PromQLLexerUNLESS               = 13
	PromQLLexerEQ                   = 14
	PromQLLexerDEQ                  = 15
	PromQLLexerNE                   = 16
	PromQLLexerGT                   = 17
	PromQLLexerLT                   = 18
	PromQLLexerGE                   = 19
	PromQLLexerLE                   = 20
	PromQLLexerRE                   = 21
	PromQLLexerNRE                  = 22
	PromQLLexerBY                   = 23
	PromQLLexerWITHOUT              = 24
	PromQLLexerON                   = 25
	PromQLLexerIGNORING             = 26
	PromQLLexerGROUP_LEFT           = 27
	PromQLLexerGROUP_RIGHT          = 28
	PromQLLexerOFFSET               = 29
	PromQLLexerBOOL                 = 30
	PromQLLexerLEFT_BRACE           = 31
	PromQLLexerRIGHT_BRACE          = 32
	PromQLLexerLEFT_PAREN           = 33
	PromQLLexerRIGHT_PAREN          = 34
	PromQLLexerLEFT_BRACKET         = 35
	PromQLLexerRIGHT_BRACKET        = 36
	PromQLLexerCOMMA                = 37
	PromQLLexerAT                   = 38
	PromQLLexerSUBQUERY_RANGE       = 39
	PromQLLexerTIME_RANGE           = 40
	PromQLLexerDURATION             = 41
	PromQLLexerMETRIC_NAME          = 42
	PromQLLexerLABEL_NAME           = 43
	PromQLLexerWS                   = 44
	PromQLLexerSL_COMMENT           = 45
)

// PromQLLexer escapedChannels.
const (
	PromQLLexerWHITESPACE = 2
	PromQLLexerCOMMENTS   = 3
)

var Functions []string
var AggregationOperators []string

func (l *PromQLLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 41:
		l.METRIC_NAME_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PromQLLexer) METRIC_NAME_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:

		if slices.Contains(AggregationOperators, l.GetText()) {
			l.SetType(PromQLLexerAGGREGATION_OPERATOR)
		}
		if slices.Contains(Functions, l.GetText()) {
			l.SetType(PromQLLexerAGGREGATION_OPERATOR)
		}

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

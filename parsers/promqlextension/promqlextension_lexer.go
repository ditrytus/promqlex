// Code generated from PromQLExtensionLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlextension

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

type PromQLExtensionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PromQLExtensionLexerLexerStaticData struct {
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

func promqlextensionlexerLexerInit() {
	staticData := &PromQLExtensionLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN", "WHITESPACE", "COMMENTS",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "RAW_STRING_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "'if'", "'true'", "'false'", "'T'", "':'", "'.'", "",
		"", "", "'metric'", "'label'", "'def'", "'$'", "", "", "", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'", "'='", "'=='",
		"'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'", "'by'", "'without'",
		"'on'", "'ignoring'", "'group_left'", "'group_right'", "'offset'", "'bool'",
		"'{'", "'}'", "'('", "')'", "'['", "']'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "ID", "IF", "TRUE", "FALSE",
		"T", "COLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS", "DIGITS", "METRIC_KEYWORD",
		"LABEL_KEYWORD", "DEF", "CALL_SIGN", "NL", "NUMBER", "STRING", "ADD",
		"SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR", "UNLESS", "EQ", "DEQ",
		"NE", "GT", "LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING",
		"GROUP_LEFT", "GROUP_RIGHT", "OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE",
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA",
		"AT", "DURATION", "METRIC_NAME", "LABEL_NAME", "WS", "SL_COMMENT", "RAW_STRING",
		"BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"SUBQUERY_RANGE", "TIME_RANGE", "ID", "IF", "TRUE", "FALSE", "T", "COLON",
		"DOT", "POSITIVE_INTEGER", "TWO_DIGITS", "DIGITS", "METRIC_KEYWORD",
		"LABEL_KEYWORD", "DEF", "CALL_SIGN", "NL", "NUMERAL", "SCIENTIFIC_NUMBER",
		"NUMBER", "STRING", "BACKTICK_OPEN", "ADD", "SUB", "MULT", "DIV", "MOD",
		"POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE",
		"RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT",
		"OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN",
		"LEFT_BRACKET", "RIGHT_BRACKET", "COMMA", "AT", "DURATION", "METRIC_NAME",
		"LABEL_NAME", "WS", "SL_COMMENT", "RAW_STRING_CONTENT", "RAW_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 60, 441, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
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
		2, 62, 7, 62, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 4, 2, 135, 8, 2, 11,
		2, 12, 2, 136, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 4, 9, 161, 8, 9, 11, 9, 12, 9, 162, 1, 9, 5, 9, 166, 8, 9, 10, 9, 12,
		9, 169, 9, 9, 3, 9, 171, 8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 177,
		8, 11, 11, 11, 12, 11, 178, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 203, 8, 16, 1, 17, 4, 17, 206,
		8, 17, 11, 17, 12, 17, 207, 1, 17, 1, 17, 4, 17, 212, 8, 17, 11, 17, 12,
		17, 213, 3, 17, 216, 8, 17, 1, 18, 1, 18, 1, 18, 3, 18, 221, 8, 18, 1,
		18, 3, 18, 224, 8, 18, 1, 19, 1, 19, 3, 19, 228, 8, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 5, 20, 234, 8, 20, 10, 20, 12, 20, 237, 9, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 244, 8, 20, 10, 20, 12, 20, 247, 9, 20, 1,
		20, 3, 20, 250, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1,
		52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 4, 56, 382, 8, 56,
		11, 56, 12, 56, 383, 1, 56, 1, 56, 1, 56, 3, 56, 389, 8, 56, 4, 56, 391,
		8, 56, 11, 56, 12, 56, 392, 1, 57, 1, 57, 5, 57, 397, 8, 57, 10, 57, 12,
		57, 400, 9, 57, 1, 57, 1, 57, 1, 58, 1, 58, 5, 58, 406, 8, 58, 10, 58,
		12, 58, 409, 9, 58, 1, 59, 4, 59, 412, 8, 59, 11, 59, 12, 59, 413, 1, 59,
		1, 59, 1, 60, 1, 60, 5, 60, 420, 8, 60, 10, 60, 12, 60, 423, 9, 60, 1,
		60, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 4, 61, 432, 8, 61, 11, 61,
		12, 61, 433, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 421, 0, 63, 2,
		0, 4, 0, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11,
		24, 12, 26, 13, 28, 14, 30, 15, 32, 16, 34, 17, 36, 0, 38, 0, 40, 18, 42,
		19, 44, 60, 46, 20, 48, 21, 50, 22, 52, 23, 54, 24, 56, 25, 58, 26, 60,
		27, 62, 28, 64, 29, 66, 30, 68, 31, 70, 32, 72, 33, 74, 34, 76, 35, 78,
		36, 80, 37, 82, 38, 84, 39, 86, 40, 88, 41, 90, 42, 92, 43, 94, 44, 96,
		45, 98, 46, 100, 47, 102, 48, 104, 49, 106, 50, 108, 51, 110, 52, 112,
		53, 114, 54, 116, 55, 118, 56, 120, 57, 122, 58, 124, 0, 126, 59, 2, 0,
		1, 33, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2,
		0, 73, 73, 105, 105, 2, 0, 70, 70, 102, 102, 2, 0, 84, 84, 116, 116, 2,
		0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117, 2, 0, 69, 69, 101, 101, 2,
		0, 65, 65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 83, 83, 115, 115, 1, 0,
		48, 57, 1, 0, 49, 57, 2, 0, 77, 77, 109, 109, 2, 0, 67, 67, 99, 99, 2,
		0, 66, 66, 98, 98, 2, 0, 68, 68, 100, 100, 2, 0, 43, 43, 45, 45, 2, 0,
		39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 2, 0, 78, 78, 110, 110, 2, 0, 79,
		79, 111, 111, 2, 0, 89, 89, 121, 121, 2, 0, 87, 87, 119, 119, 2, 0, 72,
		72, 104, 104, 2, 0, 71, 71, 103, 103, 2, 0, 80, 80, 112, 112, 12, 0, 68,
		68, 72, 72, 77, 77, 83, 83, 87, 87, 89, 89, 100, 100, 104, 104, 109, 109,
		115, 115, 119, 119, 121, 121, 4, 0, 58, 58, 65, 90, 95, 95, 97, 122, 4,
		0, 48, 58, 65, 90, 95, 95, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0,
		9, 10, 13, 13, 32, 32, 2, 0, 92, 92, 96, 96, 461, 0, 6, 1, 0, 0, 0, 0,
		8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0,
		0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0,
		0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0, 28, 1, 0, 0, 0, 0, 30, 1, 0,
		0, 0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0, 0, 40, 1, 0, 0, 0, 0, 42, 1,
		0, 0, 0, 0, 44, 1, 0, 0, 0, 0, 46, 1, 0, 0, 0, 0, 48, 1, 0, 0, 0, 0, 50,
		1, 0, 0, 0, 0, 52, 1, 0, 0, 0, 0, 54, 1, 0, 0, 0, 0, 56, 1, 0, 0, 0, 0,
		58, 1, 0, 0, 0, 0, 60, 1, 0, 0, 0, 0, 62, 1, 0, 0, 0, 0, 64, 1, 0, 0, 0,
		0, 66, 1, 0, 0, 0, 0, 68, 1, 0, 0, 0, 0, 70, 1, 0, 0, 0, 0, 72, 1, 0, 0,
		0, 0, 74, 1, 0, 0, 0, 0, 76, 1, 0, 0, 0, 0, 78, 1, 0, 0, 0, 0, 80, 1, 0,
		0, 0, 0, 82, 1, 0, 0, 0, 0, 84, 1, 0, 0, 0, 0, 86, 1, 0, 0, 0, 0, 88, 1,
		0, 0, 0, 0, 90, 1, 0, 0, 0, 0, 92, 1, 0, 0, 0, 0, 94, 1, 0, 0, 0, 0, 96,
		1, 0, 0, 0, 0, 98, 1, 0, 0, 0, 0, 100, 1, 0, 0, 0, 0, 102, 1, 0, 0, 0,
		0, 104, 1, 0, 0, 0, 0, 106, 1, 0, 0, 0, 0, 108, 1, 0, 0, 0, 0, 110, 1,
		0, 0, 0, 0, 112, 1, 0, 0, 0, 0, 114, 1, 0, 0, 0, 0, 116, 1, 0, 0, 0, 0,
		118, 1, 0, 0, 0, 0, 120, 1, 0, 0, 0, 0, 122, 1, 0, 0, 0, 1, 124, 1, 0,
		0, 0, 1, 126, 1, 0, 0, 0, 2, 128, 1, 0, 0, 0, 4, 130, 1, 0, 0, 0, 6, 132,
		1, 0, 0, 0, 8, 138, 1, 0, 0, 0, 10, 141, 1, 0, 0, 0, 12, 146, 1, 0, 0,
		0, 14, 152, 1, 0, 0, 0, 16, 154, 1, 0, 0, 0, 18, 156, 1, 0, 0, 0, 20, 170,
		1, 0, 0, 0, 22, 172, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 180, 1, 0, 0,
		0, 28, 187, 1, 0, 0, 0, 30, 193, 1, 0, 0, 0, 32, 197, 1, 0, 0, 0, 34, 202,
		1, 0, 0, 0, 36, 205, 1, 0, 0, 0, 38, 217, 1, 0, 0, 0, 40, 227, 1, 0, 0,
		0, 42, 249, 1, 0, 0, 0, 44, 251, 1, 0, 0, 0, 46, 256, 1, 0, 0, 0, 48, 258,
		1, 0, 0, 0, 50, 260, 1, 0, 0, 0, 52, 262, 1, 0, 0, 0, 54, 264, 1, 0, 0,
		0, 56, 266, 1, 0, 0, 0, 58, 268, 1, 0, 0, 0, 60, 272, 1, 0, 0, 0, 62, 275,
		1, 0, 0, 0, 64, 282, 1, 0, 0, 0, 66, 284, 1, 0, 0, 0, 68, 287, 1, 0, 0,
		0, 70, 290, 1, 0, 0, 0, 72, 292, 1, 0, 0, 0, 74, 294, 1, 0, 0, 0, 76, 297,
		1, 0, 0, 0, 78, 300, 1, 0, 0, 0, 80, 303, 1, 0, 0, 0, 82, 306, 1, 0, 0,
		0, 84, 309, 1, 0, 0, 0, 86, 317, 1, 0, 0, 0, 88, 320, 1, 0, 0, 0, 90, 329,
		1, 0, 0, 0, 92, 340, 1, 0, 0, 0, 94, 352, 1, 0, 0, 0, 96, 359, 1, 0, 0,
		0, 98, 364, 1, 0, 0, 0, 100, 366, 1, 0, 0, 0, 102, 368, 1, 0, 0, 0, 104,
		370, 1, 0, 0, 0, 106, 372, 1, 0, 0, 0, 108, 374, 1, 0, 0, 0, 110, 376,
		1, 0, 0, 0, 112, 378, 1, 0, 0, 0, 114, 390, 1, 0, 0, 0, 116, 394, 1, 0,
		0, 0, 118, 403, 1, 0, 0, 0, 120, 411, 1, 0, 0, 0, 122, 417, 1, 0, 0, 0,
		124, 431, 1, 0, 0, 0, 126, 437, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		3, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 5, 1, 0, 0, 0, 132, 134, 7, 0,
		0, 0, 133, 135, 7, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 7, 1, 0, 0, 0, 138, 139,
		7, 2, 0, 0, 139, 140, 7, 3, 0, 0, 140, 9, 1, 0, 0, 0, 141, 142, 7, 4, 0,
		0, 142, 143, 7, 5, 0, 0, 143, 144, 7, 6, 0, 0, 144, 145, 7, 7, 0, 0, 145,
		11, 1, 0, 0, 0, 146, 147, 7, 3, 0, 0, 147, 148, 7, 8, 0, 0, 148, 149, 7,
		9, 0, 0, 149, 150, 7, 10, 0, 0, 150, 151, 7, 7, 0, 0, 151, 13, 1, 0, 0,
		0, 152, 153, 7, 4, 0, 0, 153, 15, 1, 0, 0, 0, 154, 155, 5, 58, 0, 0, 155,
		17, 1, 0, 0, 0, 156, 157, 5, 46, 0, 0, 157, 19, 1, 0, 0, 0, 158, 171, 7,
		11, 0, 0, 159, 161, 7, 12, 0, 0, 160, 159, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 167, 1, 0, 0, 0,
		164, 166, 7, 11, 0, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167,
		1, 0, 0, 0, 170, 158, 1, 0, 0, 0, 170, 160, 1, 0, 0, 0, 171, 21, 1, 0,
		0, 0, 172, 173, 7, 11, 0, 0, 173, 174, 7, 11, 0, 0, 174, 23, 1, 0, 0, 0,
		175, 177, 7, 11, 0, 0, 176, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 25, 1, 0, 0, 0, 180, 181, 7,
		13, 0, 0, 181, 182, 7, 7, 0, 0, 182, 183, 7, 4, 0, 0, 183, 184, 7, 5, 0,
		0, 184, 185, 7, 2, 0, 0, 185, 186, 7, 14, 0, 0, 186, 27, 1, 0, 0, 0, 187,
		188, 7, 9, 0, 0, 188, 189, 7, 8, 0, 0, 189, 190, 7, 15, 0, 0, 190, 191,
		7, 7, 0, 0, 191, 192, 7, 9, 0, 0, 192, 29, 1, 0, 0, 0, 193, 194, 7, 16,
		0, 0, 194, 195, 7, 7, 0, 0, 195, 196, 7, 3, 0, 0, 196, 31, 1, 0, 0, 0,
		197, 198, 5, 36, 0, 0, 198, 33, 1, 0, 0, 0, 199, 203, 5, 10, 0, 0, 200,
		201, 5, 13, 0, 0, 201, 203, 5, 10, 0, 0, 202, 199, 1, 0, 0, 0, 202, 200,
		1, 0, 0, 0, 203, 35, 1, 0, 0, 0, 204, 206, 7, 11, 0, 0, 205, 204, 1, 0,
		0, 0, 206, 207, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 215, 1, 0, 0, 0, 209, 211, 5, 46, 0, 0, 210, 212, 7, 11, 0, 0, 211,
		210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214,
		1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215, 209, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 37, 1, 0, 0, 0, 217, 223, 3, 36, 17, 0, 218, 220, 7, 7, 0, 0,
		219, 221, 7, 17, 0, 0, 220, 219, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 224, 3, 36, 17, 0, 223, 218, 1, 0, 0, 0, 223, 224,
		1, 0, 0, 0, 224, 39, 1, 0, 0, 0, 225, 228, 3, 36, 17, 0, 226, 228, 3, 38,
		18, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0, 228, 41, 1, 0, 0, 0,
		229, 235, 5, 39, 0, 0, 230, 234, 8, 18, 0, 0, 231, 232, 5, 92, 0, 0, 232,
		234, 9, 0, 0, 0, 233, 230, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 237,
		1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 238, 1, 0,
		0, 0, 237, 235, 1, 0, 0, 0, 238, 250, 5, 39, 0, 0, 239, 245, 5, 34, 0,
		0, 240, 244, 8, 19, 0, 0, 241, 242, 5, 92, 0, 0, 242, 244, 9, 0, 0, 0,
		243, 240, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245,
		243, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248, 1, 0, 0, 0, 247, 245,
		1, 0, 0, 0, 248, 250, 5, 34, 0, 0, 249, 229, 1, 0, 0, 0, 249, 239, 1, 0,
		0, 0, 250, 43, 1, 0, 0, 0, 251, 252, 5, 96, 0, 0, 252, 253, 1, 0, 0, 0,
		253, 254, 6, 21, 0, 0, 254, 255, 6, 21, 1, 0, 255, 45, 1, 0, 0, 0, 256,
		257, 5, 43, 0, 0, 257, 47, 1, 0, 0, 0, 258, 259, 5, 45, 0, 0, 259, 49,
		1, 0, 0, 0, 260, 261, 5, 42, 0, 0, 261, 51, 1, 0, 0, 0, 262, 263, 5, 47,
		0, 0, 263, 53, 1, 0, 0, 0, 264, 265, 5, 37, 0, 0, 265, 55, 1, 0, 0, 0,
		266, 267, 5, 94, 0, 0, 267, 57, 1, 0, 0, 0, 268, 269, 7, 8, 0, 0, 269,
		270, 7, 20, 0, 0, 270, 271, 7, 16, 0, 0, 271, 59, 1, 0, 0, 0, 272, 273,
		7, 21, 0, 0, 273, 274, 7, 5, 0, 0, 274, 61, 1, 0, 0, 0, 275, 276, 7, 6,
		0, 0, 276, 277, 7, 20, 0, 0, 277, 278, 7, 9, 0, 0, 278, 279, 7, 7, 0, 0,
		279, 280, 7, 10, 0, 0, 280, 281, 7, 10, 0, 0, 281, 63, 1, 0, 0, 0, 282,
		283, 5, 61, 0, 0, 283, 65, 1, 0, 0, 0, 284, 285, 5, 61, 0, 0, 285, 286,
		5, 61, 0, 0, 286, 67, 1, 0, 0, 0, 287, 288, 5, 33, 0, 0, 288, 289, 5, 61,
		0, 0, 289, 69, 1, 0, 0, 0, 290, 291, 5, 62, 0, 0, 291, 71, 1, 0, 0, 0,
		292, 293, 5, 60, 0, 0, 293, 73, 1, 0, 0, 0, 294, 295, 5, 62, 0, 0, 295,
		296, 5, 61, 0, 0, 296, 75, 1, 0, 0, 0, 297, 298, 5, 60, 0, 0, 298, 299,
		5, 61, 0, 0, 299, 77, 1, 0, 0, 0, 300, 301, 5, 61, 0, 0, 301, 302, 5, 126,
		0, 0, 302, 79, 1, 0, 0, 0, 303, 304, 5, 33, 0, 0, 304, 305, 5, 126, 0,
		0, 305, 81, 1, 0, 0, 0, 306, 307, 7, 15, 0, 0, 307, 308, 7, 22, 0, 0, 308,
		83, 1, 0, 0, 0, 309, 310, 7, 23, 0, 0, 310, 311, 7, 2, 0, 0, 311, 312,
		7, 4, 0, 0, 312, 313, 7, 24, 0, 0, 313, 314, 7, 21, 0, 0, 314, 315, 7,
		6, 0, 0, 315, 316, 7, 4, 0, 0, 316, 85, 1, 0, 0, 0, 317, 318, 7, 21, 0,
		0, 318, 319, 7, 20, 0, 0, 319, 87, 1, 0, 0, 0, 320, 321, 7, 2, 0, 0, 321,
		322, 7, 25, 0, 0, 322, 323, 7, 20, 0, 0, 323, 324, 7, 21, 0, 0, 324, 325,
		7, 5, 0, 0, 325, 326, 7, 2, 0, 0, 326, 327, 7, 20, 0, 0, 327, 328, 7, 25,
		0, 0, 328, 89, 1, 0, 0, 0, 329, 330, 7, 25, 0, 0, 330, 331, 7, 5, 0, 0,
		331, 332, 7, 21, 0, 0, 332, 333, 7, 6, 0, 0, 333, 334, 7, 26, 0, 0, 334,
		335, 5, 95, 0, 0, 335, 336, 7, 9, 0, 0, 336, 337, 7, 7, 0, 0, 337, 338,
		7, 3, 0, 0, 338, 339, 7, 4, 0, 0, 339, 91, 1, 0, 0, 0, 340, 341, 7, 25,
		0, 0, 341, 342, 7, 5, 0, 0, 342, 343, 7, 21, 0, 0, 343, 344, 7, 6, 0, 0,
		344, 345, 7, 26, 0, 0, 345, 346, 5, 95, 0, 0, 346, 347, 7, 5, 0, 0, 347,
		348, 7, 2, 0, 0, 348, 349, 7, 25, 0, 0, 349, 350, 7, 24, 0, 0, 350, 351,
		7, 4, 0, 0, 351, 93, 1, 0, 0, 0, 352, 353, 7, 21, 0, 0, 353, 354, 7, 3,
		0, 0, 354, 355, 7, 3, 0, 0, 355, 356, 7, 10, 0, 0, 356, 357, 7, 7, 0, 0,
		357, 358, 7, 4, 0, 0, 358, 95, 1, 0, 0, 0, 359, 360, 7, 15, 0, 0, 360,
		361, 7, 21, 0, 0, 361, 362, 7, 21, 0, 0, 362, 363, 7, 9, 0, 0, 363, 97,
		1, 0, 0, 0, 364, 365, 5, 123, 0, 0, 365, 99, 1, 0, 0, 0, 366, 367, 5, 125,
		0, 0, 367, 101, 1, 0, 0, 0, 368, 369, 5, 40, 0, 0, 369, 103, 1, 0, 0, 0,
		370, 371, 5, 41, 0, 0, 371, 105, 1, 0, 0, 0, 372, 373, 5, 91, 0, 0, 373,
		107, 1, 0, 0, 0, 374, 375, 5, 93, 0, 0, 375, 109, 1, 0, 0, 0, 376, 377,
		5, 44, 0, 0, 377, 111, 1, 0, 0, 0, 378, 379, 5, 64, 0, 0, 379, 113, 1,
		0, 0, 0, 380, 382, 7, 11, 0, 0, 381, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0,
		0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 388, 1, 0, 0, 0, 385,
		386, 7, 13, 0, 0, 386, 389, 7, 10, 0, 0, 387, 389, 7, 27, 0, 0, 388, 385,
		1, 0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 391, 1, 0, 0, 0, 390, 381, 1, 0,
		0, 0, 391, 392, 1, 0, 0, 0, 392, 390, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0,
		393, 115, 1, 0, 0, 0, 394, 398, 7, 28, 0, 0, 395, 397, 7, 29, 0, 0, 396,
		395, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399,
		1, 0, 0, 0, 399, 401, 1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 401, 402, 6, 57,
		2, 0, 402, 117, 1, 0, 0, 0, 403, 407, 7, 30, 0, 0, 404, 406, 7, 1, 0, 0,
		405, 404, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407,
		408, 1, 0, 0, 0, 408, 119, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 412,
		7, 31, 0, 0, 411, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 411, 1, 0,
		0, 0, 413, 414, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416, 6, 59, 3, 0,
		416, 121, 1, 0, 0, 0, 417, 421, 5, 35, 0, 0, 418, 420, 9, 0, 0, 0, 419,
		418, 1, 0, 0, 0, 420, 423, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 421, 419,
		1, 0, 0, 0, 422, 424, 1, 0, 0, 0, 423, 421, 1, 0, 0, 0, 424, 425, 5, 10,
		0, 0, 425, 426, 1, 0, 0, 0, 426, 427, 6, 60, 4, 0, 427, 123, 1, 0, 0, 0,
		428, 432, 8, 32, 0, 0, 429, 430, 5, 92, 0, 0, 430, 432, 5, 96, 0, 0, 431,
		428, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 431,
		1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436, 6, 61,
		0, 0, 436, 125, 1, 0, 0, 0, 437, 438, 5, 96, 0, 0, 438, 439, 1, 0, 0, 0,
		439, 440, 6, 62, 5, 0, 440, 127, 1, 0, 0, 0, 28, 0, 1, 136, 162, 167, 170,
		178, 202, 207, 213, 215, 220, 223, 227, 233, 235, 243, 245, 249, 383, 388,
		392, 398, 407, 413, 421, 431, 433, 6, 3, 0, 0, 5, 1, 0, 1, 57, 0, 0, 2,
		0, 0, 3, 0, 4, 0, 0,
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

// PromQLExtensionLexerInit initializes any static state used to implement PromQLExtensionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPromQLExtensionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLExtensionLexerInit() {
	staticData := &PromQLExtensionLexerLexerStaticData
	staticData.once.Do(promqlextensionlexerLexerInit)
}

// NewPromQLExtensionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPromQLExtensionLexer(input antlr.CharStream) *PromQLExtensionLexer {
	PromQLExtensionLexerInit()
	l := new(PromQLExtensionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PromQLExtensionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PromQLExtensionLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PromQLExtensionLexer tokens.
const (
	PromQLExtensionLexerAGGREGATION_OPERATOR = 1
	PromQLExtensionLexerFUNCTION             = 2
	PromQLExtensionLexerID                   = 3
	PromQLExtensionLexerIF                   = 4
	PromQLExtensionLexerTRUE                 = 5
	PromQLExtensionLexerFALSE                = 6
	PromQLExtensionLexerT                    = 7
	PromQLExtensionLexerCOLON                = 8
	PromQLExtensionLexerDOT                  = 9
	PromQLExtensionLexerPOSITIVE_INTEGER     = 10
	PromQLExtensionLexerTWO_DIGITS           = 11
	PromQLExtensionLexerDIGITS               = 12
	PromQLExtensionLexerMETRIC_KEYWORD       = 13
	PromQLExtensionLexerLABEL_KEYWORD        = 14
	PromQLExtensionLexerDEF                  = 15
	PromQLExtensionLexerCALL_SIGN            = 16
	PromQLExtensionLexerNL                   = 17
	PromQLExtensionLexerNUMBER               = 18
	PromQLExtensionLexerSTRING               = 19
	PromQLExtensionLexerADD                  = 20
	PromQLExtensionLexerSUB                  = 21
	PromQLExtensionLexerMULT                 = 22
	PromQLExtensionLexerDIV                  = 23
	PromQLExtensionLexerMOD                  = 24
	PromQLExtensionLexerPOW                  = 25
	PromQLExtensionLexerAND                  = 26
	PromQLExtensionLexerOR                   = 27
	PromQLExtensionLexerUNLESS               = 28
	PromQLExtensionLexerEQ                   = 29
	PromQLExtensionLexerDEQ                  = 30
	PromQLExtensionLexerNE                   = 31
	PromQLExtensionLexerGT                   = 32
	PromQLExtensionLexerLT                   = 33
	PromQLExtensionLexerGE                   = 34
	PromQLExtensionLexerLE                   = 35
	PromQLExtensionLexerRE                   = 36
	PromQLExtensionLexerNRE                  = 37
	PromQLExtensionLexerBY                   = 38
	PromQLExtensionLexerWITHOUT              = 39
	PromQLExtensionLexerON                   = 40
	PromQLExtensionLexerIGNORING             = 41
	PromQLExtensionLexerGROUP_LEFT           = 42
	PromQLExtensionLexerGROUP_RIGHT          = 43
	PromQLExtensionLexerOFFSET               = 44
	PromQLExtensionLexerBOOL                 = 45
	PromQLExtensionLexerLEFT_BRACE           = 46
	PromQLExtensionLexerRIGHT_BRACE          = 47
	PromQLExtensionLexerLEFT_PAREN           = 48
	PromQLExtensionLexerRIGHT_PAREN          = 49
	PromQLExtensionLexerLEFT_BRACKET         = 50
	PromQLExtensionLexerRIGHT_BRACKET        = 51
	PromQLExtensionLexerCOMMA                = 52
	PromQLExtensionLexerAT                   = 53
	PromQLExtensionLexerDURATION             = 54
	PromQLExtensionLexerMETRIC_NAME          = 55
	PromQLExtensionLexerLABEL_NAME           = 56
	PromQLExtensionLexerWS                   = 57
	PromQLExtensionLexerSL_COMMENT           = 58
	PromQLExtensionLexerRAW_STRING           = 59
	PromQLExtensionLexerBACKTICK_OPEN        = 60
)

// PromQLExtensionLexer escapedChannels.
const (
	PromQLExtensionLexerWHITESPACE = 2
	PromQLExtensionLexerCOMMENTS   = 3
)

// PromQLExtensionLexerRAW_STRING_MODE is the PromQLExtensionLexer mode.
const PromQLExtensionLexerRAW_STRING_MODE = 1

type FunctionsProvider interface {
	GetTokenType(text string) (int, bool)
}

func (l *PromQLExtensionLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 57:
		l.METRIC_NAME_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PromQLExtensionLexer) METRIC_NAME_Action(localctx antlr.RuleContext, actionIndex int) {
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

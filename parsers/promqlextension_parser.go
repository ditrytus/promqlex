// Code generated from PromQLExtensionParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLExtensionParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PromQLExtensionParser struct {
	*antlr.BaseParser
}

var PromQLExtensionParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promqlextensionparserParserInit() {
	staticData := &PromQLExtensionParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'true'", "'false'", "'T'", "':'", "'.'", "", "", "", "", "",
		"'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "", "", "'{'", "'}'", "'('", "')'", "'['", "']'",
		"','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "EX_ID", "EX_TRUE", "EX_FALSE", "EX_T", "EX_COLON", "EX_DOT", "EX_POSITIVE_INTEGER",
		"EX_TWO_DIGITS", "EX_DIGITS", "NUMBER", "STRING", "ADD", "SUB", "MULT",
		"DIV", "MOD", "POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT",
		"LT", "GE", "LE", "RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT",
		"GROUP_RIGHT", "OFFSET", "BOOL", "AGGREGATION_OPERATOR", "FUNCTION",
		"LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET",
		"RIGHT_BRACKET", "COMMA", "AT", "SUBQUERY_RANGE", "TIME_RANGE", "DURATION",
		"METRIC_NAME", "LABEL_NAME", "WS", "SL_COMMENT",
	}
	staticData.RuleNames = []string{
		"promqlx", "ex_statement_list", "ex_statement", "ex_alias_def", "ex_macro_def",
		"ex_block", "ex_arg_list", "ex_if_statement", "ex_condition", "ex_compareVectorOperation",
		"ex_trueConst", "ex_falseConst", "ex_time_instant", "ex_iso_date_time",
		"ex_iso_date_time_ymdhmsf", "ex_iso_date_time_ymdhms", "ex_iso_date_time_ymdhm",
		"ex_iso_date_time_ymdh", "ex_iso_date_time_ymd", "ex_iso_date_time_ym",
		"ex_iso_date_time_y", "ex_year", "ex_month", "ex_day", "ex_hour", "ex_minutes",
		"ex_seconds", "ex_frac_sec", "expression", "vectorOperation", "unaryOp",
		"powOp", "multOp", "addOp", "compareOp", "andUnlessOp", "orOp", "vectorMatchOp",
		"subqueryOp", "offsetOp", "vector", "parens", "instantSelector", "labelMatcher",
		"labelMatcherOperator", "labelMatcherList", "matrixSelector", "offset",
		"function_", "parameter", "parameterList", "aggregation", "by", "without",
		"grouping", "on_", "ignoring", "groupLeft", "groupRight", "labelName",
		"labelNameList", "keyword", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 503, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 1,
		0, 1, 0, 1, 1, 4, 1, 130, 8, 1, 11, 1, 12, 1, 131, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 138, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 147,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 4, 6,
		159, 8, 6, 11, 6, 12, 6, 160, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 3, 8,
		169, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 188, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 268, 8, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 303,
		8, 29, 10, 29, 12, 29, 306, 9, 29, 1, 30, 1, 30, 1, 31, 1, 31, 3, 31, 312,
		8, 31, 1, 32, 1, 32, 3, 32, 316, 8, 32, 1, 33, 1, 33, 3, 33, 320, 8, 33,
		1, 34, 1, 34, 3, 34, 324, 8, 34, 1, 34, 3, 34, 327, 8, 34, 1, 35, 1, 35,
		3, 35, 331, 8, 35, 1, 36, 1, 36, 3, 36, 335, 8, 36, 1, 37, 1, 37, 3, 37,
		339, 8, 37, 1, 38, 1, 38, 3, 38, 343, 8, 38, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 355, 8, 40, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 3, 42, 364, 8, 42, 1, 42, 3, 42,
		367, 8, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 373, 8, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 5, 45, 384, 8, 45, 10,
		45, 12, 45, 387, 9, 45, 1, 45, 3, 45, 390, 8, 45, 1, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 403, 8,
		47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 410, 8, 48, 10, 48, 12, 48,
		413, 9, 48, 3, 48, 415, 8, 48, 1, 48, 1, 48, 1, 49, 1, 49, 3, 49, 421,
		8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 427, 8, 50, 10, 50, 12, 50, 430,
		9, 50, 3, 50, 432, 8, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 3, 51, 441, 8, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51,
		449, 8, 51, 3, 51, 451, 8, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53,
		1, 54, 1, 54, 3, 54, 461, 8, 54, 1, 54, 1, 54, 3, 54, 465, 8, 54, 1, 55,
		1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 3, 57, 475, 8, 57, 1,
		58, 1, 58, 3, 58, 479, 8, 58, 1, 59, 1, 59, 1, 59, 3, 59, 484, 8, 59, 1,
		60, 1, 60, 1, 60, 1, 60, 5, 60, 490, 8, 60, 10, 60, 12, 60, 493, 9, 60,
		3, 60, 495, 8, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 0,
		1, 58, 63, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
		70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
		106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 0, 8, 1, 0, 12, 13, 1,
		0, 14, 16, 1, 0, 22, 27, 2, 0, 18, 18, 20, 20, 2, 0, 20, 20, 32, 32, 3,
		0, 21, 21, 23, 23, 28, 29, 2, 0, 18, 20, 30, 39, 1, 0, 10, 11, 502, 0,
		126, 1, 0, 0, 0, 2, 129, 1, 0, 0, 0, 4, 137, 1, 0, 0, 0, 6, 139, 1, 0,
		0, 0, 8, 143, 1, 0, 0, 0, 10, 151, 1, 0, 0, 0, 12, 155, 1, 0, 0, 0, 14,
		162, 1, 0, 0, 0, 16, 168, 1, 0, 0, 0, 18, 170, 1, 0, 0, 0, 20, 174, 1,
		0, 0, 0, 22, 176, 1, 0, 0, 0, 24, 178, 1, 0, 0, 0, 26, 187, 1, 0, 0, 0,
		28, 189, 1, 0, 0, 0, 30, 203, 1, 0, 0, 0, 32, 215, 1, 0, 0, 0, 34, 225,
		1, 0, 0, 0, 36, 233, 1, 0, 0, 0, 38, 239, 1, 0, 0, 0, 40, 243, 1, 0, 0,
		0, 42, 245, 1, 0, 0, 0, 44, 247, 1, 0, 0, 0, 46, 249, 1, 0, 0, 0, 48, 251,
		1, 0, 0, 0, 50, 253, 1, 0, 0, 0, 52, 255, 1, 0, 0, 0, 54, 257, 1, 0, 0,
		0, 56, 259, 1, 0, 0, 0, 58, 267, 1, 0, 0, 0, 60, 307, 1, 0, 0, 0, 62, 309,
		1, 0, 0, 0, 64, 313, 1, 0, 0, 0, 66, 317, 1, 0, 0, 0, 68, 321, 1, 0, 0,
		0, 70, 328, 1, 0, 0, 0, 72, 332, 1, 0, 0, 0, 74, 336, 1, 0, 0, 0, 76, 340,
		1, 0, 0, 0, 78, 344, 1, 0, 0, 0, 80, 354, 1, 0, 0, 0, 82, 356, 1, 0, 0,
		0, 84, 372, 1, 0, 0, 0, 86, 374, 1, 0, 0, 0, 88, 378, 1, 0, 0, 0, 90, 380,
		1, 0, 0, 0, 92, 391, 1, 0, 0, 0, 94, 402, 1, 0, 0, 0, 96, 404, 1, 0, 0,
		0, 98, 420, 1, 0, 0, 0, 100, 422, 1, 0, 0, 0, 102, 450, 1, 0, 0, 0, 104,
		452, 1, 0, 0, 0, 106, 455, 1, 0, 0, 0, 108, 460, 1, 0, 0, 0, 110, 466,
		1, 0, 0, 0, 112, 469, 1, 0, 0, 0, 114, 472, 1, 0, 0, 0, 116, 476, 1, 0,
		0, 0, 118, 483, 1, 0, 0, 0, 120, 485, 1, 0, 0, 0, 122, 498, 1, 0, 0, 0,
		124, 500, 1, 0, 0, 0, 126, 127, 3, 2, 1, 0, 127, 1, 1, 0, 0, 0, 128, 130,
		3, 4, 2, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 129, 1, 0,
		0, 0, 131, 132, 1, 0, 0, 0, 132, 3, 1, 0, 0, 0, 133, 138, 3, 6, 3, 0, 134,
		138, 3, 8, 4, 0, 135, 138, 3, 14, 7, 0, 136, 138, 3, 56, 28, 0, 137, 133,
		1, 0, 0, 0, 137, 134, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 136, 1, 0,
		0, 0, 138, 5, 1, 0, 0, 0, 139, 140, 5, 1, 0, 0, 140, 141, 5, 21, 0, 0,
		141, 142, 3, 56, 28, 0, 142, 7, 1, 0, 0, 0, 143, 144, 5, 1, 0, 0, 144,
		146, 5, 42, 0, 0, 145, 147, 3, 12, 6, 0, 146, 145, 1, 0, 0, 0, 146, 147,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 5, 43, 0, 0, 149, 150, 3, 10,
		5, 0, 150, 9, 1, 0, 0, 0, 151, 152, 5, 40, 0, 0, 152, 153, 3, 2, 1, 0,
		153, 154, 5, 41, 0, 0, 154, 11, 1, 0, 0, 0, 155, 158, 3, 56, 28, 0, 156,
		157, 5, 46, 0, 0, 157, 159, 3, 56, 28, 0, 158, 156, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 13, 1, 0,
		0, 0, 162, 163, 3, 16, 8, 0, 163, 164, 3, 10, 5, 0, 164, 15, 1, 0, 0, 0,
		165, 169, 3, 18, 9, 0, 166, 169, 3, 20, 10, 0, 167, 169, 3, 22, 11, 0,
		168, 165, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169,
		17, 1, 0, 0, 0, 170, 171, 3, 58, 29, 0, 171, 172, 3, 68, 34, 0, 172, 173,
		3, 58, 29, 0, 173, 19, 1, 0, 0, 0, 174, 175, 5, 2, 0, 0, 175, 21, 1, 0,
		0, 0, 176, 177, 5, 3, 0, 0, 177, 23, 1, 0, 0, 0, 178, 179, 3, 26, 13, 0,
		179, 25, 1, 0, 0, 0, 180, 188, 3, 28, 14, 0, 181, 188, 3, 30, 15, 0, 182,
		188, 3, 32, 16, 0, 183, 188, 3, 34, 17, 0, 184, 188, 3, 36, 18, 0, 185,
		188, 3, 38, 19, 0, 186, 188, 3, 40, 20, 0, 187, 180, 1, 0, 0, 0, 187, 181,
		1, 0, 0, 0, 187, 182, 1, 0, 0, 0, 187, 183, 1, 0, 0, 0, 187, 184, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 187, 186, 1, 0, 0, 0, 188, 27, 1, 0, 0, 0,
		189, 190, 3, 42, 21, 0, 190, 191, 5, 13, 0, 0, 191, 192, 3, 44, 22, 0,
		192, 193, 5, 13, 0, 0, 193, 194, 3, 46, 23, 0, 194, 195, 5, 4, 0, 0, 195,
		196, 3, 48, 24, 0, 196, 197, 5, 5, 0, 0, 197, 198, 3, 50, 25, 0, 198, 199,
		5, 5, 0, 0, 199, 200, 3, 52, 26, 0, 200, 201, 5, 6, 0, 0, 201, 202, 3,
		54, 27, 0, 202, 29, 1, 0, 0, 0, 203, 204, 3, 42, 21, 0, 204, 205, 5, 13,
		0, 0, 205, 206, 3, 44, 22, 0, 206, 207, 5, 13, 0, 0, 207, 208, 3, 46, 23,
		0, 208, 209, 5, 4, 0, 0, 209, 210, 3, 48, 24, 0, 210, 211, 5, 5, 0, 0,
		211, 212, 3, 50, 25, 0, 212, 213, 5, 5, 0, 0, 213, 214, 3, 52, 26, 0, 214,
		31, 1, 0, 0, 0, 215, 216, 3, 42, 21, 0, 216, 217, 5, 13, 0, 0, 217, 218,
		3, 44, 22, 0, 218, 219, 5, 13, 0, 0, 219, 220, 3, 46, 23, 0, 220, 221,
		5, 4, 0, 0, 221, 222, 3, 48, 24, 0, 222, 223, 5, 5, 0, 0, 223, 224, 3,
		50, 25, 0, 224, 33, 1, 0, 0, 0, 225, 226, 3, 42, 21, 0, 226, 227, 5, 13,
		0, 0, 227, 228, 3, 44, 22, 0, 228, 229, 5, 13, 0, 0, 229, 230, 3, 46, 23,
		0, 230, 231, 5, 4, 0, 0, 231, 232, 3, 48, 24, 0, 232, 35, 1, 0, 0, 0, 233,
		234, 3, 42, 21, 0, 234, 235, 5, 13, 0, 0, 235, 236, 3, 44, 22, 0, 236,
		237, 5, 13, 0, 0, 237, 238, 3, 46, 23, 0, 238, 37, 1, 0, 0, 0, 239, 240,
		3, 42, 21, 0, 240, 241, 5, 13, 0, 0, 241, 242, 3, 44, 22, 0, 242, 39, 1,
		0, 0, 0, 243, 244, 3, 42, 21, 0, 244, 41, 1, 0, 0, 0, 245, 246, 5, 7, 0,
		0, 246, 43, 1, 0, 0, 0, 247, 248, 5, 8, 0, 0, 248, 45, 1, 0, 0, 0, 249,
		250, 5, 8, 0, 0, 250, 47, 1, 0, 0, 0, 251, 252, 5, 8, 0, 0, 252, 49, 1,
		0, 0, 0, 253, 254, 5, 8, 0, 0, 254, 51, 1, 0, 0, 0, 255, 256, 5, 8, 0,
		0, 256, 53, 1, 0, 0, 0, 257, 258, 5, 9, 0, 0, 258, 55, 1, 0, 0, 0, 259,
		260, 3, 58, 29, 0, 260, 261, 5, 0, 0, 1, 261, 57, 1, 0, 0, 0, 262, 263,
		6, 29, -1, 0, 263, 264, 3, 60, 30, 0, 264, 265, 3, 58, 29, 9, 265, 268,
		1, 0, 0, 0, 266, 268, 3, 80, 40, 0, 267, 262, 1, 0, 0, 0, 267, 266, 1,
		0, 0, 0, 268, 304, 1, 0, 0, 0, 269, 270, 10, 11, 0, 0, 270, 271, 3, 62,
		31, 0, 271, 272, 3, 58, 29, 11, 272, 303, 1, 0, 0, 0, 273, 274, 10, 8,
		0, 0, 274, 275, 3, 64, 32, 0, 275, 276, 3, 58, 29, 9, 276, 303, 1, 0, 0,
		0, 277, 278, 10, 7, 0, 0, 278, 279, 3, 66, 33, 0, 279, 280, 3, 58, 29,
		8, 280, 303, 1, 0, 0, 0, 281, 282, 10, 6, 0, 0, 282, 283, 3, 68, 34, 0,
		283, 284, 3, 58, 29, 7, 284, 303, 1, 0, 0, 0, 285, 286, 10, 5, 0, 0, 286,
		287, 3, 70, 35, 0, 287, 288, 3, 58, 29, 6, 288, 303, 1, 0, 0, 0, 289, 290,
		10, 4, 0, 0, 290, 291, 3, 72, 36, 0, 291, 292, 3, 58, 29, 5, 292, 303,
		1, 0, 0, 0, 293, 294, 10, 3, 0, 0, 294, 295, 3, 74, 37, 0, 295, 296, 3,
		58, 29, 4, 296, 303, 1, 0, 0, 0, 297, 298, 10, 2, 0, 0, 298, 299, 5, 47,
		0, 0, 299, 303, 3, 58, 29, 3, 300, 301, 10, 10, 0, 0, 301, 303, 3, 76,
		38, 0, 302, 269, 1, 0, 0, 0, 302, 273, 1, 0, 0, 0, 302, 277, 1, 0, 0, 0,
		302, 281, 1, 0, 0, 0, 302, 285, 1, 0, 0, 0, 302, 289, 1, 0, 0, 0, 302,
		293, 1, 0, 0, 0, 302, 297, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 306,
		1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 59, 1, 0,
		0, 0, 306, 304, 1, 0, 0, 0, 307, 308, 7, 0, 0, 0, 308, 61, 1, 0, 0, 0,
		309, 311, 5, 17, 0, 0, 310, 312, 3, 108, 54, 0, 311, 310, 1, 0, 0, 0, 311,
		312, 1, 0, 0, 0, 312, 63, 1, 0, 0, 0, 313, 315, 7, 1, 0, 0, 314, 316, 3,
		108, 54, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 65, 1, 0,
		0, 0, 317, 319, 7, 0, 0, 0, 318, 320, 3, 108, 54, 0, 319, 318, 1, 0, 0,
		0, 319, 320, 1, 0, 0, 0, 320, 67, 1, 0, 0, 0, 321, 323, 7, 2, 0, 0, 322,
		324, 5, 37, 0, 0, 323, 322, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326,
		1, 0, 0, 0, 325, 327, 3, 108, 54, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1,
		0, 0, 0, 327, 69, 1, 0, 0, 0, 328, 330, 7, 3, 0, 0, 329, 331, 3, 108, 54,
		0, 330, 329, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 71, 1, 0, 0, 0, 332,
		334, 5, 19, 0, 0, 333, 335, 3, 108, 54, 0, 334, 333, 1, 0, 0, 0, 334, 335,
		1, 0, 0, 0, 335, 73, 1, 0, 0, 0, 336, 338, 7, 4, 0, 0, 337, 339, 3, 108,
		54, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 75, 1, 0, 0, 0,
		340, 342, 5, 48, 0, 0, 341, 343, 3, 78, 39, 0, 342, 341, 1, 0, 0, 0, 342,
		343, 1, 0, 0, 0, 343, 77, 1, 0, 0, 0, 344, 345, 5, 36, 0, 0, 345, 346,
		5, 50, 0, 0, 346, 79, 1, 0, 0, 0, 347, 355, 3, 96, 48, 0, 348, 355, 3,
		102, 51, 0, 349, 355, 3, 84, 42, 0, 350, 355, 3, 92, 46, 0, 351, 355, 3,
		94, 47, 0, 352, 355, 3, 124, 62, 0, 353, 355, 3, 82, 41, 0, 354, 347, 1,
		0, 0, 0, 354, 348, 1, 0, 0, 0, 354, 349, 1, 0, 0, 0, 354, 350, 1, 0, 0,
		0, 354, 351, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 353, 1, 0, 0, 0, 355,
		81, 1, 0, 0, 0, 356, 357, 5, 42, 0, 0, 357, 358, 3, 58, 29, 0, 358, 359,
		5, 43, 0, 0, 359, 83, 1, 0, 0, 0, 360, 366, 5, 51, 0, 0, 361, 363, 5, 40,
		0, 0, 362, 364, 3, 90, 45, 0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0,
		0, 364, 365, 1, 0, 0, 0, 365, 367, 5, 41, 0, 0, 366, 361, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 373, 1, 0, 0, 0, 368, 369, 5, 40, 0, 0, 369, 370,
		3, 90, 45, 0, 370, 371, 5, 41, 0, 0, 371, 373, 1, 0, 0, 0, 372, 360, 1,
		0, 0, 0, 372, 368, 1, 0, 0, 0, 373, 85, 1, 0, 0, 0, 374, 375, 3, 118, 59,
		0, 375, 376, 3, 88, 44, 0, 376, 377, 5, 11, 0, 0, 377, 87, 1, 0, 0, 0,
		378, 379, 7, 5, 0, 0, 379, 89, 1, 0, 0, 0, 380, 385, 3, 86, 43, 0, 381,
		382, 5, 46, 0, 0, 382, 384, 3, 86, 43, 0, 383, 381, 1, 0, 0, 0, 384, 387,
		1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 389, 1, 0,
		0, 0, 387, 385, 1, 0, 0, 0, 388, 390, 5, 46, 0, 0, 389, 388, 1, 0, 0, 0,
		389, 390, 1, 0, 0, 0, 390, 91, 1, 0, 0, 0, 391, 392, 3, 84, 42, 0, 392,
		393, 5, 49, 0, 0, 393, 93, 1, 0, 0, 0, 394, 395, 3, 84, 42, 0, 395, 396,
		5, 36, 0, 0, 396, 397, 5, 50, 0, 0, 397, 403, 1, 0, 0, 0, 398, 399, 3,
		92, 46, 0, 399, 400, 5, 36, 0, 0, 400, 401, 5, 50, 0, 0, 401, 403, 1, 0,
		0, 0, 402, 394, 1, 0, 0, 0, 402, 398, 1, 0, 0, 0, 403, 95, 1, 0, 0, 0,
		404, 405, 5, 39, 0, 0, 405, 414, 5, 42, 0, 0, 406, 411, 3, 98, 49, 0, 407,
		408, 5, 46, 0, 0, 408, 410, 3, 98, 49, 0, 409, 407, 1, 0, 0, 0, 410, 413,
		1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 415, 1, 0,
		0, 0, 413, 411, 1, 0, 0, 0, 414, 406, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0,
		415, 416, 1, 0, 0, 0, 416, 417, 5, 43, 0, 0, 417, 97, 1, 0, 0, 0, 418,
		421, 3, 124, 62, 0, 419, 421, 3, 58, 29, 0, 420, 418, 1, 0, 0, 0, 420,
		419, 1, 0, 0, 0, 421, 99, 1, 0, 0, 0, 422, 431, 5, 42, 0, 0, 423, 428,
		3, 98, 49, 0, 424, 425, 5, 46, 0, 0, 425, 427, 3, 98, 49, 0, 426, 424,
		1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 428, 429, 1, 0,
		0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 423, 1, 0, 0, 0,
		431, 432, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434, 5, 43, 0, 0, 434,
		101, 1, 0, 0, 0, 435, 436, 5, 38, 0, 0, 436, 451, 3, 100, 50, 0, 437, 440,
		5, 38, 0, 0, 438, 441, 3, 104, 52, 0, 439, 441, 3, 106, 53, 0, 440, 438,
		1, 0, 0, 0, 440, 439, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 3, 100,
		50, 0, 443, 451, 1, 0, 0, 0, 444, 445, 5, 38, 0, 0, 445, 448, 3, 100, 50,
		0, 446, 449, 3, 104, 52, 0, 447, 449, 3, 106, 53, 0, 448, 446, 1, 0, 0,
		0, 448, 447, 1, 0, 0, 0, 449, 451, 1, 0, 0, 0, 450, 435, 1, 0, 0, 0, 450,
		437, 1, 0, 0, 0, 450, 444, 1, 0, 0, 0, 451, 103, 1, 0, 0, 0, 452, 453,
		5, 30, 0, 0, 453, 454, 3, 120, 60, 0, 454, 105, 1, 0, 0, 0, 455, 456, 5,
		31, 0, 0, 456, 457, 3, 120, 60, 0, 457, 107, 1, 0, 0, 0, 458, 461, 3, 110,
		55, 0, 459, 461, 3, 112, 56, 0, 460, 458, 1, 0, 0, 0, 460, 459, 1, 0, 0,
		0, 461, 464, 1, 0, 0, 0, 462, 465, 3, 114, 57, 0, 463, 465, 3, 116, 58,
		0, 464, 462, 1, 0, 0, 0, 464, 463, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465,
		109, 1, 0, 0, 0, 466, 467, 5, 32, 0, 0, 467, 468, 3, 120, 60, 0, 468, 111,
		1, 0, 0, 0, 469, 470, 5, 33, 0, 0, 470, 471, 3, 120, 60, 0, 471, 113, 1,
		0, 0, 0, 472, 474, 5, 34, 0, 0, 473, 475, 3, 120, 60, 0, 474, 473, 1, 0,
		0, 0, 474, 475, 1, 0, 0, 0, 475, 115, 1, 0, 0, 0, 476, 478, 5, 35, 0, 0,
		477, 479, 3, 120, 60, 0, 478, 477, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479,
		117, 1, 0, 0, 0, 480, 484, 3, 122, 61, 0, 481, 484, 5, 51, 0, 0, 482, 484,
		5, 52, 0, 0, 483, 480, 1, 0, 0, 0, 483, 481, 1, 0, 0, 0, 483, 482, 1, 0,
		0, 0, 484, 119, 1, 0, 0, 0, 485, 494, 5, 42, 0, 0, 486, 491, 3, 118, 59,
		0, 487, 488, 5, 46, 0, 0, 488, 490, 3, 118, 59, 0, 489, 487, 1, 0, 0, 0,
		490, 493, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 491, 492, 1, 0, 0, 0, 492,
		495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 494, 486, 1, 0, 0, 0, 494, 495,
		1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 5, 43, 0, 0, 497, 121, 1, 0,
		0, 0, 498, 499, 7, 6, 0, 0, 499, 123, 1, 0, 0, 0, 500, 501, 7, 7, 0, 0,
		501, 125, 1, 0, 0, 0, 40, 131, 137, 146, 160, 168, 187, 267, 302, 304,
		311, 315, 319, 323, 326, 330, 334, 338, 342, 354, 363, 366, 372, 385, 389,
		402, 411, 414, 420, 428, 431, 440, 448, 450, 460, 464, 474, 478, 483, 491,
		494,
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

// PromQLExtensionParserInit initializes any static state used to implement PromQLExtensionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromQLExtensionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLExtensionParserInit() {
	staticData := &PromQLExtensionParserParserStaticData
	staticData.once.Do(promqlextensionparserParserInit)
}

// NewPromQLExtensionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromQLExtensionParser(input antlr.TokenStream) *PromQLExtensionParser {
	PromQLExtensionParserInit()
	this := new(PromQLExtensionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromQLExtensionParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromQLExtensionParser.g4"

	return this
}

// PromQLExtensionParser tokens.
const (
	PromQLExtensionParserEOF                  = antlr.TokenEOF
	PromQLExtensionParserEX_ID                = 1
	PromQLExtensionParserEX_TRUE              = 2
	PromQLExtensionParserEX_FALSE             = 3
	PromQLExtensionParserEX_T                 = 4
	PromQLExtensionParserEX_COLON             = 5
	PromQLExtensionParserEX_DOT               = 6
	PromQLExtensionParserEX_POSITIVE_INTEGER  = 7
	PromQLExtensionParserEX_TWO_DIGITS        = 8
	PromQLExtensionParserEX_DIGITS            = 9
	PromQLExtensionParserNUMBER               = 10
	PromQLExtensionParserSTRING               = 11
	PromQLExtensionParserADD                  = 12
	PromQLExtensionParserSUB                  = 13
	PromQLExtensionParserMULT                 = 14
	PromQLExtensionParserDIV                  = 15
	PromQLExtensionParserMOD                  = 16
	PromQLExtensionParserPOW                  = 17
	PromQLExtensionParserAND                  = 18
	PromQLExtensionParserOR                   = 19
	PromQLExtensionParserUNLESS               = 20
	PromQLExtensionParserEQ                   = 21
	PromQLExtensionParserDEQ                  = 22
	PromQLExtensionParserNE                   = 23
	PromQLExtensionParserGT                   = 24
	PromQLExtensionParserLT                   = 25
	PromQLExtensionParserGE                   = 26
	PromQLExtensionParserLE                   = 27
	PromQLExtensionParserRE                   = 28
	PromQLExtensionParserNRE                  = 29
	PromQLExtensionParserBY                   = 30
	PromQLExtensionParserWITHOUT              = 31
	PromQLExtensionParserON                   = 32
	PromQLExtensionParserIGNORING             = 33
	PromQLExtensionParserGROUP_LEFT           = 34
	PromQLExtensionParserGROUP_RIGHT          = 35
	PromQLExtensionParserOFFSET               = 36
	PromQLExtensionParserBOOL                 = 37
	PromQLExtensionParserAGGREGATION_OPERATOR = 38
	PromQLExtensionParserFUNCTION             = 39
	PromQLExtensionParserLEFT_BRACE           = 40
	PromQLExtensionParserRIGHT_BRACE          = 41
	PromQLExtensionParserLEFT_PAREN           = 42
	PromQLExtensionParserRIGHT_PAREN          = 43
	PromQLExtensionParserLEFT_BRACKET         = 44
	PromQLExtensionParserRIGHT_BRACKET        = 45
	PromQLExtensionParserCOMMA                = 46
	PromQLExtensionParserAT                   = 47
	PromQLExtensionParserSUBQUERY_RANGE       = 48
	PromQLExtensionParserTIME_RANGE           = 49
	PromQLExtensionParserDURATION             = 50
	PromQLExtensionParserMETRIC_NAME          = 51
	PromQLExtensionParserLABEL_NAME           = 52
	PromQLExtensionParserWS                   = 53
	PromQLExtensionParserSL_COMMENT           = 54
)

// PromQLExtensionParser rules.
const (
	PromQLExtensionParserRULE_promqlx                   = 0
	PromQLExtensionParserRULE_ex_statement_list         = 1
	PromQLExtensionParserRULE_ex_statement              = 2
	PromQLExtensionParserRULE_ex_alias_def              = 3
	PromQLExtensionParserRULE_ex_macro_def              = 4
	PromQLExtensionParserRULE_ex_block                  = 5
	PromQLExtensionParserRULE_ex_arg_list               = 6
	PromQLExtensionParserRULE_ex_if_statement           = 7
	PromQLExtensionParserRULE_ex_condition              = 8
	PromQLExtensionParserRULE_ex_compareVectorOperation = 9
	PromQLExtensionParserRULE_ex_trueConst              = 10
	PromQLExtensionParserRULE_ex_falseConst             = 11
	PromQLExtensionParserRULE_ex_time_instant           = 12
	PromQLExtensionParserRULE_ex_iso_date_time          = 13
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf  = 14
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhms   = 15
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhm    = 16
	PromQLExtensionParserRULE_ex_iso_date_time_ymdh     = 17
	PromQLExtensionParserRULE_ex_iso_date_time_ymd      = 18
	PromQLExtensionParserRULE_ex_iso_date_time_ym       = 19
	PromQLExtensionParserRULE_ex_iso_date_time_y        = 20
	PromQLExtensionParserRULE_ex_year                   = 21
	PromQLExtensionParserRULE_ex_month                  = 22
	PromQLExtensionParserRULE_ex_day                    = 23
	PromQLExtensionParserRULE_ex_hour                   = 24
	PromQLExtensionParserRULE_ex_minutes                = 25
	PromQLExtensionParserRULE_ex_seconds                = 26
	PromQLExtensionParserRULE_ex_frac_sec               = 27
	PromQLExtensionParserRULE_expression                = 28
	PromQLExtensionParserRULE_vectorOperation           = 29
	PromQLExtensionParserRULE_unaryOp                   = 30
	PromQLExtensionParserRULE_powOp                     = 31
	PromQLExtensionParserRULE_multOp                    = 32
	PromQLExtensionParserRULE_addOp                     = 33
	PromQLExtensionParserRULE_compareOp                 = 34
	PromQLExtensionParserRULE_andUnlessOp               = 35
	PromQLExtensionParserRULE_orOp                      = 36
	PromQLExtensionParserRULE_vectorMatchOp             = 37
	PromQLExtensionParserRULE_subqueryOp                = 38
	PromQLExtensionParserRULE_offsetOp                  = 39
	PromQLExtensionParserRULE_vector                    = 40
	PromQLExtensionParserRULE_parens                    = 41
	PromQLExtensionParserRULE_instantSelector           = 42
	PromQLExtensionParserRULE_labelMatcher              = 43
	PromQLExtensionParserRULE_labelMatcherOperator      = 44
	PromQLExtensionParserRULE_labelMatcherList          = 45
	PromQLExtensionParserRULE_matrixSelector            = 46
	PromQLExtensionParserRULE_offset                    = 47
	PromQLExtensionParserRULE_function_                 = 48
	PromQLExtensionParserRULE_parameter                 = 49
	PromQLExtensionParserRULE_parameterList             = 50
	PromQLExtensionParserRULE_aggregation               = 51
	PromQLExtensionParserRULE_by                        = 52
	PromQLExtensionParserRULE_without                   = 53
	PromQLExtensionParserRULE_grouping                  = 54
	PromQLExtensionParserRULE_on_                       = 55
	PromQLExtensionParserRULE_ignoring                  = 56
	PromQLExtensionParserRULE_groupLeft                 = 57
	PromQLExtensionParserRULE_groupRight                = 58
	PromQLExtensionParserRULE_labelName                 = 59
	PromQLExtensionParserRULE_labelNameList             = 60
	PromQLExtensionParserRULE_keyword                   = 61
	PromQLExtensionParserRULE_literal                   = 62
)

// IPromqlxContext is an interface to support dynamic dispatch.
type IPromqlxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_statement_list() IEx_statement_listContext

	// IsPromqlxContext differentiates from other interfaces.
	IsPromqlxContext()
}

type PromqlxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromqlxContext() *PromqlxContext {
	var p = new(PromqlxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_promqlx
	return p
}

func InitEmptyPromqlxContext(p *PromqlxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_promqlx
}

func (*PromqlxContext) IsPromqlxContext() {}

func NewPromqlxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromqlxContext {
	var p = new(PromqlxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_promqlx

	return p
}

func (s *PromqlxContext) GetParser() antlr.Parser { return s.parser }

func (s *PromqlxContext) Ex_statement_list() IEx_statement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_statement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_statement_listContext)
}

func (s *PromqlxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromqlxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromqlxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterPromqlx(s)
	}
}

func (s *PromqlxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitPromqlx(s)
	}
}

func (p *PromQLExtensionParser) Promqlx() (localctx IPromqlxContext) {
	localctx = NewPromqlxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromQLExtensionParserRULE_promqlx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Ex_statement_list()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_statement_listContext is an interface to support dynamic dispatch.
type IEx_statement_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEx_statement() []IEx_statementContext
	Ex_statement(i int) IEx_statementContext

	// IsEx_statement_listContext differentiates from other interfaces.
	IsEx_statement_listContext()
}

type Ex_statement_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_statement_listContext() *Ex_statement_listContext {
	var p = new(Ex_statement_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement_list
	return p
}

func InitEmptyEx_statement_listContext(p *Ex_statement_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement_list
}

func (*Ex_statement_listContext) IsEx_statement_listContext() {}

func NewEx_statement_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_statement_listContext {
	var p = new(Ex_statement_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement_list

	return p
}

func (s *Ex_statement_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_statement_listContext) AllEx_statement() []IEx_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEx_statementContext); ok {
			len++
		}
	}

	tst := make([]IEx_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEx_statementContext); ok {
			tst[i] = t.(IEx_statementContext)
			i++
		}
	}

	return tst
}

func (s *Ex_statement_listContext) Ex_statement(i int) IEx_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_statementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_statementContext)
}

func (s *Ex_statement_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_statement_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_statement_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_statement_list(s)
	}
}

func (s *Ex_statement_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_statement_list(s)
	}
}

func (p *PromQLExtensionParser) Ex_statement_list() (localctx IEx_statement_listContext) {
	localctx = NewEx_statement_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromQLExtensionParserRULE_ex_statement_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2258122005560334) != 0) {
		{
			p.SetState(128)
			p.Ex_statement()
		}

		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_statementContext is an interface to support dynamic dispatch.
type IEx_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_alias_def() IEx_alias_defContext
	Ex_macro_def() IEx_macro_defContext
	Ex_if_statement() IEx_if_statementContext
	Expression() IExpressionContext

	// IsEx_statementContext differentiates from other interfaces.
	IsEx_statementContext()
}

type Ex_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_statementContext() *Ex_statementContext {
	var p = new(Ex_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement
	return p
}

func InitEmptyEx_statementContext(p *Ex_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement
}

func (*Ex_statementContext) IsEx_statementContext() {}

func NewEx_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_statementContext {
	var p = new(Ex_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_statement

	return p
}

func (s *Ex_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_statementContext) Ex_alias_def() IEx_alias_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_alias_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_alias_defContext)
}

func (s *Ex_statementContext) Ex_macro_def() IEx_macro_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_macro_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_macro_defContext)
}

func (s *Ex_statementContext) Ex_if_statement() IEx_if_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_if_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_if_statementContext)
}

func (s *Ex_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Ex_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_statement(s)
	}
}

func (s *Ex_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_statement(s)
	}
}

func (p *PromQLExtensionParser) Ex_statement() (localctx IEx_statementContext) {
	localctx = NewEx_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromQLExtensionParserRULE_ex_statement)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Ex_alias_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Ex_macro_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(135)
			p.Ex_if_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Expression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_alias_defContext is an interface to support dynamic dispatch.
type IEx_alias_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_ID() antlr.TerminalNode
	EQ() antlr.TerminalNode
	Expression() IExpressionContext

	// IsEx_alias_defContext differentiates from other interfaces.
	IsEx_alias_defContext()
}

type Ex_alias_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_alias_defContext() *Ex_alias_defContext {
	var p = new(Ex_alias_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_def
	return p
}

func InitEmptyEx_alias_defContext(p *Ex_alias_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_def
}

func (*Ex_alias_defContext) IsEx_alias_defContext() {}

func NewEx_alias_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_alias_defContext {
	var p = new(Ex_alias_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_def

	return p
}

func (s *Ex_alias_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_alias_defContext) EX_ID() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_ID, 0)
}

func (s *Ex_alias_defContext) EQ() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEQ, 0)
}

func (s *Ex_alias_defContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Ex_alias_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_alias_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_alias_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_alias_def(s)
	}
}

func (s *Ex_alias_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_alias_def(s)
	}
}

func (p *PromQLExtensionParser) Ex_alias_def() (localctx IEx_alias_defContext) {
	localctx = NewEx_alias_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromQLExtensionParserRULE_ex_alias_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(PromQLExtensionParserEX_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(PromQLExtensionParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Expression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_macro_defContext is an interface to support dynamic dispatch.
type IEx_macro_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_ID() antlr.TerminalNode
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	Ex_block() IEx_blockContext
	Ex_arg_list() IEx_arg_listContext

	// IsEx_macro_defContext differentiates from other interfaces.
	IsEx_macro_defContext()
}

type Ex_macro_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_macro_defContext() *Ex_macro_defContext {
	var p = new(Ex_macro_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_macro_def
	return p
}

func InitEmptyEx_macro_defContext(p *Ex_macro_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_macro_def
}

func (*Ex_macro_defContext) IsEx_macro_defContext() {}

func NewEx_macro_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_macro_defContext {
	var p = new(Ex_macro_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_macro_def

	return p
}

func (s *Ex_macro_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_macro_defContext) EX_ID() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_ID, 0)
}

func (s *Ex_macro_defContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *Ex_macro_defContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *Ex_macro_defContext) Ex_block() IEx_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_blockContext)
}

func (s *Ex_macro_defContext) Ex_arg_list() IEx_arg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_arg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_arg_listContext)
}

func (s *Ex_macro_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_macro_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_macro_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_macro_def(s)
	}
}

func (s *Ex_macro_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_macro_def(s)
	}
}

func (p *PromQLExtensionParser) Ex_macro_def() (localctx IEx_macro_defContext) {
	localctx = NewEx_macro_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromQLExtensionParserRULE_ex_macro_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(PromQLExtensionParserEX_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2258122005560320) != 0 {
		{
			p.SetState(145)
			p.Ex_arg_list()
		}

	}
	{
		p.SetState(148)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Ex_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_blockContext is an interface to support dynamic dispatch.
type IEx_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACE() antlr.TerminalNode
	Ex_statement_list() IEx_statement_listContext
	RIGHT_BRACE() antlr.TerminalNode

	// IsEx_blockContext differentiates from other interfaces.
	IsEx_blockContext()
}

type Ex_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_blockContext() *Ex_blockContext {
	var p = new(Ex_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_block
	return p
}

func InitEmptyEx_blockContext(p *Ex_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_block
}

func (*Ex_blockContext) IsEx_blockContext() {}

func NewEx_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_blockContext {
	var p = new(Ex_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_block

	return p
}

func (s *Ex_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_blockContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_BRACE, 0)
}

func (s *Ex_blockContext) Ex_statement_list() IEx_statement_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_statement_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_statement_listContext)
}

func (s *Ex_blockContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_BRACE, 0)
}

func (s *Ex_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_block(s)
	}
}

func (s *Ex_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_block(s)
	}
}

func (p *PromQLExtensionParser) Ex_block() (localctx IEx_blockContext) {
	localctx = NewEx_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromQLExtensionParserRULE_ex_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(PromQLExtensionParserLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Ex_statement_list()
	}
	{
		p.SetState(153)
		p.Match(PromQLExtensionParserRIGHT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_arg_listContext is an interface to support dynamic dispatch.
type IEx_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEx_arg_listContext differentiates from other interfaces.
	IsEx_arg_listContext()
}

type Ex_arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_arg_listContext() *Ex_arg_listContext {
	var p = new(Ex_arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_arg_list
	return p
}

func InitEmptyEx_arg_listContext(p *Ex_arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_arg_list
}

func (*Ex_arg_listContext) IsEx_arg_listContext() {}

func NewEx_arg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_arg_listContext {
	var p = new(Ex_arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_arg_list

	return p
}

func (s *Ex_arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_arg_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Ex_arg_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Ex_arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserCOMMA)
}

func (s *Ex_arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserCOMMA, i)
}

func (s *Ex_arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_arg_list(s)
	}
}

func (s *Ex_arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_arg_list(s)
	}
}

func (p *PromQLExtensionParser) Ex_arg_list() (localctx IEx_arg_listContext) {
	localctx = NewEx_arg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromQLExtensionParserRULE_ex_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Expression()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromQLExtensionParserCOMMA {
		{
			p.SetState(156)
			p.Match(PromQLExtensionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Expression()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_if_statementContext is an interface to support dynamic dispatch.
type IEx_if_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_condition() IEx_conditionContext
	Ex_block() IEx_blockContext

	// IsEx_if_statementContext differentiates from other interfaces.
	IsEx_if_statementContext()
}

type Ex_if_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_if_statementContext() *Ex_if_statementContext {
	var p = new(Ex_if_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_if_statement
	return p
}

func InitEmptyEx_if_statementContext(p *Ex_if_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_if_statement
}

func (*Ex_if_statementContext) IsEx_if_statementContext() {}

func NewEx_if_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_if_statementContext {
	var p = new(Ex_if_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_if_statement

	return p
}

func (s *Ex_if_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_if_statementContext) Ex_condition() IEx_conditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_conditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_conditionContext)
}

func (s *Ex_if_statementContext) Ex_block() IEx_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_blockContext)
}

func (s *Ex_if_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_if_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_if_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_if_statement(s)
	}
}

func (s *Ex_if_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_if_statement(s)
	}
}

func (p *PromQLExtensionParser) Ex_if_statement() (localctx IEx_if_statementContext) {
	localctx = NewEx_if_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromQLExtensionParserRULE_ex_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Ex_condition()
	}
	{
		p.SetState(163)
		p.Ex_block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_conditionContext is an interface to support dynamic dispatch.
type IEx_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_compareVectorOperation() IEx_compareVectorOperationContext
	Ex_trueConst() IEx_trueConstContext
	Ex_falseConst() IEx_falseConstContext

	// IsEx_conditionContext differentiates from other interfaces.
	IsEx_conditionContext()
}

type Ex_conditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_conditionContext() *Ex_conditionContext {
	var p = new(Ex_conditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_condition
	return p
}

func InitEmptyEx_conditionContext(p *Ex_conditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_condition
}

func (*Ex_conditionContext) IsEx_conditionContext() {}

func NewEx_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_conditionContext {
	var p = new(Ex_conditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_condition

	return p
}

func (s *Ex_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_conditionContext) Ex_compareVectorOperation() IEx_compareVectorOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_compareVectorOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_compareVectorOperationContext)
}

func (s *Ex_conditionContext) Ex_trueConst() IEx_trueConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_trueConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_trueConstContext)
}

func (s *Ex_conditionContext) Ex_falseConst() IEx_falseConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_falseConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_falseConstContext)
}

func (s *Ex_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_condition(s)
	}
}

func (s *Ex_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_condition(s)
	}
}

func (p *PromQLExtensionParser) Ex_condition() (localctx IEx_conditionContext) {
	localctx = NewEx_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromQLExtensionParserRULE_ex_condition)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserNUMBER, PromQLExtensionParserSTRING, PromQLExtensionParserADD, PromQLExtensionParserSUB, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION, PromQLExtensionParserLEFT_BRACE, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Ex_compareVectorOperation()
		}

	case PromQLExtensionParserEX_TRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Ex_trueConst()
		}

	case PromQLExtensionParserEX_FALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Ex_falseConst()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_compareVectorOperationContext is an interface to support dynamic dispatch.
type IEx_compareVectorOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVectorOperation() []IVectorOperationContext
	VectorOperation(i int) IVectorOperationContext
	CompareOp() ICompareOpContext

	// IsEx_compareVectorOperationContext differentiates from other interfaces.
	IsEx_compareVectorOperationContext()
}

type Ex_compareVectorOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_compareVectorOperationContext() *Ex_compareVectorOperationContext {
	var p = new(Ex_compareVectorOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_compareVectorOperation
	return p
}

func InitEmptyEx_compareVectorOperationContext(p *Ex_compareVectorOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_compareVectorOperation
}

func (*Ex_compareVectorOperationContext) IsEx_compareVectorOperationContext() {}

func NewEx_compareVectorOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_compareVectorOperationContext {
	var p = new(Ex_compareVectorOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_compareVectorOperation

	return p
}

func (s *Ex_compareVectorOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_compareVectorOperationContext) AllVectorOperation() []IVectorOperationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVectorOperationContext); ok {
			len++
		}
	}

	tst := make([]IVectorOperationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVectorOperationContext); ok {
			tst[i] = t.(IVectorOperationContext)
			i++
		}
	}

	return tst
}

func (s *Ex_compareVectorOperationContext) VectorOperation(i int) IVectorOperationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorOperationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorOperationContext)
}

func (s *Ex_compareVectorOperationContext) CompareOp() ICompareOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareOpContext)
}

func (s *Ex_compareVectorOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_compareVectorOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_compareVectorOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_compareVectorOperation(s)
	}
}

func (s *Ex_compareVectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_compareVectorOperation(s)
	}
}

func (p *PromQLExtensionParser) Ex_compareVectorOperation() (localctx IEx_compareVectorOperationContext) {
	localctx = NewEx_compareVectorOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromQLExtensionParserRULE_ex_compareVectorOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.vectorOperation(0)
	}
	{
		p.SetState(171)
		p.CompareOp()
	}
	{
		p.SetState(172)
		p.vectorOperation(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_trueConstContext is an interface to support dynamic dispatch.
type IEx_trueConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TRUE() antlr.TerminalNode

	// IsEx_trueConstContext differentiates from other interfaces.
	IsEx_trueConstContext()
}

type Ex_trueConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_trueConstContext() *Ex_trueConstContext {
	var p = new(Ex_trueConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_trueConst
	return p
}

func InitEmptyEx_trueConstContext(p *Ex_trueConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_trueConst
}

func (*Ex_trueConstContext) IsEx_trueConstContext() {}

func NewEx_trueConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_trueConstContext {
	var p = new(Ex_trueConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_trueConst

	return p
}

func (s *Ex_trueConstContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_trueConstContext) EX_TRUE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TRUE, 0)
}

func (s *Ex_trueConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_trueConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_trueConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_trueConst(s)
	}
}

func (s *Ex_trueConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_trueConst(s)
	}
}

func (p *PromQLExtensionParser) Ex_trueConst() (localctx IEx_trueConstContext) {
	localctx = NewEx_trueConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromQLExtensionParserRULE_ex_trueConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(PromQLExtensionParserEX_TRUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_falseConstContext is an interface to support dynamic dispatch.
type IEx_falseConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_FALSE() antlr.TerminalNode

	// IsEx_falseConstContext differentiates from other interfaces.
	IsEx_falseConstContext()
}

type Ex_falseConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_falseConstContext() *Ex_falseConstContext {
	var p = new(Ex_falseConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_falseConst
	return p
}

func InitEmptyEx_falseConstContext(p *Ex_falseConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_falseConst
}

func (*Ex_falseConstContext) IsEx_falseConstContext() {}

func NewEx_falseConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_falseConstContext {
	var p = new(Ex_falseConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_falseConst

	return p
}

func (s *Ex_falseConstContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_falseConstContext) EX_FALSE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_FALSE, 0)
}

func (s *Ex_falseConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_falseConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_falseConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_falseConst(s)
	}
}

func (s *Ex_falseConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_falseConst(s)
	}
}

func (p *PromQLExtensionParser) Ex_falseConst() (localctx IEx_falseConstContext) {
	localctx = NewEx_falseConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromQLExtensionParserRULE_ex_falseConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(PromQLExtensionParserEX_FALSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_time_instantContext is an interface to support dynamic dispatch.
type IEx_time_instantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_iso_date_time() IEx_iso_date_timeContext

	// IsEx_time_instantContext differentiates from other interfaces.
	IsEx_time_instantContext()
}

type Ex_time_instantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_time_instantContext() *Ex_time_instantContext {
	var p = new(Ex_time_instantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant
	return p
}

func InitEmptyEx_time_instantContext(p *Ex_time_instantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant
}

func (*Ex_time_instantContext) IsEx_time_instantContext() {}

func NewEx_time_instantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_time_instantContext {
	var p = new(Ex_time_instantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant

	return p
}

func (s *Ex_time_instantContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_time_instantContext) Ex_iso_date_time() IEx_iso_date_timeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_timeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_timeContext)
}

func (s *Ex_time_instantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_time_instantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_time_instantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_time_instant(s)
	}
}

func (s *Ex_time_instantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_time_instant(s)
	}
}

func (p *PromQLExtensionParser) Ex_time_instant() (localctx IEx_time_instantContext) {
	localctx = NewEx_time_instantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromQLExtensionParserRULE_ex_time_instant)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Ex_iso_date_time()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_timeContext is an interface to support dynamic dispatch.
type IEx_iso_date_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_iso_date_time_ymdhmsf() IEx_iso_date_time_ymdhmsfContext
	Ex_iso_date_time_ymdhms() IEx_iso_date_time_ymdhmsContext
	Ex_iso_date_time_ymdhm() IEx_iso_date_time_ymdhmContext
	Ex_iso_date_time_ymdh() IEx_iso_date_time_ymdhContext
	Ex_iso_date_time_ymd() IEx_iso_date_time_ymdContext
	Ex_iso_date_time_ym() IEx_iso_date_time_ymContext
	Ex_iso_date_time_y() IEx_iso_date_time_yContext

	// IsEx_iso_date_timeContext differentiates from other interfaces.
	IsEx_iso_date_timeContext()
}

type Ex_iso_date_timeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_timeContext() *Ex_iso_date_timeContext {
	var p = new(Ex_iso_date_timeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time
	return p
}

func InitEmptyEx_iso_date_timeContext(p *Ex_iso_date_timeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time
}

func (*Ex_iso_date_timeContext) IsEx_iso_date_timeContext() {}

func NewEx_iso_date_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_timeContext {
	var p = new(Ex_iso_date_timeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time

	return p
}

func (s *Ex_iso_date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ymdhmsf() IEx_iso_date_time_ymdhmsfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymdhmsfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymdhmsfContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ymdhms() IEx_iso_date_time_ymdhmsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymdhmsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymdhmsContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ymdhm() IEx_iso_date_time_ymdhmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymdhmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymdhmContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ymdh() IEx_iso_date_time_ymdhContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymdhContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymdhContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ymd() IEx_iso_date_time_ymdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymdContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_ym() IEx_iso_date_time_ymContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_ymContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_ymContext)
}

func (s *Ex_iso_date_timeContext) Ex_iso_date_time_y() IEx_iso_date_time_yContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_iso_date_time_yContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_iso_date_time_yContext)
}

func (s *Ex_iso_date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time(s)
	}
}

func (s *Ex_iso_date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time() (localctx IEx_iso_date_timeContext) {
	localctx = NewEx_iso_date_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromQLExtensionParserRULE_ex_iso_date_time)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Ex_iso_date_time_ymdhmsf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Ex_iso_date_time_ymdhms()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Ex_iso_date_time_ymdhm()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Ex_iso_date_time_ymdh()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Ex_iso_date_time_ymd()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)
			p.Ex_iso_date_time_ym()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(186)
			p.Ex_iso_date_time_y()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymdhmsfContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymdhmsfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Ex_month() IEx_monthContext
	Ex_day() IEx_dayContext
	EX_T() antlr.TerminalNode
	Ex_hour() IEx_hourContext
	AllEX_COLON() []antlr.TerminalNode
	EX_COLON(i int) antlr.TerminalNode
	Ex_minutes() IEx_minutesContext
	Ex_seconds() IEx_secondsContext
	EX_DOT() antlr.TerminalNode
	Ex_frac_sec() IEx_frac_secContext

	// IsEx_iso_date_time_ymdhmsfContext differentiates from other interfaces.
	IsEx_iso_date_time_ymdhmsfContext()
}

type Ex_iso_date_time_ymdhmsfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymdhmsfContext() *Ex_iso_date_time_ymdhmsfContext {
	var p = new(Ex_iso_date_time_ymdhmsfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf
	return p
}

func InitEmptyEx_iso_date_time_ymdhmsfContext(p *Ex_iso_date_time_ymdhmsfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf
}

func (*Ex_iso_date_time_ymdhmsfContext) IsEx_iso_date_time_ymdhmsfContext() {}

func NewEx_iso_date_time_ymdhmsfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymdhmsfContext {
	var p = new(Ex_iso_date_time_ymdhmsfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf

	return p
}

func (s *Ex_iso_date_time_ymdhmsfContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserSUB)
}

func (s *Ex_iso_date_time_ymdhmsfContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, i)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_day() IEx_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_dayContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) EX_T() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_T, 0)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_hour() IEx_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_hourContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) AllEX_COLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserEX_COLON)
}

func (s *Ex_iso_date_time_ymdhmsfContext) EX_COLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_COLON, i)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_minutes() IEx_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_minutesContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_seconds() IEx_secondsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_secondsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_secondsContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) EX_DOT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_DOT, 0)
}

func (s *Ex_iso_date_time_ymdhmsfContext) Ex_frac_sec() IEx_frac_secContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_frac_secContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_frac_secContext)
}

func (s *Ex_iso_date_time_ymdhmsfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymdhmsfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymdhmsfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ymdhmsf(s)
	}
}

func (s *Ex_iso_date_time_ymdhmsfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ymdhmsf(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ymdhmsf() (localctx IEx_iso_date_time_ymdhmsfContext) {
	localctx = NewEx_iso_date_time_ymdhmsfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Ex_year()
	}
	{
		p.SetState(190)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Ex_month()
	}
	{
		p.SetState(192)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Ex_day()
	}
	{
		p.SetState(194)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Ex_hour()
	}
	{
		p.SetState(196)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Ex_minutes()
	}
	{
		p.SetState(198)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Ex_seconds()
	}
	{
		p.SetState(200)
		p.Match(PromQLExtensionParserEX_DOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Ex_frac_sec()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymdhmsContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymdhmsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Ex_month() IEx_monthContext
	Ex_day() IEx_dayContext
	EX_T() antlr.TerminalNode
	Ex_hour() IEx_hourContext
	AllEX_COLON() []antlr.TerminalNode
	EX_COLON(i int) antlr.TerminalNode
	Ex_minutes() IEx_minutesContext
	Ex_seconds() IEx_secondsContext

	// IsEx_iso_date_time_ymdhmsContext differentiates from other interfaces.
	IsEx_iso_date_time_ymdhmsContext()
}

type Ex_iso_date_time_ymdhmsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymdhmsContext() *Ex_iso_date_time_ymdhmsContext {
	var p = new(Ex_iso_date_time_ymdhmsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhms
	return p
}

func InitEmptyEx_iso_date_time_ymdhmsContext(p *Ex_iso_date_time_ymdhmsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhms
}

func (*Ex_iso_date_time_ymdhmsContext) IsEx_iso_date_time_ymdhmsContext() {}

func NewEx_iso_date_time_ymdhmsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymdhmsContext {
	var p = new(Ex_iso_date_time_ymdhmsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhms

	return p
}

func (s *Ex_iso_date_time_ymdhmsContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymdhmsContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserSUB)
}

func (s *Ex_iso_date_time_ymdhmsContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, i)
}

func (s *Ex_iso_date_time_ymdhmsContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) Ex_day() IEx_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_dayContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) EX_T() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_T, 0)
}

func (s *Ex_iso_date_time_ymdhmsContext) Ex_hour() IEx_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_hourContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) AllEX_COLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserEX_COLON)
}

func (s *Ex_iso_date_time_ymdhmsContext) EX_COLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_COLON, i)
}

func (s *Ex_iso_date_time_ymdhmsContext) Ex_minutes() IEx_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_minutesContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) Ex_seconds() IEx_secondsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_secondsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_secondsContext)
}

func (s *Ex_iso_date_time_ymdhmsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymdhmsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymdhmsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ymdhms(s)
	}
}

func (s *Ex_iso_date_time_ymdhmsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ymdhms(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ymdhms() (localctx IEx_iso_date_time_ymdhmsContext) {
	localctx = NewEx_iso_date_time_ymdhmsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromQLExtensionParserRULE_ex_iso_date_time_ymdhms)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Ex_year()
	}
	{
		p.SetState(204)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Ex_month()
	}
	{
		p.SetState(206)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Ex_day()
	}
	{
		p.SetState(208)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Ex_hour()
	}
	{
		p.SetState(210)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Ex_minutes()
	}
	{
		p.SetState(212)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Ex_seconds()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymdhmContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymdhmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Ex_month() IEx_monthContext
	Ex_day() IEx_dayContext
	EX_T() antlr.TerminalNode
	Ex_hour() IEx_hourContext
	EX_COLON() antlr.TerminalNode
	Ex_minutes() IEx_minutesContext

	// IsEx_iso_date_time_ymdhmContext differentiates from other interfaces.
	IsEx_iso_date_time_ymdhmContext()
}

type Ex_iso_date_time_ymdhmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymdhmContext() *Ex_iso_date_time_ymdhmContext {
	var p = new(Ex_iso_date_time_ymdhmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhm
	return p
}

func InitEmptyEx_iso_date_time_ymdhmContext(p *Ex_iso_date_time_ymdhmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhm
}

func (*Ex_iso_date_time_ymdhmContext) IsEx_iso_date_time_ymdhmContext() {}

func NewEx_iso_date_time_ymdhmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymdhmContext {
	var p = new(Ex_iso_date_time_ymdhmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdhm

	return p
}

func (s *Ex_iso_date_time_ymdhmContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymdhmContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymdhmContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserSUB)
}

func (s *Ex_iso_date_time_ymdhmContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, i)
}

func (s *Ex_iso_date_time_ymdhmContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymdhmContext) Ex_day() IEx_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_dayContext)
}

func (s *Ex_iso_date_time_ymdhmContext) EX_T() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_T, 0)
}

func (s *Ex_iso_date_time_ymdhmContext) Ex_hour() IEx_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_hourContext)
}

func (s *Ex_iso_date_time_ymdhmContext) EX_COLON() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_COLON, 0)
}

func (s *Ex_iso_date_time_ymdhmContext) Ex_minutes() IEx_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_minutesContext)
}

func (s *Ex_iso_date_time_ymdhmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymdhmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymdhmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ymdhm(s)
	}
}

func (s *Ex_iso_date_time_ymdhmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ymdhm(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ymdhm() (localctx IEx_iso_date_time_ymdhmContext) {
	localctx = NewEx_iso_date_time_ymdhmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromQLExtensionParserRULE_ex_iso_date_time_ymdhm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Ex_year()
	}
	{
		p.SetState(216)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Ex_month()
	}
	{
		p.SetState(218)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Ex_day()
	}
	{
		p.SetState(220)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Ex_hour()
	}
	{
		p.SetState(222)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Ex_minutes()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymdhContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymdhContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Ex_month() IEx_monthContext
	Ex_day() IEx_dayContext
	EX_T() antlr.TerminalNode
	Ex_hour() IEx_hourContext

	// IsEx_iso_date_time_ymdhContext differentiates from other interfaces.
	IsEx_iso_date_time_ymdhContext()
}

type Ex_iso_date_time_ymdhContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymdhContext() *Ex_iso_date_time_ymdhContext {
	var p = new(Ex_iso_date_time_ymdhContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdh
	return p
}

func InitEmptyEx_iso_date_time_ymdhContext(p *Ex_iso_date_time_ymdhContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdh
}

func (*Ex_iso_date_time_ymdhContext) IsEx_iso_date_time_ymdhContext() {}

func NewEx_iso_date_time_ymdhContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymdhContext {
	var p = new(Ex_iso_date_time_ymdhContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymdh

	return p
}

func (s *Ex_iso_date_time_ymdhContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymdhContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymdhContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserSUB)
}

func (s *Ex_iso_date_time_ymdhContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, i)
}

func (s *Ex_iso_date_time_ymdhContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymdhContext) Ex_day() IEx_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_dayContext)
}

func (s *Ex_iso_date_time_ymdhContext) EX_T() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_T, 0)
}

func (s *Ex_iso_date_time_ymdhContext) Ex_hour() IEx_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_hourContext)
}

func (s *Ex_iso_date_time_ymdhContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymdhContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymdhContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ymdh(s)
	}
}

func (s *Ex_iso_date_time_ymdhContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ymdh(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ymdh() (localctx IEx_iso_date_time_ymdhContext) {
	localctx = NewEx_iso_date_time_ymdhContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromQLExtensionParserRULE_ex_iso_date_time_ymdh)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Ex_year()
	}
	{
		p.SetState(226)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
		p.Ex_month()
	}
	{
		p.SetState(228)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Ex_day()
	}
	{
		p.SetState(230)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Ex_hour()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymdContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Ex_month() IEx_monthContext
	Ex_day() IEx_dayContext

	// IsEx_iso_date_time_ymdContext differentiates from other interfaces.
	IsEx_iso_date_time_ymdContext()
}

type Ex_iso_date_time_ymdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymdContext() *Ex_iso_date_time_ymdContext {
	var p = new(Ex_iso_date_time_ymdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymd
	return p
}

func InitEmptyEx_iso_date_time_ymdContext(p *Ex_iso_date_time_ymdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymd
}

func (*Ex_iso_date_time_ymdContext) IsEx_iso_date_time_ymdContext() {}

func NewEx_iso_date_time_ymdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymdContext {
	var p = new(Ex_iso_date_time_ymdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ymd

	return p
}

func (s *Ex_iso_date_time_ymdContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymdContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymdContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserSUB)
}

func (s *Ex_iso_date_time_ymdContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, i)
}

func (s *Ex_iso_date_time_ymdContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymdContext) Ex_day() IEx_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_dayContext)
}

func (s *Ex_iso_date_time_ymdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ymd(s)
	}
}

func (s *Ex_iso_date_time_ymdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ymd(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ymd() (localctx IEx_iso_date_time_ymdContext) {
	localctx = NewEx_iso_date_time_ymdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromQLExtensionParserRULE_ex_iso_date_time_ymd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Ex_year()
	}
	{
		p.SetState(234)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Ex_month()
	}
	{
		p.SetState(236)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Ex_day()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_ymContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_ymContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext
	SUB() antlr.TerminalNode
	Ex_month() IEx_monthContext

	// IsEx_iso_date_time_ymContext differentiates from other interfaces.
	IsEx_iso_date_time_ymContext()
}

type Ex_iso_date_time_ymContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_ymContext() *Ex_iso_date_time_ymContext {
	var p = new(Ex_iso_date_time_ymContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ym
	return p
}

func InitEmptyEx_iso_date_time_ymContext(p *Ex_iso_date_time_ymContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ym
}

func (*Ex_iso_date_time_ymContext) IsEx_iso_date_time_ymContext() {}

func NewEx_iso_date_time_ymContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_ymContext {
	var p = new(Ex_iso_date_time_ymContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_ym

	return p
}

func (s *Ex_iso_date_time_ymContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_ymContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_ymContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, 0)
}

func (s *Ex_iso_date_time_ymContext) Ex_month() IEx_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_monthContext)
}

func (s *Ex_iso_date_time_ymContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_ymContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_ymContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_ym(s)
	}
}

func (s *Ex_iso_date_time_ymContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_ym(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_ym() (localctx IEx_iso_date_time_ymContext) {
	localctx = NewEx_iso_date_time_ymContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromQLExtensionParserRULE_ex_iso_date_time_ym)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Ex_year()
	}
	{
		p.SetState(240)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Ex_month()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_iso_date_time_yContext is an interface to support dynamic dispatch.
type IEx_iso_date_time_yContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_year() IEx_yearContext

	// IsEx_iso_date_time_yContext differentiates from other interfaces.
	IsEx_iso_date_time_yContext()
}

type Ex_iso_date_time_yContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_iso_date_time_yContext() *Ex_iso_date_time_yContext {
	var p = new(Ex_iso_date_time_yContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_y
	return p
}

func InitEmptyEx_iso_date_time_yContext(p *Ex_iso_date_time_yContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_y
}

func (*Ex_iso_date_time_yContext) IsEx_iso_date_time_yContext() {}

func NewEx_iso_date_time_yContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_iso_date_time_yContext {
	var p = new(Ex_iso_date_time_yContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_iso_date_time_y

	return p
}

func (s *Ex_iso_date_time_yContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_iso_date_time_yContext) Ex_year() IEx_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_yearContext)
}

func (s *Ex_iso_date_time_yContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_iso_date_time_yContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_iso_date_time_yContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_iso_date_time_y(s)
	}
}

func (s *Ex_iso_date_time_yContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_iso_date_time_y(s)
	}
}

func (p *PromQLExtensionParser) Ex_iso_date_time_y() (localctx IEx_iso_date_time_yContext) {
	localctx = NewEx_iso_date_time_yContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromQLExtensionParserRULE_ex_iso_date_time_y)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Ex_year()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_yearContext is an interface to support dynamic dispatch.
type IEx_yearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_POSITIVE_INTEGER() antlr.TerminalNode

	// IsEx_yearContext differentiates from other interfaces.
	IsEx_yearContext()
}

type Ex_yearContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_yearContext() *Ex_yearContext {
	var p = new(Ex_yearContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_year
	return p
}

func InitEmptyEx_yearContext(p *Ex_yearContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_year
}

func (*Ex_yearContext) IsEx_yearContext() {}

func NewEx_yearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_yearContext {
	var p = new(Ex_yearContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_year

	return p
}

func (s *Ex_yearContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_yearContext) EX_POSITIVE_INTEGER() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_POSITIVE_INTEGER, 0)
}

func (s *Ex_yearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_yearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_yearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_year(s)
	}
}

func (s *Ex_yearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_year(s)
	}
}

func (p *PromQLExtensionParser) Ex_year() (localctx IEx_yearContext) {
	localctx = NewEx_yearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PromQLExtensionParserRULE_ex_year)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(PromQLExtensionParserEX_POSITIVE_INTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_monthContext is an interface to support dynamic dispatch.
type IEx_monthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TWO_DIGITS() antlr.TerminalNode

	// IsEx_monthContext differentiates from other interfaces.
	IsEx_monthContext()
}

type Ex_monthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_monthContext() *Ex_monthContext {
	var p = new(Ex_monthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_month
	return p
}

func InitEmptyEx_monthContext(p *Ex_monthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_month
}

func (*Ex_monthContext) IsEx_monthContext() {}

func NewEx_monthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_monthContext {
	var p = new(Ex_monthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_month

	return p
}

func (s *Ex_monthContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_monthContext) EX_TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TWO_DIGITS, 0)
}

func (s *Ex_monthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_monthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_monthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_month(s)
	}
}

func (s *Ex_monthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_month(s)
	}
}

func (p *PromQLExtensionParser) Ex_month() (localctx IEx_monthContext) {
	localctx = NewEx_monthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PromQLExtensionParserRULE_ex_month)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(PromQLExtensionParserEX_TWO_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_dayContext is an interface to support dynamic dispatch.
type IEx_dayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TWO_DIGITS() antlr.TerminalNode

	// IsEx_dayContext differentiates from other interfaces.
	IsEx_dayContext()
}

type Ex_dayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_dayContext() *Ex_dayContext {
	var p = new(Ex_dayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_day
	return p
}

func InitEmptyEx_dayContext(p *Ex_dayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_day
}

func (*Ex_dayContext) IsEx_dayContext() {}

func NewEx_dayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_dayContext {
	var p = new(Ex_dayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_day

	return p
}

func (s *Ex_dayContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_dayContext) EX_TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TWO_DIGITS, 0)
}

func (s *Ex_dayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_dayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_dayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_day(s)
	}
}

func (s *Ex_dayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_day(s)
	}
}

func (p *PromQLExtensionParser) Ex_day() (localctx IEx_dayContext) {
	localctx = NewEx_dayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PromQLExtensionParserRULE_ex_day)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(PromQLExtensionParserEX_TWO_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_hourContext is an interface to support dynamic dispatch.
type IEx_hourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TWO_DIGITS() antlr.TerminalNode

	// IsEx_hourContext differentiates from other interfaces.
	IsEx_hourContext()
}

type Ex_hourContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_hourContext() *Ex_hourContext {
	var p = new(Ex_hourContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_hour
	return p
}

func InitEmptyEx_hourContext(p *Ex_hourContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_hour
}

func (*Ex_hourContext) IsEx_hourContext() {}

func NewEx_hourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_hourContext {
	var p = new(Ex_hourContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_hour

	return p
}

func (s *Ex_hourContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_hourContext) EX_TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TWO_DIGITS, 0)
}

func (s *Ex_hourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_hourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_hourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_hour(s)
	}
}

func (s *Ex_hourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_hour(s)
	}
}

func (p *PromQLExtensionParser) Ex_hour() (localctx IEx_hourContext) {
	localctx = NewEx_hourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PromQLExtensionParserRULE_ex_hour)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(PromQLExtensionParserEX_TWO_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_minutesContext is an interface to support dynamic dispatch.
type IEx_minutesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TWO_DIGITS() antlr.TerminalNode

	// IsEx_minutesContext differentiates from other interfaces.
	IsEx_minutesContext()
}

type Ex_minutesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_minutesContext() *Ex_minutesContext {
	var p = new(Ex_minutesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_minutes
	return p
}

func InitEmptyEx_minutesContext(p *Ex_minutesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_minutes
}

func (*Ex_minutesContext) IsEx_minutesContext() {}

func NewEx_minutesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_minutesContext {
	var p = new(Ex_minutesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_minutes

	return p
}

func (s *Ex_minutesContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_minutesContext) EX_TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TWO_DIGITS, 0)
}

func (s *Ex_minutesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_minutesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_minutesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_minutes(s)
	}
}

func (s *Ex_minutesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_minutes(s)
	}
}

func (p *PromQLExtensionParser) Ex_minutes() (localctx IEx_minutesContext) {
	localctx = NewEx_minutesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PromQLExtensionParserRULE_ex_minutes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(PromQLExtensionParserEX_TWO_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_secondsContext is an interface to support dynamic dispatch.
type IEx_secondsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_TWO_DIGITS() antlr.TerminalNode

	// IsEx_secondsContext differentiates from other interfaces.
	IsEx_secondsContext()
}

type Ex_secondsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_secondsContext() *Ex_secondsContext {
	var p = new(Ex_secondsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_seconds
	return p
}

func InitEmptyEx_secondsContext(p *Ex_secondsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_seconds
}

func (*Ex_secondsContext) IsEx_secondsContext() {}

func NewEx_secondsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_secondsContext {
	var p = new(Ex_secondsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_seconds

	return p
}

func (s *Ex_secondsContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_secondsContext) EX_TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_TWO_DIGITS, 0)
}

func (s *Ex_secondsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_secondsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_secondsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_seconds(s)
	}
}

func (s *Ex_secondsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_seconds(s)
	}
}

func (p *PromQLExtensionParser) Ex_seconds() (localctx IEx_secondsContext) {
	localctx = NewEx_secondsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromQLExtensionParserRULE_ex_seconds)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(PromQLExtensionParserEX_TWO_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEx_frac_secContext is an interface to support dynamic dispatch.
type IEx_frac_secContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_DIGITS() antlr.TerminalNode

	// IsEx_frac_secContext differentiates from other interfaces.
	IsEx_frac_secContext()
}

type Ex_frac_secContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_frac_secContext() *Ex_frac_secContext {
	var p = new(Ex_frac_secContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_frac_sec
	return p
}

func InitEmptyEx_frac_secContext(p *Ex_frac_secContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_frac_sec
}

func (*Ex_frac_secContext) IsEx_frac_secContext() {}

func NewEx_frac_secContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_frac_secContext {
	var p = new(Ex_frac_secContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_frac_sec

	return p
}

func (s *Ex_frac_secContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_frac_secContext) EX_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_DIGITS, 0)
}

func (s *Ex_frac_secContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_frac_secContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_frac_secContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_frac_sec(s)
	}
}

func (s *Ex_frac_secContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_frac_sec(s)
	}
}

func (p *PromQLExtensionParser) Ex_frac_sec() (localctx IEx_frac_secContext) {
	localctx = NewEx_frac_secContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromQLExtensionParserRULE_ex_frac_sec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(PromQLExtensionParserEX_DIGITS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VectorOperation() IVectorOperationContext
	EOF() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) VectorOperation() IVectorOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorOperationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PromQLExtensionParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromQLExtensionParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.vectorOperation(0)
	}
	{
		p.SetState(260)
		p.Match(PromQLExtensionParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorOperationContext is an interface to support dynamic dispatch.
type IVectorOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UnaryOp() IUnaryOpContext
	AllVectorOperation() []IVectorOperationContext
	VectorOperation(i int) IVectorOperationContext
	Vector() IVectorContext
	PowOp() IPowOpContext
	MultOp() IMultOpContext
	AddOp() IAddOpContext
	CompareOp() ICompareOpContext
	AndUnlessOp() IAndUnlessOpContext
	OrOp() IOrOpContext
	VectorMatchOp() IVectorMatchOpContext
	AT() antlr.TerminalNode
	SubqueryOp() ISubqueryOpContext

	// IsVectorOperationContext differentiates from other interfaces.
	IsVectorOperationContext()
}

type VectorOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorOperationContext() *VectorOperationContext {
	var p = new(VectorOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vectorOperation
	return p
}

func InitEmptyVectorOperationContext(p *VectorOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vectorOperation
}

func (*VectorOperationContext) IsVectorOperationContext() {}

func NewVectorOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorOperationContext {
	var p = new(VectorOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_vectorOperation

	return p
}

func (s *VectorOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorOperationContext) UnaryOp() IUnaryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *VectorOperationContext) AllVectorOperation() []IVectorOperationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVectorOperationContext); ok {
			len++
		}
	}

	tst := make([]IVectorOperationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVectorOperationContext); ok {
			tst[i] = t.(IVectorOperationContext)
			i++
		}
	}

	return tst
}

func (s *VectorOperationContext) VectorOperation(i int) IVectorOperationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorOperationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorOperationContext)
}

func (s *VectorOperationContext) Vector() IVectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *VectorOperationContext) PowOp() IPowOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPowOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPowOpContext)
}

func (s *VectorOperationContext) MultOp() IMultOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultOpContext)
}

func (s *VectorOperationContext) AddOp() IAddOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}

func (s *VectorOperationContext) CompareOp() ICompareOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareOpContext)
}

func (s *VectorOperationContext) AndUnlessOp() IAndUnlessOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndUnlessOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndUnlessOpContext)
}

func (s *VectorOperationContext) OrOp() IOrOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrOpContext)
}

func (s *VectorOperationContext) VectorMatchOp() IVectorMatchOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorMatchOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorMatchOpContext)
}

func (s *VectorOperationContext) AT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserAT, 0)
}

func (s *VectorOperationContext) SubqueryOp() ISubqueryOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryOpContext)
}

func (s *VectorOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterVectorOperation(s)
	}
}

func (s *VectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitVectorOperation(s)
	}
}

func (p *PromQLExtensionParser) VectorOperation() (localctx IVectorOperationContext) {
	return p.vectorOperation(0)
}

func (p *PromQLExtensionParser) vectorOperation(_p int) (localctx IVectorOperationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewVectorOperationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVectorOperationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, PromQLExtensionParserRULE_vectorOperation, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserADD, PromQLExtensionParserSUB:
		{
			p.SetState(263)
			p.UnaryOp()
		}
		{
			p.SetState(264)
			p.vectorOperation(9)
		}

	case PromQLExtensionParserNUMBER, PromQLExtensionParserSTRING, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION, PromQLExtensionParserLEFT_BRACE, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserMETRIC_NAME:
		{
			p.SetState(266)
			p.Vector()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(270)
					p.PowOp()
				}
				{
					p.SetState(271)
					p.vectorOperation(11)
				}

			case 2:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(273)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(274)
					p.MultOp()
				}
				{
					p.SetState(275)
					p.vectorOperation(9)
				}

			case 3:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(277)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(278)
					p.AddOp()
				}
				{
					p.SetState(279)
					p.vectorOperation(8)
				}

			case 4:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(282)
					p.CompareOp()
				}
				{
					p.SetState(283)
					p.vectorOperation(7)
				}

			case 5:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(286)
					p.AndUnlessOp()
				}
				{
					p.SetState(287)
					p.vectorOperation(6)
				}

			case 6:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(290)
					p.OrOp()
				}
				{
					p.SetState(291)
					p.vectorOperation(5)
				}

			case 7:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(294)
					p.VectorMatchOp()
				}
				{
					p.SetState(295)
					p.vectorOperation(4)
				}

			case 8:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(297)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(298)
					p.Match(PromQLExtensionParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(299)
					p.vectorOperation(3)
				}

			case 9:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(301)
					p.SubqueryOp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_unaryOp
	return p
}

func InitEmptyUnaryOpContext(p *UnaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_unaryOp
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserADD, 0)
}

func (s *UnaryOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (p *PromQLExtensionParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromQLExtensionParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserADD || _la == PromQLExtensionParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPowOpContext is an interface to support dynamic dispatch.
type IPowOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POW() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsPowOpContext differentiates from other interfaces.
	IsPowOpContext()
}

type PowOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowOpContext() *PowOpContext {
	var p = new(PowOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_powOp
	return p
}

func InitEmptyPowOpContext(p *PowOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_powOp
}

func (*PowOpContext) IsPowOpContext() {}

func NewPowOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowOpContext {
	var p = new(PowOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_powOp

	return p
}

func (s *PowOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PowOpContext) POW() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserPOW, 0)
}

func (s *PowOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *PowOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterPowOp(s)
	}
}

func (s *PowOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitPowOp(s)
	}
}

func (p *PromQLExtensionParser) PowOp() (localctx IPowOpContext) {
	localctx = NewPowOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromQLExtensionParserRULE_powOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(PromQLExtensionParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(310)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultOpContext is an interface to support dynamic dispatch.
type IMultOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsMultOpContext differentiates from other interfaces.
	IsMultOpContext()
}

type MultOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultOpContext() *MultOpContext {
	var p = new(MultOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_multOp
	return p
}

func InitEmptyMultOpContext(p *MultOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_multOp
}

func (*MultOpContext) IsMultOpContext() {}

func NewMultOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOpContext {
	var p = new(MultOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_multOp

	return p
}

func (s *MultOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOpContext) MULT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserMULT, 0)
}

func (s *MultOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserDIV, 0)
}

func (s *MultOpContext) MOD() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserMOD, 0)
}

func (s *MultOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *MultOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterMultOp(s)
	}
}

func (s *MultOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitMultOp(s)
	}
}

func (p *PromQLExtensionParser) MultOp() (localctx IMultOpContext) {
	localctx = NewMultOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromQLExtensionParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&114688) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(314)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserADD, 0)
}

func (s *AddOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUB, 0)
}

func (s *AddOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (p *PromQLExtensionParser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromQLExtensionParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserADD || _la == PromQLExtensionParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(318)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICompareOpContext is an interface to support dynamic dispatch.
type ICompareOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	GE() antlr.TerminalNode
	LE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsCompareOpContext differentiates from other interfaces.
	IsCompareOpContext()
}

type CompareOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOpContext() *CompareOpContext {
	var p = new(CompareOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_compareOp
	return p
}

func InitEmptyCompareOpContext(p *CompareOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_compareOp
}

func (*CompareOpContext) IsCompareOpContext() {}

func NewCompareOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOpContext {
	var p = new(CompareOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_compareOp

	return p
}

func (s *CompareOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOpContext) DEQ() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserDEQ, 0)
}

func (s *CompareOpContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserNE, 0)
}

func (s *CompareOpContext) GT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGT, 0)
}

func (s *CompareOpContext) LT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLT, 0)
}

func (s *CompareOpContext) GE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGE, 0)
}

func (s *CompareOpContext) LE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLE, 0)
}

func (s *CompareOpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserBOOL, 0)
}

func (s *CompareOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *CompareOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterCompareOp(s)
	}
}

func (s *CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitCompareOp(s)
	}
}

func (p *PromQLExtensionParser) CompareOp() (localctx ICompareOpContext) {
	localctx = NewCompareOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PromQLExtensionParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264241152) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserBOOL {
		{
			p.SetState(322)
			p.Match(PromQLExtensionParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(325)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndUnlessOpContext is an interface to support dynamic dispatch.
type IAndUnlessOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	UNLESS() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsAndUnlessOpContext differentiates from other interfaces.
	IsAndUnlessOpContext()
}

type AndUnlessOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndUnlessOpContext() *AndUnlessOpContext {
	var p = new(AndUnlessOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_andUnlessOp
	return p
}

func InitEmptyAndUnlessOpContext(p *AndUnlessOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_andUnlessOp
}

func (*AndUnlessOpContext) IsAndUnlessOpContext() {}

func NewAndUnlessOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndUnlessOpContext {
	var p = new(AndUnlessOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_andUnlessOp

	return p
}

func (s *AndUnlessOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AndUnlessOpContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserAND, 0)
}

func (s *AndUnlessOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserUNLESS, 0)
}

func (s *AndUnlessOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *AndUnlessOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndUnlessOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndUnlessOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterAndUnlessOp(s)
	}
}

func (s *AndUnlessOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitAndUnlessOp(s)
	}
}

func (p *PromQLExtensionParser) AndUnlessOp() (localctx IAndUnlessOpContext) {
	localctx = NewAndUnlessOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PromQLExtensionParserRULE_andUnlessOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserAND || _la == PromQLExtensionParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(329)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrOpContext is an interface to support dynamic dispatch.
type IOrOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsOrOpContext differentiates from other interfaces.
	IsOrOpContext()
}

type OrOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrOpContext() *OrOpContext {
	var p = new(OrOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_orOp
	return p
}

func InitEmptyOrOpContext(p *OrOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_orOp
}

func (*OrOpContext) IsOrOpContext() {}

func NewOrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrOpContext {
	var p = new(OrOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_orOp

	return p
}

func (s *OrOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OrOpContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserOR, 0)
}

func (s *OrOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *OrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterOrOp(s)
	}
}

func (s *OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitOrOp(s)
	}
}

func (p *PromQLExtensionParser) OrOp() (localctx IOrOpContext) {
	localctx = NewOrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PromQLExtensionParserRULE_orOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(PromQLExtensionParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(333)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorMatchOpContext is an interface to support dynamic dispatch.
type IVectorMatchOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ON() antlr.TerminalNode
	UNLESS() antlr.TerminalNode
	Grouping() IGroupingContext

	// IsVectorMatchOpContext differentiates from other interfaces.
	IsVectorMatchOpContext()
}

type VectorMatchOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorMatchOpContext() *VectorMatchOpContext {
	var p = new(VectorMatchOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vectorMatchOp
	return p
}

func InitEmptyVectorMatchOpContext(p *VectorMatchOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vectorMatchOp
}

func (*VectorMatchOpContext) IsVectorMatchOpContext() {}

func NewVectorMatchOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorMatchOpContext {
	var p = new(VectorMatchOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_vectorMatchOp

	return p
}

func (s *VectorMatchOpContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorMatchOpContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserON, 0)
}

func (s *VectorMatchOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserUNLESS, 0)
}

func (s *VectorMatchOpContext) Grouping() IGroupingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupingContext)
}

func (s *VectorMatchOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorMatchOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorMatchOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterVectorMatchOp(s)
	}
}

func (s *VectorMatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitVectorMatchOp(s)
	}
}

func (p *PromQLExtensionParser) VectorMatchOp() (localctx IVectorMatchOpContext) {
	localctx = NewVectorMatchOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PromQLExtensionParserRULE_vectorMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserUNLESS || _la == PromQLExtensionParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(337)
			p.Grouping()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISubqueryOpContext is an interface to support dynamic dispatch.
type ISubqueryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUBQUERY_RANGE() antlr.TerminalNode
	OffsetOp() IOffsetOpContext

	// IsSubqueryOpContext differentiates from other interfaces.
	IsSubqueryOpContext()
}

type SubqueryOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubqueryOpContext() *SubqueryOpContext {
	var p = new(SubqueryOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_subqueryOp
	return p
}

func InitEmptySubqueryOpContext(p *SubqueryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_subqueryOp
}

func (*SubqueryOpContext) IsSubqueryOpContext() {}

func NewSubqueryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryOpContext {
	var p = new(SubqueryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_subqueryOp

	return p
}

func (s *SubqueryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryOpContext) SUBQUERY_RANGE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSUBQUERY_RANGE, 0)
}

func (s *SubqueryOpContext) OffsetOp() IOffsetOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetOpContext)
}

func (s *SubqueryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubqueryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterSubqueryOp(s)
	}
}

func (s *SubqueryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitSubqueryOp(s)
	}
}

func (p *PromQLExtensionParser) SubqueryOp() (localctx ISubqueryOpContext) {
	localctx = NewSubqueryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PromQLExtensionParserRULE_subqueryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(PromQLExtensionParserSUBQUERY_RANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(341)
			p.OffsetOp()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOffsetOpContext is an interface to support dynamic dispatch.
type IOffsetOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET() antlr.TerminalNode
	DURATION() antlr.TerminalNode

	// IsOffsetOpContext differentiates from other interfaces.
	IsOffsetOpContext()
}

type OffsetOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetOpContext() *OffsetOpContext {
	var p = new(OffsetOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_offsetOp
	return p
}

func InitEmptyOffsetOpContext(p *OffsetOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_offsetOp
}

func (*OffsetOpContext) IsOffsetOpContext() {}

func NewOffsetOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetOpContext {
	var p = new(OffsetOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_offsetOp

	return p
}

func (s *OffsetOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetOpContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserOFFSET, 0)
}

func (s *OffsetOpContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserDURATION, 0)
}

func (s *OffsetOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterOffsetOp(s)
	}
}

func (s *OffsetOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitOffsetOp(s)
	}
}

func (p *PromQLExtensionParser) OffsetOp() (localctx IOffsetOpContext) {
	localctx = NewOffsetOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PromQLExtensionParserRULE_offsetOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(PromQLExtensionParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(PromQLExtensionParserDURATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_() IFunction_Context
	Aggregation() IAggregationContext
	InstantSelector() IInstantSelectorContext
	MatrixSelector() IMatrixSelectorContext
	Offset() IOffsetContext
	Literal() ILiteralContext
	Parens() IParensContext

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vector
	return p
}

func InitEmptyVectorContext(p *VectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_vector
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) Function_() IFunction_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *VectorContext) Aggregation() IAggregationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregationContext)
}

func (s *VectorContext) InstantSelector() IInstantSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstantSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstantSelectorContext)
}

func (s *VectorContext) MatrixSelector() IMatrixSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixSelectorContext)
}

func (s *VectorContext) Offset() IOffsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetContext)
}

func (s *VectorContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *VectorContext) Parens() IParensContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParensContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParensContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *PromQLExtensionParser) Vector() (localctx IVectorContext) {
	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PromQLExtensionParserRULE_vector)
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(349)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(350)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(351)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(352)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(353)
			p.Parens()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParensContext is an interface to support dynamic dispatch.
type IParensContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PAREN() antlr.TerminalNode
	VectorOperation() IVectorOperationContext
	RIGHT_PAREN() antlr.TerminalNode

	// IsParensContext differentiates from other interfaces.
	IsParensContext()
}

type ParensContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParensContext() *ParensContext {
	var p = new(ParensContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parens
	return p
}

func InitEmptyParensContext(p *ParensContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parens
}

func (*ParensContext) IsParensContext() {}

func NewParensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParensContext {
	var p = new(ParensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_parens

	return p
}

func (s *ParensContext) GetParser() antlr.Parser { return s.parser }

func (s *ParensContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *ParensContext) VectorOperation() IVectorOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorOperationContext)
}

func (s *ParensContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitParens(s)
	}
}

func (p *PromQLExtensionParser) Parens() (localctx IParensContext) {
	localctx = NewParensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PromQLExtensionParserRULE_parens)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.vectorOperation(0)
	}
	{
		p.SetState(358)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstantSelectorContext is an interface to support dynamic dispatch.
type IInstantSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	METRIC_NAME() antlr.TerminalNode
	LEFT_BRACE() antlr.TerminalNode
	RIGHT_BRACE() antlr.TerminalNode
	LabelMatcherList() ILabelMatcherListContext

	// IsInstantSelectorContext differentiates from other interfaces.
	IsInstantSelectorContext()
}

type InstantSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstantSelectorContext() *InstantSelectorContext {
	var p = new(InstantSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_instantSelector
	return p
}

func InitEmptyInstantSelectorContext(p *InstantSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_instantSelector
}

func (*InstantSelectorContext) IsInstantSelectorContext() {}

func NewInstantSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstantSelectorContext {
	var p = new(InstantSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_instantSelector

	return p
}

func (s *InstantSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *InstantSelectorContext) METRIC_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserMETRIC_NAME, 0)
}

func (s *InstantSelectorContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_BRACE, 0)
}

func (s *InstantSelectorContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_BRACE, 0)
}

func (s *InstantSelectorContext) LabelMatcherList() ILabelMatcherListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelMatcherListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelMatcherListContext)
}

func (s *InstantSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstantSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstantSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterInstantSelector(s)
	}
}

func (s *InstantSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitInstantSelector(s)
	}
}

func (p *PromQLExtensionParser) InstantSelector() (localctx IInstantSelectorContext) {
	localctx = NewInstantSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PromQLExtensionParserRULE_instantSelector)
	var _la int

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(360)
			p.Match(PromQLExtensionParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(361)
				p.Match(PromQLExtensionParserLEFT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6756497880776704) != 0 {
				{
					p.SetState(362)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(365)
				p.Match(PromQLExtensionParserRIGHT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case PromQLExtensionParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(PromQLExtensionParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.LabelMatcherList()
		}
		{
			p.SetState(370)
			p.Match(PromQLExtensionParserRIGHT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelMatcherContext is an interface to support dynamic dispatch.
type ILabelMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabelName() ILabelNameContext
	LabelMatcherOperator() ILabelMatcherOperatorContext
	STRING() antlr.TerminalNode

	// IsLabelMatcherContext differentiates from other interfaces.
	IsLabelMatcherContext()
}

type LabelMatcherContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelMatcherContext() *LabelMatcherContext {
	var p = new(LabelMatcherContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcher
	return p
}

func InitEmptyLabelMatcherContext(p *LabelMatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcher
}

func (*LabelMatcherContext) IsLabelMatcherContext() {}

func NewLabelMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherContext {
	var p = new(LabelMatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcher

	return p
}

func (s *LabelMatcherContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherContext) LabelName() ILabelNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *LabelMatcherContext) LabelMatcherOperator() ILabelMatcherOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelMatcherOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelMatcherOperatorContext)
}

func (s *LabelMatcherContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSTRING, 0)
}

func (s *LabelMatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLabelMatcher(s)
	}
}

func (p *PromQLExtensionParser) LabelMatcher() (localctx ILabelMatcherContext) {
	localctx = NewLabelMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PromQLExtensionParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.LabelName()
	}
	{
		p.SetState(375)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(376)
		p.Match(PromQLExtensionParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelMatcherOperatorContext is an interface to support dynamic dispatch.
type ILabelMatcherOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	RE() antlr.TerminalNode
	NRE() antlr.TerminalNode

	// IsLabelMatcherOperatorContext differentiates from other interfaces.
	IsLabelMatcherOperatorContext()
}

type LabelMatcherOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelMatcherOperatorContext() *LabelMatcherOperatorContext {
	var p = new(LabelMatcherOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherOperator
	return p
}

func InitEmptyLabelMatcherOperatorContext(p *LabelMatcherOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherOperator
}

func (*LabelMatcherOperatorContext) IsLabelMatcherOperatorContext() {}

func NewLabelMatcherOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherOperatorContext {
	var p = new(LabelMatcherOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherOperator

	return p
}

func (s *LabelMatcherOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEQ, 0)
}

func (s *LabelMatcherOperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserNE, 0)
}

func (s *LabelMatcherOperatorContext) RE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRE, 0)
}

func (s *LabelMatcherOperatorContext) NRE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserNRE, 0)
}

func (s *LabelMatcherOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLabelMatcherOperator(s)
	}
}

func (s *LabelMatcherOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLabelMatcherOperator(s)
	}
}

func (p *PromQLExtensionParser) LabelMatcherOperator() (localctx ILabelMatcherOperatorContext) {
	localctx = NewLabelMatcherOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PromQLExtensionParserRULE_labelMatcherOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&815792128) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelMatcherListContext is an interface to support dynamic dispatch.
type ILabelMatcherListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabelMatcher() []ILabelMatcherContext
	LabelMatcher(i int) ILabelMatcherContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLabelMatcherListContext differentiates from other interfaces.
	IsLabelMatcherListContext()
}

type LabelMatcherListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelMatcherListContext() *LabelMatcherListContext {
	var p = new(LabelMatcherListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherList
	return p
}

func InitEmptyLabelMatcherListContext(p *LabelMatcherListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherList
}

func (*LabelMatcherListContext) IsLabelMatcherListContext() {}

func NewLabelMatcherListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherListContext {
	var p = new(LabelMatcherListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_labelMatcherList

	return p
}

func (s *LabelMatcherListContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherListContext) AllLabelMatcher() []ILabelMatcherContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelMatcherContext); ok {
			len++
		}
	}

	tst := make([]ILabelMatcherContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelMatcherContext); ok {
			tst[i] = t.(ILabelMatcherContext)
			i++
		}
	}

	return tst
}

func (s *LabelMatcherListContext) LabelMatcher(i int) ILabelMatcherContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelMatcherContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelMatcherContext)
}

func (s *LabelMatcherListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserCOMMA)
}

func (s *LabelMatcherListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserCOMMA, i)
}

func (s *LabelMatcherListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLabelMatcherList(s)
	}
}

func (s *LabelMatcherListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLabelMatcherList(s)
	}
}

func (p *PromQLExtensionParser) LabelMatcherList() (localctx ILabelMatcherListContext) {
	localctx = NewLabelMatcherListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PromQLExtensionParserRULE_labelMatcherList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.LabelMatcher()
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(382)
				p.LabelMatcher()
			}

		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserCOMMA {
		{
			p.SetState(388)
			p.Match(PromQLExtensionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrixSelectorContext is an interface to support dynamic dispatch.
type IMatrixSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InstantSelector() IInstantSelectorContext
	TIME_RANGE() antlr.TerminalNode

	// IsMatrixSelectorContext differentiates from other interfaces.
	IsMatrixSelectorContext()
}

type MatrixSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrixSelectorContext() *MatrixSelectorContext {
	var p = new(MatrixSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_matrixSelector
	return p
}

func InitEmptyMatrixSelectorContext(p *MatrixSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_matrixSelector
}

func (*MatrixSelectorContext) IsMatrixSelectorContext() {}

func NewMatrixSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixSelectorContext {
	var p = new(MatrixSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_matrixSelector

	return p
}

func (s *MatrixSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *MatrixSelectorContext) InstantSelector() IInstantSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstantSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstantSelectorContext)
}

func (s *MatrixSelectorContext) TIME_RANGE() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserTIME_RANGE, 0)
}

func (s *MatrixSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterMatrixSelector(s)
	}
}

func (s *MatrixSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitMatrixSelector(s)
	}
}

func (p *PromQLExtensionParser) MatrixSelector() (localctx IMatrixSelectorContext) {
	localctx = NewMatrixSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PromQLExtensionParserRULE_matrixSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.InstantSelector()
	}
	{
		p.SetState(392)
		p.Match(PromQLExtensionParserTIME_RANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOffsetContext is an interface to support dynamic dispatch.
type IOffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	InstantSelector() IInstantSelectorContext
	OFFSET() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	MatrixSelector() IMatrixSelectorContext

	// IsOffsetContext differentiates from other interfaces.
	IsOffsetContext()
}

type OffsetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetContext() *OffsetContext {
	var p = new(OffsetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_offset
	return p
}

func InitEmptyOffsetContext(p *OffsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_offset
}

func (*OffsetContext) IsOffsetContext() {}

func NewOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetContext {
	var p = new(OffsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_offset

	return p
}

func (s *OffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetContext) InstantSelector() IInstantSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstantSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstantSelectorContext)
}

func (s *OffsetContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserOFFSET, 0)
}

func (s *OffsetContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserDURATION, 0)
}

func (s *OffsetContext) MatrixSelector() IMatrixSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrixSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrixSelectorContext)
}

func (s *OffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterOffset(s)
	}
}

func (s *OffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitOffset(s)
	}
}

func (p *PromQLExtensionParser) Offset() (localctx IOffsetContext) {
	localctx = NewOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PromQLExtensionParserRULE_offset)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.InstantSelector()
		}
		{
			p.SetState(395)
			p.Match(PromQLExtensionParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(PromQLExtensionParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.MatrixSelector()
		}
		{
			p.SetState(399)
			p.Match(PromQLExtensionParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Match(PromQLExtensionParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_function_
	return p
}

func InitEmptyFunction_Context(p *Function_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_function_
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserFUNCTION, 0)
}

func (s *Function_Context) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *Function_Context) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *Function_Context) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *Function_Context) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Function_Context) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserCOMMA)
}

func (s *Function_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserCOMMA, i)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *PromQLExtensionParser) Function_() (localctx IFunction_Context) {
	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, PromQLExtensionParserRULE_function_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(404)
		p.Match(PromQLExtensionParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2258122005560320) != 0 {
		{
			p.SetState(406)
			p.Parameter()
		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(407)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(408)
				p.Parameter()
			}

			p.SetState(413)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(416)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	VectorOperation() IVectorOperationContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ParameterContext) VectorOperation() IVectorOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorOperationContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *PromQLExtensionParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, PromQLExtensionParserRULE_parameter)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)
			p.vectorOperation(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *ParameterListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *PromQLExtensionParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, PromQLExtensionParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2258122005560320) != 0 {
		{
			p.SetState(423)
			p.Parameter()
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(424)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(425)
				p.Parameter()
			}

			p.SetState(430)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(433)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregationContext is an interface to support dynamic dispatch.
type IAggregationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AGGREGATION_OPERATOR() antlr.TerminalNode
	ParameterList() IParameterListContext
	By() IByContext
	Without() IWithoutContext

	// IsAggregationContext differentiates from other interfaces.
	IsAggregationContext()
}

type AggregationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregationContext() *AggregationContext {
	var p = new(AggregationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserAGGREGATION_OPERATOR, 0)
}

func (s *AggregationContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *AggregationContext) By() IByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IByContext)
}

func (s *AggregationContext) Without() IWithoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWithoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWithoutContext)
}

func (s *AggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (p *PromQLExtensionParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, PromQLExtensionParserRULE_aggregation)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExtensionParserBY:
			{
				p.SetState(438)
				p.By()
			}

		case PromQLExtensionParserWITHOUT:
			{
				p.SetState(439)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(442)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.ParameterList()
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExtensionParserBY:
			{
				p.SetState(446)
				p.By()
			}

		case PromQLExtensionParserWITHOUT:
			{
				p.SetState(447)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IByContext is an interface to support dynamic dispatch.
type IByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BY() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsByContext differentiates from other interfaces.
	IsByContext()
}

type ByContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByContext() *ByContext {
	var p = new(ByContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_by
	return p
}

func InitEmptyByContext(p *ByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_by
}

func (*ByContext) IsByContext() {}

func NewByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByContext {
	var p = new(ByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_by

	return p
}

func (s *ByContext) GetParser() antlr.Parser { return s.parser }

func (s *ByContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserBY, 0)
}

func (s *ByContext) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *ByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterBy(s)
	}
}

func (s *ByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitBy(s)
	}
}

func (p *PromQLExtensionParser) By() (localctx IByContext) {
	localctx = NewByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, PromQLExtensionParserRULE_by)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(PromQLExtensionParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.LabelNameList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWithoutContext is an interface to support dynamic dispatch.
type IWithoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WITHOUT() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsWithoutContext differentiates from other interfaces.
	IsWithoutContext()
}

type WithoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithoutContext() *WithoutContext {
	var p = new(WithoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_without
	return p
}

func InitEmptyWithoutContext(p *WithoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_without
}

func (*WithoutContext) IsWithoutContext() {}

func NewWithoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithoutContext {
	var p = new(WithoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_without

	return p
}

func (s *WithoutContext) GetParser() antlr.Parser { return s.parser }

func (s *WithoutContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserWITHOUT, 0)
}

func (s *WithoutContext) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *WithoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterWithout(s)
	}
}

func (s *WithoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitWithout(s)
	}
}

func (p *PromQLExtensionParser) Without() (localctx IWithoutContext) {
	localctx = NewWithoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, PromQLExtensionParserRULE_without)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(PromQLExtensionParserWITHOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(456)
		p.LabelNameList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupingContext is an interface to support dynamic dispatch.
type IGroupingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	On_() IOn_Context
	Ignoring() IIgnoringContext
	GroupLeft() IGroupLeftContext
	GroupRight() IGroupRightContext

	// IsGroupingContext differentiates from other interfaces.
	IsGroupingContext()
}

type GroupingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupingContext() *GroupingContext {
	var p = new(GroupingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_grouping
	return p
}

func InitEmptyGroupingContext(p *GroupingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_grouping
}

func (*GroupingContext) IsGroupingContext() {}

func NewGroupingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupingContext {
	var p = new(GroupingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_grouping

	return p
}

func (s *GroupingContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupingContext) On_() IOn_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOn_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOn_Context)
}

func (s *GroupingContext) Ignoring() IIgnoringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIgnoringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIgnoringContext)
}

func (s *GroupingContext) GroupLeft() IGroupLeftContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupLeftContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupLeftContext)
}

func (s *GroupingContext) GroupRight() IGroupRightContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupRightContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupRightContext)
}

func (s *GroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (p *PromQLExtensionParser) Grouping() (localctx IGroupingContext) {
	localctx = NewGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, PromQLExtensionParserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserON:
		{
			p.SetState(458)
			p.On_()
		}

	case PromQLExtensionParserIGNORING:
		{
			p.SetState(459)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserGROUP_LEFT:
		{
			p.SetState(462)
			p.GroupLeft()
		}

	case PromQLExtensionParserGROUP_RIGHT:
		{
			p.SetState(463)
			p.GroupRight()
		}

	case PromQLExtensionParserNUMBER, PromQLExtensionParserSTRING, PromQLExtensionParserADD, PromQLExtensionParserSUB, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION, PromQLExtensionParserLEFT_BRACE, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserMETRIC_NAME:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOn_Context is an interface to support dynamic dispatch.
type IOn_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ON() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsOn_Context differentiates from other interfaces.
	IsOn_Context()
}

type On_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOn_Context() *On_Context {
	var p = new(On_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_on_
	return p
}

func InitEmptyOn_Context(p *On_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_on_
}

func (*On_Context) IsOn_Context() {}

func NewOn_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *On_Context {
	var p = new(On_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_on_

	return p
}

func (s *On_Context) GetParser() antlr.Parser { return s.parser }

func (s *On_Context) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserON, 0)
}

func (s *On_Context) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *On_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *On_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *On_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterOn_(s)
	}
}

func (s *On_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitOn_(s)
	}
}

func (p *PromQLExtensionParser) On_() (localctx IOn_Context) {
	localctx = NewOn_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, PromQLExtensionParserRULE_on_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(PromQLExtensionParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(467)
		p.LabelNameList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIgnoringContext is an interface to support dynamic dispatch.
type IIgnoringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IGNORING() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsIgnoringContext differentiates from other interfaces.
	IsIgnoringContext()
}

type IgnoringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIgnoringContext() *IgnoringContext {
	var p = new(IgnoringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ignoring
	return p
}

func InitEmptyIgnoringContext(p *IgnoringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ignoring
}

func (*IgnoringContext) IsIgnoringContext() {}

func NewIgnoringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoringContext {
	var p = new(IgnoringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ignoring

	return p
}

func (s *IgnoringContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoringContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserIGNORING, 0)
}

func (s *IgnoringContext) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *IgnoringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgnoringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgnoringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterIgnoring(s)
	}
}

func (s *IgnoringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitIgnoring(s)
	}
}

func (p *PromQLExtensionParser) Ignoring() (localctx IIgnoringContext) {
	localctx = NewIgnoringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, PromQLExtensionParserRULE_ignoring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Match(PromQLExtensionParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(470)
		p.LabelNameList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupLeftContext is an interface to support dynamic dispatch.
type IGroupLeftContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP_LEFT() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsGroupLeftContext differentiates from other interfaces.
	IsGroupLeftContext()
}

type GroupLeftContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupLeftContext() *GroupLeftContext {
	var p = new(GroupLeftContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_groupLeft
	return p
}

func InitEmptyGroupLeftContext(p *GroupLeftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_groupLeft
}

func (*GroupLeftContext) IsGroupLeftContext() {}

func NewGroupLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupLeftContext {
	var p = new(GroupLeftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_groupLeft

	return p
}

func (s *GroupLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupLeftContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGROUP_LEFT, 0)
}

func (s *GroupLeftContext) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *GroupLeftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupLeftContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupLeftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterGroupLeft(s)
	}
}

func (s *GroupLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitGroupLeft(s)
	}
}

func (p *PromQLExtensionParser) GroupLeft() (localctx IGroupLeftContext) {
	localctx = NewGroupLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, PromQLExtensionParserRULE_groupLeft)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.Match(PromQLExtensionParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(473)
			p.LabelNameList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupRightContext is an interface to support dynamic dispatch.
type IGroupRightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP_RIGHT() antlr.TerminalNode
	LabelNameList() ILabelNameListContext

	// IsGroupRightContext differentiates from other interfaces.
	IsGroupRightContext()
}

type GroupRightContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupRightContext() *GroupRightContext {
	var p = new(GroupRightContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_groupRight
	return p
}

func InitEmptyGroupRightContext(p *GroupRightContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_groupRight
}

func (*GroupRightContext) IsGroupRightContext() {}

func NewGroupRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupRightContext {
	var p = new(GroupRightContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_groupRight

	return p
}

func (s *GroupRightContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupRightContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGROUP_RIGHT, 0)
}

func (s *GroupRightContext) LabelNameList() ILabelNameListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameListContext)
}

func (s *GroupRightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupRightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupRightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterGroupRight(s)
	}
}

func (s *GroupRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitGroupRight(s)
	}
}

func (p *PromQLExtensionParser) GroupRight() (localctx IGroupRightContext) {
	localctx = NewGroupRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, PromQLExtensionParserRULE_groupRight)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)
		p.Match(PromQLExtensionParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(477)
			p.LabelNameList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelNameContext is an interface to support dynamic dispatch.
type ILabelNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() IKeywordContext
	METRIC_NAME() antlr.TerminalNode
	LABEL_NAME() antlr.TerminalNode

	// IsLabelNameContext differentiates from other interfaces.
	IsLabelNameContext()
}

type LabelNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelNameContext() *LabelNameContext {
	var p = new(LabelNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelName
	return p
}

func InitEmptyLabelNameContext(p *LabelNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelName
}

func (*LabelNameContext) IsLabelNameContext() {}

func NewLabelNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameContext {
	var p = new(LabelNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_labelName

	return p
}

func (s *LabelNameContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelNameContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *LabelNameContext) METRIC_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserMETRIC_NAME, 0)
}

func (s *LabelNameContext) LABEL_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLABEL_NAME, 0)
}

func (s *LabelNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLabelName(s)
	}
}

func (s *LabelNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLabelName(s)
	}
}

func (p *PromQLExtensionParser) LabelName() (localctx ILabelNameContext) {
	localctx = NewLabelNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, PromQLExtensionParserRULE_labelName)
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserAND, PromQLExtensionParserOR, PromQLExtensionParserUNLESS, PromQLExtensionParserBY, PromQLExtensionParserWITHOUT, PromQLExtensionParserON, PromQLExtensionParserIGNORING, PromQLExtensionParserGROUP_LEFT, PromQLExtensionParserGROUP_RIGHT, PromQLExtensionParserOFFSET, PromQLExtensionParserBOOL, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			p.Keyword()
		}

	case PromQLExtensionParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(481)
			p.Match(PromQLExtensionParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExtensionParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(482)
			p.Match(PromQLExtensionParserLABEL_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelNameListContext is an interface to support dynamic dispatch.
type ILabelNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	AllLabelName() []ILabelNameContext
	LabelName(i int) ILabelNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLabelNameListContext differentiates from other interfaces.
	IsLabelNameListContext()
}

type LabelNameListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelNameListContext() *LabelNameListContext {
	var p = new(LabelNameListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelNameList
	return p
}

func InitEmptyLabelNameListContext(p *LabelNameListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_labelNameList
}

func (*LabelNameListContext) IsLabelNameListContext() {}

func NewLabelNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameListContext {
	var p = new(LabelNameListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_labelNameList

	return p
}

func (s *LabelNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelNameListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *LabelNameListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *LabelNameListContext) AllLabelName() []ILabelNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelNameContext); ok {
			len++
		}
	}

	tst := make([]ILabelNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelNameContext); ok {
			tst[i] = t.(ILabelNameContext)
			i++
		}
	}

	return tst
}

func (s *LabelNameListContext) LabelName(i int) ILabelNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelNameContext)
}

func (s *LabelNameListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserCOMMA)
}

func (s *LabelNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserCOMMA, i)
}

func (s *LabelNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLabelNameList(s)
	}
}

func (s *LabelNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLabelNameList(s)
	}
}

func (p *PromQLExtensionParser) LabelNameList() (localctx ILabelNameListContext) {
	localctx = NewLabelNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, PromQLExtensionParserRULE_labelNameList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6756497880776704) != 0 {
		{
			p.SetState(486)
			p.LabelName()
		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(487)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(488)
				p.LabelName()
			}

			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(496)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	UNLESS() antlr.TerminalNode
	BY() antlr.TerminalNode
	WITHOUT() antlr.TerminalNode
	ON() antlr.TerminalNode
	IGNORING() antlr.TerminalNode
	GROUP_LEFT() antlr.TerminalNode
	GROUP_RIGHT() antlr.TerminalNode
	OFFSET() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	AGGREGATION_OPERATOR() antlr.TerminalNode
	FUNCTION() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserAND, 0)
}

func (s *KeywordContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserOR, 0)
}

func (s *KeywordContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserUNLESS, 0)
}

func (s *KeywordContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserBY, 0)
}

func (s *KeywordContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserWITHOUT, 0)
}

func (s *KeywordContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserON, 0)
}

func (s *KeywordContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserIGNORING, 0)
}

func (s *KeywordContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGROUP_LEFT, 0)
}

func (s *KeywordContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserGROUP_RIGHT, 0)
}

func (s *KeywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserOFFSET, 0)
}

func (s *KeywordContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserBOOL, 0)
}

func (s *KeywordContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserAGGREGATION_OPERATOR, 0)
}

func (s *KeywordContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserFUNCTION, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *PromQLExtensionParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, PromQLExtensionParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1098439720960) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserNUMBER, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *PromQLExtensionParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, PromQLExtensionParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserNUMBER || _la == PromQLExtensionParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PromQLExtensionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 29:
		var t *VectorOperationContext = nil
		if localctx != nil {
			t = localctx.(*VectorOperationContext)
		}
		return p.VectorOperation_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLExtensionParser) VectorOperation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
		"", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'",
		"'unless'", "'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'",
		"'!~'", "'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "", "", "'{'", "'}'", "'('", "')'", "'['", "']'",
		"','", "'@'", "", "", "", "", "", "", "'if'", "'true'", "'false'", "'T'",
		"':'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND",
		"OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE",
		"BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET",
		"BOOL", "AGGREGATION_OPERATOR", "FUNCTION", "LEFT_BRACE", "RIGHT_BRACE",
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA",
		"AT", "DURATION", "METRIC_NAME", "LABEL_NAME", "WS", "SL_COMMENT", "EX_ID",
		"EX_IF", "EX_TRUE", "EX_FALSE", "EX_T", "EX_COLON", "EX_DOT", "EX_POSITIVE_INTEGER",
		"EX_TWO_DIGITS", "EX_DIGITS", "EX_NL",
	}
	staticData.RuleNames = []string{
		"promqlx", "ex_statement", "ex_alias_def", "ex_macro_def", "ex_block",
		"ex_arg_list", "ex_if_statement", "ex_condition", "ex_compareVectorOperation",
		"ex_trueConst", "ex_falseConst", "ex_time_instant_literal", "ex_iso_date_time",
		"ex_iso_date_time_ymdhmsf", "ex_iso_date_time_ymdhms", "ex_iso_date_time_ymdhm",
		"ex_iso_date_time_ymdh", "ex_iso_date_time_ymd", "ex_iso_date_time_ym",
		"ex_iso_date_time_y", "ex_year", "ex_month", "ex_day", "ex_hour", "ex_minutes",
		"ex_seconds", "ex_frac_sec", "ex_unix_timestamp", "ex_const_num_expression",
		"ex_num_literal", "ex_alias_call", "ex_duration", "ex_time_range", "ex_subquery_range",
		"vectorOperation", "unaryOp", "powOp", "multOp", "addOp", "compareOp",
		"andUnlessOp", "orOp", "vectorMatchOp", "subqueryOp", "offsetOp", "vector",
		"parens", "instantSelector", "labelMatcher", "labelMatcherOperator",
		"labelMatcherList", "matrixSelector", "offset", "function_", "parameter",
		"parameterList", "aggregation", "by", "without", "grouping", "on_",
		"ignoring", "groupLeft", "groupRight", "labelName", "labelNameList",
		"keyword", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 563, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 1, 0,
		1, 0, 1, 0, 4, 0, 140, 8, 0, 11, 0, 12, 0, 141, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 148, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 157, 8,
		3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 4, 5, 169,
		8, 5, 11, 5, 12, 5, 170, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7,
		180, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 3, 11, 192, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 201, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 3, 28, 296, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 302, 8,
		29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33,
		1, 33, 1, 33, 3, 33, 316, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 3, 34, 326, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 361, 8, 34, 10,
		34, 12, 34, 364, 9, 34, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 370, 8, 36,
		1, 37, 1, 37, 3, 37, 374, 8, 37, 1, 38, 1, 38, 3, 38, 378, 8, 38, 1, 39,
		1, 39, 3, 39, 382, 8, 39, 1, 39, 3, 39, 385, 8, 39, 1, 40, 1, 40, 3, 40,
		389, 8, 40, 1, 41, 1, 41, 3, 41, 393, 8, 41, 1, 42, 1, 42, 3, 42, 397,
		8, 42, 1, 43, 1, 43, 3, 43, 401, 8, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 413, 8, 45, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 3, 47, 422, 8, 47, 1, 47, 3, 47, 425,
		8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 431, 8, 47, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 5, 50, 442, 8, 50, 10, 50,
		12, 50, 445, 9, 50, 1, 50, 3, 50, 448, 8, 50, 1, 51, 1, 51, 1, 51, 1, 52,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 461, 8, 52, 1,
		53, 1, 53, 1, 53, 1, 53, 1, 53, 5, 53, 468, 8, 53, 10, 53, 12, 53, 471,
		9, 53, 3, 53, 473, 8, 53, 1, 53, 1, 53, 1, 54, 1, 54, 3, 54, 479, 8, 54,
		1, 55, 1, 55, 1, 55, 1, 55, 5, 55, 485, 8, 55, 10, 55, 12, 55, 488, 9,
		55, 3, 55, 490, 8, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		3, 56, 499, 8, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 507,
		8, 56, 3, 56, 509, 8, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1,
		59, 1, 59, 3, 59, 519, 8, 59, 1, 59, 1, 59, 3, 59, 523, 8, 59, 1, 60, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 3, 62, 533, 8, 62, 1, 63,
		1, 63, 3, 63, 537, 8, 63, 1, 64, 1, 64, 1, 64, 3, 64, 542, 8, 64, 1, 65,
		1, 65, 1, 65, 1, 65, 5, 65, 548, 8, 65, 10, 65, 12, 65, 551, 9, 65, 3,
		65, 553, 8, 65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 3, 67, 561, 8,
		67, 1, 67, 0, 1, 68, 68, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
		98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126,
		128, 130, 132, 134, 0, 7, 1, 0, 3, 4, 1, 0, 5, 7, 1, 0, 13, 18, 2, 0, 9,
		9, 11, 11, 2, 0, 11, 11, 23, 23, 3, 0, 12, 12, 14, 14, 19, 20, 2, 0, 9,
		11, 21, 30, 570, 0, 139, 1, 0, 0, 0, 2, 147, 1, 0, 0, 0, 4, 149, 1, 0,
		0, 0, 6, 153, 1, 0, 0, 0, 8, 161, 1, 0, 0, 0, 10, 165, 1, 0, 0, 0, 12,
		172, 1, 0, 0, 0, 14, 179, 1, 0, 0, 0, 16, 181, 1, 0, 0, 0, 18, 185, 1,
		0, 0, 0, 20, 187, 1, 0, 0, 0, 22, 191, 1, 0, 0, 0, 24, 200, 1, 0, 0, 0,
		26, 202, 1, 0, 0, 0, 28, 216, 1, 0, 0, 0, 30, 228, 1, 0, 0, 0, 32, 238,
		1, 0, 0, 0, 34, 246, 1, 0, 0, 0, 36, 252, 1, 0, 0, 0, 38, 256, 1, 0, 0,
		0, 40, 258, 1, 0, 0, 0, 42, 260, 1, 0, 0, 0, 44, 262, 1, 0, 0, 0, 46, 264,
		1, 0, 0, 0, 48, 266, 1, 0, 0, 0, 50, 268, 1, 0, 0, 0, 52, 270, 1, 0, 0,
		0, 54, 272, 1, 0, 0, 0, 56, 295, 1, 0, 0, 0, 58, 301, 1, 0, 0, 0, 60, 303,
		1, 0, 0, 0, 62, 305, 1, 0, 0, 0, 64, 307, 1, 0, 0, 0, 66, 311, 1, 0, 0,
		0, 68, 325, 1, 0, 0, 0, 70, 365, 1, 0, 0, 0, 72, 367, 1, 0, 0, 0, 74, 371,
		1, 0, 0, 0, 76, 375, 1, 0, 0, 0, 78, 379, 1, 0, 0, 0, 80, 386, 1, 0, 0,
		0, 82, 390, 1, 0, 0, 0, 84, 394, 1, 0, 0, 0, 86, 398, 1, 0, 0, 0, 88, 402,
		1, 0, 0, 0, 90, 412, 1, 0, 0, 0, 92, 414, 1, 0, 0, 0, 94, 430, 1, 0, 0,
		0, 96, 432, 1, 0, 0, 0, 98, 436, 1, 0, 0, 0, 100, 438, 1, 0, 0, 0, 102,
		449, 1, 0, 0, 0, 104, 460, 1, 0, 0, 0, 106, 462, 1, 0, 0, 0, 108, 478,
		1, 0, 0, 0, 110, 480, 1, 0, 0, 0, 112, 508, 1, 0, 0, 0, 114, 510, 1, 0,
		0, 0, 116, 513, 1, 0, 0, 0, 118, 518, 1, 0, 0, 0, 120, 524, 1, 0, 0, 0,
		122, 527, 1, 0, 0, 0, 124, 530, 1, 0, 0, 0, 126, 534, 1, 0, 0, 0, 128,
		541, 1, 0, 0, 0, 130, 543, 1, 0, 0, 0, 132, 556, 1, 0, 0, 0, 134, 560,
		1, 0, 0, 0, 136, 137, 3, 2, 1, 0, 137, 138, 5, 54, 0, 0, 138, 140, 1, 0,
		0, 0, 139, 136, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0,
		141, 142, 1, 0, 0, 0, 142, 1, 1, 0, 0, 0, 143, 148, 3, 4, 2, 0, 144, 148,
		3, 6, 3, 0, 145, 148, 3, 12, 6, 0, 146, 148, 3, 68, 34, 0, 147, 143, 1,
		0, 0, 0, 147, 144, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 146, 1, 0, 0,
		0, 148, 3, 1, 0, 0, 0, 149, 150, 5, 44, 0, 0, 150, 151, 5, 12, 0, 0, 151,
		152, 3, 68, 34, 0, 152, 5, 1, 0, 0, 0, 153, 154, 5, 44, 0, 0, 154, 156,
		5, 33, 0, 0, 155, 157, 3, 10, 5, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1,
		0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 5, 34, 0, 0, 159, 160, 3, 8, 4,
		0, 160, 7, 1, 0, 0, 0, 161, 162, 5, 31, 0, 0, 162, 163, 3, 0, 0, 0, 163,
		164, 5, 32, 0, 0, 164, 9, 1, 0, 0, 0, 165, 168, 3, 68, 34, 0, 166, 167,
		5, 37, 0, 0, 167, 169, 3, 68, 34, 0, 168, 166, 1, 0, 0, 0, 169, 170, 1,
		0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 11, 1, 0, 0,
		0, 172, 173, 5, 45, 0, 0, 173, 174, 3, 14, 7, 0, 174, 175, 3, 8, 4, 0,
		175, 13, 1, 0, 0, 0, 176, 180, 3, 16, 8, 0, 177, 180, 3, 18, 9, 0, 178,
		180, 3, 20, 10, 0, 179, 176, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 178,
		1, 0, 0, 0, 180, 15, 1, 0, 0, 0, 181, 182, 3, 68, 34, 0, 182, 183, 3, 78,
		39, 0, 183, 184, 3, 68, 34, 0, 184, 17, 1, 0, 0, 0, 185, 186, 5, 46, 0,
		0, 186, 19, 1, 0, 0, 0, 187, 188, 5, 47, 0, 0, 188, 21, 1, 0, 0, 0, 189,
		192, 3, 24, 12, 0, 190, 192, 3, 54, 27, 0, 191, 189, 1, 0, 0, 0, 191, 190,
		1, 0, 0, 0, 192, 23, 1, 0, 0, 0, 193, 201, 3, 26, 13, 0, 194, 201, 3, 28,
		14, 0, 195, 201, 3, 30, 15, 0, 196, 201, 3, 32, 16, 0, 197, 201, 3, 34,
		17, 0, 198, 201, 3, 36, 18, 0, 199, 201, 3, 38, 19, 0, 200, 193, 1, 0,
		0, 0, 200, 194, 1, 0, 0, 0, 200, 195, 1, 0, 0, 0, 200, 196, 1, 0, 0, 0,
		200, 197, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		25, 1, 0, 0, 0, 202, 203, 3, 40, 20, 0, 203, 204, 5, 4, 0, 0, 204, 205,
		3, 42, 21, 0, 205, 206, 5, 4, 0, 0, 206, 207, 3, 44, 22, 0, 207, 208, 5,
		48, 0, 0, 208, 209, 3, 46, 23, 0, 209, 210, 5, 49, 0, 0, 210, 211, 3, 48,
		24, 0, 211, 212, 5, 49, 0, 0, 212, 213, 3, 50, 25, 0, 213, 214, 5, 50,
		0, 0, 214, 215, 3, 52, 26, 0, 215, 27, 1, 0, 0, 0, 216, 217, 3, 40, 20,
		0, 217, 218, 5, 4, 0, 0, 218, 219, 3, 42, 21, 0, 219, 220, 5, 4, 0, 0,
		220, 221, 3, 44, 22, 0, 221, 222, 5, 48, 0, 0, 222, 223, 3, 46, 23, 0,
		223, 224, 5, 49, 0, 0, 224, 225, 3, 48, 24, 0, 225, 226, 5, 49, 0, 0, 226,
		227, 3, 50, 25, 0, 227, 29, 1, 0, 0, 0, 228, 229, 3, 40, 20, 0, 229, 230,
		5, 4, 0, 0, 230, 231, 3, 42, 21, 0, 231, 232, 5, 4, 0, 0, 232, 233, 3,
		44, 22, 0, 233, 234, 5, 48, 0, 0, 234, 235, 3, 46, 23, 0, 235, 236, 5,
		49, 0, 0, 236, 237, 3, 48, 24, 0, 237, 31, 1, 0, 0, 0, 238, 239, 3, 40,
		20, 0, 239, 240, 5, 4, 0, 0, 240, 241, 3, 42, 21, 0, 241, 242, 5, 4, 0,
		0, 242, 243, 3, 44, 22, 0, 243, 244, 5, 48, 0, 0, 244, 245, 3, 46, 23,
		0, 245, 33, 1, 0, 0, 0, 246, 247, 3, 40, 20, 0, 247, 248, 5, 4, 0, 0, 248,
		249, 3, 42, 21, 0, 249, 250, 5, 4, 0, 0, 250, 251, 3, 44, 22, 0, 251, 35,
		1, 0, 0, 0, 252, 253, 3, 40, 20, 0, 253, 254, 5, 4, 0, 0, 254, 255, 3,
		42, 21, 0, 255, 37, 1, 0, 0, 0, 256, 257, 3, 40, 20, 0, 257, 39, 1, 0,
		0, 0, 258, 259, 5, 51, 0, 0, 259, 41, 1, 0, 0, 0, 260, 261, 5, 52, 0, 0,
		261, 43, 1, 0, 0, 0, 262, 263, 5, 52, 0, 0, 263, 45, 1, 0, 0, 0, 264, 265,
		5, 52, 0, 0, 265, 47, 1, 0, 0, 0, 266, 267, 5, 52, 0, 0, 267, 49, 1, 0,
		0, 0, 268, 269, 5, 52, 0, 0, 269, 51, 1, 0, 0, 0, 270, 271, 5, 53, 0, 0,
		271, 53, 1, 0, 0, 0, 272, 273, 5, 51, 0, 0, 273, 55, 1, 0, 0, 0, 274, 275,
		3, 58, 29, 0, 275, 276, 3, 72, 36, 0, 276, 277, 3, 58, 29, 0, 277, 296,
		1, 0, 0, 0, 278, 279, 3, 70, 35, 0, 279, 280, 3, 58, 29, 0, 280, 296, 1,
		0, 0, 0, 281, 282, 3, 58, 29, 0, 282, 283, 3, 74, 37, 0, 283, 284, 3, 58,
		29, 0, 284, 296, 1, 0, 0, 0, 285, 286, 3, 58, 29, 0, 286, 287, 3, 76, 38,
		0, 287, 288, 3, 58, 29, 0, 288, 296, 1, 0, 0, 0, 289, 290, 5, 33, 0, 0,
		290, 291, 3, 56, 28, 0, 291, 292, 5, 34, 0, 0, 292, 296, 1, 0, 0, 0, 293,
		296, 3, 60, 30, 0, 294, 296, 3, 58, 29, 0, 295, 274, 1, 0, 0, 0, 295, 278,
		1, 0, 0, 0, 295, 281, 1, 0, 0, 0, 295, 285, 1, 0, 0, 0, 295, 289, 1, 0,
		0, 0, 295, 293, 1, 0, 0, 0, 295, 294, 1, 0, 0, 0, 296, 57, 1, 0, 0, 0,
		297, 302, 5, 1, 0, 0, 298, 302, 5, 39, 0, 0, 299, 302, 3, 22, 11, 0, 300,
		302, 3, 60, 30, 0, 301, 297, 1, 0, 0, 0, 301, 298, 1, 0, 0, 0, 301, 299,
		1, 0, 0, 0, 301, 300, 1, 0, 0, 0, 302, 59, 1, 0, 0, 0, 303, 304, 5, 44,
		0, 0, 304, 61, 1, 0, 0, 0, 305, 306, 3, 56, 28, 0, 306, 63, 1, 0, 0, 0,
		307, 308, 5, 35, 0, 0, 308, 309, 3, 62, 31, 0, 309, 310, 5, 36, 0, 0, 310,
		65, 1, 0, 0, 0, 311, 312, 5, 35, 0, 0, 312, 313, 3, 62, 31, 0, 313, 315,
		5, 49, 0, 0, 314, 316, 3, 62, 31, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1,
		0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 318, 5, 36, 0, 0, 318, 67, 1, 0, 0,
		0, 319, 320, 6, 34, -1, 0, 320, 321, 3, 70, 35, 0, 321, 322, 3, 68, 34,
		10, 322, 326, 1, 0, 0, 0, 323, 326, 3, 90, 45, 0, 324, 326, 3, 60, 30,
		0, 325, 319, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 324, 1, 0, 0, 0, 326,
		362, 1, 0, 0, 0, 327, 328, 10, 12, 0, 0, 328, 329, 3, 72, 36, 0, 329, 330,
		3, 68, 34, 12, 330, 361, 1, 0, 0, 0, 331, 332, 10, 9, 0, 0, 332, 333, 3,
		74, 37, 0, 333, 334, 3, 68, 34, 10, 334, 361, 1, 0, 0, 0, 335, 336, 10,
		8, 0, 0, 336, 337, 3, 76, 38, 0, 337, 338, 3, 68, 34, 9, 338, 361, 1, 0,
		0, 0, 339, 340, 10, 7, 0, 0, 340, 341, 3, 78, 39, 0, 341, 342, 3, 68, 34,
		8, 342, 361, 1, 0, 0, 0, 343, 344, 10, 6, 0, 0, 344, 345, 3, 80, 40, 0,
		345, 346, 3, 68, 34, 7, 346, 361, 1, 0, 0, 0, 347, 348, 10, 5, 0, 0, 348,
		349, 3, 82, 41, 0, 349, 350, 3, 68, 34, 6, 350, 361, 1, 0, 0, 0, 351, 352,
		10, 4, 0, 0, 352, 353, 3, 84, 42, 0, 353, 354, 3, 68, 34, 5, 354, 361,
		1, 0, 0, 0, 355, 356, 10, 3, 0, 0, 356, 357, 5, 38, 0, 0, 357, 361, 3,
		68, 34, 4, 358, 359, 10, 11, 0, 0, 359, 361, 3, 86, 43, 0, 360, 327, 1,
		0, 0, 0, 360, 331, 1, 0, 0, 0, 360, 335, 1, 0, 0, 0, 360, 339, 1, 0, 0,
		0, 360, 343, 1, 0, 0, 0, 360, 347, 1, 0, 0, 0, 360, 351, 1, 0, 0, 0, 360,
		355, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 364, 1, 0, 0, 0, 362, 360,
		1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 69, 1, 0, 0, 0, 364, 362, 1, 0,
		0, 0, 365, 366, 7, 0, 0, 0, 366, 71, 1, 0, 0, 0, 367, 369, 5, 8, 0, 0,
		368, 370, 3, 118, 59, 0, 369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370,
		73, 1, 0, 0, 0, 371, 373, 7, 1, 0, 0, 372, 374, 3, 118, 59, 0, 373, 372,
		1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 75, 1, 0, 0, 0, 375, 377, 7, 0,
		0, 0, 376, 378, 3, 118, 59, 0, 377, 376, 1, 0, 0, 0, 377, 378, 1, 0, 0,
		0, 378, 77, 1, 0, 0, 0, 379, 381, 7, 2, 0, 0, 380, 382, 5, 28, 0, 0, 381,
		380, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 384, 1, 0, 0, 0, 383, 385,
		3, 118, 59, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 79, 1,
		0, 0, 0, 386, 388, 7, 3, 0, 0, 387, 389, 3, 118, 59, 0, 388, 387, 1, 0,
		0, 0, 388, 389, 1, 0, 0, 0, 389, 81, 1, 0, 0, 0, 390, 392, 5, 10, 0, 0,
		391, 393, 3, 118, 59, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393,
		83, 1, 0, 0, 0, 394, 396, 7, 4, 0, 0, 395, 397, 3, 118, 59, 0, 396, 395,
		1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 85, 1, 0, 0, 0, 398, 400, 3, 66,
		33, 0, 399, 401, 3, 88, 44, 0, 400, 399, 1, 0, 0, 0, 400, 401, 1, 0, 0,
		0, 401, 87, 1, 0, 0, 0, 402, 403, 5, 27, 0, 0, 403, 404, 3, 62, 31, 0,
		404, 89, 1, 0, 0, 0, 405, 413, 3, 106, 53, 0, 406, 413, 3, 112, 56, 0,
		407, 413, 3, 94, 47, 0, 408, 413, 3, 102, 51, 0, 409, 413, 3, 104, 52,
		0, 410, 413, 3, 134, 67, 0, 411, 413, 3, 92, 46, 0, 412, 405, 1, 0, 0,
		0, 412, 406, 1, 0, 0, 0, 412, 407, 1, 0, 0, 0, 412, 408, 1, 0, 0, 0, 412,
		409, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 91, 1,
		0, 0, 0, 414, 415, 5, 33, 0, 0, 415, 416, 3, 68, 34, 0, 416, 417, 5, 34,
		0, 0, 417, 93, 1, 0, 0, 0, 418, 424, 5, 40, 0, 0, 419, 421, 5, 31, 0, 0,
		420, 422, 3, 100, 50, 0, 421, 420, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422,
		423, 1, 0, 0, 0, 423, 425, 5, 32, 0, 0, 424, 419, 1, 0, 0, 0, 424, 425,
		1, 0, 0, 0, 425, 431, 1, 0, 0, 0, 426, 427, 5, 31, 0, 0, 427, 428, 3, 100,
		50, 0, 428, 429, 5, 32, 0, 0, 429, 431, 1, 0, 0, 0, 430, 418, 1, 0, 0,
		0, 430, 426, 1, 0, 0, 0, 431, 95, 1, 0, 0, 0, 432, 433, 3, 128, 64, 0,
		433, 434, 3, 98, 49, 0, 434, 435, 5, 2, 0, 0, 435, 97, 1, 0, 0, 0, 436,
		437, 7, 5, 0, 0, 437, 99, 1, 0, 0, 0, 438, 443, 3, 96, 48, 0, 439, 440,
		5, 37, 0, 0, 440, 442, 3, 96, 48, 0, 441, 439, 1, 0, 0, 0, 442, 445, 1,
		0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 447, 1, 0, 0,
		0, 445, 443, 1, 0, 0, 0, 446, 448, 5, 37, 0, 0, 447, 446, 1, 0, 0, 0, 447,
		448, 1, 0, 0, 0, 448, 101, 1, 0, 0, 0, 449, 450, 3, 94, 47, 0, 450, 451,
		3, 64, 32, 0, 451, 103, 1, 0, 0, 0, 452, 453, 3, 94, 47, 0, 453, 454, 5,
		27, 0, 0, 454, 455, 3, 62, 31, 0, 455, 461, 1, 0, 0, 0, 456, 457, 3, 102,
		51, 0, 457, 458, 5, 27, 0, 0, 458, 459, 3, 62, 31, 0, 459, 461, 1, 0, 0,
		0, 460, 452, 1, 0, 0, 0, 460, 456, 1, 0, 0, 0, 461, 105, 1, 0, 0, 0, 462,
		463, 5, 30, 0, 0, 463, 472, 5, 33, 0, 0, 464, 469, 3, 108, 54, 0, 465,
		466, 5, 37, 0, 0, 466, 468, 3, 108, 54, 0, 467, 465, 1, 0, 0, 0, 468, 471,
		1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 473, 1, 0,
		0, 0, 471, 469, 1, 0, 0, 0, 472, 464, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0,
		473, 474, 1, 0, 0, 0, 474, 475, 5, 34, 0, 0, 475, 107, 1, 0, 0, 0, 476,
		479, 3, 134, 67, 0, 477, 479, 3, 68, 34, 0, 478, 476, 1, 0, 0, 0, 478,
		477, 1, 0, 0, 0, 479, 109, 1, 0, 0, 0, 480, 489, 5, 33, 0, 0, 481, 486,
		3, 108, 54, 0, 482, 483, 5, 37, 0, 0, 483, 485, 3, 108, 54, 0, 484, 482,
		1, 0, 0, 0, 485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0,
		0, 0, 487, 490, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 489, 481, 1, 0, 0, 0,
		489, 490, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 492, 5, 34, 0, 0, 492,
		111, 1, 0, 0, 0, 493, 494, 5, 29, 0, 0, 494, 509, 3, 110, 55, 0, 495, 498,
		5, 29, 0, 0, 496, 499, 3, 114, 57, 0, 497, 499, 3, 116, 58, 0, 498, 496,
		1, 0, 0, 0, 498, 497, 1, 0, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 3, 110,
		55, 0, 501, 509, 1, 0, 0, 0, 502, 503, 5, 29, 0, 0, 503, 506, 3, 110, 55,
		0, 504, 507, 3, 114, 57, 0, 505, 507, 3, 116, 58, 0, 506, 504, 1, 0, 0,
		0, 506, 505, 1, 0, 0, 0, 507, 509, 1, 0, 0, 0, 508, 493, 1, 0, 0, 0, 508,
		495, 1, 0, 0, 0, 508, 502, 1, 0, 0, 0, 509, 113, 1, 0, 0, 0, 510, 511,
		5, 21, 0, 0, 511, 512, 3, 130, 65, 0, 512, 115, 1, 0, 0, 0, 513, 514, 5,
		22, 0, 0, 514, 515, 3, 130, 65, 0, 515, 117, 1, 0, 0, 0, 516, 519, 3, 120,
		60, 0, 517, 519, 3, 122, 61, 0, 518, 516, 1, 0, 0, 0, 518, 517, 1, 0, 0,
		0, 519, 522, 1, 0, 0, 0, 520, 523, 3, 124, 62, 0, 521, 523, 3, 126, 63,
		0, 522, 520, 1, 0, 0, 0, 522, 521, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523,
		119, 1, 0, 0, 0, 524, 525, 5, 23, 0, 0, 525, 526, 3, 130, 65, 0, 526, 121,
		1, 0, 0, 0, 527, 528, 5, 24, 0, 0, 528, 529, 3, 130, 65, 0, 529, 123, 1,
		0, 0, 0, 530, 532, 5, 25, 0, 0, 531, 533, 3, 130, 65, 0, 532, 531, 1, 0,
		0, 0, 532, 533, 1, 0, 0, 0, 533, 125, 1, 0, 0, 0, 534, 536, 5, 26, 0, 0,
		535, 537, 3, 130, 65, 0, 536, 535, 1, 0, 0, 0, 536, 537, 1, 0, 0, 0, 537,
		127, 1, 0, 0, 0, 538, 542, 3, 132, 66, 0, 539, 542, 5, 40, 0, 0, 540, 542,
		5, 41, 0, 0, 541, 538, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 541, 540, 1, 0,
		0, 0, 542, 129, 1, 0, 0, 0, 543, 552, 5, 33, 0, 0, 544, 549, 3, 128, 64,
		0, 545, 546, 5, 37, 0, 0, 546, 548, 3, 128, 64, 0, 547, 545, 1, 0, 0, 0,
		548, 551, 1, 0, 0, 0, 549, 547, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550,
		553, 1, 0, 0, 0, 551, 549, 1, 0, 0, 0, 552, 544, 1, 0, 0, 0, 552, 553,
		1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 555, 5, 34, 0, 0, 555, 131, 1, 0,
		0, 0, 556, 557, 7, 6, 0, 0, 557, 133, 1, 0, 0, 0, 558, 561, 3, 56, 28,
		0, 559, 561, 5, 2, 0, 0, 560, 558, 1, 0, 0, 0, 560, 559, 1, 0, 0, 0, 561,
		135, 1, 0, 0, 0, 45, 141, 147, 156, 170, 179, 191, 200, 295, 301, 315,
		325, 360, 362, 369, 373, 377, 381, 384, 388, 392, 396, 400, 412, 421, 424,
		430, 443, 447, 460, 469, 472, 478, 486, 489, 498, 506, 508, 518, 522, 532,
		536, 541, 549, 552, 560,
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
	PromQLExtensionParserNUMBER               = 1
	PromQLExtensionParserSTRING               = 2
	PromQLExtensionParserADD                  = 3
	PromQLExtensionParserSUB                  = 4
	PromQLExtensionParserMULT                 = 5
	PromQLExtensionParserDIV                  = 6
	PromQLExtensionParserMOD                  = 7
	PromQLExtensionParserPOW                  = 8
	PromQLExtensionParserAND                  = 9
	PromQLExtensionParserOR                   = 10
	PromQLExtensionParserUNLESS               = 11
	PromQLExtensionParserEQ                   = 12
	PromQLExtensionParserDEQ                  = 13
	PromQLExtensionParserNE                   = 14
	PromQLExtensionParserGT                   = 15
	PromQLExtensionParserLT                   = 16
	PromQLExtensionParserGE                   = 17
	PromQLExtensionParserLE                   = 18
	PromQLExtensionParserRE                   = 19
	PromQLExtensionParserNRE                  = 20
	PromQLExtensionParserBY                   = 21
	PromQLExtensionParserWITHOUT              = 22
	PromQLExtensionParserON                   = 23
	PromQLExtensionParserIGNORING             = 24
	PromQLExtensionParserGROUP_LEFT           = 25
	PromQLExtensionParserGROUP_RIGHT          = 26
	PromQLExtensionParserOFFSET               = 27
	PromQLExtensionParserBOOL                 = 28
	PromQLExtensionParserAGGREGATION_OPERATOR = 29
	PromQLExtensionParserFUNCTION             = 30
	PromQLExtensionParserLEFT_BRACE           = 31
	PromQLExtensionParserRIGHT_BRACE          = 32
	PromQLExtensionParserLEFT_PAREN           = 33
	PromQLExtensionParserRIGHT_PAREN          = 34
	PromQLExtensionParserLEFT_BRACKET         = 35
	PromQLExtensionParserRIGHT_BRACKET        = 36
	PromQLExtensionParserCOMMA                = 37
	PromQLExtensionParserAT                   = 38
	PromQLExtensionParserDURATION             = 39
	PromQLExtensionParserMETRIC_NAME          = 40
	PromQLExtensionParserLABEL_NAME           = 41
	PromQLExtensionParserWS                   = 42
	PromQLExtensionParserSL_COMMENT           = 43
	PromQLExtensionParserEX_ID                = 44
	PromQLExtensionParserEX_IF                = 45
	PromQLExtensionParserEX_TRUE              = 46
	PromQLExtensionParserEX_FALSE             = 47
	PromQLExtensionParserEX_T                 = 48
	PromQLExtensionParserEX_COLON             = 49
	PromQLExtensionParserEX_DOT               = 50
	PromQLExtensionParserEX_POSITIVE_INTEGER  = 51
	PromQLExtensionParserEX_TWO_DIGITS        = 52
	PromQLExtensionParserEX_DIGITS            = 53
	PromQLExtensionParserEX_NL                = 54
)

// PromQLExtensionParser rules.
const (
	PromQLExtensionParserRULE_promqlx                   = 0
	PromQLExtensionParserRULE_ex_statement              = 1
	PromQLExtensionParserRULE_ex_alias_def              = 2
	PromQLExtensionParserRULE_ex_macro_def              = 3
	PromQLExtensionParserRULE_ex_block                  = 4
	PromQLExtensionParserRULE_ex_arg_list               = 5
	PromQLExtensionParserRULE_ex_if_statement           = 6
	PromQLExtensionParserRULE_ex_condition              = 7
	PromQLExtensionParserRULE_ex_compareVectorOperation = 8
	PromQLExtensionParserRULE_ex_trueConst              = 9
	PromQLExtensionParserRULE_ex_falseConst             = 10
	PromQLExtensionParserRULE_ex_time_instant_literal   = 11
	PromQLExtensionParserRULE_ex_iso_date_time          = 12
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf  = 13
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhms   = 14
	PromQLExtensionParserRULE_ex_iso_date_time_ymdhm    = 15
	PromQLExtensionParserRULE_ex_iso_date_time_ymdh     = 16
	PromQLExtensionParserRULE_ex_iso_date_time_ymd      = 17
	PromQLExtensionParserRULE_ex_iso_date_time_ym       = 18
	PromQLExtensionParserRULE_ex_iso_date_time_y        = 19
	PromQLExtensionParserRULE_ex_year                   = 20
	PromQLExtensionParserRULE_ex_month                  = 21
	PromQLExtensionParserRULE_ex_day                    = 22
	PromQLExtensionParserRULE_ex_hour                   = 23
	PromQLExtensionParserRULE_ex_minutes                = 24
	PromQLExtensionParserRULE_ex_seconds                = 25
	PromQLExtensionParserRULE_ex_frac_sec               = 26
	PromQLExtensionParserRULE_ex_unix_timestamp         = 27
	PromQLExtensionParserRULE_ex_const_num_expression   = 28
	PromQLExtensionParserRULE_ex_num_literal            = 29
	PromQLExtensionParserRULE_ex_alias_call             = 30
	PromQLExtensionParserRULE_ex_duration               = 31
	PromQLExtensionParserRULE_ex_time_range             = 32
	PromQLExtensionParserRULE_ex_subquery_range         = 33
	PromQLExtensionParserRULE_vectorOperation           = 34
	PromQLExtensionParserRULE_unaryOp                   = 35
	PromQLExtensionParserRULE_powOp                     = 36
	PromQLExtensionParserRULE_multOp                    = 37
	PromQLExtensionParserRULE_addOp                     = 38
	PromQLExtensionParserRULE_compareOp                 = 39
	PromQLExtensionParserRULE_andUnlessOp               = 40
	PromQLExtensionParserRULE_orOp                      = 41
	PromQLExtensionParserRULE_vectorMatchOp             = 42
	PromQLExtensionParserRULE_subqueryOp                = 43
	PromQLExtensionParserRULE_offsetOp                  = 44
	PromQLExtensionParserRULE_vector                    = 45
	PromQLExtensionParserRULE_parens                    = 46
	PromQLExtensionParserRULE_instantSelector           = 47
	PromQLExtensionParserRULE_labelMatcher              = 48
	PromQLExtensionParserRULE_labelMatcherOperator      = 49
	PromQLExtensionParserRULE_labelMatcherList          = 50
	PromQLExtensionParserRULE_matrixSelector            = 51
	PromQLExtensionParserRULE_offset                    = 52
	PromQLExtensionParserRULE_function_                 = 53
	PromQLExtensionParserRULE_parameter                 = 54
	PromQLExtensionParserRULE_parameterList             = 55
	PromQLExtensionParserRULE_aggregation               = 56
	PromQLExtensionParserRULE_by                        = 57
	PromQLExtensionParserRULE_without                   = 58
	PromQLExtensionParserRULE_grouping                  = 59
	PromQLExtensionParserRULE_on_                       = 60
	PromQLExtensionParserRULE_ignoring                  = 61
	PromQLExtensionParserRULE_groupLeft                 = 62
	PromQLExtensionParserRULE_groupRight                = 63
	PromQLExtensionParserRULE_labelName                 = 64
	PromQLExtensionParserRULE_labelNameList             = 65
	PromQLExtensionParserRULE_keyword                   = 66
	PromQLExtensionParserRULE_literal                   = 67
)

// IPromqlxContext is an interface to support dynamic dispatch.
type IPromqlxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEx_statement() []IEx_statementContext
	Ex_statement(i int) IEx_statementContext
	AllEX_NL() []antlr.TerminalNode
	EX_NL(i int) antlr.TerminalNode

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

func (s *PromqlxContext) AllEx_statement() []IEx_statementContext {
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

func (s *PromqlxContext) Ex_statement(i int) IEx_statementContext {
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

func (s *PromqlxContext) AllEX_NL() []antlr.TerminalNode {
	return s.GetTokens(PromQLExtensionParserEX_NL)
}

func (s *PromqlxContext) EX_NL(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_NL, i)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2306237987291166) != 0) {
		{
			p.SetState(136)
			p.Ex_statement()
		}
		{
			p.SetState(137)
			p.Match(PromQLExtensionParserEX_NL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(141)
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
	VectorOperation() IVectorOperationContext

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

func (s *Ex_statementContext) VectorOperation() IVectorOperationContext {
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
	p.EnterRule(localctx, 2, PromQLExtensionParserRULE_ex_statement)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Ex_alias_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Ex_macro_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Ex_if_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(146)
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

// IEx_alias_defContext is an interface to support dynamic dispatch.
type IEx_alias_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_ID() antlr.TerminalNode
	EQ() antlr.TerminalNode
	VectorOperation() IVectorOperationContext

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

func (s *Ex_alias_defContext) VectorOperation() IVectorOperationContext {
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
	p.EnterRule(localctx, 4, PromQLExtensionParserRULE_ex_alias_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(PromQLExtensionParserEX_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(PromQLExtensionParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
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
	p.EnterRule(localctx, 6, PromQLExtensionParserRULE_ex_macro_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(PromQLExtensionParserEX_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2271053615202334) != 0 {
		{
			p.SetState(155)
			p.Ex_arg_list()
		}

	}
	{
		p.SetState(158)
		p.Match(PromQLExtensionParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
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
	Promqlx() IPromqlxContext
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

func (s *Ex_blockContext) Promqlx() IPromqlxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromqlxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromqlxContext)
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
	p.EnterRule(localctx, 8, PromQLExtensionParserRULE_ex_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(PromQLExtensionParserLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Promqlx()
	}
	{
		p.SetState(163)
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
	AllVectorOperation() []IVectorOperationContext
	VectorOperation(i int) IVectorOperationContext
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

func (s *Ex_arg_listContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *Ex_arg_listContext) VectorOperation(i int) IVectorOperationContext {
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
	p.EnterRule(localctx, 10, PromQLExtensionParserRULE_ex_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.vectorOperation(0)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromQLExtensionParserCOMMA {
		{
			p.SetState(166)
			p.Match(PromQLExtensionParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.vectorOperation(0)
		}

		p.SetState(170)
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
	EX_IF() antlr.TerminalNode
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

func (s *Ex_if_statementContext) EX_IF() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_IF, 0)
}

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
	p.EnterRule(localctx, 12, PromQLExtensionParserRULE_ex_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(PromQLExtensionParserEX_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Ex_condition()
	}
	{
		p.SetState(174)
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
	p.EnterRule(localctx, 14, PromQLExtensionParserRULE_ex_condition)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserNUMBER, PromQLExtensionParserSTRING, PromQLExtensionParserADD, PromQLExtensionParserSUB, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION, PromQLExtensionParserLEFT_BRACE, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserDURATION, PromQLExtensionParserMETRIC_NAME, PromQLExtensionParserEX_ID, PromQLExtensionParserEX_POSITIVE_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Ex_compareVectorOperation()
		}

	case PromQLExtensionParserEX_TRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Ex_trueConst()
		}

	case PromQLExtensionParserEX_FALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
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
	p.EnterRule(localctx, 16, PromQLExtensionParserRULE_ex_compareVectorOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.vectorOperation(0)
	}
	{
		p.SetState(182)
		p.CompareOp()
	}
	{
		p.SetState(183)
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
	p.EnterRule(localctx, 18, PromQLExtensionParserRULE_ex_trueConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
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
	p.EnterRule(localctx, 20, PromQLExtensionParserRULE_ex_falseConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
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

// IEx_time_instant_literalContext is an interface to support dynamic dispatch.
type IEx_time_instant_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_iso_date_time() IEx_iso_date_timeContext
	Ex_unix_timestamp() IEx_unix_timestampContext

	// IsEx_time_instant_literalContext differentiates from other interfaces.
	IsEx_time_instant_literalContext()
}

type Ex_time_instant_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_time_instant_literalContext() *Ex_time_instant_literalContext {
	var p = new(Ex_time_instant_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant_literal
	return p
}

func InitEmptyEx_time_instant_literalContext(p *Ex_time_instant_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant_literal
}

func (*Ex_time_instant_literalContext) IsEx_time_instant_literalContext() {}

func NewEx_time_instant_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_time_instant_literalContext {
	var p = new(Ex_time_instant_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_instant_literal

	return p
}

func (s *Ex_time_instant_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_time_instant_literalContext) Ex_iso_date_time() IEx_iso_date_timeContext {
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

func (s *Ex_time_instant_literalContext) Ex_unix_timestamp() IEx_unix_timestampContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_unix_timestampContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_unix_timestampContext)
}

func (s *Ex_time_instant_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_time_instant_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_time_instant_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_time_instant_literal(s)
	}
}

func (s *Ex_time_instant_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_time_instant_literal(s)
	}
}

func (p *PromQLExtensionParser) Ex_time_instant_literal() (localctx IEx_time_instant_literalContext) {
	localctx = NewEx_time_instant_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromQLExtensionParserRULE_ex_time_instant_literal)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Ex_iso_date_time()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Ex_unix_timestamp()
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
	p.EnterRule(localctx, 24, PromQLExtensionParserRULE_ex_iso_date_time)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Ex_iso_date_time_ymdhmsf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Ex_iso_date_time_ymdhms()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Ex_iso_date_time_ymdhm()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(196)
			p.Ex_iso_date_time_ymdh()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(197)
			p.Ex_iso_date_time_ymd()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(198)
			p.Ex_iso_date_time_ym()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(199)
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
	p.EnterRule(localctx, 26, PromQLExtensionParserRULE_ex_iso_date_time_ymdhmsf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Ex_year()
	}
	{
		p.SetState(203)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Ex_month()
	}
	{
		p.SetState(205)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Ex_day()
	}
	{
		p.SetState(207)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Ex_hour()
	}
	{
		p.SetState(209)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Ex_minutes()
	}
	{
		p.SetState(211)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Ex_seconds()
	}
	{
		p.SetState(213)
		p.Match(PromQLExtensionParserEX_DOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
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
	p.EnterRule(localctx, 28, PromQLExtensionParserRULE_ex_iso_date_time_ymdhms)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Ex_year()
	}
	{
		p.SetState(217)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Ex_month()
	}
	{
		p.SetState(219)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Ex_day()
	}
	{
		p.SetState(221)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Ex_hour()
	}
	{
		p.SetState(223)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Ex_minutes()
	}
	{
		p.SetState(225)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
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
	p.EnterRule(localctx, 30, PromQLExtensionParserRULE_ex_iso_date_time_ymdhm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Ex_year()
	}
	{
		p.SetState(229)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Ex_month()
	}
	{
		p.SetState(231)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Ex_day()
	}
	{
		p.SetState(233)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.Ex_hour()
	}
	{
		p.SetState(235)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
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
	p.EnterRule(localctx, 32, PromQLExtensionParserRULE_ex_iso_date_time_ymdh)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Ex_year()
	}
	{
		p.SetState(239)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Ex_month()
	}
	{
		p.SetState(241)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Ex_day()
	}
	{
		p.SetState(243)
		p.Match(PromQLExtensionParserEX_T)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
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
	p.EnterRule(localctx, 34, PromQLExtensionParserRULE_ex_iso_date_time_ymd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Ex_year()
	}
	{
		p.SetState(247)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Ex_month()
	}
	{
		p.SetState(249)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
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
	p.EnterRule(localctx, 36, PromQLExtensionParserRULE_ex_iso_date_time_ym)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Ex_year()
	}
	{
		p.SetState(253)
		p.Match(PromQLExtensionParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
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
	p.EnterRule(localctx, 38, PromQLExtensionParserRULE_ex_iso_date_time_y)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
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
	p.EnterRule(localctx, 40, PromQLExtensionParserRULE_ex_year)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 42, PromQLExtensionParserRULE_ex_month)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
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
	p.EnterRule(localctx, 44, PromQLExtensionParserRULE_ex_day)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
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
	p.EnterRule(localctx, 46, PromQLExtensionParserRULE_ex_hour)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
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
	p.EnterRule(localctx, 48, PromQLExtensionParserRULE_ex_minutes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
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
	p.EnterRule(localctx, 50, PromQLExtensionParserRULE_ex_seconds)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
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
	p.EnterRule(localctx, 52, PromQLExtensionParserRULE_ex_frac_sec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
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

// IEx_unix_timestampContext is an interface to support dynamic dispatch.
type IEx_unix_timestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_POSITIVE_INTEGER() antlr.TerminalNode

	// IsEx_unix_timestampContext differentiates from other interfaces.
	IsEx_unix_timestampContext()
}

type Ex_unix_timestampContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_unix_timestampContext() *Ex_unix_timestampContext {
	var p = new(Ex_unix_timestampContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_unix_timestamp
	return p
}

func InitEmptyEx_unix_timestampContext(p *Ex_unix_timestampContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_unix_timestamp
}

func (*Ex_unix_timestampContext) IsEx_unix_timestampContext() {}

func NewEx_unix_timestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_unix_timestampContext {
	var p = new(Ex_unix_timestampContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_unix_timestamp

	return p
}

func (s *Ex_unix_timestampContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_unix_timestampContext) EX_POSITIVE_INTEGER() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_POSITIVE_INTEGER, 0)
}

func (s *Ex_unix_timestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_unix_timestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_unix_timestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_unix_timestamp(s)
	}
}

func (s *Ex_unix_timestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_unix_timestamp(s)
	}
}

func (p *PromQLExtensionParser) Ex_unix_timestamp() (localctx IEx_unix_timestampContext) {
	localctx = NewEx_unix_timestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromQLExtensionParserRULE_ex_unix_timestamp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
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

// IEx_const_num_expressionContext is an interface to support dynamic dispatch.
type IEx_const_num_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEx_num_literal() []IEx_num_literalContext
	Ex_num_literal(i int) IEx_num_literalContext
	PowOp() IPowOpContext
	UnaryOp() IUnaryOpContext
	MultOp() IMultOpContext
	AddOp() IAddOpContext
	LEFT_PAREN() antlr.TerminalNode
	Ex_const_num_expression() IEx_const_num_expressionContext
	RIGHT_PAREN() antlr.TerminalNode
	Ex_alias_call() IEx_alias_callContext

	// IsEx_const_num_expressionContext differentiates from other interfaces.
	IsEx_const_num_expressionContext()
}

type Ex_const_num_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_const_num_expressionContext() *Ex_const_num_expressionContext {
	var p = new(Ex_const_num_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_const_num_expression
	return p
}

func InitEmptyEx_const_num_expressionContext(p *Ex_const_num_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_const_num_expression
}

func (*Ex_const_num_expressionContext) IsEx_const_num_expressionContext() {}

func NewEx_const_num_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_const_num_expressionContext {
	var p = new(Ex_const_num_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_const_num_expression

	return p
}

func (s *Ex_const_num_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_const_num_expressionContext) AllEx_num_literal() []IEx_num_literalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEx_num_literalContext); ok {
			len++
		}
	}

	tst := make([]IEx_num_literalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEx_num_literalContext); ok {
			tst[i] = t.(IEx_num_literalContext)
			i++
		}
	}

	return tst
}

func (s *Ex_const_num_expressionContext) Ex_num_literal(i int) IEx_num_literalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_num_literalContext); ok {
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

	return t.(IEx_num_literalContext)
}

func (s *Ex_const_num_expressionContext) PowOp() IPowOpContext {
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

func (s *Ex_const_num_expressionContext) UnaryOp() IUnaryOpContext {
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

func (s *Ex_const_num_expressionContext) MultOp() IMultOpContext {
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

func (s *Ex_const_num_expressionContext) AddOp() IAddOpContext {
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

func (s *Ex_const_num_expressionContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_PAREN, 0)
}

func (s *Ex_const_num_expressionContext) Ex_const_num_expression() IEx_const_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_const_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_const_num_expressionContext)
}

func (s *Ex_const_num_expressionContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_PAREN, 0)
}

func (s *Ex_const_num_expressionContext) Ex_alias_call() IEx_alias_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_alias_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_alias_callContext)
}

func (s *Ex_const_num_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_const_num_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_const_num_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_const_num_expression(s)
	}
}

func (s *Ex_const_num_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_const_num_expression(s)
	}
}

func (p *PromQLExtensionParser) Ex_const_num_expression() (localctx IEx_const_num_expressionContext) {
	localctx = NewEx_const_num_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromQLExtensionParserRULE_ex_const_num_expression)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Ex_num_literal()
		}
		{
			p.SetState(275)
			p.PowOp()
		}
		{
			p.SetState(276)
			p.Ex_num_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.UnaryOp()
		}
		{
			p.SetState(279)
			p.Ex_num_literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Ex_num_literal()
		}
		{
			p.SetState(282)
			p.MultOp()
		}
		{
			p.SetState(283)
			p.Ex_num_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(285)
			p.Ex_num_literal()
		}
		{
			p.SetState(286)
			p.AddOp()
		}
		{
			p.SetState(287)
			p.Ex_num_literal()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(289)
			p.Match(PromQLExtensionParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Ex_const_num_expression()
		}
		{
			p.SetState(291)
			p.Match(PromQLExtensionParserRIGHT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(293)
			p.Ex_alias_call()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(294)
			p.Ex_num_literal()
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

// IEx_num_literalContext is an interface to support dynamic dispatch.
type IEx_num_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	Ex_time_instant_literal() IEx_time_instant_literalContext
	Ex_alias_call() IEx_alias_callContext

	// IsEx_num_literalContext differentiates from other interfaces.
	IsEx_num_literalContext()
}

type Ex_num_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_num_literalContext() *Ex_num_literalContext {
	var p = new(Ex_num_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_num_literal
	return p
}

func InitEmptyEx_num_literalContext(p *Ex_num_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_num_literal
}

func (*Ex_num_literalContext) IsEx_num_literalContext() {}

func NewEx_num_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_num_literalContext {
	var p = new(Ex_num_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_num_literal

	return p
}

func (s *Ex_num_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_num_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserNUMBER, 0)
}

func (s *Ex_num_literalContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserDURATION, 0)
}

func (s *Ex_num_literalContext) Ex_time_instant_literal() IEx_time_instant_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_time_instant_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_time_instant_literalContext)
}

func (s *Ex_num_literalContext) Ex_alias_call() IEx_alias_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_alias_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_alias_callContext)
}

func (s *Ex_num_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_num_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_num_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_num_literal(s)
	}
}

func (s *Ex_num_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_num_literal(s)
	}
}

func (p *PromQLExtensionParser) Ex_num_literal() (localctx IEx_num_literalContext) {
	localctx = NewEx_num_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PromQLExtensionParserRULE_ex_num_literal)
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(PromQLExtensionParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExtensionParserDURATION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Match(PromQLExtensionParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExtensionParserEX_POSITIVE_INTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.Ex_time_instant_literal()
		}

	case PromQLExtensionParserEX_ID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(300)
			p.Ex_alias_call()
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

// IEx_alias_callContext is an interface to support dynamic dispatch.
type IEx_alias_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EX_ID() antlr.TerminalNode

	// IsEx_alias_callContext differentiates from other interfaces.
	IsEx_alias_callContext()
}

type Ex_alias_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_alias_callContext() *Ex_alias_callContext {
	var p = new(Ex_alias_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_call
	return p
}

func InitEmptyEx_alias_callContext(p *Ex_alias_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_call
}

func (*Ex_alias_callContext) IsEx_alias_callContext() {}

func NewEx_alias_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_alias_callContext {
	var p = new(Ex_alias_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_alias_call

	return p
}

func (s *Ex_alias_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_alias_callContext) EX_ID() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_ID, 0)
}

func (s *Ex_alias_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_alias_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_alias_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_alias_call(s)
	}
}

func (s *Ex_alias_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_alias_call(s)
	}
}

func (p *PromQLExtensionParser) Ex_alias_call() (localctx IEx_alias_callContext) {
	localctx = NewEx_alias_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromQLExtensionParserRULE_ex_alias_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(PromQLExtensionParserEX_ID)
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

// IEx_durationContext is an interface to support dynamic dispatch.
type IEx_durationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ex_const_num_expression() IEx_const_num_expressionContext

	// IsEx_durationContext differentiates from other interfaces.
	IsEx_durationContext()
}

type Ex_durationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_durationContext() *Ex_durationContext {
	var p = new(Ex_durationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_duration
	return p
}

func InitEmptyEx_durationContext(p *Ex_durationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_duration
}

func (*Ex_durationContext) IsEx_durationContext() {}

func NewEx_durationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_durationContext {
	var p = new(Ex_durationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_duration

	return p
}

func (s *Ex_durationContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_durationContext) Ex_const_num_expression() IEx_const_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_const_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_const_num_expressionContext)
}

func (s *Ex_durationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_durationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_durationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_duration(s)
	}
}

func (s *Ex_durationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_duration(s)
	}
}

func (p *PromQLExtensionParser) Ex_duration() (localctx IEx_durationContext) {
	localctx = NewEx_durationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromQLExtensionParserRULE_ex_duration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Ex_const_num_expression()
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

// IEx_time_rangeContext is an interface to support dynamic dispatch.
type IEx_time_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	Ex_duration() IEx_durationContext
	RIGHT_BRACKET() antlr.TerminalNode

	// IsEx_time_rangeContext differentiates from other interfaces.
	IsEx_time_rangeContext()
}

type Ex_time_rangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_time_rangeContext() *Ex_time_rangeContext {
	var p = new(Ex_time_rangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_range
	return p
}

func InitEmptyEx_time_rangeContext(p *Ex_time_rangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_range
}

func (*Ex_time_rangeContext) IsEx_time_rangeContext() {}

func NewEx_time_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_time_rangeContext {
	var p = new(Ex_time_rangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_time_range

	return p
}

func (s *Ex_time_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_time_rangeContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_BRACKET, 0)
}

func (s *Ex_time_rangeContext) Ex_duration() IEx_durationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_durationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_durationContext)
}

func (s *Ex_time_rangeContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_BRACKET, 0)
}

func (s *Ex_time_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_time_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_time_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_time_range(s)
	}
}

func (s *Ex_time_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_time_range(s)
	}
}

func (p *PromQLExtensionParser) Ex_time_range() (localctx IEx_time_rangeContext) {
	localctx = NewEx_time_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromQLExtensionParserRULE_ex_time_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(PromQLExtensionParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Ex_duration()
	}
	{
		p.SetState(309)
		p.Match(PromQLExtensionParserRIGHT_BRACKET)
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

// IEx_subquery_rangeContext is an interface to support dynamic dispatch.
type IEx_subquery_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	AllEx_duration() []IEx_durationContext
	Ex_duration(i int) IEx_durationContext
	EX_COLON() antlr.TerminalNode
	RIGHT_BRACKET() antlr.TerminalNode

	// IsEx_subquery_rangeContext differentiates from other interfaces.
	IsEx_subquery_rangeContext()
}

type Ex_subquery_rangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEx_subquery_rangeContext() *Ex_subquery_rangeContext {
	var p = new(Ex_subquery_rangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_subquery_range
	return p
}

func InitEmptyEx_subquery_rangeContext(p *Ex_subquery_rangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExtensionParserRULE_ex_subquery_range
}

func (*Ex_subquery_rangeContext) IsEx_subquery_rangeContext() {}

func NewEx_subquery_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ex_subquery_rangeContext {
	var p = new(Ex_subquery_rangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExtensionParserRULE_ex_subquery_range

	return p
}

func (s *Ex_subquery_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ex_subquery_rangeContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserLEFT_BRACKET, 0)
}

func (s *Ex_subquery_rangeContext) AllEx_duration() []IEx_durationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEx_durationContext); ok {
			len++
		}
	}

	tst := make([]IEx_durationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEx_durationContext); ok {
			tst[i] = t.(IEx_durationContext)
			i++
		}
	}

	return tst
}

func (s *Ex_subquery_rangeContext) Ex_duration(i int) IEx_durationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_durationContext); ok {
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

	return t.(IEx_durationContext)
}

func (s *Ex_subquery_rangeContext) EX_COLON() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserEX_COLON, 0)
}

func (s *Ex_subquery_rangeContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExtensionParserRIGHT_BRACKET, 0)
}

func (s *Ex_subquery_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ex_subquery_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ex_subquery_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.EnterEx_subquery_range(s)
	}
}

func (s *Ex_subquery_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExtensionParserListener); ok {
		listenerT.ExitEx_subquery_range(s)
	}
}

func (p *PromQLExtensionParser) Ex_subquery_range() (localctx IEx_subquery_rangeContext) {
	localctx = NewEx_subquery_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromQLExtensionParserRULE_ex_subquery_range)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(PromQLExtensionParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Ex_duration()
	}
	{
		p.SetState(313)
		p.Match(PromQLExtensionParserEX_COLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2269950345478170) != 0 {
		{
			p.SetState(314)
			p.Ex_duration()
		}

	}
	{
		p.SetState(317)
		p.Match(PromQLExtensionParserRIGHT_BRACKET)
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
	Ex_alias_call() IEx_alias_callContext
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

func (s *VectorOperationContext) Ex_alias_call() IEx_alias_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_alias_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_alias_callContext)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, PromQLExtensionParserRULE_vectorOperation, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(320)
			p.UnaryOp()
		}
		{
			p.SetState(321)
			p.vectorOperation(10)
		}

	case 2:
		{
			p.SetState(323)
			p.Vector()
		}

	case 3:
		{
			p.SetState(324)
			p.Ex_alias_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(328)
					p.PowOp()
				}
				{
					p.SetState(329)
					p.vectorOperation(12)
				}

			case 2:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(332)
					p.MultOp()
				}
				{
					p.SetState(333)
					p.vectorOperation(10)
				}

			case 3:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(335)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(336)
					p.AddOp()
				}
				{
					p.SetState(337)
					p.vectorOperation(9)
				}

			case 4:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(340)
					p.CompareOp()
				}
				{
					p.SetState(341)
					p.vectorOperation(8)
				}

			case 5:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(343)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(344)
					p.AndUnlessOp()
				}
				{
					p.SetState(345)
					p.vectorOperation(7)
				}

			case 6:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(348)
					p.OrOp()
				}
				{
					p.SetState(349)
					p.vectorOperation(6)
				}

			case 7:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(352)
					p.VectorMatchOp()
				}
				{
					p.SetState(353)
					p.vectorOperation(5)
				}

			case 8:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(356)
					p.Match(PromQLExtensionParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(357)
					p.vectorOperation(4)
				}

			case 9:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExtensionParserRULE_vectorOperation)
				p.SetState(358)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(359)
					p.SubqueryOp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, PromQLExtensionParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
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
	p.EnterRule(localctx, 72, PromQLExtensionParserRULE_powOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(PromQLExtensionParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(368)
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
	p.EnterRule(localctx, 74, PromQLExtensionParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&224) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(372)
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
	p.EnterRule(localctx, 76, PromQLExtensionParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserADD || _la == PromQLExtensionParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(376)
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
	p.EnterRule(localctx, 78, PromQLExtensionParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516096) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserBOOL {
		{
			p.SetState(380)
			p.Match(PromQLExtensionParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(383)
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
	p.EnterRule(localctx, 80, PromQLExtensionParserRULE_andUnlessOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserAND || _la == PromQLExtensionParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(387)
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
	p.EnterRule(localctx, 82, PromQLExtensionParserRULE_orOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(PromQLExtensionParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(391)
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
	p.EnterRule(localctx, 84, PromQLExtensionParserRULE_vectorMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(394)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExtensionParserUNLESS || _la == PromQLExtensionParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserON || _la == PromQLExtensionParserIGNORING {
		{
			p.SetState(395)
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
	Ex_subquery_range() IEx_subquery_rangeContext
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

func (s *SubqueryOpContext) Ex_subquery_range() IEx_subquery_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_subquery_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_subquery_rangeContext)
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
	p.EnterRule(localctx, 86, PromQLExtensionParserRULE_subqueryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(398)
		p.Ex_subquery_range()
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(399)
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
	Ex_duration() IEx_durationContext

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

func (s *OffsetOpContext) Ex_duration() IEx_durationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_durationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_durationContext)
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
	p.EnterRule(localctx, 88, PromQLExtensionParserRULE_offsetOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(PromQLExtensionParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Ex_duration()
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
	p.EnterRule(localctx, 90, PromQLExtensionParserRULE_vector)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(408)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(409)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(410)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(411)
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
	p.EnterRule(localctx, 92, PromQLExtensionParserRULE_parens)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.vectorOperation(0)
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
	p.EnterRule(localctx, 94, PromQLExtensionParserRULE_instantSelector)
	var _la int

	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Match(PromQLExtensionParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(419)
				p.Match(PromQLExtensionParserLEFT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3300680273408) != 0 {
				{
					p.SetState(420)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(423)
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
			p.SetState(426)
			p.Match(PromQLExtensionParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.LabelMatcherList()
		}
		{
			p.SetState(428)
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
	p.EnterRule(localctx, 96, PromQLExtensionParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.LabelName()
	}
	{
		p.SetState(433)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 98, PromQLExtensionParserRULE_labelMatcherOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1593344) != 0) {
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
	p.EnterRule(localctx, 100, PromQLExtensionParserRULE_labelMatcherList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.LabelMatcher()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(439)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(440)
				p.LabelMatcher()
			}

		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExtensionParserCOMMA {
		{
			p.SetState(446)
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
	Ex_time_range() IEx_time_rangeContext

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

func (s *MatrixSelectorContext) Ex_time_range() IEx_time_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_time_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_time_rangeContext)
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
	p.EnterRule(localctx, 102, PromQLExtensionParserRULE_matrixSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.InstantSelector()
	}
	{
		p.SetState(450)
		p.Ex_time_range()
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
	Ex_duration() IEx_durationContext
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

func (s *OffsetContext) Ex_duration() IEx_durationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_durationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_durationContext)
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
	p.EnterRule(localctx, 104, PromQLExtensionParserRULE_offset)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.InstantSelector()
		}
		{
			p.SetState(453)
			p.Match(PromQLExtensionParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Ex_duration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
			p.MatrixSelector()
		}
		{
			p.SetState(457)
			p.Match(PromQLExtensionParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)
			p.Ex_duration()
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
	p.EnterRule(localctx, 106, PromQLExtensionParserRULE_function_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
		p.Match(PromQLExtensionParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(463)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2271053615202334) != 0 {
		{
			p.SetState(464)
			p.Parameter()
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(465)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(466)
				p.Parameter()
			}

			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(474)
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
	p.EnterRule(localctx, 108, PromQLExtensionParserRULE_parameter)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
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
	p.EnterRule(localctx, 110, PromQLExtensionParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2271053615202334) != 0 {
		{
			p.SetState(481)
			p.Parameter()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(482)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(483)
				p.Parameter()
			}

			p.SetState(488)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(491)
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
	p.EnterRule(localctx, 112, PromQLExtensionParserRULE_aggregation)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(493)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(495)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExtensionParserBY:
			{
				p.SetState(496)
				p.By()
			}

		case PromQLExtensionParserWITHOUT:
			{
				p.SetState(497)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(500)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(502)
			p.Match(PromQLExtensionParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(503)
			p.ParameterList()
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExtensionParserBY:
			{
				p.SetState(504)
				p.By()
			}

		case PromQLExtensionParserWITHOUT:
			{
				p.SetState(505)
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
	p.EnterRule(localctx, 114, PromQLExtensionParserRULE_by)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Match(PromQLExtensionParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
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
	p.EnterRule(localctx, 116, PromQLExtensionParserRULE_without)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(PromQLExtensionParserWITHOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
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
	p.EnterRule(localctx, 118, PromQLExtensionParserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserON:
		{
			p.SetState(516)
			p.On_()
		}

	case PromQLExtensionParserIGNORING:
		{
			p.SetState(517)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserGROUP_LEFT:
		{
			p.SetState(520)
			p.GroupLeft()
		}

	case PromQLExtensionParserGROUP_RIGHT:
		{
			p.SetState(521)
			p.GroupRight()
		}

	case PromQLExtensionParserNUMBER, PromQLExtensionParserSTRING, PromQLExtensionParserADD, PromQLExtensionParserSUB, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION, PromQLExtensionParserLEFT_BRACE, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserDURATION, PromQLExtensionParserMETRIC_NAME, PromQLExtensionParserEX_ID, PromQLExtensionParserEX_POSITIVE_INTEGER:

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
	p.EnterRule(localctx, 120, PromQLExtensionParserRULE_on_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(PromQLExtensionParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
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
	p.EnterRule(localctx, 122, PromQLExtensionParserRULE_ignoring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(527)
		p.Match(PromQLExtensionParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
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
	p.EnterRule(localctx, 124, PromQLExtensionParserRULE_groupLeft)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Match(PromQLExtensionParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(531)
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
	p.EnterRule(localctx, 126, PromQLExtensionParserRULE_groupRight)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(PromQLExtensionParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(536)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(535)
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
	p.EnterRule(localctx, 128, PromQLExtensionParserRULE_labelName)
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserAND, PromQLExtensionParserOR, PromQLExtensionParserUNLESS, PromQLExtensionParserBY, PromQLExtensionParserWITHOUT, PromQLExtensionParserON, PromQLExtensionParserIGNORING, PromQLExtensionParserGROUP_LEFT, PromQLExtensionParserGROUP_RIGHT, PromQLExtensionParserOFFSET, PromQLExtensionParserBOOL, PromQLExtensionParserAGGREGATION_OPERATOR, PromQLExtensionParserFUNCTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Keyword()
		}

	case PromQLExtensionParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Match(PromQLExtensionParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExtensionParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)
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
	p.EnterRule(localctx, 130, PromQLExtensionParserRULE_labelNameList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(543)
		p.Match(PromQLExtensionParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3300680273408) != 0 {
		{
			p.SetState(544)
			p.LabelName()
		}
		p.SetState(549)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExtensionParserCOMMA {
			{
				p.SetState(545)
				p.Match(PromQLExtensionParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(546)
				p.LabelName()
			}

			p.SetState(551)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(554)
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
	p.EnterRule(localctx, 132, PromQLExtensionParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(556)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2145390080) != 0) {
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
	Ex_const_num_expression() IEx_const_num_expressionContext
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

func (s *LiteralContext) Ex_const_num_expression() IEx_const_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEx_const_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEx_const_num_expressionContext)
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
	p.EnterRule(localctx, 134, PromQLExtensionParserRULE_literal)
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExtensionParserNUMBER, PromQLExtensionParserADD, PromQLExtensionParserSUB, PromQLExtensionParserLEFT_PAREN, PromQLExtensionParserDURATION, PromQLExtensionParserEX_ID, PromQLExtensionParserEX_POSITIVE_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(558)
			p.Ex_const_num_expression()
		}

	case PromQLExtensionParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(559)
			p.Match(PromQLExtensionParserSTRING)
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

func (p *PromQLExtensionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 34:
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
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

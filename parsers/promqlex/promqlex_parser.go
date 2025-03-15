// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
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

type PromQLExParser struct {
	*antlr.BaseParser
}

var PromQLExParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promqlexparserParserInit() {
	staticData := &PromQLExParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'if'", "'true'", "'false'", "'T'", "':'", "'.'", "",
		"", "", "'metric'", "'label'", "'def'", "'$'", "", "'['", "']'", "",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'", "'unless'",
		"'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'", "'!~'",
		"'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
		"'offset'", "'bool'", "'{'", "'}'", "'('", "')'", "','", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "AGGREGATION_OPERATOR", "FUNCTION", "METRIC_NAME", "IF", "TRUE",
		"FALSE", "T", "COLON", "DOT", "POSITIVE_INTEGER", "TWO_DIGITS", "DIGITS",
		"METRIC_KEYWORD", "LABEL_KEYWORD", "DEF", "CALL_SIGN", "NL", "LEFT_BRACKET",
		"RIGHT_BRACKET", "NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD",
		"POW", "AND", "OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE",
		"RE", "NRE", "BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT",
		"OFFSET", "BOOL", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN", "RIGHT_PAREN",
		"COMMA", "AT", "DURATION", "LABEL_NAME", "WS", "SL_COMMENT", "ID", "RAW_STRING",
		"BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"promqlx", "statement", "alias_def", "alias_call", "macro_def", "macro_call",
		"args_decl", "arg_name", "statement_block", "arg_list", "if_statement",
		"condition", "compareVectorOperation", "trueConst", "falseConst", "time_instant_literal",
		"iso_date_time", "iso_date_time_ymdhmsf", "iso_date_time_ymdhms", "iso_date_time_ymdhm",
		"iso_date_time_ymdh", "iso_date_time_ymd", "iso_date_time_ym", "iso_date_time_y",
		"iso_year", "iso_month", "iso_day", "iso_hour", "iso_minutes", "iso_seconds",
		"is_frac_sec", "unix_timestamp", "const_num_expression", "num_literal",
		"duration", "time_range", "subquery_range", "vectorOperation", "subqueryOp",
		"offsetOp", "matrixSelector", "offset", "literal", "instantSelector",
		"labelName", "metric_name", "expression", "unaryOp", "powOp", "multOp",
		"addOp", "compareOp", "andUnlessOp", "orOp", "vectorMatchOp", "vector",
		"parens", "labelMatcher", "labelMatcherOperator", "labelMatcherList",
		"function_", "parameter", "parameterList", "aggregation", "by", "without",
		"grouping", "on_", "ignoring", "groupLeft", "groupRight", "labelNameList",
		"keyword", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 60, 607, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 1, 0, 1, 0, 1, 0, 5, 0, 152, 8, 0, 10, 0, 12, 0, 155, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 163, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 177, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 187, 8, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 6, 5, 6, 194, 8, 6, 10, 6, 12, 6, 197, 9, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 4, 9, 208, 8, 9, 11, 9, 12, 9, 209,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 219, 8, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15,
		231, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 240,
		8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 334,
		8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 340, 8, 33, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 352, 8, 36,
		1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 363,
		8, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 5, 37, 398, 8, 37, 10, 37, 12, 37, 401, 9, 37, 1, 38,
		1, 38, 3, 38, 405, 8, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 421, 8, 41,
		1, 42, 1, 42, 3, 42, 425, 8, 42, 1, 43, 1, 43, 1, 43, 3, 43, 430, 8, 43,
		1, 43, 3, 43, 433, 8, 43, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43, 439, 8, 43,
		1, 44, 1, 44, 1, 44, 3, 44, 444, 8, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1,
		46, 1, 47, 1, 47, 1, 48, 1, 48, 3, 48, 455, 8, 48, 1, 49, 1, 49, 3, 49,
		459, 8, 49, 1, 50, 1, 50, 3, 50, 463, 8, 50, 1, 51, 1, 51, 3, 51, 467,
		8, 51, 1, 51, 3, 51, 470, 8, 51, 1, 52, 1, 52, 3, 52, 474, 8, 52, 1, 53,
		1, 53, 3, 53, 478, 8, 53, 1, 54, 1, 54, 3, 54, 482, 8, 54, 1, 55, 1, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 491, 8, 55, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59,
		5, 59, 506, 8, 59, 10, 59, 12, 59, 509, 9, 59, 1, 59, 3, 59, 512, 8, 59,
		1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 519, 8, 60, 10, 60, 12, 60, 522,
		9, 60, 3, 60, 524, 8, 60, 1, 60, 1, 60, 1, 61, 1, 61, 3, 61, 530, 8, 61,
		1, 62, 1, 62, 1, 62, 1, 62, 5, 62, 536, 8, 62, 10, 62, 12, 62, 539, 9,
		62, 3, 62, 541, 8, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63,
		3, 63, 550, 8, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 1, 63, 3, 63, 558,
		8, 63, 3, 63, 560, 8, 63, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1,
		66, 1, 66, 3, 66, 570, 8, 66, 1, 66, 1, 66, 3, 66, 574, 8, 66, 1, 67, 1,
		67, 1, 67, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 3, 69, 584, 8, 69, 1, 70,
		1, 70, 3, 70, 588, 8, 70, 1, 71, 1, 71, 1, 71, 1, 71, 5, 71, 594, 8, 71,
		10, 71, 12, 71, 597, 9, 71, 3, 71, 599, 8, 71, 1, 71, 1, 71, 1, 72, 1,
		72, 1, 73, 1, 73, 1, 73, 0, 1, 74, 74, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
		92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120,
		122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 0, 9,
		2, 0, 3, 3, 58, 58, 1, 0, 22, 23, 1, 0, 24, 26, 1, 0, 32, 37, 2, 0, 28,
		28, 30, 30, 2, 0, 30, 30, 42, 42, 3, 0, 31, 31, 33, 33, 38, 39, 3, 0, 1,
		2, 28, 30, 40, 47, 2, 0, 21, 21, 59, 59, 610, 0, 148, 1, 0, 0, 0, 2, 162,
		1, 0, 0, 0, 4, 164, 1, 0, 0, 0, 6, 169, 1, 0, 0, 0, 8, 172, 1, 0, 0, 0,
		10, 181, 1, 0, 0, 0, 12, 190, 1, 0, 0, 0, 14, 198, 1, 0, 0, 0, 16, 200,
		1, 0, 0, 0, 18, 204, 1, 0, 0, 0, 20, 211, 1, 0, 0, 0, 22, 218, 1, 0, 0,
		0, 24, 220, 1, 0, 0, 0, 26, 224, 1, 0, 0, 0, 28, 226, 1, 0, 0, 0, 30, 230,
		1, 0, 0, 0, 32, 239, 1, 0, 0, 0, 34, 241, 1, 0, 0, 0, 36, 255, 1, 0, 0,
		0, 38, 267, 1, 0, 0, 0, 40, 277, 1, 0, 0, 0, 42, 285, 1, 0, 0, 0, 44, 291,
		1, 0, 0, 0, 46, 295, 1, 0, 0, 0, 48, 297, 1, 0, 0, 0, 50, 299, 1, 0, 0,
		0, 52, 301, 1, 0, 0, 0, 54, 303, 1, 0, 0, 0, 56, 305, 1, 0, 0, 0, 58, 307,
		1, 0, 0, 0, 60, 309, 1, 0, 0, 0, 62, 311, 1, 0, 0, 0, 64, 333, 1, 0, 0,
		0, 66, 339, 1, 0, 0, 0, 68, 341, 1, 0, 0, 0, 70, 343, 1, 0, 0, 0, 72, 347,
		1, 0, 0, 0, 74, 362, 1, 0, 0, 0, 76, 402, 1, 0, 0, 0, 78, 406, 1, 0, 0,
		0, 80, 409, 1, 0, 0, 0, 82, 420, 1, 0, 0, 0, 84, 424, 1, 0, 0, 0, 86, 438,
		1, 0, 0, 0, 88, 443, 1, 0, 0, 0, 90, 445, 1, 0, 0, 0, 92, 447, 1, 0, 0,
		0, 94, 450, 1, 0, 0, 0, 96, 452, 1, 0, 0, 0, 98, 456, 1, 0, 0, 0, 100,
		460, 1, 0, 0, 0, 102, 464, 1, 0, 0, 0, 104, 471, 1, 0, 0, 0, 106, 475,
		1, 0, 0, 0, 108, 479, 1, 0, 0, 0, 110, 490, 1, 0, 0, 0, 112, 492, 1, 0,
		0, 0, 114, 496, 1, 0, 0, 0, 116, 500, 1, 0, 0, 0, 118, 502, 1, 0, 0, 0,
		120, 513, 1, 0, 0, 0, 122, 529, 1, 0, 0, 0, 124, 531, 1, 0, 0, 0, 126,
		559, 1, 0, 0, 0, 128, 561, 1, 0, 0, 0, 130, 564, 1, 0, 0, 0, 132, 569,
		1, 0, 0, 0, 134, 575, 1, 0, 0, 0, 136, 578, 1, 0, 0, 0, 138, 581, 1, 0,
		0, 0, 140, 585, 1, 0, 0, 0, 142, 589, 1, 0, 0, 0, 144, 602, 1, 0, 0, 0,
		146, 604, 1, 0, 0, 0, 148, 153, 3, 2, 1, 0, 149, 150, 5, 17, 0, 0, 150,
		152, 3, 2, 1, 0, 151, 149, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151,
		1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 156, 157, 5, 0, 0, 1, 157, 1, 1, 0, 0, 0, 158, 163, 3, 4, 2, 0, 159,
		163, 3, 8, 4, 0, 160, 163, 3, 20, 10, 0, 161, 163, 3, 74, 37, 0, 162, 158,
		1, 0, 0, 0, 162, 159, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0,
		0, 0, 163, 3, 1, 0, 0, 0, 164, 165, 5, 15, 0, 0, 165, 166, 5, 58, 0, 0,
		166, 167, 5, 31, 0, 0, 167, 168, 3, 74, 37, 0, 168, 5, 1, 0, 0, 0, 169,
		170, 5, 16, 0, 0, 170, 171, 5, 58, 0, 0, 171, 7, 1, 0, 0, 0, 172, 173,
		5, 15, 0, 0, 173, 174, 5, 58, 0, 0, 174, 176, 5, 50, 0, 0, 175, 177, 3,
		12, 6, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0,
		0, 178, 179, 5, 51, 0, 0, 179, 180, 3, 16, 8, 0, 180, 9, 1, 0, 0, 0, 181,
		182, 5, 16, 0, 0, 182, 183, 5, 58, 0, 0, 183, 184, 5, 58, 0, 0, 184, 186,
		5, 50, 0, 0, 185, 187, 3, 18, 9, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1,
		0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 5, 51, 0, 0, 189, 11, 1, 0, 0,
		0, 190, 195, 3, 14, 7, 0, 191, 192, 5, 52, 0, 0, 192, 194, 3, 14, 7, 0,
		193, 191, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 13, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 199, 5,
		58, 0, 0, 199, 15, 1, 0, 0, 0, 200, 201, 5, 48, 0, 0, 201, 202, 3, 0, 0,
		0, 202, 203, 5, 49, 0, 0, 203, 17, 1, 0, 0, 0, 204, 207, 3, 74, 37, 0,
		205, 206, 5, 52, 0, 0, 206, 208, 3, 74, 37, 0, 207, 205, 1, 0, 0, 0, 208,
		209, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 19, 1,
		0, 0, 0, 211, 212, 5, 4, 0, 0, 212, 213, 3, 22, 11, 0, 213, 214, 3, 16,
		8, 0, 214, 21, 1, 0, 0, 0, 215, 219, 3, 24, 12, 0, 216, 219, 3, 26, 13,
		0, 217, 219, 3, 28, 14, 0, 218, 215, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0,
		218, 217, 1, 0, 0, 0, 219, 23, 1, 0, 0, 0, 220, 221, 3, 74, 37, 0, 221,
		222, 3, 102, 51, 0, 222, 223, 3, 74, 37, 0, 223, 25, 1, 0, 0, 0, 224, 225,
		5, 5, 0, 0, 225, 27, 1, 0, 0, 0, 226, 227, 5, 6, 0, 0, 227, 29, 1, 0, 0,
		0, 228, 231, 3, 32, 16, 0, 229, 231, 3, 62, 31, 0, 230, 228, 1, 0, 0, 0,
		230, 229, 1, 0, 0, 0, 231, 31, 1, 0, 0, 0, 232, 240, 3, 34, 17, 0, 233,
		240, 3, 36, 18, 0, 234, 240, 3, 38, 19, 0, 235, 240, 3, 40, 20, 0, 236,
		240, 3, 42, 21, 0, 237, 240, 3, 44, 22, 0, 238, 240, 3, 46, 23, 0, 239,
		232, 1, 0, 0, 0, 239, 233, 1, 0, 0, 0, 239, 234, 1, 0, 0, 0, 239, 235,
		1, 0, 0, 0, 239, 236, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 238, 1, 0,
		0, 0, 240, 33, 1, 0, 0, 0, 241, 242, 3, 48, 24, 0, 242, 243, 5, 23, 0,
		0, 243, 244, 3, 50, 25, 0, 244, 245, 5, 23, 0, 0, 245, 246, 3, 52, 26,
		0, 246, 247, 5, 7, 0, 0, 247, 248, 3, 54, 27, 0, 248, 249, 5, 8, 0, 0,
		249, 250, 3, 56, 28, 0, 250, 251, 5, 8, 0, 0, 251, 252, 3, 58, 29, 0, 252,
		253, 5, 9, 0, 0, 253, 254, 3, 60, 30, 0, 254, 35, 1, 0, 0, 0, 255, 256,
		3, 48, 24, 0, 256, 257, 5, 23, 0, 0, 257, 258, 3, 50, 25, 0, 258, 259,
		5, 23, 0, 0, 259, 260, 3, 52, 26, 0, 260, 261, 5, 7, 0, 0, 261, 262, 3,
		54, 27, 0, 262, 263, 5, 8, 0, 0, 263, 264, 3, 56, 28, 0, 264, 265, 5, 8,
		0, 0, 265, 266, 3, 58, 29, 0, 266, 37, 1, 0, 0, 0, 267, 268, 3, 48, 24,
		0, 268, 269, 5, 23, 0, 0, 269, 270, 3, 50, 25, 0, 270, 271, 5, 23, 0, 0,
		271, 272, 3, 52, 26, 0, 272, 273, 5, 7, 0, 0, 273, 274, 3, 54, 27, 0, 274,
		275, 5, 8, 0, 0, 275, 276, 3, 56, 28, 0, 276, 39, 1, 0, 0, 0, 277, 278,
		3, 48, 24, 0, 278, 279, 5, 23, 0, 0, 279, 280, 3, 50, 25, 0, 280, 281,
		5, 23, 0, 0, 281, 282, 3, 52, 26, 0, 282, 283, 5, 7, 0, 0, 283, 284, 3,
		54, 27, 0, 284, 41, 1, 0, 0, 0, 285, 286, 3, 48, 24, 0, 286, 287, 5, 23,
		0, 0, 287, 288, 3, 50, 25, 0, 288, 289, 5, 23, 0, 0, 289, 290, 3, 52, 26,
		0, 290, 43, 1, 0, 0, 0, 291, 292, 3, 48, 24, 0, 292, 293, 5, 23, 0, 0,
		293, 294, 3, 50, 25, 0, 294, 45, 1, 0, 0, 0, 295, 296, 3, 48, 24, 0, 296,
		47, 1, 0, 0, 0, 297, 298, 5, 10, 0, 0, 298, 49, 1, 0, 0, 0, 299, 300, 5,
		11, 0, 0, 300, 51, 1, 0, 0, 0, 301, 302, 5, 11, 0, 0, 302, 53, 1, 0, 0,
		0, 303, 304, 5, 11, 0, 0, 304, 55, 1, 0, 0, 0, 305, 306, 5, 11, 0, 0, 306,
		57, 1, 0, 0, 0, 307, 308, 5, 11, 0, 0, 308, 59, 1, 0, 0, 0, 309, 310, 5,
		12, 0, 0, 310, 61, 1, 0, 0, 0, 311, 312, 5, 10, 0, 0, 312, 63, 1, 0, 0,
		0, 313, 314, 3, 66, 33, 0, 314, 315, 3, 96, 48, 0, 315, 316, 3, 66, 33,
		0, 316, 334, 1, 0, 0, 0, 317, 318, 3, 94, 47, 0, 318, 319, 3, 66, 33, 0,
		319, 334, 1, 0, 0, 0, 320, 321, 3, 66, 33, 0, 321, 322, 3, 98, 49, 0, 322,
		323, 3, 66, 33, 0, 323, 334, 1, 0, 0, 0, 324, 325, 3, 66, 33, 0, 325, 326,
		3, 100, 50, 0, 326, 327, 3, 66, 33, 0, 327, 334, 1, 0, 0, 0, 328, 329,
		5, 50, 0, 0, 329, 330, 3, 64, 32, 0, 330, 331, 5, 51, 0, 0, 331, 334, 1,
		0, 0, 0, 332, 334, 3, 66, 33, 0, 333, 313, 1, 0, 0, 0, 333, 317, 1, 0,
		0, 0, 333, 320, 1, 0, 0, 0, 333, 324, 1, 0, 0, 0, 333, 328, 1, 0, 0, 0,
		333, 332, 1, 0, 0, 0, 334, 65, 1, 0, 0, 0, 335, 340, 5, 20, 0, 0, 336,
		340, 5, 54, 0, 0, 337, 340, 3, 30, 15, 0, 338, 340, 3, 6, 3, 0, 339, 335,
		1, 0, 0, 0, 339, 336, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 338, 1, 0,
		0, 0, 340, 67, 1, 0, 0, 0, 341, 342, 3, 64, 32, 0, 342, 69, 1, 0, 0, 0,
		343, 344, 5, 18, 0, 0, 344, 345, 3, 68, 34, 0, 345, 346, 5, 19, 0, 0, 346,
		71, 1, 0, 0, 0, 347, 348, 5, 18, 0, 0, 348, 349, 3, 68, 34, 0, 349, 351,
		5, 8, 0, 0, 350, 352, 3, 68, 34, 0, 351, 350, 1, 0, 0, 0, 351, 352, 1,
		0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 5, 19, 0, 0, 354, 73, 1, 0, 0,
		0, 355, 356, 6, 37, -1, 0, 356, 357, 3, 94, 47, 0, 357, 358, 3, 74, 37,
		11, 358, 363, 1, 0, 0, 0, 359, 363, 3, 110, 55, 0, 360, 363, 3, 10, 5,
		0, 361, 363, 3, 6, 3, 0, 362, 355, 1, 0, 0, 0, 362, 359, 1, 0, 0, 0, 362,
		360, 1, 0, 0, 0, 362, 361, 1, 0, 0, 0, 363, 399, 1, 0, 0, 0, 364, 365,
		10, 13, 0, 0, 365, 366, 3, 96, 48, 0, 366, 367, 3, 74, 37, 13, 367, 398,
		1, 0, 0, 0, 368, 369, 10, 10, 0, 0, 369, 370, 3, 98, 49, 0, 370, 371, 3,
		74, 37, 11, 371, 398, 1, 0, 0, 0, 372, 373, 10, 9, 0, 0, 373, 374, 3, 100,
		50, 0, 374, 375, 3, 74, 37, 10, 375, 398, 1, 0, 0, 0, 376, 377, 10, 8,
		0, 0, 377, 378, 3, 102, 51, 0, 378, 379, 3, 74, 37, 9, 379, 398, 1, 0,
		0, 0, 380, 381, 10, 7, 0, 0, 381, 382, 3, 104, 52, 0, 382, 383, 3, 74,
		37, 8, 383, 398, 1, 0, 0, 0, 384, 385, 10, 6, 0, 0, 385, 386, 3, 106, 53,
		0, 386, 387, 3, 74, 37, 7, 387, 398, 1, 0, 0, 0, 388, 389, 10, 5, 0, 0,
		389, 390, 3, 108, 54, 0, 390, 391, 3, 74, 37, 6, 391, 398, 1, 0, 0, 0,
		392, 393, 10, 4, 0, 0, 393, 394, 5, 53, 0, 0, 394, 398, 3, 74, 37, 5, 395,
		396, 10, 12, 0, 0, 396, 398, 3, 76, 38, 0, 397, 364, 1, 0, 0, 0, 397, 368,
		1, 0, 0, 0, 397, 372, 1, 0, 0, 0, 397, 376, 1, 0, 0, 0, 397, 380, 1, 0,
		0, 0, 397, 384, 1, 0, 0, 0, 397, 388, 1, 0, 0, 0, 397, 392, 1, 0, 0, 0,
		397, 395, 1, 0, 0, 0, 398, 401, 1, 0, 0, 0, 399, 397, 1, 0, 0, 0, 399,
		400, 1, 0, 0, 0, 400, 75, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 402, 404, 3,
		72, 36, 0, 403, 405, 3, 78, 39, 0, 404, 403, 1, 0, 0, 0, 404, 405, 1, 0,
		0, 0, 405, 77, 1, 0, 0, 0, 406, 407, 5, 46, 0, 0, 407, 408, 3, 68, 34,
		0, 408, 79, 1, 0, 0, 0, 409, 410, 3, 86, 43, 0, 410, 411, 3, 70, 35, 0,
		411, 81, 1, 0, 0, 0, 412, 413, 3, 86, 43, 0, 413, 414, 5, 46, 0, 0, 414,
		415, 3, 68, 34, 0, 415, 421, 1, 0, 0, 0, 416, 417, 3, 80, 40, 0, 417, 418,
		5, 46, 0, 0, 418, 419, 3, 68, 34, 0, 419, 421, 1, 0, 0, 0, 420, 412, 1,
		0, 0, 0, 420, 416, 1, 0, 0, 0, 421, 83, 1, 0, 0, 0, 422, 425, 3, 64, 32,
		0, 423, 425, 3, 146, 73, 0, 424, 422, 1, 0, 0, 0, 424, 423, 1, 0, 0, 0,
		425, 85, 1, 0, 0, 0, 426, 432, 3, 90, 45, 0, 427, 429, 5, 48, 0, 0, 428,
		430, 3, 118, 59, 0, 429, 428, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 431,
		1, 0, 0, 0, 431, 433, 5, 49, 0, 0, 432, 427, 1, 0, 0, 0, 432, 433, 1, 0,
		0, 0, 433, 439, 1, 0, 0, 0, 434, 435, 5, 48, 0, 0, 435, 436, 3, 118, 59,
		0, 436, 437, 5, 49, 0, 0, 437, 439, 1, 0, 0, 0, 438, 426, 1, 0, 0, 0, 438,
		434, 1, 0, 0, 0, 439, 87, 1, 0, 0, 0, 440, 444, 3, 144, 72, 0, 441, 444,
		3, 90, 45, 0, 442, 444, 5, 55, 0, 0, 443, 440, 1, 0, 0, 0, 443, 441, 1,
		0, 0, 0, 443, 442, 1, 0, 0, 0, 444, 89, 1, 0, 0, 0, 445, 446, 7, 0, 0,
		0, 446, 91, 1, 0, 0, 0, 447, 448, 3, 74, 37, 0, 448, 449, 5, 0, 0, 1, 449,
		93, 1, 0, 0, 0, 450, 451, 7, 1, 0, 0, 451, 95, 1, 0, 0, 0, 452, 454, 5,
		27, 0, 0, 453, 455, 3, 132, 66, 0, 454, 453, 1, 0, 0, 0, 454, 455, 1, 0,
		0, 0, 455, 97, 1, 0, 0, 0, 456, 458, 7, 2, 0, 0, 457, 459, 3, 132, 66,
		0, 458, 457, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 99, 1, 0, 0, 0, 460,
		462, 7, 1, 0, 0, 461, 463, 3, 132, 66, 0, 462, 461, 1, 0, 0, 0, 462, 463,
		1, 0, 0, 0, 463, 101, 1, 0, 0, 0, 464, 466, 7, 3, 0, 0, 465, 467, 5, 47,
		0, 0, 466, 465, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 469, 1, 0, 0, 0,
		468, 470, 3, 132, 66, 0, 469, 468, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470,
		103, 1, 0, 0, 0, 471, 473, 7, 4, 0, 0, 472, 474, 3, 132, 66, 0, 473, 472,
		1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 105, 1, 0, 0, 0, 475, 477, 5, 29,
		0, 0, 476, 478, 3, 132, 66, 0, 477, 476, 1, 0, 0, 0, 477, 478, 1, 0, 0,
		0, 478, 107, 1, 0, 0, 0, 479, 481, 7, 5, 0, 0, 480, 482, 3, 132, 66, 0,
		481, 480, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 109, 1, 0, 0, 0, 483,
		491, 3, 120, 60, 0, 484, 491, 3, 126, 63, 0, 485, 491, 3, 86, 43, 0, 486,
		491, 3, 80, 40, 0, 487, 491, 3, 82, 41, 0, 488, 491, 3, 84, 42, 0, 489,
		491, 3, 112, 56, 0, 490, 483, 1, 0, 0, 0, 490, 484, 1, 0, 0, 0, 490, 485,
		1, 0, 0, 0, 490, 486, 1, 0, 0, 0, 490, 487, 1, 0, 0, 0, 490, 488, 1, 0,
		0, 0, 490, 489, 1, 0, 0, 0, 491, 111, 1, 0, 0, 0, 492, 493, 5, 50, 0, 0,
		493, 494, 3, 74, 37, 0, 494, 495, 5, 51, 0, 0, 495, 113, 1, 0, 0, 0, 496,
		497, 3, 88, 44, 0, 497, 498, 3, 116, 58, 0, 498, 499, 3, 146, 73, 0, 499,
		115, 1, 0, 0, 0, 500, 501, 7, 6, 0, 0, 501, 117, 1, 0, 0, 0, 502, 507,
		3, 114, 57, 0, 503, 504, 5, 52, 0, 0, 504, 506, 3, 114, 57, 0, 505, 503,
		1, 0, 0, 0, 506, 509, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 507, 508, 1, 0,
		0, 0, 508, 511, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 510, 512, 5, 52, 0, 0,
		511, 510, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512, 119, 1, 0, 0, 0, 513,
		514, 5, 2, 0, 0, 514, 523, 5, 50, 0, 0, 515, 520, 3, 122, 61, 0, 516, 517,
		5, 52, 0, 0, 517, 519, 3, 122, 61, 0, 518, 516, 1, 0, 0, 0, 519, 522, 1,
		0, 0, 0, 520, 518, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 521, 524, 1, 0, 0,
		0, 522, 520, 1, 0, 0, 0, 523, 515, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0, 524,
		525, 1, 0, 0, 0, 525, 526, 5, 51, 0, 0, 526, 121, 1, 0, 0, 0, 527, 530,
		3, 84, 42, 0, 528, 530, 3, 74, 37, 0, 529, 527, 1, 0, 0, 0, 529, 528, 1,
		0, 0, 0, 530, 123, 1, 0, 0, 0, 531, 540, 5, 50, 0, 0, 532, 537, 3, 122,
		61, 0, 533, 534, 5, 52, 0, 0, 534, 536, 3, 122, 61, 0, 535, 533, 1, 0,
		0, 0, 536, 539, 1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0,
		538, 541, 1, 0, 0, 0, 539, 537, 1, 0, 0, 0, 540, 532, 1, 0, 0, 0, 540,
		541, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 543, 5, 51, 0, 0, 543, 125,
		1, 0, 0, 0, 544, 545, 5, 1, 0, 0, 545, 560, 3, 124, 62, 0, 546, 549, 5,
		1, 0, 0, 547, 550, 3, 128, 64, 0, 548, 550, 3, 130, 65, 0, 549, 547, 1,
		0, 0, 0, 549, 548, 1, 0, 0, 0, 550, 551, 1, 0, 0, 0, 551, 552, 3, 124,
		62, 0, 552, 560, 1, 0, 0, 0, 553, 554, 5, 1, 0, 0, 554, 557, 3, 124, 62,
		0, 555, 558, 3, 128, 64, 0, 556, 558, 3, 130, 65, 0, 557, 555, 1, 0, 0,
		0, 557, 556, 1, 0, 0, 0, 558, 560, 1, 0, 0, 0, 559, 544, 1, 0, 0, 0, 559,
		546, 1, 0, 0, 0, 559, 553, 1, 0, 0, 0, 560, 127, 1, 0, 0, 0, 561, 562,
		5, 40, 0, 0, 562, 563, 3, 142, 71, 0, 563, 129, 1, 0, 0, 0, 564, 565, 5,
		41, 0, 0, 565, 566, 3, 142, 71, 0, 566, 131, 1, 0, 0, 0, 567, 570, 3, 134,
		67, 0, 568, 570, 3, 136, 68, 0, 569, 567, 1, 0, 0, 0, 569, 568, 1, 0, 0,
		0, 570, 573, 1, 0, 0, 0, 571, 574, 3, 138, 69, 0, 572, 574, 3, 140, 70,
		0, 573, 571, 1, 0, 0, 0, 573, 572, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574,
		133, 1, 0, 0, 0, 575, 576, 5, 42, 0, 0, 576, 577, 3, 142, 71, 0, 577, 135,
		1, 0, 0, 0, 578, 579, 5, 43, 0, 0, 579, 580, 3, 142, 71, 0, 580, 137, 1,
		0, 0, 0, 581, 583, 5, 44, 0, 0, 582, 584, 3, 142, 71, 0, 583, 582, 1, 0,
		0, 0, 583, 584, 1, 0, 0, 0, 584, 139, 1, 0, 0, 0, 585, 587, 5, 45, 0, 0,
		586, 588, 3, 142, 71, 0, 587, 586, 1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588,
		141, 1, 0, 0, 0, 589, 598, 5, 50, 0, 0, 590, 595, 3, 88, 44, 0, 591, 592,
		5, 52, 0, 0, 592, 594, 3, 88, 44, 0, 593, 591, 1, 0, 0, 0, 594, 597, 1,
		0, 0, 0, 595, 593, 1, 0, 0, 0, 595, 596, 1, 0, 0, 0, 596, 599, 1, 0, 0,
		0, 597, 595, 1, 0, 0, 0, 598, 590, 1, 0, 0, 0, 598, 599, 1, 0, 0, 0, 599,
		600, 1, 0, 0, 0, 600, 601, 5, 51, 0, 0, 601, 143, 1, 0, 0, 0, 602, 603,
		7, 7, 0, 0, 603, 145, 1, 0, 0, 0, 604, 605, 7, 8, 0, 0, 605, 147, 1, 0,
		0, 0, 47, 153, 162, 176, 186, 195, 209, 218, 230, 239, 333, 339, 351, 362,
		397, 399, 404, 420, 424, 429, 432, 438, 443, 454, 458, 462, 466, 469, 473,
		477, 481, 490, 507, 511, 520, 523, 529, 537, 540, 549, 557, 559, 569, 573,
		583, 587, 595, 598,
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

// PromQLExParserInit initializes any static state used to implement PromQLExParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromQLExParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLExParserInit() {
	staticData := &PromQLExParserParserStaticData
	staticData.once.Do(promqlexparserParserInit)
}

// NewPromQLExParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromQLExParser(input antlr.TokenStream) *PromQLExParser {
	PromQLExParserInit()
	this := new(PromQLExParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromQLExParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromQLExParser.g4"

	return this
}

// PromQLExParser tokens.
const (
	PromQLExParserEOF                  = antlr.TokenEOF
	PromQLExParserAGGREGATION_OPERATOR = 1
	PromQLExParserFUNCTION             = 2
	PromQLExParserMETRIC_NAME          = 3
	PromQLExParserIF                   = 4
	PromQLExParserTRUE                 = 5
	PromQLExParserFALSE                = 6
	PromQLExParserT                    = 7
	PromQLExParserCOLON                = 8
	PromQLExParserDOT                  = 9
	PromQLExParserPOSITIVE_INTEGER     = 10
	PromQLExParserTWO_DIGITS           = 11
	PromQLExParserDIGITS               = 12
	PromQLExParserMETRIC_KEYWORD       = 13
	PromQLExParserLABEL_KEYWORD        = 14
	PromQLExParserDEF                  = 15
	PromQLExParserCALL_SIGN            = 16
	PromQLExParserNL                   = 17
	PromQLExParserLEFT_BRACKET         = 18
	PromQLExParserRIGHT_BRACKET        = 19
	PromQLExParserNUMBER               = 20
	PromQLExParserSTRING               = 21
	PromQLExParserADD                  = 22
	PromQLExParserSUB                  = 23
	PromQLExParserMULT                 = 24
	PromQLExParserDIV                  = 25
	PromQLExParserMOD                  = 26
	PromQLExParserPOW                  = 27
	PromQLExParserAND                  = 28
	PromQLExParserOR                   = 29
	PromQLExParserUNLESS               = 30
	PromQLExParserEQ                   = 31
	PromQLExParserDEQ                  = 32
	PromQLExParserNE                   = 33
	PromQLExParserGT                   = 34
	PromQLExParserLT                   = 35
	PromQLExParserGE                   = 36
	PromQLExParserLE                   = 37
	PromQLExParserRE                   = 38
	PromQLExParserNRE                  = 39
	PromQLExParserBY                   = 40
	PromQLExParserWITHOUT              = 41
	PromQLExParserON                   = 42
	PromQLExParserIGNORING             = 43
	PromQLExParserGROUP_LEFT           = 44
	PromQLExParserGROUP_RIGHT          = 45
	PromQLExParserOFFSET               = 46
	PromQLExParserBOOL                 = 47
	PromQLExParserLEFT_BRACE           = 48
	PromQLExParserRIGHT_BRACE          = 49
	PromQLExParserLEFT_PAREN           = 50
	PromQLExParserRIGHT_PAREN          = 51
	PromQLExParserCOMMA                = 52
	PromQLExParserAT                   = 53
	PromQLExParserDURATION             = 54
	PromQLExParserLABEL_NAME           = 55
	PromQLExParserWS                   = 56
	PromQLExParserSL_COMMENT           = 57
	PromQLExParserID                   = 58
	PromQLExParserRAW_STRING           = 59
	PromQLExParserBACKTICK_OPEN        = 60
)

// PromQLExParser rules.
const (
	PromQLExParserRULE_promqlx                = 0
	PromQLExParserRULE_statement              = 1
	PromQLExParserRULE_alias_def              = 2
	PromQLExParserRULE_alias_call             = 3
	PromQLExParserRULE_macro_def              = 4
	PromQLExParserRULE_macro_call             = 5
	PromQLExParserRULE_args_decl              = 6
	PromQLExParserRULE_arg_name               = 7
	PromQLExParserRULE_statement_block        = 8
	PromQLExParserRULE_arg_list               = 9
	PromQLExParserRULE_if_statement           = 10
	PromQLExParserRULE_condition              = 11
	PromQLExParserRULE_compareVectorOperation = 12
	PromQLExParserRULE_trueConst              = 13
	PromQLExParserRULE_falseConst             = 14
	PromQLExParserRULE_time_instant_literal   = 15
	PromQLExParserRULE_iso_date_time          = 16
	PromQLExParserRULE_iso_date_time_ymdhmsf  = 17
	PromQLExParserRULE_iso_date_time_ymdhms   = 18
	PromQLExParserRULE_iso_date_time_ymdhm    = 19
	PromQLExParserRULE_iso_date_time_ymdh     = 20
	PromQLExParserRULE_iso_date_time_ymd      = 21
	PromQLExParserRULE_iso_date_time_ym       = 22
	PromQLExParserRULE_iso_date_time_y        = 23
	PromQLExParserRULE_iso_year               = 24
	PromQLExParserRULE_iso_month              = 25
	PromQLExParserRULE_iso_day                = 26
	PromQLExParserRULE_iso_hour               = 27
	PromQLExParserRULE_iso_minutes            = 28
	PromQLExParserRULE_iso_seconds            = 29
	PromQLExParserRULE_is_frac_sec            = 30
	PromQLExParserRULE_unix_timestamp         = 31
	PromQLExParserRULE_const_num_expression   = 32
	PromQLExParserRULE_num_literal            = 33
	PromQLExParserRULE_duration               = 34
	PromQLExParserRULE_time_range             = 35
	PromQLExParserRULE_subquery_range         = 36
	PromQLExParserRULE_vectorOperation        = 37
	PromQLExParserRULE_subqueryOp             = 38
	PromQLExParserRULE_offsetOp               = 39
	PromQLExParserRULE_matrixSelector         = 40
	PromQLExParserRULE_offset                 = 41
	PromQLExParserRULE_literal                = 42
	PromQLExParserRULE_instantSelector        = 43
	PromQLExParserRULE_labelName              = 44
	PromQLExParserRULE_metric_name            = 45
	PromQLExParserRULE_expression             = 46
	PromQLExParserRULE_unaryOp                = 47
	PromQLExParserRULE_powOp                  = 48
	PromQLExParserRULE_multOp                 = 49
	PromQLExParserRULE_addOp                  = 50
	PromQLExParserRULE_compareOp              = 51
	PromQLExParserRULE_andUnlessOp            = 52
	PromQLExParserRULE_orOp                   = 53
	PromQLExParserRULE_vectorMatchOp          = 54
	PromQLExParserRULE_vector                 = 55
	PromQLExParserRULE_parens                 = 56
	PromQLExParserRULE_labelMatcher           = 57
	PromQLExParserRULE_labelMatcherOperator   = 58
	PromQLExParserRULE_labelMatcherList       = 59
	PromQLExParserRULE_function_              = 60
	PromQLExParserRULE_parameter              = 61
	PromQLExParserRULE_parameterList          = 62
	PromQLExParserRULE_aggregation            = 63
	PromQLExParserRULE_by                     = 64
	PromQLExParserRULE_without                = 65
	PromQLExParserRULE_grouping               = 66
	PromQLExParserRULE_on_                    = 67
	PromQLExParserRULE_ignoring               = 68
	PromQLExParserRULE_groupLeft              = 69
	PromQLExParserRULE_groupRight             = 70
	PromQLExParserRULE_labelNameList          = 71
	PromQLExParserRULE_keyword                = 72
	PromQLExParserRULE_string                 = 73
)

// IPromqlxContext is an interface to support dynamic dispatch.
type IPromqlxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	EOF() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

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
	p.RuleIndex = PromQLExParserRULE_promqlx
	return p
}

func InitEmptyPromqlxContext(p *PromqlxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_promqlx
}

func (*PromqlxContext) IsPromqlxContext() {}

func NewPromqlxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromqlxContext {
	var p = new(PromqlxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_promqlx

	return p
}

func (s *PromqlxContext) GetParser() antlr.Parser { return s.parser }

func (s *PromqlxContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *PromqlxContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *PromqlxContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromQLExParserEOF, 0)
}

func (s *PromqlxContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserNL)
}

func (s *PromqlxContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserNL, i)
}

func (s *PromqlxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromqlxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromqlxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterPromqlx(s)
	}
}

func (s *PromqlxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitPromqlx(s)
	}
}

func (p *PromQLExParser) Promqlx() (localctx IPromqlxContext) {
	localctx = NewPromqlxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromQLExParserRULE_promqlx)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Statement()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLExParserNL {
		{
			p.SetState(149)
			p.Match(PromQLExParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Statement()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(PromQLExParserEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Alias_def() IAlias_defContext
	Macro_def() IMacro_defContext
	If_statement() IIf_statementContext
	VectorOperation() IVectorOperationContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Alias_def() IAlias_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlias_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlias_defContext)
}

func (s *StatementContext) Macro_def() IMacro_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacro_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacro_defContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) VectorOperation() IVectorOperationContext {
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

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *PromQLExParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromQLExParserRULE_statement)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Alias_def()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Macro_def()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
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

// IAlias_defContext is an interface to support dynamic dispatch.
type IAlias_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEF() antlr.TerminalNode
	ID() antlr.TerminalNode
	EQ() antlr.TerminalNode
	VectorOperation() IVectorOperationContext

	// IsAlias_defContext differentiates from other interfaces.
	IsAlias_defContext()
}

type Alias_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlias_defContext() *Alias_defContext {
	var p = new(Alias_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_alias_def
	return p
}

func InitEmptyAlias_defContext(p *Alias_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_alias_def
}

func (*Alias_defContext) IsAlias_defContext() {}

func NewAlias_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_defContext {
	var p = new(Alias_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_alias_def

	return p
}

func (s *Alias_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_defContext) DEF() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDEF, 0)
}

func (s *Alias_defContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
}

func (s *Alias_defContext) EQ() antlr.TerminalNode {
	return s.GetToken(PromQLExParserEQ, 0)
}

func (s *Alias_defContext) VectorOperation() IVectorOperationContext {
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

func (s *Alias_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAlias_def(s)
	}
}

func (s *Alias_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAlias_def(s)
	}
}

func (p *PromQLExParser) Alias_def() (localctx IAlias_defContext) {
	localctx = NewAlias_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromQLExParserRULE_alias_def)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(PromQLExParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
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

// IAlias_callContext is an interface to support dynamic dispatch.
type IAlias_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALL_SIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsAlias_callContext differentiates from other interfaces.
	IsAlias_callContext()
}

type Alias_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlias_callContext() *Alias_callContext {
	var p = new(Alias_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_alias_call
	return p
}

func InitEmptyAlias_callContext(p *Alias_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_alias_call
}

func (*Alias_callContext) IsAlias_callContext() {}

func NewAlias_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_callContext {
	var p = new(Alias_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_alias_call

	return p
}

func (s *Alias_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_callContext) CALL_SIGN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserCALL_SIGN, 0)
}

func (s *Alias_callContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
}

func (s *Alias_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAlias_call(s)
	}
}

func (s *Alias_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAlias_call(s)
	}
}

func (p *PromQLExParser) Alias_call() (localctx IAlias_callContext) {
	localctx = NewAlias_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromQLExParserRULE_alias_call)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(PromQLExParserID)
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

// IMacro_defContext is an interface to support dynamic dispatch.
type IMacro_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEF() antlr.TerminalNode
	ID() antlr.TerminalNode
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	Statement_block() IStatement_blockContext
	Args_decl() IArgs_declContext

	// IsMacro_defContext differentiates from other interfaces.
	IsMacro_defContext()
}

type Macro_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacro_defContext() *Macro_defContext {
	var p = new(Macro_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_macro_def
	return p
}

func InitEmptyMacro_defContext(p *Macro_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_macro_def
}

func (*Macro_defContext) IsMacro_defContext() {}

func NewMacro_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_defContext {
	var p = new(Macro_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_macro_def

	return p
}

func (s *Macro_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_defContext) DEF() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDEF, 0)
}

func (s *Macro_defContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
}

func (s *Macro_defContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *Macro_defContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *Macro_defContext) Statement_block() IStatement_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_blockContext)
}

func (s *Macro_defContext) Args_decl() IArgs_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgs_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgs_declContext)
}

func (s *Macro_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Macro_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterMacro_def(s)
	}
}

func (s *Macro_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitMacro_def(s)
	}
}

func (p *PromQLExParser) Macro_def() (localctx IMacro_defContext) {
	localctx = NewMacro_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromQLExParserRULE_macro_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserID {
		{
			p.SetState(175)
			p.Args_decl()
		}

	}
	{
		p.SetState(178)
		p.Match(PromQLExParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Statement_block()
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

// IMacro_callContext is an interface to support dynamic dispatch.
type IMacro_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALL_SIGN() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	LEFT_PAREN() antlr.TerminalNode
	RIGHT_PAREN() antlr.TerminalNode
	Arg_list() IArg_listContext

	// IsMacro_callContext differentiates from other interfaces.
	IsMacro_callContext()
}

type Macro_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacro_callContext() *Macro_callContext {
	var p = new(Macro_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_macro_call
	return p
}

func InitEmptyMacro_callContext(p *Macro_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_macro_call
}

func (*Macro_callContext) IsMacro_callContext() {}

func NewMacro_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_callContext {
	var p = new(Macro_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_macro_call

	return p
}

func (s *Macro_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_callContext) CALL_SIGN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserCALL_SIGN, 0)
}

func (s *Macro_callContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserID)
}

func (s *Macro_callContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, i)
}

func (s *Macro_callContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *Macro_callContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *Macro_callContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *Macro_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Macro_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterMacro_call(s)
	}
}

func (s *Macro_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitMacro_call(s)
	}
}

func (p *PromQLExParser) Macro_call() (localctx IMacro_callContext) {
	localctx = NewMacro_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromQLExParserRULE_macro_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&884112901863965710) != 0 {
		{
			p.SetState(185)
			p.Arg_list()
		}

	}
	{
		p.SetState(188)
		p.Match(PromQLExParserRIGHT_PAREN)
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

// IArgs_declContext is an interface to support dynamic dispatch.
type IArgs_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg_name() []IArg_nameContext
	Arg_name(i int) IArg_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgs_declContext differentiates from other interfaces.
	IsArgs_declContext()
}

type Args_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgs_declContext() *Args_declContext {
	var p = new(Args_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_args_decl
	return p
}

func InitEmptyArgs_declContext(p *Args_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_args_decl
}

func (*Args_declContext) IsArgs_declContext() {}

func NewArgs_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Args_declContext {
	var p = new(Args_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_args_decl

	return p
}

func (s *Args_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Args_declContext) AllArg_name() []IArg_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArg_nameContext); ok {
			len++
		}
	}

	tst := make([]IArg_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArg_nameContext); ok {
			tst[i] = t.(IArg_nameContext)
			i++
		}
	}

	return tst
}

func (s *Args_declContext) Arg_name(i int) IArg_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_nameContext); ok {
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

	return t.(IArg_nameContext)
}

func (s *Args_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *Args_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *Args_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Args_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Args_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterArgs_decl(s)
	}
}

func (s *Args_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitArgs_decl(s)
	}
}

func (p *PromQLExParser) Args_decl() (localctx IArgs_declContext) {
	localctx = NewArgs_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromQLExParserRULE_args_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Arg_name()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLExParserCOMMA {
		{
			p.SetState(191)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Arg_name()
		}

		p.SetState(197)
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

// IArg_nameContext is an interface to support dynamic dispatch.
type IArg_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsArg_nameContext differentiates from other interfaces.
	IsArg_nameContext()
}

type Arg_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_nameContext() *Arg_nameContext {
	var p = new(Arg_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_arg_name
	return p
}

func InitEmptyArg_nameContext(p *Arg_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_arg_name
}

func (*Arg_nameContext) IsArg_nameContext() {}

func NewArg_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_nameContext {
	var p = new(Arg_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_arg_name

	return p
}

func (s *Arg_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_nameContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
}

func (s *Arg_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterArg_name(s)
	}
}

func (s *Arg_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitArg_name(s)
	}
}

func (p *PromQLExParser) Arg_name() (localctx IArg_nameContext) {
	localctx = NewArg_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromQLExParserRULE_arg_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(PromQLExParserID)
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

// IStatement_blockContext is an interface to support dynamic dispatch.
type IStatement_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACE() antlr.TerminalNode
	Promqlx() IPromqlxContext
	RIGHT_BRACE() antlr.TerminalNode

	// IsStatement_blockContext differentiates from other interfaces.
	IsStatement_blockContext()
}

type Statement_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatement_blockContext() *Statement_blockContext {
	var p = new(Statement_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_statement_block
	return p
}

func InitEmptyStatement_blockContext(p *Statement_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_statement_block
}

func (*Statement_blockContext) IsStatement_blockContext() {}

func NewStatement_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Statement_blockContext {
	var p = new(Statement_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_statement_block

	return p
}

func (s *Statement_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Statement_blockContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_BRACE, 0)
}

func (s *Statement_blockContext) Promqlx() IPromqlxContext {
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

func (s *Statement_blockContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_BRACE, 0)
}

func (s *Statement_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Statement_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Statement_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterStatement_block(s)
	}
}

func (s *Statement_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitStatement_block(s)
	}
}

func (p *PromQLExParser) Statement_block() (localctx IStatement_blockContext) {
	localctx = NewStatement_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromQLExParserRULE_statement_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(PromQLExParserLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Promqlx()
	}
	{
		p.SetState(202)
		p.Match(PromQLExParserRIGHT_BRACE)
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

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVectorOperation() []IVectorOperationContext
	VectorOperation(i int) IVectorOperationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *Arg_listContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *Arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *Arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arg_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterArg_list(s)
	}
}

func (s *Arg_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitArg_list(s)
	}
}

func (p *PromQLExParser) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromQLExParserRULE_arg_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.vectorOperation(0)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromQLExParserCOMMA {
		{
			p.SetState(205)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.vectorOperation(0)
		}

		p.SetState(209)
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Condition() IConditionContext
	Statement_block() IStatement_blockContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(PromQLExParserIF, 0)
}

func (s *If_statementContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *If_statementContext) Statement_block() IStatement_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatement_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatement_blockContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *PromQLExParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromQLExParserRULE_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(PromQLExParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Condition()
	}
	{
		p.SetState(213)
		p.Statement_block()
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

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CompareVectorOperation() ICompareVectorOperationContext
	TrueConst() ITrueConstContext
	FalseConst() IFalseConstContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CompareVectorOperation() ICompareVectorOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompareVectorOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompareVectorOperationContext)
}

func (s *ConditionContext) TrueConst() ITrueConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrueConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrueConstContext)
}

func (s *ConditionContext) FalseConst() IFalseConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFalseConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFalseConstContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *PromQLExParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromQLExParserRULE_condition)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserMETRIC_NAME, PromQLExParserPOSITIVE_INTEGER, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserSTRING, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_BRACE, PromQLExParserLEFT_PAREN, PromQLExParserDURATION, PromQLExParserID, PromQLExParserRAW_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)
			p.CompareVectorOperation()
		}

	case PromQLExParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.TrueConst()
		}

	case PromQLExParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(217)
			p.FalseConst()
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

// ICompareVectorOperationContext is an interface to support dynamic dispatch.
type ICompareVectorOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVectorOperation() []IVectorOperationContext
	VectorOperation(i int) IVectorOperationContext
	CompareOp() ICompareOpContext

	// IsCompareVectorOperationContext differentiates from other interfaces.
	IsCompareVectorOperationContext()
}

type CompareVectorOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareVectorOperationContext() *CompareVectorOperationContext {
	var p = new(CompareVectorOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_compareVectorOperation
	return p
}

func InitEmptyCompareVectorOperationContext(p *CompareVectorOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_compareVectorOperation
}

func (*CompareVectorOperationContext) IsCompareVectorOperationContext() {}

func NewCompareVectorOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareVectorOperationContext {
	var p = new(CompareVectorOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_compareVectorOperation

	return p
}

func (s *CompareVectorOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareVectorOperationContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *CompareVectorOperationContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *CompareVectorOperationContext) CompareOp() ICompareOpContext {
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

func (s *CompareVectorOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareVectorOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareVectorOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterCompareVectorOperation(s)
	}
}

func (s *CompareVectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitCompareVectorOperation(s)
	}
}

func (p *PromQLExParser) CompareVectorOperation() (localctx ICompareVectorOperationContext) {
	localctx = NewCompareVectorOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromQLExParserRULE_compareVectorOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.vectorOperation(0)
	}
	{
		p.SetState(221)
		p.CompareOp()
	}
	{
		p.SetState(222)
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

// ITrueConstContext is an interface to support dynamic dispatch.
type ITrueConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode

	// IsTrueConstContext differentiates from other interfaces.
	IsTrueConstContext()
}

type TrueConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrueConstContext() *TrueConstContext {
	var p = new(TrueConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_trueConst
	return p
}

func InitEmptyTrueConstContext(p *TrueConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_trueConst
}

func (*TrueConstContext) IsTrueConstContext() {}

func NewTrueConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueConstContext {
	var p = new(TrueConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_trueConst

	return p
}

func (s *TrueConstContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueConstContext) TRUE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTRUE, 0)
}

func (s *TrueConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterTrueConst(s)
	}
}

func (s *TrueConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitTrueConst(s)
	}
}

func (p *PromQLExParser) TrueConst() (localctx ITrueConstContext) {
	localctx = NewTrueConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromQLExParserRULE_trueConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(PromQLExParserTRUE)
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

// IFalseConstContext is an interface to support dynamic dispatch.
type IFalseConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FALSE() antlr.TerminalNode

	// IsFalseConstContext differentiates from other interfaces.
	IsFalseConstContext()
}

type FalseConstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFalseConstContext() *FalseConstContext {
	var p = new(FalseConstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_falseConst
	return p
}

func InitEmptyFalseConstContext(p *FalseConstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_falseConst
}

func (*FalseConstContext) IsFalseConstContext() {}

func NewFalseConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FalseConstContext {
	var p = new(FalseConstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_falseConst

	return p
}

func (s *FalseConstContext) GetParser() antlr.Parser { return s.parser }

func (s *FalseConstContext) FALSE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserFALSE, 0)
}

func (s *FalseConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FalseConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterFalseConst(s)
	}
}

func (s *FalseConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitFalseConst(s)
	}
}

func (p *PromQLExParser) FalseConst() (localctx IFalseConstContext) {
	localctx = NewFalseConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromQLExParserRULE_falseConst)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(PromQLExParserFALSE)
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

// ITime_instant_literalContext is an interface to support dynamic dispatch.
type ITime_instant_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_date_time() IIso_date_timeContext
	Unix_timestamp() IUnix_timestampContext

	// IsTime_instant_literalContext differentiates from other interfaces.
	IsTime_instant_literalContext()
}

type Time_instant_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTime_instant_literalContext() *Time_instant_literalContext {
	var p = new(Time_instant_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_time_instant_literal
	return p
}

func InitEmptyTime_instant_literalContext(p *Time_instant_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_time_instant_literal
}

func (*Time_instant_literalContext) IsTime_instant_literalContext() {}

func NewTime_instant_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Time_instant_literalContext {
	var p = new(Time_instant_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_time_instant_literal

	return p
}

func (s *Time_instant_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Time_instant_literalContext) Iso_date_time() IIso_date_timeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_timeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_timeContext)
}

func (s *Time_instant_literalContext) Unix_timestamp() IUnix_timestampContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnix_timestampContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnix_timestampContext)
}

func (s *Time_instant_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Time_instant_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Time_instant_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterTime_instant_literal(s)
	}
}

func (s *Time_instant_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitTime_instant_literal(s)
	}
}

func (p *PromQLExParser) Time_instant_literal() (localctx ITime_instant_literalContext) {
	localctx = NewTime_instant_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromQLExParserRULE_time_instant_literal)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Iso_date_time()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Unix_timestamp()
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

// IIso_date_timeContext is an interface to support dynamic dispatch.
type IIso_date_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_date_time_ymdhmsf() IIso_date_time_ymdhmsfContext
	Iso_date_time_ymdhms() IIso_date_time_ymdhmsContext
	Iso_date_time_ymdhm() IIso_date_time_ymdhmContext
	Iso_date_time_ymdh() IIso_date_time_ymdhContext
	Iso_date_time_ymd() IIso_date_time_ymdContext
	Iso_date_time_ym() IIso_date_time_ymContext
	Iso_date_time_y() IIso_date_time_yContext

	// IsIso_date_timeContext differentiates from other interfaces.
	IsIso_date_timeContext()
}

type Iso_date_timeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_timeContext() *Iso_date_timeContext {
	var p = new(Iso_date_timeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time
	return p
}

func InitEmptyIso_date_timeContext(p *Iso_date_timeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time
}

func (*Iso_date_timeContext) IsIso_date_timeContext() {}

func NewIso_date_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_timeContext {
	var p = new(Iso_date_timeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time

	return p
}

func (s *Iso_date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_timeContext) Iso_date_time_ymdhmsf() IIso_date_time_ymdhmsfContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymdhmsfContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymdhmsfContext)
}

func (s *Iso_date_timeContext) Iso_date_time_ymdhms() IIso_date_time_ymdhmsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymdhmsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymdhmsContext)
}

func (s *Iso_date_timeContext) Iso_date_time_ymdhm() IIso_date_time_ymdhmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymdhmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymdhmContext)
}

func (s *Iso_date_timeContext) Iso_date_time_ymdh() IIso_date_time_ymdhContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymdhContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymdhContext)
}

func (s *Iso_date_timeContext) Iso_date_time_ymd() IIso_date_time_ymdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymdContext)
}

func (s *Iso_date_timeContext) Iso_date_time_ym() IIso_date_time_ymContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_ymContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_ymContext)
}

func (s *Iso_date_timeContext) Iso_date_time_y() IIso_date_time_yContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_date_time_yContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_date_time_yContext)
}

func (s *Iso_date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time(s)
	}
}

func (s *Iso_date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time(s)
	}
}

func (p *PromQLExParser) Iso_date_time() (localctx IIso_date_timeContext) {
	localctx = NewIso_date_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromQLExParserRULE_iso_date_time)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Iso_date_time_ymdhmsf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Iso_date_time_ymdhms()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			p.Iso_date_time_ymdhm()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(235)
			p.Iso_date_time_ymdh()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(236)
			p.Iso_date_time_ymd()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(237)
			p.Iso_date_time_ym()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(238)
			p.Iso_date_time_y()
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

// IIso_date_time_ymdhmsfContext is an interface to support dynamic dispatch.
type IIso_date_time_ymdhmsfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Iso_month() IIso_monthContext
	Iso_day() IIso_dayContext
	T() antlr.TerminalNode
	Iso_hour() IIso_hourContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	Iso_minutes() IIso_minutesContext
	Iso_seconds() IIso_secondsContext
	DOT() antlr.TerminalNode
	Is_frac_sec() IIs_frac_secContext

	// IsIso_date_time_ymdhmsfContext differentiates from other interfaces.
	IsIso_date_time_ymdhmsfContext()
}

type Iso_date_time_ymdhmsfContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymdhmsfContext() *Iso_date_time_ymdhmsfContext {
	var p = new(Iso_date_time_ymdhmsfContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhmsf
	return p
}

func InitEmptyIso_date_time_ymdhmsfContext(p *Iso_date_time_ymdhmsfContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhmsf
}

func (*Iso_date_time_ymdhmsfContext) IsIso_date_time_ymdhmsfContext() {}

func NewIso_date_time_ymdhmsfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymdhmsfContext {
	var p = new(Iso_date_time_ymdhmsfContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhmsf

	return p
}

func (s *Iso_date_time_ymdhmsfContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymdhmsfContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymdhmsfContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSUB)
}

func (s *Iso_date_time_ymdhmsfContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, i)
}

func (s *Iso_date_time_ymdhmsfContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymdhmsfContext) Iso_day() IIso_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_dayContext)
}

func (s *Iso_date_time_ymdhmsfContext) T() antlr.TerminalNode {
	return s.GetToken(PromQLExParserT, 0)
}

func (s *Iso_date_time_ymdhmsfContext) Iso_hour() IIso_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_hourContext)
}

func (s *Iso_date_time_ymdhmsfContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserCOLON)
}

func (s *Iso_date_time_ymdhmsfContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOLON, i)
}

func (s *Iso_date_time_ymdhmsfContext) Iso_minutes() IIso_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_minutesContext)
}

func (s *Iso_date_time_ymdhmsfContext) Iso_seconds() IIso_secondsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_secondsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_secondsContext)
}

func (s *Iso_date_time_ymdhmsfContext) DOT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDOT, 0)
}

func (s *Iso_date_time_ymdhmsfContext) Is_frac_sec() IIs_frac_secContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIs_frac_secContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIs_frac_secContext)
}

func (s *Iso_date_time_ymdhmsfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymdhmsfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymdhmsfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ymdhmsf(s)
	}
}

func (s *Iso_date_time_ymdhmsfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ymdhmsf(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ymdhmsf() (localctx IIso_date_time_ymdhmsfContext) {
	localctx = NewIso_date_time_ymdhmsfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromQLExParserRULE_iso_date_time_ymdhmsf)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Iso_year()
	}
	{
		p.SetState(242)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Iso_month()
	}
	{
		p.SetState(244)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Iso_day()
	}
	{
		p.SetState(246)
		p.Match(PromQLExParserT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Iso_hour()
	}
	{
		p.SetState(248)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Iso_minutes()
	}
	{
		p.SetState(250)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Iso_seconds()
	}
	{
		p.SetState(252)
		p.Match(PromQLExParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Is_frac_sec()
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

// IIso_date_time_ymdhmsContext is an interface to support dynamic dispatch.
type IIso_date_time_ymdhmsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Iso_month() IIso_monthContext
	Iso_day() IIso_dayContext
	T() antlr.TerminalNode
	Iso_hour() IIso_hourContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	Iso_minutes() IIso_minutesContext
	Iso_seconds() IIso_secondsContext

	// IsIso_date_time_ymdhmsContext differentiates from other interfaces.
	IsIso_date_time_ymdhmsContext()
}

type Iso_date_time_ymdhmsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymdhmsContext() *Iso_date_time_ymdhmsContext {
	var p = new(Iso_date_time_ymdhmsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhms
	return p
}

func InitEmptyIso_date_time_ymdhmsContext(p *Iso_date_time_ymdhmsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhms
}

func (*Iso_date_time_ymdhmsContext) IsIso_date_time_ymdhmsContext() {}

func NewIso_date_time_ymdhmsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymdhmsContext {
	var p = new(Iso_date_time_ymdhmsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhms

	return p
}

func (s *Iso_date_time_ymdhmsContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymdhmsContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymdhmsContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSUB)
}

func (s *Iso_date_time_ymdhmsContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, i)
}

func (s *Iso_date_time_ymdhmsContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymdhmsContext) Iso_day() IIso_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_dayContext)
}

func (s *Iso_date_time_ymdhmsContext) T() antlr.TerminalNode {
	return s.GetToken(PromQLExParserT, 0)
}

func (s *Iso_date_time_ymdhmsContext) Iso_hour() IIso_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_hourContext)
}

func (s *Iso_date_time_ymdhmsContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserCOLON)
}

func (s *Iso_date_time_ymdhmsContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOLON, i)
}

func (s *Iso_date_time_ymdhmsContext) Iso_minutes() IIso_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_minutesContext)
}

func (s *Iso_date_time_ymdhmsContext) Iso_seconds() IIso_secondsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_secondsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_secondsContext)
}

func (s *Iso_date_time_ymdhmsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymdhmsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymdhmsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ymdhms(s)
	}
}

func (s *Iso_date_time_ymdhmsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ymdhms(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ymdhms() (localctx IIso_date_time_ymdhmsContext) {
	localctx = NewIso_date_time_ymdhmsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromQLExParserRULE_iso_date_time_ymdhms)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Iso_year()
	}
	{
		p.SetState(256)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Iso_month()
	}
	{
		p.SetState(258)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Iso_day()
	}
	{
		p.SetState(260)
		p.Match(PromQLExParserT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Iso_hour()
	}
	{
		p.SetState(262)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Iso_minutes()
	}
	{
		p.SetState(264)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Iso_seconds()
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

// IIso_date_time_ymdhmContext is an interface to support dynamic dispatch.
type IIso_date_time_ymdhmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Iso_month() IIso_monthContext
	Iso_day() IIso_dayContext
	T() antlr.TerminalNode
	Iso_hour() IIso_hourContext
	COLON() antlr.TerminalNode
	Iso_minutes() IIso_minutesContext

	// IsIso_date_time_ymdhmContext differentiates from other interfaces.
	IsIso_date_time_ymdhmContext()
}

type Iso_date_time_ymdhmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymdhmContext() *Iso_date_time_ymdhmContext {
	var p = new(Iso_date_time_ymdhmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhm
	return p
}

func InitEmptyIso_date_time_ymdhmContext(p *Iso_date_time_ymdhmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhm
}

func (*Iso_date_time_ymdhmContext) IsIso_date_time_ymdhmContext() {}

func NewIso_date_time_ymdhmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymdhmContext {
	var p = new(Iso_date_time_ymdhmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdhm

	return p
}

func (s *Iso_date_time_ymdhmContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymdhmContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymdhmContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSUB)
}

func (s *Iso_date_time_ymdhmContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, i)
}

func (s *Iso_date_time_ymdhmContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymdhmContext) Iso_day() IIso_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_dayContext)
}

func (s *Iso_date_time_ymdhmContext) T() antlr.TerminalNode {
	return s.GetToken(PromQLExParserT, 0)
}

func (s *Iso_date_time_ymdhmContext) Iso_hour() IIso_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_hourContext)
}

func (s *Iso_date_time_ymdhmContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOLON, 0)
}

func (s *Iso_date_time_ymdhmContext) Iso_minutes() IIso_minutesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_minutesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_minutesContext)
}

func (s *Iso_date_time_ymdhmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymdhmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymdhmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ymdhm(s)
	}
}

func (s *Iso_date_time_ymdhmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ymdhm(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ymdhm() (localctx IIso_date_time_ymdhmContext) {
	localctx = NewIso_date_time_ymdhmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromQLExParserRULE_iso_date_time_ymdhm)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Iso_year()
	}
	{
		p.SetState(268)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Iso_month()
	}
	{
		p.SetState(270)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Iso_day()
	}
	{
		p.SetState(272)
		p.Match(PromQLExParserT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Iso_hour()
	}
	{
		p.SetState(274)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(275)
		p.Iso_minutes()
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

// IIso_date_time_ymdhContext is an interface to support dynamic dispatch.
type IIso_date_time_ymdhContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Iso_month() IIso_monthContext
	Iso_day() IIso_dayContext
	T() antlr.TerminalNode
	Iso_hour() IIso_hourContext

	// IsIso_date_time_ymdhContext differentiates from other interfaces.
	IsIso_date_time_ymdhContext()
}

type Iso_date_time_ymdhContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymdhContext() *Iso_date_time_ymdhContext {
	var p = new(Iso_date_time_ymdhContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdh
	return p
}

func InitEmptyIso_date_time_ymdhContext(p *Iso_date_time_ymdhContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdh
}

func (*Iso_date_time_ymdhContext) IsIso_date_time_ymdhContext() {}

func NewIso_date_time_ymdhContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymdhContext {
	var p = new(Iso_date_time_ymdhContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymdh

	return p
}

func (s *Iso_date_time_ymdhContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymdhContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymdhContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSUB)
}

func (s *Iso_date_time_ymdhContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, i)
}

func (s *Iso_date_time_ymdhContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymdhContext) Iso_day() IIso_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_dayContext)
}

func (s *Iso_date_time_ymdhContext) T() antlr.TerminalNode {
	return s.GetToken(PromQLExParserT, 0)
}

func (s *Iso_date_time_ymdhContext) Iso_hour() IIso_hourContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_hourContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_hourContext)
}

func (s *Iso_date_time_ymdhContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymdhContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymdhContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ymdh(s)
	}
}

func (s *Iso_date_time_ymdhContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ymdh(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ymdh() (localctx IIso_date_time_ymdhContext) {
	localctx = NewIso_date_time_ymdhContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromQLExParserRULE_iso_date_time_ymdh)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Iso_year()
	}
	{
		p.SetState(278)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Iso_month()
	}
	{
		p.SetState(280)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Iso_day()
	}
	{
		p.SetState(282)
		p.Match(PromQLExParserT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Iso_hour()
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

// IIso_date_time_ymdContext is an interface to support dynamic dispatch.
type IIso_date_time_ymdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	AllSUB() []antlr.TerminalNode
	SUB(i int) antlr.TerminalNode
	Iso_month() IIso_monthContext
	Iso_day() IIso_dayContext

	// IsIso_date_time_ymdContext differentiates from other interfaces.
	IsIso_date_time_ymdContext()
}

type Iso_date_time_ymdContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymdContext() *Iso_date_time_ymdContext {
	var p = new(Iso_date_time_ymdContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymd
	return p
}

func InitEmptyIso_date_time_ymdContext(p *Iso_date_time_ymdContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymd
}

func (*Iso_date_time_ymdContext) IsIso_date_time_ymdContext() {}

func NewIso_date_time_ymdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymdContext {
	var p = new(Iso_date_time_ymdContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ymd

	return p
}

func (s *Iso_date_time_ymdContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymdContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymdContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSUB)
}

func (s *Iso_date_time_ymdContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, i)
}

func (s *Iso_date_time_ymdContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymdContext) Iso_day() IIso_dayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_dayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_dayContext)
}

func (s *Iso_date_time_ymdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ymd(s)
	}
}

func (s *Iso_date_time_ymdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ymd(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ymd() (localctx IIso_date_time_ymdContext) {
	localctx = NewIso_date_time_ymdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PromQLExParserRULE_iso_date_time_ymd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Iso_year()
	}
	{
		p.SetState(286)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Iso_month()
	}
	{
		p.SetState(288)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Iso_day()
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

// IIso_date_time_ymContext is an interface to support dynamic dispatch.
type IIso_date_time_ymContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext
	SUB() antlr.TerminalNode
	Iso_month() IIso_monthContext

	// IsIso_date_time_ymContext differentiates from other interfaces.
	IsIso_date_time_ymContext()
}

type Iso_date_time_ymContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_ymContext() *Iso_date_time_ymContext {
	var p = new(Iso_date_time_ymContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ym
	return p
}

func InitEmptyIso_date_time_ymContext(p *Iso_date_time_ymContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ym
}

func (*Iso_date_time_ymContext) IsIso_date_time_ymContext() {}

func NewIso_date_time_ymContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_ymContext {
	var p = new(Iso_date_time_ymContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_ym

	return p
}

func (s *Iso_date_time_ymContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_ymContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_ymContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, 0)
}

func (s *Iso_date_time_ymContext) Iso_month() IIso_monthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_monthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_monthContext)
}

func (s *Iso_date_time_ymContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_ymContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_ymContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_ym(s)
	}
}

func (s *Iso_date_time_ymContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_ym(s)
	}
}

func (p *PromQLExParser) Iso_date_time_ym() (localctx IIso_date_time_ymContext) {
	localctx = NewIso_date_time_ymContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PromQLExParserRULE_iso_date_time_ym)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Iso_year()
	}
	{
		p.SetState(292)
		p.Match(PromQLExParserSUB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Iso_month()
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

// IIso_date_time_yContext is an interface to support dynamic dispatch.
type IIso_date_time_yContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Iso_year() IIso_yearContext

	// IsIso_date_time_yContext differentiates from other interfaces.
	IsIso_date_time_yContext()
}

type Iso_date_time_yContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_date_time_yContext() *Iso_date_time_yContext {
	var p = new(Iso_date_time_yContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_y
	return p
}

func InitEmptyIso_date_time_yContext(p *Iso_date_time_yContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_date_time_y
}

func (*Iso_date_time_yContext) IsIso_date_time_yContext() {}

func NewIso_date_time_yContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_date_time_yContext {
	var p = new(Iso_date_time_yContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_date_time_y

	return p
}

func (s *Iso_date_time_yContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_date_time_yContext) Iso_year() IIso_yearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIso_yearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIso_yearContext)
}

func (s *Iso_date_time_yContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_date_time_yContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_date_time_yContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_date_time_y(s)
	}
}

func (s *Iso_date_time_yContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_date_time_y(s)
	}
}

func (p *PromQLExParser) Iso_date_time_y() (localctx IIso_date_time_yContext) {
	localctx = NewIso_date_time_yContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PromQLExParserRULE_iso_date_time_y)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Iso_year()
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

// IIso_yearContext is an interface to support dynamic dispatch.
type IIso_yearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POSITIVE_INTEGER() antlr.TerminalNode

	// IsIso_yearContext differentiates from other interfaces.
	IsIso_yearContext()
}

type Iso_yearContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_yearContext() *Iso_yearContext {
	var p = new(Iso_yearContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_year
	return p
}

func InitEmptyIso_yearContext(p *Iso_yearContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_year
}

func (*Iso_yearContext) IsIso_yearContext() {}

func NewIso_yearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_yearContext {
	var p = new(Iso_yearContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_year

	return p
}

func (s *Iso_yearContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_yearContext) POSITIVE_INTEGER() antlr.TerminalNode {
	return s.GetToken(PromQLExParserPOSITIVE_INTEGER, 0)
}

func (s *Iso_yearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_yearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_yearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_year(s)
	}
}

func (s *Iso_yearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_year(s)
	}
}

func (p *PromQLExParser) Iso_year() (localctx IIso_yearContext) {
	localctx = NewIso_yearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PromQLExParserRULE_iso_year)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(PromQLExParserPOSITIVE_INTEGER)
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

// IIso_monthContext is an interface to support dynamic dispatch.
type IIso_monthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TWO_DIGITS() antlr.TerminalNode

	// IsIso_monthContext differentiates from other interfaces.
	IsIso_monthContext()
}

type Iso_monthContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_monthContext() *Iso_monthContext {
	var p = new(Iso_monthContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_month
	return p
}

func InitEmptyIso_monthContext(p *Iso_monthContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_month
}

func (*Iso_monthContext) IsIso_monthContext() {}

func NewIso_monthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_monthContext {
	var p = new(Iso_monthContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_month

	return p
}

func (s *Iso_monthContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_monthContext) TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTWO_DIGITS, 0)
}

func (s *Iso_monthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_monthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_monthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_month(s)
	}
}

func (s *Iso_monthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_month(s)
	}
}

func (p *PromQLExParser) Iso_month() (localctx IIso_monthContext) {
	localctx = NewIso_monthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PromQLExParserRULE_iso_month)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(PromQLExParserTWO_DIGITS)
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

// IIso_dayContext is an interface to support dynamic dispatch.
type IIso_dayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TWO_DIGITS() antlr.TerminalNode

	// IsIso_dayContext differentiates from other interfaces.
	IsIso_dayContext()
}

type Iso_dayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_dayContext() *Iso_dayContext {
	var p = new(Iso_dayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_day
	return p
}

func InitEmptyIso_dayContext(p *Iso_dayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_day
}

func (*Iso_dayContext) IsIso_dayContext() {}

func NewIso_dayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_dayContext {
	var p = new(Iso_dayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_day

	return p
}

func (s *Iso_dayContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_dayContext) TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTWO_DIGITS, 0)
}

func (s *Iso_dayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_dayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_dayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_day(s)
	}
}

func (s *Iso_dayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_day(s)
	}
}

func (p *PromQLExParser) Iso_day() (localctx IIso_dayContext) {
	localctx = NewIso_dayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromQLExParserRULE_iso_day)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(PromQLExParserTWO_DIGITS)
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

// IIso_hourContext is an interface to support dynamic dispatch.
type IIso_hourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TWO_DIGITS() antlr.TerminalNode

	// IsIso_hourContext differentiates from other interfaces.
	IsIso_hourContext()
}

type Iso_hourContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_hourContext() *Iso_hourContext {
	var p = new(Iso_hourContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_hour
	return p
}

func InitEmptyIso_hourContext(p *Iso_hourContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_hour
}

func (*Iso_hourContext) IsIso_hourContext() {}

func NewIso_hourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_hourContext {
	var p = new(Iso_hourContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_hour

	return p
}

func (s *Iso_hourContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_hourContext) TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTWO_DIGITS, 0)
}

func (s *Iso_hourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_hourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_hourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_hour(s)
	}
}

func (s *Iso_hourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_hour(s)
	}
}

func (p *PromQLExParser) Iso_hour() (localctx IIso_hourContext) {
	localctx = NewIso_hourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromQLExParserRULE_iso_hour)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(PromQLExParserTWO_DIGITS)
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

// IIso_minutesContext is an interface to support dynamic dispatch.
type IIso_minutesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TWO_DIGITS() antlr.TerminalNode

	// IsIso_minutesContext differentiates from other interfaces.
	IsIso_minutesContext()
}

type Iso_minutesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_minutesContext() *Iso_minutesContext {
	var p = new(Iso_minutesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_minutes
	return p
}

func InitEmptyIso_minutesContext(p *Iso_minutesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_minutes
}

func (*Iso_minutesContext) IsIso_minutesContext() {}

func NewIso_minutesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_minutesContext {
	var p = new(Iso_minutesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_minutes

	return p
}

func (s *Iso_minutesContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_minutesContext) TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTWO_DIGITS, 0)
}

func (s *Iso_minutesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_minutesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_minutesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_minutes(s)
	}
}

func (s *Iso_minutesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_minutes(s)
	}
}

func (p *PromQLExParser) Iso_minutes() (localctx IIso_minutesContext) {
	localctx = NewIso_minutesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromQLExParserRULE_iso_minutes)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(PromQLExParserTWO_DIGITS)
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

// IIso_secondsContext is an interface to support dynamic dispatch.
type IIso_secondsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TWO_DIGITS() antlr.TerminalNode

	// IsIso_secondsContext differentiates from other interfaces.
	IsIso_secondsContext()
}

type Iso_secondsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIso_secondsContext() *Iso_secondsContext {
	var p = new(Iso_secondsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_seconds
	return p
}

func InitEmptyIso_secondsContext(p *Iso_secondsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_iso_seconds
}

func (*Iso_secondsContext) IsIso_secondsContext() {}

func NewIso_secondsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iso_secondsContext {
	var p = new(Iso_secondsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_iso_seconds

	return p
}

func (s *Iso_secondsContext) GetParser() antlr.Parser { return s.parser }

func (s *Iso_secondsContext) TWO_DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserTWO_DIGITS, 0)
}

func (s *Iso_secondsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iso_secondsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iso_secondsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIso_seconds(s)
	}
}

func (s *Iso_secondsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIso_seconds(s)
	}
}

func (p *PromQLExParser) Iso_seconds() (localctx IIso_secondsContext) {
	localctx = NewIso_secondsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PromQLExParserRULE_iso_seconds)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(PromQLExParserTWO_DIGITS)
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

// IIs_frac_secContext is an interface to support dynamic dispatch.
type IIs_frac_secContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGITS() antlr.TerminalNode

	// IsIs_frac_secContext differentiates from other interfaces.
	IsIs_frac_secContext()
}

type Is_frac_secContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIs_frac_secContext() *Is_frac_secContext {
	var p = new(Is_frac_secContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_is_frac_sec
	return p
}

func InitEmptyIs_frac_secContext(p *Is_frac_secContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_is_frac_sec
}

func (*Is_frac_secContext) IsIs_frac_secContext() {}

func NewIs_frac_secContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Is_frac_secContext {
	var p = new(Is_frac_secContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_is_frac_sec

	return p
}

func (s *Is_frac_secContext) GetParser() antlr.Parser { return s.parser }

func (s *Is_frac_secContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDIGITS, 0)
}

func (s *Is_frac_secContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Is_frac_secContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Is_frac_secContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIs_frac_sec(s)
	}
}

func (s *Is_frac_secContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIs_frac_sec(s)
	}
}

func (p *PromQLExParser) Is_frac_sec() (localctx IIs_frac_secContext) {
	localctx = NewIs_frac_secContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromQLExParserRULE_is_frac_sec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(PromQLExParserDIGITS)
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

// IUnix_timestampContext is an interface to support dynamic dispatch.
type IUnix_timestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	POSITIVE_INTEGER() antlr.TerminalNode

	// IsUnix_timestampContext differentiates from other interfaces.
	IsUnix_timestampContext()
}

type Unix_timestampContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnix_timestampContext() *Unix_timestampContext {
	var p = new(Unix_timestampContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_unix_timestamp
	return p
}

func InitEmptyUnix_timestampContext(p *Unix_timestampContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_unix_timestamp
}

func (*Unix_timestampContext) IsUnix_timestampContext() {}

func NewUnix_timestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unix_timestampContext {
	var p = new(Unix_timestampContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_unix_timestamp

	return p
}

func (s *Unix_timestampContext) GetParser() antlr.Parser { return s.parser }

func (s *Unix_timestampContext) POSITIVE_INTEGER() antlr.TerminalNode {
	return s.GetToken(PromQLExParserPOSITIVE_INTEGER, 0)
}

func (s *Unix_timestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unix_timestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unix_timestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterUnix_timestamp(s)
	}
}

func (s *Unix_timestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitUnix_timestamp(s)
	}
}

func (p *PromQLExParser) Unix_timestamp() (localctx IUnix_timestampContext) {
	localctx = NewUnix_timestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromQLExParserRULE_unix_timestamp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(PromQLExParserPOSITIVE_INTEGER)
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

// IConst_num_expressionContext is an interface to support dynamic dispatch.
type IConst_num_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNum_literal() []INum_literalContext
	Num_literal(i int) INum_literalContext
	PowOp() IPowOpContext
	UnaryOp() IUnaryOpContext
	MultOp() IMultOpContext
	AddOp() IAddOpContext
	LEFT_PAREN() antlr.TerminalNode
	Const_num_expression() IConst_num_expressionContext
	RIGHT_PAREN() antlr.TerminalNode

	// IsConst_num_expressionContext differentiates from other interfaces.
	IsConst_num_expressionContext()
}

type Const_num_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_num_expressionContext() *Const_num_expressionContext {
	var p = new(Const_num_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_const_num_expression
	return p
}

func InitEmptyConst_num_expressionContext(p *Const_num_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_const_num_expression
}

func (*Const_num_expressionContext) IsConst_num_expressionContext() {}

func NewConst_num_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_num_expressionContext {
	var p = new(Const_num_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_const_num_expression

	return p
}

func (s *Const_num_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_num_expressionContext) AllNum_literal() []INum_literalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INum_literalContext); ok {
			len++
		}
	}

	tst := make([]INum_literalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INum_literalContext); ok {
			tst[i] = t.(INum_literalContext)
			i++
		}
	}

	return tst
}

func (s *Const_num_expressionContext) Num_literal(i int) INum_literalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INum_literalContext); ok {
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

	return t.(INum_literalContext)
}

func (s *Const_num_expressionContext) PowOp() IPowOpContext {
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

func (s *Const_num_expressionContext) UnaryOp() IUnaryOpContext {
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

func (s *Const_num_expressionContext) MultOp() IMultOpContext {
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

func (s *Const_num_expressionContext) AddOp() IAddOpContext {
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

func (s *Const_num_expressionContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *Const_num_expressionContext) Const_num_expression() IConst_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_num_expressionContext)
}

func (s *Const_num_expressionContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *Const_num_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_num_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_num_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConst_num_expression(s)
	}
}

func (s *Const_num_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConst_num_expression(s)
	}
}

func (p *PromQLExParser) Const_num_expression() (localctx IConst_num_expressionContext) {
	localctx = NewConst_num_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromQLExParserRULE_const_num_expression)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Num_literal()
		}
		{
			p.SetState(314)
			p.PowOp()
		}
		{
			p.SetState(315)
			p.Num_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.UnaryOp()
		}
		{
			p.SetState(318)
			p.Num_literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(320)
			p.Num_literal()
		}
		{
			p.SetState(321)
			p.MultOp()
		}
		{
			p.SetState(322)
			p.Num_literal()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(324)
			p.Num_literal()
		}
		{
			p.SetState(325)
			p.AddOp()
		}
		{
			p.SetState(326)
			p.Num_literal()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(328)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Const_num_expression()
		}
		{
			p.SetState(330)
			p.Match(PromQLExParserRIGHT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(332)
			p.Num_literal()
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

// INum_literalContext is an interface to support dynamic dispatch.
type INum_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	Time_instant_literal() ITime_instant_literalContext
	Alias_call() IAlias_callContext

	// IsNum_literalContext differentiates from other interfaces.
	IsNum_literalContext()
}

type Num_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNum_literalContext() *Num_literalContext {
	var p = new(Num_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_num_literal
	return p
}

func InitEmptyNum_literalContext(p *Num_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_num_literal
}

func (*Num_literalContext) IsNum_literalContext() {}

func NewNum_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Num_literalContext {
	var p = new(Num_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_num_literal

	return p
}

func (s *Num_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Num_literalContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNUMBER, 0)
}

func (s *Num_literalContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDURATION, 0)
}

func (s *Num_literalContext) Time_instant_literal() ITime_instant_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITime_instant_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITime_instant_literalContext)
}

func (s *Num_literalContext) Alias_call() IAlias_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlias_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlias_callContext)
}

func (s *Num_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Num_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Num_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterNum_literal(s)
	}
}

func (s *Num_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitNum_literal(s)
	}
}

func (p *PromQLExParser) Num_literal() (localctx INum_literalContext) {
	localctx = NewNum_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromQLExParserRULE_num_literal)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Match(PromQLExParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExParserDURATION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Match(PromQLExParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExParserPOSITIVE_INTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(337)
			p.Time_instant_literal()
		}

	case PromQLExParserCALL_SIGN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(338)
			p.Alias_call()
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

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_num_expression() IConst_num_expressionContext

	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_duration
	return p
}

func InitEmptyDurationContext(p *DurationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_duration
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) Const_num_expression() IConst_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_num_expressionContext)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitDuration(s)
	}
}

func (p *PromQLExParser) Duration() (localctx IDurationContext) {
	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PromQLExParserRULE_duration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Const_num_expression()
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

// ITime_rangeContext is an interface to support dynamic dispatch.
type ITime_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	Duration() IDurationContext
	RIGHT_BRACKET() antlr.TerminalNode

	// IsTime_rangeContext differentiates from other interfaces.
	IsTime_rangeContext()
}

type Time_rangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTime_rangeContext() *Time_rangeContext {
	var p = new(Time_rangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_time_range
	return p
}

func InitEmptyTime_rangeContext(p *Time_rangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_time_range
}

func (*Time_rangeContext) IsTime_rangeContext() {}

func NewTime_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Time_rangeContext {
	var p = new(Time_rangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_time_range

	return p
}

func (s *Time_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Time_rangeContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_BRACKET, 0)
}

func (s *Time_rangeContext) Duration() IDurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *Time_rangeContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_BRACKET, 0)
}

func (s *Time_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Time_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Time_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterTime_range(s)
	}
}

func (s *Time_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitTime_range(s)
	}
}

func (p *PromQLExParser) Time_range() (localctx ITime_rangeContext) {
	localctx = NewTime_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PromQLExParserRULE_time_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.Duration()
	}
	{
		p.SetState(345)
		p.Match(PromQLExParserRIGHT_BRACKET)
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

// ISubquery_rangeContext is an interface to support dynamic dispatch.
type ISubquery_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_BRACKET() antlr.TerminalNode
	AllDuration() []IDurationContext
	Duration(i int) IDurationContext
	COLON() antlr.TerminalNode
	RIGHT_BRACKET() antlr.TerminalNode

	// IsSubquery_rangeContext differentiates from other interfaces.
	IsSubquery_rangeContext()
}

type Subquery_rangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubquery_rangeContext() *Subquery_rangeContext {
	var p = new(Subquery_rangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_subquery_range
	return p
}

func InitEmptySubquery_rangeContext(p *Subquery_rangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_subquery_range
}

func (*Subquery_rangeContext) IsSubquery_rangeContext() {}

func NewSubquery_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subquery_rangeContext {
	var p = new(Subquery_rangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_subquery_range

	return p
}

func (s *Subquery_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Subquery_rangeContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_BRACKET, 0)
}

func (s *Subquery_rangeContext) AllDuration() []IDurationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDurationContext); ok {
			len++
		}
	}

	tst := make([]IDurationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDurationContext); ok {
			tst[i] = t.(IDurationContext)
			i++
		}
	}

	return tst
}

func (s *Subquery_rangeContext) Duration(i int) IDurationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
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

	return t.(IDurationContext)
}

func (s *Subquery_rangeContext) COLON() antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOLON, 0)
}

func (s *Subquery_rangeContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_BRACKET, 0)
}

func (s *Subquery_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subquery_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subquery_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterSubquery_range(s)
	}
}

func (s *Subquery_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitSubquery_range(s)
	}
}

func (p *PromQLExParser) Subquery_range() (localctx ISubquery_rangeContext) {
	localctx = NewSubquery_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PromQLExParserRULE_subquery_range)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.Duration()
	}
	{
		p.SetState(349)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&19140298430022656) != 0 {
		{
			p.SetState(350)
			p.Duration()
		}

	}
	{
		p.SetState(353)
		p.Match(PromQLExParserRIGHT_BRACKET)
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
	Macro_call() IMacro_callContext
	Alias_call() IAlias_callContext
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
	p.RuleIndex = PromQLExParserRULE_vectorOperation
	return p
}

func InitEmptyVectorOperationContext(p *VectorOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_vectorOperation
}

func (*VectorOperationContext) IsVectorOperationContext() {}

func NewVectorOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorOperationContext {
	var p = new(VectorOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_vectorOperation

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

func (s *VectorOperationContext) Macro_call() IMacro_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacro_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacro_callContext)
}

func (s *VectorOperationContext) Alias_call() IAlias_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlias_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlias_callContext)
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
	return s.GetToken(PromQLExParserAT, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVectorOperation(s)
	}
}

func (s *VectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVectorOperation(s)
	}
}

func (p *PromQLExParser) VectorOperation() (localctx IVectorOperationContext) {
	return p.vectorOperation(0)
}

func (p *PromQLExParser) vectorOperation(_p int) (localctx IVectorOperationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewVectorOperationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVectorOperationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 74
	p.EnterRecursionRule(localctx, 74, PromQLExParserRULE_vectorOperation, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(356)
			p.UnaryOp()
		}
		{
			p.SetState(357)
			p.vectorOperation(11)
		}

	case 2:
		{
			p.SetState(359)
			p.Vector()
		}

	case 3:
		{
			p.SetState(360)
			p.Macro_call()
		}

	case 4:
		{
			p.SetState(361)
			p.Alias_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(364)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(365)
					p.PowOp()
				}
				{
					p.SetState(366)
					p.vectorOperation(13)
				}

			case 2:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(368)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(369)
					p.MultOp()
				}
				{
					p.SetState(370)
					p.vectorOperation(11)
				}

			case 3:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(372)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(373)
					p.AddOp()
				}
				{
					p.SetState(374)
					p.vectorOperation(10)
				}

			case 4:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(376)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(377)
					p.CompareOp()
				}
				{
					p.SetState(378)
					p.vectorOperation(9)
				}

			case 5:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(380)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(381)
					p.AndUnlessOp()
				}
				{
					p.SetState(382)
					p.vectorOperation(8)
				}

			case 6:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(384)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(385)
					p.OrOp()
				}
				{
					p.SetState(386)
					p.vectorOperation(7)
				}

			case 7:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(388)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(389)
					p.VectorMatchOp()
				}
				{
					p.SetState(390)
					p.vectorOperation(6)
				}

			case 8:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(393)
					p.Match(PromQLExParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(394)
					p.vectorOperation(5)
				}

			case 9:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(395)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(396)
					p.SubqueryOp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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

// ISubqueryOpContext is an interface to support dynamic dispatch.
type ISubqueryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Subquery_range() ISubquery_rangeContext
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
	p.RuleIndex = PromQLExParserRULE_subqueryOp
	return p
}

func InitEmptySubqueryOpContext(p *SubqueryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_subqueryOp
}

func (*SubqueryOpContext) IsSubqueryOpContext() {}

func NewSubqueryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryOpContext {
	var p = new(SubqueryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_subqueryOp

	return p
}

func (s *SubqueryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryOpContext) Subquery_range() ISubquery_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubquery_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubquery_rangeContext)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterSubqueryOp(s)
	}
}

func (s *SubqueryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitSubqueryOp(s)
	}
}

func (p *PromQLExParser) SubqueryOp() (localctx ISubqueryOpContext) {
	localctx = NewSubqueryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PromQLExParserRULE_subqueryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Subquery_range()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(403)
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
	Duration() IDurationContext

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
	p.RuleIndex = PromQLExParserRULE_offsetOp
	return p
}

func InitEmptyOffsetOpContext(p *OffsetOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_offsetOp
}

func (*OffsetOpContext) IsOffsetOpContext() {}

func NewOffsetOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetOpContext {
	var p = new(OffsetOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_offsetOp

	return p
}

func (s *OffsetOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetOpContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserOFFSET, 0)
}

func (s *OffsetOpContext) Duration() IDurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *OffsetOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterOffsetOp(s)
	}
}

func (s *OffsetOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitOffsetOp(s)
	}
}

func (p *PromQLExParser) OffsetOp() (localctx IOffsetOpContext) {
	localctx = NewOffsetOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PromQLExParserRULE_offsetOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(PromQLExParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Duration()
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
	Time_range() ITime_rangeContext

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
	p.RuleIndex = PromQLExParserRULE_matrixSelector
	return p
}

func InitEmptyMatrixSelectorContext(p *MatrixSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_matrixSelector
}

func (*MatrixSelectorContext) IsMatrixSelectorContext() {}

func NewMatrixSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixSelectorContext {
	var p = new(MatrixSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_matrixSelector

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

func (s *MatrixSelectorContext) Time_range() ITime_rangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITime_rangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITime_rangeContext)
}

func (s *MatrixSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterMatrixSelector(s)
	}
}

func (s *MatrixSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitMatrixSelector(s)
	}
}

func (p *PromQLExParser) MatrixSelector() (localctx IMatrixSelectorContext) {
	localctx = NewMatrixSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PromQLExParserRULE_matrixSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.InstantSelector()
	}
	{
		p.SetState(410)
		p.Time_range()
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
	Duration() IDurationContext
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
	p.RuleIndex = PromQLExParserRULE_offset
	return p
}

func InitEmptyOffsetContext(p *OffsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_offset
}

func (*OffsetContext) IsOffsetContext() {}

func NewOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetContext {
	var p = new(OffsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_offset

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
	return s.GetToken(PromQLExParserOFFSET, 0)
}

func (s *OffsetContext) Duration() IDurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterOffset(s)
	}
}

func (s *OffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitOffset(s)
	}
}

func (p *PromQLExParser) Offset() (localctx IOffsetContext) {
	localctx = NewOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PromQLExParserRULE_offset)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.InstantSelector()
		}
		{
			p.SetState(413)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Duration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.MatrixSelector()
		}
		{
			p.SetState(417)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Duration()
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_num_expression() IConst_num_expressionContext
	String_() IStringContext

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
	p.RuleIndex = PromQLExParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Const_num_expression() IConst_num_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_num_expressionContext)
}

func (s *LiteralContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *PromQLExParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PromQLExParserRULE_literal)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserPOSITIVE_INTEGER, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_PAREN, PromQLExParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.Const_num_expression()
		}

	case PromQLExParserSTRING, PromQLExParserRAW_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)
			p.String_()
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

// IInstantSelectorContext is an interface to support dynamic dispatch.
type IInstantSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Metric_name() IMetric_nameContext
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
	p.RuleIndex = PromQLExParserRULE_instantSelector
	return p
}

func InitEmptyInstantSelectorContext(p *InstantSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_instantSelector
}

func (*InstantSelectorContext) IsInstantSelectorContext() {}

func NewInstantSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstantSelectorContext {
	var p = new(InstantSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_instantSelector

	return p
}

func (s *InstantSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *InstantSelectorContext) Metric_name() IMetric_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetric_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetric_nameContext)
}

func (s *InstantSelectorContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_BRACE, 0)
}

func (s *InstantSelectorContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_BRACE, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterInstantSelector(s)
	}
}

func (s *InstantSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitInstantSelector(s)
	}
}

func (p *PromQLExParser) InstantSelector() (localctx IInstantSelectorContext) {
	localctx = NewInstantSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PromQLExParserRULE_instantSelector)
	var _la int

	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)
			p.Metric_name()
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(427)
				p.Match(PromQLExParserLEFT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(429)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&324539550514806798) != 0 {
				{
					p.SetState(428)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(431)
				p.Match(PromQLExParserRIGHT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case PromQLExParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.Match(PromQLExParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.LabelMatcherList()
		}
		{
			p.SetState(436)
			p.Match(PromQLExParserRIGHT_BRACE)
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

// ILabelNameContext is an interface to support dynamic dispatch.
type ILabelNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Keyword() IKeywordContext
	Metric_name() IMetric_nameContext
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
	p.RuleIndex = PromQLExParserRULE_labelName
	return p
}

func InitEmptyLabelNameContext(p *LabelNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_labelName
}

func (*LabelNameContext) IsLabelNameContext() {}

func NewLabelNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameContext {
	var p = new(LabelNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_labelName

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

func (s *LabelNameContext) Metric_name() IMetric_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetric_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetric_nameContext)
}

func (s *LabelNameContext) LABEL_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLABEL_NAME, 0)
}

func (s *LabelNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLabelName(s)
	}
}

func (s *LabelNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLabelName(s)
	}
}

func (p *PromQLExParser) LabelName() (localctx ILabelNameContext) {
	localctx = NewLabelNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PromQLExParserRULE_labelName)
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserAND, PromQLExParserOR, PromQLExParserUNLESS, PromQLExParserBY, PromQLExParserWITHOUT, PromQLExParserON, PromQLExParserIGNORING, PromQLExParserGROUP_LEFT, PromQLExParserGROUP_RIGHT, PromQLExParserOFFSET, PromQLExParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(440)
			p.Keyword()
		}

	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(441)
			p.Metric_name()
		}

	case PromQLExParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(442)
			p.Match(PromQLExParserLABEL_NAME)
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

// IMetric_nameContext is an interface to support dynamic dispatch.
type IMetric_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	METRIC_NAME() antlr.TerminalNode

	// IsMetric_nameContext differentiates from other interfaces.
	IsMetric_nameContext()
}

type Metric_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetric_nameContext() *Metric_nameContext {
	var p = new(Metric_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_metric_name
	return p
}

func InitEmptyMetric_nameContext(p *Metric_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_metric_name
}

func (*Metric_nameContext) IsMetric_nameContext() {}

func NewMetric_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Metric_nameContext {
	var p = new(Metric_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_metric_name

	return p
}

func (s *Metric_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Metric_nameContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
}

func (s *Metric_nameContext) METRIC_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLExParserMETRIC_NAME, 0)
}

func (s *Metric_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Metric_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Metric_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterMetric_name(s)
	}
}

func (s *Metric_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitMetric_name(s)
	}
}

func (p *PromQLExParser) Metric_name() (localctx IMetric_nameContext) {
	localctx = NewMetric_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PromQLExParserRULE_metric_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserMETRIC_NAME || _la == PromQLExParserID) {
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
	p.RuleIndex = PromQLExParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_expression

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
	return s.GetToken(PromQLExParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PromQLExParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PromQLExParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.vectorOperation(0)
	}
	{
		p.SetState(448)
		p.Match(PromQLExParserEOF)
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
	p.RuleIndex = PromQLExParserRULE_unaryOp
	return p
}

func InitEmptyUnaryOpContext(p *UnaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_unaryOp
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLExParserADD, 0)
}

func (s *UnaryOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (p *PromQLExParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PromQLExParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserADD || _la == PromQLExParserSUB) {
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
	p.RuleIndex = PromQLExParserRULE_powOp
	return p
}

func InitEmptyPowOpContext(p *PowOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_powOp
}

func (*PowOpContext) IsPowOpContext() {}

func NewPowOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowOpContext {
	var p = new(PowOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_powOp

	return p
}

func (s *PowOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PowOpContext) POW() antlr.TerminalNode {
	return s.GetToken(PromQLExParserPOW, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterPowOp(s)
	}
}

func (s *PowOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitPowOp(s)
	}
}

func (p *PromQLExParser) PowOp() (localctx IPowOpContext) {
	localctx = NewPowOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, PromQLExParserRULE_powOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(PromQLExParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(453)
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
	p.RuleIndex = PromQLExParserRULE_multOp
	return p
}

func InitEmptyMultOpContext(p *MultOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_multOp
}

func (*MultOpContext) IsMultOpContext() {}

func NewMultOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOpContext {
	var p = new(MultOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_multOp

	return p
}

func (s *MultOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOpContext) MULT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserMULT, 0)
}

func (s *MultOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDIV, 0)
}

func (s *MultOpContext) MOD() antlr.TerminalNode {
	return s.GetToken(PromQLExParserMOD, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterMultOp(s)
	}
}

func (s *MultOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitMultOp(s)
	}
}

func (p *PromQLExParser) MultOp() (localctx IMultOpContext) {
	localctx = NewMultOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, PromQLExParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&117440512) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(457)
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
	p.RuleIndex = PromQLExParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLExParserADD, 0)
}

func (s *AddOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLExParserSUB, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (p *PromQLExParser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, PromQLExParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserADD || _la == PromQLExParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(461)
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
	p.RuleIndex = PromQLExParserRULE_compareOp
	return p
}

func InitEmptyCompareOpContext(p *CompareOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_compareOp
}

func (*CompareOpContext) IsCompareOpContext() {}

func NewCompareOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOpContext {
	var p = new(CompareOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_compareOp

	return p
}

func (s *CompareOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOpContext) DEQ() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDEQ, 0)
}

func (s *CompareOpContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNE, 0)
}

func (s *CompareOpContext) GT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGT, 0)
}

func (s *CompareOpContext) LT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLT, 0)
}

func (s *CompareOpContext) GE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGE, 0)
}

func (s *CompareOpContext) LE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLE, 0)
}

func (s *CompareOpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLExParserBOOL, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterCompareOp(s)
	}
}

func (s *CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitCompareOp(s)
	}
}

func (p *PromQLExParser) CompareOp() (localctx ICompareOpContext) {
	localctx = NewCompareOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, PromQLExParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&270582939648) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserBOOL {
		{
			p.SetState(465)
			p.Match(PromQLExParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(468)
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
	p.RuleIndex = PromQLExParserRULE_andUnlessOp
	return p
}

func InitEmptyAndUnlessOpContext(p *AndUnlessOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_andUnlessOp
}

func (*AndUnlessOpContext) IsAndUnlessOpContext() {}

func NewAndUnlessOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndUnlessOpContext {
	var p = new(AndUnlessOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_andUnlessOp

	return p
}

func (s *AndUnlessOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AndUnlessOpContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLExParserAND, 0)
}

func (s *AndUnlessOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserUNLESS, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAndUnlessOp(s)
	}
}

func (s *AndUnlessOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAndUnlessOp(s)
	}
}

func (p *PromQLExParser) AndUnlessOp() (localctx IAndUnlessOpContext) {
	localctx = NewAndUnlessOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, PromQLExParserRULE_andUnlessOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserAND || _la == PromQLExParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(472)
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
	p.RuleIndex = PromQLExParserRULE_orOp
	return p
}

func InitEmptyOrOpContext(p *OrOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_orOp
}

func (*OrOpContext) IsOrOpContext() {}

func NewOrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrOpContext {
	var p = new(OrOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_orOp

	return p
}

func (s *OrOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OrOpContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLExParserOR, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterOrOp(s)
	}
}

func (s *OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitOrOp(s)
	}
}

func (p *PromQLExParser) OrOp() (localctx IOrOpContext) {
	localctx = NewOrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, PromQLExParserRULE_orOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(PromQLExParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(476)
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
	p.RuleIndex = PromQLExParserRULE_vectorMatchOp
	return p
}

func InitEmptyVectorMatchOpContext(p *VectorMatchOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_vectorMatchOp
}

func (*VectorMatchOpContext) IsVectorMatchOpContext() {}

func NewVectorMatchOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorMatchOpContext {
	var p = new(VectorMatchOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_vectorMatchOp

	return p
}

func (s *VectorMatchOpContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorMatchOpContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExParserON, 0)
}

func (s *VectorMatchOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserUNLESS, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVectorMatchOp(s)
	}
}

func (s *VectorMatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVectorMatchOp(s)
	}
}

func (p *PromQLExParser) VectorMatchOp() (localctx IVectorMatchOpContext) {
	localctx = NewVectorMatchOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, PromQLExParserRULE_vectorMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserUNLESS || _la == PromQLExParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(480)
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
	p.RuleIndex = PromQLExParserRULE_vector
	return p
}

func InitEmptyVectorContext(p *VectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_vector
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_vector

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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *PromQLExParser) Vector() (localctx IVectorContext) {
	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, PromQLExParserRULE_vector)
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(484)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(485)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(486)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(487)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(488)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(489)
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
	p.RuleIndex = PromQLExParserRULE_parens
	return p
}

func InitEmptyParensContext(p *ParensContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_parens
}

func (*ParensContext) IsParensContext() {}

func NewParensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParensContext {
	var p = new(ParensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_parens

	return p
}

func (s *ParensContext) GetParser() antlr.Parser { return s.parser }

func (s *ParensContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
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
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitParens(s)
	}
}

func (p *PromQLExParser) Parens() (localctx IParensContext) {
	localctx = NewParensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, PromQLExParserRULE_parens)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(493)
		p.vectorOperation(0)
	}
	{
		p.SetState(494)
		p.Match(PromQLExParserRIGHT_PAREN)
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

// ILabelMatcherContext is an interface to support dynamic dispatch.
type ILabelMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabelName() ILabelNameContext
	LabelMatcherOperator() ILabelMatcherOperatorContext
	String_() IStringContext

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
	p.RuleIndex = PromQLExParserRULE_labelMatcher
	return p
}

func InitEmptyLabelMatcherContext(p *LabelMatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_labelMatcher
}

func (*LabelMatcherContext) IsLabelMatcherContext() {}

func NewLabelMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherContext {
	var p = new(LabelMatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_labelMatcher

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

func (s *LabelMatcherContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LabelMatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLabelMatcher(s)
	}
}

func (p *PromQLExParser) LabelMatcher() (localctx ILabelMatcherContext) {
	localctx = NewLabelMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, PromQLExParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.LabelName()
	}
	{
		p.SetState(497)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(498)
		p.String_()
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
	p.RuleIndex = PromQLExParserRULE_labelMatcherOperator
	return p
}

func InitEmptyLabelMatcherOperatorContext(p *LabelMatcherOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_labelMatcherOperator
}

func (*LabelMatcherOperatorContext) IsLabelMatcherOperatorContext() {}

func NewLabelMatcherOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherOperatorContext {
	var p = new(LabelMatcherOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_labelMatcherOperator

	return p
}

func (s *LabelMatcherOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(PromQLExParserEQ, 0)
}

func (s *LabelMatcherOperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNE, 0)
}

func (s *LabelMatcherOperatorContext) RE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRE, 0)
}

func (s *LabelMatcherOperatorContext) NRE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNRE, 0)
}

func (s *LabelMatcherOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLabelMatcherOperator(s)
	}
}

func (s *LabelMatcherOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLabelMatcherOperator(s)
	}
}

func (p *PromQLExParser) LabelMatcherOperator() (localctx ILabelMatcherOperatorContext) {
	localctx = NewLabelMatcherOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, PromQLExParserRULE_labelMatcherOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&835371139072) != 0) {
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
	p.RuleIndex = PromQLExParserRULE_labelMatcherList
	return p
}

func InitEmptyLabelMatcherListContext(p *LabelMatcherListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_labelMatcherList
}

func (*LabelMatcherListContext) IsLabelMatcherListContext() {}

func NewLabelMatcherListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherListContext {
	var p = new(LabelMatcherListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_labelMatcherList

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
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *LabelMatcherListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *LabelMatcherListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLabelMatcherList(s)
	}
}

func (s *LabelMatcherListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLabelMatcherList(s)
	}
}

func (p *PromQLExParser) LabelMatcherList() (localctx ILabelMatcherListContext) {
	localctx = NewLabelMatcherListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, PromQLExParserRULE_labelMatcherList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.LabelMatcher()
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(503)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(504)
				p.LabelMatcher()
			}

		}
		p.SetState(509)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserCOMMA {
		{
			p.SetState(510)
			p.Match(PromQLExParserCOMMA)
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
	p.RuleIndex = PromQLExParserRULE_function_
	return p
}

func InitEmptyFunction_Context(p *Function_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_function_
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLExParserFUNCTION, 0)
}

func (s *Function_Context) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *Function_Context) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *Function_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *PromQLExParser) Function_() (localctx IFunction_Context) {
	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, PromQLExParserRULE_function_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(PromQLExParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&884112901863965710) != 0 {
		{
			p.SetState(515)
			p.Parameter()
		}
		p.SetState(520)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(516)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(517)
				p.Parameter()
			}

			p.SetState(522)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(525)
		p.Match(PromQLExParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLExParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_parameter

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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *PromQLExParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, PromQLExParserRULE_parameter)
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(527)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(528)
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
	p.RuleIndex = PromQLExParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *ParameterListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *PromQLExParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, PromQLExParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&884112901863965710) != 0 {
		{
			p.SetState(532)
			p.Parameter()
		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(533)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(534)
				p.Parameter()
			}

			p.SetState(539)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(542)
		p.Match(PromQLExParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLExParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLExParserAGGREGATION_OPERATOR, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (p *PromQLExParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, PromQLExParserRULE_aggregation)
	p.SetState(559)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(546)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(549)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(547)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(548)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(551)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(554)
			p.ParameterList()
		}
		p.SetState(557)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(555)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(556)
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
	p.RuleIndex = PromQLExParserRULE_by
	return p
}

func InitEmptyByContext(p *ByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_by
}

func (*ByContext) IsByContext() {}

func NewByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByContext {
	var p = new(ByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_by

	return p
}

func (s *ByContext) GetParser() antlr.Parser { return s.parser }

func (s *ByContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLExParserBY, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterBy(s)
	}
}

func (s *ByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitBy(s)
	}
}

func (p *PromQLExParser) By() (localctx IByContext) {
	localctx = NewByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, PromQLExParserRULE_by)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.Match(PromQLExParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(562)
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
	p.RuleIndex = PromQLExParserRULE_without
	return p
}

func InitEmptyWithoutContext(p *WithoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_without
}

func (*WithoutContext) IsWithoutContext() {}

func NewWithoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithoutContext {
	var p = new(WithoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_without

	return p
}

func (s *WithoutContext) GetParser() antlr.Parser { return s.parser }

func (s *WithoutContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserWITHOUT, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterWithout(s)
	}
}

func (s *WithoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitWithout(s)
	}
}

func (p *PromQLExParser) Without() (localctx IWithoutContext) {
	localctx = NewWithoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, PromQLExParserRULE_without)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(564)
		p.Match(PromQLExParserWITHOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(565)
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
	p.RuleIndex = PromQLExParserRULE_grouping
	return p
}

func InitEmptyGroupingContext(p *GroupingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_grouping
}

func (*GroupingContext) IsGroupingContext() {}

func NewGroupingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupingContext {
	var p = new(GroupingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_grouping

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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (p *PromQLExParser) Grouping() (localctx IGroupingContext) {
	localctx = NewGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, PromQLExParserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(569)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserON:
		{
			p.SetState(567)
			p.On_()
		}

	case PromQLExParserIGNORING:
		{
			p.SetState(568)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLExParserGROUP_LEFT:
		{
			p.SetState(571)
			p.GroupLeft()
		}

	case PromQLExParserGROUP_RIGHT:
		{
			p.SetState(572)
			p.GroupRight()
		}

	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserMETRIC_NAME, PromQLExParserPOSITIVE_INTEGER, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserSTRING, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_BRACE, PromQLExParserLEFT_PAREN, PromQLExParserDURATION, PromQLExParserID, PromQLExParserRAW_STRING:

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
	p.RuleIndex = PromQLExParserRULE_on_
	return p
}

func InitEmptyOn_Context(p *On_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_on_
}

func (*On_Context) IsOn_Context() {}

func NewOn_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *On_Context {
	var p = new(On_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_on_

	return p
}

func (s *On_Context) GetParser() antlr.Parser { return s.parser }

func (s *On_Context) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExParserON, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterOn_(s)
	}
}

func (s *On_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitOn_(s)
	}
}

func (p *PromQLExParser) On_() (localctx IOn_Context) {
	localctx = NewOn_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, PromQLExParserRULE_on_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(575)
		p.Match(PromQLExParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(576)
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
	p.RuleIndex = PromQLExParserRULE_ignoring
	return p
}

func InitEmptyIgnoringContext(p *IgnoringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_ignoring
}

func (*IgnoringContext) IsIgnoringContext() {}

func NewIgnoringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoringContext {
	var p = new(IgnoringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_ignoring

	return p
}

func (s *IgnoringContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoringContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLExParserIGNORING, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterIgnoring(s)
	}
}

func (s *IgnoringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitIgnoring(s)
	}
}

func (p *PromQLExParser) Ignoring() (localctx IIgnoringContext) {
	localctx = NewIgnoringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, PromQLExParserRULE_ignoring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)
		p.Match(PromQLExParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(579)
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
	p.RuleIndex = PromQLExParserRULE_groupLeft
	return p
}

func InitEmptyGroupLeftContext(p *GroupLeftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_groupLeft
}

func (*GroupLeftContext) IsGroupLeftContext() {}

func NewGroupLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupLeftContext {
	var p = new(GroupLeftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_groupLeft

	return p
}

func (s *GroupLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupLeftContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGROUP_LEFT, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterGroupLeft(s)
	}
}

func (s *GroupLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitGroupLeft(s)
	}
}

func (p *PromQLExParser) GroupLeft() (localctx IGroupLeftContext) {
	localctx = NewGroupLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, PromQLExParserRULE_groupLeft)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(581)
		p.Match(PromQLExParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(583)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(582)
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
	p.RuleIndex = PromQLExParserRULE_groupRight
	return p
}

func InitEmptyGroupRightContext(p *GroupRightContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_groupRight
}

func (*GroupRightContext) IsGroupRightContext() {}

func NewGroupRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupRightContext {
	var p = new(GroupRightContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_groupRight

	return p
}

func (s *GroupRightContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupRightContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGROUP_RIGHT, 0)
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
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterGroupRight(s)
	}
}

func (s *GroupRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitGroupRight(s)
	}
}

func (p *PromQLExParser) GroupRight() (localctx IGroupRightContext) {
	localctx = NewGroupRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, PromQLExParserRULE_groupRight)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(585)
		p.Match(PromQLExParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(587)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(586)
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
	p.RuleIndex = PromQLExParserRULE_labelNameList
	return p
}

func InitEmptyLabelNameListContext(p *LabelNameListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_labelNameList
}

func (*LabelNameListContext) IsLabelNameListContext() {}

func NewLabelNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameListContext {
	var p = new(LabelNameListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_labelNameList

	return p
}

func (s *LabelNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelNameListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *LabelNameListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLExParserCOMMA)
}

func (s *LabelNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserCOMMA, i)
}

func (s *LabelNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLabelNameList(s)
	}
}

func (s *LabelNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLabelNameList(s)
	}
}

func (p *PromQLExParser) LabelNameList() (localctx ILabelNameListContext) {
	localctx = NewLabelNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, PromQLExParserRULE_labelNameList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&324539550514806798) != 0 {
		{
			p.SetState(590)
			p.LabelName()
		}
		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(591)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(592)
				p.LabelName()
			}

			p.SetState(597)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(600)
		p.Match(PromQLExParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLExParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLExParserAND, 0)
}

func (s *KeywordContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLExParserOR, 0)
}

func (s *KeywordContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLExParserUNLESS, 0)
}

func (s *KeywordContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLExParserBY, 0)
}

func (s *KeywordContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserWITHOUT, 0)
}

func (s *KeywordContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLExParserON, 0)
}

func (s *KeywordContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLExParserIGNORING, 0)
}

func (s *KeywordContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGROUP_LEFT, 0)
}

func (s *KeywordContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserGROUP_RIGHT, 0)
}

func (s *KeywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLExParserOFFSET, 0)
}

func (s *KeywordContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLExParserBOOL, 0)
}

func (s *KeywordContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLExParserAGGREGATION_OPERATOR, 0)
}

func (s *KeywordContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLExParserFUNCTION, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *PromQLExParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, PromQLExParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(602)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&280377344131078) != 0) {
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

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	RAW_STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLExParserSTRING, 0)
}

func (s *StringContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRAW_STRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *PromQLExParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, PromQLExParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(604)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserSTRING || _la == PromQLExParserRAW_STRING) {
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

func (p *PromQLExParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 37:
		var t *VectorOperationContext = nil
		if localctx != nil {
			t = localctx.(*VectorOperationContext)
		}
		return p.VectorOperation_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLExParser) VectorOperation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

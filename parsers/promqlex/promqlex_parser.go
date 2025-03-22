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
		"ID", "RAW_STRING", "BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"promqlx", "statement", "alias_def", "alias_call", "macro_def", "macro_call",
		"args_decl", "arg_name", "statement_block", "arg_list", "if_statement",
		"condition", "compareVectorOperation", "trueConst", "falseConst", "time_instant_literal",
		"const_num_expression", "num_literal", "duration", "time_range", "subquery_range",
		"vectorOperation", "subqueryOp", "offsetOp", "matrixSelector", "offset",
		"literal", "instantSelector", "labelName", "metric_name", "at_modifier_timestamp",
		"expression", "unaryOp", "powOp", "multOp", "addOp", "compareOp", "andUnlessOp",
		"orOp", "vectorMatchOp", "vector", "parens", "labelMatcher", "labelMatcherOperator",
		"labelMatcherList", "function_", "parameter", "parameterList", "aggregation",
		"by", "without", "grouping", "on_", "ignoring", "groupLeft", "groupRight",
		"labelNameList", "keyword", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 512, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		2, 58, 7, 58, 1, 0, 1, 0, 1, 0, 5, 0, 122, 8, 0, 10, 0, 12, 0, 125, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 133, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 147, 8,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 157, 8, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 164, 8, 6, 10, 6, 12, 6, 167, 9, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 4, 9, 178, 8, 9, 11,
		9, 12, 9, 179, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11,
		189, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 3, 15, 201, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 212, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 226, 8, 16,
		10, 16, 12, 16, 229, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 235, 8,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 247, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 259, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 294, 8, 21, 10,
		21, 12, 21, 297, 9, 21, 1, 22, 1, 22, 3, 22, 301, 8, 22, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 3, 25, 317, 8, 25, 1, 26, 1, 26, 3, 26, 321, 8, 26, 1, 27, 1,
		27, 1, 27, 3, 27, 326, 8, 27, 1, 27, 3, 27, 329, 8, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 3, 27, 335, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 340, 8, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 351,
		8, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 360, 8,
		33, 1, 34, 1, 34, 3, 34, 364, 8, 34, 1, 35, 1, 35, 3, 35, 368, 8, 35, 1,
		36, 1, 36, 3, 36, 372, 8, 36, 1, 36, 3, 36, 375, 8, 36, 1, 37, 1, 37, 3,
		37, 379, 8, 37, 1, 38, 1, 38, 3, 38, 383, 8, 38, 1, 39, 1, 39, 3, 39, 387,
		8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 396, 8,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 44, 1, 44, 1, 44, 5, 44, 411, 8, 44, 10, 44, 12, 44, 414, 9, 44, 1,
		44, 3, 44, 417, 8, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 424, 8,
		45, 10, 45, 12, 45, 427, 9, 45, 3, 45, 429, 8, 45, 1, 45, 1, 45, 1, 46,
		1, 46, 3, 46, 435, 8, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 441, 8, 47,
		10, 47, 12, 47, 444, 9, 47, 3, 47, 446, 8, 47, 1, 47, 1, 47, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 3, 48, 455, 8, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 3, 48, 463, 8, 48, 3, 48, 465, 8, 48, 1, 49, 1, 49, 1, 49,
		1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 3, 51, 475, 8, 51, 1, 51, 1, 51, 3,
		51, 479, 8, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54,
		3, 54, 489, 8, 54, 1, 55, 1, 55, 3, 55, 493, 8, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 5, 56, 499, 8, 56, 10, 56, 12, 56, 502, 9, 56, 3, 56, 504, 8, 56,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 0, 2, 32, 42, 59, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
		78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
		112, 114, 116, 0, 9, 2, 0, 5, 5, 56, 56, 1, 0, 18, 19, 1, 0, 20, 22, 1,
		0, 30, 35, 2, 0, 24, 24, 26, 26, 2, 0, 26, 26, 40, 40, 3, 0, 29, 29, 31,
		31, 36, 37, 3, 0, 1, 2, 24, 26, 38, 45, 2, 0, 17, 17, 57, 57, 527, 0, 118,
		1, 0, 0, 0, 2, 132, 1, 0, 0, 0, 4, 134, 1, 0, 0, 0, 6, 139, 1, 0, 0, 0,
		8, 142, 1, 0, 0, 0, 10, 151, 1, 0, 0, 0, 12, 160, 1, 0, 0, 0, 14, 168,
		1, 0, 0, 0, 16, 170, 1, 0, 0, 0, 18, 174, 1, 0, 0, 0, 20, 181, 1, 0, 0,
		0, 22, 188, 1, 0, 0, 0, 24, 190, 1, 0, 0, 0, 26, 194, 1, 0, 0, 0, 28, 196,
		1, 0, 0, 0, 30, 200, 1, 0, 0, 0, 32, 211, 1, 0, 0, 0, 34, 234, 1, 0, 0,
		0, 36, 236, 1, 0, 0, 0, 38, 238, 1, 0, 0, 0, 40, 242, 1, 0, 0, 0, 42, 258,
		1, 0, 0, 0, 44, 298, 1, 0, 0, 0, 46, 302, 1, 0, 0, 0, 48, 305, 1, 0, 0,
		0, 50, 316, 1, 0, 0, 0, 52, 320, 1, 0, 0, 0, 54, 334, 1, 0, 0, 0, 56, 339,
		1, 0, 0, 0, 58, 341, 1, 0, 0, 0, 60, 350, 1, 0, 0, 0, 62, 352, 1, 0, 0,
		0, 64, 355, 1, 0, 0, 0, 66, 357, 1, 0, 0, 0, 68, 361, 1, 0, 0, 0, 70, 365,
		1, 0, 0, 0, 72, 369, 1, 0, 0, 0, 74, 376, 1, 0, 0, 0, 76, 380, 1, 0, 0,
		0, 78, 384, 1, 0, 0, 0, 80, 395, 1, 0, 0, 0, 82, 397, 1, 0, 0, 0, 84, 401,
		1, 0, 0, 0, 86, 405, 1, 0, 0, 0, 88, 407, 1, 0, 0, 0, 90, 418, 1, 0, 0,
		0, 92, 434, 1, 0, 0, 0, 94, 436, 1, 0, 0, 0, 96, 464, 1, 0, 0, 0, 98, 466,
		1, 0, 0, 0, 100, 469, 1, 0, 0, 0, 102, 474, 1, 0, 0, 0, 104, 480, 1, 0,
		0, 0, 106, 483, 1, 0, 0, 0, 108, 486, 1, 0, 0, 0, 110, 490, 1, 0, 0, 0,
		112, 494, 1, 0, 0, 0, 114, 507, 1, 0, 0, 0, 116, 509, 1, 0, 0, 0, 118,
		123, 3, 2, 1, 0, 119, 120, 5, 9, 0, 0, 120, 122, 3, 2, 1, 0, 121, 119,
		1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0,
		0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5, 0, 0, 1,
		127, 1, 1, 0, 0, 0, 128, 133, 3, 4, 2, 0, 129, 133, 3, 8, 4, 0, 130, 133,
		3, 20, 10, 0, 131, 133, 3, 42, 21, 0, 132, 128, 1, 0, 0, 0, 132, 129, 1,
		0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 3, 1, 0, 0, 0,
		134, 135, 5, 12, 0, 0, 135, 136, 5, 56, 0, 0, 136, 137, 5, 29, 0, 0, 137,
		138, 3, 42, 21, 0, 138, 5, 1, 0, 0, 0, 139, 140, 5, 13, 0, 0, 140, 141,
		5, 56, 0, 0, 141, 7, 1, 0, 0, 0, 142, 143, 5, 12, 0, 0, 143, 144, 5, 56,
		0, 0, 144, 146, 5, 48, 0, 0, 145, 147, 3, 12, 6, 0, 146, 145, 1, 0, 0,
		0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 5, 49, 0, 0, 149,
		150, 3, 16, 8, 0, 150, 9, 1, 0, 0, 0, 151, 152, 5, 13, 0, 0, 152, 153,
		5, 56, 0, 0, 153, 154, 5, 56, 0, 0, 154, 156, 5, 48, 0, 0, 155, 157, 3,
		18, 9, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 1, 0, 0,
		0, 158, 159, 5, 49, 0, 0, 159, 11, 1, 0, 0, 0, 160, 165, 3, 14, 7, 0, 161,
		162, 5, 50, 0, 0, 162, 164, 3, 14, 7, 0, 163, 161, 1, 0, 0, 0, 164, 167,
		1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 13, 1, 0,
		0, 0, 167, 165, 1, 0, 0, 0, 168, 169, 5, 56, 0, 0, 169, 15, 1, 0, 0, 0,
		170, 171, 5, 46, 0, 0, 171, 172, 3, 0, 0, 0, 172, 173, 5, 47, 0, 0, 173,
		17, 1, 0, 0, 0, 174, 177, 3, 42, 21, 0, 175, 176, 5, 50, 0, 0, 176, 178,
		3, 42, 21, 0, 177, 175, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1,
		0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 19, 1, 0, 0, 0, 181, 182, 5, 6, 0,
		0, 182, 183, 3, 22, 11, 0, 183, 184, 3, 16, 8, 0, 184, 21, 1, 0, 0, 0,
		185, 189, 3, 24, 12, 0, 186, 189, 3, 26, 13, 0, 187, 189, 3, 28, 14, 0,
		188, 185, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 187, 1, 0, 0, 0, 189,
		23, 1, 0, 0, 0, 190, 191, 3, 42, 21, 0, 191, 192, 3, 72, 36, 0, 192, 193,
		3, 42, 21, 0, 193, 25, 1, 0, 0, 0, 194, 195, 5, 7, 0, 0, 195, 27, 1, 0,
		0, 0, 196, 197, 5, 8, 0, 0, 197, 29, 1, 0, 0, 0, 198, 201, 5, 3, 0, 0,
		199, 201, 5, 16, 0, 0, 200, 198, 1, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		31, 1, 0, 0, 0, 202, 203, 6, 16, -1, 0, 203, 204, 3, 64, 32, 0, 204, 205,
		3, 32, 16, 5, 205, 212, 1, 0, 0, 0, 206, 207, 5, 48, 0, 0, 207, 208, 3,
		32, 16, 0, 208, 209, 5, 49, 0, 0, 209, 212, 1, 0, 0, 0, 210, 212, 3, 34,
		17, 0, 211, 202, 1, 0, 0, 0, 211, 206, 1, 0, 0, 0, 211, 210, 1, 0, 0, 0,
		212, 227, 1, 0, 0, 0, 213, 214, 10, 6, 0, 0, 214, 215, 3, 66, 33, 0, 215,
		216, 3, 32, 16, 6, 216, 226, 1, 0, 0, 0, 217, 218, 10, 4, 0, 0, 218, 219,
		3, 68, 34, 0, 219, 220, 3, 32, 16, 5, 220, 226, 1, 0, 0, 0, 221, 222, 10,
		3, 0, 0, 222, 223, 3, 70, 35, 0, 223, 224, 3, 32, 16, 4, 224, 226, 1, 0,
		0, 0, 225, 213, 1, 0, 0, 0, 225, 217, 1, 0, 0, 0, 225, 221, 1, 0, 0, 0,
		226, 229, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		33, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 230, 235, 5, 16, 0, 0, 231, 235,
		5, 52, 0, 0, 232, 235, 3, 30, 15, 0, 233, 235, 3, 6, 3, 0, 234, 230, 1,
		0, 0, 0, 234, 231, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 233, 1, 0, 0,
		0, 235, 35, 1, 0, 0, 0, 236, 237, 3, 32, 16, 0, 237, 37, 1, 0, 0, 0, 238,
		239, 5, 14, 0, 0, 239, 240, 3, 36, 18, 0, 240, 241, 5, 15, 0, 0, 241, 39,
		1, 0, 0, 0, 242, 243, 5, 14, 0, 0, 243, 244, 3, 36, 18, 0, 244, 246, 5,
		4, 0, 0, 245, 247, 3, 36, 18, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1, 0,
		0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 5, 15, 0, 0, 249, 41, 1, 0, 0, 0,
		250, 251, 6, 21, -1, 0, 251, 259, 3, 32, 16, 0, 252, 253, 3, 64, 32, 0,
		253, 254, 3, 42, 21, 11, 254, 259, 1, 0, 0, 0, 255, 259, 3, 80, 40, 0,
		256, 259, 3, 10, 5, 0, 257, 259, 3, 6, 3, 0, 258, 250, 1, 0, 0, 0, 258,
		252, 1, 0, 0, 0, 258, 255, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 258, 257,
		1, 0, 0, 0, 259, 295, 1, 0, 0, 0, 260, 261, 10, 13, 0, 0, 261, 262, 3,
		66, 33, 0, 262, 263, 3, 42, 21, 13, 263, 294, 1, 0, 0, 0, 264, 265, 10,
		10, 0, 0, 265, 266, 3, 68, 34, 0, 266, 267, 3, 42, 21, 11, 267, 294, 1,
		0, 0, 0, 268, 269, 10, 9, 0, 0, 269, 270, 3, 70, 35, 0, 270, 271, 3, 42,
		21, 10, 271, 294, 1, 0, 0, 0, 272, 273, 10, 8, 0, 0, 273, 274, 3, 72, 36,
		0, 274, 275, 3, 42, 21, 9, 275, 294, 1, 0, 0, 0, 276, 277, 10, 7, 0, 0,
		277, 278, 3, 74, 37, 0, 278, 279, 3, 42, 21, 8, 279, 294, 1, 0, 0, 0, 280,
		281, 10, 6, 0, 0, 281, 282, 3, 76, 38, 0, 282, 283, 3, 42, 21, 7, 283,
		294, 1, 0, 0, 0, 284, 285, 10, 5, 0, 0, 285, 286, 3, 78, 39, 0, 286, 287,
		3, 42, 21, 6, 287, 294, 1, 0, 0, 0, 288, 289, 10, 12, 0, 0, 289, 294, 3,
		44, 22, 0, 290, 291, 10, 4, 0, 0, 291, 292, 5, 51, 0, 0, 292, 294, 3, 60,
		30, 0, 293, 260, 1, 0, 0, 0, 293, 264, 1, 0, 0, 0, 293, 268, 1, 0, 0, 0,
		293, 272, 1, 0, 0, 0, 293, 276, 1, 0, 0, 0, 293, 280, 1, 0, 0, 0, 293,
		284, 1, 0, 0, 0, 293, 288, 1, 0, 0, 0, 293, 290, 1, 0, 0, 0, 294, 297,
		1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 43, 1, 0,
		0, 0, 297, 295, 1, 0, 0, 0, 298, 300, 3, 40, 20, 0, 299, 301, 3, 46, 23,
		0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 45, 1, 0, 0, 0, 302,
		303, 5, 44, 0, 0, 303, 304, 3, 36, 18, 0, 304, 47, 1, 0, 0, 0, 305, 306,
		3, 54, 27, 0, 306, 307, 3, 38, 19, 0, 307, 49, 1, 0, 0, 0, 308, 309, 3,
		54, 27, 0, 309, 310, 5, 44, 0, 0, 310, 311, 3, 36, 18, 0, 311, 317, 1,
		0, 0, 0, 312, 313, 3, 48, 24, 0, 313, 314, 5, 44, 0, 0, 314, 315, 3, 36,
		18, 0, 315, 317, 1, 0, 0, 0, 316, 308, 1, 0, 0, 0, 316, 312, 1, 0, 0, 0,
		317, 51, 1, 0, 0, 0, 318, 321, 3, 32, 16, 0, 319, 321, 3, 116, 58, 0, 320,
		318, 1, 0, 0, 0, 320, 319, 1, 0, 0, 0, 321, 53, 1, 0, 0, 0, 322, 328, 3,
		58, 29, 0, 323, 325, 5, 46, 0, 0, 324, 326, 3, 88, 44, 0, 325, 324, 1,
		0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 5, 47, 0,
		0, 328, 323, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 335, 1, 0, 0, 0, 330,
		331, 5, 46, 0, 0, 331, 332, 3, 88, 44, 0, 332, 333, 5, 47, 0, 0, 333, 335,
		1, 0, 0, 0, 334, 322, 1, 0, 0, 0, 334, 330, 1, 0, 0, 0, 335, 55, 1, 0,
		0, 0, 336, 340, 3, 114, 57, 0, 337, 340, 3, 58, 29, 0, 338, 340, 5, 53,
		0, 0, 339, 336, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 338, 1, 0, 0, 0,
		340, 57, 1, 0, 0, 0, 341, 342, 7, 0, 0, 0, 342, 59, 1, 0, 0, 0, 343, 351,
		3, 32, 16, 0, 344, 345, 5, 27, 0, 0, 345, 346, 5, 48, 0, 0, 346, 351, 5,
		49, 0, 0, 347, 348, 5, 28, 0, 0, 348, 349, 5, 48, 0, 0, 349, 351, 5, 49,
		0, 0, 350, 343, 1, 0, 0, 0, 350, 344, 1, 0, 0, 0, 350, 347, 1, 0, 0, 0,
		351, 61, 1, 0, 0, 0, 352, 353, 3, 42, 21, 0, 353, 354, 5, 0, 0, 1, 354,
		63, 1, 0, 0, 0, 355, 356, 7, 1, 0, 0, 356, 65, 1, 0, 0, 0, 357, 359, 5,
		23, 0, 0, 358, 360, 3, 102, 51, 0, 359, 358, 1, 0, 0, 0, 359, 360, 1, 0,
		0, 0, 360, 67, 1, 0, 0, 0, 361, 363, 7, 2, 0, 0, 362, 364, 3, 102, 51,
		0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 69, 1, 0, 0, 0, 365,
		367, 7, 1, 0, 0, 366, 368, 3, 102, 51, 0, 367, 366, 1, 0, 0, 0, 367, 368,
		1, 0, 0, 0, 368, 71, 1, 0, 0, 0, 369, 371, 7, 3, 0, 0, 370, 372, 5, 45,
		0, 0, 371, 370, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0,
		373, 375, 3, 102, 51, 0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375,
		73, 1, 0, 0, 0, 376, 378, 7, 4, 0, 0, 377, 379, 3, 102, 51, 0, 378, 377,
		1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 75, 1, 0, 0, 0, 380, 382, 5, 25,
		0, 0, 381, 383, 3, 102, 51, 0, 382, 381, 1, 0, 0, 0, 382, 383, 1, 0, 0,
		0, 383, 77, 1, 0, 0, 0, 384, 386, 7, 5, 0, 0, 385, 387, 3, 102, 51, 0,
		386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 79, 1, 0, 0, 0, 388, 396,
		3, 90, 45, 0, 389, 396, 3, 96, 48, 0, 390, 396, 3, 54, 27, 0, 391, 396,
		3, 48, 24, 0, 392, 396, 3, 50, 25, 0, 393, 396, 3, 52, 26, 0, 394, 396,
		3, 82, 41, 0, 395, 388, 1, 0, 0, 0, 395, 389, 1, 0, 0, 0, 395, 390, 1,
		0, 0, 0, 395, 391, 1, 0, 0, 0, 395, 392, 1, 0, 0, 0, 395, 393, 1, 0, 0,
		0, 395, 394, 1, 0, 0, 0, 396, 81, 1, 0, 0, 0, 397, 398, 5, 48, 0, 0, 398,
		399, 3, 42, 21, 0, 399, 400, 5, 49, 0, 0, 400, 83, 1, 0, 0, 0, 401, 402,
		3, 56, 28, 0, 402, 403, 3, 86, 43, 0, 403, 404, 3, 116, 58, 0, 404, 85,
		1, 0, 0, 0, 405, 406, 7, 6, 0, 0, 406, 87, 1, 0, 0, 0, 407, 412, 3, 84,
		42, 0, 408, 409, 5, 50, 0, 0, 409, 411, 3, 84, 42, 0, 410, 408, 1, 0, 0,
		0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413,
		416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 417, 5, 50, 0, 0, 416, 415,
		1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 89, 1, 0, 0, 0, 418, 419, 5, 2,
		0, 0, 419, 428, 5, 48, 0, 0, 420, 425, 3, 92, 46, 0, 421, 422, 5, 50, 0,
		0, 422, 424, 3, 92, 46, 0, 423, 421, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0,
		425, 423, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0, 427,
		425, 1, 0, 0, 0, 428, 420, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430,
		1, 0, 0, 0, 430, 431, 5, 49, 0, 0, 431, 91, 1, 0, 0, 0, 432, 435, 3, 52,
		26, 0, 433, 435, 3, 42, 21, 0, 434, 432, 1, 0, 0, 0, 434, 433, 1, 0, 0,
		0, 435, 93, 1, 0, 0, 0, 436, 445, 5, 48, 0, 0, 437, 442, 3, 92, 46, 0,
		438, 439, 5, 50, 0, 0, 439, 441, 3, 92, 46, 0, 440, 438, 1, 0, 0, 0, 441,
		444, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 446,
		1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 437, 1, 0, 0, 0, 445, 446, 1, 0,
		0, 0, 446, 447, 1, 0, 0, 0, 447, 448, 5, 49, 0, 0, 448, 95, 1, 0, 0, 0,
		449, 450, 5, 1, 0, 0, 450, 465, 3, 94, 47, 0, 451, 454, 5, 1, 0, 0, 452,
		455, 3, 98, 49, 0, 453, 455, 3, 100, 50, 0, 454, 452, 1, 0, 0, 0, 454,
		453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 457, 3, 94, 47, 0, 457, 465,
		1, 0, 0, 0, 458, 459, 5, 1, 0, 0, 459, 462, 3, 94, 47, 0, 460, 463, 3,
		98, 49, 0, 461, 463, 3, 100, 50, 0, 462, 460, 1, 0, 0, 0, 462, 461, 1,
		0, 0, 0, 463, 465, 1, 0, 0, 0, 464, 449, 1, 0, 0, 0, 464, 451, 1, 0, 0,
		0, 464, 458, 1, 0, 0, 0, 465, 97, 1, 0, 0, 0, 466, 467, 5, 38, 0, 0, 467,
		468, 3, 112, 56, 0, 468, 99, 1, 0, 0, 0, 469, 470, 5, 39, 0, 0, 470, 471,
		3, 112, 56, 0, 471, 101, 1, 0, 0, 0, 472, 475, 3, 104, 52, 0, 473, 475,
		3, 106, 53, 0, 474, 472, 1, 0, 0, 0, 474, 473, 1, 0, 0, 0, 475, 478, 1,
		0, 0, 0, 476, 479, 3, 108, 54, 0, 477, 479, 3, 110, 55, 0, 478, 476, 1,
		0, 0, 0, 478, 477, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 103, 1, 0, 0,
		0, 480, 481, 5, 40, 0, 0, 481, 482, 3, 112, 56, 0, 482, 105, 1, 0, 0, 0,
		483, 484, 5, 41, 0, 0, 484, 485, 3, 112, 56, 0, 485, 107, 1, 0, 0, 0, 486,
		488, 5, 42, 0, 0, 487, 489, 3, 112, 56, 0, 488, 487, 1, 0, 0, 0, 488, 489,
		1, 0, 0, 0, 489, 109, 1, 0, 0, 0, 490, 492, 5, 43, 0, 0, 491, 493, 3, 112,
		56, 0, 492, 491, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 111, 1, 0, 0, 0,
		494, 503, 5, 48, 0, 0, 495, 500, 3, 56, 28, 0, 496, 497, 5, 50, 0, 0, 497,
		499, 3, 56, 28, 0, 498, 496, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 498,
		1, 0, 0, 0, 500, 501, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0,
		0, 0, 503, 495, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0,
		505, 506, 5, 49, 0, 0, 506, 113, 1, 0, 0, 0, 507, 508, 7, 7, 0, 0, 508,
		115, 1, 0, 0, 0, 509, 510, 7, 8, 0, 0, 510, 117, 1, 0, 0, 0, 49, 123, 132,
		146, 156, 165, 179, 188, 200, 211, 225, 227, 234, 246, 258, 293, 295, 300,
		316, 320, 325, 328, 334, 339, 350, 359, 363, 367, 371, 374, 378, 382, 386,
		395, 412, 416, 425, 428, 434, 442, 445, 454, 462, 464, 474, 478, 488, 492,
		500, 503,
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
	PromQLExParserRFC_3339_TIMESTAMP   = 3
	PromQLExParserCOLON                = 4
	PromQLExParserMETRIC_NAME          = 5
	PromQLExParserIF                   = 6
	PromQLExParserTRUE                 = 7
	PromQLExParserFALSE                = 8
	PromQLExParserSEMICOLON            = 9
	PromQLExParserMETRIC_KEYWORD       = 10
	PromQLExParserLABEL_KEYWORD        = 11
	PromQLExParserDEF                  = 12
	PromQLExParserCALL_SIGN            = 13
	PromQLExParserLEFT_BRACKET         = 14
	PromQLExParserRIGHT_BRACKET        = 15
	PromQLExParserNUMBER               = 16
	PromQLExParserSTRING               = 17
	PromQLExParserADD                  = 18
	PromQLExParserSUB                  = 19
	PromQLExParserMULT                 = 20
	PromQLExParserDIV                  = 21
	PromQLExParserMOD                  = 22
	PromQLExParserPOW                  = 23
	PromQLExParserAND                  = 24
	PromQLExParserOR                   = 25
	PromQLExParserUNLESS               = 26
	PromQLExParserSTART                = 27
	PromQLExParserEND                  = 28
	PromQLExParserEQ                   = 29
	PromQLExParserDEQ                  = 30
	PromQLExParserNE                   = 31
	PromQLExParserGT                   = 32
	PromQLExParserLT                   = 33
	PromQLExParserGE                   = 34
	PromQLExParserLE                   = 35
	PromQLExParserRE                   = 36
	PromQLExParserNRE                  = 37
	PromQLExParserBY                   = 38
	PromQLExParserWITHOUT              = 39
	PromQLExParserON                   = 40
	PromQLExParserIGNORING             = 41
	PromQLExParserGROUP_LEFT           = 42
	PromQLExParserGROUP_RIGHT          = 43
	PromQLExParserOFFSET               = 44
	PromQLExParserBOOL                 = 45
	PromQLExParserLEFT_BRACE           = 46
	PromQLExParserRIGHT_BRACE          = 47
	PromQLExParserLEFT_PAREN           = 48
	PromQLExParserRIGHT_PAREN          = 49
	PromQLExParserCOMMA                = 50
	PromQLExParserAT                   = 51
	PromQLExParserDURATION             = 52
	PromQLExParserLABEL_NAME           = 53
	PromQLExParserWS                   = 54
	PromQLExParserSL_COMMENT           = 55
	PromQLExParserID                   = 56
	PromQLExParserRAW_STRING           = 57
	PromQLExParserBACKTICK_OPEN        = 58
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
	PromQLExParserRULE_const_num_expression   = 16
	PromQLExParserRULE_num_literal            = 17
	PromQLExParserRULE_duration               = 18
	PromQLExParserRULE_time_range             = 19
	PromQLExParserRULE_subquery_range         = 20
	PromQLExParserRULE_vectorOperation        = 21
	PromQLExParserRULE_subqueryOp             = 22
	PromQLExParserRULE_offsetOp               = 23
	PromQLExParserRULE_matrixSelector         = 24
	PromQLExParserRULE_offset                 = 25
	PromQLExParserRULE_literal                = 26
	PromQLExParserRULE_instantSelector        = 27
	PromQLExParserRULE_labelName              = 28
	PromQLExParserRULE_metric_name            = 29
	PromQLExParserRULE_at_modifier_timestamp  = 30
	PromQLExParserRULE_expression             = 31
	PromQLExParserRULE_unaryOp                = 32
	PromQLExParserRULE_powOp                  = 33
	PromQLExParserRULE_multOp                 = 34
	PromQLExParserRULE_addOp                  = 35
	PromQLExParserRULE_compareOp              = 36
	PromQLExParserRULE_andUnlessOp            = 37
	PromQLExParserRULE_orOp                   = 38
	PromQLExParserRULE_vectorMatchOp          = 39
	PromQLExParserRULE_vector                 = 40
	PromQLExParserRULE_parens                 = 41
	PromQLExParserRULE_labelMatcher           = 42
	PromQLExParserRULE_labelMatcherOperator   = 43
	PromQLExParserRULE_labelMatcherList       = 44
	PromQLExParserRULE_function_              = 45
	PromQLExParserRULE_parameter              = 46
	PromQLExParserRULE_parameterList          = 47
	PromQLExParserRULE_aggregation            = 48
	PromQLExParserRULE_by                     = 49
	PromQLExParserRULE_without                = 50
	PromQLExParserRULE_grouping               = 51
	PromQLExParserRULE_on_                    = 52
	PromQLExParserRULE_ignoring               = 53
	PromQLExParserRULE_groupLeft              = 54
	PromQLExParserRULE_groupRight             = 55
	PromQLExParserRULE_labelNameList          = 56
	PromQLExParserRULE_keyword                = 57
	PromQLExParserRULE_string                 = 58
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
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

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

func (s *PromqlxContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSEMICOLON)
}

func (s *PromqlxContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSEMICOLON, i)
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
		p.SetState(118)
		p.Statement()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLExParserSEMICOLON {
		{
			p.SetState(119)
			p.Match(PromQLExParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Statement()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
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

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type State_MacroDefContext struct {
	StatementContext
}

func NewState_MacroDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *State_MacroDefContext {
	var p = new(State_MacroDefContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *State_MacroDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_MacroDefContext) Macro_def() IMacro_defContext {
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

func (s *State_MacroDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterState_MacroDef(s)
	}
}

func (s *State_MacroDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitState_MacroDef(s)
	}
}

type State_VectorOperationContext struct {
	StatementContext
}

func NewState_VectorOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *State_VectorOperationContext {
	var p = new(State_VectorOperationContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *State_VectorOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_VectorOperationContext) VectorOperation() IVectorOperationContext {
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

func (s *State_VectorOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterState_VectorOperation(s)
	}
}

func (s *State_VectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitState_VectorOperation(s)
	}
}

type State_IfStatementContext struct {
	StatementContext
}

func NewState_IfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *State_IfStatementContext {
	var p = new(State_IfStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *State_IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_IfStatementContext) If_statement() IIf_statementContext {
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

func (s *State_IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterState_IfStatement(s)
	}
}

func (s *State_IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitState_IfStatement(s)
	}
}

type State_AliasDefContext struct {
	StatementContext
}

func NewState_AliasDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *State_AliasDefContext {
	var p = new(State_AliasDefContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *State_AliasDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_AliasDefContext) Alias_def() IAlias_defContext {
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

func (s *State_AliasDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterState_AliasDef(s)
	}
}

func (s *State_AliasDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitState_AliasDef(s)
	}
}

func (p *PromQLExParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromQLExParserRULE_statement)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewState_AliasDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Alias_def()
		}

	case 2:
		localctx = NewState_MacroDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Macro_def()
		}

	case 3:
		localctx = NewState_IfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.If_statement()
		}

	case 4:
		localctx = NewState_VectorOperationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
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
		p.SetState(134)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(PromQLExParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
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
		p.SetState(139)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
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
		p.SetState(142)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(PromQLExParserLEFT_PAREN)
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

	if _la == PromQLExParserID {
		{
			p.SetState(145)
			p.Args_decl()
		}

	}
	{
		p.SetState(148)
		p.Match(PromQLExParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
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
		p.SetState(151)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(PromQLExParserLEFT_PAREN)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&221028225463033902) != 0 {
		{
			p.SetState(155)
			p.Arg_list()
		}

	}
	{
		p.SetState(158)
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
		p.SetState(160)
		p.Arg_name()
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLExParserCOMMA {
		{
			p.SetState(161)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Arg_name()
		}

		p.SetState(167)
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
		p.SetState(168)
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
		p.SetState(170)
		p.Match(PromQLExParserLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Promqlx()
	}
	{
		p.SetState(172)
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
		p.SetState(174)
		p.vectorOperation(0)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromQLExParserCOMMA {
		{
			p.SetState(175)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.vectorOperation(0)
		}

		p.SetState(179)
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
		p.SetState(181)
		p.Match(PromQLExParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Condition()
	}
	{
		p.SetState(183)
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
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserMETRIC_NAME, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserSTRING, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_BRACE, PromQLExParserLEFT_PAREN, PromQLExParserDURATION, PromQLExParserID, PromQLExParserRAW_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.CompareVectorOperation()
		}

	case PromQLExParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.TrueConst()
		}

	case PromQLExParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
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
		p.SetState(190)
		p.vectorOperation(0)
	}
	{
		p.SetState(191)
		p.CompareOp()
	}
	{
		p.SetState(192)
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
		p.SetState(194)
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
		p.SetState(196)
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

func (s *Time_instant_literalContext) CopyAll(ctx *Time_instant_literalContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Time_instant_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Time_instant_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TimeInstLit_UnixTimestampContext struct {
	Time_instant_literalContext
}

func NewTimeInstLit_UnixTimestampContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeInstLit_UnixTimestampContext {
	var p = new(TimeInstLit_UnixTimestampContext)

	InitEmptyTime_instant_literalContext(&p.Time_instant_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Time_instant_literalContext))

	return p
}

func (s *TimeInstLit_UnixTimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeInstLit_UnixTimestampContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNUMBER, 0)
}

func (s *TimeInstLit_UnixTimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterTimeInstLit_UnixTimestamp(s)
	}
}

func (s *TimeInstLit_UnixTimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitTimeInstLit_UnixTimestamp(s)
	}
}

type TimeInstLit_IsoDateTimeContext struct {
	Time_instant_literalContext
}

func NewTimeInstLit_IsoDateTimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeInstLit_IsoDateTimeContext {
	var p = new(TimeInstLit_IsoDateTimeContext)

	InitEmptyTime_instant_literalContext(&p.Time_instant_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Time_instant_literalContext))

	return p
}

func (s *TimeInstLit_IsoDateTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeInstLit_IsoDateTimeContext) RFC_3339_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRFC_3339_TIMESTAMP, 0)
}

func (s *TimeInstLit_IsoDateTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterTimeInstLit_IsoDateTime(s)
	}
}

func (s *TimeInstLit_IsoDateTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitTimeInstLit_IsoDateTime(s)
	}
}

func (p *PromQLExParser) Time_instant_literal() (localctx ITime_instant_literalContext) {
	localctx = NewTime_instant_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromQLExParserRULE_time_instant_literal)
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP:
		localctx = NewTimeInstLit_IsoDateTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Match(PromQLExParserRFC_3339_TIMESTAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExParserNUMBER:
		localctx = NewTimeInstLit_UnixTimestampContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(PromQLExParserNUMBER)
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

// IConst_num_expressionContext is an interface to support dynamic dispatch.
type IConst_num_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *Const_num_expressionContext) CopyAll(ctx *Const_num_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Const_num_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_num_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConsNumExpr_NumLiteralContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_NumLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_NumLiteralContext {
	var p = new(ConsNumExpr_NumLiteralContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_NumLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_NumLiteralContext) Num_literal() INum_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INum_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INum_literalContext)
}

func (s *ConsNumExpr_NumLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_NumLiteral(s)
	}
}

func (s *ConsNumExpr_NumLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_NumLiteral(s)
	}
}

type ConsNumExpr_MulOpContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_MulOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_MulOpContext {
	var p = new(ConsNumExpr_MulOpContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_MulOpContext) AllConst_num_expression() []IConst_num_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			len++
		}
	}

	tst := make([]IConst_num_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_num_expressionContext); ok {
			tst[i] = t.(IConst_num_expressionContext)
			i++
		}
	}

	return tst
}

func (s *ConsNumExpr_MulOpContext) Const_num_expression(i int) IConst_num_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
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

	return t.(IConst_num_expressionContext)
}

func (s *ConsNumExpr_MulOpContext) MultOp() IMultOpContext {
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

func (s *ConsNumExpr_MulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_MulOp(s)
	}
}

func (s *ConsNumExpr_MulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_MulOp(s)
	}
}

type ConsNumExpr_PowerOpContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_PowerOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_PowerOpContext {
	var p = new(ConsNumExpr_PowerOpContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_PowerOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_PowerOpContext) AllConst_num_expression() []IConst_num_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			len++
		}
	}

	tst := make([]IConst_num_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_num_expressionContext); ok {
			tst[i] = t.(IConst_num_expressionContext)
			i++
		}
	}

	return tst
}

func (s *ConsNumExpr_PowerOpContext) Const_num_expression(i int) IConst_num_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
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

	return t.(IConst_num_expressionContext)
}

func (s *ConsNumExpr_PowerOpContext) PowOp() IPowOpContext {
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

func (s *ConsNumExpr_PowerOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_PowerOp(s)
	}
}

func (s *ConsNumExpr_PowerOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_PowerOp(s)
	}
}

type ConsNumExpr_UnaryOpContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_UnaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_UnaryOpContext {
	var p = new(ConsNumExpr_UnaryOpContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_UnaryOpContext) UnaryOp() IUnaryOpContext {
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

func (s *ConsNumExpr_UnaryOpContext) Const_num_expression() IConst_num_expressionContext {
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

func (s *ConsNumExpr_UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_UnaryOp(s)
	}
}

func (s *ConsNumExpr_UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_UnaryOp(s)
	}
}

type ConsNumExpr_ParenOpContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_ParenOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_ParenOpContext {
	var p = new(ConsNumExpr_ParenOpContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_ParenOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_ParenOpContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *ConsNumExpr_ParenOpContext) Const_num_expression() IConst_num_expressionContext {
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

func (s *ConsNumExpr_ParenOpContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *ConsNumExpr_ParenOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_ParenOp(s)
	}
}

func (s *ConsNumExpr_ParenOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_ParenOp(s)
	}
}

type ConsNumExpr_AddOpContext struct {
	Const_num_expressionContext
}

func NewConsNumExpr_AddOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConsNumExpr_AddOpContext {
	var p = new(ConsNumExpr_AddOpContext)

	InitEmptyConst_num_expressionContext(&p.Const_num_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Const_num_expressionContext))

	return p
}

func (s *ConsNumExpr_AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsNumExpr_AddOpContext) AllConst_num_expression() []IConst_num_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
			len++
		}
	}

	tst := make([]IConst_num_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_num_expressionContext); ok {
			tst[i] = t.(IConst_num_expressionContext)
			i++
		}
	}

	return tst
}

func (s *ConsNumExpr_AddOpContext) Const_num_expression(i int) IConst_num_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_num_expressionContext); ok {
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

	return t.(IConst_num_expressionContext)
}

func (s *ConsNumExpr_AddOpContext) AddOp() IAddOpContext {
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

func (s *ConsNumExpr_AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterConsNumExpr_AddOp(s)
	}
}

func (s *ConsNumExpr_AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitConsNumExpr_AddOp(s)
	}
}

func (p *PromQLExParser) Const_num_expression() (localctx IConst_num_expressionContext) {
	return p.const_num_expression(0)
}

func (p *PromQLExParser) const_num_expression(_p int) (localctx IConst_num_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConst_num_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConst_num_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, PromQLExParserRULE_const_num_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserADD, PromQLExParserSUB:
		localctx = NewConsNumExpr_UnaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(203)
			p.UnaryOp()
		}
		{
			p.SetState(204)
			p.const_num_expression(5)
		}

	case PromQLExParserLEFT_PAREN:
		localctx = NewConsNumExpr_ParenOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.const_num_expression(0)
		}
		{
			p.SetState(208)
			p.Match(PromQLExParserRIGHT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserDURATION:
		localctx = NewConsNumExpr_NumLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Num_literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConsNumExpr_PowerOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(214)
					p.PowOp()
				}
				{
					p.SetState(215)
					p.const_num_expression(6)
				}

			case 2:
				localctx = NewConsNumExpr_MulOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(218)
					p.MultOp()
				}
				{
					p.SetState(219)
					p.const_num_expression(5)
				}

			case 3:
				localctx = NewConsNumExpr_AddOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(222)
					p.AddOp()
				}
				{
					p.SetState(223)
					p.const_num_expression(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
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

// INum_literalContext is an interface to support dynamic dispatch.
type INum_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *Num_literalContext) CopyAll(ctx *Num_literalContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Num_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Num_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumLit_DurationContext struct {
	Num_literalContext
}

func NewNumLit_DurationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumLit_DurationContext {
	var p = new(NumLit_DurationContext)

	InitEmptyNum_literalContext(&p.Num_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Num_literalContext))

	return p
}

func (s *NumLit_DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumLit_DurationContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLExParserDURATION, 0)
}

func (s *NumLit_DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterNumLit_Duration(s)
	}
}

func (s *NumLit_DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitNumLit_Duration(s)
	}
}

type NumLit_AliasCallContext struct {
	Num_literalContext
}

func NewNumLit_AliasCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumLit_AliasCallContext {
	var p = new(NumLit_AliasCallContext)

	InitEmptyNum_literalContext(&p.Num_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Num_literalContext))

	return p
}

func (s *NumLit_AliasCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumLit_AliasCallContext) Alias_call() IAlias_callContext {
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

func (s *NumLit_AliasCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterNumLit_AliasCall(s)
	}
}

func (s *NumLit_AliasCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitNumLit_AliasCall(s)
	}
}

type NumLit_NumberContext struct {
	Num_literalContext
}

func NewNumLit_NumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumLit_NumberContext {
	var p = new(NumLit_NumberContext)

	InitEmptyNum_literalContext(&p.Num_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Num_literalContext))

	return p
}

func (s *NumLit_NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumLit_NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLExParserNUMBER, 0)
}

func (s *NumLit_NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterNumLit_Number(s)
	}
}

func (s *NumLit_NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitNumLit_Number(s)
	}
}

type NumLit_TimeInstantLitContext struct {
	Num_literalContext
}

func NewNumLit_TimeInstantLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumLit_TimeInstantLitContext {
	var p = new(NumLit_TimeInstantLitContext)

	InitEmptyNum_literalContext(&p.Num_literalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Num_literalContext))

	return p
}

func (s *NumLit_TimeInstantLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumLit_TimeInstantLitContext) Time_instant_literal() ITime_instant_literalContext {
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

func (s *NumLit_TimeInstantLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterNumLit_TimeInstantLit(s)
	}
}

func (s *NumLit_TimeInstantLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitNumLit_TimeInstantLit(s)
	}
}

func (p *PromQLExParser) Num_literal() (localctx INum_literalContext) {
	localctx = NewNum_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromQLExParserRULE_num_literal)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumLit_NumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Match(PromQLExParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNumLit_DurationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Match(PromQLExParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNumLit_TimeInstantLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(232)
			p.Time_instant_literal()
		}

	case 4:
		localctx = NewNumLit_AliasCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.Alias_call()
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
	p.EnterRule(localctx, 36, PromQLExParserRULE_duration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.const_num_expression(0)
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
	p.EnterRule(localctx, 38, PromQLExParserRULE_time_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Duration()
	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 40, PromQLExParserRULE_subquery_range)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Duration()
	}
	{
		p.SetState(244)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4785074604941320) != 0 {
		{
			p.SetState(245)
			p.Duration()
		}

	}
	{
		p.SetState(248)
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

func (s *VectorOperationContext) CopyAll(ctx *VectorOperationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VectorOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VecOp_MacroContext struct {
	VectorOperationContext
}

func NewVecOp_MacroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_MacroContext {
	var p = new(VecOp_MacroContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_MacroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_MacroContext) Macro_call() IMacro_callContext {
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

func (s *VecOp_MacroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_Macro(s)
	}
}

func (s *VecOp_MacroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_Macro(s)
	}
}

type VecOp_AddOpContext struct {
	VectorOperationContext
}

func NewVecOp_AddOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_AddOpContext {
	var p = new(VecOp_AddOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_AddOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_AddOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_AddOpContext) AddOp() IAddOpContext {
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

func (s *VecOp_AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_AddOp(s)
	}
}

func (s *VecOp_AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_AddOp(s)
	}
}

type VecOp_SubqueryOpContext struct {
	VectorOperationContext
}

func NewVecOp_SubqueryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_SubqueryOpContext {
	var p = new(VecOp_SubqueryOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_SubqueryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_SubqueryOpContext) VectorOperation() IVectorOperationContext {
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

func (s *VecOp_SubqueryOpContext) SubqueryOp() ISubqueryOpContext {
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

func (s *VecOp_SubqueryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_SubqueryOp(s)
	}
}

func (s *VecOp_SubqueryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_SubqueryOp(s)
	}
}

type VecOp_OrOpContext struct {
	VectorOperationContext
}

func NewVecOp_OrOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_OrOpContext {
	var p = new(VecOp_OrOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_OrOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_OrOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_OrOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_OrOpContext) OrOp() IOrOpContext {
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

func (s *VecOp_OrOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_OrOp(s)
	}
}

func (s *VecOp_OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_OrOp(s)
	}
}

type VecOp_VecContext struct {
	VectorOperationContext
}

func NewVecOp_VecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_VecContext {
	var p = new(VecOp_VecContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_VecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_VecContext) Vector() IVectorContext {
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

func (s *VecOp_VecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_Vec(s)
	}
}

func (s *VecOp_VecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_Vec(s)
	}
}

type VecOp_AtContext struct {
	VectorOperationContext
}

func NewVecOp_AtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_AtContext {
	var p = new(VecOp_AtContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_AtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_AtContext) VectorOperation() IVectorOperationContext {
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

func (s *VecOp_AtContext) AT() antlr.TerminalNode {
	return s.GetToken(PromQLExParserAT, 0)
}

func (s *VecOp_AtContext) At_modifier_timestamp() IAt_modifier_timestampContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAt_modifier_timestampContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAt_modifier_timestampContext)
}

func (s *VecOp_AtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_At(s)
	}
}

func (s *VecOp_AtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_At(s)
	}
}

type VecOp_UnaryOpContext struct {
	VectorOperationContext
}

func NewVecOp_UnaryOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_UnaryOpContext {
	var p = new(VecOp_UnaryOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_UnaryOpContext) UnaryOp() IUnaryOpContext {
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

func (s *VecOp_UnaryOpContext) VectorOperation() IVectorOperationContext {
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

func (s *VecOp_UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_UnaryOp(s)
	}
}

func (s *VecOp_UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_UnaryOp(s)
	}
}

type VecOp_VecMatchOpContext struct {
	VectorOperationContext
}

func NewVecOp_VecMatchOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_VecMatchOpContext {
	var p = new(VecOp_VecMatchOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_VecMatchOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_VecMatchOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_VecMatchOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_VecMatchOpContext) VectorMatchOp() IVectorMatchOpContext {
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

func (s *VecOp_VecMatchOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_VecMatchOp(s)
	}
}

func (s *VecOp_VecMatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_VecMatchOp(s)
	}
}

type VecOp_MulOpContext struct {
	VectorOperationContext
}

func NewVecOp_MulOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_MulOpContext {
	var p = new(VecOp_MulOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_MulOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_MulOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_MulOpContext) MultOp() IMultOpContext {
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

func (s *VecOp_MulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_MulOp(s)
	}
}

func (s *VecOp_MulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_MulOp(s)
	}
}

type VecOp_ConstNumExprContext struct {
	VectorOperationContext
}

func NewVecOp_ConstNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_ConstNumExprContext {
	var p = new(VecOp_ConstNumExprContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_ConstNumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_ConstNumExprContext) Const_num_expression() IConst_num_expressionContext {
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

func (s *VecOp_ConstNumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_ConstNumExpr(s)
	}
}

func (s *VecOp_ConstNumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_ConstNumExpr(s)
	}
}

type VecOp_AliasContext struct {
	VectorOperationContext
}

func NewVecOp_AliasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_AliasContext {
	var p = new(VecOp_AliasContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_AliasContext) Alias_call() IAlias_callContext {
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

func (s *VecOp_AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_Alias(s)
	}
}

func (s *VecOp_AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_Alias(s)
	}
}

type VecOp_PowOpContext struct {
	VectorOperationContext
}

func NewVecOp_PowOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_PowOpContext {
	var p = new(VecOp_PowOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_PowOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_PowOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_PowOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_PowOpContext) PowOp() IPowOpContext {
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

func (s *VecOp_PowOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_PowOp(s)
	}
}

func (s *VecOp_PowOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_PowOp(s)
	}
}

type VecOp_CompareOpContext struct {
	VectorOperationContext
}

func NewVecOp_CompareOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_CompareOpContext {
	var p = new(VecOp_CompareOpContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_CompareOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_CompareOpContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_CompareOpContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_CompareOpContext) CompareOp() ICompareOpContext {
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

func (s *VecOp_CompareOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_CompareOp(s)
	}
}

func (s *VecOp_CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_CompareOp(s)
	}
}

type VecOp_AndUnlessContext struct {
	VectorOperationContext
}

func NewVecOp_AndUnlessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VecOp_AndUnlessContext {
	var p = new(VecOp_AndUnlessContext)

	InitEmptyVectorOperationContext(&p.VectorOperationContext)
	p.parser = parser
	p.CopyAll(ctx.(*VectorOperationContext))

	return p
}

func (s *VecOp_AndUnlessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecOp_AndUnlessContext) AllVectorOperation() []IVectorOperationContext {
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

func (s *VecOp_AndUnlessContext) VectorOperation(i int) IVectorOperationContext {
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

func (s *VecOp_AndUnlessContext) AndUnlessOp() IAndUnlessOpContext {
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

func (s *VecOp_AndUnlessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterVecOp_AndUnless(s)
	}
}

func (s *VecOp_AndUnlessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitVecOp_AndUnless(s)
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
	_startState := 42
	p.EnterRecursionRule(localctx, 42, PromQLExParserRULE_vectorOperation, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVecOp_ConstNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(251)
			p.const_num_expression(0)
		}

	case 2:
		localctx = NewVecOp_UnaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(252)
			p.UnaryOp()
		}
		{
			p.SetState(253)
			p.vectorOperation(11)
		}

	case 3:
		localctx = NewVecOp_VecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(255)
			p.Vector()
		}

	case 4:
		localctx = NewVecOp_MacroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Macro_call()
		}

	case 5:
		localctx = NewVecOp_AliasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(257)
			p.Alias_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVecOp_PowOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(261)
					p.PowOp()
				}
				{
					p.SetState(262)
					p.vectorOperation(13)
				}

			case 2:
				localctx = NewVecOp_MulOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(264)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(265)
					p.MultOp()
				}
				{
					p.SetState(266)
					p.vectorOperation(11)
				}

			case 3:
				localctx = NewVecOp_AddOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(268)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(269)
					p.AddOp()
				}
				{
					p.SetState(270)
					p.vectorOperation(10)
				}

			case 4:
				localctx = NewVecOp_CompareOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(273)
					p.CompareOp()
				}
				{
					p.SetState(274)
					p.vectorOperation(9)
				}

			case 5:
				localctx = NewVecOp_AndUnlessContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(276)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(277)
					p.AndUnlessOp()
				}
				{
					p.SetState(278)
					p.vectorOperation(8)
				}

			case 6:
				localctx = NewVecOp_OrOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(281)
					p.OrOp()
				}
				{
					p.SetState(282)
					p.vectorOperation(7)
				}

			case 7:
				localctx = NewVecOp_VecMatchOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(285)
					p.VectorMatchOp()
				}
				{
					p.SetState(286)
					p.vectorOperation(6)
				}

			case 8:
				localctx = NewVecOp_SubqueryOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(289)
					p.SubqueryOp()
				}

			case 9:
				localctx = NewVecOp_AtContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(291)
					p.Match(PromQLExParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)
					p.At_modifier_timestamp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 44, PromQLExParserRULE_subqueryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Subquery_range()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(299)
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
	p.EnterRule(localctx, 46, PromQLExParserRULE_offsetOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(PromQLExParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
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
	p.EnterRule(localctx, 48, PromQLExParserRULE_matrixSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.InstantSelector()
	}
	{
		p.SetState(306)
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
	p.EnterRule(localctx, 50, PromQLExParserRULE_offset)
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.InstantSelector()
		}
		{
			p.SetState(309)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Duration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.MatrixSelector()
		}
		{
			p.SetState(313)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
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

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Lit_ConstNumExprContext struct {
	LiteralContext
}

func NewLit_ConstNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Lit_ConstNumExprContext {
	var p = new(Lit_ConstNumExprContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *Lit_ConstNumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lit_ConstNumExprContext) Const_num_expression() IConst_num_expressionContext {
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

func (s *Lit_ConstNumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLit_ConstNumExpr(s)
	}
}

func (s *Lit_ConstNumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLit_ConstNumExpr(s)
	}
}

type Lit_StringContext struct {
	LiteralContext
}

func NewLit_StringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Lit_StringContext {
	var p = new(Lit_StringContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *Lit_StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lit_StringContext) String_() IStringContext {
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

func (s *Lit_StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterLit_String(s)
	}
}

func (s *Lit_StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitLit_String(s)
	}
}

func (p *PromQLExParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromQLExParserRULE_literal)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_PAREN, PromQLExParserDURATION:
		localctx = NewLit_ConstNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.const_num_expression(0)
		}

	case PromQLExParserSTRING, PromQLExParserRAW_STRING:
		localctx = NewLit_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
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
	p.EnterRule(localctx, 54, PromQLExParserRULE_instantSelector)
	var _la int

	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Metric_name()
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(323)
				p.Match(PromQLExParserLEFT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81134887276380198) != 0 {
				{
					p.SetState(324)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(327)
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
			p.SetState(330)
			p.Match(PromQLExParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.LabelMatcherList()
		}
		{
			p.SetState(332)
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
	p.EnterRule(localctx, 56, PromQLExParserRULE_labelName)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserAND, PromQLExParserOR, PromQLExParserUNLESS, PromQLExParserBY, PromQLExParserWITHOUT, PromQLExParserON, PromQLExParserIGNORING, PromQLExParserGROUP_LEFT, PromQLExParserGROUP_RIGHT, PromQLExParserOFFSET, PromQLExParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Keyword()
		}

	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.Metric_name()
		}

	case PromQLExParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
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
	p.EnterRule(localctx, 58, PromQLExParserRULE_metric_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
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

// IAt_modifier_timestampContext is an interface to support dynamic dispatch.
type IAt_modifier_timestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAt_modifier_timestampContext differentiates from other interfaces.
	IsAt_modifier_timestampContext()
}

type At_modifier_timestampContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAt_modifier_timestampContext() *At_modifier_timestampContext {
	var p = new(At_modifier_timestampContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_at_modifier_timestamp
	return p
}

func InitEmptyAt_modifier_timestampContext(p *At_modifier_timestampContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLExParserRULE_at_modifier_timestamp
}

func (*At_modifier_timestampContext) IsAt_modifier_timestampContext() {}

func NewAt_modifier_timestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *At_modifier_timestampContext {
	var p = new(At_modifier_timestampContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLExParserRULE_at_modifier_timestamp

	return p
}

func (s *At_modifier_timestampContext) GetParser() antlr.Parser { return s.parser }

func (s *At_modifier_timestampContext) CopyAll(ctx *At_modifier_timestampContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *At_modifier_timestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *At_modifier_timestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AtModTime_StartContext struct {
	At_modifier_timestampContext
}

func NewAtModTime_StartContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtModTime_StartContext {
	var p = new(AtModTime_StartContext)

	InitEmptyAt_modifier_timestampContext(&p.At_modifier_timestampContext)
	p.parser = parser
	p.CopyAll(ctx.(*At_modifier_timestampContext))

	return p
}

func (s *AtModTime_StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtModTime_StartContext) START() antlr.TerminalNode {
	return s.GetToken(PromQLExParserSTART, 0)
}

func (s *AtModTime_StartContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *AtModTime_StartContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *AtModTime_StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAtModTime_Start(s)
	}
}

func (s *AtModTime_StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAtModTime_Start(s)
	}
}

type AtModTime_EndContext struct {
	At_modifier_timestampContext
}

func NewAtModTime_EndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtModTime_EndContext {
	var p = new(AtModTime_EndContext)

	InitEmptyAt_modifier_timestampContext(&p.At_modifier_timestampContext)
	p.parser = parser
	p.CopyAll(ctx.(*At_modifier_timestampContext))

	return p
}

func (s *AtModTime_EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtModTime_EndContext) END() antlr.TerminalNode {
	return s.GetToken(PromQLExParserEND, 0)
}

func (s *AtModTime_EndContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_PAREN, 0)
}

func (s *AtModTime_EndContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_PAREN, 0)
}

func (s *AtModTime_EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAtModTime_End(s)
	}
}

func (s *AtModTime_EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAtModTime_End(s)
	}
}

type AtModTime_ConstNumExprContext struct {
	At_modifier_timestampContext
}

func NewAtModTime_ConstNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtModTime_ConstNumExprContext {
	var p = new(AtModTime_ConstNumExprContext)

	InitEmptyAt_modifier_timestampContext(&p.At_modifier_timestampContext)
	p.parser = parser
	p.CopyAll(ctx.(*At_modifier_timestampContext))

	return p
}

func (s *AtModTime_ConstNumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtModTime_ConstNumExprContext) Const_num_expression() IConst_num_expressionContext {
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

func (s *AtModTime_ConstNumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.EnterAtModTime_ConstNumExpr(s)
	}
}

func (s *AtModTime_ConstNumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLExParserListener); ok {
		listenerT.ExitAtModTime_ConstNumExpr(s)
	}
}

func (p *PromQLExParser) At_modifier_timestamp() (localctx IAt_modifier_timestampContext) {
	localctx = NewAt_modifier_timestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromQLExParserRULE_at_modifier_timestamp)
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_PAREN, PromQLExParserDURATION:
		localctx = NewAtModTime_ConstNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.const_num_expression(0)
		}

	case PromQLExParserSTART:
		localctx = NewAtModTime_StartContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.Match(PromQLExParserSTART)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(PromQLExParserRIGHT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLExParserEND:
		localctx = NewAtModTime_EndContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(347)
			p.Match(PromQLExParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(PromQLExParserRIGHT_PAREN)
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
	p.EnterRule(localctx, 62, PromQLExParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.vectorOperation(0)
	}
	{
		p.SetState(353)
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
	p.EnterRule(localctx, 64, PromQLExParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
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
	p.EnterRule(localctx, 66, PromQLExParserRULE_powOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(PromQLExParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(358)
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
	p.EnterRule(localctx, 68, PromQLExParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(362)
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
	p.EnterRule(localctx, 70, PromQLExParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserADD || _la == PromQLExParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(366)
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
	p.EnterRule(localctx, 72, PromQLExParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67645734912) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserBOOL {
		{
			p.SetState(370)
			p.Match(PromQLExParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(373)
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
	p.EnterRule(localctx, 74, PromQLExParserRULE_andUnlessOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserAND || _la == PromQLExParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(377)
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
	p.EnterRule(localctx, 76, PromQLExParserRULE_orOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(PromQLExParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(381)
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
	p.EnterRule(localctx, 78, PromQLExParserRULE_vectorMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserUNLESS || _la == PromQLExParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(385)
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
	p.EnterRule(localctx, 80, PromQLExParserRULE_vector)
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(390)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(391)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(392)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(393)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(394)
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
	p.EnterRule(localctx, 82, PromQLExParserRULE_parens)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.vectorOperation(0)
	}
	{
		p.SetState(399)
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
	p.EnterRule(localctx, 84, PromQLExParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.LabelName()
	}
	{
		p.SetState(402)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(403)
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
	p.EnterRule(localctx, 86, PromQLExParserRULE_labelMatcherOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&208842784768) != 0) {
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
	p.EnterRule(localctx, 88, PromQLExParserRULE_labelMatcherList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.LabelMatcher()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(408)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(409)
				p.LabelMatcher()
			}

		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserCOMMA {
		{
			p.SetState(415)
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
	p.EnterRule(localctx, 90, PromQLExParserRULE_function_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(PromQLExParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&221028225463033902) != 0 {
		{
			p.SetState(420)
			p.Parameter()
		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(421)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(422)
				p.Parameter()
			}

			p.SetState(427)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(430)
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
	p.EnterRule(localctx, 92, PromQLExParserRULE_parameter)
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
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
	p.EnterRule(localctx, 94, PromQLExParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&221028225463033902) != 0 {
		{
			p.SetState(437)
			p.Parameter()
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(438)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(439)
				p.Parameter()
			}

			p.SetState(444)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(447)
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
	p.EnterRule(localctx, 96, PromQLExParserRULE_aggregation)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(449)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(451)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
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

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(452)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(453)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(456)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(458)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.ParameterList()
		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(460)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(461)
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
	p.EnterRule(localctx, 98, PromQLExParserRULE_by)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(PromQLExParserBY)
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
	p.EnterRule(localctx, 100, PromQLExParserRULE_without)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Match(PromQLExParserWITHOUT)
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
	p.EnterRule(localctx, 102, PromQLExParserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserON:
		{
			p.SetState(472)
			p.On_()
		}

	case PromQLExParserIGNORING:
		{
			p.SetState(473)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLExParserGROUP_LEFT:
		{
			p.SetState(476)
			p.GroupLeft()
		}

	case PromQLExParserGROUP_RIGHT:
		{
			p.SetState(477)
			p.GroupRight()
		}

	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserMETRIC_NAME, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserSTRING, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_BRACE, PromQLExParserLEFT_PAREN, PromQLExParserDURATION, PromQLExParserID, PromQLExParserRAW_STRING:

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
	p.EnterRule(localctx, 104, PromQLExParserRULE_on_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(480)
		p.Match(PromQLExParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)
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
	p.EnterRule(localctx, 106, PromQLExParserRULE_ignoring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(PromQLExParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(484)
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
	p.EnterRule(localctx, 108, PromQLExParserRULE_groupLeft)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		p.Match(PromQLExParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(487)
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
	p.EnterRule(localctx, 110, PromQLExParserRULE_groupRight)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Match(PromQLExParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(491)
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
	p.EnterRule(localctx, 112, PromQLExParserRULE_labelNameList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81134887276380198) != 0 {
		{
			p.SetState(495)
			p.LabelName()
		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(496)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(497)
				p.LabelName()
			}

			p.SetState(502)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(505)
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
	p.EnterRule(localctx, 114, PromQLExParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70093983711238) != 0) {
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
	p.EnterRule(localctx, 116, PromQLExParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
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
	case 16:
		var t *Const_num_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Const_num_expressionContext)
		}
		return p.Const_num_expression_Sempred(t, predIndex)

	case 21:
		var t *VectorOperationContext = nil
		if localctx != nil {
			t = localctx.(*VectorOperationContext)
		}
		return p.VectorOperation_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLExParser) Const_num_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PromQLExParser) VectorOperation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

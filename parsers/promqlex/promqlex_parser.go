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
		"ID", "ID_WS", "RAW_STRING", "BACKTICK_OPEN",
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
		4, 1, 59, 510, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		2, 58, 7, 58, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 126, 8, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 140, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 151, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 158, 8, 6, 10, 6,
		12, 6, 161, 9, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 168, 8, 8, 10, 8,
		12, 8, 171, 9, 8, 1, 9, 1, 9, 1, 9, 4, 9, 176, 8, 9, 11, 9, 12, 9, 177,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 187, 8, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3, 15,
		199, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 210, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 224, 8, 16, 10, 16, 12, 16, 227,
		9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 233, 8, 17, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 245, 8, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 257, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 292, 8, 21, 10, 21, 12, 21, 295, 9,
		21, 1, 22, 1, 22, 3, 22, 299, 8, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 315,
		8, 25, 1, 26, 1, 26, 3, 26, 319, 8, 26, 1, 27, 1, 27, 1, 27, 3, 27, 324,
		8, 27, 1, 27, 3, 27, 327, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 333,
		8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 338, 8, 28, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 349, 8, 30, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 3, 33, 358, 8, 33, 1, 34, 1, 34, 3,
		34, 362, 8, 34, 1, 35, 1, 35, 3, 35, 366, 8, 35, 1, 36, 1, 36, 3, 36, 370,
		8, 36, 1, 36, 3, 36, 373, 8, 36, 1, 37, 1, 37, 3, 37, 377, 8, 37, 1, 38,
		1, 38, 3, 38, 381, 8, 38, 1, 39, 1, 39, 3, 39, 385, 8, 39, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 394, 8, 40, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44,
		5, 44, 409, 8, 44, 10, 44, 12, 44, 412, 9, 44, 1, 44, 3, 44, 415, 8, 44,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 422, 8, 45, 10, 45, 12, 45, 425,
		9, 45, 3, 45, 427, 8, 45, 1, 45, 1, 45, 1, 46, 1, 46, 3, 46, 433, 8, 46,
		1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 439, 8, 47, 10, 47, 12, 47, 442, 9,
		47, 3, 47, 444, 8, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		3, 48, 453, 8, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 461,
		8, 48, 3, 48, 463, 8, 48, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		51, 1, 51, 3, 51, 473, 8, 51, 1, 51, 1, 51, 3, 51, 477, 8, 51, 1, 52, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 3, 54, 487, 8, 54, 1, 55,
		1, 55, 3, 55, 491, 8, 55, 1, 56, 1, 56, 1, 56, 1, 56, 5, 56, 497, 8, 56,
		10, 56, 12, 56, 500, 9, 56, 3, 56, 502, 8, 56, 1, 56, 1, 56, 1, 57, 1,
		57, 1, 58, 1, 58, 1, 58, 0, 2, 32, 42, 59, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 0, 9,
		2, 0, 5, 5, 56, 56, 1, 0, 18, 19, 1, 0, 20, 22, 1, 0, 30, 35, 2, 0, 24,
		24, 26, 26, 2, 0, 26, 26, 40, 40, 3, 0, 29, 29, 31, 31, 36, 37, 3, 0, 1,
		2, 24, 26, 38, 45, 2, 0, 17, 17, 58, 58, 525, 0, 118, 1, 0, 0, 0, 2, 125,
		1, 0, 0, 0, 4, 127, 1, 0, 0, 0, 6, 132, 1, 0, 0, 0, 8, 135, 1, 0, 0, 0,
		10, 146, 1, 0, 0, 0, 12, 154, 1, 0, 0, 0, 14, 162, 1, 0, 0, 0, 16, 164,
		1, 0, 0, 0, 18, 172, 1, 0, 0, 0, 20, 179, 1, 0, 0, 0, 22, 186, 1, 0, 0,
		0, 24, 188, 1, 0, 0, 0, 26, 192, 1, 0, 0, 0, 28, 194, 1, 0, 0, 0, 30, 198,
		1, 0, 0, 0, 32, 209, 1, 0, 0, 0, 34, 232, 1, 0, 0, 0, 36, 234, 1, 0, 0,
		0, 38, 236, 1, 0, 0, 0, 40, 240, 1, 0, 0, 0, 42, 256, 1, 0, 0, 0, 44, 296,
		1, 0, 0, 0, 46, 300, 1, 0, 0, 0, 48, 303, 1, 0, 0, 0, 50, 314, 1, 0, 0,
		0, 52, 318, 1, 0, 0, 0, 54, 332, 1, 0, 0, 0, 56, 337, 1, 0, 0, 0, 58, 339,
		1, 0, 0, 0, 60, 348, 1, 0, 0, 0, 62, 350, 1, 0, 0, 0, 64, 353, 1, 0, 0,
		0, 66, 355, 1, 0, 0, 0, 68, 359, 1, 0, 0, 0, 70, 363, 1, 0, 0, 0, 72, 367,
		1, 0, 0, 0, 74, 374, 1, 0, 0, 0, 76, 378, 1, 0, 0, 0, 78, 382, 1, 0, 0,
		0, 80, 393, 1, 0, 0, 0, 82, 395, 1, 0, 0, 0, 84, 399, 1, 0, 0, 0, 86, 403,
		1, 0, 0, 0, 88, 405, 1, 0, 0, 0, 90, 416, 1, 0, 0, 0, 92, 432, 1, 0, 0,
		0, 94, 434, 1, 0, 0, 0, 96, 462, 1, 0, 0, 0, 98, 464, 1, 0, 0, 0, 100,
		467, 1, 0, 0, 0, 102, 472, 1, 0, 0, 0, 104, 478, 1, 0, 0, 0, 106, 481,
		1, 0, 0, 0, 108, 484, 1, 0, 0, 0, 110, 488, 1, 0, 0, 0, 112, 492, 1, 0,
		0, 0, 114, 505, 1, 0, 0, 0, 116, 507, 1, 0, 0, 0, 118, 119, 3, 16, 8, 0,
		119, 120, 5, 0, 0, 1, 120, 1, 1, 0, 0, 0, 121, 126, 3, 4, 2, 0, 122, 126,
		3, 8, 4, 0, 123, 126, 3, 20, 10, 0, 124, 126, 3, 42, 21, 0, 125, 121, 1,
		0, 0, 0, 125, 122, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 124, 1, 0, 0,
		0, 126, 3, 1, 0, 0, 0, 127, 128, 5, 12, 0, 0, 128, 129, 5, 56, 0, 0, 129,
		130, 5, 29, 0, 0, 130, 131, 3, 42, 21, 0, 131, 5, 1, 0, 0, 0, 132, 133,
		5, 13, 0, 0, 133, 134, 5, 56, 0, 0, 134, 7, 1, 0, 0, 0, 135, 136, 5, 12,
		0, 0, 136, 137, 5, 56, 0, 0, 137, 139, 5, 48, 0, 0, 138, 140, 3, 12, 6,
		0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		142, 5, 49, 0, 0, 142, 143, 5, 46, 0, 0, 143, 144, 3, 16, 8, 0, 144, 145,
		5, 47, 0, 0, 145, 9, 1, 0, 0, 0, 146, 147, 5, 13, 0, 0, 147, 148, 5, 56,
		0, 0, 148, 150, 5, 48, 0, 0, 149, 151, 3, 18, 9, 0, 150, 149, 1, 0, 0,
		0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 5, 49, 0, 0, 153,
		11, 1, 0, 0, 0, 154, 159, 3, 14, 7, 0, 155, 156, 5, 50, 0, 0, 156, 158,
		3, 14, 7, 0, 157, 155, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0,
		0, 0, 159, 160, 1, 0, 0, 0, 160, 13, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0,
		162, 163, 5, 56, 0, 0, 163, 15, 1, 0, 0, 0, 164, 169, 3, 2, 1, 0, 165,
		166, 5, 9, 0, 0, 166, 168, 3, 2, 1, 0, 167, 165, 1, 0, 0, 0, 168, 171,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 17, 1, 0,
		0, 0, 171, 169, 1, 0, 0, 0, 172, 175, 3, 42, 21, 0, 173, 174, 5, 50, 0,
		0, 174, 176, 3, 42, 21, 0, 175, 173, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 19, 1, 0, 0, 0, 179, 180,
		5, 6, 0, 0, 180, 181, 3, 22, 11, 0, 181, 182, 3, 16, 8, 0, 182, 21, 1,
		0, 0, 0, 183, 187, 3, 24, 12, 0, 184, 187, 3, 26, 13, 0, 185, 187, 3, 28,
		14, 0, 186, 183, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0,
		187, 23, 1, 0, 0, 0, 188, 189, 3, 42, 21, 0, 189, 190, 3, 72, 36, 0, 190,
		191, 3, 42, 21, 0, 191, 25, 1, 0, 0, 0, 192, 193, 5, 7, 0, 0, 193, 27,
		1, 0, 0, 0, 194, 195, 5, 8, 0, 0, 195, 29, 1, 0, 0, 0, 196, 199, 5, 3,
		0, 0, 197, 199, 5, 16, 0, 0, 198, 196, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0,
		199, 31, 1, 0, 0, 0, 200, 201, 6, 16, -1, 0, 201, 202, 3, 64, 32, 0, 202,
		203, 3, 32, 16, 5, 203, 210, 1, 0, 0, 0, 204, 205, 5, 48, 0, 0, 205, 206,
		3, 32, 16, 0, 206, 207, 5, 49, 0, 0, 207, 210, 1, 0, 0, 0, 208, 210, 3,
		34, 17, 0, 209, 200, 1, 0, 0, 0, 209, 204, 1, 0, 0, 0, 209, 208, 1, 0,
		0, 0, 210, 225, 1, 0, 0, 0, 211, 212, 10, 6, 0, 0, 212, 213, 3, 66, 33,
		0, 213, 214, 3, 32, 16, 6, 214, 224, 1, 0, 0, 0, 215, 216, 10, 4, 0, 0,
		216, 217, 3, 68, 34, 0, 217, 218, 3, 32, 16, 5, 218, 224, 1, 0, 0, 0, 219,
		220, 10, 3, 0, 0, 220, 221, 3, 70, 35, 0, 221, 222, 3, 32, 16, 4, 222,
		224, 1, 0, 0, 0, 223, 211, 1, 0, 0, 0, 223, 215, 1, 0, 0, 0, 223, 219,
		1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0,
		0, 0, 226, 33, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 233, 5, 16, 0, 0,
		229, 233, 5, 52, 0, 0, 230, 233, 3, 30, 15, 0, 231, 233, 3, 6, 3, 0, 232,
		228, 1, 0, 0, 0, 232, 229, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232, 231,
		1, 0, 0, 0, 233, 35, 1, 0, 0, 0, 234, 235, 3, 32, 16, 0, 235, 37, 1, 0,
		0, 0, 236, 237, 5, 14, 0, 0, 237, 238, 3, 36, 18, 0, 238, 239, 5, 15, 0,
		0, 239, 39, 1, 0, 0, 0, 240, 241, 5, 14, 0, 0, 241, 242, 3, 36, 18, 0,
		242, 244, 5, 4, 0, 0, 243, 245, 3, 36, 18, 0, 244, 243, 1, 0, 0, 0, 244,
		245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 5, 15, 0, 0, 247, 41,
		1, 0, 0, 0, 248, 249, 6, 21, -1, 0, 249, 257, 3, 32, 16, 0, 250, 251, 3,
		64, 32, 0, 251, 252, 3, 42, 21, 11, 252, 257, 1, 0, 0, 0, 253, 257, 3,
		80, 40, 0, 254, 257, 3, 10, 5, 0, 255, 257, 3, 6, 3, 0, 256, 248, 1, 0,
		0, 0, 256, 250, 1, 0, 0, 0, 256, 253, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0,
		256, 255, 1, 0, 0, 0, 257, 293, 1, 0, 0, 0, 258, 259, 10, 13, 0, 0, 259,
		260, 3, 66, 33, 0, 260, 261, 3, 42, 21, 13, 261, 292, 1, 0, 0, 0, 262,
		263, 10, 10, 0, 0, 263, 264, 3, 68, 34, 0, 264, 265, 3, 42, 21, 11, 265,
		292, 1, 0, 0, 0, 266, 267, 10, 9, 0, 0, 267, 268, 3, 70, 35, 0, 268, 269,
		3, 42, 21, 10, 269, 292, 1, 0, 0, 0, 270, 271, 10, 8, 0, 0, 271, 272, 3,
		72, 36, 0, 272, 273, 3, 42, 21, 9, 273, 292, 1, 0, 0, 0, 274, 275, 10,
		7, 0, 0, 275, 276, 3, 74, 37, 0, 276, 277, 3, 42, 21, 8, 277, 292, 1, 0,
		0, 0, 278, 279, 10, 6, 0, 0, 279, 280, 3, 76, 38, 0, 280, 281, 3, 42, 21,
		7, 281, 292, 1, 0, 0, 0, 282, 283, 10, 5, 0, 0, 283, 284, 3, 78, 39, 0,
		284, 285, 3, 42, 21, 6, 285, 292, 1, 0, 0, 0, 286, 287, 10, 12, 0, 0, 287,
		292, 3, 44, 22, 0, 288, 289, 10, 4, 0, 0, 289, 290, 5, 51, 0, 0, 290, 292,
		3, 60, 30, 0, 291, 258, 1, 0, 0, 0, 291, 262, 1, 0, 0, 0, 291, 266, 1,
		0, 0, 0, 291, 270, 1, 0, 0, 0, 291, 274, 1, 0, 0, 0, 291, 278, 1, 0, 0,
		0, 291, 282, 1, 0, 0, 0, 291, 286, 1, 0, 0, 0, 291, 288, 1, 0, 0, 0, 292,
		295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 43, 1,
		0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 298, 3, 40, 20, 0, 297, 299, 3, 46,
		23, 0, 298, 297, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 45, 1, 0, 0, 0,
		300, 301, 5, 44, 0, 0, 301, 302, 3, 36, 18, 0, 302, 47, 1, 0, 0, 0, 303,
		304, 3, 54, 27, 0, 304, 305, 3, 38, 19, 0, 305, 49, 1, 0, 0, 0, 306, 307,
		3, 54, 27, 0, 307, 308, 5, 44, 0, 0, 308, 309, 3, 36, 18, 0, 309, 315,
		1, 0, 0, 0, 310, 311, 3, 48, 24, 0, 311, 312, 5, 44, 0, 0, 312, 313, 3,
		36, 18, 0, 313, 315, 1, 0, 0, 0, 314, 306, 1, 0, 0, 0, 314, 310, 1, 0,
		0, 0, 315, 51, 1, 0, 0, 0, 316, 319, 3, 32, 16, 0, 317, 319, 3, 116, 58,
		0, 318, 316, 1, 0, 0, 0, 318, 317, 1, 0, 0, 0, 319, 53, 1, 0, 0, 0, 320,
		326, 3, 58, 29, 0, 321, 323, 5, 46, 0, 0, 322, 324, 3, 88, 44, 0, 323,
		322, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 327,
		5, 47, 0, 0, 326, 321, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 333, 1, 0,
		0, 0, 328, 329, 5, 46, 0, 0, 329, 330, 3, 88, 44, 0, 330, 331, 5, 47, 0,
		0, 331, 333, 1, 0, 0, 0, 332, 320, 1, 0, 0, 0, 332, 328, 1, 0, 0, 0, 333,
		55, 1, 0, 0, 0, 334, 338, 3, 114, 57, 0, 335, 338, 3, 58, 29, 0, 336, 338,
		5, 53, 0, 0, 337, 334, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 337, 336, 1, 0,
		0, 0, 338, 57, 1, 0, 0, 0, 339, 340, 7, 0, 0, 0, 340, 59, 1, 0, 0, 0, 341,
		349, 3, 32, 16, 0, 342, 343, 5, 27, 0, 0, 343, 344, 5, 48, 0, 0, 344, 349,
		5, 49, 0, 0, 345, 346, 5, 28, 0, 0, 346, 347, 5, 48, 0, 0, 347, 349, 5,
		49, 0, 0, 348, 341, 1, 0, 0, 0, 348, 342, 1, 0, 0, 0, 348, 345, 1, 0, 0,
		0, 349, 61, 1, 0, 0, 0, 350, 351, 3, 42, 21, 0, 351, 352, 5, 0, 0, 1, 352,
		63, 1, 0, 0, 0, 353, 354, 7, 1, 0, 0, 354, 65, 1, 0, 0, 0, 355, 357, 5,
		23, 0, 0, 356, 358, 3, 102, 51, 0, 357, 356, 1, 0, 0, 0, 357, 358, 1, 0,
		0, 0, 358, 67, 1, 0, 0, 0, 359, 361, 7, 2, 0, 0, 360, 362, 3, 102, 51,
		0, 361, 360, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 69, 1, 0, 0, 0, 363,
		365, 7, 1, 0, 0, 364, 366, 3, 102, 51, 0, 365, 364, 1, 0, 0, 0, 365, 366,
		1, 0, 0, 0, 366, 71, 1, 0, 0, 0, 367, 369, 7, 3, 0, 0, 368, 370, 5, 45,
		0, 0, 369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 372, 1, 0, 0, 0,
		371, 373, 3, 102, 51, 0, 372, 371, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373,
		73, 1, 0, 0, 0, 374, 376, 7, 4, 0, 0, 375, 377, 3, 102, 51, 0, 376, 375,
		1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 75, 1, 0, 0, 0, 378, 380, 5, 25,
		0, 0, 379, 381, 3, 102, 51, 0, 380, 379, 1, 0, 0, 0, 380, 381, 1, 0, 0,
		0, 381, 77, 1, 0, 0, 0, 382, 384, 7, 5, 0, 0, 383, 385, 3, 102, 51, 0,
		384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 79, 1, 0, 0, 0, 386, 394,
		3, 90, 45, 0, 387, 394, 3, 96, 48, 0, 388, 394, 3, 54, 27, 0, 389, 394,
		3, 48, 24, 0, 390, 394, 3, 50, 25, 0, 391, 394, 3, 52, 26, 0, 392, 394,
		3, 82, 41, 0, 393, 386, 1, 0, 0, 0, 393, 387, 1, 0, 0, 0, 393, 388, 1,
		0, 0, 0, 393, 389, 1, 0, 0, 0, 393, 390, 1, 0, 0, 0, 393, 391, 1, 0, 0,
		0, 393, 392, 1, 0, 0, 0, 394, 81, 1, 0, 0, 0, 395, 396, 5, 48, 0, 0, 396,
		397, 3, 42, 21, 0, 397, 398, 5, 49, 0, 0, 398, 83, 1, 0, 0, 0, 399, 400,
		3, 56, 28, 0, 400, 401, 3, 86, 43, 0, 401, 402, 3, 116, 58, 0, 402, 85,
		1, 0, 0, 0, 403, 404, 7, 6, 0, 0, 404, 87, 1, 0, 0, 0, 405, 410, 3, 84,
		42, 0, 406, 407, 5, 50, 0, 0, 407, 409, 3, 84, 42, 0, 408, 406, 1, 0, 0,
		0, 409, 412, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411,
		414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 413, 415, 5, 50, 0, 0, 414, 413,
		1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 89, 1, 0, 0, 0, 416, 417, 5, 2,
		0, 0, 417, 426, 5, 48, 0, 0, 418, 423, 3, 92, 46, 0, 419, 420, 5, 50, 0,
		0, 420, 422, 3, 92, 46, 0, 421, 419, 1, 0, 0, 0, 422, 425, 1, 0, 0, 0,
		423, 421, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425,
		423, 1, 0, 0, 0, 426, 418, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428,
		1, 0, 0, 0, 428, 429, 5, 49, 0, 0, 429, 91, 1, 0, 0, 0, 430, 433, 3, 52,
		26, 0, 431, 433, 3, 42, 21, 0, 432, 430, 1, 0, 0, 0, 432, 431, 1, 0, 0,
		0, 433, 93, 1, 0, 0, 0, 434, 443, 5, 48, 0, 0, 435, 440, 3, 92, 46, 0,
		436, 437, 5, 50, 0, 0, 437, 439, 3, 92, 46, 0, 438, 436, 1, 0, 0, 0, 439,
		442, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 444,
		1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 443, 435, 1, 0, 0, 0, 443, 444, 1, 0,
		0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 5, 49, 0, 0, 446, 95, 1, 0, 0, 0,
		447, 448, 5, 1, 0, 0, 448, 463, 3, 94, 47, 0, 449, 452, 5, 1, 0, 0, 450,
		453, 3, 98, 49, 0, 451, 453, 3, 100, 50, 0, 452, 450, 1, 0, 0, 0, 452,
		451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 455, 3, 94, 47, 0, 455, 463,
		1, 0, 0, 0, 456, 457, 5, 1, 0, 0, 457, 460, 3, 94, 47, 0, 458, 461, 3,
		98, 49, 0, 459, 461, 3, 100, 50, 0, 460, 458, 1, 0, 0, 0, 460, 459, 1,
		0, 0, 0, 461, 463, 1, 0, 0, 0, 462, 447, 1, 0, 0, 0, 462, 449, 1, 0, 0,
		0, 462, 456, 1, 0, 0, 0, 463, 97, 1, 0, 0, 0, 464, 465, 5, 38, 0, 0, 465,
		466, 3, 112, 56, 0, 466, 99, 1, 0, 0, 0, 467, 468, 5, 39, 0, 0, 468, 469,
		3, 112, 56, 0, 469, 101, 1, 0, 0, 0, 470, 473, 3, 104, 52, 0, 471, 473,
		3, 106, 53, 0, 472, 470, 1, 0, 0, 0, 472, 471, 1, 0, 0, 0, 473, 476, 1,
		0, 0, 0, 474, 477, 3, 108, 54, 0, 475, 477, 3, 110, 55, 0, 476, 474, 1,
		0, 0, 0, 476, 475, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 103, 1, 0, 0,
		0, 478, 479, 5, 40, 0, 0, 479, 480, 3, 112, 56, 0, 480, 105, 1, 0, 0, 0,
		481, 482, 5, 41, 0, 0, 482, 483, 3, 112, 56, 0, 483, 107, 1, 0, 0, 0, 484,
		486, 5, 42, 0, 0, 485, 487, 3, 112, 56, 0, 486, 485, 1, 0, 0, 0, 486, 487,
		1, 0, 0, 0, 487, 109, 1, 0, 0, 0, 488, 490, 5, 43, 0, 0, 489, 491, 3, 112,
		56, 0, 490, 489, 1, 0, 0, 0, 490, 491, 1, 0, 0, 0, 491, 111, 1, 0, 0, 0,
		492, 501, 5, 48, 0, 0, 493, 498, 3, 56, 28, 0, 494, 495, 5, 50, 0, 0, 495,
		497, 3, 56, 28, 0, 496, 494, 1, 0, 0, 0, 497, 500, 1, 0, 0, 0, 498, 496,
		1, 0, 0, 0, 498, 499, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 498, 1, 0,
		0, 0, 501, 493, 1, 0, 0, 0, 501, 502, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0,
		503, 504, 5, 49, 0, 0, 504, 113, 1, 0, 0, 0, 505, 506, 7, 7, 0, 0, 506,
		115, 1, 0, 0, 0, 507, 508, 7, 8, 0, 0, 508, 117, 1, 0, 0, 0, 49, 125, 139,
		150, 159, 169, 177, 186, 198, 209, 223, 225, 232, 244, 256, 291, 293, 298,
		314, 318, 323, 326, 332, 337, 348, 357, 361, 365, 369, 372, 376, 380, 384,
		393, 410, 414, 423, 426, 432, 440, 443, 452, 460, 462, 472, 476, 486, 490,
		498, 501,
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
	PromQLExParserID_WS                = 57
	PromQLExParserRAW_STRING           = 58
	PromQLExParserBACKTICK_OPEN        = 59
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
	Statement_block() IStatement_blockContext
	EOF() antlr.TerminalNode

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

func (s *PromqlxContext) Statement_block() IStatement_blockContext {
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

func (s *PromqlxContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromQLExParserEOF, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Statement_block()
	}
	{
		p.SetState(119)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewState_AliasDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Alias_def()
		}

	case 2:
		localctx = NewState_MacroDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Macro_def()
		}

	case 3:
		localctx = NewState_IfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.If_statement()
		}

	case 4:
		localctx = NewState_VectorOperationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
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
		p.SetState(127)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(PromQLExParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
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
		p.SetState(132)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
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
	LEFT_BRACE() antlr.TerminalNode
	Statement_block() IStatement_blockContext
	RIGHT_BRACE() antlr.TerminalNode
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

func (s *Macro_defContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserLEFT_BRACE, 0)
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

func (s *Macro_defContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLExParserRIGHT_BRACE, 0)
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
		p.SetState(135)
		p.Match(PromQLExParserDEF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserID {
		{
			p.SetState(138)
			p.Args_decl()
		}

	}
	{
		p.SetState(141)
		p.Match(PromQLExParserRIGHT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(PromQLExParserLEFT_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
		p.Statement_block()
	}
	{
		p.SetState(144)
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

// IMacro_callContext is an interface to support dynamic dispatch.
type IMacro_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALL_SIGN() antlr.TerminalNode
	ID() antlr.TerminalNode
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

func (s *Macro_callContext) ID() antlr.TerminalNode {
	return s.GetToken(PromQLExParserID, 0)
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
		p.SetState(146)
		p.Match(PromQLExParserCALL_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(PromQLExParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&365143413538889774) != 0 {
		{
			p.SetState(149)
			p.Arg_list()
		}

	}
	{
		p.SetState(152)
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
		p.SetState(154)
		p.Arg_name()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLExParserCOMMA {
		{
			p.SetState(155)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Arg_name()
		}

		p.SetState(161)
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
		p.SetState(162)
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
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

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

func (s *Statement_blockContext) AllStatement() []IStatementContext {
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

func (s *Statement_blockContext) Statement(i int) IStatementContext {
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

func (s *Statement_blockContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(PromQLExParserSEMICOLON)
}

func (s *Statement_blockContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(PromQLExParserSEMICOLON, i)
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Statement()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(165)
				p.Match(PromQLExParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(166)
				p.Statement()
			}

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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
		p.SetState(172)
		p.vectorOperation(0)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PromQLExParserCOMMA {
		{
			p.SetState(173)
			p.Match(PromQLExParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
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
		p.SetState(179)
		p.Match(PromQLExParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Condition()
	}
	{
		p.SetState(181)
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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserMETRIC_NAME, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserSTRING, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_BRACE, PromQLExParserLEFT_PAREN, PromQLExParserDURATION, PromQLExParserID, PromQLExParserRAW_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.CompareVectorOperation()
		}

	case PromQLExParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.TrueConst()
		}

	case PromQLExParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
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
		p.SetState(188)
		p.vectorOperation(0)
	}
	{
		p.SetState(189)
		p.CompareOp()
	}
	{
		p.SetState(190)
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
		p.SetState(192)
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
		p.SetState(194)
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
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP:
		localctx = NewTimeInstLit_IsoDateTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(196)
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
			p.SetState(197)
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
	p.SetState(209)
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
			p.SetState(201)
			p.UnaryOp()
		}
		{
			p.SetState(202)
			p.const_num_expression(5)
		}

	case PromQLExParserLEFT_PAREN:
		localctx = NewConsNumExpr_ParenOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.const_num_expression(0)
		}
		{
			p.SetState(206)
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
			p.SetState(208)
			p.Num_literal()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(225)
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
			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConsNumExpr_PowerOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(212)
					p.PowOp()
				}
				{
					p.SetState(213)
					p.const_num_expression(6)
				}

			case 2:
				localctx = NewConsNumExpr_MulOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(216)
					p.MultOp()
				}
				{
					p.SetState(217)
					p.const_num_expression(5)
				}

			case 3:
				localctx = NewConsNumExpr_AddOpContext(p, NewConst_num_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_const_num_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(220)
					p.AddOp()
				}
				{
					p.SetState(221)
					p.const_num_expression(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(227)
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
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumLit_NumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
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
			p.SetState(229)
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
			p.SetState(230)
			p.Time_instant_literal()
		}

	case 4:
		localctx = NewNumLit_AliasCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
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
		p.SetState(234)
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
		p.SetState(236)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Duration()
	}
	{
		p.SetState(238)
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
		p.SetState(240)
		p.Match(PromQLExParserLEFT_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Duration()
	}
	{
		p.SetState(242)
		p.Match(PromQLExParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4785074604941320) != 0 {
		{
			p.SetState(243)
			p.Duration()
		}

	}
	{
		p.SetState(246)
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
	p.SetState(256)
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
			p.SetState(249)
			p.const_num_expression(0)
		}

	case 2:
		localctx = NewVecOp_UnaryOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(250)
			p.UnaryOp()
		}
		{
			p.SetState(251)
			p.vectorOperation(11)
		}

	case 3:
		localctx = NewVecOp_VecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(253)
			p.Vector()
		}

	case 4:
		localctx = NewVecOp_MacroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.Macro_call()
		}

	case 5:
		localctx = NewVecOp_AliasContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(255)
			p.Alias_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(293)
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
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVecOp_PowOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(259)
					p.PowOp()
				}
				{
					p.SetState(260)
					p.vectorOperation(13)
				}

			case 2:
				localctx = NewVecOp_MulOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(262)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(263)
					p.MultOp()
				}
				{
					p.SetState(264)
					p.vectorOperation(11)
				}

			case 3:
				localctx = NewVecOp_AddOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(267)
					p.AddOp()
				}
				{
					p.SetState(268)
					p.vectorOperation(10)
				}

			case 4:
				localctx = NewVecOp_CompareOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(270)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(271)
					p.CompareOp()
				}
				{
					p.SetState(272)
					p.vectorOperation(9)
				}

			case 5:
				localctx = NewVecOp_AndUnlessContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(275)
					p.AndUnlessOp()
				}
				{
					p.SetState(276)
					p.vectorOperation(8)
				}

			case 6:
				localctx = NewVecOp_OrOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(279)
					p.OrOp()
				}
				{
					p.SetState(280)
					p.vectorOperation(7)
				}

			case 7:
				localctx = NewVecOp_VecMatchOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					p.VectorMatchOp()
				}
				{
					p.SetState(284)
					p.vectorOperation(6)
				}

			case 8:
				localctx = NewVecOp_SubqueryOpContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(286)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(287)
					p.SubqueryOp()
				}

			case 9:
				localctx = NewVecOp_AtContext(p, NewVectorOperationContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PromQLExParserRULE_vectorOperation)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(289)
					p.Match(PromQLExParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(290)
					p.At_modifier_timestamp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(295)
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
		p.SetState(296)
		p.Subquery_range()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(297)
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
		p.SetState(300)
		p.Match(PromQLExParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
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
		p.SetState(303)
		p.InstantSelector()
	}
	{
		p.SetState(304)
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
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.InstantSelector()
		}
		{
			p.SetState(307)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Duration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.MatrixSelector()
		}
		{
			p.SetState(311)
			p.Match(PromQLExParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
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
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_PAREN, PromQLExParserDURATION:
		localctx = NewLit_ConstNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.const_num_expression(0)
		}

	case PromQLExParserSTRING, PromQLExParserRAW_STRING:
		localctx = NewLit_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
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

	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Metric_name()
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(321)
				p.Match(PromQLExParserLEFT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81134887276380198) != 0 {
				{
					p.SetState(322)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(325)
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
			p.SetState(328)
			p.Match(PromQLExParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.LabelMatcherList()
		}
		{
			p.SetState(330)
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
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserAGGREGATION_OPERATOR, PromQLExParserFUNCTION, PromQLExParserAND, PromQLExParserOR, PromQLExParserUNLESS, PromQLExParserBY, PromQLExParserWITHOUT, PromQLExParserON, PromQLExParserIGNORING, PromQLExParserGROUP_LEFT, PromQLExParserGROUP_RIGHT, PromQLExParserOFFSET, PromQLExParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.Keyword()
		}

	case PromQLExParserMETRIC_NAME, PromQLExParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(335)
			p.Metric_name()
		}

	case PromQLExParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(336)
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
		p.SetState(339)
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
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserRFC_3339_TIMESTAMP, PromQLExParserCALL_SIGN, PromQLExParserNUMBER, PromQLExParserADD, PromQLExParserSUB, PromQLExParserLEFT_PAREN, PromQLExParserDURATION:
		localctx = NewAtModTime_ConstNumExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.const_num_expression(0)
		}

	case PromQLExParserSTART:
		localctx = NewAtModTime_StartContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(PromQLExParserSTART)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
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
			p.SetState(345)
			p.Match(PromQLExParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)
			p.Match(PromQLExParserLEFT_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
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
		p.SetState(350)
		p.vectorOperation(0)
	}
	{
		p.SetState(351)
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
		p.SetState(353)
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
		p.SetState(355)
		p.Match(PromQLExParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(356)
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
		p.SetState(359)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(360)
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
		p.SetState(363)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserADD || _la == PromQLExParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(364)
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
		p.SetState(367)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67645734912) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserBOOL {
		{
			p.SetState(368)
			p.Match(PromQLExParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(371)
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
		p.SetState(374)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserAND || _la == PromQLExParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(375)
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
		p.SetState(378)
		p.Match(PromQLExParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
		{
			p.SetState(379)
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
		p.SetState(382)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLExParserUNLESS || _la == PromQLExParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserON || _la == PromQLExParserIGNORING {
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
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(387)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(388)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(389)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(390)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(391)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(392)
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
		p.SetState(395)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(396)
		p.vectorOperation(0)
	}
	{
		p.SetState(397)
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
		p.SetState(399)
		p.LabelName()
	}
	{
		p.SetState(400)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(401)
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
		p.SetState(403)
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
		p.SetState(405)
		p.LabelMatcher()
	}
	p.SetState(410)
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
				p.SetState(406)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(407)
				p.LabelMatcher()
			}

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
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLExParserCOMMA {
		{
			p.SetState(413)
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
		p.SetState(416)
		p.Match(PromQLExParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&365143413538889774) != 0 {
		{
			p.SetState(418)
			p.Parameter()
		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(419)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
		}

	}
	{
		p.SetState(428)
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
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(430)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(431)
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
		p.SetState(434)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&365143413538889774) != 0 {
		{
			p.SetState(435)
			p.Parameter()
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(436)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
		}

	}
	{
		p.SetState(445)
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
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(447)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(450)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(451)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(454)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
			p.Match(PromQLExParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.ParameterList()
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLExParserBY:
			{
				p.SetState(458)
				p.By()
			}

		case PromQLExParserWITHOUT:
			{
				p.SetState(459)
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
		p.SetState(464)
		p.Match(PromQLExParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
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
		p.SetState(467)
		p.Match(PromQLExParserWITHOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
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
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLExParserON:
		{
			p.SetState(470)
			p.On_()
		}

	case PromQLExParserIGNORING:
		{
			p.SetState(471)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLExParserGROUP_LEFT:
		{
			p.SetState(474)
			p.GroupLeft()
		}

	case PromQLExParserGROUP_RIGHT:
		{
			p.SetState(475)
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
		p.SetState(478)
		p.Match(PromQLExParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(479)
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
		p.SetState(481)
		p.Match(PromQLExParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(482)
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
		p.SetState(484)
		p.Match(PromQLExParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(485)
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
		p.SetState(488)
		p.Match(PromQLExParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(489)
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
		p.SetState(492)
		p.Match(PromQLExParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&81134887276380198) != 0 {
		{
			p.SetState(493)
			p.LabelName()
		}
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLExParserCOMMA {
			{
				p.SetState(494)
				p.Match(PromQLExParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
		}

	}
	{
		p.SetState(503)
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
		p.SetState(505)
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
		p.SetState(507)
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

// Code generated from PromQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promql // PromQLParser
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

type PromQLParser struct {
	*antlr.BaseParser
}

var PromQLParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promqlparserParserInit() {
	staticData := &PromQLParserParserStaticData
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
		"WS", "SL_COMMENT", "RAW_STRING", "BACKTICK_OPEN",
	}
	staticData.RuleNames = []string{
		"expression", "vectorOperation", "unaryOp", "powOp", "multOp", "addOp",
		"compareOp", "andUnlessOp", "orOp", "vectorMatchOp", "subqueryOp", "offsetOp",
		"vector", "parens", "instantSelector", "labelMatcher", "labelMatcherOperator",
		"labelMatcherList", "matrixSelector", "offset", "function_", "parameter",
		"parameterList", "aggregation", "by", "without", "grouping", "on_",
		"ignoring", "groupLeft", "groupRight", "labelName", "labelNameList",
		"keyword", "literal", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 320, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 116, 8, 1, 10, 1, 12, 1, 119, 9,
		1, 1, 2, 1, 2, 1, 3, 1, 3, 3, 3, 125, 8, 3, 1, 4, 1, 4, 3, 4, 129, 8, 4,
		1, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6, 3, 6, 137, 8, 6, 1, 6, 3, 6, 140,
		8, 6, 1, 7, 1, 7, 3, 7, 144, 8, 7, 1, 8, 1, 8, 3, 8, 148, 8, 8, 1, 9, 1,
		9, 3, 9, 152, 8, 9, 1, 10, 1, 10, 3, 10, 156, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 168, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 177, 8, 14, 1, 14,
		3, 14, 180, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 186, 8, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 5, 17, 197, 8,
		17, 10, 17, 12, 17, 200, 9, 17, 1, 17, 3, 17, 203, 8, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 216,
		8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 223, 8, 20, 10, 20, 12,
		20, 226, 9, 20, 3, 20, 228, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 234,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 240, 8, 22, 10, 22, 12, 22, 243,
		9, 22, 3, 22, 245, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 3, 23, 254, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		262, 8, 23, 3, 23, 264, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 3, 26, 274, 8, 26, 1, 26, 1, 26, 3, 26, 278, 8, 26, 1, 27,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 3, 29, 288, 8, 29, 1,
		30, 1, 30, 3, 30, 292, 8, 30, 1, 31, 1, 31, 1, 31, 3, 31, 297, 8, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 5, 32, 303, 8, 32, 10, 32, 12, 32, 306, 9, 32,
		3, 32, 308, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 3, 34, 316,
		8, 34, 1, 35, 1, 35, 1, 35, 0, 1, 2, 36, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 0, 8, 1, 0, 5, 6, 1, 0, 7, 9, 1, 0,
		15, 20, 2, 0, 11, 11, 13, 13, 2, 0, 13, 13, 25, 25, 3, 0, 14, 14, 16, 16,
		21, 22, 3, 0, 1, 2, 11, 13, 23, 30, 2, 0, 4, 4, 46, 46, 333, 0, 72, 1,
		0, 0, 0, 2, 80, 1, 0, 0, 0, 4, 120, 1, 0, 0, 0, 6, 122, 1, 0, 0, 0, 8,
		126, 1, 0, 0, 0, 10, 130, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 141, 1,
		0, 0, 0, 16, 145, 1, 0, 0, 0, 18, 149, 1, 0, 0, 0, 20, 153, 1, 0, 0, 0,
		22, 157, 1, 0, 0, 0, 24, 167, 1, 0, 0, 0, 26, 169, 1, 0, 0, 0, 28, 185,
		1, 0, 0, 0, 30, 187, 1, 0, 0, 0, 32, 191, 1, 0, 0, 0, 34, 193, 1, 0, 0,
		0, 36, 204, 1, 0, 0, 0, 38, 215, 1, 0, 0, 0, 40, 217, 1, 0, 0, 0, 42, 233,
		1, 0, 0, 0, 44, 235, 1, 0, 0, 0, 46, 263, 1, 0, 0, 0, 48, 265, 1, 0, 0,
		0, 50, 268, 1, 0, 0, 0, 52, 273, 1, 0, 0, 0, 54, 279, 1, 0, 0, 0, 56, 282,
		1, 0, 0, 0, 58, 285, 1, 0, 0, 0, 60, 289, 1, 0, 0, 0, 62, 296, 1, 0, 0,
		0, 64, 298, 1, 0, 0, 0, 66, 311, 1, 0, 0, 0, 68, 315, 1, 0, 0, 0, 70, 317,
		1, 0, 0, 0, 72, 73, 3, 2, 1, 0, 73, 74, 5, 0, 0, 1, 74, 1, 1, 0, 0, 0,
		75, 76, 6, 1, -1, 0, 76, 77, 3, 4, 2, 0, 77, 78, 3, 2, 1, 9, 78, 81, 1,
		0, 0, 0, 79, 81, 3, 24, 12, 0, 80, 75, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0,
		81, 117, 1, 0, 0, 0, 82, 83, 10, 11, 0, 0, 83, 84, 3, 6, 3, 0, 84, 85,
		3, 2, 1, 11, 85, 116, 1, 0, 0, 0, 86, 87, 10, 8, 0, 0, 87, 88, 3, 8, 4,
		0, 88, 89, 3, 2, 1, 9, 89, 116, 1, 0, 0, 0, 90, 91, 10, 7, 0, 0, 91, 92,
		3, 10, 5, 0, 92, 93, 3, 2, 1, 8, 93, 116, 1, 0, 0, 0, 94, 95, 10, 6, 0,
		0, 95, 96, 3, 12, 6, 0, 96, 97, 3, 2, 1, 7, 97, 116, 1, 0, 0, 0, 98, 99,
		10, 5, 0, 0, 99, 100, 3, 14, 7, 0, 100, 101, 3, 2, 1, 6, 101, 116, 1, 0,
		0, 0, 102, 103, 10, 4, 0, 0, 103, 104, 3, 16, 8, 0, 104, 105, 3, 2, 1,
		5, 105, 116, 1, 0, 0, 0, 106, 107, 10, 3, 0, 0, 107, 108, 3, 18, 9, 0,
		108, 109, 3, 2, 1, 4, 109, 116, 1, 0, 0, 0, 110, 111, 10, 2, 0, 0, 111,
		112, 5, 38, 0, 0, 112, 116, 3, 2, 1, 3, 113, 114, 10, 10, 0, 0, 114, 116,
		3, 20, 10, 0, 115, 82, 1, 0, 0, 0, 115, 86, 1, 0, 0, 0, 115, 90, 1, 0,
		0, 0, 115, 94, 1, 0, 0, 0, 115, 98, 1, 0, 0, 0, 115, 102, 1, 0, 0, 0, 115,
		106, 1, 0, 0, 0, 115, 110, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 119,
		1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 3, 1, 0, 0,
		0, 119, 117, 1, 0, 0, 0, 120, 121, 7, 0, 0, 0, 121, 5, 1, 0, 0, 0, 122,
		124, 5, 10, 0, 0, 123, 125, 3, 52, 26, 0, 124, 123, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 7, 1, 0, 0, 0, 126, 128, 7, 1, 0, 0, 127, 129, 3, 52,
		26, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 9, 1, 0, 0, 0,
		130, 132, 7, 0, 0, 0, 131, 133, 3, 52, 26, 0, 132, 131, 1, 0, 0, 0, 132,
		133, 1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134, 136, 7, 2, 0, 0, 135, 137, 5,
		30, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0,
		0, 138, 140, 3, 52, 26, 0, 139, 138, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 13, 1, 0, 0, 0, 141, 143, 7, 3, 0, 0, 142, 144, 3, 52, 26, 0, 143,
		142, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 15, 1, 0, 0, 0, 145, 147, 5,
		12, 0, 0, 146, 148, 3, 52, 26, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 17, 1, 0, 0, 0, 149, 151, 7, 4, 0, 0, 150, 152, 3, 52, 26, 0,
		151, 150, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 19, 1, 0, 0, 0, 153, 155,
		5, 39, 0, 0, 154, 156, 3, 22, 11, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 21, 1, 0, 0, 0, 157, 158, 5, 29, 0, 0, 158, 159, 5, 41, 0,
		0, 159, 23, 1, 0, 0, 0, 160, 168, 3, 40, 20, 0, 161, 168, 3, 46, 23, 0,
		162, 168, 3, 28, 14, 0, 163, 168, 3, 36, 18, 0, 164, 168, 3, 38, 19, 0,
		165, 168, 3, 68, 34, 0, 166, 168, 3, 26, 13, 0, 167, 160, 1, 0, 0, 0, 167,
		161, 1, 0, 0, 0, 167, 162, 1, 0, 0, 0, 167, 163, 1, 0, 0, 0, 167, 164,
		1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 25, 1, 0,
		0, 0, 169, 170, 5, 33, 0, 0, 170, 171, 3, 2, 1, 0, 171, 172, 5, 34, 0,
		0, 172, 27, 1, 0, 0, 0, 173, 179, 5, 42, 0, 0, 174, 176, 5, 31, 0, 0, 175,
		177, 3, 34, 17, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 180, 5, 32, 0, 0, 179, 174, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 186, 1, 0, 0, 0, 181, 182, 5, 31, 0, 0, 182, 183, 3, 34, 17,
		0, 183, 184, 5, 32, 0, 0, 184, 186, 1, 0, 0, 0, 185, 173, 1, 0, 0, 0, 185,
		181, 1, 0, 0, 0, 186, 29, 1, 0, 0, 0, 187, 188, 3, 62, 31, 0, 188, 189,
		3, 32, 16, 0, 189, 190, 3, 70, 35, 0, 190, 31, 1, 0, 0, 0, 191, 192, 7,
		5, 0, 0, 192, 33, 1, 0, 0, 0, 193, 198, 3, 30, 15, 0, 194, 195, 5, 37,
		0, 0, 195, 197, 3, 30, 15, 0, 196, 194, 1, 0, 0, 0, 197, 200, 1, 0, 0,
		0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200,
		198, 1, 0, 0, 0, 201, 203, 5, 37, 0, 0, 202, 201, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 35, 1, 0, 0, 0, 204, 205, 3, 28, 14, 0, 205, 206, 5, 40,
		0, 0, 206, 37, 1, 0, 0, 0, 207, 208, 3, 28, 14, 0, 208, 209, 5, 29, 0,
		0, 209, 210, 5, 41, 0, 0, 210, 216, 1, 0, 0, 0, 211, 212, 3, 36, 18, 0,
		212, 213, 5, 29, 0, 0, 213, 214, 5, 41, 0, 0, 214, 216, 1, 0, 0, 0, 215,
		207, 1, 0, 0, 0, 215, 211, 1, 0, 0, 0, 216, 39, 1, 0, 0, 0, 217, 218, 5,
		2, 0, 0, 218, 227, 5, 33, 0, 0, 219, 224, 3, 42, 21, 0, 220, 221, 5, 37,
		0, 0, 221, 223, 3, 42, 21, 0, 222, 220, 1, 0, 0, 0, 223, 226, 1, 0, 0,
		0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226,
		224, 1, 0, 0, 0, 227, 219, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229,
		1, 0, 0, 0, 229, 230, 5, 34, 0, 0, 230, 41, 1, 0, 0, 0, 231, 234, 3, 68,
		34, 0, 232, 234, 3, 2, 1, 0, 233, 231, 1, 0, 0, 0, 233, 232, 1, 0, 0, 0,
		234, 43, 1, 0, 0, 0, 235, 244, 5, 33, 0, 0, 236, 241, 3, 42, 21, 0, 237,
		238, 5, 37, 0, 0, 238, 240, 3, 42, 21, 0, 239, 237, 1, 0, 0, 0, 240, 243,
		1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 245, 1, 0,
		0, 0, 243, 241, 1, 0, 0, 0, 244, 236, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0,
		245, 246, 1, 0, 0, 0, 246, 247, 5, 34, 0, 0, 247, 45, 1, 0, 0, 0, 248,
		249, 5, 1, 0, 0, 249, 264, 3, 44, 22, 0, 250, 253, 5, 1, 0, 0, 251, 254,
		3, 48, 24, 0, 252, 254, 3, 50, 25, 0, 253, 251, 1, 0, 0, 0, 253, 252, 1,
		0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 3, 44, 22, 0, 256, 264, 1, 0,
		0, 0, 257, 258, 5, 1, 0, 0, 258, 261, 3, 44, 22, 0, 259, 262, 3, 48, 24,
		0, 260, 262, 3, 50, 25, 0, 261, 259, 1, 0, 0, 0, 261, 260, 1, 0, 0, 0,
		262, 264, 1, 0, 0, 0, 263, 248, 1, 0, 0, 0, 263, 250, 1, 0, 0, 0, 263,
		257, 1, 0, 0, 0, 264, 47, 1, 0, 0, 0, 265, 266, 5, 23, 0, 0, 266, 267,
		3, 64, 32, 0, 267, 49, 1, 0, 0, 0, 268, 269, 5, 24, 0, 0, 269, 270, 3,
		64, 32, 0, 270, 51, 1, 0, 0, 0, 271, 274, 3, 54, 27, 0, 272, 274, 3, 56,
		28, 0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0,
		275, 278, 3, 58, 29, 0, 276, 278, 3, 60, 30, 0, 277, 275, 1, 0, 0, 0, 277,
		276, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 53, 1, 0, 0, 0, 279, 280, 5,
		25, 0, 0, 280, 281, 3, 64, 32, 0, 281, 55, 1, 0, 0, 0, 282, 283, 5, 26,
		0, 0, 283, 284, 3, 64, 32, 0, 284, 57, 1, 0, 0, 0, 285, 287, 5, 27, 0,
		0, 286, 288, 3, 64, 32, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0,
		288, 59, 1, 0, 0, 0, 289, 291, 5, 28, 0, 0, 290, 292, 3, 64, 32, 0, 291,
		290, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 61, 1, 0, 0, 0, 293, 297, 3,
		66, 33, 0, 294, 297, 5, 42, 0, 0, 295, 297, 5, 43, 0, 0, 296, 293, 1, 0,
		0, 0, 296, 294, 1, 0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 63, 1, 0, 0, 0,
		298, 307, 5, 33, 0, 0, 299, 304, 3, 62, 31, 0, 300, 301, 5, 37, 0, 0, 301,
		303, 3, 62, 31, 0, 302, 300, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302,
		1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0,
		0, 0, 307, 299, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0,
		309, 310, 5, 34, 0, 0, 310, 65, 1, 0, 0, 0, 311, 312, 7, 6, 0, 0, 312,
		67, 1, 0, 0, 0, 313, 316, 5, 3, 0, 0, 314, 316, 3, 70, 35, 0, 315, 313,
		1, 0, 0, 0, 315, 314, 1, 0, 0, 0, 316, 69, 1, 0, 0, 0, 317, 318, 7, 7,
		0, 0, 318, 71, 1, 0, 0, 0, 35, 80, 115, 117, 124, 128, 132, 136, 139, 143,
		147, 151, 155, 167, 176, 179, 185, 198, 202, 215, 224, 227, 233, 241, 244,
		253, 261, 263, 273, 277, 287, 291, 296, 304, 307, 315,
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

// PromQLParserInit initializes any static state used to implement PromQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLParserInit() {
	staticData := &PromQLParserParserStaticData
	staticData.once.Do(promqlparserParserInit)
}

// NewPromQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromQLParser(input antlr.TokenStream) *PromQLParser {
	PromQLParserInit()
	this := new(PromQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromQLParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromQLParser.g4"

	return this
}

// PromQLParser tokens.
const (
	PromQLParserEOF                  = antlr.TokenEOF
	PromQLParserAGGREGATION_OPERATOR = 1
	PromQLParserFUNCTION             = 2
	PromQLParserNUMBER               = 3
	PromQLParserSTRING               = 4
	PromQLParserADD                  = 5
	PromQLParserSUB                  = 6
	PromQLParserMULT                 = 7
	PromQLParserDIV                  = 8
	PromQLParserMOD                  = 9
	PromQLParserPOW                  = 10
	PromQLParserAND                  = 11
	PromQLParserOR                   = 12
	PromQLParserUNLESS               = 13
	PromQLParserEQ                   = 14
	PromQLParserDEQ                  = 15
	PromQLParserNE                   = 16
	PromQLParserGT                   = 17
	PromQLParserLT                   = 18
	PromQLParserGE                   = 19
	PromQLParserLE                   = 20
	PromQLParserRE                   = 21
	PromQLParserNRE                  = 22
	PromQLParserBY                   = 23
	PromQLParserWITHOUT              = 24
	PromQLParserON                   = 25
	PromQLParserIGNORING             = 26
	PromQLParserGROUP_LEFT           = 27
	PromQLParserGROUP_RIGHT          = 28
	PromQLParserOFFSET               = 29
	PromQLParserBOOL                 = 30
	PromQLParserLEFT_BRACE           = 31
	PromQLParserRIGHT_BRACE          = 32
	PromQLParserLEFT_PAREN           = 33
	PromQLParserRIGHT_PAREN          = 34
	PromQLParserLEFT_BRACKET         = 35
	PromQLParserRIGHT_BRACKET        = 36
	PromQLParserCOMMA                = 37
	PromQLParserAT                   = 38
	PromQLParserSUBQUERY_RANGE       = 39
	PromQLParserTIME_RANGE           = 40
	PromQLParserDURATION             = 41
	PromQLParserMETRIC_NAME          = 42
	PromQLParserLABEL_NAME           = 43
	PromQLParserWS                   = 44
	PromQLParserSL_COMMENT           = 45
	PromQLParserRAW_STRING           = 46
	PromQLParserBACKTICK_OPEN        = 47
)

// PromQLParser rules.
const (
	PromQLParserRULE_expression           = 0
	PromQLParserRULE_vectorOperation      = 1
	PromQLParserRULE_unaryOp              = 2
	PromQLParserRULE_powOp                = 3
	PromQLParserRULE_multOp               = 4
	PromQLParserRULE_addOp                = 5
	PromQLParserRULE_compareOp            = 6
	PromQLParserRULE_andUnlessOp          = 7
	PromQLParserRULE_orOp                 = 8
	PromQLParserRULE_vectorMatchOp        = 9
	PromQLParserRULE_subqueryOp           = 10
	PromQLParserRULE_offsetOp             = 11
	PromQLParserRULE_vector               = 12
	PromQLParserRULE_parens               = 13
	PromQLParserRULE_instantSelector      = 14
	PromQLParserRULE_labelMatcher         = 15
	PromQLParserRULE_labelMatcherOperator = 16
	PromQLParserRULE_labelMatcherList     = 17
	PromQLParserRULE_matrixSelector       = 18
	PromQLParserRULE_offset               = 19
	PromQLParserRULE_function_            = 20
	PromQLParserRULE_parameter            = 21
	PromQLParserRULE_parameterList        = 22
	PromQLParserRULE_aggregation          = 23
	PromQLParserRULE_by                   = 24
	PromQLParserRULE_without              = 25
	PromQLParserRULE_grouping             = 26
	PromQLParserRULE_on_                  = 27
	PromQLParserRULE_ignoring             = 28
	PromQLParserRULE_groupLeft            = 29
	PromQLParserRULE_groupRight           = 30
	PromQLParserRULE_labelName            = 31
	PromQLParserRULE_labelNameList        = 32
	PromQLParserRULE_keyword              = 33
	PromQLParserRULE_literal              = 34
	PromQLParserRULE_string               = 35
)

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
	p.RuleIndex = PromQLParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_expression

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
	return s.GetToken(PromQLParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PromQLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromQLParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.vectorOperation(0)
	}
	{
		p.SetState(73)
		p.Match(PromQLParserEOF)
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
	p.RuleIndex = PromQLParserRULE_vectorOperation
	return p
}

func InitEmptyVectorOperationContext(p *VectorOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_vectorOperation
}

func (*VectorOperationContext) IsVectorOperationContext() {}

func NewVectorOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorOperationContext {
	var p = new(VectorOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_vectorOperation

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
	return s.GetToken(PromQLParserAT, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterVectorOperation(s)
	}
}

func (s *VectorOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitVectorOperation(s)
	}
}

func (p *PromQLParser) VectorOperation() (localctx IVectorOperationContext) {
	return p.vectorOperation(0)
}

func (p *PromQLParser) vectorOperation(_p int) (localctx IVectorOperationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewVectorOperationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IVectorOperationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PromQLParserRULE_vectorOperation, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLParserADD, PromQLParserSUB:
		{
			p.SetState(76)
			p.UnaryOp()
		}
		{
			p.SetState(77)
			p.vectorOperation(9)
		}

	case PromQLParserAGGREGATION_OPERATOR, PromQLParserFUNCTION, PromQLParserNUMBER, PromQLParserSTRING, PromQLParserLEFT_BRACE, PromQLParserLEFT_PAREN, PromQLParserMETRIC_NAME, PromQLParserRAW_STRING:
		{
			p.SetState(79)
			p.Vector()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(83)
					p.PowOp()
				}
				{
					p.SetState(84)
					p.vectorOperation(11)
				}

			case 2:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(87)
					p.MultOp()
				}
				{
					p.SetState(88)
					p.vectorOperation(9)
				}

			case 3:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.AddOp()
				}
				{
					p.SetState(92)
					p.vectorOperation(8)
				}

			case 4:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(94)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(95)
					p.CompareOp()
				}
				{
					p.SetState(96)
					p.vectorOperation(7)
				}

			case 5:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.AndUnlessOp()
				}
				{
					p.SetState(100)
					p.vectorOperation(6)
				}

			case 6:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(102)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(103)
					p.OrOp()
				}
				{
					p.SetState(104)
					p.vectorOperation(5)
				}

			case 7:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(107)
					p.VectorMatchOp()
				}
				{
					p.SetState(108)
					p.vectorOperation(4)
				}

			case 8:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(111)
					p.Match(PromQLParserAT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(112)
					p.vectorOperation(3)
				}

			case 9:
				localctx = NewVectorOperationContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PromQLParserRULE_vectorOperation)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					p.SubqueryOp()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.RuleIndex = PromQLParserRULE_unaryOp
	return p
}

func InitEmptyUnaryOpContext(p *UnaryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_unaryOp
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLParserADD, 0)
}

func (s *UnaryOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLParserSUB, 0)
}

func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterUnaryOp(s)
	}
}

func (s *UnaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitUnaryOp(s)
	}
}

func (p *PromQLParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromQLParserRULE_unaryOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLParserADD || _la == PromQLParserSUB) {
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
	p.RuleIndex = PromQLParserRULE_powOp
	return p
}

func InitEmptyPowOpContext(p *PowOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_powOp
}

func (*PowOpContext) IsPowOpContext() {}

func NewPowOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowOpContext {
	var p = new(PowOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_powOp

	return p
}

func (s *PowOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PowOpContext) POW() antlr.TerminalNode {
	return s.GetToken(PromQLParserPOW, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterPowOp(s)
	}
}

func (s *PowOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitPowOp(s)
	}
}

func (p *PromQLParser) PowOp() (localctx IPowOpContext) {
	localctx = NewPowOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromQLParserRULE_powOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(PromQLParserPOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(123)
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
	p.RuleIndex = PromQLParserRULE_multOp
	return p
}

func InitEmptyMultOpContext(p *MultOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_multOp
}

func (*MultOpContext) IsMultOpContext() {}

func NewMultOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOpContext {
	var p = new(MultOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_multOp

	return p
}

func (s *MultOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOpContext) MULT() antlr.TerminalNode {
	return s.GetToken(PromQLParserMULT, 0)
}

func (s *MultOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(PromQLParserDIV, 0)
}

func (s *MultOpContext) MOD() antlr.TerminalNode {
	return s.GetToken(PromQLParserMOD, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterMultOp(s)
	}
}

func (s *MultOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitMultOp(s)
	}
}

func (p *PromQLParser) MultOp() (localctx IMultOpContext) {
	localctx = NewMultOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromQLParserRULE_multOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(127)
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
	p.RuleIndex = PromQLParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(PromQLParserADD, 0)
}

func (s *AddOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(PromQLParserSUB, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (p *PromQLParser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromQLParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLParserADD || _la == PromQLParserSUB) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(131)
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
	p.RuleIndex = PromQLParserRULE_compareOp
	return p
}

func InitEmptyCompareOpContext(p *CompareOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_compareOp
}

func (*CompareOpContext) IsCompareOpContext() {}

func NewCompareOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOpContext {
	var p = new(CompareOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_compareOp

	return p
}

func (s *CompareOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompareOpContext) DEQ() antlr.TerminalNode {
	return s.GetToken(PromQLParserDEQ, 0)
}

func (s *CompareOpContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLParserNE, 0)
}

func (s *CompareOpContext) GT() antlr.TerminalNode {
	return s.GetToken(PromQLParserGT, 0)
}

func (s *CompareOpContext) LT() antlr.TerminalNode {
	return s.GetToken(PromQLParserLT, 0)
}

func (s *CompareOpContext) GE() antlr.TerminalNode {
	return s.GetToken(PromQLParserGE, 0)
}

func (s *CompareOpContext) LE() antlr.TerminalNode {
	return s.GetToken(PromQLParserLE, 0)
}

func (s *CompareOpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLParserBOOL, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterCompareOp(s)
	}
}

func (s *CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitCompareOp(s)
	}
}

func (p *PromQLParser) CompareOp() (localctx ICompareOpContext) {
	localctx = NewCompareOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromQLParserRULE_compareOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2064384) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserBOOL {
		{
			p.SetState(135)
			p.Match(PromQLParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(138)
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
	p.RuleIndex = PromQLParserRULE_andUnlessOp
	return p
}

func InitEmptyAndUnlessOpContext(p *AndUnlessOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_andUnlessOp
}

func (*AndUnlessOpContext) IsAndUnlessOpContext() {}

func NewAndUnlessOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndUnlessOpContext {
	var p = new(AndUnlessOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_andUnlessOp

	return p
}

func (s *AndUnlessOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AndUnlessOpContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLParserAND, 0)
}

func (s *AndUnlessOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLParserUNLESS, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterAndUnlessOp(s)
	}
}

func (s *AndUnlessOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitAndUnlessOp(s)
	}
}

func (p *PromQLParser) AndUnlessOp() (localctx IAndUnlessOpContext) {
	localctx = NewAndUnlessOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromQLParserRULE_andUnlessOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLParserAND || _la == PromQLParserUNLESS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(142)
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
	p.RuleIndex = PromQLParserRULE_orOp
	return p
}

func InitEmptyOrOpContext(p *OrOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_orOp
}

func (*OrOpContext) IsOrOpContext() {}

func NewOrOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrOpContext {
	var p = new(OrOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_orOp

	return p
}

func (s *OrOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OrOpContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLParserOR, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterOrOp(s)
	}
}

func (s *OrOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitOrOp(s)
	}
}

func (p *PromQLParser) OrOp() (localctx IOrOpContext) {
	localctx = NewOrOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromQLParserRULE_orOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(PromQLParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(146)
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
	p.RuleIndex = PromQLParserRULE_vectorMatchOp
	return p
}

func InitEmptyVectorMatchOpContext(p *VectorMatchOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_vectorMatchOp
}

func (*VectorMatchOpContext) IsVectorMatchOpContext() {}

func NewVectorMatchOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorMatchOpContext {
	var p = new(VectorMatchOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_vectorMatchOp

	return p
}

func (s *VectorMatchOpContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorMatchOpContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLParserON, 0)
}

func (s *VectorMatchOpContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLParserUNLESS, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterVectorMatchOp(s)
	}
}

func (s *VectorMatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitVectorMatchOp(s)
	}
}

func (p *PromQLParser) VectorMatchOp() (localctx IVectorMatchOpContext) {
	localctx = NewVectorMatchOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromQLParserRULE_vectorMatchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLParserUNLESS || _la == PromQLParserON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserON || _la == PromQLParserIGNORING {
		{
			p.SetState(150)
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
	p.RuleIndex = PromQLParserRULE_subqueryOp
	return p
}

func InitEmptySubqueryOpContext(p *SubqueryOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_subqueryOp
}

func (*SubqueryOpContext) IsSubqueryOpContext() {}

func NewSubqueryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryOpContext {
	var p = new(SubqueryOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_subqueryOp

	return p
}

func (s *SubqueryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryOpContext) SUBQUERY_RANGE() antlr.TerminalNode {
	return s.GetToken(PromQLParserSUBQUERY_RANGE, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterSubqueryOp(s)
	}
}

func (s *SubqueryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitSubqueryOp(s)
	}
}

func (p *PromQLParser) SubqueryOp() (localctx ISubqueryOpContext) {
	localctx = NewSubqueryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromQLParserRULE_subqueryOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(PromQLParserSUBQUERY_RANGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(154)
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
	p.RuleIndex = PromQLParserRULE_offsetOp
	return p
}

func InitEmptyOffsetOpContext(p *OffsetOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_offsetOp
}

func (*OffsetOpContext) IsOffsetOpContext() {}

func NewOffsetOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetOpContext {
	var p = new(OffsetOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_offsetOp

	return p
}

func (s *OffsetOpContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetOpContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLParserOFFSET, 0)
}

func (s *OffsetOpContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLParserDURATION, 0)
}

func (s *OffsetOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterOffsetOp(s)
	}
}

func (s *OffsetOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitOffsetOp(s)
	}
}

func (p *PromQLParser) OffsetOp() (localctx IOffsetOpContext) {
	localctx = NewOffsetOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PromQLParserRULE_offsetOp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(PromQLParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(PromQLParserDURATION)
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
	p.RuleIndex = PromQLParserRULE_vector
	return p
}

func InitEmptyVectorContext(p *VectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_vector
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_vector

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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitVector(s)
	}
}

func (p *PromQLParser) Vector() (localctx IVectorContext) {
	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromQLParserRULE_vector)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Function_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Aggregation()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.InstantSelector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(163)
			p.MatrixSelector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(164)
			p.Offset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(165)
			p.Literal()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(166)
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
	p.RuleIndex = PromQLParserRULE_parens
	return p
}

func InitEmptyParensContext(p *ParensContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_parens
}

func (*ParensContext) IsParensContext() {}

func NewParensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParensContext {
	var p = new(ParensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_parens

	return p
}

func (s *ParensContext) GetParser() antlr.Parser { return s.parser }

func (s *ParensContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserLEFT_PAREN, 0)
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
	return s.GetToken(PromQLParserRIGHT_PAREN, 0)
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitParens(s)
	}
}

func (p *PromQLParser) Parens() (localctx IParensContext) {
	localctx = NewParensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromQLParserRULE_parens)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(PromQLParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.vectorOperation(0)
	}
	{
		p.SetState(171)
		p.Match(PromQLParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLParserRULE_instantSelector
	return p
}

func InitEmptyInstantSelectorContext(p *InstantSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_instantSelector
}

func (*InstantSelectorContext) IsInstantSelectorContext() {}

func NewInstantSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstantSelectorContext {
	var p = new(InstantSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_instantSelector

	return p
}

func (s *InstantSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *InstantSelectorContext) METRIC_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLParserMETRIC_NAME, 0)
}

func (s *InstantSelectorContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLParserLEFT_BRACE, 0)
}

func (s *InstantSelectorContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(PromQLParserRIGHT_BRACE, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterInstantSelector(s)
	}
}

func (s *InstantSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitInstantSelector(s)
	}
}

func (p *PromQLParser) InstantSelector() (localctx IInstantSelectorContext) {
	localctx = NewInstantSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PromQLParserRULE_instantSelector)
	var _la int

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(PromQLParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(174)
				p.Match(PromQLParserLEFT_BRACE)
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

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13196278642694) != 0 {
				{
					p.SetState(175)
					p.LabelMatcherList()
				}

			}
			{
				p.SetState(178)
				p.Match(PromQLParserRIGHT_BRACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case PromQLParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Match(PromQLParserLEFT_BRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.LabelMatcherList()
		}
		{
			p.SetState(183)
			p.Match(PromQLParserRIGHT_BRACE)
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
	p.RuleIndex = PromQLParserRULE_labelMatcher
	return p
}

func InitEmptyLabelMatcherContext(p *LabelMatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_labelMatcher
}

func (*LabelMatcherContext) IsLabelMatcherContext() {}

func NewLabelMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherContext {
	var p = new(LabelMatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_labelMatcher

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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLabelMatcher(s)
	}
}

func (p *PromQLParser) LabelMatcher() (localctx ILabelMatcherContext) {
	localctx = NewLabelMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromQLParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.LabelName()
	}
	{
		p.SetState(188)
		p.LabelMatcherOperator()
	}
	{
		p.SetState(189)
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
	p.RuleIndex = PromQLParserRULE_labelMatcherOperator
	return p
}

func InitEmptyLabelMatcherOperatorContext(p *LabelMatcherOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_labelMatcherOperator
}

func (*LabelMatcherOperatorContext) IsLabelMatcherOperatorContext() {}

func NewLabelMatcherOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherOperatorContext {
	var p = new(LabelMatcherOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_labelMatcherOperator

	return p
}

func (s *LabelMatcherOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(PromQLParserEQ, 0)
}

func (s *LabelMatcherOperatorContext) NE() antlr.TerminalNode {
	return s.GetToken(PromQLParserNE, 0)
}

func (s *LabelMatcherOperatorContext) RE() antlr.TerminalNode {
	return s.GetToken(PromQLParserRE, 0)
}

func (s *LabelMatcherOperatorContext) NRE() antlr.TerminalNode {
	return s.GetToken(PromQLParserNRE, 0)
}

func (s *LabelMatcherOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLabelMatcherOperator(s)
	}
}

func (s *LabelMatcherOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLabelMatcherOperator(s)
	}
}

func (p *PromQLParser) LabelMatcherOperator() (localctx ILabelMatcherOperatorContext) {
	localctx = NewLabelMatcherOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PromQLParserRULE_labelMatcherOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6373376) != 0) {
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
	p.RuleIndex = PromQLParserRULE_labelMatcherList
	return p
}

func InitEmptyLabelMatcherListContext(p *LabelMatcherListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_labelMatcherList
}

func (*LabelMatcherListContext) IsLabelMatcherListContext() {}

func NewLabelMatcherListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherListContext {
	var p = new(LabelMatcherListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_labelMatcherList

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
	return s.GetTokens(PromQLParserCOMMA)
}

func (s *LabelMatcherListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLParserCOMMA, i)
}

func (s *LabelMatcherListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLabelMatcherList(s)
	}
}

func (s *LabelMatcherListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLabelMatcherList(s)
	}
}

func (p *PromQLParser) LabelMatcherList() (localctx ILabelMatcherListContext) {
	localctx = NewLabelMatcherListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PromQLParserRULE_labelMatcherList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.LabelMatcher()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(194)
				p.Match(PromQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(195)
				p.LabelMatcher()
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLParserCOMMA {
		{
			p.SetState(201)
			p.Match(PromQLParserCOMMA)
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
	p.RuleIndex = PromQLParserRULE_matrixSelector
	return p
}

func InitEmptyMatrixSelectorContext(p *MatrixSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_matrixSelector
}

func (*MatrixSelectorContext) IsMatrixSelectorContext() {}

func NewMatrixSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatrixSelectorContext {
	var p = new(MatrixSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_matrixSelector

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
	return s.GetToken(PromQLParserTIME_RANGE, 0)
}

func (s *MatrixSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrixSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatrixSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterMatrixSelector(s)
	}
}

func (s *MatrixSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitMatrixSelector(s)
	}
}

func (p *PromQLParser) MatrixSelector() (localctx IMatrixSelectorContext) {
	localctx = NewMatrixSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PromQLParserRULE_matrixSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.InstantSelector()
	}
	{
		p.SetState(205)
		p.Match(PromQLParserTIME_RANGE)
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
	p.RuleIndex = PromQLParserRULE_offset
	return p
}

func InitEmptyOffsetContext(p *OffsetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_offset
}

func (*OffsetContext) IsOffsetContext() {}

func NewOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetContext {
	var p = new(OffsetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_offset

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
	return s.GetToken(PromQLParserOFFSET, 0)
}

func (s *OffsetContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLParserDURATION, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterOffset(s)
	}
}

func (s *OffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitOffset(s)
	}
}

func (p *PromQLParser) Offset() (localctx IOffsetContext) {
	localctx = NewOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PromQLParserRULE_offset)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.InstantSelector()
		}
		{
			p.SetState(208)
			p.Match(PromQLParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(PromQLParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.MatrixSelector()
		}
		{
			p.SetState(212)
			p.Match(PromQLParserOFFSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(PromQLParserDURATION)
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
	p.RuleIndex = PromQLParserRULE_function_
	return p
}

func InitEmptyFunction_Context(p *Function_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_function_
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLParserFUNCTION, 0)
}

func (s *Function_Context) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserLEFT_PAREN, 0)
}

func (s *Function_Context) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLParserCOMMA)
}

func (s *Function_Context) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLParserCOMMA, i)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *PromQLParser) Function_() (localctx IFunction_Context) {
	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PromQLParserRULE_function_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(PromQLParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(PromQLParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&74777528107134) != 0 {
		{
			p.SetState(219)
			p.Parameter()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLParserCOMMA {
			{
				p.SetState(220)
				p.Match(PromQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(221)
				p.Parameter()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(229)
		p.Match(PromQLParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_parameter

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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *PromQLParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PromQLParserRULE_parameter)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
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
	p.RuleIndex = PromQLParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserLEFT_PAREN, 0)
}

func (s *ParameterListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLParserCOMMA)
}

func (s *ParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLParserCOMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *PromQLParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PromQLParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(PromQLParserLEFT_PAREN)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&74777528107134) != 0 {
		{
			p.SetState(236)
			p.Parameter()
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLParserCOMMA {
			{
				p.SetState(237)
				p.Match(PromQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(238)
				p.Parameter()
			}

			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(246)
		p.Match(PromQLParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLParserAGGREGATION_OPERATOR, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (p *PromQLParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PromQLParserRULE_aggregation)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(PromQLParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.ParameterList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(PromQLParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLParserBY:
			{
				p.SetState(251)
				p.By()
			}

		case PromQLParserWITHOUT:
			{
				p.SetState(252)
				p.Without()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(255)
			p.ParameterList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(257)
			p.Match(PromQLParserAGGREGATION_OPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.ParameterList()
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLParserBY:
			{
				p.SetState(259)
				p.By()
			}

		case PromQLParserWITHOUT:
			{
				p.SetState(260)
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
	p.RuleIndex = PromQLParserRULE_by
	return p
}

func InitEmptyByContext(p *ByContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_by
}

func (*ByContext) IsByContext() {}

func NewByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByContext {
	var p = new(ByContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_by

	return p
}

func (s *ByContext) GetParser() antlr.Parser { return s.parser }

func (s *ByContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLParserBY, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterBy(s)
	}
}

func (s *ByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitBy(s)
	}
}

func (p *PromQLParser) By() (localctx IByContext) {
	localctx = NewByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PromQLParserRULE_by)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.Match(PromQLParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
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
	p.RuleIndex = PromQLParserRULE_without
	return p
}

func InitEmptyWithoutContext(p *WithoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_without
}

func (*WithoutContext) IsWithoutContext() {}

func NewWithoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithoutContext {
	var p = new(WithoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_without

	return p
}

func (s *WithoutContext) GetParser() antlr.Parser { return s.parser }

func (s *WithoutContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLParserWITHOUT, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterWithout(s)
	}
}

func (s *WithoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitWithout(s)
	}
}

func (p *PromQLParser) Without() (localctx IWithoutContext) {
	localctx = NewWithoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PromQLParserRULE_without)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(PromQLParserWITHOUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
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
	p.RuleIndex = PromQLParserRULE_grouping
	return p
}

func InitEmptyGroupingContext(p *GroupingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_grouping
}

func (*GroupingContext) IsGroupingContext() {}

func NewGroupingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupingContext {
	var p = new(GroupingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_grouping

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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (p *PromQLParser) Grouping() (localctx IGroupingContext) {
	localctx = NewGroupingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PromQLParserRULE_grouping)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLParserON:
		{
			p.SetState(271)
			p.On_()
		}

	case PromQLParserIGNORING:
		{
			p.SetState(272)
			p.Ignoring()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case PromQLParserGROUP_LEFT:
		{
			p.SetState(275)
			p.GroupLeft()
		}

	case PromQLParserGROUP_RIGHT:
		{
			p.SetState(276)
			p.GroupRight()
		}

	case PromQLParserAGGREGATION_OPERATOR, PromQLParserFUNCTION, PromQLParserNUMBER, PromQLParserSTRING, PromQLParserADD, PromQLParserSUB, PromQLParserLEFT_BRACE, PromQLParserLEFT_PAREN, PromQLParserMETRIC_NAME, PromQLParserRAW_STRING:

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
	p.RuleIndex = PromQLParserRULE_on_
	return p
}

func InitEmptyOn_Context(p *On_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_on_
}

func (*On_Context) IsOn_Context() {}

func NewOn_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *On_Context {
	var p = new(On_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_on_

	return p
}

func (s *On_Context) GetParser() antlr.Parser { return s.parser }

func (s *On_Context) ON() antlr.TerminalNode {
	return s.GetToken(PromQLParserON, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterOn_(s)
	}
}

func (s *On_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitOn_(s)
	}
}

func (p *PromQLParser) On_() (localctx IOn_Context) {
	localctx = NewOn_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PromQLParserRULE_on_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(PromQLParserON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
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
	p.RuleIndex = PromQLParserRULE_ignoring
	return p
}

func InitEmptyIgnoringContext(p *IgnoringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_ignoring
}

func (*IgnoringContext) IsIgnoringContext() {}

func NewIgnoringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoringContext {
	var p = new(IgnoringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_ignoring

	return p
}

func (s *IgnoringContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoringContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLParserIGNORING, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterIgnoring(s)
	}
}

func (s *IgnoringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitIgnoring(s)
	}
}

func (p *PromQLParser) Ignoring() (localctx IIgnoringContext) {
	localctx = NewIgnoringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PromQLParserRULE_ignoring)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(PromQLParserIGNORING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
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
	p.RuleIndex = PromQLParserRULE_groupLeft
	return p
}

func InitEmptyGroupLeftContext(p *GroupLeftContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_groupLeft
}

func (*GroupLeftContext) IsGroupLeftContext() {}

func NewGroupLeftContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupLeftContext {
	var p = new(GroupLeftContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_groupLeft

	return p
}

func (s *GroupLeftContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupLeftContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLParserGROUP_LEFT, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterGroupLeft(s)
	}
}

func (s *GroupLeftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitGroupLeft(s)
	}
}

func (p *PromQLParser) GroupLeft() (localctx IGroupLeftContext) {
	localctx = NewGroupLeftContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PromQLParserRULE_groupLeft)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(PromQLParserGROUP_LEFT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(286)
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
	p.RuleIndex = PromQLParserRULE_groupRight
	return p
}

func InitEmptyGroupRightContext(p *GroupRightContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_groupRight
}

func (*GroupRightContext) IsGroupRightContext() {}

func NewGroupRightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupRightContext {
	var p = new(GroupRightContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_groupRight

	return p
}

func (s *GroupRightContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupRightContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLParserGROUP_RIGHT, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterGroupRight(s)
	}
}

func (s *GroupRightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitGroupRight(s)
	}
}

func (p *PromQLParser) GroupRight() (localctx IGroupRightContext) {
	localctx = NewGroupRightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PromQLParserRULE_groupRight)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(PromQLParserGROUP_RIGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(290)
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
	p.RuleIndex = PromQLParserRULE_labelName
	return p
}

func InitEmptyLabelNameContext(p *LabelNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_labelName
}

func (*LabelNameContext) IsLabelNameContext() {}

func NewLabelNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameContext {
	var p = new(LabelNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_labelName

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
	return s.GetToken(PromQLParserMETRIC_NAME, 0)
}

func (s *LabelNameContext) LABEL_NAME() antlr.TerminalNode {
	return s.GetToken(PromQLParserLABEL_NAME, 0)
}

func (s *LabelNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLabelName(s)
	}
}

func (s *LabelNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLabelName(s)
	}
}

func (p *PromQLParser) LabelName() (localctx ILabelNameContext) {
	localctx = NewLabelNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PromQLParserRULE_labelName)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLParserAGGREGATION_OPERATOR, PromQLParserFUNCTION, PromQLParserAND, PromQLParserOR, PromQLParserUNLESS, PromQLParserBY, PromQLParserWITHOUT, PromQLParserON, PromQLParserIGNORING, PromQLParserGROUP_LEFT, PromQLParserGROUP_RIGHT, PromQLParserOFFSET, PromQLParserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Keyword()
		}

	case PromQLParserMETRIC_NAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(PromQLParserMETRIC_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLParserLABEL_NAME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(PromQLParserLABEL_NAME)
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
	p.RuleIndex = PromQLParserRULE_labelNameList
	return p
}

func InitEmptyLabelNameListContext(p *LabelNameListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_labelNameList
}

func (*LabelNameListContext) IsLabelNameListContext() {}

func NewLabelNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelNameListContext {
	var p = new(LabelNameListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_labelNameList

	return p
}

func (s *LabelNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelNameListContext) LEFT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserLEFT_PAREN, 0)
}

func (s *LabelNameListContext) RIGHT_PAREN() antlr.TerminalNode {
	return s.GetToken(PromQLParserRIGHT_PAREN, 0)
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
	return s.GetTokens(PromQLParserCOMMA)
}

func (s *LabelNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PromQLParserCOMMA, i)
}

func (s *LabelNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLabelNameList(s)
	}
}

func (s *LabelNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLabelNameList(s)
	}
}

func (p *PromQLParser) LabelNameList() (localctx ILabelNameListContext) {
	localctx = NewLabelNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PromQLParserRULE_labelNameList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(PromQLParserLEFT_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13196278642694) != 0 {
		{
			p.SetState(299)
			p.LabelName()
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PromQLParserCOMMA {
			{
				p.SetState(300)
				p.Match(PromQLParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(301)
				p.LabelName()
			}

			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(309)
		p.Match(PromQLParserRIGHT_PAREN)
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
	p.RuleIndex = PromQLParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) AND() antlr.TerminalNode {
	return s.GetToken(PromQLParserAND, 0)
}

func (s *KeywordContext) OR() antlr.TerminalNode {
	return s.GetToken(PromQLParserOR, 0)
}

func (s *KeywordContext) UNLESS() antlr.TerminalNode {
	return s.GetToken(PromQLParserUNLESS, 0)
}

func (s *KeywordContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLParserBY, 0)
}

func (s *KeywordContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLParserWITHOUT, 0)
}

func (s *KeywordContext) ON() antlr.TerminalNode {
	return s.GetToken(PromQLParserON, 0)
}

func (s *KeywordContext) IGNORING() antlr.TerminalNode {
	return s.GetToken(PromQLParserIGNORING, 0)
}

func (s *KeywordContext) GROUP_LEFT() antlr.TerminalNode {
	return s.GetToken(PromQLParserGROUP_LEFT, 0)
}

func (s *KeywordContext) GROUP_RIGHT() antlr.TerminalNode {
	return s.GetToken(PromQLParserGROUP_RIGHT, 0)
}

func (s *KeywordContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(PromQLParserOFFSET, 0)
}

func (s *KeywordContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PromQLParserBOOL, 0)
}

func (s *KeywordContext) AGGREGATION_OPERATOR() antlr.TerminalNode {
	return s.GetToken(PromQLParserAGGREGATION_OPERATOR, 0)
}

func (s *KeywordContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PromQLParserFUNCTION, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *PromQLParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PromQLParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2139109382) != 0) {
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
	p.RuleIndex = PromQLParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLParserNUMBER, 0)
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
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *PromQLParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PromQLParserRULE_literal)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Match(PromQLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLParserSTRING, PromQLParserRAW_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
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
	p.RuleIndex = PromQLParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLParserSTRING, 0)
}

func (s *StringContext) RAW_STRING() antlr.TerminalNode {
	return s.GetToken(PromQLParserRAW_STRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *PromQLParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PromQLParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLParserSTRING || _la == PromQLParserRAW_STRING) {
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

func (p *PromQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *VectorOperationContext = nil
		if localctx != nil {
			t = localctx.(*VectorOperationContext)
		}
		return p.VectorOperation_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PromQLParser) VectorOperation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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

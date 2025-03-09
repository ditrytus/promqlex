package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

type FunctionsSet struct {
	Functions                    map[string]struct{}
	AggregationOperators         map[string]struct{}
	functionTokenType            int
	aggregationOperatorTokenType int
}

func NewFunctionsSet(
	functions, aggregationOperators []string,
	functionTokenType, aggregationOperatorTokenType int,
) FunctionsSet {
	functionsSet := FunctionsSet{
		Functions:                    make(map[string]struct{}),
		AggregationOperators:         make(map[string]struct{}),
		functionTokenType:            functionTokenType,
		aggregationOperatorTokenType: aggregationOperatorTokenType,
	}
	for _, function := range functions {
		functionsSet.Functions[function] = struct{}{}
	}
	for _, aggregationOperator := range aggregationOperators {
		functionsSet.AggregationOperators[aggregationOperator] = struct{}{}
	}
	return functionsSet
}

func (f FunctionsSet) GetTokenType(text string) (int, bool) {
	if _, ok := f.Functions[text]; ok {
		return f.functionTokenType, true
	}
	if _, ok := f.AggregationOperators[strings.ToLower(text)]; ok {
		return f.aggregationOperatorTokenType, true
	}
	return antlr.TokenInvalidType, false
}

var prometheusFunctions = []string{
	"abs",
	"absent",
	"absent_over_time",
	"ceil",
	"changes",
	"clamp",
	"clamp_max",
	"clamp_min",
	"day_of_month",
	"day_of_week",
	"day_of_year",
	"days_in_month",
	"delta",
	"deriv",
	"exp",
	"floor",
	"histogram_count",
	"histogram_sum",
	"histogram_fraction",
	"histogram_quantile",
	"holt_winters",
	"hour",
	"idelta",
	"increase",
	"irate",
	"label_join",
	"label_replace",
	"ln",
	"log2",
	"log10",
	"minute",
	"month",
	"predict_linear",
	"rate",
	"resets",
	"round",
	"scalar",
	"sgn",
	"sort",
	"sort_desc",
	"sqrt",
	"time",
	"timestamp",
	"vector",
	"year",
	"avg_over_time",
	"min_over_time",
	"max_over_time",
	"sum_over_time",
	"count_over_time",
	"quantile_over_time",
	"stddev_over_time",
	"stdvar_over_time",
	"last_over_time",
	"present_over_time",
	"acos",
	"acosh",
	"asin",
	"asinh",
	"atan",
	"atanh",
	"cos",
	"cosh",
	"sin",
	"sinh",
	"tan",
	"tanh",
	"deg",
	"pi",
	"rad",
}

var prometheusAggregationOperators = []string{
	"sum",
	"min",
	"max",
	"avg",
	"group",
	"stddev",
	"stdvar",
	"count",
	"count_values",
	"bottomk",
	"topk",
	"quantile",
}

func NewPrometheusFunctionSet(functionTokenType, aggregationOperatorTokenType int) FunctionsSet {
	return NewFunctionsSet(
		prometheusFunctions,
		prometheusAggregationOperators,
		functionTokenType,
		aggregationOperatorTokenType,
	)
}

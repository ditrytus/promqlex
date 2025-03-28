package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
	math "math"
	"strconv"
	"time"
)

type replaceableNode interface {
	GetStart() antlr.Token
	GetStop() antlr.Token
}

type ConstNumExprEvaluator struct {
	promqlex.BasePromQLExParserListener
	rewriter           *antlr.TokenStreamRewriter
	constNumExprStack  Stack[float64]
	unixTimestamp      float64
	timeInstantLiteral time.Time
}

func NewConstNumExprEvaluator(tokenStream *antlr.TokenStreamRewriter) *ConstNumExprEvaluator {
	return &ConstNumExprEvaluator{
		rewriter: tokenStream,
	}
}

var _ promqlex.PromQLExParserListener = &ConstNumExprEvaluator{}

func (t *ConstNumExprEvaluator) setFloat64FieldFromToken(parser antlr.Parser, field **float64, token antlr.Token) {
	if token != nil {
		val, err := strconv.ParseFloat(token.GetText(), 64)
		if err != nil {
			parser.NotifyErrorListeners(err.Error(), token, nil)
		}
		*field = &val
	}
}

func (t *ConstNumExprEvaluator) ExitTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {
	txt := c.RFC_3339_TIMESTAMP().GetText()
	timestamp, err := time.Parse(time.RFC3339, txt)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.RFC_3339_TIMESTAMP().GetSymbol(), nil)
	}
	t.timeInstantLiteral = timestamp
}

func (t *ConstNumExprEvaluator) ExitTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
	txt := c.NUMBER().GetText()
	unixTimestamp, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.NUMBER().GetSymbol(), nil)
	}
	timeVal := unixFloatToTime(unixTimestamp)
	t.timeInstantLiteral = timeVal
}

func (t *ConstNumExprEvaluator) ExitConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {
	b := t.constNumExprStack.MustPop()
	a := t.constNumExprStack.MustPop()
	result := math.Pow(a, b)
	t.constNumExprStack.Push(result)
}

func (t *ConstNumExprEvaluator) ExitConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {
	if c.UnaryOp().ADD() != nil {
		return
	}
	t.constNumExprStack.Push(-t.constNumExprStack.MustPop())
}

func (t *ConstNumExprEvaluator) ExitConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {
	b := t.constNumExprStack.MustPop()
	a := t.constNumExprStack.MustPop()
	var result float64
	switch {
	case c.MultOp().MULT() != nil:
		result = a * b
	case c.MultOp().DIV() != nil:
		if b == 0 {
			c.GetParser().NotifyErrorListeners("division by zero", c.MultOp().DIV().GetSymbol(), nil)
		}
		result = a / b
	}
	t.constNumExprStack.Push(result)
}

func (t *ConstNumExprEvaluator) ExitConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {
	b := t.constNumExprStack.MustPop()
	a := t.constNumExprStack.MustPop()
	var result float64
	switch {
	case c.AddOp().ADD() != nil:
		result = a + b
	case c.AddOp().SUB() != nil:
		result = a - b
	}
	t.constNumExprStack.Push(result)
}

func (t *ConstNumExprEvaluator) ExitNumLit_Number(c *promqlex.NumLit_NumberContext) {
	txt := c.NUMBER().GetText()
	num, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.NUMBER().GetSymbol(), nil)
	}
	t.constNumExprStack.Push(num)
}

func (t *ConstNumExprEvaluator) ExitNumLit_Duration(c *promqlex.NumLit_DurationContext) {
	txt := c.DURATION().GetText()
	duration, err := ParseDuration(txt)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.DURATION().GetSymbol(), nil)
	}
	t.constNumExprStack.Push(duration.Seconds())
}

func (t *ConstNumExprEvaluator) ExitNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {
	t.constNumExprStack.Push(timeToUnixFloat(t.timeInstantLiteral))
}
func (t *ConstNumExprEvaluator) ExitDuration(c *promqlex.DurationContext) {
	t.replaceWithConstNumExprAsDuration(c)
}

func (t *ConstNumExprEvaluator) ExitVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) ExitLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) ExitAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) replaceWithConstNumExprAsFloat(c replaceableNode) {
	num := t.constNumExprStack.MustPop()
	txt := strconv.FormatFloat(num, 'g', 12, 64)
	t.rewriter.ReplaceTokenDefault(c.GetStart(), c.GetStop(), txt)
}

func (t *ConstNumExprEvaluator) replaceWithConstNumExprAsDuration(c *promqlex.DurationContext) {
	num := t.constNumExprStack.MustPop()
	txt := FormatDuration(time.Duration(float64(time.Second) * num))
	t.rewriter.ReplaceTokenDefault(c.GetStart(), c.GetStop(), txt)
}

func unixFloatToTime(unixTimestamp float64) time.Time {
	sec, nano := secondsToSecAndNano(unixTimestamp)
	return time.Unix(int64(sec), int64(nano))
}

func secondsToSecAndNano(seconds float64) (int, int) {
	sec := int(seconds)
	nano := int((seconds - float64(sec)) * 1e9)
	return sec, nano
}

func timeToUnixFloat(t time.Time) float64 {
	return float64(t.Unix()) + float64(t.Nanosecond())/1e9
}

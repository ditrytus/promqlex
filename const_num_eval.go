package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
	math "math"
	"strconv"
	"time"
)

type ConstNumExprEvaluator struct {
	rewriter           *antlr.TokenStreamRewriter
	constNumExprStack  Stack[float64]
	unixTimestamp      float64
	timeInstantLiteral time.Time
}

func NewConstNumExporEvaluator(tokenStream *antlr.TokenStreamRewriter) *ConstNumExprEvaluator {
	return &ConstNumExprEvaluator{
		rewriter: tokenStream,
	}
}

var _ promqlex.PromQLExParserListener = &ConstNumExprEvaluator{}

func (t *ConstNumExprEvaluator) VisitTerminal(node antlr.TerminalNode) {}

func (t *ConstNumExprEvaluator) VisitErrorNode(node antlr.ErrorNode) {}

func (t *ConstNumExprEvaluator) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (t *ConstNumExprEvaluator) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (t *ConstNumExprEvaluator) EnterPromqlx(c *promqlex.PromqlxContext) {}

func (t *ConstNumExprEvaluator) EnterState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (t *ConstNumExprEvaluator) EnterState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (t *ConstNumExprEvaluator) EnterState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (t *ConstNumExprEvaluator) EnterState_VectorOperation(c *promqlex.State_VectorOperationContext) {
}

func (t *ConstNumExprEvaluator) EnterAlias_def(c *promqlex.Alias_defContext) {}

func (t *ConstNumExprEvaluator) EnterAlias_call(c *promqlex.Alias_callContext) {}

func (t *ConstNumExprEvaluator) EnterMacro_def(c *promqlex.Macro_defContext) {}

func (t *ConstNumExprEvaluator) EnterMacro_call(c *promqlex.Macro_callContext) {}

func (t *ConstNumExprEvaluator) EnterArgs_decl(c *promqlex.Args_declContext) {}

func (t *ConstNumExprEvaluator) EnterArg_name(c *promqlex.Arg_nameContext) {}

func (t *ConstNumExprEvaluator) EnterStatement_block(c *promqlex.Statement_blockContext) {}

func (t *ConstNumExprEvaluator) EnterArg_list(c *promqlex.Arg_listContext) {}

func (t *ConstNumExprEvaluator) EnterIf_statement(c *promqlex.If_statementContext) {}

func (t *ConstNumExprEvaluator) EnterCondition(c *promqlex.ConditionContext) {}

func (t *ConstNumExprEvaluator) EnterCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {
}

func (t *ConstNumExprEvaluator) EnterTrueConst(c *promqlex.TrueConstContext) {}

func (t *ConstNumExprEvaluator) EnterFalseConst(c *promqlex.FalseConstContext) {}

func (t *ConstNumExprEvaluator) EnterTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {
}

func (t *ConstNumExprEvaluator) EnterTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
}

func (t *ConstNumExprEvaluator) setFloat64FieldFromToken(parser antlr.Parser, field **float64, token antlr.Token) {
	if token != nil {
		val, err := strconv.ParseFloat(token.GetText(), 64)
		if err != nil {
			parser.NotifyErrorListeners(err.Error(), token, nil)
		}
		*field = &val
	}
}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (t *ConstNumExprEvaluator) EnterConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {
}

func (t *ConstNumExprEvaluator) EnterNumLit_Number(c *promqlex.NumLit_NumberContext) {}

func (t *ConstNumExprEvaluator) EnterNumLit_Duration(c *promqlex.NumLit_DurationContext) {}

func (t *ConstNumExprEvaluator) EnterNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {
}

func (t *ConstNumExprEvaluator) EnterNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {}

func (t *ConstNumExprEvaluator) EnterDuration(c *promqlex.DurationContext) {}

func (t *ConstNumExprEvaluator) EnterTime_range(c *promqlex.Time_rangeContext) {}

func (t *ConstNumExprEvaluator) EnterSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_At(c *promqlex.VecOp_AtContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (t *ConstNumExprEvaluator) EnterVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {}

func (t *ConstNumExprEvaluator) EnterSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (t *ConstNumExprEvaluator) EnterOffsetOp(c *promqlex.OffsetOpContext) {}

func (t *ConstNumExprEvaluator) EnterMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (t *ConstNumExprEvaluator) EnterOffset(c *promqlex.OffsetContext) {}

func (t *ConstNumExprEvaluator) EnterLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {}

func (t *ConstNumExprEvaluator) EnterLit_String(c *promqlex.Lit_StringContext) {}

func (t *ConstNumExprEvaluator) EnterInstantSelector(c *promqlex.InstantSelectorContext) {}

func (t *ConstNumExprEvaluator) EnterLabelName(c *promqlex.LabelNameContext) {}

func (t *ConstNumExprEvaluator) EnterMetric_name(c *promqlex.Metric_nameContext) {}

func (t *ConstNumExprEvaluator) EnterAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {
}

func (t *ConstNumExprEvaluator) EnterAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (t *ConstNumExprEvaluator) EnterAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (t *ConstNumExprEvaluator) EnterExpression(c *promqlex.ExpressionContext) {}

func (t *ConstNumExprEvaluator) EnterUnaryOp(c *promqlex.UnaryOpContext) {}

func (t *ConstNumExprEvaluator) EnterPowOp(c *promqlex.PowOpContext) {}

func (t *ConstNumExprEvaluator) EnterMultOp(c *promqlex.MultOpContext) {}

func (t *ConstNumExprEvaluator) EnterAddOp(c *promqlex.AddOpContext) {}

func (t *ConstNumExprEvaluator) EnterCompareOp(c *promqlex.CompareOpContext) {}

func (t *ConstNumExprEvaluator) EnterAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (t *ConstNumExprEvaluator) EnterOrOp(c *promqlex.OrOpContext) {}

func (t *ConstNumExprEvaluator) EnterVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (t *ConstNumExprEvaluator) EnterVector(c *promqlex.VectorContext) {}

func (t *ConstNumExprEvaluator) EnterParens(c *promqlex.ParensContext) {}

func (t *ConstNumExprEvaluator) EnterLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (t *ConstNumExprEvaluator) EnterLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (t *ConstNumExprEvaluator) EnterLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (t *ConstNumExprEvaluator) EnterFunction_(c *promqlex.Function_Context) {}

func (t *ConstNumExprEvaluator) EnterParameter(c *promqlex.ParameterContext) {}

func (t *ConstNumExprEvaluator) EnterParameterList(c *promqlex.ParameterListContext) {}

func (t *ConstNumExprEvaluator) EnterAggregation(c *promqlex.AggregationContext) {}

func (t *ConstNumExprEvaluator) EnterBy(c *promqlex.ByContext) {}

func (t *ConstNumExprEvaluator) EnterWithout(c *promqlex.WithoutContext) {}

func (t *ConstNumExprEvaluator) EnterGrouping(c *promqlex.GroupingContext) {}

func (t *ConstNumExprEvaluator) EnterOn_(c *promqlex.On_Context) {}

func (t *ConstNumExprEvaluator) EnterIgnoring(c *promqlex.IgnoringContext) {}

func (t *ConstNumExprEvaluator) EnterGroupLeft(c *promqlex.GroupLeftContext) {}

func (t *ConstNumExprEvaluator) EnterGroupRight(c *promqlex.GroupRightContext) {}

func (t *ConstNumExprEvaluator) EnterLabelNameList(c *promqlex.LabelNameListContext) {}

func (t *ConstNumExprEvaluator) EnterKeyword(c *promqlex.KeywordContext) {}

func (t *ConstNumExprEvaluator) EnterString(c *promqlex.StringContext) {}

func (t *ConstNumExprEvaluator) ExitPromqlx(c *promqlex.PromqlxContext) {}

func (t *ConstNumExprEvaluator) ExitState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (t *ConstNumExprEvaluator) ExitState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (t *ConstNumExprEvaluator) ExitState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (t *ConstNumExprEvaluator) ExitState_VectorOperation(c *promqlex.State_VectorOperationContext) {}

func (t *ConstNumExprEvaluator) ExitAlias_def(c *promqlex.Alias_defContext) {}

func (t *ConstNumExprEvaluator) ExitAlias_call(c *promqlex.Alias_callContext) {}

func (t *ConstNumExprEvaluator) ExitMacro_def(c *promqlex.Macro_defContext) {}

func (t *ConstNumExprEvaluator) ExitMacro_call(c *promqlex.Macro_callContext) {}

func (t *ConstNumExprEvaluator) ExitArgs_decl(c *promqlex.Args_declContext) {}

func (t *ConstNumExprEvaluator) ExitArg_name(c *promqlex.Arg_nameContext) {}

func (t *ConstNumExprEvaluator) ExitStatement_block(c *promqlex.Statement_blockContext) {}

func (t *ConstNumExprEvaluator) ExitArg_list(c *promqlex.Arg_listContext) {}

func (t *ConstNumExprEvaluator) ExitIf_statement(c *promqlex.If_statementContext) {}

func (t *ConstNumExprEvaluator) ExitCondition(c *promqlex.ConditionContext) {}

func (t *ConstNumExprEvaluator) ExitCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {
}

func (t *ConstNumExprEvaluator) ExitTrueConst(c *promqlex.TrueConstContext) {}

func (t *ConstNumExprEvaluator) ExitFalseConst(c *promqlex.FalseConstContext) {}

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

func (t *ConstNumExprEvaluator) ExitConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (t *ConstNumExprEvaluator) ExitConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {
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

func (t *ConstNumExprEvaluator) ExitNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {
	// TODO: implement
	panic("alias call as numeric literal is not yet supported")
}

func (t *ConstNumExprEvaluator) ExitDuration(c *promqlex.DurationContext) {
	t.replaceWithConstNumExprAsDuration(c)
}

func (t *ConstNumExprEvaluator) ExitTime_range(c *promqlex.Time_rangeContext) {}

func (t *ConstNumExprEvaluator) ExitSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_At(c *promqlex.VecOp_AtContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (t *ConstNumExprEvaluator) ExitVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) ExitSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (t *ConstNumExprEvaluator) ExitOffsetOp(c *promqlex.OffsetOpContext) {}

func (t *ConstNumExprEvaluator) ExitMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (t *ConstNumExprEvaluator) ExitOffset(c *promqlex.OffsetContext) {}

func (t *ConstNumExprEvaluator) ExitLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) ExitLit_String(c *promqlex.Lit_StringContext) {}

func (t *ConstNumExprEvaluator) ExitInstantSelector(c *promqlex.InstantSelectorContext) {}

func (t *ConstNumExprEvaluator) ExitLabelName(c *promqlex.LabelNameContext) {}

func (t *ConstNumExprEvaluator) ExitMetric_name(c *promqlex.Metric_nameContext) {}

func (t *ConstNumExprEvaluator) ExitAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *ConstNumExprEvaluator) ExitAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (t *ConstNumExprEvaluator) ExitAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (t *ConstNumExprEvaluator) ExitExpression(c *promqlex.ExpressionContext) {}

func (t *ConstNumExprEvaluator) ExitUnaryOp(c *promqlex.UnaryOpContext) {}

func (t *ConstNumExprEvaluator) ExitPowOp(c *promqlex.PowOpContext) {}

func (t *ConstNumExprEvaluator) ExitMultOp(c *promqlex.MultOpContext) {}

func (t *ConstNumExprEvaluator) ExitAddOp(c *promqlex.AddOpContext) {}

func (t *ConstNumExprEvaluator) ExitCompareOp(c *promqlex.CompareOpContext) {}

func (t *ConstNumExprEvaluator) ExitAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (t *ConstNumExprEvaluator) ExitOrOp(c *promqlex.OrOpContext) {}

func (t *ConstNumExprEvaluator) ExitVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (t *ConstNumExprEvaluator) ExitVector(c *promqlex.VectorContext) {}

func (t *ConstNumExprEvaluator) ExitParens(c *promqlex.ParensContext) {}

func (t *ConstNumExprEvaluator) ExitLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (t *ConstNumExprEvaluator) ExitLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (t *ConstNumExprEvaluator) ExitLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (t *ConstNumExprEvaluator) ExitFunction_(c *promqlex.Function_Context) {}

func (t *ConstNumExprEvaluator) ExitParameter(c *promqlex.ParameterContext) {}

func (t *ConstNumExprEvaluator) ExitParameterList(c *promqlex.ParameterListContext) {}

func (t *ConstNumExprEvaluator) ExitAggregation(c *promqlex.AggregationContext) {}

func (t *ConstNumExprEvaluator) ExitBy(c *promqlex.ByContext) {}

func (t *ConstNumExprEvaluator) ExitWithout(c *promqlex.WithoutContext) {}

func (t *ConstNumExprEvaluator) ExitGrouping(c *promqlex.GroupingContext) {}

func (t *ConstNumExprEvaluator) ExitOn_(c *promqlex.On_Context) {}

func (t *ConstNumExprEvaluator) ExitIgnoring(c *promqlex.IgnoringContext) {}

func (t *ConstNumExprEvaluator) ExitGroupLeft(c *promqlex.GroupLeftContext) {}

func (t *ConstNumExprEvaluator) ExitGroupRight(c *promqlex.GroupRightContext) {}

func (t *ConstNumExprEvaluator) ExitLabelNameList(c *promqlex.LabelNameListContext) {}

func (t *ConstNumExprEvaluator) ExitKeyword(c *promqlex.KeywordContext) {}

func (t *ConstNumExprEvaluator) ExitString(c *promqlex.StringContext) {}

type replaceableNode interface {
	GetStart() antlr.Token
	GetStop() antlr.Token
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

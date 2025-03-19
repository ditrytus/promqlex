package promqlex

import (
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
	math "math"
	"strconv"
	"time"
)

type IsoDateTimeElements struct {
	Year   *float64
	Month  *float64
	Day    *float64
	Hour   *float64
	Minute *float64
	Second *float64
}

func (i *IsoDateTimeElements) Elements() []*float64 {
	return []*float64{i.Year, i.Month, i.Day, i.Hour, i.Minute, i.Second}
}

func (i *IsoDateTimeElements) Time() (time.Time, error) {
	if i.Year == nil {
		return time.Time{}, errors.New("missing year")
	}
	sec, nano := secondsToSecAndNano(coalesce(i.Second, 0))
	date := time.Date(
		int(*i.Year),
		time.Month(coalesce(i.Month, 1)),
		int(coalesce(i.Day, 1)),
		int(coalesce(i.Hour, 0)),
		int(coalesce(i.Minute, 0)),
		sec,
		nano,
		time.UTC,
	)
	return date, nil
}

func coalesce[T any](a *T, b T) T {
	if a == nil {
		return b
	}
	return *a
}

type Transpiler struct {
	rewriter            *antlr.TokenStreamRewriter
	constNumExprStack   Stack[float64]
	unixTimestamp       float64
	timeInstantLiteral  time.Time
	isoDateTimeElements IsoDateTimeElements
}

func NewTranspiler(tokenStream *antlr.TokenStreamRewriter) *Transpiler {
	return &Transpiler{
		rewriter: tokenStream,
	}
}

var _ promqlex.PromQLExParserListener = &Transpiler{}

func (t *Transpiler) VisitTerminal(node antlr.TerminalNode) {}

func (t *Transpiler) VisitErrorNode(node antlr.ErrorNode) {}

func (t *Transpiler) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (t *Transpiler) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (t *Transpiler) EnterPromqlx(c *promqlex.PromqlxContext) {}

func (t *Transpiler) EnterState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (t *Transpiler) EnterState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (t *Transpiler) EnterState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (t *Transpiler) EnterState_VectorOperation(c *promqlex.State_VectorOperationContext) {}

func (t *Transpiler) EnterAlias_def(c *promqlex.Alias_defContext) {}

func (t *Transpiler) EnterAlias_call(c *promqlex.Alias_callContext) {}

func (t *Transpiler) EnterMacro_def(c *promqlex.Macro_defContext) {}

func (t *Transpiler) EnterMacro_call(c *promqlex.Macro_callContext) {}

func (t *Transpiler) EnterArgs_decl(c *promqlex.Args_declContext) {}

func (t *Transpiler) EnterArg_name(c *promqlex.Arg_nameContext) {}

func (t *Transpiler) EnterStatement_block(c *promqlex.Statement_blockContext) {}

func (t *Transpiler) EnterArg_list(c *promqlex.Arg_listContext) {}

func (t *Transpiler) EnterIf_statement(c *promqlex.If_statementContext) {}

func (t *Transpiler) EnterCondition(c *promqlex.ConditionContext) {}

func (t *Transpiler) EnterCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {}

func (t *Transpiler) EnterTrueConst(c *promqlex.TrueConstContext) {}

func (t *Transpiler) EnterFalseConst(c *promqlex.FalseConstContext) {}

func (t *Transpiler) EnterTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {
}

func (t *Transpiler) EnterTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {}

func (t *Transpiler) EnterIso_date_time(c *promqlex.Iso_date_timeContext) {}

func (t *Transpiler) setFloat64FieldFromToken(parser antlr.Parser, field **float64, token antlr.Token) {
	if token != nil {
		val, err := strconv.ParseFloat(token.GetText(), 64)
		if err != nil {
			parser.NotifyErrorListeners(err.Error(), token, nil)
		}
		*field = &val
	}
}

func (t *Transpiler) EnterUnix_timestamp(c *promqlex.Unix_timestampContext) {}

func (t *Transpiler) EnterConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {}

func (t *Transpiler) EnterConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {}

func (t *Transpiler) EnterConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {}

func (t *Transpiler) EnterConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {}

func (t *Transpiler) EnterConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (t *Transpiler) EnterConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {}

func (t *Transpiler) EnterNumLit_Number(c *promqlex.NumLit_NumberContext) {}

func (t *Transpiler) EnterNumLit_Duration(c *promqlex.NumLit_DurationContext) {}

func (t *Transpiler) EnterNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {}

func (t *Transpiler) EnterNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {}

func (t *Transpiler) EnterDuration(c *promqlex.DurationContext) {}

func (t *Transpiler) EnterTime_range(c *promqlex.Time_rangeContext) {}

func (t *Transpiler) EnterSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (t *Transpiler) EnterVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (t *Transpiler) EnterVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (t *Transpiler) EnterVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (t *Transpiler) EnterVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (t *Transpiler) EnterVecOp_At(c *promqlex.VecOp_AtContext) {}

func (t *Transpiler) EnterVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (t *Transpiler) EnterVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (t *Transpiler) EnterVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (t *Transpiler) EnterVecOp_SubOp(c *promqlex.VecOp_SubOpContext) {}

func (t *Transpiler) EnterVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *Transpiler) EnterVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *Transpiler) EnterVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *Transpiler) EnterVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *Transpiler) EnterSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (t *Transpiler) EnterOffsetOp(c *promqlex.OffsetOpContext) {}

func (t *Transpiler) EnterMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (t *Transpiler) EnterOffset(c *promqlex.OffsetContext) {}

func (t *Transpiler) EnterLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {}

func (t *Transpiler) EnterLit_String(c *promqlex.Lit_StringContext) {}

func (t *Transpiler) EnterInstantSelector(c *promqlex.InstantSelectorContext) {}

func (t *Transpiler) EnterLabelName(c *promqlex.LabelNameContext) {}

func (t *Transpiler) EnterMetric_name(c *promqlex.Metric_nameContext) {}

func (t *Transpiler) EnterAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {}

func (t *Transpiler) EnterAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (t *Transpiler) EnterAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (t *Transpiler) EnterExpression(c *promqlex.ExpressionContext) {}

func (t *Transpiler) EnterUnaryOp(c *promqlex.UnaryOpContext) {}

func (t *Transpiler) EnterPowOp(c *promqlex.PowOpContext) {}

func (t *Transpiler) EnterMultOp(c *promqlex.MultOpContext) {}

func (t *Transpiler) EnterAddOp(c *promqlex.AddOpContext) {}

func (t *Transpiler) EnterCompareOp(c *promqlex.CompareOpContext) {}

func (t *Transpiler) EnterAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (t *Transpiler) EnterOrOp(c *promqlex.OrOpContext) {}

func (t *Transpiler) EnterVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (t *Transpiler) EnterVector(c *promqlex.VectorContext) {}

func (t *Transpiler) EnterParens(c *promqlex.ParensContext) {}

func (t *Transpiler) EnterLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (t *Transpiler) EnterLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (t *Transpiler) EnterLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (t *Transpiler) EnterFunction_(c *promqlex.Function_Context) {}

func (t *Transpiler) EnterParameter(c *promqlex.ParameterContext) {}

func (t *Transpiler) EnterParameterList(c *promqlex.ParameterListContext) {}

func (t *Transpiler) EnterAggregation(c *promqlex.AggregationContext) {}

func (t *Transpiler) EnterBy(c *promqlex.ByContext) {}

func (t *Transpiler) EnterWithout(c *promqlex.WithoutContext) {}

func (t *Transpiler) EnterGrouping(c *promqlex.GroupingContext) {}

func (t *Transpiler) EnterOn_(c *promqlex.On_Context) {}

func (t *Transpiler) EnterIgnoring(c *promqlex.IgnoringContext) {}

func (t *Transpiler) EnterGroupLeft(c *promqlex.GroupLeftContext) {}

func (t *Transpiler) EnterGroupRight(c *promqlex.GroupRightContext) {}

func (t *Transpiler) EnterLabelNameList(c *promqlex.LabelNameListContext) {}

func (t *Transpiler) EnterKeyword(c *promqlex.KeywordContext) {}

func (t *Transpiler) EnterString(c *promqlex.StringContext) {}

func (t *Transpiler) ExitPromqlx(c *promqlex.PromqlxContext) {}

func (t *Transpiler) ExitState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (t *Transpiler) ExitState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (t *Transpiler) ExitState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (t *Transpiler) ExitState_VectorOperation(c *promqlex.State_VectorOperationContext) {}

func (t *Transpiler) ExitAlias_def(c *promqlex.Alias_defContext) {}

func (t *Transpiler) ExitAlias_call(c *promqlex.Alias_callContext) {}

func (t *Transpiler) ExitMacro_def(c *promqlex.Macro_defContext) {}

func (t *Transpiler) ExitMacro_call(c *promqlex.Macro_callContext) {}

func (t *Transpiler) ExitArgs_decl(c *promqlex.Args_declContext) {}

func (t *Transpiler) ExitArg_name(c *promqlex.Arg_nameContext) {}

func (t *Transpiler) ExitStatement_block(c *promqlex.Statement_blockContext) {}

func (t *Transpiler) ExitArg_list(c *promqlex.Arg_listContext) {}

func (t *Transpiler) ExitIf_statement(c *promqlex.If_statementContext) {}

func (t *Transpiler) ExitCondition(c *promqlex.ConditionContext) {}

func (t *Transpiler) ExitCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {}

func (t *Transpiler) ExitTrueConst(c *promqlex.TrueConstContext) {}

func (t *Transpiler) ExitFalseConst(c *promqlex.FalseConstContext) {}

func (t *Transpiler) ExitTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {
	timestamp, err := t.isoDateTimeElements.Time()
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.GetStart(), nil)
	}
	t.timeInstantLiteral = timestamp
}

func (t *Transpiler) ExitTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
	timeVal := unixFloatToTime(t.unixTimestamp)
	t.timeInstantLiteral = timeVal
}

func (t *Transpiler) ExitIso_date_time(c *promqlex.Iso_date_timeContext) {
	isoDateTimeElements := IsoDateTimeElements{}
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Year, c.GetYear())
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Month, c.GetMonth())
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Day, c.GetDay())
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Hour, c.GetHour())
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Minute, c.GetMinutes())
	t.setFloat64FieldFromToken(c.GetParser(), &isoDateTimeElements.Second, c.GetSeconds())
	t.isoDateTimeElements = isoDateTimeElements
}

func (t *Transpiler) ExitUnix_timestamp(c *promqlex.Unix_timestampContext) {
	txt := c.NUMBER().GetText()
	float, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.NUMBER().GetSymbol(), nil)
	}
	t.unixTimestamp = float
}

func (t *Transpiler) ExitConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {
	b := t.constNumExprStack.MustPop()
	a := t.constNumExprStack.MustPop()
	result := math.Pow(a, b)
	t.constNumExprStack.Push(result)
}

func (t *Transpiler) ExitConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {
	if c.UnaryOp().ADD() != nil {
		return
	}
	t.constNumExprStack.Push(-t.constNumExprStack.MustPop())
}

func (t *Transpiler) ExitConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {
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

func (t *Transpiler) ExitConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {
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

func (t *Transpiler) ExitConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (t *Transpiler) ExitConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {}

func (t *Transpiler) ExitNumLit_Number(c *promqlex.NumLit_NumberContext) {
	txt := c.NUMBER().GetText()
	num, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.NUMBER().GetSymbol(), nil)
	}
	t.constNumExprStack.Push(num)
}

func (t *Transpiler) ExitNumLit_Duration(c *promqlex.NumLit_DurationContext) {
	txt := c.DURATION().GetText()
	duration, err := ParseDuration(txt)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.DURATION().GetSymbol(), nil)
	}
	t.constNumExprStack.Push(float64(duration))
}

func (t *Transpiler) ExitNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {}

func (t *Transpiler) ExitNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {
	// TODO: implement
	panic("alias call as numeric literal is not yet supported")
}

func (t *Transpiler) ExitDuration(c *promqlex.DurationContext) {}

func (t *Transpiler) ExitTime_range(c *promqlex.Time_rangeContext) {}

func (t *Transpiler) ExitSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (t *Transpiler) ExitVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (t *Transpiler) ExitVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (t *Transpiler) ExitVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (t *Transpiler) ExitVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (t *Transpiler) ExitVecOp_At(c *promqlex.VecOp_AtContext) {}

func (t *Transpiler) ExitVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (t *Transpiler) ExitVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (t *Transpiler) ExitVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (t *Transpiler) ExitVecOp_SubOp(c *promqlex.VecOp_SubOpContext) {}

func (t *Transpiler) ExitVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *Transpiler) ExitVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *Transpiler) ExitVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *Transpiler) ExitVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *Transpiler) ExitSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (t *Transpiler) ExitOffsetOp(c *promqlex.OffsetOpContext) {}

func (t *Transpiler) ExitMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (t *Transpiler) ExitOffset(c *promqlex.OffsetContext) {}

func (t *Transpiler) ExitLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {
	num := t.constNumExprStack.MustPop()
	txt := strconv.FormatFloat(num, 'f', -1, 64)
	t.rewriter.ReplaceTokenDefault(c.GetStart(), c.GetStop(), txt)
}

func (t *Transpiler) ExitLit_String(c *promqlex.Lit_StringContext) {}

func (t *Transpiler) ExitInstantSelector(c *promqlex.InstantSelectorContext) {}

func (t *Transpiler) ExitLabelName(c *promqlex.LabelNameContext) {}

func (t *Transpiler) ExitMetric_name(c *promqlex.Metric_nameContext) {}

func (t *Transpiler) ExitAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {}

func (t *Transpiler) ExitAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (t *Transpiler) ExitAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (t *Transpiler) ExitExpression(c *promqlex.ExpressionContext) {}

func (t *Transpiler) ExitUnaryOp(c *promqlex.UnaryOpContext) {}

func (t *Transpiler) ExitPowOp(c *promqlex.PowOpContext) {}

func (t *Transpiler) ExitMultOp(c *promqlex.MultOpContext) {}

func (t *Transpiler) ExitAddOp(c *promqlex.AddOpContext) {}

func (t *Transpiler) ExitCompareOp(c *promqlex.CompareOpContext) {}

func (t *Transpiler) ExitAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (t *Transpiler) ExitOrOp(c *promqlex.OrOpContext) {}

func (t *Transpiler) ExitVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (t *Transpiler) ExitVector(c *promqlex.VectorContext) {}

func (t *Transpiler) ExitParens(c *promqlex.ParensContext) {}

func (t *Transpiler) ExitLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (t *Transpiler) ExitLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (t *Transpiler) ExitLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (t *Transpiler) ExitFunction_(c *promqlex.Function_Context) {}

func (t *Transpiler) ExitParameter(c *promqlex.ParameterContext) {}

func (t *Transpiler) ExitParameterList(c *promqlex.ParameterListContext) {}

func (t *Transpiler) ExitAggregation(c *promqlex.AggregationContext) {}

func (t *Transpiler) ExitBy(c *promqlex.ByContext) {}

func (t *Transpiler) ExitWithout(c *promqlex.WithoutContext) {}

func (t *Transpiler) ExitGrouping(c *promqlex.GroupingContext) {}

func (t *Transpiler) ExitOn_(c *promqlex.On_Context) {}

func (t *Transpiler) ExitIgnoring(c *promqlex.IgnoringContext) {}

func (t *Transpiler) ExitGroupLeft(c *promqlex.GroupLeftContext) {}

func (t *Transpiler) ExitGroupRight(c *promqlex.GroupRightContext) {}

func (t *Transpiler) ExitLabelNameList(c *promqlex.LabelNameListContext) {}

func (t *Transpiler) ExitKeyword(c *promqlex.KeywordContext) {}

func (t *Transpiler) ExitString(c *promqlex.StringContext) {}

func unixFloatToTime(unixTimestamp float64) time.Time {
	sec, nano := secondsToSecAndNano(unixTimestamp)
	return time.Unix(int64(sec), int64(nano))
}

func secondsToSecAndNano(seconds float64) (int, int) {
	sec := int(seconds)
	nano := int((seconds - float64(sec)) * 1e9)
	return sec, nano
}

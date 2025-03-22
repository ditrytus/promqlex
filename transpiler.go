package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
	math "math"
	"strconv"
	"time"
)

type Transpiler struct {
	rewriter           *antlr.TokenStreamRewriter
	constNumExprStack  Stack[float64]
	unixTimestamp      float64
	timeInstantLiteral time.Time
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

func (t *Transpiler) setFloat64FieldFromToken(parser antlr.Parser, field **float64, token antlr.Token) {
	if token != nil {
		val, err := strconv.ParseFloat(token.GetText(), 64)
		if err != nil {
			parser.NotifyErrorListeners(err.Error(), token, nil)
		}
		*field = &val
	}
}

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

func (t *Transpiler) EnterVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *Transpiler) EnterVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *Transpiler) EnterVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *Transpiler) EnterVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *Transpiler) EnterVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (t *Transpiler) EnterVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {}

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
	txt := c.RFC_3339_TIMESTAMP().GetText()
	timestamp, err := time.Parse(time.RFC3339, txt)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.RFC_3339_TIMESTAMP().GetSymbol(), nil)
	}
	t.timeInstantLiteral = timestamp
}

func (t *Transpiler) ExitTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
	txt := c.NUMBER().GetText()
	unixTimestamp, err := strconv.ParseFloat(txt, 64)
	if err != nil {
		c.GetParser().NotifyErrorListeners(err.Error(), c.NUMBER().GetSymbol(), nil)
	}
	timeVal := unixFloatToTime(unixTimestamp)
	t.timeInstantLiteral = timeVal
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
	t.constNumExprStack.Push(duration.Seconds())
}

func (t *Transpiler) ExitNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {
	t.constNumExprStack.Push(timeToUnixFloat(t.timeInstantLiteral))
}

func (t *Transpiler) ExitNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {
	// TODO: implement
	panic("alias call as numeric literal is not yet supported")
}

func (t *Transpiler) ExitDuration(c *promqlex.DurationContext) {
	t.replaceWithConstNumExprAsDuration(c)
}

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

func (t *Transpiler) ExitVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (t *Transpiler) ExitVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (t *Transpiler) ExitVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (t *Transpiler) ExitVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (t *Transpiler) ExitVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (t *Transpiler) ExitVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *Transpiler) ExitSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (t *Transpiler) ExitOffsetOp(c *promqlex.OffsetOpContext) {}

func (t *Transpiler) ExitMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (t *Transpiler) ExitOffset(c *promqlex.OffsetContext) {}

func (t *Transpiler) ExitLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

func (t *Transpiler) ExitLit_String(c *promqlex.Lit_StringContext) {}

func (t *Transpiler) ExitInstantSelector(c *promqlex.InstantSelectorContext) {}

func (t *Transpiler) ExitLabelName(c *promqlex.LabelNameContext) {}

func (t *Transpiler) ExitMetric_name(c *promqlex.Metric_nameContext) {}

func (t *Transpiler) ExitAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {
	t.replaceWithConstNumExprAsFloat(c)
}

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

type replaceableNode interface {
	GetStart() antlr.Token
	GetStop() antlr.Token
}

func (t *Transpiler) replaceWithConstNumExprAsFloat(c replaceableNode) {
	num := t.constNumExprStack.MustPop()
	txt := strconv.FormatFloat(num, 'g', 12, 64)
	t.rewriter.ReplaceTokenDefault(c.GetStart(), c.GetStop(), txt)
}

func (t *Transpiler) replaceWithConstNumExprAsDuration(c *promqlex.DurationContext) {
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

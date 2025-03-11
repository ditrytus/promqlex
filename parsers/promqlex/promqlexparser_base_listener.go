// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
import "github.com/antlr4-go/antlr/v4"

// BasePromQLExParserListener is a complete listener for a parse tree produced by PromQLExParser.
type BasePromQLExParserListener struct{}

var _ PromQLExParserListener = &BasePromQLExParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromQLExParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromQLExParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromQLExParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromQLExParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPromqlx is called when production promqlx is entered.
func (s *BasePromQLExParserListener) EnterPromqlx(ctx *PromqlxContext) {}

// ExitPromqlx is called when production promqlx is exited.
func (s *BasePromQLExParserListener) ExitPromqlx(ctx *PromqlxContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePromQLExParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePromQLExParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAlias_def is called when production alias_def is entered.
func (s *BasePromQLExParserListener) EnterAlias_def(ctx *Alias_defContext) {}

// ExitAlias_def is called when production alias_def is exited.
func (s *BasePromQLExParserListener) ExitAlias_def(ctx *Alias_defContext) {}

// EnterAlias_call is called when production alias_call is entered.
func (s *BasePromQLExParserListener) EnterAlias_call(ctx *Alias_callContext) {}

// ExitAlias_call is called when production alias_call is exited.
func (s *BasePromQLExParserListener) ExitAlias_call(ctx *Alias_callContext) {}

// EnterMacro_def is called when production macro_def is entered.
func (s *BasePromQLExParserListener) EnterMacro_def(ctx *Macro_defContext) {}

// ExitMacro_def is called when production macro_def is exited.
func (s *BasePromQLExParserListener) ExitMacro_def(ctx *Macro_defContext) {}

// EnterMacro_call is called when production macro_call is entered.
func (s *BasePromQLExParserListener) EnterMacro_call(ctx *Macro_callContext) {}

// ExitMacro_call is called when production macro_call is exited.
func (s *BasePromQLExParserListener) ExitMacro_call(ctx *Macro_callContext) {}

// EnterArgs_decl is called when production args_decl is entered.
func (s *BasePromQLExParserListener) EnterArgs_decl(ctx *Args_declContext) {}

// ExitArgs_decl is called when production args_decl is exited.
func (s *BasePromQLExParserListener) ExitArgs_decl(ctx *Args_declContext) {}

// EnterArg_name is called when production arg_name is entered.
func (s *BasePromQLExParserListener) EnterArg_name(ctx *Arg_nameContext) {}

// ExitArg_name is called when production arg_name is exited.
func (s *BasePromQLExParserListener) ExitArg_name(ctx *Arg_nameContext) {}

// EnterStatement_block is called when production statement_block is entered.
func (s *BasePromQLExParserListener) EnterStatement_block(ctx *Statement_blockContext) {}

// ExitStatement_block is called when production statement_block is exited.
func (s *BasePromQLExParserListener) ExitStatement_block(ctx *Statement_blockContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BasePromQLExParserListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BasePromQLExParserListener) ExitArg_list(ctx *Arg_listContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasePromQLExParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasePromQLExParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasePromQLExParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasePromQLExParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterCompareVectorOperation is called when production compareVectorOperation is entered.
func (s *BasePromQLExParserListener) EnterCompareVectorOperation(ctx *CompareVectorOperationContext) {
}

// ExitCompareVectorOperation is called when production compareVectorOperation is exited.
func (s *BasePromQLExParserListener) ExitCompareVectorOperation(ctx *CompareVectorOperationContext) {}

// EnterTrueConst is called when production trueConst is entered.
func (s *BasePromQLExParserListener) EnterTrueConst(ctx *TrueConstContext) {}

// ExitTrueConst is called when production trueConst is exited.
func (s *BasePromQLExParserListener) ExitTrueConst(ctx *TrueConstContext) {}

// EnterFalseConst is called when production falseConst is entered.
func (s *BasePromQLExParserListener) EnterFalseConst(ctx *FalseConstContext) {}

// ExitFalseConst is called when production falseConst is exited.
func (s *BasePromQLExParserListener) ExitFalseConst(ctx *FalseConstContext) {}

// EnterTime_instant_literal is called when production time_instant_literal is entered.
func (s *BasePromQLExParserListener) EnterTime_instant_literal(ctx *Time_instant_literalContext) {}

// ExitTime_instant_literal is called when production time_instant_literal is exited.
func (s *BasePromQLExParserListener) ExitTime_instant_literal(ctx *Time_instant_literalContext) {}

// EnterIso_date_time is called when production iso_date_time is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time(ctx *Iso_date_timeContext) {}

// ExitIso_date_time is called when production iso_date_time is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time(ctx *Iso_date_timeContext) {}

// EnterIso_date_time_ymdhmsf is called when production iso_date_time_ymdhmsf is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ymdhmsf(ctx *Iso_date_time_ymdhmsfContext) {}

// ExitIso_date_time_ymdhmsf is called when production iso_date_time_ymdhmsf is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ymdhmsf(ctx *Iso_date_time_ymdhmsfContext) {}

// EnterIso_date_time_ymdhms is called when production iso_date_time_ymdhms is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ymdhms(ctx *Iso_date_time_ymdhmsContext) {}

// ExitIso_date_time_ymdhms is called when production iso_date_time_ymdhms is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ymdhms(ctx *Iso_date_time_ymdhmsContext) {}

// EnterIso_date_time_ymdhm is called when production iso_date_time_ymdhm is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ymdhm(ctx *Iso_date_time_ymdhmContext) {}

// ExitIso_date_time_ymdhm is called when production iso_date_time_ymdhm is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ymdhm(ctx *Iso_date_time_ymdhmContext) {}

// EnterIso_date_time_ymdh is called when production iso_date_time_ymdh is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ymdh(ctx *Iso_date_time_ymdhContext) {}

// ExitIso_date_time_ymdh is called when production iso_date_time_ymdh is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ymdh(ctx *Iso_date_time_ymdhContext) {}

// EnterIso_date_time_ymd is called when production iso_date_time_ymd is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ymd(ctx *Iso_date_time_ymdContext) {}

// ExitIso_date_time_ymd is called when production iso_date_time_ymd is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ymd(ctx *Iso_date_time_ymdContext) {}

// EnterIso_date_time_ym is called when production iso_date_time_ym is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_ym(ctx *Iso_date_time_ymContext) {}

// ExitIso_date_time_ym is called when production iso_date_time_ym is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_ym(ctx *Iso_date_time_ymContext) {}

// EnterIso_date_time_y is called when production iso_date_time_y is entered.
func (s *BasePromQLExParserListener) EnterIso_date_time_y(ctx *Iso_date_time_yContext) {}

// ExitIso_date_time_y is called when production iso_date_time_y is exited.
func (s *BasePromQLExParserListener) ExitIso_date_time_y(ctx *Iso_date_time_yContext) {}

// EnterIso_year is called when production iso_year is entered.
func (s *BasePromQLExParserListener) EnterIso_year(ctx *Iso_yearContext) {}

// ExitIso_year is called when production iso_year is exited.
func (s *BasePromQLExParserListener) ExitIso_year(ctx *Iso_yearContext) {}

// EnterIso_month is called when production iso_month is entered.
func (s *BasePromQLExParserListener) EnterIso_month(ctx *Iso_monthContext) {}

// ExitIso_month is called when production iso_month is exited.
func (s *BasePromQLExParserListener) ExitIso_month(ctx *Iso_monthContext) {}

// EnterIso_day is called when production iso_day is entered.
func (s *BasePromQLExParserListener) EnterIso_day(ctx *Iso_dayContext) {}

// ExitIso_day is called when production iso_day is exited.
func (s *BasePromQLExParserListener) ExitIso_day(ctx *Iso_dayContext) {}

// EnterIso_hour is called when production iso_hour is entered.
func (s *BasePromQLExParserListener) EnterIso_hour(ctx *Iso_hourContext) {}

// ExitIso_hour is called when production iso_hour is exited.
func (s *BasePromQLExParserListener) ExitIso_hour(ctx *Iso_hourContext) {}

// EnterIso_minutes is called when production iso_minutes is entered.
func (s *BasePromQLExParserListener) EnterIso_minutes(ctx *Iso_minutesContext) {}

// ExitIso_minutes is called when production iso_minutes is exited.
func (s *BasePromQLExParserListener) ExitIso_minutes(ctx *Iso_minutesContext) {}

// EnterIso_seconds is called when production iso_seconds is entered.
func (s *BasePromQLExParserListener) EnterIso_seconds(ctx *Iso_secondsContext) {}

// ExitIso_seconds is called when production iso_seconds is exited.
func (s *BasePromQLExParserListener) ExitIso_seconds(ctx *Iso_secondsContext) {}

// EnterIs_frac_sec is called when production is_frac_sec is entered.
func (s *BasePromQLExParserListener) EnterIs_frac_sec(ctx *Is_frac_secContext) {}

// ExitIs_frac_sec is called when production is_frac_sec is exited.
func (s *BasePromQLExParserListener) ExitIs_frac_sec(ctx *Is_frac_secContext) {}

// EnterUnix_timestamp is called when production unix_timestamp is entered.
func (s *BasePromQLExParserListener) EnterUnix_timestamp(ctx *Unix_timestampContext) {}

// ExitUnix_timestamp is called when production unix_timestamp is exited.
func (s *BasePromQLExParserListener) ExitUnix_timestamp(ctx *Unix_timestampContext) {}

// EnterConst_num_expression is called when production const_num_expression is entered.
func (s *BasePromQLExParserListener) EnterConst_num_expression(ctx *Const_num_expressionContext) {}

// ExitConst_num_expression is called when production const_num_expression is exited.
func (s *BasePromQLExParserListener) ExitConst_num_expression(ctx *Const_num_expressionContext) {}

// EnterNum_literal is called when production num_literal is entered.
func (s *BasePromQLExParserListener) EnterNum_literal(ctx *Num_literalContext) {}

// ExitNum_literal is called when production num_literal is exited.
func (s *BasePromQLExParserListener) ExitNum_literal(ctx *Num_literalContext) {}

// EnterDuration is called when production duration is entered.
func (s *BasePromQLExParserListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BasePromQLExParserListener) ExitDuration(ctx *DurationContext) {}

// EnterTime_range is called when production time_range is entered.
func (s *BasePromQLExParserListener) EnterTime_range(ctx *Time_rangeContext) {}

// ExitTime_range is called when production time_range is exited.
func (s *BasePromQLExParserListener) ExitTime_range(ctx *Time_rangeContext) {}

// EnterSubquery_range is called when production subquery_range is entered.
func (s *BasePromQLExParserListener) EnterSubquery_range(ctx *Subquery_rangeContext) {}

// ExitSubquery_range is called when production subquery_range is exited.
func (s *BasePromQLExParserListener) ExitSubquery_range(ctx *Subquery_rangeContext) {}

// EnterVectorOperation is called when production vectorOperation is entered.
func (s *BasePromQLExParserListener) EnterVectorOperation(ctx *VectorOperationContext) {}

// ExitVectorOperation is called when production vectorOperation is exited.
func (s *BasePromQLExParserListener) ExitVectorOperation(ctx *VectorOperationContext) {}

// EnterSubqueryOp is called when production subqueryOp is entered.
func (s *BasePromQLExParserListener) EnterSubqueryOp(ctx *SubqueryOpContext) {}

// ExitSubqueryOp is called when production subqueryOp is exited.
func (s *BasePromQLExParserListener) ExitSubqueryOp(ctx *SubqueryOpContext) {}

// EnterOffsetOp is called when production offsetOp is entered.
func (s *BasePromQLExParserListener) EnterOffsetOp(ctx *OffsetOpContext) {}

// ExitOffsetOp is called when production offsetOp is exited.
func (s *BasePromQLExParserListener) ExitOffsetOp(ctx *OffsetOpContext) {}

// EnterMatrixSelector is called when production matrixSelector is entered.
func (s *BasePromQLExParserListener) EnterMatrixSelector(ctx *MatrixSelectorContext) {}

// ExitMatrixSelector is called when production matrixSelector is exited.
func (s *BasePromQLExParserListener) ExitMatrixSelector(ctx *MatrixSelectorContext) {}

// EnterOffset is called when production offset is entered.
func (s *BasePromQLExParserListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BasePromQLExParserListener) ExitOffset(ctx *OffsetContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePromQLExParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePromQLExParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInstantSelector is called when production instantSelector is entered.
func (s *BasePromQLExParserListener) EnterInstantSelector(ctx *InstantSelectorContext) {}

// ExitInstantSelector is called when production instantSelector is exited.
func (s *BasePromQLExParserListener) ExitInstantSelector(ctx *InstantSelectorContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BasePromQLExParserListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BasePromQLExParserListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterMetric_name is called when production metric_name is entered.
func (s *BasePromQLExParserListener) EnterMetric_name(ctx *Metric_nameContext) {}

// ExitMetric_name is called when production metric_name is exited.
func (s *BasePromQLExParserListener) ExitMetric_name(ctx *Metric_nameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePromQLExParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePromQLExParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BasePromQLExParserListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BasePromQLExParserListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterPowOp is called when production powOp is entered.
func (s *BasePromQLExParserListener) EnterPowOp(ctx *PowOpContext) {}

// ExitPowOp is called when production powOp is exited.
func (s *BasePromQLExParserListener) ExitPowOp(ctx *PowOpContext) {}

// EnterMultOp is called when production multOp is entered.
func (s *BasePromQLExParserListener) EnterMultOp(ctx *MultOpContext) {}

// ExitMultOp is called when production multOp is exited.
func (s *BasePromQLExParserListener) ExitMultOp(ctx *MultOpContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BasePromQLExParserListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BasePromQLExParserListener) ExitAddOp(ctx *AddOpContext) {}

// EnterCompareOp is called when production compareOp is entered.
func (s *BasePromQLExParserListener) EnterCompareOp(ctx *CompareOpContext) {}

// ExitCompareOp is called when production compareOp is exited.
func (s *BasePromQLExParserListener) ExitCompareOp(ctx *CompareOpContext) {}

// EnterAndUnlessOp is called when production andUnlessOp is entered.
func (s *BasePromQLExParserListener) EnterAndUnlessOp(ctx *AndUnlessOpContext) {}

// ExitAndUnlessOp is called when production andUnlessOp is exited.
func (s *BasePromQLExParserListener) ExitAndUnlessOp(ctx *AndUnlessOpContext) {}

// EnterOrOp is called when production orOp is entered.
func (s *BasePromQLExParserListener) EnterOrOp(ctx *OrOpContext) {}

// ExitOrOp is called when production orOp is exited.
func (s *BasePromQLExParserListener) ExitOrOp(ctx *OrOpContext) {}

// EnterVectorMatchOp is called when production vectorMatchOp is entered.
func (s *BasePromQLExParserListener) EnterVectorMatchOp(ctx *VectorMatchOpContext) {}

// ExitVectorMatchOp is called when production vectorMatchOp is exited.
func (s *BasePromQLExParserListener) ExitVectorMatchOp(ctx *VectorMatchOpContext) {}

// EnterVector is called when production vector is entered.
func (s *BasePromQLExParserListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BasePromQLExParserListener) ExitVector(ctx *VectorContext) {}

// EnterParens is called when production parens is entered.
func (s *BasePromQLExParserListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BasePromQLExParserListener) ExitParens(ctx *ParensContext) {}

// EnterLabelMatcher is called when production labelMatcher is entered.
func (s *BasePromQLExParserListener) EnterLabelMatcher(ctx *LabelMatcherContext) {}

// ExitLabelMatcher is called when production labelMatcher is exited.
func (s *BasePromQLExParserListener) ExitLabelMatcher(ctx *LabelMatcherContext) {}

// EnterLabelMatcherOperator is called when production labelMatcherOperator is entered.
func (s *BasePromQLExParserListener) EnterLabelMatcherOperator(ctx *LabelMatcherOperatorContext) {}

// ExitLabelMatcherOperator is called when production labelMatcherOperator is exited.
func (s *BasePromQLExParserListener) ExitLabelMatcherOperator(ctx *LabelMatcherOperatorContext) {}

// EnterLabelMatcherList is called when production labelMatcherList is entered.
func (s *BasePromQLExParserListener) EnterLabelMatcherList(ctx *LabelMatcherListContext) {}

// ExitLabelMatcherList is called when production labelMatcherList is exited.
func (s *BasePromQLExParserListener) ExitLabelMatcherList(ctx *LabelMatcherListContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BasePromQLExParserListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BasePromQLExParserListener) ExitFunction_(ctx *Function_Context) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePromQLExParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePromQLExParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasePromQLExParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasePromQLExParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BasePromQLExParserListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BasePromQLExParserListener) ExitAggregation(ctx *AggregationContext) {}

// EnterBy is called when production by is entered.
func (s *BasePromQLExParserListener) EnterBy(ctx *ByContext) {}

// ExitBy is called when production by is exited.
func (s *BasePromQLExParserListener) ExitBy(ctx *ByContext) {}

// EnterWithout is called when production without is entered.
func (s *BasePromQLExParserListener) EnterWithout(ctx *WithoutContext) {}

// ExitWithout is called when production without is exited.
func (s *BasePromQLExParserListener) ExitWithout(ctx *WithoutContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BasePromQLExParserListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BasePromQLExParserListener) ExitGrouping(ctx *GroupingContext) {}

// EnterOn_ is called when production on_ is entered.
func (s *BasePromQLExParserListener) EnterOn_(ctx *On_Context) {}

// ExitOn_ is called when production on_ is exited.
func (s *BasePromQLExParserListener) ExitOn_(ctx *On_Context) {}

// EnterIgnoring is called when production ignoring is entered.
func (s *BasePromQLExParserListener) EnterIgnoring(ctx *IgnoringContext) {}

// ExitIgnoring is called when production ignoring is exited.
func (s *BasePromQLExParserListener) ExitIgnoring(ctx *IgnoringContext) {}

// EnterGroupLeft is called when production groupLeft is entered.
func (s *BasePromQLExParserListener) EnterGroupLeft(ctx *GroupLeftContext) {}

// ExitGroupLeft is called when production groupLeft is exited.
func (s *BasePromQLExParserListener) ExitGroupLeft(ctx *GroupLeftContext) {}

// EnterGroupRight is called when production groupRight is entered.
func (s *BasePromQLExParserListener) EnterGroupRight(ctx *GroupRightContext) {}

// ExitGroupRight is called when production groupRight is exited.
func (s *BasePromQLExParserListener) ExitGroupRight(ctx *GroupRightContext) {}

// EnterLabelNameList is called when production labelNameList is entered.
func (s *BasePromQLExParserListener) EnterLabelNameList(ctx *LabelNameListContext) {}

// ExitLabelNameList is called when production labelNameList is exited.
func (s *BasePromQLExParserListener) ExitLabelNameList(ctx *LabelNameListContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BasePromQLExParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BasePromQLExParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterString is called when production string is entered.
func (s *BasePromQLExParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasePromQLExParserListener) ExitString(ctx *StringContext) {}

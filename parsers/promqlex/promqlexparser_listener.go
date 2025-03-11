// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
import "github.com/antlr4-go/antlr/v4"

// PromQLExParserListener is a complete listener for a parse tree produced by PromQLExParser.
type PromQLExParserListener interface {
	antlr.ParseTreeListener

	// EnterPromqlx is called when entering the promqlx production.
	EnterPromqlx(c *PromqlxContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAlias_def is called when entering the alias_def production.
	EnterAlias_def(c *Alias_defContext)

	// EnterAlias_call is called when entering the alias_call production.
	EnterAlias_call(c *Alias_callContext)

	// EnterMacro_def is called when entering the macro_def production.
	EnterMacro_def(c *Macro_defContext)

	// EnterMacro_call is called when entering the macro_call production.
	EnterMacro_call(c *Macro_callContext)

	// EnterArgs_decl is called when entering the args_decl production.
	EnterArgs_decl(c *Args_declContext)

	// EnterArg_name is called when entering the arg_name production.
	EnterArg_name(c *Arg_nameContext)

	// EnterStatement_block is called when entering the statement_block production.
	EnterStatement_block(c *Statement_blockContext)

	// EnterArg_list is called when entering the arg_list production.
	EnterArg_list(c *Arg_listContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterCompareVectorOperation is called when entering the compareVectorOperation production.
	EnterCompareVectorOperation(c *CompareVectorOperationContext)

	// EnterTrueConst is called when entering the trueConst production.
	EnterTrueConst(c *TrueConstContext)

	// EnterFalseConst is called when entering the falseConst production.
	EnterFalseConst(c *FalseConstContext)

	// EnterTime_instant_literal is called when entering the time_instant_literal production.
	EnterTime_instant_literal(c *Time_instant_literalContext)

	// EnterIso_date_time is called when entering the iso_date_time production.
	EnterIso_date_time(c *Iso_date_timeContext)

	// EnterIso_date_time_ymdhmsf is called when entering the iso_date_time_ymdhmsf production.
	EnterIso_date_time_ymdhmsf(c *Iso_date_time_ymdhmsfContext)

	// EnterIso_date_time_ymdhms is called when entering the iso_date_time_ymdhms production.
	EnterIso_date_time_ymdhms(c *Iso_date_time_ymdhmsContext)

	// EnterIso_date_time_ymdhm is called when entering the iso_date_time_ymdhm production.
	EnterIso_date_time_ymdhm(c *Iso_date_time_ymdhmContext)

	// EnterIso_date_time_ymdh is called when entering the iso_date_time_ymdh production.
	EnterIso_date_time_ymdh(c *Iso_date_time_ymdhContext)

	// EnterIso_date_time_ymd is called when entering the iso_date_time_ymd production.
	EnterIso_date_time_ymd(c *Iso_date_time_ymdContext)

	// EnterIso_date_time_ym is called when entering the iso_date_time_ym production.
	EnterIso_date_time_ym(c *Iso_date_time_ymContext)

	// EnterIso_date_time_y is called when entering the iso_date_time_y production.
	EnterIso_date_time_y(c *Iso_date_time_yContext)

	// EnterIso_year is called when entering the iso_year production.
	EnterIso_year(c *Iso_yearContext)

	// EnterIso_month is called when entering the iso_month production.
	EnterIso_month(c *Iso_monthContext)

	// EnterIso_day is called when entering the iso_day production.
	EnterIso_day(c *Iso_dayContext)

	// EnterIso_hour is called when entering the iso_hour production.
	EnterIso_hour(c *Iso_hourContext)

	// EnterIso_minutes is called when entering the iso_minutes production.
	EnterIso_minutes(c *Iso_minutesContext)

	// EnterIso_seconds is called when entering the iso_seconds production.
	EnterIso_seconds(c *Iso_secondsContext)

	// EnterIs_frac_sec is called when entering the is_frac_sec production.
	EnterIs_frac_sec(c *Is_frac_secContext)

	// EnterUnix_timestamp is called when entering the unix_timestamp production.
	EnterUnix_timestamp(c *Unix_timestampContext)

	// EnterConst_num_expression is called when entering the const_num_expression production.
	EnterConst_num_expression(c *Const_num_expressionContext)

	// EnterNum_literal is called when entering the num_literal production.
	EnterNum_literal(c *Num_literalContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// EnterTime_range is called when entering the time_range production.
	EnterTime_range(c *Time_rangeContext)

	// EnterSubquery_range is called when entering the subquery_range production.
	EnterSubquery_range(c *Subquery_rangeContext)

	// EnterVectorOperation is called when entering the vectorOperation production.
	EnterVectorOperation(c *VectorOperationContext)

	// EnterSubqueryOp is called when entering the subqueryOp production.
	EnterSubqueryOp(c *SubqueryOpContext)

	// EnterOffsetOp is called when entering the offsetOp production.
	EnterOffsetOp(c *OffsetOpContext)

	// EnterMatrixSelector is called when entering the matrixSelector production.
	EnterMatrixSelector(c *MatrixSelectorContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInstantSelector is called when entering the instantSelector production.
	EnterInstantSelector(c *InstantSelectorContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterMetric_name is called when entering the metric_name production.
	EnterMetric_name(c *Metric_nameContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterPowOp is called when entering the powOp production.
	EnterPowOp(c *PowOpContext)

	// EnterMultOp is called when entering the multOp production.
	EnterMultOp(c *MultOpContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterCompareOp is called when entering the compareOp production.
	EnterCompareOp(c *CompareOpContext)

	// EnterAndUnlessOp is called when entering the andUnlessOp production.
	EnterAndUnlessOp(c *AndUnlessOpContext)

	// EnterOrOp is called when entering the orOp production.
	EnterOrOp(c *OrOpContext)

	// EnterVectorMatchOp is called when entering the vectorMatchOp production.
	EnterVectorMatchOp(c *VectorMatchOpContext)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterLabelMatcherOperator is called when entering the labelMatcherOperator production.
	EnterLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// EnterLabelMatcherList is called when entering the labelMatcherList production.
	EnterLabelMatcherList(c *LabelMatcherListContext)

	// EnterFunction_ is called when entering the function_ production.
	EnterFunction_(c *Function_Context)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterBy is called when entering the by production.
	EnterBy(c *ByContext)

	// EnterWithout is called when entering the without production.
	EnterWithout(c *WithoutContext)

	// EnterGrouping is called when entering the grouping production.
	EnterGrouping(c *GroupingContext)

	// EnterOn_ is called when entering the on_ production.
	EnterOn_(c *On_Context)

	// EnterIgnoring is called when entering the ignoring production.
	EnterIgnoring(c *IgnoringContext)

	// EnterGroupLeft is called when entering the groupLeft production.
	EnterGroupLeft(c *GroupLeftContext)

	// EnterGroupRight is called when entering the groupRight production.
	EnterGroupRight(c *GroupRightContext)

	// EnterLabelNameList is called when entering the labelNameList production.
	EnterLabelNameList(c *LabelNameListContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// ExitPromqlx is called when exiting the promqlx production.
	ExitPromqlx(c *PromqlxContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAlias_def is called when exiting the alias_def production.
	ExitAlias_def(c *Alias_defContext)

	// ExitAlias_call is called when exiting the alias_call production.
	ExitAlias_call(c *Alias_callContext)

	// ExitMacro_def is called when exiting the macro_def production.
	ExitMacro_def(c *Macro_defContext)

	// ExitMacro_call is called when exiting the macro_call production.
	ExitMacro_call(c *Macro_callContext)

	// ExitArgs_decl is called when exiting the args_decl production.
	ExitArgs_decl(c *Args_declContext)

	// ExitArg_name is called when exiting the arg_name production.
	ExitArg_name(c *Arg_nameContext)

	// ExitStatement_block is called when exiting the statement_block production.
	ExitStatement_block(c *Statement_blockContext)

	// ExitArg_list is called when exiting the arg_list production.
	ExitArg_list(c *Arg_listContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitCompareVectorOperation is called when exiting the compareVectorOperation production.
	ExitCompareVectorOperation(c *CompareVectorOperationContext)

	// ExitTrueConst is called when exiting the trueConst production.
	ExitTrueConst(c *TrueConstContext)

	// ExitFalseConst is called when exiting the falseConst production.
	ExitFalseConst(c *FalseConstContext)

	// ExitTime_instant_literal is called when exiting the time_instant_literal production.
	ExitTime_instant_literal(c *Time_instant_literalContext)

	// ExitIso_date_time is called when exiting the iso_date_time production.
	ExitIso_date_time(c *Iso_date_timeContext)

	// ExitIso_date_time_ymdhmsf is called when exiting the iso_date_time_ymdhmsf production.
	ExitIso_date_time_ymdhmsf(c *Iso_date_time_ymdhmsfContext)

	// ExitIso_date_time_ymdhms is called when exiting the iso_date_time_ymdhms production.
	ExitIso_date_time_ymdhms(c *Iso_date_time_ymdhmsContext)

	// ExitIso_date_time_ymdhm is called when exiting the iso_date_time_ymdhm production.
	ExitIso_date_time_ymdhm(c *Iso_date_time_ymdhmContext)

	// ExitIso_date_time_ymdh is called when exiting the iso_date_time_ymdh production.
	ExitIso_date_time_ymdh(c *Iso_date_time_ymdhContext)

	// ExitIso_date_time_ymd is called when exiting the iso_date_time_ymd production.
	ExitIso_date_time_ymd(c *Iso_date_time_ymdContext)

	// ExitIso_date_time_ym is called when exiting the iso_date_time_ym production.
	ExitIso_date_time_ym(c *Iso_date_time_ymContext)

	// ExitIso_date_time_y is called when exiting the iso_date_time_y production.
	ExitIso_date_time_y(c *Iso_date_time_yContext)

	// ExitIso_year is called when exiting the iso_year production.
	ExitIso_year(c *Iso_yearContext)

	// ExitIso_month is called when exiting the iso_month production.
	ExitIso_month(c *Iso_monthContext)

	// ExitIso_day is called when exiting the iso_day production.
	ExitIso_day(c *Iso_dayContext)

	// ExitIso_hour is called when exiting the iso_hour production.
	ExitIso_hour(c *Iso_hourContext)

	// ExitIso_minutes is called when exiting the iso_minutes production.
	ExitIso_minutes(c *Iso_minutesContext)

	// ExitIso_seconds is called when exiting the iso_seconds production.
	ExitIso_seconds(c *Iso_secondsContext)

	// ExitIs_frac_sec is called when exiting the is_frac_sec production.
	ExitIs_frac_sec(c *Is_frac_secContext)

	// ExitUnix_timestamp is called when exiting the unix_timestamp production.
	ExitUnix_timestamp(c *Unix_timestampContext)

	// ExitConst_num_expression is called when exiting the const_num_expression production.
	ExitConst_num_expression(c *Const_num_expressionContext)

	// ExitNum_literal is called when exiting the num_literal production.
	ExitNum_literal(c *Num_literalContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)

	// ExitTime_range is called when exiting the time_range production.
	ExitTime_range(c *Time_rangeContext)

	// ExitSubquery_range is called when exiting the subquery_range production.
	ExitSubquery_range(c *Subquery_rangeContext)

	// ExitVectorOperation is called when exiting the vectorOperation production.
	ExitVectorOperation(c *VectorOperationContext)

	// ExitSubqueryOp is called when exiting the subqueryOp production.
	ExitSubqueryOp(c *SubqueryOpContext)

	// ExitOffsetOp is called when exiting the offsetOp production.
	ExitOffsetOp(c *OffsetOpContext)

	// ExitMatrixSelector is called when exiting the matrixSelector production.
	ExitMatrixSelector(c *MatrixSelectorContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInstantSelector is called when exiting the instantSelector production.
	ExitInstantSelector(c *InstantSelectorContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitMetric_name is called when exiting the metric_name production.
	ExitMetric_name(c *Metric_nameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitPowOp is called when exiting the powOp production.
	ExitPowOp(c *PowOpContext)

	// ExitMultOp is called when exiting the multOp production.
	ExitMultOp(c *MultOpContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitCompareOp is called when exiting the compareOp production.
	ExitCompareOp(c *CompareOpContext)

	// ExitAndUnlessOp is called when exiting the andUnlessOp production.
	ExitAndUnlessOp(c *AndUnlessOpContext)

	// ExitOrOp is called when exiting the orOp production.
	ExitOrOp(c *OrOpContext)

	// ExitVectorMatchOp is called when exiting the vectorMatchOp production.
	ExitVectorMatchOp(c *VectorMatchOpContext)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitLabelMatcherOperator is called when exiting the labelMatcherOperator production.
	ExitLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// ExitLabelMatcherList is called when exiting the labelMatcherList production.
	ExitLabelMatcherList(c *LabelMatcherListContext)

	// ExitFunction_ is called when exiting the function_ production.
	ExitFunction_(c *Function_Context)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitBy is called when exiting the by production.
	ExitBy(c *ByContext)

	// ExitWithout is called when exiting the without production.
	ExitWithout(c *WithoutContext)

	// ExitGrouping is called when exiting the grouping production.
	ExitGrouping(c *GroupingContext)

	// ExitOn_ is called when exiting the on_ production.
	ExitOn_(c *On_Context)

	// ExitIgnoring is called when exiting the ignoring production.
	ExitIgnoring(c *IgnoringContext)

	// ExitGroupLeft is called when exiting the groupLeft production.
	ExitGroupLeft(c *GroupLeftContext)

	// ExitGroupRight is called when exiting the groupRight production.
	ExitGroupRight(c *GroupRightContext)

	// ExitLabelNameList is called when exiting the labelNameList production.
	ExitLabelNameList(c *LabelNameListContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)
}

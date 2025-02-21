// Code generated from PromQLExtensionParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLExtensionParser

import "github.com/antlr4-go/antlr/v4"

// PromQLExtensionParserListener is a complete listener for a parse tree produced by PromQLExtensionParser.
type PromQLExtensionParserListener interface {
	antlr.ParseTreeListener

	// EnterPromqlx is called when entering the promqlx production.
	EnterPromqlx(c *PromqlxContext)

	// EnterEx_statement is called when entering the ex_statement production.
	EnterEx_statement(c *Ex_statementContext)

	// EnterEx_alias_def is called when entering the ex_alias_def production.
	EnterEx_alias_def(c *Ex_alias_defContext)

	// EnterEx_alias_call is called when entering the ex_alias_call production.
	EnterEx_alias_call(c *Ex_alias_callContext)

	// EnterEx_macro_def is called when entering the ex_macro_def production.
	EnterEx_macro_def(c *Ex_macro_defContext)

	// EnterEx_macro_call is called when entering the ex_macro_call production.
	EnterEx_macro_call(c *Ex_macro_callContext)

	// EnterEx_args_decl is called when entering the ex_args_decl production.
	EnterEx_args_decl(c *Ex_args_declContext)

	// EnterEx_arg_name is called when entering the ex_arg_name production.
	EnterEx_arg_name(c *Ex_arg_nameContext)

	// EnterEx_block is called when entering the ex_block production.
	EnterEx_block(c *Ex_blockContext)

	// EnterEx_arg_list is called when entering the ex_arg_list production.
	EnterEx_arg_list(c *Ex_arg_listContext)

	// EnterEx_if_statement is called when entering the ex_if_statement production.
	EnterEx_if_statement(c *Ex_if_statementContext)

	// EnterEx_condition is called when entering the ex_condition production.
	EnterEx_condition(c *Ex_conditionContext)

	// EnterEx_compareVectorOperation is called when entering the ex_compareVectorOperation production.
	EnterEx_compareVectorOperation(c *Ex_compareVectorOperationContext)

	// EnterEx_trueConst is called when entering the ex_trueConst production.
	EnterEx_trueConst(c *Ex_trueConstContext)

	// EnterEx_falseConst is called when entering the ex_falseConst production.
	EnterEx_falseConst(c *Ex_falseConstContext)

	// EnterEx_time_instant_literal is called when entering the ex_time_instant_literal production.
	EnterEx_time_instant_literal(c *Ex_time_instant_literalContext)

	// EnterEx_iso_date_time is called when entering the ex_iso_date_time production.
	EnterEx_iso_date_time(c *Ex_iso_date_timeContext)

	// EnterEx_iso_date_time_ymdhmsf is called when entering the ex_iso_date_time_ymdhmsf production.
	EnterEx_iso_date_time_ymdhmsf(c *Ex_iso_date_time_ymdhmsfContext)

	// EnterEx_iso_date_time_ymdhms is called when entering the ex_iso_date_time_ymdhms production.
	EnterEx_iso_date_time_ymdhms(c *Ex_iso_date_time_ymdhmsContext)

	// EnterEx_iso_date_time_ymdhm is called when entering the ex_iso_date_time_ymdhm production.
	EnterEx_iso_date_time_ymdhm(c *Ex_iso_date_time_ymdhmContext)

	// EnterEx_iso_date_time_ymdh is called when entering the ex_iso_date_time_ymdh production.
	EnterEx_iso_date_time_ymdh(c *Ex_iso_date_time_ymdhContext)

	// EnterEx_iso_date_time_ymd is called when entering the ex_iso_date_time_ymd production.
	EnterEx_iso_date_time_ymd(c *Ex_iso_date_time_ymdContext)

	// EnterEx_iso_date_time_ym is called when entering the ex_iso_date_time_ym production.
	EnterEx_iso_date_time_ym(c *Ex_iso_date_time_ymContext)

	// EnterEx_iso_date_time_y is called when entering the ex_iso_date_time_y production.
	EnterEx_iso_date_time_y(c *Ex_iso_date_time_yContext)

	// EnterEx_year is called when entering the ex_year production.
	EnterEx_year(c *Ex_yearContext)

	// EnterEx_month is called when entering the ex_month production.
	EnterEx_month(c *Ex_monthContext)

	// EnterEx_day is called when entering the ex_day production.
	EnterEx_day(c *Ex_dayContext)

	// EnterEx_hour is called when entering the ex_hour production.
	EnterEx_hour(c *Ex_hourContext)

	// EnterEx_minutes is called when entering the ex_minutes production.
	EnterEx_minutes(c *Ex_minutesContext)

	// EnterEx_seconds is called when entering the ex_seconds production.
	EnterEx_seconds(c *Ex_secondsContext)

	// EnterEx_frac_sec is called when entering the ex_frac_sec production.
	EnterEx_frac_sec(c *Ex_frac_secContext)

	// EnterEx_unix_timestamp is called when entering the ex_unix_timestamp production.
	EnterEx_unix_timestamp(c *Ex_unix_timestampContext)

	// EnterEx_const_num_expression is called when entering the ex_const_num_expression production.
	EnterEx_const_num_expression(c *Ex_const_num_expressionContext)

	// EnterEx_num_literal is called when entering the ex_num_literal production.
	EnterEx_num_literal(c *Ex_num_literalContext)

	// EnterEx_duration is called when entering the ex_duration production.
	EnterEx_duration(c *Ex_durationContext)

	// EnterEx_time_range is called when entering the ex_time_range production.
	EnterEx_time_range(c *Ex_time_rangeContext)

	// EnterEx_subquery_range is called when entering the ex_subquery_range production.
	EnterEx_subquery_range(c *Ex_subquery_rangeContext)

	// EnterVectorOperation is called when entering the vectorOperation production.
	EnterVectorOperation(c *VectorOperationContext)

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

	// EnterSubqueryOp is called when entering the subqueryOp production.
	EnterSubqueryOp(c *SubqueryOpContext)

	// EnterOffsetOp is called when entering the offsetOp production.
	EnterOffsetOp(c *OffsetOpContext)

	// EnterVector is called when entering the vector production.
	EnterVector(c *VectorContext)

	// EnterParens is called when entering the parens production.
	EnterParens(c *ParensContext)

	// EnterInstantSelector is called when entering the instantSelector production.
	EnterInstantSelector(c *InstantSelectorContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterLabelMatcherOperator is called when entering the labelMatcherOperator production.
	EnterLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// EnterLabelMatcherList is called when entering the labelMatcherList production.
	EnterLabelMatcherList(c *LabelMatcherListContext)

	// EnterMatrixSelector is called when entering the matrixSelector production.
	EnterMatrixSelector(c *MatrixSelectorContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

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

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterLabelNameList is called when entering the labelNameList production.
	EnterLabelNameList(c *LabelNameListContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// ExitPromqlx is called when exiting the promqlx production.
	ExitPromqlx(c *PromqlxContext)

	// ExitEx_statement is called when exiting the ex_statement production.
	ExitEx_statement(c *Ex_statementContext)

	// ExitEx_alias_def is called when exiting the ex_alias_def production.
	ExitEx_alias_def(c *Ex_alias_defContext)

	// ExitEx_alias_call is called when exiting the ex_alias_call production.
	ExitEx_alias_call(c *Ex_alias_callContext)

	// ExitEx_macro_def is called when exiting the ex_macro_def production.
	ExitEx_macro_def(c *Ex_macro_defContext)

	// ExitEx_macro_call is called when exiting the ex_macro_call production.
	ExitEx_macro_call(c *Ex_macro_callContext)

	// ExitEx_args_decl is called when exiting the ex_args_decl production.
	ExitEx_args_decl(c *Ex_args_declContext)

	// ExitEx_arg_name is called when exiting the ex_arg_name production.
	ExitEx_arg_name(c *Ex_arg_nameContext)

	// ExitEx_block is called when exiting the ex_block production.
	ExitEx_block(c *Ex_blockContext)

	// ExitEx_arg_list is called when exiting the ex_arg_list production.
	ExitEx_arg_list(c *Ex_arg_listContext)

	// ExitEx_if_statement is called when exiting the ex_if_statement production.
	ExitEx_if_statement(c *Ex_if_statementContext)

	// ExitEx_condition is called when exiting the ex_condition production.
	ExitEx_condition(c *Ex_conditionContext)

	// ExitEx_compareVectorOperation is called when exiting the ex_compareVectorOperation production.
	ExitEx_compareVectorOperation(c *Ex_compareVectorOperationContext)

	// ExitEx_trueConst is called when exiting the ex_trueConst production.
	ExitEx_trueConst(c *Ex_trueConstContext)

	// ExitEx_falseConst is called when exiting the ex_falseConst production.
	ExitEx_falseConst(c *Ex_falseConstContext)

	// ExitEx_time_instant_literal is called when exiting the ex_time_instant_literal production.
	ExitEx_time_instant_literal(c *Ex_time_instant_literalContext)

	// ExitEx_iso_date_time is called when exiting the ex_iso_date_time production.
	ExitEx_iso_date_time(c *Ex_iso_date_timeContext)

	// ExitEx_iso_date_time_ymdhmsf is called when exiting the ex_iso_date_time_ymdhmsf production.
	ExitEx_iso_date_time_ymdhmsf(c *Ex_iso_date_time_ymdhmsfContext)

	// ExitEx_iso_date_time_ymdhms is called when exiting the ex_iso_date_time_ymdhms production.
	ExitEx_iso_date_time_ymdhms(c *Ex_iso_date_time_ymdhmsContext)

	// ExitEx_iso_date_time_ymdhm is called when exiting the ex_iso_date_time_ymdhm production.
	ExitEx_iso_date_time_ymdhm(c *Ex_iso_date_time_ymdhmContext)

	// ExitEx_iso_date_time_ymdh is called when exiting the ex_iso_date_time_ymdh production.
	ExitEx_iso_date_time_ymdh(c *Ex_iso_date_time_ymdhContext)

	// ExitEx_iso_date_time_ymd is called when exiting the ex_iso_date_time_ymd production.
	ExitEx_iso_date_time_ymd(c *Ex_iso_date_time_ymdContext)

	// ExitEx_iso_date_time_ym is called when exiting the ex_iso_date_time_ym production.
	ExitEx_iso_date_time_ym(c *Ex_iso_date_time_ymContext)

	// ExitEx_iso_date_time_y is called when exiting the ex_iso_date_time_y production.
	ExitEx_iso_date_time_y(c *Ex_iso_date_time_yContext)

	// ExitEx_year is called when exiting the ex_year production.
	ExitEx_year(c *Ex_yearContext)

	// ExitEx_month is called when exiting the ex_month production.
	ExitEx_month(c *Ex_monthContext)

	// ExitEx_day is called when exiting the ex_day production.
	ExitEx_day(c *Ex_dayContext)

	// ExitEx_hour is called when exiting the ex_hour production.
	ExitEx_hour(c *Ex_hourContext)

	// ExitEx_minutes is called when exiting the ex_minutes production.
	ExitEx_minutes(c *Ex_minutesContext)

	// ExitEx_seconds is called when exiting the ex_seconds production.
	ExitEx_seconds(c *Ex_secondsContext)

	// ExitEx_frac_sec is called when exiting the ex_frac_sec production.
	ExitEx_frac_sec(c *Ex_frac_secContext)

	// ExitEx_unix_timestamp is called when exiting the ex_unix_timestamp production.
	ExitEx_unix_timestamp(c *Ex_unix_timestampContext)

	// ExitEx_const_num_expression is called when exiting the ex_const_num_expression production.
	ExitEx_const_num_expression(c *Ex_const_num_expressionContext)

	// ExitEx_num_literal is called when exiting the ex_num_literal production.
	ExitEx_num_literal(c *Ex_num_literalContext)

	// ExitEx_duration is called when exiting the ex_duration production.
	ExitEx_duration(c *Ex_durationContext)

	// ExitEx_time_range is called when exiting the ex_time_range production.
	ExitEx_time_range(c *Ex_time_rangeContext)

	// ExitEx_subquery_range is called when exiting the ex_subquery_range production.
	ExitEx_subquery_range(c *Ex_subquery_rangeContext)

	// ExitVectorOperation is called when exiting the vectorOperation production.
	ExitVectorOperation(c *VectorOperationContext)

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

	// ExitSubqueryOp is called when exiting the subqueryOp production.
	ExitSubqueryOp(c *SubqueryOpContext)

	// ExitOffsetOp is called when exiting the offsetOp production.
	ExitOffsetOp(c *OffsetOpContext)

	// ExitVector is called when exiting the vector production.
	ExitVector(c *VectorContext)

	// ExitParens is called when exiting the parens production.
	ExitParens(c *ParensContext)

	// ExitInstantSelector is called when exiting the instantSelector production.
	ExitInstantSelector(c *InstantSelectorContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitLabelMatcherOperator is called when exiting the labelMatcherOperator production.
	ExitLabelMatcherOperator(c *LabelMatcherOperatorContext)

	// ExitLabelMatcherList is called when exiting the labelMatcherList production.
	ExitLabelMatcherList(c *LabelMatcherListContext)

	// ExitMatrixSelector is called when exiting the matrixSelector production.
	ExitMatrixSelector(c *MatrixSelectorContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

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

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitLabelNameList is called when exiting the labelNameList production.
	ExitLabelNameList(c *LabelNameListContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)
}

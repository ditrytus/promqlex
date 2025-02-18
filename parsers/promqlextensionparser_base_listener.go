// Code generated from PromQLExtensionParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLExtensionParser

import "github.com/antlr4-go/antlr/v4"

// BasePromQLExtensionParserListener is a complete listener for a parse tree produced by PromQLExtensionParser.
type BasePromQLExtensionParserListener struct{}

var _ PromQLExtensionParserListener = &BasePromQLExtensionParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromQLExtensionParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromQLExtensionParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromQLExtensionParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromQLExtensionParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPromqlx is called when production promqlx is entered.
func (s *BasePromQLExtensionParserListener) EnterPromqlx(ctx *PromqlxContext) {}

// ExitPromqlx is called when production promqlx is exited.
func (s *BasePromQLExtensionParserListener) ExitPromqlx(ctx *PromqlxContext) {}

// EnterEx_statement_list is called when production ex_statement_list is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_statement_list(ctx *Ex_statement_listContext) {}

// ExitEx_statement_list is called when production ex_statement_list is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_statement_list(ctx *Ex_statement_listContext) {}

// EnterEx_statement is called when production ex_statement is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_statement(ctx *Ex_statementContext) {}

// ExitEx_statement is called when production ex_statement is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_statement(ctx *Ex_statementContext) {}

// EnterEx_alias_def is called when production ex_alias_def is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_alias_def(ctx *Ex_alias_defContext) {}

// ExitEx_alias_def is called when production ex_alias_def is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_alias_def(ctx *Ex_alias_defContext) {}

// EnterEx_macro_def is called when production ex_macro_def is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_macro_def(ctx *Ex_macro_defContext) {}

// ExitEx_macro_def is called when production ex_macro_def is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_macro_def(ctx *Ex_macro_defContext) {}

// EnterEx_block is called when production ex_block is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_block(ctx *Ex_blockContext) {}

// ExitEx_block is called when production ex_block is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_block(ctx *Ex_blockContext) {}

// EnterEx_arg_list is called when production ex_arg_list is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_arg_list(ctx *Ex_arg_listContext) {}

// ExitEx_arg_list is called when production ex_arg_list is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_arg_list(ctx *Ex_arg_listContext) {}

// EnterEx_if_statement is called when production ex_if_statement is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_if_statement(ctx *Ex_if_statementContext) {}

// ExitEx_if_statement is called when production ex_if_statement is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_if_statement(ctx *Ex_if_statementContext) {}

// EnterEx_condition is called when production ex_condition is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_condition(ctx *Ex_conditionContext) {}

// ExitEx_condition is called when production ex_condition is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_condition(ctx *Ex_conditionContext) {}

// EnterEx_compareVectorOperation is called when production ex_compareVectorOperation is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_compareVectorOperation(ctx *Ex_compareVectorOperationContext) {
}

// ExitEx_compareVectorOperation is called when production ex_compareVectorOperation is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_compareVectorOperation(ctx *Ex_compareVectorOperationContext) {
}

// EnterEx_trueConst is called when production ex_trueConst is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_trueConst(ctx *Ex_trueConstContext) {}

// ExitEx_trueConst is called when production ex_trueConst is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_trueConst(ctx *Ex_trueConstContext) {}

// EnterEx_falseConst is called when production ex_falseConst is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_falseConst(ctx *Ex_falseConstContext) {}

// ExitEx_falseConst is called when production ex_falseConst is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_falseConst(ctx *Ex_falseConstContext) {}

// EnterEx_time_instant is called when production ex_time_instant is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_time_instant(ctx *Ex_time_instantContext) {}

// ExitEx_time_instant is called when production ex_time_instant is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_time_instant(ctx *Ex_time_instantContext) {}

// EnterEx_iso_date_time is called when production ex_iso_date_time is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time(ctx *Ex_iso_date_timeContext) {}

// ExitEx_iso_date_time is called when production ex_iso_date_time is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time(ctx *Ex_iso_date_timeContext) {}

// EnterEx_iso_date_time_ymdhmsf is called when production ex_iso_date_time_ymdhmsf is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ymdhmsf(ctx *Ex_iso_date_time_ymdhmsfContext) {
}

// ExitEx_iso_date_time_ymdhmsf is called when production ex_iso_date_time_ymdhmsf is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ymdhmsf(ctx *Ex_iso_date_time_ymdhmsfContext) {
}

// EnterEx_iso_date_time_ymdhms is called when production ex_iso_date_time_ymdhms is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ymdhms(ctx *Ex_iso_date_time_ymdhmsContext) {
}

// ExitEx_iso_date_time_ymdhms is called when production ex_iso_date_time_ymdhms is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ymdhms(ctx *Ex_iso_date_time_ymdhmsContext) {
}

// EnterEx_iso_date_time_ymdhm is called when production ex_iso_date_time_ymdhm is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ymdhm(ctx *Ex_iso_date_time_ymdhmContext) {
}

// ExitEx_iso_date_time_ymdhm is called when production ex_iso_date_time_ymdhm is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ymdhm(ctx *Ex_iso_date_time_ymdhmContext) {
}

// EnterEx_iso_date_time_ymdh is called when production ex_iso_date_time_ymdh is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ymdh(ctx *Ex_iso_date_time_ymdhContext) {
}

// ExitEx_iso_date_time_ymdh is called when production ex_iso_date_time_ymdh is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ymdh(ctx *Ex_iso_date_time_ymdhContext) {
}

// EnterEx_iso_date_time_ymd is called when production ex_iso_date_time_ymd is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ymd(ctx *Ex_iso_date_time_ymdContext) {
}

// ExitEx_iso_date_time_ymd is called when production ex_iso_date_time_ymd is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ymd(ctx *Ex_iso_date_time_ymdContext) {
}

// EnterEx_iso_date_time_ym is called when production ex_iso_date_time_ym is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_ym(ctx *Ex_iso_date_time_ymContext) {
}

// ExitEx_iso_date_time_ym is called when production ex_iso_date_time_ym is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_ym(ctx *Ex_iso_date_time_ymContext) {
}

// EnterEx_iso_date_time_y is called when production ex_iso_date_time_y is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_iso_date_time_y(ctx *Ex_iso_date_time_yContext) {}

// ExitEx_iso_date_time_y is called when production ex_iso_date_time_y is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_iso_date_time_y(ctx *Ex_iso_date_time_yContext) {}

// EnterEx_year is called when production ex_year is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_year(ctx *Ex_yearContext) {}

// ExitEx_year is called when production ex_year is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_year(ctx *Ex_yearContext) {}

// EnterEx_month is called when production ex_month is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_month(ctx *Ex_monthContext) {}

// ExitEx_month is called when production ex_month is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_month(ctx *Ex_monthContext) {}

// EnterEx_day is called when production ex_day is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_day(ctx *Ex_dayContext) {}

// ExitEx_day is called when production ex_day is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_day(ctx *Ex_dayContext) {}

// EnterEx_hour is called when production ex_hour is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_hour(ctx *Ex_hourContext) {}

// ExitEx_hour is called when production ex_hour is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_hour(ctx *Ex_hourContext) {}

// EnterEx_minutes is called when production ex_minutes is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_minutes(ctx *Ex_minutesContext) {}

// ExitEx_minutes is called when production ex_minutes is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_minutes(ctx *Ex_minutesContext) {}

// EnterEx_seconds is called when production ex_seconds is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_seconds(ctx *Ex_secondsContext) {}

// ExitEx_seconds is called when production ex_seconds is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_seconds(ctx *Ex_secondsContext) {}

// EnterEx_frac_sec is called when production ex_frac_sec is entered.
func (s *BasePromQLExtensionParserListener) EnterEx_frac_sec(ctx *Ex_frac_secContext) {}

// ExitEx_frac_sec is called when production ex_frac_sec is exited.
func (s *BasePromQLExtensionParserListener) ExitEx_frac_sec(ctx *Ex_frac_secContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePromQLExtensionParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePromQLExtensionParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVectorOperation is called when production vectorOperation is entered.
func (s *BasePromQLExtensionParserListener) EnterVectorOperation(ctx *VectorOperationContext) {}

// ExitVectorOperation is called when production vectorOperation is exited.
func (s *BasePromQLExtensionParserListener) ExitVectorOperation(ctx *VectorOperationContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BasePromQLExtensionParserListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BasePromQLExtensionParserListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterPowOp is called when production powOp is entered.
func (s *BasePromQLExtensionParserListener) EnterPowOp(ctx *PowOpContext) {}

// ExitPowOp is called when production powOp is exited.
func (s *BasePromQLExtensionParserListener) ExitPowOp(ctx *PowOpContext) {}

// EnterMultOp is called when production multOp is entered.
func (s *BasePromQLExtensionParserListener) EnterMultOp(ctx *MultOpContext) {}

// ExitMultOp is called when production multOp is exited.
func (s *BasePromQLExtensionParserListener) ExitMultOp(ctx *MultOpContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BasePromQLExtensionParserListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BasePromQLExtensionParserListener) ExitAddOp(ctx *AddOpContext) {}

// EnterCompareOp is called when production compareOp is entered.
func (s *BasePromQLExtensionParserListener) EnterCompareOp(ctx *CompareOpContext) {}

// ExitCompareOp is called when production compareOp is exited.
func (s *BasePromQLExtensionParserListener) ExitCompareOp(ctx *CompareOpContext) {}

// EnterAndUnlessOp is called when production andUnlessOp is entered.
func (s *BasePromQLExtensionParserListener) EnterAndUnlessOp(ctx *AndUnlessOpContext) {}

// ExitAndUnlessOp is called when production andUnlessOp is exited.
func (s *BasePromQLExtensionParserListener) ExitAndUnlessOp(ctx *AndUnlessOpContext) {}

// EnterOrOp is called when production orOp is entered.
func (s *BasePromQLExtensionParserListener) EnterOrOp(ctx *OrOpContext) {}

// ExitOrOp is called when production orOp is exited.
func (s *BasePromQLExtensionParserListener) ExitOrOp(ctx *OrOpContext) {}

// EnterVectorMatchOp is called when production vectorMatchOp is entered.
func (s *BasePromQLExtensionParserListener) EnterVectorMatchOp(ctx *VectorMatchOpContext) {}

// ExitVectorMatchOp is called when production vectorMatchOp is exited.
func (s *BasePromQLExtensionParserListener) ExitVectorMatchOp(ctx *VectorMatchOpContext) {}

// EnterSubqueryOp is called when production subqueryOp is entered.
func (s *BasePromQLExtensionParserListener) EnterSubqueryOp(ctx *SubqueryOpContext) {}

// ExitSubqueryOp is called when production subqueryOp is exited.
func (s *BasePromQLExtensionParserListener) ExitSubqueryOp(ctx *SubqueryOpContext) {}

// EnterOffsetOp is called when production offsetOp is entered.
func (s *BasePromQLExtensionParserListener) EnterOffsetOp(ctx *OffsetOpContext) {}

// ExitOffsetOp is called when production offsetOp is exited.
func (s *BasePromQLExtensionParserListener) ExitOffsetOp(ctx *OffsetOpContext) {}

// EnterVector is called when production vector is entered.
func (s *BasePromQLExtensionParserListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BasePromQLExtensionParserListener) ExitVector(ctx *VectorContext) {}

// EnterParens is called when production parens is entered.
func (s *BasePromQLExtensionParserListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BasePromQLExtensionParserListener) ExitParens(ctx *ParensContext) {}

// EnterInstantSelector is called when production instantSelector is entered.
func (s *BasePromQLExtensionParserListener) EnterInstantSelector(ctx *InstantSelectorContext) {}

// ExitInstantSelector is called when production instantSelector is exited.
func (s *BasePromQLExtensionParserListener) ExitInstantSelector(ctx *InstantSelectorContext) {}

// EnterLabelMatcher is called when production labelMatcher is entered.
func (s *BasePromQLExtensionParserListener) EnterLabelMatcher(ctx *LabelMatcherContext) {}

// ExitLabelMatcher is called when production labelMatcher is exited.
func (s *BasePromQLExtensionParserListener) ExitLabelMatcher(ctx *LabelMatcherContext) {}

// EnterLabelMatcherOperator is called when production labelMatcherOperator is entered.
func (s *BasePromQLExtensionParserListener) EnterLabelMatcherOperator(ctx *LabelMatcherOperatorContext) {
}

// ExitLabelMatcherOperator is called when production labelMatcherOperator is exited.
func (s *BasePromQLExtensionParserListener) ExitLabelMatcherOperator(ctx *LabelMatcherOperatorContext) {
}

// EnterLabelMatcherList is called when production labelMatcherList is entered.
func (s *BasePromQLExtensionParserListener) EnterLabelMatcherList(ctx *LabelMatcherListContext) {}

// ExitLabelMatcherList is called when production labelMatcherList is exited.
func (s *BasePromQLExtensionParserListener) ExitLabelMatcherList(ctx *LabelMatcherListContext) {}

// EnterMatrixSelector is called when production matrixSelector is entered.
func (s *BasePromQLExtensionParserListener) EnterMatrixSelector(ctx *MatrixSelectorContext) {}

// ExitMatrixSelector is called when production matrixSelector is exited.
func (s *BasePromQLExtensionParserListener) ExitMatrixSelector(ctx *MatrixSelectorContext) {}

// EnterOffset is called when production offset is entered.
func (s *BasePromQLExtensionParserListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BasePromQLExtensionParserListener) ExitOffset(ctx *OffsetContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BasePromQLExtensionParserListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BasePromQLExtensionParserListener) ExitFunction_(ctx *Function_Context) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePromQLExtensionParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePromQLExtensionParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasePromQLExtensionParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasePromQLExtensionParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BasePromQLExtensionParserListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BasePromQLExtensionParserListener) ExitAggregation(ctx *AggregationContext) {}

// EnterBy is called when production by is entered.
func (s *BasePromQLExtensionParserListener) EnterBy(ctx *ByContext) {}

// ExitBy is called when production by is exited.
func (s *BasePromQLExtensionParserListener) ExitBy(ctx *ByContext) {}

// EnterWithout is called when production without is entered.
func (s *BasePromQLExtensionParserListener) EnterWithout(ctx *WithoutContext) {}

// ExitWithout is called when production without is exited.
func (s *BasePromQLExtensionParserListener) ExitWithout(ctx *WithoutContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BasePromQLExtensionParserListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BasePromQLExtensionParserListener) ExitGrouping(ctx *GroupingContext) {}

// EnterOn_ is called when production on_ is entered.
func (s *BasePromQLExtensionParserListener) EnterOn_(ctx *On_Context) {}

// ExitOn_ is called when production on_ is exited.
func (s *BasePromQLExtensionParserListener) ExitOn_(ctx *On_Context) {}

// EnterIgnoring is called when production ignoring is entered.
func (s *BasePromQLExtensionParserListener) EnterIgnoring(ctx *IgnoringContext) {}

// ExitIgnoring is called when production ignoring is exited.
func (s *BasePromQLExtensionParserListener) ExitIgnoring(ctx *IgnoringContext) {}

// EnterGroupLeft is called when production groupLeft is entered.
func (s *BasePromQLExtensionParserListener) EnterGroupLeft(ctx *GroupLeftContext) {}

// ExitGroupLeft is called when production groupLeft is exited.
func (s *BasePromQLExtensionParserListener) ExitGroupLeft(ctx *GroupLeftContext) {}

// EnterGroupRight is called when production groupRight is entered.
func (s *BasePromQLExtensionParserListener) EnterGroupRight(ctx *GroupRightContext) {}

// ExitGroupRight is called when production groupRight is exited.
func (s *BasePromQLExtensionParserListener) ExitGroupRight(ctx *GroupRightContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BasePromQLExtensionParserListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BasePromQLExtensionParserListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterLabelNameList is called when production labelNameList is entered.
func (s *BasePromQLExtensionParserListener) EnterLabelNameList(ctx *LabelNameListContext) {}

// ExitLabelNameList is called when production labelNameList is exited.
func (s *BasePromQLExtensionParserListener) ExitLabelNameList(ctx *LabelNameListContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BasePromQLExtensionParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BasePromQLExtensionParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePromQLExtensionParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePromQLExtensionParserListener) ExitLiteral(ctx *LiteralContext) {}

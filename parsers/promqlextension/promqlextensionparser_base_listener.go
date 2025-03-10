// Code generated from PromQLExtensionParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlextension // PromQLExtensionParser
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

// EnterStatement is called when production statement is entered.
func (s *BasePromQLExtensionParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePromQLExtensionParserListener) ExitStatement(ctx *StatementContext) {}

// EnterAlias_def is called when production alias_def is entered.
func (s *BasePromQLExtensionParserListener) EnterAlias_def(ctx *Alias_defContext) {}

// ExitAlias_def is called when production alias_def is exited.
func (s *BasePromQLExtensionParserListener) ExitAlias_def(ctx *Alias_defContext) {}

// EnterAlias_call is called when production alias_call is entered.
func (s *BasePromQLExtensionParserListener) EnterAlias_call(ctx *Alias_callContext) {}

// ExitAlias_call is called when production alias_call is exited.
func (s *BasePromQLExtensionParserListener) ExitAlias_call(ctx *Alias_callContext) {}

// EnterMacro_def is called when production macro_def is entered.
func (s *BasePromQLExtensionParserListener) EnterMacro_def(ctx *Macro_defContext) {}

// ExitMacro_def is called when production macro_def is exited.
func (s *BasePromQLExtensionParserListener) ExitMacro_def(ctx *Macro_defContext) {}

// EnterMacro_call is called when production macro_call is entered.
func (s *BasePromQLExtensionParserListener) EnterMacro_call(ctx *Macro_callContext) {}

// ExitMacro_call is called when production macro_call is exited.
func (s *BasePromQLExtensionParserListener) ExitMacro_call(ctx *Macro_callContext) {}

// EnterArgs_decl is called when production args_decl is entered.
func (s *BasePromQLExtensionParserListener) EnterArgs_decl(ctx *Args_declContext) {}

// ExitArgs_decl is called when production args_decl is exited.
func (s *BasePromQLExtensionParserListener) ExitArgs_decl(ctx *Args_declContext) {}

// EnterArg_name is called when production arg_name is entered.
func (s *BasePromQLExtensionParserListener) EnterArg_name(ctx *Arg_nameContext) {}

// ExitArg_name is called when production arg_name is exited.
func (s *BasePromQLExtensionParserListener) ExitArg_name(ctx *Arg_nameContext) {}

// EnterStatement_block is called when production statement_block is entered.
func (s *BasePromQLExtensionParserListener) EnterStatement_block(ctx *Statement_blockContext) {}

// ExitStatement_block is called when production statement_block is exited.
func (s *BasePromQLExtensionParserListener) ExitStatement_block(ctx *Statement_blockContext) {}

// EnterArg_list is called when production arg_list is entered.
func (s *BasePromQLExtensionParserListener) EnterArg_list(ctx *Arg_listContext) {}

// ExitArg_list is called when production arg_list is exited.
func (s *BasePromQLExtensionParserListener) ExitArg_list(ctx *Arg_listContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BasePromQLExtensionParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BasePromQLExtensionParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasePromQLExtensionParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasePromQLExtensionParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterCompareVectorOperation is called when production compareVectorOperation is entered.
func (s *BasePromQLExtensionParserListener) EnterCompareVectorOperation(ctx *CompareVectorOperationContext) {
}

// ExitCompareVectorOperation is called when production compareVectorOperation is exited.
func (s *BasePromQLExtensionParserListener) ExitCompareVectorOperation(ctx *CompareVectorOperationContext) {
}

// EnterTrueConst is called when production trueConst is entered.
func (s *BasePromQLExtensionParserListener) EnterTrueConst(ctx *TrueConstContext) {}

// ExitTrueConst is called when production trueConst is exited.
func (s *BasePromQLExtensionParserListener) ExitTrueConst(ctx *TrueConstContext) {}

// EnterFalseConst is called when production falseConst is entered.
func (s *BasePromQLExtensionParserListener) EnterFalseConst(ctx *FalseConstContext) {}

// ExitFalseConst is called when production falseConst is exited.
func (s *BasePromQLExtensionParserListener) ExitFalseConst(ctx *FalseConstContext) {}

// EnterTime_instant_literal is called when production time_instant_literal is entered.
func (s *BasePromQLExtensionParserListener) EnterTime_instant_literal(ctx *Time_instant_literalContext) {
}

// ExitTime_instant_literal is called when production time_instant_literal is exited.
func (s *BasePromQLExtensionParserListener) ExitTime_instant_literal(ctx *Time_instant_literalContext) {
}

// EnterIso_date_time is called when production iso_date_time is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time(ctx *Iso_date_timeContext) {}

// ExitIso_date_time is called when production iso_date_time is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time(ctx *Iso_date_timeContext) {}

// EnterIso_date_time_ymdhmsf is called when production iso_date_time_ymdhmsf is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ymdhmsf(ctx *Iso_date_time_ymdhmsfContext) {
}

// ExitIso_date_time_ymdhmsf is called when production iso_date_time_ymdhmsf is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ymdhmsf(ctx *Iso_date_time_ymdhmsfContext) {
}

// EnterIso_date_time_ymdhms is called when production iso_date_time_ymdhms is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ymdhms(ctx *Iso_date_time_ymdhmsContext) {
}

// ExitIso_date_time_ymdhms is called when production iso_date_time_ymdhms is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ymdhms(ctx *Iso_date_time_ymdhmsContext) {
}

// EnterIso_date_time_ymdhm is called when production iso_date_time_ymdhm is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ymdhm(ctx *Iso_date_time_ymdhmContext) {
}

// ExitIso_date_time_ymdhm is called when production iso_date_time_ymdhm is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ymdhm(ctx *Iso_date_time_ymdhmContext) {
}

// EnterIso_date_time_ymdh is called when production iso_date_time_ymdh is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ymdh(ctx *Iso_date_time_ymdhContext) {}

// ExitIso_date_time_ymdh is called when production iso_date_time_ymdh is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ymdh(ctx *Iso_date_time_ymdhContext) {}

// EnterIso_date_time_ymd is called when production iso_date_time_ymd is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ymd(ctx *Iso_date_time_ymdContext) {}

// ExitIso_date_time_ymd is called when production iso_date_time_ymd is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ymd(ctx *Iso_date_time_ymdContext) {}

// EnterIso_date_time_ym is called when production iso_date_time_ym is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_ym(ctx *Iso_date_time_ymContext) {}

// ExitIso_date_time_ym is called when production iso_date_time_ym is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_ym(ctx *Iso_date_time_ymContext) {}

// EnterIso_date_time_y is called when production iso_date_time_y is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_date_time_y(ctx *Iso_date_time_yContext) {}

// ExitIso_date_time_y is called when production iso_date_time_y is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_date_time_y(ctx *Iso_date_time_yContext) {}

// EnterIso_year is called when production iso_year is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_year(ctx *Iso_yearContext) {}

// ExitIso_year is called when production iso_year is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_year(ctx *Iso_yearContext) {}

// EnterIso_month is called when production iso_month is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_month(ctx *Iso_monthContext) {}

// ExitIso_month is called when production iso_month is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_month(ctx *Iso_monthContext) {}

// EnterIso_day is called when production iso_day is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_day(ctx *Iso_dayContext) {}

// ExitIso_day is called when production iso_day is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_day(ctx *Iso_dayContext) {}

// EnterIso_hour is called when production iso_hour is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_hour(ctx *Iso_hourContext) {}

// ExitIso_hour is called when production iso_hour is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_hour(ctx *Iso_hourContext) {}

// EnterIso_minutes is called when production iso_minutes is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_minutes(ctx *Iso_minutesContext) {}

// ExitIso_minutes is called when production iso_minutes is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_minutes(ctx *Iso_minutesContext) {}

// EnterIso_seconds is called when production iso_seconds is entered.
func (s *BasePromQLExtensionParserListener) EnterIso_seconds(ctx *Iso_secondsContext) {}

// ExitIso_seconds is called when production iso_seconds is exited.
func (s *BasePromQLExtensionParserListener) ExitIso_seconds(ctx *Iso_secondsContext) {}

// EnterIs_frac_sec is called when production is_frac_sec is entered.
func (s *BasePromQLExtensionParserListener) EnterIs_frac_sec(ctx *Is_frac_secContext) {}

// ExitIs_frac_sec is called when production is_frac_sec is exited.
func (s *BasePromQLExtensionParserListener) ExitIs_frac_sec(ctx *Is_frac_secContext) {}

// EnterUnix_timestamp is called when production unix_timestamp is entered.
func (s *BasePromQLExtensionParserListener) EnterUnix_timestamp(ctx *Unix_timestampContext) {}

// ExitUnix_timestamp is called when production unix_timestamp is exited.
func (s *BasePromQLExtensionParserListener) ExitUnix_timestamp(ctx *Unix_timestampContext) {}

// EnterConst_num_expression is called when production const_num_expression is entered.
func (s *BasePromQLExtensionParserListener) EnterConst_num_expression(ctx *Const_num_expressionContext) {
}

// ExitConst_num_expression is called when production const_num_expression is exited.
func (s *BasePromQLExtensionParserListener) ExitConst_num_expression(ctx *Const_num_expressionContext) {
}

// EnterNum_literal is called when production num_literal is entered.
func (s *BasePromQLExtensionParserListener) EnterNum_literal(ctx *Num_literalContext) {}

// ExitNum_literal is called when production num_literal is exited.
func (s *BasePromQLExtensionParserListener) ExitNum_literal(ctx *Num_literalContext) {}

// EnterDuration is called when production duration is entered.
func (s *BasePromQLExtensionParserListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BasePromQLExtensionParserListener) ExitDuration(ctx *DurationContext) {}

// EnterTime_range is called when production time_range is entered.
func (s *BasePromQLExtensionParserListener) EnterTime_range(ctx *Time_rangeContext) {}

// ExitTime_range is called when production time_range is exited.
func (s *BasePromQLExtensionParserListener) ExitTime_range(ctx *Time_rangeContext) {}

// EnterSubquery_range is called when production subquery_range is entered.
func (s *BasePromQLExtensionParserListener) EnterSubquery_range(ctx *Subquery_rangeContext) {}

// ExitSubquery_range is called when production subquery_range is exited.
func (s *BasePromQLExtensionParserListener) ExitSubquery_range(ctx *Subquery_rangeContext) {}

// EnterVectorOperation is called when production vectorOperation is entered.
func (s *BasePromQLExtensionParserListener) EnterVectorOperation(ctx *VectorOperationContext) {}

// ExitVectorOperation is called when production vectorOperation is exited.
func (s *BasePromQLExtensionParserListener) ExitVectorOperation(ctx *VectorOperationContext) {}

// EnterSubqueryOp is called when production subqueryOp is entered.
func (s *BasePromQLExtensionParserListener) EnterSubqueryOp(ctx *SubqueryOpContext) {}

// ExitSubqueryOp is called when production subqueryOp is exited.
func (s *BasePromQLExtensionParserListener) ExitSubqueryOp(ctx *SubqueryOpContext) {}

// EnterOffsetOp is called when production offsetOp is entered.
func (s *BasePromQLExtensionParserListener) EnterOffsetOp(ctx *OffsetOpContext) {}

// ExitOffsetOp is called when production offsetOp is exited.
func (s *BasePromQLExtensionParserListener) ExitOffsetOp(ctx *OffsetOpContext) {}

// EnterMatrixSelector is called when production matrixSelector is entered.
func (s *BasePromQLExtensionParserListener) EnterMatrixSelector(ctx *MatrixSelectorContext) {}

// ExitMatrixSelector is called when production matrixSelector is exited.
func (s *BasePromQLExtensionParserListener) ExitMatrixSelector(ctx *MatrixSelectorContext) {}

// EnterOffset is called when production offset is entered.
func (s *BasePromQLExtensionParserListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BasePromQLExtensionParserListener) ExitOffset(ctx *OffsetContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePromQLExtensionParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePromQLExtensionParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePromQLExtensionParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePromQLExtensionParserListener) ExitExpression(ctx *ExpressionContext) {}

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

// EnterString is called when production string is entered.
func (s *BasePromQLExtensionParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasePromQLExtensionParserListener) ExitString(ctx *StringContext) {}

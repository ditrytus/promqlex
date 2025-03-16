// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PromQLExParser.
type PromQLExParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PromQLExParser#promqlx.
	VisitPromqlx(ctx *PromqlxContext) interface{}

	// Visit a parse tree produced by PromQLExParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by PromQLExParser#alias_def.
	VisitAlias_def(ctx *Alias_defContext) interface{}

	// Visit a parse tree produced by PromQLExParser#alias_call.
	VisitAlias_call(ctx *Alias_callContext) interface{}

	// Visit a parse tree produced by PromQLExParser#macro_def.
	VisitMacro_def(ctx *Macro_defContext) interface{}

	// Visit a parse tree produced by PromQLExParser#macro_call.
	VisitMacro_call(ctx *Macro_callContext) interface{}

	// Visit a parse tree produced by PromQLExParser#args_decl.
	VisitArgs_decl(ctx *Args_declContext) interface{}

	// Visit a parse tree produced by PromQLExParser#arg_name.
	VisitArg_name(ctx *Arg_nameContext) interface{}

	// Visit a parse tree produced by PromQLExParser#statement_block.
	VisitStatement_block(ctx *Statement_blockContext) interface{}

	// Visit a parse tree produced by PromQLExParser#arg_list.
	VisitArg_list(ctx *Arg_listContext) interface{}

	// Visit a parse tree produced by PromQLExParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by PromQLExParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by PromQLExParser#compareVectorOperation.
	VisitCompareVectorOperation(ctx *CompareVectorOperationContext) interface{}

	// Visit a parse tree produced by PromQLExParser#trueConst.
	VisitTrueConst(ctx *TrueConstContext) interface{}

	// Visit a parse tree produced by PromQLExParser#falseConst.
	VisitFalseConst(ctx *FalseConstContext) interface{}

	// Visit a parse tree produced by PromQLExParser#time_instant_literal.
	VisitTime_instant_literal(ctx *Time_instant_literalContext) interface{}

	// Visit a parse tree produced by PromQLExParser#iso_date_time.
	VisitIso_date_time(ctx *Iso_date_timeContext) interface{}

	// Visit a parse tree produced by PromQLExParser#unix_timestamp.
	VisitUnix_timestamp(ctx *Unix_timestampContext) interface{}

	// Visit a parse tree produced by PromQLExParser#const_num_expression.
	VisitConst_num_expression(ctx *Const_num_expressionContext) interface{}

	// Visit a parse tree produced by PromQLExParser#num_literal.
	VisitNum_literal(ctx *Num_literalContext) interface{}

	// Visit a parse tree produced by PromQLExParser#duration.
	VisitDuration(ctx *DurationContext) interface{}

	// Visit a parse tree produced by PromQLExParser#time_range.
	VisitTime_range(ctx *Time_rangeContext) interface{}

	// Visit a parse tree produced by PromQLExParser#subquery_range.
	VisitSubquery_range(ctx *Subquery_rangeContext) interface{}

	// Visit a parse tree produced by PromQLExParser#vectorOperation.
	VisitVectorOperation(ctx *VectorOperationContext) interface{}

	// Visit a parse tree produced by PromQLExParser#subqueryOp.
	VisitSubqueryOp(ctx *SubqueryOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#offsetOp.
	VisitOffsetOp(ctx *OffsetOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#matrixSelector.
	VisitMatrixSelector(ctx *MatrixSelectorContext) interface{}

	// Visit a parse tree produced by PromQLExParser#offset.
	VisitOffset(ctx *OffsetContext) interface{}

	// Visit a parse tree produced by PromQLExParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PromQLExParser#instantSelector.
	VisitInstantSelector(ctx *InstantSelectorContext) interface{}

	// Visit a parse tree produced by PromQLExParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by PromQLExParser#metric_name.
	VisitMetric_name(ctx *Metric_nameContext) interface{}

	// Visit a parse tree produced by PromQLExParser#at_modifier_timestamp.
	VisitAt_modifier_timestamp(ctx *At_modifier_timestampContext) interface{}

	// Visit a parse tree produced by PromQLExParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PromQLExParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#powOp.
	VisitPowOp(ctx *PowOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#multOp.
	VisitMultOp(ctx *MultOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#compareOp.
	VisitCompareOp(ctx *CompareOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#andUnlessOp.
	VisitAndUnlessOp(ctx *AndUnlessOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#orOp.
	VisitOrOp(ctx *OrOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#vectorMatchOp.
	VisitVectorMatchOp(ctx *VectorMatchOpContext) interface{}

	// Visit a parse tree produced by PromQLExParser#vector.
	VisitVector(ctx *VectorContext) interface{}

	// Visit a parse tree produced by PromQLExParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by PromQLExParser#labelMatcher.
	VisitLabelMatcher(ctx *LabelMatcherContext) interface{}

	// Visit a parse tree produced by PromQLExParser#labelMatcherOperator.
	VisitLabelMatcherOperator(ctx *LabelMatcherOperatorContext) interface{}

	// Visit a parse tree produced by PromQLExParser#labelMatcherList.
	VisitLabelMatcherList(ctx *LabelMatcherListContext) interface{}

	// Visit a parse tree produced by PromQLExParser#function_.
	VisitFunction_(ctx *Function_Context) interface{}

	// Visit a parse tree produced by PromQLExParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by PromQLExParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by PromQLExParser#aggregation.
	VisitAggregation(ctx *AggregationContext) interface{}

	// Visit a parse tree produced by PromQLExParser#by.
	VisitBy(ctx *ByContext) interface{}

	// Visit a parse tree produced by PromQLExParser#without.
	VisitWithout(ctx *WithoutContext) interface{}

	// Visit a parse tree produced by PromQLExParser#grouping.
	VisitGrouping(ctx *GroupingContext) interface{}

	// Visit a parse tree produced by PromQLExParser#on_.
	VisitOn_(ctx *On_Context) interface{}

	// Visit a parse tree produced by PromQLExParser#ignoring.
	VisitIgnoring(ctx *IgnoringContext) interface{}

	// Visit a parse tree produced by PromQLExParser#groupLeft.
	VisitGroupLeft(ctx *GroupLeftContext) interface{}

	// Visit a parse tree produced by PromQLExParser#groupRight.
	VisitGroupRight(ctx *GroupRightContext) interface{}

	// Visit a parse tree produced by PromQLExParser#labelNameList.
	VisitLabelNameList(ctx *LabelNameListContext) interface{}

	// Visit a parse tree produced by PromQLExParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by PromQLExParser#string.
	VisitString(ctx *StringContext) interface{}
}

// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
import "github.com/antlr4-go/antlr/v4"

type BasePromQLExParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePromQLExParserVisitor) VisitPromqlx(ctx *PromqlxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAlias_def(ctx *Alias_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAlias_call(ctx *Alias_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitMacro_def(ctx *Macro_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitMacro_call(ctx *Macro_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitArgs_decl(ctx *Args_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitArg_name(ctx *Arg_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitStatement_block(ctx *Statement_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitArg_list(ctx *Arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitCompareVectorOperation(ctx *CompareVectorOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitTrueConst(ctx *TrueConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitFalseConst(ctx *FalseConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitTime_instant_literal(ctx *Time_instant_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitIso_date_time(ctx *Iso_date_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitUnix_timestamp(ctx *Unix_timestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitConst_num_expression(ctx *Const_num_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitNum_literal(ctx *Num_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitDuration(ctx *DurationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitTime_range(ctx *Time_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitSubquery_range(ctx *Subquery_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitVectorOperation(ctx *VectorOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitSubqueryOp(ctx *SubqueryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitOffsetOp(ctx *OffsetOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitMatrixSelector(ctx *MatrixSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitOffset(ctx *OffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitInstantSelector(ctx *InstantSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitMetric_name(ctx *Metric_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAt_modifier_timestamp(ctx *At_modifier_timestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitPowOp(ctx *PowOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitMultOp(ctx *MultOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitCompareOp(ctx *CompareOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAndUnlessOp(ctx *AndUnlessOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitOrOp(ctx *OrOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitVectorMatchOp(ctx *VectorMatchOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitVector(ctx *VectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLabelMatcher(ctx *LabelMatcherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLabelMatcherOperator(ctx *LabelMatcherOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLabelMatcherList(ctx *LabelMatcherListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitFunction_(ctx *Function_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitAggregation(ctx *AggregationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitBy(ctx *ByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitWithout(ctx *WithoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitGrouping(ctx *GroupingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitOn_(ctx *On_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitIgnoring(ctx *IgnoringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitGroupLeft(ctx *GroupLeftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitGroupRight(ctx *GroupRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitLabelNameList(ctx *LabelNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLExParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

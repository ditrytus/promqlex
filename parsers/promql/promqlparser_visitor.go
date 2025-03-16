// Code generated from PromQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promql // PromQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PromQLParser.
type PromQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PromQLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PromQLParser#vectorOperation.
	VisitVectorOperation(ctx *VectorOperationContext) interface{}

	// Visit a parse tree produced by PromQLParser#at_modifier_timestamp.
	VisitAt_modifier_timestamp(ctx *At_modifier_timestampContext) interface{}

	// Visit a parse tree produced by PromQLParser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#powOp.
	VisitPowOp(ctx *PowOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#multOp.
	VisitMultOp(ctx *MultOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#compareOp.
	VisitCompareOp(ctx *CompareOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#andUnlessOp.
	VisitAndUnlessOp(ctx *AndUnlessOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#orOp.
	VisitOrOp(ctx *OrOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#vectorMatchOp.
	VisitVectorMatchOp(ctx *VectorMatchOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#subqueryOp.
	VisitSubqueryOp(ctx *SubqueryOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#offsetOp.
	VisitOffsetOp(ctx *OffsetOpContext) interface{}

	// Visit a parse tree produced by PromQLParser#vector.
	VisitVector(ctx *VectorContext) interface{}

	// Visit a parse tree produced by PromQLParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by PromQLParser#instantSelector.
	VisitInstantSelector(ctx *InstantSelectorContext) interface{}

	// Visit a parse tree produced by PromQLParser#labelMatcher.
	VisitLabelMatcher(ctx *LabelMatcherContext) interface{}

	// Visit a parse tree produced by PromQLParser#labelMatcherOperator.
	VisitLabelMatcherOperator(ctx *LabelMatcherOperatorContext) interface{}

	// Visit a parse tree produced by PromQLParser#labelMatcherList.
	VisitLabelMatcherList(ctx *LabelMatcherListContext) interface{}

	// Visit a parse tree produced by PromQLParser#matrixSelector.
	VisitMatrixSelector(ctx *MatrixSelectorContext) interface{}

	// Visit a parse tree produced by PromQLParser#offset.
	VisitOffset(ctx *OffsetContext) interface{}

	// Visit a parse tree produced by PromQLParser#function_.
	VisitFunction_(ctx *Function_Context) interface{}

	// Visit a parse tree produced by PromQLParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by PromQLParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by PromQLParser#aggregation.
	VisitAggregation(ctx *AggregationContext) interface{}

	// Visit a parse tree produced by PromQLParser#by.
	VisitBy(ctx *ByContext) interface{}

	// Visit a parse tree produced by PromQLParser#without.
	VisitWithout(ctx *WithoutContext) interface{}

	// Visit a parse tree produced by PromQLParser#grouping.
	VisitGrouping(ctx *GroupingContext) interface{}

	// Visit a parse tree produced by PromQLParser#on_.
	VisitOn_(ctx *On_Context) interface{}

	// Visit a parse tree produced by PromQLParser#ignoring.
	VisitIgnoring(ctx *IgnoringContext) interface{}

	// Visit a parse tree produced by PromQLParser#groupLeft.
	VisitGroupLeft(ctx *GroupLeftContext) interface{}

	// Visit a parse tree produced by PromQLParser#groupRight.
	VisitGroupRight(ctx *GroupRightContext) interface{}

	// Visit a parse tree produced by PromQLParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by PromQLParser#labelNameList.
	VisitLabelNameList(ctx *LabelNameListContext) interface{}

	// Visit a parse tree produced by PromQLParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by PromQLParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by PromQLParser#string.
	VisitString(ctx *StringContext) interface{}
}

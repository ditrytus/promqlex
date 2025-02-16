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

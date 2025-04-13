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

// EnterState_AliasDef is called when production State_AliasDef is entered.
func (s *BasePromQLExParserListener) EnterState_AliasDef(ctx *State_AliasDefContext) {}

// ExitState_AliasDef is called when production State_AliasDef is exited.
func (s *BasePromQLExParserListener) ExitState_AliasDef(ctx *State_AliasDefContext) {}

// EnterState_MacroDef is called when production State_MacroDef is entered.
func (s *BasePromQLExParserListener) EnterState_MacroDef(ctx *State_MacroDefContext) {}

// ExitState_MacroDef is called when production State_MacroDef is exited.
func (s *BasePromQLExParserListener) ExitState_MacroDef(ctx *State_MacroDefContext) {}

// EnterState_IfStatement is called when production State_IfStatement is entered.
func (s *BasePromQLExParserListener) EnterState_IfStatement(ctx *State_IfStatementContext) {}

// ExitState_IfStatement is called when production State_IfStatement is exited.
func (s *BasePromQLExParserListener) ExitState_IfStatement(ctx *State_IfStatementContext) {}

// EnterState_VectorOperation is called when production State_VectorOperation is entered.
func (s *BasePromQLExParserListener) EnterState_VectorOperation(ctx *State_VectorOperationContext) {}

// ExitState_VectorOperation is called when production State_VectorOperation is exited.
func (s *BasePromQLExParserListener) ExitState_VectorOperation(ctx *State_VectorOperationContext) {}

// EnterAlias_def is called when production alias_def is entered.
func (s *BasePromQLExParserListener) EnterAlias_def(ctx *Alias_defContext) {}

// ExitAlias_def is called when production alias_def is exited.
func (s *BasePromQLExParserListener) ExitAlias_def(ctx *Alias_defContext) {}

// EnterMacro_def is called when production macro_def is entered.
func (s *BasePromQLExParserListener) EnterMacro_def(ctx *Macro_defContext) {}

// ExitMacro_def is called when production macro_def is exited.
func (s *BasePromQLExParserListener) ExitMacro_def(ctx *Macro_defContext) {}

// EnterSubstitute is called when production substitute is entered.
func (s *BasePromQLExParserListener) EnterSubstitute(ctx *SubstituteContext) {}

// ExitSubstitute is called when production substitute is exited.
func (s *BasePromQLExParserListener) ExitSubstitute(ctx *SubstituteContext) {}

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

// EnterTimeInstLit_IsoDateTime is called when production TimeInstLit_IsoDateTime is entered.
func (s *BasePromQLExParserListener) EnterTimeInstLit_IsoDateTime(ctx *TimeInstLit_IsoDateTimeContext) {
}

// ExitTimeInstLit_IsoDateTime is called when production TimeInstLit_IsoDateTime is exited.
func (s *BasePromQLExParserListener) ExitTimeInstLit_IsoDateTime(ctx *TimeInstLit_IsoDateTimeContext) {
}

// EnterTimeInstLit_UnixTimestamp is called when production TimeInstLit_UnixTimestamp is entered.
func (s *BasePromQLExParserListener) EnterTimeInstLit_UnixTimestamp(ctx *TimeInstLit_UnixTimestampContext) {
}

// ExitTimeInstLit_UnixTimestamp is called when production TimeInstLit_UnixTimestamp is exited.
func (s *BasePromQLExParserListener) ExitTimeInstLit_UnixTimestamp(ctx *TimeInstLit_UnixTimestampContext) {
}

// EnterConsNumExpr_NumLiteral is called when production ConsNumExpr_NumLiteral is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_NumLiteral(ctx *ConsNumExpr_NumLiteralContext) {
}

// ExitConsNumExpr_NumLiteral is called when production ConsNumExpr_NumLiteral is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_NumLiteral(ctx *ConsNumExpr_NumLiteralContext) {}

// EnterConsNumExpr_MulOp is called when production ConsNumExpr_MulOp is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_MulOp(ctx *ConsNumExpr_MulOpContext) {}

// ExitConsNumExpr_MulOp is called when production ConsNumExpr_MulOp is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_MulOp(ctx *ConsNumExpr_MulOpContext) {}

// EnterConsNumExpr_PowerOp is called when production ConsNumExpr_PowerOp is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_PowerOp(ctx *ConsNumExpr_PowerOpContext) {}

// ExitConsNumExpr_PowerOp is called when production ConsNumExpr_PowerOp is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_PowerOp(ctx *ConsNumExpr_PowerOpContext) {}

// EnterConsNumExpr_UnaryOp is called when production ConsNumExpr_UnaryOp is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_UnaryOp(ctx *ConsNumExpr_UnaryOpContext) {}

// ExitConsNumExpr_UnaryOp is called when production ConsNumExpr_UnaryOp is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_UnaryOp(ctx *ConsNumExpr_UnaryOpContext) {}

// EnterConsNumExpr_ParenOp is called when production ConsNumExpr_ParenOp is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_ParenOp(ctx *ConsNumExpr_ParenOpContext) {}

// ExitConsNumExpr_ParenOp is called when production ConsNumExpr_ParenOp is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_ParenOp(ctx *ConsNumExpr_ParenOpContext) {}

// EnterConsNumExpr_AddOp is called when production ConsNumExpr_AddOp is entered.
func (s *BasePromQLExParserListener) EnterConsNumExpr_AddOp(ctx *ConsNumExpr_AddOpContext) {}

// ExitConsNumExpr_AddOp is called when production ConsNumExpr_AddOp is exited.
func (s *BasePromQLExParserListener) ExitConsNumExpr_AddOp(ctx *ConsNumExpr_AddOpContext) {}

// EnterNumLit_Number is called when production NumLit_Number is entered.
func (s *BasePromQLExParserListener) EnterNumLit_Number(ctx *NumLit_NumberContext) {}

// ExitNumLit_Number is called when production NumLit_Number is exited.
func (s *BasePromQLExParserListener) ExitNumLit_Number(ctx *NumLit_NumberContext) {}

// EnterNumLit_Duration is called when production NumLit_Duration is entered.
func (s *BasePromQLExParserListener) EnterNumLit_Duration(ctx *NumLit_DurationContext) {}

// ExitNumLit_Duration is called when production NumLit_Duration is exited.
func (s *BasePromQLExParserListener) ExitNumLit_Duration(ctx *NumLit_DurationContext) {}

// EnterNumLit_TimeInstantLit is called when production NumLit_TimeInstantLit is entered.
func (s *BasePromQLExParserListener) EnterNumLit_TimeInstantLit(ctx *NumLit_TimeInstantLitContext) {}

// ExitNumLit_TimeInstantLit is called when production NumLit_TimeInstantLit is exited.
func (s *BasePromQLExParserListener) ExitNumLit_TimeInstantLit(ctx *NumLit_TimeInstantLitContext) {}

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

// EnterVecOp_AddOp is called when production VecOp_AddOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_AddOp(ctx *VecOp_AddOpContext) {}

// ExitVecOp_AddOp is called when production VecOp_AddOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_AddOp(ctx *VecOp_AddOpContext) {}

// EnterVecOp_SubqueryOp is called when production VecOp_SubqueryOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_SubqueryOp(ctx *VecOp_SubqueryOpContext) {}

// ExitVecOp_SubqueryOp is called when production VecOp_SubqueryOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_SubqueryOp(ctx *VecOp_SubqueryOpContext) {}

// EnterVecOp_OrOp is called when production VecOp_OrOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_OrOp(ctx *VecOp_OrOpContext) {}

// ExitVecOp_OrOp is called when production VecOp_OrOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_OrOp(ctx *VecOp_OrOpContext) {}

// EnterVecOp_Substitute is called when production VecOp_Substitute is entered.
func (s *BasePromQLExParserListener) EnterVecOp_Substitute(ctx *VecOp_SubstituteContext) {}

// ExitVecOp_Substitute is called when production VecOp_Substitute is exited.
func (s *BasePromQLExParserListener) ExitVecOp_Substitute(ctx *VecOp_SubstituteContext) {}

// EnterVecOp_Vec is called when production VecOp_Vec is entered.
func (s *BasePromQLExParserListener) EnterVecOp_Vec(ctx *VecOp_VecContext) {}

// ExitVecOp_Vec is called when production VecOp_Vec is exited.
func (s *BasePromQLExParserListener) ExitVecOp_Vec(ctx *VecOp_VecContext) {}

// EnterVecOp_At is called when production VecOp_At is entered.
func (s *BasePromQLExParserListener) EnterVecOp_At(ctx *VecOp_AtContext) {}

// ExitVecOp_At is called when production VecOp_At is exited.
func (s *BasePromQLExParserListener) ExitVecOp_At(ctx *VecOp_AtContext) {}

// EnterVecOp_UnaryOp is called when production VecOp_UnaryOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_UnaryOp(ctx *VecOp_UnaryOpContext) {}

// ExitVecOp_UnaryOp is called when production VecOp_UnaryOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_UnaryOp(ctx *VecOp_UnaryOpContext) {}

// EnterVecOp_VecMatchOp is called when production VecOp_VecMatchOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_VecMatchOp(ctx *VecOp_VecMatchOpContext) {}

// ExitVecOp_VecMatchOp is called when production VecOp_VecMatchOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_VecMatchOp(ctx *VecOp_VecMatchOpContext) {}

// EnterVecOp_MulOp is called when production VecOp_MulOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_MulOp(ctx *VecOp_MulOpContext) {}

// ExitVecOp_MulOp is called when production VecOp_MulOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_MulOp(ctx *VecOp_MulOpContext) {}

// EnterVecOp_ConstNumExpr is called when production VecOp_ConstNumExpr is entered.
func (s *BasePromQLExParserListener) EnterVecOp_ConstNumExpr(ctx *VecOp_ConstNumExprContext) {}

// ExitVecOp_ConstNumExpr is called when production VecOp_ConstNumExpr is exited.
func (s *BasePromQLExParserListener) ExitVecOp_ConstNumExpr(ctx *VecOp_ConstNumExprContext) {}

// EnterVecOp_PowOp is called when production VecOp_PowOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_PowOp(ctx *VecOp_PowOpContext) {}

// ExitVecOp_PowOp is called when production VecOp_PowOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_PowOp(ctx *VecOp_PowOpContext) {}

// EnterVecOp_CompareOp is called when production VecOp_CompareOp is entered.
func (s *BasePromQLExParserListener) EnterVecOp_CompareOp(ctx *VecOp_CompareOpContext) {}

// ExitVecOp_CompareOp is called when production VecOp_CompareOp is exited.
func (s *BasePromQLExParserListener) ExitVecOp_CompareOp(ctx *VecOp_CompareOpContext) {}

// EnterVecOp_AndUnless is called when production VecOp_AndUnless is entered.
func (s *BasePromQLExParserListener) EnterVecOp_AndUnless(ctx *VecOp_AndUnlessContext) {}

// ExitVecOp_AndUnless is called when production VecOp_AndUnless is exited.
func (s *BasePromQLExParserListener) ExitVecOp_AndUnless(ctx *VecOp_AndUnlessContext) {}

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

// EnterLit_ConstNumExpr is called when production Lit_ConstNumExpr is entered.
func (s *BasePromQLExParserListener) EnterLit_ConstNumExpr(ctx *Lit_ConstNumExprContext) {}

// ExitLit_ConstNumExpr is called when production Lit_ConstNumExpr is exited.
func (s *BasePromQLExParserListener) ExitLit_ConstNumExpr(ctx *Lit_ConstNumExprContext) {}

// EnterLit_String is called when production Lit_String is entered.
func (s *BasePromQLExParserListener) EnterLit_String(ctx *Lit_StringContext) {}

// ExitLit_String is called when production Lit_String is exited.
func (s *BasePromQLExParserListener) ExitLit_String(ctx *Lit_StringContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BasePromQLExParserListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BasePromQLExParserListener) ExitAggregation(ctx *AggregationContext) {}

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

// EnterAtModTime_ConstNumExpr is called when production AtModTime_ConstNumExpr is entered.
func (s *BasePromQLExParserListener) EnterAtModTime_ConstNumExpr(ctx *AtModTime_ConstNumExprContext) {
}

// ExitAtModTime_ConstNumExpr is called when production AtModTime_ConstNumExpr is exited.
func (s *BasePromQLExParserListener) ExitAtModTime_ConstNumExpr(ctx *AtModTime_ConstNumExprContext) {}

// EnterAtModTime_Start is called when production AtModTime_Start is entered.
func (s *BasePromQLExParserListener) EnterAtModTime_Start(ctx *AtModTime_StartContext) {}

// ExitAtModTime_Start is called when production AtModTime_Start is exited.
func (s *BasePromQLExParserListener) ExitAtModTime_Start(ctx *AtModTime_StartContext) {}

// EnterAtModTime_End is called when production AtModTime_End is entered.
func (s *BasePromQLExParserListener) EnterAtModTime_End(ctx *AtModTime_EndContext) {}

// ExitAtModTime_End is called when production AtModTime_End is exited.
func (s *BasePromQLExParserListener) ExitAtModTime_End(ctx *AtModTime_EndContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *BasePromQLExParserListener) EnterFunction_(ctx *Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *BasePromQLExParserListener) ExitFunction_(ctx *Function_Context) {}

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

// EnterParameter is called when production parameter is entered.
func (s *BasePromQLExParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePromQLExParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasePromQLExParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasePromQLExParserListener) ExitParameterList(ctx *ParameterListContext) {}

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

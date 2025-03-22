// Code generated from PromQLExParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package promqlex // PromQLExParser
import "github.com/antlr4-go/antlr/v4"

// PromQLExParserListener is a complete listener for a parse tree produced by PromQLExParser.
type PromQLExParserListener interface {
	antlr.ParseTreeListener

	// EnterPromqlx is called when entering the promqlx production.
	EnterPromqlx(c *PromqlxContext)

	// EnterState_AliasDef is called when entering the State_AliasDef production.
	EnterState_AliasDef(c *State_AliasDefContext)

	// EnterState_MacroDef is called when entering the State_MacroDef production.
	EnterState_MacroDef(c *State_MacroDefContext)

	// EnterState_IfStatement is called when entering the State_IfStatement production.
	EnterState_IfStatement(c *State_IfStatementContext)

	// EnterState_VectorOperation is called when entering the State_VectorOperation production.
	EnterState_VectorOperation(c *State_VectorOperationContext)

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

	// EnterTimeInstLit_IsoDateTime is called when entering the TimeInstLit_IsoDateTime production.
	EnterTimeInstLit_IsoDateTime(c *TimeInstLit_IsoDateTimeContext)

	// EnterTimeInstLit_UnixTimestamp is called when entering the TimeInstLit_UnixTimestamp production.
	EnterTimeInstLit_UnixTimestamp(c *TimeInstLit_UnixTimestampContext)

	// EnterConsNumExpr_NumLiteral is called when entering the ConsNumExpr_NumLiteral production.
	EnterConsNumExpr_NumLiteral(c *ConsNumExpr_NumLiteralContext)

	// EnterConsNumExpr_MulOp is called when entering the ConsNumExpr_MulOp production.
	EnterConsNumExpr_MulOp(c *ConsNumExpr_MulOpContext)

	// EnterConsNumExpr_PowerOp is called when entering the ConsNumExpr_PowerOp production.
	EnterConsNumExpr_PowerOp(c *ConsNumExpr_PowerOpContext)

	// EnterConsNumExpr_UnaryOp is called when entering the ConsNumExpr_UnaryOp production.
	EnterConsNumExpr_UnaryOp(c *ConsNumExpr_UnaryOpContext)

	// EnterConsNumExpr_ParenOp is called when entering the ConsNumExpr_ParenOp production.
	EnterConsNumExpr_ParenOp(c *ConsNumExpr_ParenOpContext)

	// EnterConsNumExpr_AddOp is called when entering the ConsNumExpr_AddOp production.
	EnterConsNumExpr_AddOp(c *ConsNumExpr_AddOpContext)

	// EnterNumLit_Number is called when entering the NumLit_Number production.
	EnterNumLit_Number(c *NumLit_NumberContext)

	// EnterNumLit_Duration is called when entering the NumLit_Duration production.
	EnterNumLit_Duration(c *NumLit_DurationContext)

	// EnterNumLit_TimeInstantLit is called when entering the NumLit_TimeInstantLit production.
	EnterNumLit_TimeInstantLit(c *NumLit_TimeInstantLitContext)

	// EnterNumLit_AliasCall is called when entering the NumLit_AliasCall production.
	EnterNumLit_AliasCall(c *NumLit_AliasCallContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// EnterTime_range is called when entering the time_range production.
	EnterTime_range(c *Time_rangeContext)

	// EnterSubquery_range is called when entering the subquery_range production.
	EnterSubquery_range(c *Subquery_rangeContext)

	// EnterVecOp_Macro is called when entering the VecOp_Macro production.
	EnterVecOp_Macro(c *VecOp_MacroContext)

	// EnterVecOp_AddOp is called when entering the VecOp_AddOp production.
	EnterVecOp_AddOp(c *VecOp_AddOpContext)

	// EnterVecOp_SubqueryOp is called when entering the VecOp_SubqueryOp production.
	EnterVecOp_SubqueryOp(c *VecOp_SubqueryOpContext)

	// EnterVecOp_OrOp is called when entering the VecOp_OrOp production.
	EnterVecOp_OrOp(c *VecOp_OrOpContext)

	// EnterVecOp_Vec is called when entering the VecOp_Vec production.
	EnterVecOp_Vec(c *VecOp_VecContext)

	// EnterVecOp_At is called when entering the VecOp_At production.
	EnterVecOp_At(c *VecOp_AtContext)

	// EnterVecOp_UnaryOp is called when entering the VecOp_UnaryOp production.
	EnterVecOp_UnaryOp(c *VecOp_UnaryOpContext)

	// EnterVecOp_VecMatchOp is called when entering the VecOp_VecMatchOp production.
	EnterVecOp_VecMatchOp(c *VecOp_VecMatchOpContext)

	// EnterVecOp_MulOp is called when entering the VecOp_MulOp production.
	EnterVecOp_MulOp(c *VecOp_MulOpContext)

	// EnterVecOp_ConstNumExpr is called when entering the VecOp_ConstNumExpr production.
	EnterVecOp_ConstNumExpr(c *VecOp_ConstNumExprContext)

	// EnterVecOp_OffsetOp is called when entering the VecOp_OffsetOp production.
	EnterVecOp_OffsetOp(c *VecOp_OffsetOpContext)

	// EnterVecOp_Alias is called when entering the VecOp_Alias production.
	EnterVecOp_Alias(c *VecOp_AliasContext)

	// EnterVecOp_PowOp is called when entering the VecOp_PowOp production.
	EnterVecOp_PowOp(c *VecOp_PowOpContext)

	// EnterVecOp_CompareOp is called when entering the VecOp_CompareOp production.
	EnterVecOp_CompareOp(c *VecOp_CompareOpContext)

	// EnterVecOp_AndUnless is called when entering the VecOp_AndUnless production.
	EnterVecOp_AndUnless(c *VecOp_AndUnlessContext)

	// EnterSubqueryOp is called when entering the subqueryOp production.
	EnterSubqueryOp(c *SubqueryOpContext)

	// EnterOffsetOp is called when entering the offsetOp production.
	EnterOffsetOp(c *OffsetOpContext)

	// EnterMatrixSelector is called when entering the matrixSelector production.
	EnterMatrixSelector(c *MatrixSelectorContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterLit_ConstNumExpr is called when entering the Lit_ConstNumExpr production.
	EnterLit_ConstNumExpr(c *Lit_ConstNumExprContext)

	// EnterLit_String is called when entering the Lit_String production.
	EnterLit_String(c *Lit_StringContext)

	// EnterInstantSelector is called when entering the instantSelector production.
	EnterInstantSelector(c *InstantSelectorContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterMetric_name is called when entering the metric_name production.
	EnterMetric_name(c *Metric_nameContext)

	// EnterAtModTime_ConstNumExpr is called when entering the AtModTime_ConstNumExpr production.
	EnterAtModTime_ConstNumExpr(c *AtModTime_ConstNumExprContext)

	// EnterAtModTime_Start is called when entering the AtModTime_Start production.
	EnterAtModTime_Start(c *AtModTime_StartContext)

	// EnterAtModTime_End is called when entering the AtModTime_End production.
	EnterAtModTime_End(c *AtModTime_EndContext)

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

	// ExitState_AliasDef is called when exiting the State_AliasDef production.
	ExitState_AliasDef(c *State_AliasDefContext)

	// ExitState_MacroDef is called when exiting the State_MacroDef production.
	ExitState_MacroDef(c *State_MacroDefContext)

	// ExitState_IfStatement is called when exiting the State_IfStatement production.
	ExitState_IfStatement(c *State_IfStatementContext)

	// ExitState_VectorOperation is called when exiting the State_VectorOperation production.
	ExitState_VectorOperation(c *State_VectorOperationContext)

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

	// ExitTimeInstLit_IsoDateTime is called when exiting the TimeInstLit_IsoDateTime production.
	ExitTimeInstLit_IsoDateTime(c *TimeInstLit_IsoDateTimeContext)

	// ExitTimeInstLit_UnixTimestamp is called when exiting the TimeInstLit_UnixTimestamp production.
	ExitTimeInstLit_UnixTimestamp(c *TimeInstLit_UnixTimestampContext)

	// ExitConsNumExpr_NumLiteral is called when exiting the ConsNumExpr_NumLiteral production.
	ExitConsNumExpr_NumLiteral(c *ConsNumExpr_NumLiteralContext)

	// ExitConsNumExpr_MulOp is called when exiting the ConsNumExpr_MulOp production.
	ExitConsNumExpr_MulOp(c *ConsNumExpr_MulOpContext)

	// ExitConsNumExpr_PowerOp is called when exiting the ConsNumExpr_PowerOp production.
	ExitConsNumExpr_PowerOp(c *ConsNumExpr_PowerOpContext)

	// ExitConsNumExpr_UnaryOp is called when exiting the ConsNumExpr_UnaryOp production.
	ExitConsNumExpr_UnaryOp(c *ConsNumExpr_UnaryOpContext)

	// ExitConsNumExpr_ParenOp is called when exiting the ConsNumExpr_ParenOp production.
	ExitConsNumExpr_ParenOp(c *ConsNumExpr_ParenOpContext)

	// ExitConsNumExpr_AddOp is called when exiting the ConsNumExpr_AddOp production.
	ExitConsNumExpr_AddOp(c *ConsNumExpr_AddOpContext)

	// ExitNumLit_Number is called when exiting the NumLit_Number production.
	ExitNumLit_Number(c *NumLit_NumberContext)

	// ExitNumLit_Duration is called when exiting the NumLit_Duration production.
	ExitNumLit_Duration(c *NumLit_DurationContext)

	// ExitNumLit_TimeInstantLit is called when exiting the NumLit_TimeInstantLit production.
	ExitNumLit_TimeInstantLit(c *NumLit_TimeInstantLitContext)

	// ExitNumLit_AliasCall is called when exiting the NumLit_AliasCall production.
	ExitNumLit_AliasCall(c *NumLit_AliasCallContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)

	// ExitTime_range is called when exiting the time_range production.
	ExitTime_range(c *Time_rangeContext)

	// ExitSubquery_range is called when exiting the subquery_range production.
	ExitSubquery_range(c *Subquery_rangeContext)

	// ExitVecOp_Macro is called when exiting the VecOp_Macro production.
	ExitVecOp_Macro(c *VecOp_MacroContext)

	// ExitVecOp_AddOp is called when exiting the VecOp_AddOp production.
	ExitVecOp_AddOp(c *VecOp_AddOpContext)

	// ExitVecOp_SubqueryOp is called when exiting the VecOp_SubqueryOp production.
	ExitVecOp_SubqueryOp(c *VecOp_SubqueryOpContext)

	// ExitVecOp_OrOp is called when exiting the VecOp_OrOp production.
	ExitVecOp_OrOp(c *VecOp_OrOpContext)

	// ExitVecOp_Vec is called when exiting the VecOp_Vec production.
	ExitVecOp_Vec(c *VecOp_VecContext)

	// ExitVecOp_At is called when exiting the VecOp_At production.
	ExitVecOp_At(c *VecOp_AtContext)

	// ExitVecOp_UnaryOp is called when exiting the VecOp_UnaryOp production.
	ExitVecOp_UnaryOp(c *VecOp_UnaryOpContext)

	// ExitVecOp_VecMatchOp is called when exiting the VecOp_VecMatchOp production.
	ExitVecOp_VecMatchOp(c *VecOp_VecMatchOpContext)

	// ExitVecOp_MulOp is called when exiting the VecOp_MulOp production.
	ExitVecOp_MulOp(c *VecOp_MulOpContext)

	// ExitVecOp_ConstNumExpr is called when exiting the VecOp_ConstNumExpr production.
	ExitVecOp_ConstNumExpr(c *VecOp_ConstNumExprContext)

	// ExitVecOp_OffsetOp is called when exiting the VecOp_OffsetOp production.
	ExitVecOp_OffsetOp(c *VecOp_OffsetOpContext)

	// ExitVecOp_Alias is called when exiting the VecOp_Alias production.
	ExitVecOp_Alias(c *VecOp_AliasContext)

	// ExitVecOp_PowOp is called when exiting the VecOp_PowOp production.
	ExitVecOp_PowOp(c *VecOp_PowOpContext)

	// ExitVecOp_CompareOp is called when exiting the VecOp_CompareOp production.
	ExitVecOp_CompareOp(c *VecOp_CompareOpContext)

	// ExitVecOp_AndUnless is called when exiting the VecOp_AndUnless production.
	ExitVecOp_AndUnless(c *VecOp_AndUnlessContext)

	// ExitSubqueryOp is called when exiting the subqueryOp production.
	ExitSubqueryOp(c *SubqueryOpContext)

	// ExitOffsetOp is called when exiting the offsetOp production.
	ExitOffsetOp(c *OffsetOpContext)

	// ExitMatrixSelector is called when exiting the matrixSelector production.
	ExitMatrixSelector(c *MatrixSelectorContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitLit_ConstNumExpr is called when exiting the Lit_ConstNumExpr production.
	ExitLit_ConstNumExpr(c *Lit_ConstNumExprContext)

	// ExitLit_String is called when exiting the Lit_String production.
	ExitLit_String(c *Lit_StringContext)

	// ExitInstantSelector is called when exiting the instantSelector production.
	ExitInstantSelector(c *InstantSelectorContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitMetric_name is called when exiting the metric_name production.
	ExitMetric_name(c *Metric_nameContext)

	// ExitAtModTime_ConstNumExpr is called when exiting the AtModTime_ConstNumExpr production.
	ExitAtModTime_ConstNumExpr(c *AtModTime_ConstNumExprContext)

	// ExitAtModTime_Start is called when exiting the AtModTime_Start production.
	ExitAtModTime_Start(c *AtModTime_StartContext)

	// ExitAtModTime_End is called when exiting the AtModTime_End production.
	ExitAtModTime_End(c *AtModTime_EndContext)

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

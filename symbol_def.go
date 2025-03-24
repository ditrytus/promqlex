package promqlex

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ditrytus/promqlex/parsers/promqlex"
)

type SymbolsDefiner struct {
	scopes       map[antlr.ParseTree]*Scope
	errors       []*ParseTreeError
	currentScope *Scope
}

func NewSymbolsDefiner() *SymbolsDefiner {
	return &SymbolsDefiner{
		scopes: make(map[antlr.ParseTree]*Scope),
	}
}

func (m *SymbolsDefiner) VisitTerminal(node antlr.TerminalNode) {}

func (m *SymbolsDefiner) VisitErrorNode(node antlr.ErrorNode) {}

func (m *SymbolsDefiner) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (m *SymbolsDefiner) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (m *SymbolsDefiner) EnterPromqlx(c *promqlex.PromqlxContext) {
	m.enterNewScope(c)
}

func (m *SymbolsDefiner) enterNewScope(c antlr.ParseTree) {
	newScope := NewScope(m.currentScope)
	if m.currentScope == nil {
		m.currentScope = newScope
	}
	m.scopes[c] = newScope
}

func (m *SymbolsDefiner) EnterState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (m *SymbolsDefiner) EnterState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (m *SymbolsDefiner) EnterState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (m *SymbolsDefiner) EnterState_VectorOperation(c *promqlex.State_VectorOperationContext) {}

func (m *SymbolsDefiner) EnterAlias_def(c *promqlex.Alias_defContext) {
	m.defineAliasSymbol(c.ID())
}

func (m *SymbolsDefiner) defineAliasSymbol(idTerminal antlr.TerminalNode) {
	err := m.currentScope.Define(NewAliasSymbol(idTerminal.GetText()))
	if err != nil {
		m.addError(err, idTerminal)
	}
}

func (m *SymbolsDefiner) addError(err error, idTerminal antlr.ParseTree) {
	m.errors = append(m.errors, NewParseTreeError(err, idTerminal))
}

func (m *SymbolsDefiner) EnterAlias_call(c *promqlex.Alias_callContext) {}

func (m *SymbolsDefiner) EnterMacro_def(c *promqlex.Macro_defContext) {
	var ary int
	if c.Args_decl() != nil {
		ary = len(c.Args_decl().AllArg_name())
	}
	err := m.currentScope.Define(NewMacroSymbol(c.ID().GetText(), ary))
	if err != nil {
		m.addError(err, c.ID())
	}
	m.enterNewScope(c)
}

func (m *SymbolsDefiner) EnterMacro_call(c *promqlex.Macro_callContext) {}

func (m *SymbolsDefiner) EnterArgs_decl(c *promqlex.Args_declContext) {}

func (m *SymbolsDefiner) EnterArg_name(c *promqlex.Arg_nameContext) {
	m.defineAliasSymbol(c.ID())
}

func (m *SymbolsDefiner) EnterStatement_block(c *promqlex.Statement_blockContext) {}

func (m *SymbolsDefiner) EnterArg_list(c *promqlex.Arg_listContext) {}

func (m *SymbolsDefiner) EnterIf_statement(c *promqlex.If_statementContext) {}

func (m *SymbolsDefiner) EnterCondition(c *promqlex.ConditionContext) {}

func (m *SymbolsDefiner) EnterCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {}

func (m *SymbolsDefiner) EnterTrueConst(c *promqlex.TrueConstContext) {}

func (m *SymbolsDefiner) EnterFalseConst(c *promqlex.FalseConstContext) {}

func (m *SymbolsDefiner) EnterTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {}

func (m *SymbolsDefiner) EnterTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
}

func (m *SymbolsDefiner) EnterConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {}

func (m *SymbolsDefiner) EnterConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {}

func (m *SymbolsDefiner) EnterConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {}

func (m *SymbolsDefiner) EnterConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {}

func (m *SymbolsDefiner) EnterConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (m *SymbolsDefiner) EnterConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {}

func (m *SymbolsDefiner) EnterNumLit_Number(c *promqlex.NumLit_NumberContext) {}

func (m *SymbolsDefiner) EnterNumLit_Duration(c *promqlex.NumLit_DurationContext) {}

func (m *SymbolsDefiner) EnterNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {}

func (m *SymbolsDefiner) EnterNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {}

func (m *SymbolsDefiner) EnterDuration(c *promqlex.DurationContext) {}

func (m *SymbolsDefiner) EnterTime_range(c *promqlex.Time_rangeContext) {}

func (m *SymbolsDefiner) EnterSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (m *SymbolsDefiner) EnterVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (m *SymbolsDefiner) EnterVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (m *SymbolsDefiner) EnterVecOp_At(c *promqlex.VecOp_AtContext) {}

func (m *SymbolsDefiner) EnterVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {}

func (m *SymbolsDefiner) EnterVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (m *SymbolsDefiner) EnterVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (m *SymbolsDefiner) EnterVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (m *SymbolsDefiner) EnterSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (m *SymbolsDefiner) EnterOffsetOp(c *promqlex.OffsetOpContext) {}

func (m *SymbolsDefiner) EnterMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (m *SymbolsDefiner) EnterOffset(c *promqlex.OffsetContext) {}

func (m *SymbolsDefiner) EnterLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {}

func (m *SymbolsDefiner) EnterLit_String(c *promqlex.Lit_StringContext) {}

func (m *SymbolsDefiner) EnterInstantSelector(c *promqlex.InstantSelectorContext) {}

func (m *SymbolsDefiner) EnterLabelName(c *promqlex.LabelNameContext) {}

func (m *SymbolsDefiner) EnterMetric_name(c *promqlex.Metric_nameContext) {}

func (m *SymbolsDefiner) EnterAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {}

func (m *SymbolsDefiner) EnterAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (m *SymbolsDefiner) EnterAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (m *SymbolsDefiner) EnterExpression(c *promqlex.ExpressionContext) {}

func (m *SymbolsDefiner) EnterUnaryOp(c *promqlex.UnaryOpContext) {}

func (m *SymbolsDefiner) EnterPowOp(c *promqlex.PowOpContext) {}

func (m *SymbolsDefiner) EnterMultOp(c *promqlex.MultOpContext) {}

func (m *SymbolsDefiner) EnterAddOp(c *promqlex.AddOpContext) {}

func (m *SymbolsDefiner) EnterCompareOp(c *promqlex.CompareOpContext) {}

func (m *SymbolsDefiner) EnterAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (m *SymbolsDefiner) EnterOrOp(c *promqlex.OrOpContext) {}

func (m *SymbolsDefiner) EnterVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (m *SymbolsDefiner) EnterVector(c *promqlex.VectorContext) {}

func (m *SymbolsDefiner) EnterParens(c *promqlex.ParensContext) {}

func (m *SymbolsDefiner) EnterLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (m *SymbolsDefiner) EnterLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (m *SymbolsDefiner) EnterLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (m *SymbolsDefiner) EnterFunction_(c *promqlex.Function_Context) {}

func (m *SymbolsDefiner) EnterParameter(c *promqlex.ParameterContext) {}

func (m *SymbolsDefiner) EnterParameterList(c *promqlex.ParameterListContext) {}

func (m *SymbolsDefiner) EnterAggregation(c *promqlex.AggregationContext) {}

func (m *SymbolsDefiner) EnterBy(c *promqlex.ByContext) {}

func (m *SymbolsDefiner) EnterWithout(c *promqlex.WithoutContext) {}

func (m *SymbolsDefiner) EnterGrouping(c *promqlex.GroupingContext) {}

func (m *SymbolsDefiner) EnterOn_(c *promqlex.On_Context) {}

func (m *SymbolsDefiner) EnterIgnoring(c *promqlex.IgnoringContext) {}

func (m *SymbolsDefiner) EnterGroupLeft(c *promqlex.GroupLeftContext) {}

func (m *SymbolsDefiner) EnterGroupRight(c *promqlex.GroupRightContext) {}

func (m *SymbolsDefiner) EnterLabelNameList(c *promqlex.LabelNameListContext) {}

func (m *SymbolsDefiner) EnterKeyword(c *promqlex.KeywordContext) {}

func (m *SymbolsDefiner) EnterString(c *promqlex.StringContext) {}

func (m *SymbolsDefiner) ExitPromqlx(c *promqlex.PromqlxContext) {
	m.exitScope()
}

func (m *SymbolsDefiner) exitScope() {
	m.currentScope = m.currentScope.Parent()
}

func (m *SymbolsDefiner) ExitState_AliasDef(c *promqlex.State_AliasDefContext) {}

func (m *SymbolsDefiner) ExitState_MacroDef(c *promqlex.State_MacroDefContext) {}

func (m *SymbolsDefiner) ExitState_IfStatement(c *promqlex.State_IfStatementContext) {}

func (m *SymbolsDefiner) ExitState_VectorOperation(c *promqlex.State_VectorOperationContext) {}

func (m *SymbolsDefiner) ExitAlias_def(c *promqlex.Alias_defContext) {}

func (m *SymbolsDefiner) ExitAlias_call(c *promqlex.Alias_callContext) {}

func (m *SymbolsDefiner) ExitMacro_def(c *promqlex.Macro_defContext) {
	m.exitScope()
}

func (m *SymbolsDefiner) ExitMacro_call(c *promqlex.Macro_callContext) {}

func (m *SymbolsDefiner) ExitArgs_decl(c *promqlex.Args_declContext) {}

func (m *SymbolsDefiner) ExitArg_name(c *promqlex.Arg_nameContext) {}

func (m *SymbolsDefiner) ExitStatement_block(c *promqlex.Statement_blockContext) {}

func (m *SymbolsDefiner) ExitArg_list(c *promqlex.Arg_listContext) {}

func (m *SymbolsDefiner) ExitIf_statement(c *promqlex.If_statementContext) {}

func (m *SymbolsDefiner) ExitCondition(c *promqlex.ConditionContext) {}

func (m *SymbolsDefiner) ExitCompareVectorOperation(c *promqlex.CompareVectorOperationContext) {}

func (m *SymbolsDefiner) ExitTrueConst(c *promqlex.TrueConstContext) {}

func (m *SymbolsDefiner) ExitFalseConst(c *promqlex.FalseConstContext) {}

func (m *SymbolsDefiner) ExitTimeInstLit_IsoDateTime(c *promqlex.TimeInstLit_IsoDateTimeContext) {}

func (m *SymbolsDefiner) ExitTimeInstLit_UnixTimestamp(c *promqlex.TimeInstLit_UnixTimestampContext) {
}

func (m *SymbolsDefiner) ExitConsNumExpr_NumLiteral(c *promqlex.ConsNumExpr_NumLiteralContext) {}

func (m *SymbolsDefiner) ExitConsNumExpr_MulOp(c *promqlex.ConsNumExpr_MulOpContext) {}

func (m *SymbolsDefiner) ExitConsNumExpr_PowerOp(c *promqlex.ConsNumExpr_PowerOpContext) {}

func (m *SymbolsDefiner) ExitConsNumExpr_UnaryOp(c *promqlex.ConsNumExpr_UnaryOpContext) {}

func (m *SymbolsDefiner) ExitConsNumExpr_ParenOp(c *promqlex.ConsNumExpr_ParenOpContext) {}

func (m *SymbolsDefiner) ExitConsNumExpr_AddOp(c *promqlex.ConsNumExpr_AddOpContext) {}

func (m *SymbolsDefiner) ExitNumLit_Number(c *promqlex.NumLit_NumberContext) {}

func (m *SymbolsDefiner) ExitNumLit_Duration(c *promqlex.NumLit_DurationContext) {}

func (m *SymbolsDefiner) ExitNumLit_TimeInstantLit(c *promqlex.NumLit_TimeInstantLitContext) {}

func (m *SymbolsDefiner) ExitNumLit_AliasCall(c *promqlex.NumLit_AliasCallContext) {}

func (m *SymbolsDefiner) ExitDuration(c *promqlex.DurationContext) {}

func (m *SymbolsDefiner) ExitTime_range(c *promqlex.Time_rangeContext) {}

func (m *SymbolsDefiner) ExitSubquery_range(c *promqlex.Subquery_rangeContext) {}

func (m *SymbolsDefiner) ExitVecOp_Macro(c *promqlex.VecOp_MacroContext) {}

func (m *SymbolsDefiner) ExitVecOp_AddOp(c *promqlex.VecOp_AddOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_SubqueryOp(c *promqlex.VecOp_SubqueryOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_OrOp(c *promqlex.VecOp_OrOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_Vec(c *promqlex.VecOp_VecContext) {}

func (m *SymbolsDefiner) ExitVecOp_At(c *promqlex.VecOp_AtContext) {}

func (m *SymbolsDefiner) ExitVecOp_UnaryOp(c *promqlex.VecOp_UnaryOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_VecMatchOp(c *promqlex.VecOp_VecMatchOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_MulOp(c *promqlex.VecOp_MulOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_ConstNumExpr(c *promqlex.VecOp_ConstNumExprContext) {}

func (m *SymbolsDefiner) ExitVecOp_Alias(c *promqlex.VecOp_AliasContext) {}

func (m *SymbolsDefiner) ExitVecOp_PowOp(c *promqlex.VecOp_PowOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_CompareOp(c *promqlex.VecOp_CompareOpContext) {}

func (m *SymbolsDefiner) ExitVecOp_AndUnless(c *promqlex.VecOp_AndUnlessContext) {}

func (m *SymbolsDefiner) ExitSubqueryOp(c *promqlex.SubqueryOpContext) {}

func (m *SymbolsDefiner) ExitOffsetOp(c *promqlex.OffsetOpContext) {}

func (m *SymbolsDefiner) ExitMatrixSelector(c *promqlex.MatrixSelectorContext) {}

func (m *SymbolsDefiner) ExitOffset(c *promqlex.OffsetContext) {}

func (m *SymbolsDefiner) ExitLit_ConstNumExpr(c *promqlex.Lit_ConstNumExprContext) {}

func (m *SymbolsDefiner) ExitLit_String(c *promqlex.Lit_StringContext) {}

func (m *SymbolsDefiner) ExitInstantSelector(c *promqlex.InstantSelectorContext) {}

func (m *SymbolsDefiner) ExitLabelName(c *promqlex.LabelNameContext) {}

func (m *SymbolsDefiner) ExitMetric_name(c *promqlex.Metric_nameContext) {}

func (m *SymbolsDefiner) ExitAtModTime_ConstNumExpr(c *promqlex.AtModTime_ConstNumExprContext) {}

func (m *SymbolsDefiner) ExitAtModTime_Start(c *promqlex.AtModTime_StartContext) {}

func (m *SymbolsDefiner) ExitAtModTime_End(c *promqlex.AtModTime_EndContext) {}

func (m *SymbolsDefiner) ExitExpression(c *promqlex.ExpressionContext) {}

func (m *SymbolsDefiner) ExitUnaryOp(c *promqlex.UnaryOpContext) {}

func (m *SymbolsDefiner) ExitPowOp(c *promqlex.PowOpContext) {}

func (m *SymbolsDefiner) ExitMultOp(c *promqlex.MultOpContext) {}

func (m *SymbolsDefiner) ExitAddOp(c *promqlex.AddOpContext) {}

func (m *SymbolsDefiner) ExitCompareOp(c *promqlex.CompareOpContext) {}

func (m *SymbolsDefiner) ExitAndUnlessOp(c *promqlex.AndUnlessOpContext) {}

func (m *SymbolsDefiner) ExitOrOp(c *promqlex.OrOpContext) {}

func (m *SymbolsDefiner) ExitVectorMatchOp(c *promqlex.VectorMatchOpContext) {}

func (m *SymbolsDefiner) ExitVector(c *promqlex.VectorContext) {}

func (m *SymbolsDefiner) ExitParens(c *promqlex.ParensContext) {}

func (m *SymbolsDefiner) ExitLabelMatcher(c *promqlex.LabelMatcherContext) {}

func (m *SymbolsDefiner) ExitLabelMatcherOperator(c *promqlex.LabelMatcherOperatorContext) {}

func (m *SymbolsDefiner) ExitLabelMatcherList(c *promqlex.LabelMatcherListContext) {}

func (m *SymbolsDefiner) ExitFunction_(c *promqlex.Function_Context) {}

func (m *SymbolsDefiner) ExitParameter(c *promqlex.ParameterContext) {}

func (m *SymbolsDefiner) ExitParameterList(c *promqlex.ParameterListContext) {}

func (m *SymbolsDefiner) ExitAggregation(c *promqlex.AggregationContext) {}

func (m *SymbolsDefiner) ExitBy(c *promqlex.ByContext) {}

func (m *SymbolsDefiner) ExitWithout(c *promqlex.WithoutContext) {}

func (m *SymbolsDefiner) ExitGrouping(c *promqlex.GroupingContext) {}

func (m *SymbolsDefiner) ExitOn_(c *promqlex.On_Context) {}

func (m *SymbolsDefiner) ExitIgnoring(c *promqlex.IgnoringContext) {}

func (m *SymbolsDefiner) ExitGroupLeft(c *promqlex.GroupLeftContext) {}

func (m *SymbolsDefiner) ExitGroupRight(c *promqlex.GroupRightContext) {}

func (m *SymbolsDefiner) ExitLabelNameList(c *promqlex.LabelNameListContext) {}

func (m *SymbolsDefiner) ExitKeyword(c *promqlex.KeywordContext) {}

func (m *SymbolsDefiner) ExitString(c *promqlex.StringContext) {}

func (m *SymbolsDefiner) HasErrors() bool {
	return len(m.errors) > 0
}

func (m *SymbolsDefiner) Errors() []*ParseTreeError {
	return m.errors
}

var _ promqlex.PromQLExParserListener = &SymbolsDefiner{}

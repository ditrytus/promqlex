// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar PromQLExParser;

import PromQLParser;

options {
    tokenVocab = PromQLExLexer;
}

promqlx : statement (SEMICOLON statement)* EOF;

statement
    : alias_def # State_AliasDef
    | macro_def # State_MacroDef
    | if_statement # State_IfStatement
    | vectorOperation # State_VectorOperation
    ;

alias_def : DEF ID EQ vectorOperation;
alias_call : CALL_SIGN ID;

macro_def : DEF ID LEFT_PAREN args_decl? RIGHT_PAREN statement_block;
macro_call: CALL_SIGN ID ID LEFT_PAREN arg_list? RIGHT_PAREN;

args_decl : arg_name (COMMA arg_name)*;

arg_name : ID;

statement_block : LEFT_BRACE promqlx RIGHT_BRACE;

arg_list : vectorOperation (COMMA vectorOperation)+;

if_statement : IF condition statement_block;

condition
    : compareVectorOperation
    | trueConst
    | falseConst
    ;

compareVectorOperation: vectorOperation compareOp vectorOperation;

trueConst : TRUE;
falseConst: FALSE;

// Time literals

time_instant_literal
    : RFC_3339_TIMESTAMP # TimeInstLit_IsoDateTime
    | NUMBER # TimeInstLit_UnixTimestamp
    ;

// Constant expressions

const_num_expression
    : <assoc = right> num_literal powOp num_literal # ConsNumExpr_PowerOp
    | unaryOp num_literal # ConsNumExpr_UnaryOp
    | num_literal multOp num_literal # ConsNumExpr_MulOp
    | num_literal addOp num_literal # ConsNumExpr_AddOp
    | LEFT_PAREN const_num_expression RIGHT_PAREN # ConsNumExpr_ParenOp
    | num_literal # ConsNumExpr_NumLiteral
    ;

num_literal
    : NUMBER # NumLit_Number
    | DURATION # NumLit_Duration
    | time_instant_literal # NumLit_TimeInstantLit
    | alias_call # NumLit_AliasCall
    ;

duration : const_num_expression;

// Lexer rules converted to parser rules

// This rule was converted from lexer rule TIME_RANGE.
time_range: LEFT_BRACKET duration RIGHT_BRACKET;

// This rule was converted from lexer rule SUBQUERY_RANGE.
subquery_range: LEFT_BRACKET duration COLON duration? RIGHT_BRACKET;

// Overrides of PromQLParser parser rules


vectorOperation
    : <assoc = right> vectorOperation powOp vectorOperation # VecOp_PowOp
    | <assoc = right> vectorOperation subqueryOp # VecOp_SubOp
    | unaryOp vectorOperation # VecOp_UnaryOp
    | vectorOperation multOp vectorOperation # VecOp_MulOp
    | vectorOperation addOp vectorOperation # VecOp_AddOp
    | vectorOperation compareOp vectorOperation # VecOp_CompareOp
    | vectorOperation andUnlessOp vectorOperation # VecOp_AndUnless
    | vectorOperation orOp vectorOperation # VecOp_OrOp
    | vectorOperation vectorMatchOp vectorOperation # VecOp_VecMatchOp
    | vectorOperation AT at_modifier_timestamp # VecOp_At
    | vector # VecOp_Vec
    // PROMQLX: an alias or macro call can be used anywhere in the context of vector operation.
    | macro_call # VecOp_Macro
    | alias_call # VecOp_Alias
    ;

subqueryOp
    // PROMQLX: SUBQUERY_RANGE was converted from token to rule.
    : subquery_range offsetOp?
    ;

offsetOp
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : OFFSET duration
    ;

matrixSelector
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : instantSelector time_range
    ;

offset
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : instantSelector OFFSET duration
    | matrixSelector OFFSET duration
    ;

literal
    // PROMQLX: everything that can be converted to the number of seconds
    // can be used as literal
    : const_num_expression # Lit_ConstNumExpr
    | string # Lit_String
    ;

instantSelector
    : metric_name (LEFT_BRACE labelMatcherList? RIGHT_BRACE)?
    | LEFT_BRACE labelMatcherList RIGHT_BRACE
    ;

labelName
    : keyword
    | metric_name
    | LABEL_NAME
    ;

metric_name
    : ID
    | METRIC_NAME
    ;

at_modifier_timestamp
    : const_num_expression # AtModTime_ConstNumExpr
    | START LEFT_PAREN RIGHT_PAREN # AtModTime_Start
    | END LEFT_PAREN RIGHT_PAREN # AtModTime_End
    ;
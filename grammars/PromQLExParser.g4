// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar PromQLExParser;

import PromQLParser;

options {
    tokenVocab = PromQLExLexer;
}

promqlx : statement_block EOF;

statement
    : alias_def # State_AliasDef
    | macro_def # State_MacroDef
    | if_statement # State_IfStatement
    | vectorOperation # State_VectorOperation
    ;

alias_def : DEF ID EQ vectorOperation;

macro_def : DEF ID LEFT_PAREN args_decl? RIGHT_PAREN LEFT_BRACE statement_block RIGHT_BRACE;
substitute: DOLLAR ID parameterList?;

args_decl : arg_name (COMMA arg_name)*;

arg_name : ID;

statement_block : statement (SEMICOLON statement)* ;

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
    : <assoc = right> const_num_expression powOp const_num_expression # ConsNumExpr_PowerOp
    | unaryOp const_num_expression # ConsNumExpr_UnaryOp
    | const_num_expression multOp const_num_expression # ConsNumExpr_MulOp
    | const_num_expression addOp const_num_expression # ConsNumExpr_AddOp
    | LEFT_PAREN const_num_expression RIGHT_PAREN # ConsNumExpr_ParenOp
    | num_literal # ConsNumExpr_NumLiteral
    ;

num_literal
    : NUMBER # NumLit_Number
    | DURATION # NumLit_Duration
    | time_instant_literal # NumLit_TimeInstantLit
    | substitute # NumLit_Substitute
    ;

duration : const_num_expression;

// Lexer rules converted to parser rules

// This rule was converted from lexer rule TIME_RANGE.
time_range: LEFT_BRACKET duration RIGHT_BRACKET;

// This rule was converted from lexer rule SUBQUERY_RANGE.
subquery_range: LEFT_BRACKET duration COLON duration? RIGHT_BRACKET;

// Overrides of PromQLParser parser rules

vectorOperation
    : const_num_expression # VecOp_ConstNumExpr
    | <assoc = right> vectorOperation powOp vectorOperation # VecOp_PowOp
    | <assoc = right> vectorOperation subqueryOp # VecOp_SubqueryOp
    | unaryOp vectorOperation # VecOp_UnaryOp
    | vectorOperation multOp vectorOperation # VecOp_MulOp
    | vectorOperation addOp vectorOperation # VecOp_AddOp
    | vectorOperation compareOp vectorOperation # VecOp_CompareOp
    | vectorOperation andUnlessOp vectorOperation # VecOp_AndUnless
    | vectorOperation orOp vectorOperation # VecOp_OrOp
    | vectorOperation vectorMatchOp vectorOperation # VecOp_VecMatchOp
    | vectorOperation AT at_modifier_timestamp # VecOp_At
    | vector # VecOp_Vec
    // PROMQLX: a substitute can be used anywhere in the context of vector operation.
    | substitute # VecOp_Substitute
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

aggregation
    : ID parameterList
    | ID (by | without) parameterList
    | ID parameterList ( by | without)
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
    | COLON
    | (COLON ID)+ COLON?
    | (ID COLON)+ ID?
    ;

at_modifier_timestamp
    : const_num_expression # AtModTime_ConstNumExpr
    | START LEFT_PAREN RIGHT_PAREN # AtModTime_Start
    | END LEFT_PAREN RIGHT_PAREN # AtModTime_End
    ;

function_
    : ID LEFT_PAREN (parameter (COMMA parameter)*)? RIGHT_PAREN
    ;
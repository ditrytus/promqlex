// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar PromQLExParser;

import PromQLParser;

options {
    tokenVocab = PromQLExLexer;
}

promqlx : statement (SEMICOLON statement)* EOF;

statement
    : alias_def
    | macro_def
    | if_statement
    | vectorOperation
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
    : iso_date_time
    | unix_timestamp
    ;

iso_date_time
    : year=NUMBER SUB (
        month=NUMBER (
            SUB day=NUMBER (
                T hour=NUMBER (
                    COLON minutes=NUMBER (
                        COLON seconds=NUMBER
                    )?
                )?
            )?
        )?
    )?;

unix_timestamp: NUMBER;

// Constant expressions

const_num_expression
    : <assoc = right> num_literal powOp num_literal
    | unaryOp num_literal
    | num_literal multOp num_literal
    | num_literal addOp num_literal
    | LEFT_PAREN const_num_expression RIGHT_PAREN
    | num_literal
    ;

num_literal
    : NUMBER
    | DURATION
    | time_instant_literal
    | alias_call
    ;

duration : const_num_expression;

// Lexer rules converted to parser rules

// This rule was converted from lexer rule TIME_RANGE.
time_range: LEFT_BRACKET duration RIGHT_BRACKET;

// This rule was converted from lexer rule SUBQUERY_RANGE.
subquery_range: LEFT_BRACKET duration COLON duration? RIGHT_BRACKET;

// Overrides of PromQLParser parser rules

vectorOperation
    : <assoc = right> vectorOperation powOp vectorOperation
    | <assoc = right> vectorOperation subqueryOp
    | unaryOp vectorOperation
    | vectorOperation multOp vectorOperation
    | vectorOperation addOp vectorOperation
    | vectorOperation compareOp vectorOperation
    | vectorOperation andUnlessOp vectorOperation
    | vectorOperation orOp vectorOperation
    | vectorOperation vectorMatchOp vectorOperation
    | vectorOperation AT at_modifier_timestamp
    | vector
    // PROMQLX: an alias or macro call can be used anywhere in the context of vector operation.
    | macro_call
    | alias_call
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
    : const_num_expression
    | string
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
    : const_num_expression
    | START LEFT_PAREN RIGHT_PAREN
    | END LEFT_PAREN RIGHT_PAREN
    ;
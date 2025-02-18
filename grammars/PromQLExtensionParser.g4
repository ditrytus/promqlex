/*
 [The "BSD licence"]
 Copyright (c) 2013 Terence Parr
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions
 are met:
 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in the
    documentation and/or other materials provided with the distribution.
 3. The name of the author may not be used to endorse or promote products
    derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar PromQLExtensionParser;

options {
    tokenVocab = PromQLExtensionLexer;
}

promqlx : ex_statement_list;

ex_statement_list : ex_statement+ ;

ex_statement
    : ex_alias_def
    | ex_macro_def
    | ex_if_statement
    | expression
    ;

ex_alias_def : EX_ID EQ expression;

ex_macro_def : EX_ID LEFT_PAREN ex_arg_list? RIGHT_PAREN ex_block;

ex_block : LEFT_BRACE ex_statement_list RIGHT_BRACE;

ex_arg_list : expression (COMMA expression)+;

ex_if_statement : ex_condition ex_block;

ex_condition
    : ex_compareVectorOperation
    | ex_trueConst
    | ex_falseConst
    ;

ex_compareVectorOperation: vectorOperation compareOp vectorOperation;

ex_trueConst : EX_TRUE;
ex_falseConst: EX_FALSE;

// Time literals

ex_time_instant
    : ex_iso_date_time
    ;

ex_iso_date_time
    : ex_iso_date_time_ymdhmsf
    | ex_iso_date_time_ymdhms
    | ex_iso_date_time_ymdhm
    | ex_iso_date_time_ymdh
    | ex_iso_date_time_ymd
    | ex_iso_date_time_ym
    | ex_iso_date_time_y
    ;

ex_iso_date_time_ymdhmsf: ex_year SUB ex_month SUB ex_day EX_T ex_hour EX_COLON ex_minutes EX_COLON ex_seconds EX_DOT ex_frac_sec;
ex_iso_date_time_ymdhms: ex_year SUB ex_month SUB ex_day EX_T ex_hour EX_COLON ex_minutes EX_COLON ex_seconds;
ex_iso_date_time_ymdhm: ex_year SUB ex_month SUB ex_day EX_T ex_hour EX_COLON ex_minutes;
ex_iso_date_time_ymdh: ex_year SUB ex_month SUB ex_day EX_T ex_hour;
ex_iso_date_time_ymd: ex_year SUB ex_month SUB ex_day;
ex_iso_date_time_ym: ex_year SUB ex_month;
ex_iso_date_time_y: ex_year;

ex_year : EX_POSITIVE_INTEGER;
ex_month: EX_TWO_DIGITS;
ex_day: EX_TWO_DIGITS;
ex_hour: EX_TWO_DIGITS;
ex_minutes: EX_TWO_DIGITS;
ex_seconds: EX_TWO_DIGITS;
ex_frac_sec: EX_DIGITS;

// Original PromQLGrammar. All alterations have additional comments
// starting with PROMQLX.

expression
    : vectorOperation EOF
    ;

// Binary operations are ordered by precedence

// Unary operations have the same precedence as multiplications

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
    | vectorOperation AT vectorOperation
    | vector
    ;

// Operators

unaryOp
    : (ADD | SUB)
    ;

powOp
    : POW grouping?
    ;

multOp
    : (MULT | DIV | MOD) grouping?
    ;

addOp
    : (ADD | SUB) grouping?
    ;

compareOp
    : (DEQ | NE | GT | LT | GE | LE) BOOL? grouping?
    ;

andUnlessOp
    : (AND | UNLESS) grouping?
    ;

orOp
    : OR grouping?
    ;

vectorMatchOp
    : (ON | UNLESS) grouping?
    ;

subqueryOp
    : SUBQUERY_RANGE offsetOp?
    ;

offsetOp
    : OFFSET DURATION
    ;

vector
    : function_
    | aggregation
    | instantSelector
    | matrixSelector
    | offset
    | literal
    | parens
    ;

parens
    : LEFT_PAREN vectorOperation RIGHT_PAREN
    ;

// Selectors

instantSelector
    : METRIC_NAME (LEFT_BRACE labelMatcherList? RIGHT_BRACE)?
    | LEFT_BRACE labelMatcherList RIGHT_BRACE
    ;

labelMatcher
    : labelName labelMatcherOperator STRING
    ;

labelMatcherOperator
    : EQ
    | NE
    | RE
    | NRE
    ;

labelMatcherList
    : labelMatcher (COMMA labelMatcher)* COMMA?
    ;

matrixSelector
    : instantSelector TIME_RANGE
    ;

offset
    : instantSelector OFFSET DURATION
    | matrixSelector OFFSET DURATION
    ;

// Functions

function_
    : FUNCTION LEFT_PAREN (parameter (COMMA parameter)*)? RIGHT_PAREN
    ;

parameter
    : literal
    | vectorOperation
    ;

parameterList
    : LEFT_PAREN (parameter (COMMA parameter)*)? RIGHT_PAREN
    ;

// Aggregations

aggregation
    : AGGREGATION_OPERATOR parameterList
    | AGGREGATION_OPERATOR (by | without) parameterList
    | AGGREGATION_OPERATOR parameterList ( by | without)
    ;

by
    : BY labelNameList
    ;

without
    : WITHOUT labelNameList
    ;

// Vector one-to-one/one-to-many joins

grouping
    : (on_ | ignoring) (groupLeft | groupRight)?
    ;

on_
    : ON labelNameList
    ;

ignoring
    : IGNORING labelNameList
    ;

groupLeft
    : GROUP_LEFT labelNameList?
    ;

groupRight
    : GROUP_RIGHT labelNameList?
    ;

// Label names

labelName
    : keyword
    | METRIC_NAME
    | LABEL_NAME
    ;

labelNameList
    : LEFT_PAREN (labelName (COMMA labelName)*)? RIGHT_PAREN
    ;

keyword
    : AND
    | OR
    | UNLESS
    | BY
    | WITHOUT
    | ON
    | IGNORING
    | GROUP_LEFT
    | GROUP_RIGHT
    | OFFSET
    | BOOL
    | AGGREGATION_OPERATOR
    | FUNCTION
    ;

literal
    : NUMBER
    | STRING
    ;
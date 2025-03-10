// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar PromQLExtensionParser;

import PromQLParser;

options {
    tokenVocab = PromQLExtensionLexer;
}

promqlx : (ex_statement EX_NL)+ ;

ex_statement
    : ex_alias_def
    | ex_macro_def
    | ex_if_statement
    | vectorOperation
    ;

ex_alias_def : EX_DEF EX_ID EQ vectorOperation;
ex_alias_call : EX_CALL_SIGN EX_ID;

ex_macro_def : EX_DEF EX_ID LEFT_PAREN ex_args_decl? RIGHT_PAREN ex_block;
ex_macro_call: EX_CALL_SIGN EX_ID EX_ID LEFT_PAREN ex_arg_list? RIGHT_PAREN;

ex_args_decl : ex_arg_name (COMMA ex_arg_name)*;

ex_arg_name : EX_ID;

ex_block : LEFT_BRACE promqlx RIGHT_BRACE;

ex_arg_list : vectorOperation (COMMA vectorOperation)+;

ex_if_statement : EX_IF ex_condition ex_block;

ex_condition
    : ex_compareVectorOperation
    | ex_trueConst
    | ex_falseConst
    ;

ex_compareVectorOperation: vectorOperation compareOp vectorOperation;

ex_trueConst : EX_TRUE;
ex_falseConst: EX_FALSE;

// Time literals

ex_time_instant_literal
    : ex_iso_date_time
    | ex_unix_timestamp
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

ex_unix_timestamp: EX_POSITIVE_INTEGER;

// Constant expressions

ex_const_num_expression
    : <assoc = right> ex_num_literal powOp ex_num_literal
    | unaryOp ex_num_literal
    | ex_num_literal multOp ex_num_literal
    | ex_num_literal addOp ex_num_literal
    | LEFT_PAREN ex_const_num_expression RIGHT_PAREN
    | ex_num_literal
    ;

ex_num_literal
    : NUMBER
    | DURATION
    | ex_time_instant_literal
    | ex_alias_call
    ;

ex_duration : ex_const_num_expression;

// This rule was converted from lexer rule TIME_RANGE.
ex_time_range: LEFT_BRACKET ex_duration RIGHT_BRACKET;

// This rule was converted from lexer rule SUBQUERY_RANGE.
ex_subquery_range: LEFT_BRACKET ex_duration ':' ex_duration? RIGHT_BRACKET;

// PROMQLX: entry rule now is promqlx which can consist of multiple vectorOperation.
// expression
//    : vectorOperation EOF
//    ;

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
    // PROMQLX: an alias or macro call can be used anywhere in the context of vector operation.
    | ex_macro_call
    | ex_alias_call
    ;

subqueryOp
    // PROMQLX: SUBQUERY_RANGE was converted from token to rule.
    : ex_subquery_range offsetOp?
    ;

offsetOp
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : OFFSET ex_duration
    ;

matrixSelector
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : instantSelector ex_time_range
    ;

offset
    // PROMQLX: Constant duration expressions can be used everywhere
    // where PromQL exprects duration.
    : instantSelector OFFSET ex_duration
    | matrixSelector OFFSET ex_duration
    ;

literal
    // PROMQLX: everything that can be converted to the number of seconds
    // can be used as literal
    : ex_const_num_expression
    | string
    ;

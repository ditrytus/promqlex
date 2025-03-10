// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar PromQLExtensionLexer;

import PromQLLexer;

// All keywords in PromQL are case insensitive, it is just function,
// label and metric names that are not.
options {
    caseInsensitive = true;
}

// PROMQLX: this rule got converted into ex_subquery_range parser rule.
fragment SUBQUERY_RANGE: ;

// PROMQLX: this rule got converted into ex_time_range parser rule.
fragment TIME_RANGE: ;

// PROMQX: extensions

EX_ID options {
  caseInsensitive = false;
}: [a-zA-Z] [0-9a-zA-Z_]+;

EX_IF: 'if';

EX_TRUE: 'true';
EX_FALSE: 'false';

EX_T: 'T';
EX_COLON: ':';
EX_DOT: '.';

EX_POSITIVE_INTEGER: [0-9] | [1-9]+ [0-9]*;
EX_TWO_DIGITS: [0-9] [0-9];
EX_DIGITS: [0-9]+;

EX_METRIC_KEYWORD: 'metric';
EX_LABEL_KEYWORD: 'label';

EX_DEF: 'def';
EX_CALL_SIGN: '$';

EX_NL: '\n' | '\r\n' ;
// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar PromQLExLexer;

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

ID options {
  caseInsensitive = false;
}: [a-zA-Z] [0-9a-zA-Z_]+;

IF: 'if';

TRUE: 'true';
FALSE: 'false';

T: 'T';
COLON: ':';
DOT: '.';

POSITIVE_INTEGER: [0-9] | [1-9]+ [0-9]*;
TWO_DIGITS: [0-9] [0-9];
DIGITS: [0-9]+;

METRIC_KEYWORD: 'metric';
LABEL_KEYWORD: 'label';

DEF: 'def';
CALL_SIGN: '$';

NL: '\n' | '\r\n' ;
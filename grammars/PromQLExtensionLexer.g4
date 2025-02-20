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

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar PromQLExtensionLexer;

channels {
    WHITESPACE,
    COMMENTS
}

// All keywords in PromQL are case insensitive, it is just function,
// label and metric names that are not.
options {
    caseInsensitive = true;
}

// Original PROMQL lexer

fragment NUMERAL: [0-9]+ ('.' [0-9]+)?;

fragment SCIENTIFIC_NUMBER: NUMERAL ('e' [-+]? NUMERAL)?;

NUMBER: NUMERAL | SCIENTIFIC_NUMBER;

STRING: '\'' (~('\'' | '\\') | '\\' .)* '\'' | '"' (~('"' | '\\') | '\\' .)* '"';

// Binary operators

ADD  : '+';
SUB  : '-';
MULT : '*';
DIV  : '/';
MOD  : '%';
POW  : '^';

AND    : 'and';
OR     : 'or';
UNLESS : 'unless';

// Comparison operators

EQ  : '=';
DEQ : '==';
NE  : '!=';
GT  : '>';
LT  : '<';
GE  : '>=';
LE  : '<=';
RE  : '=~';
NRE : '!~';

// Aggregation modifiers

BY      : 'by';
WITHOUT : 'without';

// Join modifiers

ON          : 'on';
IGNORING    : 'ignoring';
GROUP_LEFT  : 'group_left';
GROUP_RIGHT : 'group_right';

OFFSET: 'offset';

BOOL: 'bool';

AGGREGATION_OPERATOR:
    'sum'
    | 'min'
    | 'max'
    | 'avg'
    | 'group'
    | 'stddev'
    | 'stdvar'
    | 'count'
    | 'count_values'
    | 'bottomk'
    | 'topk'
    | 'quantile'
;

FUNCTION options {
    caseInsensitive = false;
}: LABEL_NAME ;

LEFT_BRACE  : '{';
RIGHT_BRACE : '}';

LEFT_PAREN  : '(';
RIGHT_PAREN : ')';

LEFT_BRACKET  : '[';
RIGHT_BRACKET : ']';

COMMA: ',';

AT: '@';

// PROMQLX: this rule got converted into ex_subquery_range parser rule.
// SUBQUERY_RANGE: LEFT_BRACKET DURATION ':' DURATION? RIGHT_BRACKET;

// PROMQLX: this rule got converted into ex_time_range parser rule.
// TIME_RANGE: LEFT_BRACKET DURATION RIGHT_BRACKET;

// The proper order (longest to the shortest) must be validated after parsing
DURATION: ([0-9]+ ('ms' | [smhdwy]))+;

METRIC_NAME : [a-z_:] [a-z0-9_:]*;
LABEL_NAME  : [a-z_] [a-z0-9_]*;

WS         : [\r\t\n ]+   -> channel(WHITESPACE);
SL_COMMENT : '#' .*? '\n' -> channel(COMMENTS);

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
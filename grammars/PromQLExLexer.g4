// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar PromQLExLexer;

import PromQLLexer;

@header {

import (
    "slices"
    "strings"
)

}

@members {

type BracketCounter interface {
    BracketCount() int
    SetBracketCount(val int)
}

func (l *PromQLExLexer) IsLiteralName(text string) bool {
	return slices.ContainsFunc(l.LiteralNames, func(literlName string) bool {
		if strings.ToLower(strings.Trim(literlName, "'")) == strings.ToLower(text) {
			return true
		}
		return false
	})
}

}

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

METRIC_NAME : { func() bool {
        cnt, ok := p.GetInputStream().(BracketCounter)
        return (!ok || (ok && cnt.BracketCount() == 0))
    }()
}? [a-z_:] [a-z0-9_:]* { !p.IsLiteralName(p.GetText()) }? {
   if prov, ok := l.GetInputStream().(FunctionsProvider); ok {
       if tokenType, ok := prov.GetTokenType(l.GetText()); ok {
           l.SetType(tokenType)
       }
   }
};

IF: 'if';

TRUE: 'true';
FALSE: 'false';

T: 'T';
COLON: ':';
SEMICOLON: ';';
DOT: '.';

POSITIVE_INTEGER: [0-9] | [1-9]+ [0-9]*;
TWO_DIGITS: [0-9] [0-9];
DIGITS: [0-9]+;

METRIC_KEYWORD: 'metric';
LABEL_KEYWORD: 'label';

DEF: 'def' -> pushMode(ID_MODE);
CALL_SIGN: '$' -> pushMode(ID_MODE);

LEFT_BRACKET  : '[' {
    if cnt, ok := l.GetInputStream().(BracketCounter); ok {
        cnt.SetBracketCount(cnt.BracketCount()+1)
    }
};

RIGHT_BRACKET : ']' {
    if cnt, ok := l.GetInputStream().(BracketCounter); ok {
        cnt.SetBracketCount(cnt.BracketCount()-1)
    }
};

mode ID_MODE;

ID options {
  caseInsensitive = false;
}: [a-zA-Z] [0-9a-zA-Z_]* -> popMode;
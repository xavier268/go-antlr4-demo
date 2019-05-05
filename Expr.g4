// Expr.g4
grammar Expr;


//============ Rules =================

// The start rule; begin parsing here.
prog:   stat+ ;

stat:   expr NEWLINE          # printExpr
    |   ID '=' expr NEWLINE   # assign
    |   NEWLINE               # blank
    ;

expr:   '-' expr	            # Unary
    |   expr '*' expr         # Mul
    |   expr '/' expr         # Div
    |   expr '+' expr         # Add
    |   expr '-' expr         # Sub
    |   INT                   # int
    |   ID                    # ident
    |   '(' expr ')'          # parenth
    ;

// ============ Tokens ===============

ID  :   [a-zA-Z]+ ;      // match identifiers
INT :   [0-9]+ ;         // match integers
NEWLINE:'\r'? '\n' ;     // return newlines to parser (is end-statement signal)
WS  :   [ \t]+ -> skip ; // toss out whitespace

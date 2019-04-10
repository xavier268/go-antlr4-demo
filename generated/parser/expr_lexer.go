// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 61, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 6, 9, 41, 10, 9, 13, 9, 14, 9, 42, 3, 10, 6, 10, 46, 10,
	10, 13, 10, 14, 10, 47, 3, 11, 5, 11, 51, 10, 11, 3, 11, 3, 11, 3, 12,
	6, 12, 56, 10, 12, 13, 12, 14, 12, 57, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3,
	2, 5, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 64,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2,
	2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33,
	3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 37, 3, 2, 2, 2, 17, 40, 3, 2, 2, 2,
	19, 45, 3, 2, 2, 2, 21, 50, 3, 2, 2, 2, 23, 55, 3, 2, 2, 2, 25, 26, 7,
	63, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 44, 2, 2, 28, 6, 3, 2, 2, 2, 29,
	30, 7, 49, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 45, 2, 2, 32, 10, 3, 2,
	2, 2, 33, 34, 7, 47, 2, 2, 34, 12, 3, 2, 2, 2, 35, 36, 7, 42, 2, 2, 36,
	14, 3, 2, 2, 2, 37, 38, 7, 43, 2, 2, 38, 16, 3, 2, 2, 2, 39, 41, 9, 2,
	2, 2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 18, 3, 2, 2, 2, 44, 46, 9, 3, 2, 2, 45, 44, 3, 2, 2, 2,
	46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 20, 3,
	2, 2, 2, 49, 51, 7, 15, 2, 2, 50, 49, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 53, 7, 12, 2, 2, 53, 22, 3, 2, 2, 2, 54, 56, 9, 4,
	2, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 8, 12, 2, 2, 60, 24, 3, 2, 2, 2,
	7, 2, 42, 47, 50, 57, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ID", "INT", "NEWLINE", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "INT", "NEWLINE",
	"WS",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExprLexer(input antlr.CharStream) *ExprLexer {

	l := new(ExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0    = 1
	ExprLexerT__1    = 2
	ExprLexerT__2    = 3
	ExprLexerT__3    = 4
	ExprLexerT__4    = 5
	ExprLexerT__5    = 6
	ExprLexerT__6    = 7
	ExprLexerID      = 8
	ExprLexerINT     = 9
	ExprLexerNEWLINE = 10
	ExprLexerWS      = 11
)

// Code generated from json.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 87, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 58, 10, 11,
	12, 11, 14, 11, 61, 11, 11, 3, 11, 3, 11, 3, 12, 5, 12, 66, 10, 12, 3,
	12, 6, 12, 69, 10, 12, 13, 12, 14, 12, 70, 3, 12, 5, 12, 74, 10, 12, 3,
	12, 6, 12, 77, 10, 12, 13, 12, 14, 12, 78, 3, 13, 6, 13, 82, 10, 13, 13,
	13, 14, 13, 83, 3, 13, 3, 13, 2, 2, 14, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 3, 2, 5, 7, 2, 34,
	34, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 92, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 3, 27, 3, 2, 2, 2, 5, 32, 3, 2, 2, 2, 7, 38, 3, 2,
	2, 2, 9, 43, 3, 2, 2, 2, 11, 45, 3, 2, 2, 2, 13, 47, 3, 2, 2, 2, 15, 49,
	3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19, 53, 3, 2, 2, 2, 21, 55, 3, 2, 2, 2,
	23, 65, 3, 2, 2, 2, 25, 81, 3, 2, 2, 2, 27, 28, 7, 112, 2, 2, 28, 29, 7,
	119, 2, 2, 29, 30, 7, 110, 2, 2, 30, 31, 7, 110, 2, 2, 31, 4, 3, 2, 2,
	2, 32, 33, 7, 104, 2, 2, 33, 34, 7, 99, 2, 2, 34, 35, 7, 110, 2, 2, 35,
	36, 7, 117, 2, 2, 36, 37, 7, 103, 2, 2, 37, 6, 3, 2, 2, 2, 38, 39, 7, 118,
	2, 2, 39, 40, 7, 116, 2, 2, 40, 41, 7, 119, 2, 2, 41, 42, 7, 103, 2, 2,
	42, 8, 3, 2, 2, 2, 43, 44, 7, 60, 2, 2, 44, 10, 3, 2, 2, 2, 45, 46, 7,
	125, 2, 2, 46, 12, 3, 2, 2, 2, 47, 48, 7, 46, 2, 2, 48, 14, 3, 2, 2, 2,
	49, 50, 7, 127, 2, 2, 50, 16, 3, 2, 2, 2, 51, 52, 7, 93, 2, 2, 52, 18,
	3, 2, 2, 2, 53, 54, 7, 95, 2, 2, 54, 20, 3, 2, 2, 2, 55, 59, 7, 36, 2,
	2, 56, 58, 9, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	62, 63, 7, 36, 2, 2, 63, 22, 3, 2, 2, 2, 64, 66, 7, 47, 2, 2, 65, 64, 3,
	2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 73, 3, 2, 2, 2, 67, 69, 9, 3, 2, 2, 68,
	67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2,
	2, 71, 72, 3, 2, 2, 2, 72, 74, 7, 48, 2, 2, 73, 68, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 77, 9, 3, 2, 2, 76, 75, 3, 2, 2, 2,
	77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 24, 3,
	2, 2, 2, 80, 82, 9, 4, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 8, 13,
	2, 2, 86, 26, 3, 2, 2, 2, 9, 2, 59, 65, 70, 73, 78, 83, 3, 8, 2, 2,
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
	"", "'null'", "'false'", "'true'", "':'", "'{'", "','", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "Number", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"STRING", "Number", "WS",
}

type jsonLexer struct {
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

func NewjsonLexer(input antlr.CharStream) *jsonLexer {

	l := new(jsonLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "json.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// jsonLexer tokens.
const (
	jsonLexerT__0   = 1
	jsonLexerT__1   = 2
	jsonLexerT__2   = 3
	jsonLexerT__3   = 4
	jsonLexerT__4   = 5
	jsonLexerT__5   = 6
	jsonLexerT__6   = 7
	jsonLexerT__7   = 8
	jsonLexerT__8   = 9
	jsonLexerSTRING = 10
	jsonLexerNumber = 11
	jsonLexerWS     = 12
)

// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 78, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 6, 11, 52, 10, 11, 13, 11, 14, 11, 53, 3, 11, 3, 11, 3, 12,
	3, 12, 7, 12, 60, 10, 12, 12, 12, 14, 12, 63, 11, 12, 3, 12, 3, 12, 3,
	13, 6, 13, 68, 10, 13, 13, 13, 14, 13, 69, 3, 14, 6, 14, 73, 10, 14, 13,
	14, 14, 14, 74, 3, 14, 3, 14, 3, 61, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 3, 2,
	5, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 81, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2,
	2, 7, 34, 3, 2, 2, 2, 9, 36, 3, 2, 2, 2, 11, 38, 3, 2, 2, 2, 13, 40, 3,
	2, 2, 2, 15, 43, 3, 2, 2, 2, 17, 45, 3, 2, 2, 2, 19, 47, 3, 2, 2, 2, 21,
	49, 3, 2, 2, 2, 23, 57, 3, 2, 2, 2, 25, 67, 3, 2, 2, 2, 27, 72, 3, 2, 2,
	2, 29, 30, 7, 60, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 125, 2, 2, 32, 33,
	7, 127, 2, 2, 33, 6, 3, 2, 2, 2, 34, 35, 7, 125, 2, 2, 35, 8, 3, 2, 2,
	2, 36, 37, 7, 127, 2, 2, 37, 10, 3, 2, 2, 2, 38, 39, 7, 46, 2, 2, 39, 12,
	3, 2, 2, 2, 40, 41, 7, 93, 2, 2, 41, 42, 7, 95, 2, 2, 42, 14, 3, 2, 2,
	2, 43, 44, 7, 93, 2, 2, 44, 16, 3, 2, 2, 2, 45, 46, 7, 95, 2, 2, 46, 18,
	3, 2, 2, 2, 47, 48, 7, 36, 2, 2, 48, 20, 3, 2, 2, 2, 49, 51, 5, 19, 10,
	2, 50, 52, 9, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 5, 19, 10,
	2, 56, 22, 3, 2, 2, 2, 57, 61, 5, 19, 10, 2, 58, 60, 11, 2, 2, 2, 59, 58,
	3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2,
	62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 5, 19, 10, 2, 65, 24, 3,
	2, 2, 2, 66, 68, 9, 3, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69,
	67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 26, 3, 2, 2, 2, 71, 73, 9, 4, 2,
	2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 8, 14, 2, 2, 77, 28, 3, 2, 2, 2,
	7, 2, 53, 61, 69, 74, 3, 8, 2, 2,
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
	"", "':'", "'{}'", "'{'", "'}'", "','", "'[]'", "'['", "']'", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "DQM", "Key", "Str", "Int", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "DQM",
	"Key", "Str", "Int", "WS",
}

type JsonLexer struct {
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

func NewJsonLexer(input antlr.CharStream) *JsonLexer {

	l := new(JsonLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Json.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonLexer tokens.
const (
	JsonLexerT__0 = 1
	JsonLexerT__1 = 2
	JsonLexerT__2 = 3
	JsonLexerT__3 = 4
	JsonLexerT__4 = 5
	JsonLexerT__5 = 6
	JsonLexerT__6 = 7
	JsonLexerT__7 = 8
	JsonLexerDQM  = 9
	JsonLexerKey  = 10
	JsonLexerStr  = 11
	JsonLexerInt  = 12
	JsonLexerWS   = 13
)

// Code generated from Json.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 6, 9, 42, 10, 9, 13, 9, 14, 9, 43, 3, 9, 3, 9, 3,
	10, 3, 10, 7, 10, 50, 10, 10, 12, 10, 14, 10, 53, 11, 10, 3, 10, 3, 10,
	3, 11, 6, 11, 58, 10, 11, 13, 11, 14, 11, 59, 3, 12, 6, 12, 63, 10, 12,
	13, 12, 14, 12, 64, 3, 12, 3, 12, 3, 51, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 5, 5, 2, 67,
	92, 97, 97, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 71,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2,
	2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33,
	3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 37, 3, 2, 2, 2, 17, 39, 3, 2, 2, 2,
	19, 47, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 62, 3, 2, 2, 2, 25, 26, 7,
	60, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 125, 2, 2, 28, 6, 3, 2, 2, 2, 29,
	30, 7, 127, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 46, 2, 2, 32, 10, 3, 2,
	2, 2, 33, 34, 7, 93, 2, 2, 34, 12, 3, 2, 2, 2, 35, 36, 7, 95, 2, 2, 36,
	14, 3, 2, 2, 2, 37, 38, 7, 36, 2, 2, 38, 16, 3, 2, 2, 2, 39, 41, 5, 15,
	8, 2, 40, 42, 9, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41,
	3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 5, 15, 8, 2,
	46, 18, 3, 2, 2, 2, 47, 51, 5, 15, 8, 2, 48, 50, 11, 2, 2, 2, 49, 48, 3,
	2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52,
	54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 5, 15, 8, 2, 55, 20, 3, 2,
	2, 2, 56, 58, 9, 3, 2, 2, 57, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57,
	3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 22, 3, 2, 2, 2, 61, 63, 9, 4, 2, 2,
	62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3,
	2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 8, 12, 2, 2, 67, 24, 3, 2, 2, 2, 7,
	2, 43, 51, 59, 64, 3, 8, 2, 2,
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
	"", "':'", "'{'", "'}'", "','", "'['", "']'", "'\"'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "DQM", "Key", "Str", "Int", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "DQM", "Key", "Str", "Int",
	"WS",
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
	JsonLexerDQM  = 7
	JsonLexerKey  = 8
	JsonLexerStr  = 9
	JsonLexerInt  = 10
	JsonLexerWS   = 11
)

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 89, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11,
	60, 10, 11, 12, 11, 14, 11, 63, 11, 11, 3, 11, 3, 11, 3, 12, 6, 12, 68,
	10, 12, 13, 12, 14, 12, 69, 3, 13, 6, 13, 73, 10, 13, 13, 13, 14, 13, 74,
	3, 13, 3, 13, 6, 13, 79, 10, 13, 13, 13, 14, 13, 80, 3, 14, 6, 14, 84,
	10, 14, 13, 14, 14, 14, 85, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	3, 2, 5, 7, 2, 34, 34, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 5,
	2, 11, 12, 15, 15, 34, 34, 2, 93, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2,
	2, 2, 5, 34, 3, 2, 2, 2, 7, 40, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 47,
	3, 2, 2, 2, 13, 49, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 53, 3, 2, 2, 2,
	19, 55, 3, 2, 2, 2, 21, 57, 3, 2, 2, 2, 23, 67, 3, 2, 2, 2, 25, 72, 3,
	2, 2, 2, 27, 83, 3, 2, 2, 2, 29, 30, 7, 118, 2, 2, 30, 31, 7, 116, 2, 2,
	31, 32, 7, 119, 2, 2, 32, 33, 7, 103, 2, 2, 33, 4, 3, 2, 2, 2, 34, 35,
	7, 104, 2, 2, 35, 36, 7, 99, 2, 2, 36, 37, 7, 110, 2, 2, 37, 38, 7, 117,
	2, 2, 38, 39, 7, 103, 2, 2, 39, 6, 3, 2, 2, 2, 40, 41, 7, 112, 2, 2, 41,
	42, 7, 119, 2, 2, 42, 43, 7, 110, 2, 2, 43, 44, 7, 110, 2, 2, 44, 8, 3,
	2, 2, 2, 45, 46, 7, 60, 2, 2, 46, 10, 3, 2, 2, 2, 47, 48, 7, 125, 2, 2,
	48, 12, 3, 2, 2, 2, 49, 50, 7, 46, 2, 2, 50, 14, 3, 2, 2, 2, 51, 52, 7,
	127, 2, 2, 52, 16, 3, 2, 2, 2, 53, 54, 7, 93, 2, 2, 54, 18, 3, 2, 2, 2,
	55, 56, 7, 95, 2, 2, 56, 20, 3, 2, 2, 2, 57, 61, 7, 36, 2, 2, 58, 60, 9,
	2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61,
	62, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 36,
	2, 2, 65, 22, 3, 2, 2, 2, 66, 68, 9, 3, 2, 2, 67, 66, 3, 2, 2, 2, 68, 69,
	3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 24, 3, 2, 2, 2,
	71, 73, 9, 3, 2, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3,
	2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 7, 48, 2, 2, 77,
	79, 9, 3, 2, 2, 78, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 78, 3, 2, 2,
	2, 80, 81, 3, 2, 2, 2, 81, 26, 3, 2, 2, 2, 82, 84, 9, 4, 2, 2, 83, 82,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2,
	86, 87, 3, 2, 2, 2, 87, 88, 8, 14, 2, 2, 88, 28, 3, 2, 2, 2, 8, 2, 61,
	69, 74, 80, 85, 3, 8, 2, 2,
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
	"", "'true'", "'false'", "'null'", "':'", "'{'", "','", "'}'", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "Int", "Float64", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"STRING", "Int", "Float64", "WS",
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
	JsonLexerT__0    = 1
	JsonLexerT__1    = 2
	JsonLexerT__2    = 3
	JsonLexerT__3    = 4
	JsonLexerT__4    = 5
	JsonLexerT__5    = 6
	JsonLexerT__6    = 7
	JsonLexerT__7    = 8
	JsonLexerT__8    = 9
	JsonLexerSTRING  = 10
	JsonLexerInt     = 11
	JsonLexerFloat64 = 12
	JsonLexerWS      = 13
)

// Code generated from json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // json

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 5, 2, 16, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7,
	4, 26, 10, 4, 12, 4, 14, 4, 29, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 35,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5, 44, 11, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 50, 10, 5, 3, 6, 3, 6, 5, 6, 54, 10, 6, 3,
	6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 3, 4, 2, 3, 5, 12, 13, 2, 57, 2, 15, 3,
	2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 53,
	3, 2, 2, 2, 12, 16, 9, 2, 2, 2, 13, 16, 5, 8, 5, 2, 14, 16, 5, 6, 4, 2,
	15, 12, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 14, 3, 2, 2, 2, 16, 3, 3, 2,
	2, 2, 17, 18, 7, 12, 2, 2, 18, 19, 7, 6, 2, 2, 19, 20, 5, 2, 2, 2, 20,
	5, 3, 2, 2, 2, 21, 22, 7, 7, 2, 2, 22, 27, 5, 4, 3, 2, 23, 24, 7, 8, 2,
	2, 24, 26, 5, 4, 3, 2, 25, 23, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25,
	3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 30, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2,
	30, 31, 7, 9, 2, 2, 31, 35, 3, 2, 2, 2, 32, 33, 7, 7, 2, 2, 33, 35, 7,
	9, 2, 2, 34, 21, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 7, 3, 2, 2, 2, 36,
	37, 7, 10, 2, 2, 37, 42, 5, 2, 2, 2, 38, 39, 7, 8, 2, 2, 39, 41, 5, 2,
	2, 2, 40, 38, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7, 11, 2, 2,
	46, 50, 3, 2, 2, 2, 47, 48, 7, 10, 2, 2, 48, 50, 7, 11, 2, 2, 49, 36, 3,
	2, 2, 2, 49, 47, 3, 2, 2, 2, 50, 9, 3, 2, 2, 2, 51, 54, 5, 6, 4, 2, 52,
	54, 5, 8, 5, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 11, 3, 2, 2,
	2, 8, 15, 27, 34, 42, 49, 53,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'null'", "'false'", "'true'", "':'", "'{'", "','", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "Number", "WS",
}

var ruleNames = []string{
	"value", "pair", "object", "array", "json",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type jsonParser struct {
	*antlr.BaseParser
}

func NewjsonParser(input antlr.TokenStream) *jsonParser {
	this := new(jsonParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "json.g4"

	return this
}

// jsonParser tokens.
const (
	jsonParserEOF    = antlr.TokenEOF
	jsonParserT__0   = 1
	jsonParserT__1   = 2
	jsonParserT__2   = 3
	jsonParserT__3   = 4
	jsonParserT__4   = 5
	jsonParserT__5   = 6
	jsonParserT__6   = 7
	jsonParserT__7   = 8
	jsonParserT__8   = 9
	jsonParserSTRING = 10
	jsonParserNumber = 11
	jsonParserWS     = 12
)

// jsonParser rules.
const (
	jsonParserRULE_value  = 0
	jsonParserRULE_pair   = 1
	jsonParserRULE_object = 2
	jsonParserRULE_array  = 3
	jsonParserRULE_json   = 4
)

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleValueContext struct {
	*ValueContext
}

func NewSingleValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleValueContext {
	var p = new(SingleValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *SingleValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(jsonParserSTRING, 0)
}

func (s *SingleValueContext) Number() antlr.TerminalNode {
	return s.GetToken(jsonParserNumber, 0)
}

func (s *SingleValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterSingleValue(s)
	}
}

func (s *SingleValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitSingleValue(s)
	}
}

type ARRAYContext struct {
	*ValueContext
}

func NewARRAYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ARRAYContext {
	var p = new(ARRAYContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ARRAYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ARRAYContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ARRAYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterARRAY(s)
	}
}

func (s *ARRAYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitARRAY(s)
	}
}

type OBJECTContext struct {
	*ValueContext
}

func NewOBJECTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OBJECTContext {
	var p = new(OBJECTContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *OBJECTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OBJECTContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *OBJECTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterOBJECT(s)
	}
}

func (s *OBJECTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitOBJECT(s)
	}
}

func (p *jsonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, jsonParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jsonParserT__0, jsonParserT__1, jsonParserT__2, jsonParserSTRING, jsonParserNumber:
		localctx = NewSingleValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<jsonParserT__0)|(1<<jsonParserT__1)|(1<<jsonParserT__2)|(1<<jsonParserSTRING)|(1<<jsonParserNumber))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case jsonParserT__7:
		localctx = NewARRAYContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Array()
		}

	case jsonParserT__4:
		localctx = NewOBJECTContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Object()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(jsonParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *jsonParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, jsonParserRULE_pair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.Match(jsonParserSTRING)
	}
	{
		p.SetState(16)
		p.Match(jsonParserT__3)
	}
	{
		p.SetState(17)
		p.Value()
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *ObjectContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *jsonParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, jsonParserRULE_object)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Match(jsonParserT__4)
		}
		{
			p.SetState(20)
			p.Pair()
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == jsonParserT__5 {
			{
				p.SetState(21)
				p.Match(jsonParserT__5)
			}
			{
				p.SetState(22)
				p.Pair()
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(28)
			p.Match(jsonParserT__6)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(jsonParserT__4)
		}
		{
			p.SetState(31)
			p.Match(jsonParserT__6)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *jsonParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, jsonParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(jsonParserT__7)
		}
		{
			p.SetState(35)
			p.Value()
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == jsonParserT__5 {
			{
				p.SetState(36)
				p.Match(jsonParserT__5)
			}
			{
				p.SetState(37)
				p.Value()
			}

			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(43)
			p.Match(jsonParserT__8)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(jsonParserT__7)
		}
		{
			p.SetState(46)
			p.Match(jsonParserT__8)
		}

	}

	return localctx
}

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = jsonParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = jsonParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *JsonContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(jsonListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *jsonParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, jsonParserRULE_json)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case jsonParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Object()
		}

	case jsonParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

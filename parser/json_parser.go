// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Json

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 63, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 23, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 33, 10, 5, 12, 5, 14, 5, 36,
	11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 42, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 48, 10, 6, 12, 6, 14, 6, 51, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	57, 10, 6, 3, 7, 3, 7, 5, 7, 61, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8, 10,
	12, 2, 4, 3, 2, 13, 14, 3, 2, 3, 4, 2, 66, 2, 14, 3, 2, 2, 2, 4, 22, 3,
	2, 2, 2, 6, 24, 3, 2, 2, 2, 8, 41, 3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12,
	60, 3, 2, 2, 2, 14, 15, 9, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 23, 9, 2, 2,
	2, 17, 23, 7, 12, 2, 2, 18, 23, 9, 3, 2, 2, 19, 23, 7, 5, 2, 2, 20, 23,
	5, 10, 6, 2, 21, 23, 5, 8, 5, 2, 22, 16, 3, 2, 2, 2, 22, 17, 3, 2, 2, 2,
	22, 18, 3, 2, 2, 2, 22, 19, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 21, 3,
	2, 2, 2, 23, 5, 3, 2, 2, 2, 24, 25, 7, 12, 2, 2, 25, 26, 7, 6, 2, 2, 26,
	27, 5, 4, 3, 2, 27, 7, 3, 2, 2, 2, 28, 29, 7, 7, 2, 2, 29, 34, 5, 6, 4,
	2, 30, 31, 7, 8, 2, 2, 31, 33, 5, 6, 4, 2, 32, 30, 3, 2, 2, 2, 33, 36,
	3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2,
	36, 34, 3, 2, 2, 2, 37, 38, 7, 9, 2, 2, 38, 42, 3, 2, 2, 2, 39, 40, 7,
	7, 2, 2, 40, 42, 7, 9, 2, 2, 41, 28, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42,
	9, 3, 2, 2, 2, 43, 44, 7, 10, 2, 2, 44, 49, 5, 4, 3, 2, 45, 46, 7, 8, 2,
	2, 46, 48, 5, 4, 3, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2,
	52, 53, 7, 11, 2, 2, 53, 57, 3, 2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 57, 7,
	11, 2, 2, 56, 43, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 11, 3, 2, 2, 2, 58,
	61, 5, 8, 5, 2, 59, 61, 5, 10, 6, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2,
	2, 2, 61, 13, 3, 2, 2, 2, 8, 22, 34, 41, 49, 56, 60,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'true'", "'false'", "'null'", "':'", "'{'", "','", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "Int", "Float64", "WS",
}

var ruleNames = []string{
	"number", "value", "pair", "object", "array", "json",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JsonParser struct {
	*antlr.BaseParser
}

func NewJsonParser(input antlr.TokenStream) *JsonParser {
	this := new(JsonParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Json.g4"

	return this
}

// JsonParser tokens.
const (
	JsonParserEOF     = antlr.TokenEOF
	JsonParserT__0    = 1
	JsonParserT__1    = 2
	JsonParserT__2    = 3
	JsonParserT__3    = 4
	JsonParserT__4    = 5
	JsonParserT__5    = 6
	JsonParserT__6    = 7
	JsonParserT__7    = 8
	JsonParserT__8    = 9
	JsonParserSTRING  = 10
	JsonParserInt     = 11
	JsonParserFloat64 = 12
	JsonParserWS      = 13
)

// JsonParser rules.
const (
	JsonParserRULE_number = 0
	JsonParserRULE_value  = 1
	JsonParserRULE_pair   = 2
	JsonParserRULE_object = 3
	JsonParserRULE_array  = 4
	JsonParserRULE_json   = 5
)

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Int() antlr.TerminalNode {
	return s.GetToken(JsonParserInt, 0)
}

func (s *NumberContext) Float64() antlr.TerminalNode {
	return s.GetToken(JsonParserFloat64, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *JsonParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonParserRULE_number)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JsonParserInt || _la == JsonParserFloat64) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

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
	p.RuleIndex = JsonParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_value

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

type NUMBERContext struct {
	*ValueContext
	nu antlr.Token
}

func NewNUMBERContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NUMBERContext {
	var p = new(NUMBERContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NUMBERContext) GetNu() antlr.Token { return s.nu }

func (s *NUMBERContext) SetNu(v antlr.Token) { s.nu = v }

func (s *NUMBERContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NUMBERContext) Int() antlr.TerminalNode {
	return s.GetToken(JsonParserInt, 0)
}

func (s *NUMBERContext) Float64() antlr.TerminalNode {
	return s.GetToken(JsonParserFloat64, 0)
}

func (s *NUMBERContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterNUMBER(s)
	}
}

func (s *NUMBERContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitNUMBER(s)
	}
}

type NULLContext struct {
	*ValueContext
}

func NewNULLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NULLContext {
	var p = new(NULLContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NULLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NULLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterNULL(s)
	}
}

func (s *NULLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitNULL(s)
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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterARRAY(s)
	}
}

func (s *ARRAYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitARRAY(s)
	}
}

type BOOLContext struct {
	*ValueContext
}

func NewBOOLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BOOLContext {
	var p = new(BOOLContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *BOOLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BOOLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterBOOL(s)
	}
}

func (s *BOOLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitBOOL(s)
	}
}

type STRINGContext struct {
	*ValueContext
}

func NewSTRINGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *STRINGContext {
	var p = new(STRINGContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *STRINGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *STRINGContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonParserSTRING, 0)
}

func (s *STRINGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterSTRING(s)
	}
}

func (s *STRINGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitSTRING(s)
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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterOBJECT(s)
	}
}

func (s *OBJECTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitOBJECT(s)
	}
}

func (p *JsonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonParserRULE_value)
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

	p.SetState(20)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonParserInt, JsonParserFloat64:
		localctx = NewNUMBERContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*NUMBERContext).nu = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonParserInt || _la == JsonParserFloat64) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*NUMBERContext).nu = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case JsonParserSTRING:
		localctx = NewSTRINGContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Match(JsonParserSTRING)
		}

	case JsonParserT__0, JsonParserT__1:
		localctx = NewBOOLContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(16)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonParserT__0 || _la == JsonParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case JsonParserT__2:
		localctx = NewNULLContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(17)
			p.Match(JsonParserT__2)
		}

	case JsonParserT__7:
		localctx = NewARRAYContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(18)
			p.Array()
		}

	case JsonParserT__4:
		localctx = NewOBJECTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(19)
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
	p.RuleIndex = JsonParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonParserSTRING, 0)
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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *JsonParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonParserRULE_pair)

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
		p.SetState(22)
		p.Match(JsonParserSTRING)
	}
	{
		p.SetState(23)
		p.Match(JsonParserT__3)
	}
	{
		p.SetState(24)
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
	p.RuleIndex = JsonParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_object

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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *JsonParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonParserRULE_object)
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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(JsonParserT__4)
		}
		{
			p.SetState(27)
			p.Pair()
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonParserT__5 {
			{
				p.SetState(28)
				p.Match(JsonParserT__5)
			}
			{
				p.SetState(29)
				p.Pair()
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(35)
			p.Match(JsonParserT__6)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(JsonParserT__4)
		}
		{
			p.SetState(38)
			p.Match(JsonParserT__6)
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
	p.RuleIndex = JsonParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_array

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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *JsonParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonParserRULE_array)
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

	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(JsonParserT__7)
		}
		{
			p.SetState(42)
			p.Value()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JsonParserT__5 {
			{
				p.SetState(43)
				p.Match(JsonParserT__5)
			}
			{
				p.SetState(44)
				p.Value()
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(50)
			p.Match(JsonParserT__8)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(JsonParserT__7)
		}
		{
			p.SetState(53)
			p.Match(JsonParserT__8)
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
	p.RuleIndex = JsonParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_json

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
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *JsonParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonParserRULE_json)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Object()
		}

	case JsonParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

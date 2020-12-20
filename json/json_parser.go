// Code generated from Json.g4 by ANTLR 4.7.2. DO NOT EDIT.

package json // Json

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 5, 2, 18, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 33, 10, 4, 13, 4, 14, 4, 34,
	3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	6, 5, 48, 10, 5, 13, 5, 14, 5, 49, 3, 5, 3, 5, 5, 5, 54, 10, 5, 3, 6, 3,
	6, 5, 6, 58, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 65, 2, 17,
	3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10,
	57, 3, 2, 2, 2, 12, 18, 7, 14, 2, 2, 13, 18, 7, 12, 2, 2, 14, 18, 7, 13,
	2, 2, 15, 18, 5, 8, 5, 2, 16, 18, 5, 6, 4, 2, 17, 12, 3, 2, 2, 2, 17, 13,
	3, 2, 2, 2, 17, 14, 3, 2, 2, 2, 17, 15, 3, 2, 2, 2, 17, 16, 3, 2, 2, 2,
	18, 3, 3, 2, 2, 2, 19, 20, 7, 13, 2, 2, 20, 21, 7, 3, 2, 2, 21, 22, 5,
	2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 39, 7, 4, 2, 2, 24, 25, 7, 5, 2, 2, 25,
	26, 5, 4, 3, 2, 26, 27, 7, 6, 2, 2, 27, 39, 3, 2, 2, 2, 28, 29, 7, 5, 2,
	2, 29, 30, 5, 4, 3, 2, 30, 32, 7, 7, 2, 2, 31, 33, 5, 4, 3, 2, 32, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 36, 3, 2, 2, 2, 36, 37, 7, 6, 2, 2, 37, 39, 3, 2, 2, 2, 38, 23, 3,
	2, 2, 2, 38, 24, 3, 2, 2, 2, 38, 28, 3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40,
	54, 7, 8, 2, 2, 41, 42, 7, 9, 2, 2, 42, 43, 5, 2, 2, 2, 43, 44, 7, 10,
	2, 2, 44, 54, 3, 2, 2, 2, 45, 47, 7, 9, 2, 2, 46, 48, 5, 2, 2, 2, 47, 46,
	3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2,
	50, 51, 3, 2, 2, 2, 51, 52, 7, 10, 2, 2, 52, 54, 3, 2, 2, 2, 53, 40, 3,
	2, 2, 2, 53, 41, 3, 2, 2, 2, 53, 45, 3, 2, 2, 2, 54, 9, 3, 2, 2, 2, 55,
	58, 5, 8, 5, 2, 56, 58, 5, 6, 4, 2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2,
	2, 58, 11, 3, 2, 2, 2, 8, 17, 34, 38, 49, 53, 57,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'{}'", "'{'", "'}'", "','", "'[]'", "'['", "']'", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "DQM", "Key", "Str", "Int", "WS",
}

var ruleNames = []string{
	"value", "keyValue", "object", "array", "init",
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
	JsonParserEOF  = antlr.TokenEOF
	JsonParserT__0 = 1
	JsonParserT__1 = 2
	JsonParserT__2 = 3
	JsonParserT__3 = 4
	JsonParserT__4 = 5
	JsonParserT__5 = 6
	JsonParserT__6 = 7
	JsonParserT__7 = 8
	JsonParserDQM  = 9
	JsonParserKey  = 10
	JsonParserStr  = 11
	JsonParserInt  = 12
	JsonParserWS   = 13
)

// JsonParser rules.
const (
	JsonParserRULE_value    = 0
	JsonParserRULE_keyValue = 1
	JsonParserRULE_object   = 2
	JsonParserRULE_array    = 3
	JsonParserRULE_init     = 4
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

func (s *ValueContext) Int() antlr.TerminalNode {
	return s.GetToken(JsonParserInt, 0)
}

func (s *ValueContext) Key() antlr.TerminalNode {
	return s.GetToken(JsonParserKey, 0)
}

func (s *ValueContext) Str() antlr.TerminalNode {
	return s.GetToken(JsonParserStr, 0)
}

func (s *ValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *JsonParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonParserRULE_value)

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

	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonParserInt:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Match(JsonParserInt)
		}

	case JsonParserKey:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Match(JsonParserKey)
		}

	case JsonParserStr:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Match(JsonParserStr)
		}

	case JsonParserT__5, JsonParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(13)
			p.Array()
		}

	case JsonParserT__1, JsonParserT__2:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(14)
			p.Object()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeyValueContext is an interface to support dynamic dispatch.
type IKeyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValueContext differentiates from other interfaces.
	IsKeyValueContext()
}

type KeyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValueContext() *KeyValueContext {
	var p = new(KeyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonParserRULE_keyValue
	return p
}

func (*KeyValueContext) IsKeyValueContext() {}

func NewKeyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValueContext {
	var p = new(KeyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_keyValue

	return p
}

func (s *KeyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValueContext) Str() antlr.TerminalNode {
	return s.GetToken(JsonParserStr, 0)
}

func (s *KeyValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *KeyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterKeyValue(s)
	}
}

func (s *KeyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitKeyValue(s)
	}
}

func (p *JsonParser) KeyValue() (localctx IKeyValueContext) {
	localctx = NewKeyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonParserRULE_keyValue)

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
		p.SetState(17)
		p.Match(JsonParserStr)
	}
	{
		p.SetState(18)
		p.Match(JsonParserT__0)
	}
	{
		p.SetState(19)
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

func (s *ObjectContext) AllKeyValue() []IKeyValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValueContext)(nil)).Elem())
	var tst = make([]IKeyValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValueContext)
		}
	}

	return tst
}

func (s *ObjectContext) KeyValue(i int) IKeyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValueContext)
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
	p.EnterRule(localctx, 4, JsonParserRULE_object)
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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Match(JsonParserT__1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(22)
			p.Match(JsonParserT__2)
		}
		{
			p.SetState(23)
			p.KeyValue()
		}
		{
			p.SetState(24)
			p.Match(JsonParserT__3)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(JsonParserT__2)
		}
		{
			p.SetState(27)
			p.KeyValue()
		}
		{
			p.SetState(28)
			p.Match(JsonParserT__4)
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == JsonParserStr {
			{
				p.SetState(29)
				p.KeyValue()
			}

			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(34)
			p.Match(JsonParserT__3)
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
	p.EnterRule(localctx, 6, JsonParserRULE_array)
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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(JsonParserT__5)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(JsonParserT__6)
		}
		{
			p.SetState(40)
			p.Value()
		}
		{
			p.SetState(41)
			p.Match(JsonParserT__7)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Match(JsonParserT__6)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonParserT__1)|(1<<JsonParserT__2)|(1<<JsonParserT__5)|(1<<JsonParserT__6)|(1<<JsonParserKey)|(1<<JsonParserStr)|(1<<JsonParserInt))) != 0) {
			{
				p.SetState(44)
				p.Value()
			}

			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(49)
			p.Match(JsonParserT__7)
		}

	}

	return localctx
}

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *InitContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *JsonParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonParserRULE_init)

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonParserT__5, JsonParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Array()
		}

	case JsonParserT__1, JsonParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Object()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

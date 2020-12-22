// Generated from Json.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class JsonParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, String=7, Number=8, WS=9, 
		True=10, False=11, Null=12;
	public static final int
		RULE_bool = 0, RULE_value = 1, RULE_pair = 2, RULE_object = 3, RULE_array = 4, 
		RULE_init = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"bool", "value", "pair", "object", "array", "init"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'{'", "'}'", "','", "'['", "']'", null, null, null, "'true'", 
			"'false'", "'null'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "String", "Number", "WS", "True", 
			"False", "Null"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Json.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public JsonParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class BoolContext extends ParserRuleContext {
		public TerminalNode True() { return getToken(JsonParser.True, 0); }
		public TerminalNode False() { return getToken(JsonParser.False, 0); }
		public BoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterBool(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitBool(this);
		}
	}

	public final BoolContext bool() throws RecognitionException {
		BoolContext _localctx = new BoolContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_bool);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			_la = _input.LA(1);
			if ( !(_la==True || _la==False) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TerminalNode Number() { return getToken(JsonParser.Number, 0); }
		public TerminalNode String() { return getToken(JsonParser.String, 0); }
		public ArrayContext array() {
			return getRuleContext(ArrayContext.class,0);
		}
		public ObjectContext object() {
			return getRuleContext(ObjectContext.class,0);
		}
		public BoolContext bool() {
			return getRuleContext(BoolContext.class,0);
		}
		public TerminalNode Null() { return getToken(JsonParser.Null, 0); }
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterValue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitValue(this);
		}
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_value);
		try {
			setState(20);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Number:
				enterOuterAlt(_localctx, 1);
				{
				setState(14);
				match(Number);
				}
				break;
			case String:
				enterOuterAlt(_localctx, 2);
				{
				setState(15);
				match(String);
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(16);
				array();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 4);
				{
				setState(17);
				object();
				}
				break;
			case True:
			case False:
				enterOuterAlt(_localctx, 5);
				{
				setState(18);
				bool();
				}
				break;
			case Null:
				enterOuterAlt(_localctx, 6);
				{
				setState(19);
				match(Null);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PairContext extends ParserRuleContext {
		public TerminalNode String() { return getToken(JsonParser.String, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public PairContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pair; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterPair(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitPair(this);
		}
	}

	public final PairContext pair() throws RecognitionException {
		PairContext _localctx = new PairContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_pair);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			match(String);
			setState(23);
			match(T__0);
			setState(24);
			value();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectContext extends ParserRuleContext {
		public List<PairContext> pair() {
			return getRuleContexts(PairContext.class);
		}
		public PairContext pair(int i) {
			return getRuleContext(PairContext.class,i);
		}
		public ObjectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_object; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterObject(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitObject(this);
		}
	}

	public final ObjectContext object() throws RecognitionException {
		ObjectContext _localctx = new ObjectContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_object);
		int _la;
		try {
			setState(39);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				match(T__1);
				setState(27);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(28);
				match(T__1);
				setState(29);
				pair();
				setState(34);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(30);
					match(T__3);
					setState(31);
					pair();
					}
					}
					setState(36);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(37);
				match(T__2);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterArray(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitArray(this);
		}
	}

	public final ArrayContext array() throws RecognitionException {
		ArrayContext _localctx = new ArrayContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_array);
		int _la;
		try {
			setState(54);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(41);
				match(T__4);
				setState(42);
				match(T__5);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(43);
				match(T__4);
				setState(44);
				value();
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__3) {
					{
					{
					setState(45);
					match(T__3);
					setState(46);
					value();
					}
					}
					setState(51);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(52);
				match(T__5);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InitContext extends ParserRuleContext {
		public ArrayContext array() {
			return getRuleContext(ArrayContext.class,0);
		}
		public ObjectContext object() {
			return getRuleContext(ObjectContext.class,0);
		}
		public InitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_init; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).enterInit(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof JsonListener ) ((JsonListener)listener).exitInit(this);
		}
	}

	public final InitContext init() throws RecognitionException {
		InitContext _localctx = new InitContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_init);
		try {
			setState(58);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				array();
				}
				break;
			case T__1:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				object();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\16?\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\5\3"+
		"\27\n\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\7\5#\n\5\f\5\16\5&\13"+
		"\5\3\5\3\5\5\5*\n\5\3\6\3\6\3\6\3\6\3\6\3\6\7\6\62\n\6\f\6\16\6\65\13"+
		"\6\3\6\3\6\5\69\n\6\3\7\3\7\5\7=\n\7\3\7\2\2\b\2\4\6\b\n\f\2\3\3\2\f\r"+
		"\2B\2\16\3\2\2\2\4\26\3\2\2\2\6\30\3\2\2\2\b)\3\2\2\2\n8\3\2\2\2\f<\3"+
		"\2\2\2\16\17\t\2\2\2\17\3\3\2\2\2\20\27\7\n\2\2\21\27\7\t\2\2\22\27\5"+
		"\n\6\2\23\27\5\b\5\2\24\27\5\2\2\2\25\27\7\16\2\2\26\20\3\2\2\2\26\21"+
		"\3\2\2\2\26\22\3\2\2\2\26\23\3\2\2\2\26\24\3\2\2\2\26\25\3\2\2\2\27\5"+
		"\3\2\2\2\30\31\7\t\2\2\31\32\7\3\2\2\32\33\5\4\3\2\33\7\3\2\2\2\34\35"+
		"\7\4\2\2\35*\7\5\2\2\36\37\7\4\2\2\37$\5\6\4\2 !\7\6\2\2!#\5\6\4\2\" "+
		"\3\2\2\2#&\3\2\2\2$\"\3\2\2\2$%\3\2\2\2%\'\3\2\2\2&$\3\2\2\2\'(\7\5\2"+
		"\2(*\3\2\2\2)\34\3\2\2\2)\36\3\2\2\2*\t\3\2\2\2+,\7\7\2\2,9\7\b\2\2-."+
		"\7\7\2\2.\63\5\4\3\2/\60\7\6\2\2\60\62\5\4\3\2\61/\3\2\2\2\62\65\3\2\2"+
		"\2\63\61\3\2\2\2\63\64\3\2\2\2\64\66\3\2\2\2\65\63\3\2\2\2\66\67\7\b\2"+
		"\2\679\3\2\2\28+\3\2\2\28-\3\2\2\29\13\3\2\2\2:=\5\n\6\2;=\5\b\5\2<:\3"+
		"\2\2\2<;\3\2\2\2=\r\3\2\2\2\b\26$)\638<";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
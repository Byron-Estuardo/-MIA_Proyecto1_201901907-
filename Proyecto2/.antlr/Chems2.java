// Generated from /home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/Proyecto2/Chems2.g4 by ANTLR 4.8

    import "Proyecto2/nodos"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Chems2 extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		MKDISK=1, RMDISK=2, FDISK=3, MOUNT=4, MKFS=5, LOGIN=6, LOGOUT=7, MKGRP=8, 
		RMGRP=9, MKUSER=10, RMUSR=11, MKFILE=12, MKDIR=13, EXEC=14, REP=15, SIZE=16, 
		FIT=17, UNIT=18, PATH=19, NAME=20, TYPE=21, ID=22, BF=23, FF=24, WF=25, 
		FAST=26, FULL=27, USR=28, PWD=29, PWD1=30, GRP=31, R=32, CONT=33, P=34, 
		PAUSA=35, DISK=36, TREE=37, FILE=38, LETRA=39, ENTERO=40, CARACTER=41, 
		IDENTIFICADOR=42, PASSWORD=43, CADENA=44, EXTENSION=45, DIAGONAL=46, RUTA=47, 
		DIRECTORIO=48, IGUAL=49, WHITESPACE=50;
	public static final int
		RULE_start = 0, RULE_adminDiscos = 1, RULE_mkdisk = 2, RULE_parametrosmkdisk = 3, 
		RULE_ajuste = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "adminDiscos", "mkdisk", "parametrosmkdisk", "ajuste"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'mkfs'", "'login'", 
			"'logout'", "'mkgrp'", "'rmgrp'", "'mkuser'", "'rmusr'", "'mkfile'", 
			"'mkdir'", "'exec'", "'rep'", "'-size'", "'-fit'", "'-unit'", "'-path'", 
			"'-name'", "'-type'", "'-id'", "'bf'", "'ff'", "'wf'", "'fast'", "'FULL'", 
			"'-usuario'", "'-password'", "'-pwd'", "'-grp'", "'-r'", "'-cont'", "'-p'", 
			"'pause'", "'disk'", "'tree'", "'file'", null, null, null, null, null, 
			null, "'.'", "'/'", null, null, "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS", "LOGIN", "LOGOUT", 
			"MKGRP", "RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR", "EXEC", "REP", 
			"SIZE", "FIT", "UNIT", "PATH", "NAME", "TYPE", "ID", "BF", "FF", "WF", 
			"FAST", "FULL", "USR", "PWD", "PWD1", "GRP", "R", "CONT", "P", "PAUSA", 
			"DISK", "TREE", "FILE", "LETRA", "ENTERO", "CARACTER", "IDENTIFICADOR", 
			"PASSWORD", "CADENA", "EXTENSION", "DIAGONAL", "RUTA", "DIRECTORIO", 
			"IGUAL", "WHITESPACE"
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
	public String getGrammarFileName() { return "Chems2.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Chems2(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *nodos.Nodo Envio;
		public AdminDiscosContext adminDiscos;
		public AdminDiscosContext adminDiscos() {
			return getRuleContext(AdminDiscosContext.class,0);
		}
		public TerminalNode PAUSA() { return getToken(Chems2.PAUSA, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			setState(14);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
			case RMDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(10);
				((StartContext)_localctx).adminDiscos = adminDiscos();
				 _localctx.Envio = ((StartContext)_localctx).adminDiscos.Raiz  
				}
				break;
			case PAUSA:
				enterOuterAlt(_localctx, 2);
				{
				setState(13);
				match(PAUSA);
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

	public static class AdminDiscosContext extends ParserRuleContext {
		public *nodos.Nodo Raiz;
		public TerminalNode MKDISK() { return getToken(Chems2.MKDISK, 0); }
		public MkdiskContext mkdisk() {
			return getRuleContext(MkdiskContext.class,0);
		}
		public TerminalNode RMDISK() { return getToken(Chems2.RMDISK, 0); }
		public AdminDiscosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminDiscos; }
	}

	public final AdminDiscosContext adminDiscos() throws RecognitionException {
		AdminDiscosContext _localctx = new AdminDiscosContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_adminDiscos);
		try {
			setState(21);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(16);
				match(MKDISK);
				setState(17);
				mkdisk(0);
				  fmt.Println("hOLA")
				                  nodos.NewNodo("MKDISK", "")
				                  _localctx.Raiz = nodos.Raiz
				              
				}
				break;
			case RMDISK:
				enterOuterAlt(_localctx, 2);
				{
				setState(20);
				match(RMDISK);
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

	public static class MkdiskContext extends ParserRuleContext {
		public ParametrosmkdiskContext parametrosmkdisk;
		public ParametrosmkdiskContext parametrosmkdisk() {
			return getRuleContext(ParametrosmkdiskContext.class,0);
		}
		public MkdiskContext mkdisk() {
			return getRuleContext(MkdiskContext.class,0);
		}
		public MkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdisk; }
	}

	public final MkdiskContext mkdisk() throws RecognitionException {
		return mkdisk(0);
	}

	private MkdiskContext mkdisk(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		MkdiskContext _localctx = new MkdiskContext(_ctx, _parentState);
		MkdiskContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_mkdisk, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(24);
			((MkdiskContext)_localctx).parametrosmkdisk = parametrosmkdisk();
			  fmt.Println("Parametros")
			                      nodos.Add(nodos.Nodoo("PARAMETRO", ""))
			                      nodos.AddParameters(((MkdiskContext)_localctx).parametrosmkdisk.NODS) 
			}
			_ctx.stop = _input.LT(-1);
			setState(33);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new MkdiskContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_mkdisk);
					setState(27);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(28);
					((MkdiskContext)_localctx).parametrosmkdisk = parametrosmkdisk();
					 fmt.Println("hOLA 2")
					                                      nodos.AddParameters(((MkdiskContext)_localctx).parametrosmkdisk.NODS) 
					}
					} 
				}
				setState(35);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParametrosmkdiskContext extends ParserRuleContext {
		public nodos.Nodo NODS;
		public Token ENTERO;
		public AjusteContext ajuste;
		public Token CARACTER;
		public Token RUTA;
		public TerminalNode SIZE() { return getToken(Chems2.SIZE, 0); }
		public TerminalNode IGUAL() { return getToken(Chems2.IGUAL, 0); }
		public TerminalNode ENTERO() { return getToken(Chems2.ENTERO, 0); }
		public TerminalNode FIT() { return getToken(Chems2.FIT, 0); }
		public AjusteContext ajuste() {
			return getRuleContext(AjusteContext.class,0);
		}
		public TerminalNode UNIT() { return getToken(Chems2.UNIT, 0); }
		public TerminalNode CARACTER() { return getToken(Chems2.CARACTER, 0); }
		public TerminalNode PATH() { return getToken(Chems2.PATH, 0); }
		public TerminalNode RUTA() { return getToken(Chems2.RUTA, 0); }
		public ParametrosmkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkdisk; }
	}

	public final ParametrosmkdiskContext parametrosmkdisk() throws RecognitionException {
		ParametrosmkdiskContext _localctx = new ParametrosmkdiskContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_parametrosmkdisk);
		try {
			setState(53);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SIZE:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				match(SIZE);
				setState(37);
				match(IGUAL);
				setState(38);
				((ParametrosmkdiskContext)_localctx).ENTERO = match(ENTERO);
				 _localctx.NODS = nodos.Nodoo("size", (((ParametrosmkdiskContext)_localctx).ENTERO!=null?((ParametrosmkdiskContext)_localctx).ENTERO.getText():null)) 
				}
				break;
			case FIT:
				enterOuterAlt(_localctx, 2);
				{
				setState(40);
				match(FIT);
				setState(41);
				match(IGUAL);
				setState(42);
				((ParametrosmkdiskContext)_localctx).ajuste = ajuste();
				  _localctx.NODS = nodos.Nodoo("fit", "") 
				                      nodos.AddParameters(((ParametrosmkdiskContext)_localctx).ajuste.AJ) 
				}
				break;
			case UNIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(45);
				match(UNIT);
				setState(46);
				match(IGUAL);
				setState(47);
				((ParametrosmkdiskContext)_localctx).CARACTER = match(CARACTER);
				 _localctx.NODS = nodos.Nodoo("unit", (((ParametrosmkdiskContext)_localctx).CARACTER!=null?((ParametrosmkdiskContext)_localctx).CARACTER.getText():null)) 
				}
				break;
			case PATH:
				enterOuterAlt(_localctx, 4);
				{
				setState(49);
				match(PATH);
				setState(50);
				match(IGUAL);
				setState(51);
				((ParametrosmkdiskContext)_localctx).RUTA = match(RUTA);
				 _localctx.NODS = nodos.Nodoo("path", (((ParametrosmkdiskContext)_localctx).RUTA!=null?((ParametrosmkdiskContext)_localctx).RUTA.getText():null)) 
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

	public static class AjusteContext extends ParserRuleContext {
		public nodos.Nodo AJ;
		public TerminalNode BF() { return getToken(Chems2.BF, 0); }
		public TerminalNode FF() { return getToken(Chems2.FF, 0); }
		public TerminalNode WF() { return getToken(Chems2.WF, 0); }
		public AjusteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ajuste; }
	}

	public final AjusteContext ajuste() throws RecognitionException {
		AjusteContext _localctx = new AjusteContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_ajuste);
		try {
			setState(61);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BF:
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				match(BF);
				 _localctx.AJ = nodos.Nodoo("ajuste", "bf") 
				}
				break;
			case FF:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				match(FF);
				 _localctx.AJ = nodos.Nodoo("ajuste", "ff") 
				}
				break;
			case WF:
				enterOuterAlt(_localctx, 3);
				{
				setState(59);
				match(WF);
				 _localctx.AJ = nodos.Nodoo("ajuste", "wf") 
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 2:
			return mkdisk_sempred((MkdiskContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean mkdisk_sempred(MkdiskContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\64B\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\2\5\2\21\n\2\3\3\3\3\3\3\3\3"+
		"\3\3\5\3\30\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4\"\n\4\f\4\16\4%\13"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\5\58\n\5\3\6\3\6\3\6\3\6\3\6\3\6\5\6@\n\6\3\6\2\3\6\7\2\4\6\b\n\2\2\2"+
		"D\2\20\3\2\2\2\4\27\3\2\2\2\6\31\3\2\2\2\b\67\3\2\2\2\n?\3\2\2\2\f\r\5"+
		"\4\3\2\r\16\b\2\1\2\16\21\3\2\2\2\17\21\7%\2\2\20\f\3\2\2\2\20\17\3\2"+
		"\2\2\21\3\3\2\2\2\22\23\7\3\2\2\23\24\5\6\4\2\24\25\b\3\1\2\25\30\3\2"+
		"\2\2\26\30\7\4\2\2\27\22\3\2\2\2\27\26\3\2\2\2\30\5\3\2\2\2\31\32\b\4"+
		"\1\2\32\33\5\b\5\2\33\34\b\4\1\2\34#\3\2\2\2\35\36\f\4\2\2\36\37\5\b\5"+
		"\2\37 \b\4\1\2 \"\3\2\2\2!\35\3\2\2\2\"%\3\2\2\2#!\3\2\2\2#$\3\2\2\2$"+
		"\7\3\2\2\2%#\3\2\2\2&\'\7\22\2\2\'(\7\63\2\2()\7*\2\2)8\b\5\1\2*+\7\23"+
		"\2\2+,\7\63\2\2,-\5\n\6\2-.\b\5\1\2.8\3\2\2\2/\60\7\24\2\2\60\61\7\63"+
		"\2\2\61\62\7+\2\2\628\b\5\1\2\63\64\7\25\2\2\64\65\7\63\2\2\65\66\7\61"+
		"\2\2\668\b\5\1\2\67&\3\2\2\2\67*\3\2\2\2\67/\3\2\2\2\67\63\3\2\2\28\t"+
		"\3\2\2\29:\7\31\2\2:@\b\6\1\2;<\7\32\2\2<@\b\6\1\2=>\7\33\2\2>@\b\6\1"+
		"\2?9\3\2\2\2?;\3\2\2\2?=\3\2\2\2@\13\3\2\2\2\7\20\27#\67?";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
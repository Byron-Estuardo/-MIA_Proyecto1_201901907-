// Generated from /home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/MiaProyecto2/Chems.g4 by ANTLR 4.8

  import AD"MiaProyecto2/AdminDisks"
  var ValSize string = ""
  var ValFit  string = ""
  var ValUnit string = ""
  var ValPath string = ""
  var ValType string = ""
  var ValName string = ""

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Chems extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		MKDISK=1, RMDISK=2, FDISK=3, MOUNT=4, MKFS=5, LOGIN=6, LOGOUT=7, MKGRP=8, 
		RMGRP=9, MKUSER=10, RMUSR=11, MKFILE=12, MKDIR=13, EXEC=14, REP=15, SIZE=16, 
		FIT=17, UNIT=18, PATH=19, NAME=20, TYPE=21, ID=22, BF=23, FF=24, WF=25, 
		FAST=26, FULL=27, USR=28, PWD=29, PWD1=30, GRP=31, GR=32, CONT=33, GP=34, 
		PAUSA=35, DISK=36, TREE=37, FILE=38, LETRA=39, ENTERO=40, CADENA=41, IDENTIFICADOR=42, 
		RUTA=43, IGUAL=44, WHITESPACE=45, COMENTARIOS=46;
	public static final int
		RULE_start = 0, RULE_adminDiscos = 1, RULE_pmkdisk = 2, RULE_parametrosmkdisk = 3, 
		RULE_prmdisk = 4, RULE_pfdisk = 5, RULE_parametrosfdisk = 6, RULE_pmount = 7, 
		RULE_parametrosmount = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "adminDiscos", "pmkdisk", "parametrosmkdisk", "prmdisk", "pfdisk", 
			"parametrosfdisk", "pmount", "parametrosmount"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS", "LOGIN", "LOGOUT", 
			"MKGRP", "RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR", "EXEC", "REP", 
			"SIZE", "FIT", "UNIT", "PATH", "NAME", "TYPE", "ID", "BF", "FF", "WF", 
			"FAST", "FULL", "USR", "PWD", "PWD1", "GRP", "GR", "CONT", "GP", "PAUSA", 
			"DISK", "TREE", "FILE", "LETRA", "ENTERO", "CADENA", "IDENTIFICADOR", 
			"RUTA", "IGUAL", "WHITESPACE", "COMENTARIOS"
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
	public String getGrammarFileName() { return "Chems.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Chems(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public AdminDiscosContext adminDiscos() {
			return getRuleContext(AdminDiscosContext.class,0);
		}
		public TerminalNode PAUSA() { return getToken(Chems.PAUSA, 0); }
		public TerminalNode EOF() { return getToken(Chems.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			setState(22);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
			case RMDISK:
			case FDISK:
			case MOUNT:
				enterOuterAlt(_localctx, 1);
				{
				setState(18);
				adminDiscos();
				}
				break;
			case PAUSA:
				enterOuterAlt(_localctx, 2);
				{
				setState(19);
				match(PAUSA);
				 AD.Pausa() 
				}
				break;
			case EOF:
				enterOuterAlt(_localctx, 3);
				{
				setState(21);
				match(EOF);
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
		public TerminalNode MKDISK() { return getToken(Chems.MKDISK, 0); }
		public PmkdiskContext pmkdisk() {
			return getRuleContext(PmkdiskContext.class,0);
		}
		public TerminalNode RMDISK() { return getToken(Chems.RMDISK, 0); }
		public PrmdiskContext prmdisk() {
			return getRuleContext(PrmdiskContext.class,0);
		}
		public TerminalNode FDISK() { return getToken(Chems.FDISK, 0); }
		public PfdiskContext pfdisk() {
			return getRuleContext(PfdiskContext.class,0);
		}
		public TerminalNode MOUNT() { return getToken(Chems.MOUNT, 0); }
		public PmountContext pmount() {
			return getRuleContext(PmountContext.class,0);
		}
		public AdminDiscosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminDiscos; }
	}

	public final AdminDiscosContext adminDiscos() throws RecognitionException {
		AdminDiscosContext _localctx = new AdminDiscosContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_adminDiscos);
		try {
			setState(32);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(24);
				match(MKDISK);
				setState(25);
				pmkdisk(0);
				}
				break;
			case RMDISK:
				enterOuterAlt(_localctx, 2);
				{
				setState(26);
				match(RMDISK);
				setState(27);
				prmdisk();
				}
				break;
			case FDISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(28);
				match(FDISK);
				setState(29);
				pfdisk(0);
				}
				break;
			case MOUNT:
				enterOuterAlt(_localctx, 4);
				{
				setState(30);
				match(MOUNT);
				setState(31);
				pmount(0);
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

	public static class PmkdiskContext extends ParserRuleContext {
		public ParametrosmkdiskContext parametrosmkdisk() {
			return getRuleContext(ParametrosmkdiskContext.class,0);
		}
		public PmkdiskContext pmkdisk() {
			return getRuleContext(PmkdiskContext.class,0);
		}
		public PmkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmkdisk; }
	}

	public final PmkdiskContext pmkdisk() throws RecognitionException {
		return pmkdisk(0);
	}

	private PmkdiskContext pmkdisk(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PmkdiskContext _localctx = new PmkdiskContext(_ctx, _parentState);
		PmkdiskContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_pmkdisk, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(35);
			parametrosmkdisk();
			 AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath); ValSize=""  
			}
			_ctx.stop = _input.LT(-1);
			setState(44);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmkdiskContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmkdisk);
					setState(38);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(39);
					parametrosmkdisk();
					  
					}
					} 
				}
				setState(46);
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
		public Token ENTERO;
		public Token CADENA;
		public Token RUTA;
		public TerminalNode SIZE() { return getToken(Chems.SIZE, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode ENTERO() { return getToken(Chems.ENTERO, 0); }
		public TerminalNode FIT() { return getToken(Chems.FIT, 0); }
		public TerminalNode BF() { return getToken(Chems.BF, 0); }
		public TerminalNode FF() { return getToken(Chems.FF, 0); }
		public TerminalNode WF() { return getToken(Chems.WF, 0); }
		public TerminalNode UNIT() { return getToken(Chems.UNIT, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public ParametrosmkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkdisk; }
	}

	public final ParametrosmkdiskContext parametrosmkdisk() throws RecognitionException {
		ParametrosmkdiskContext _localctx = new ParametrosmkdiskContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_parametrosmkdisk);
		try {
			setState(71);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				match(SIZE);
				setState(48);
				match(IGUAL);
				setState(49);
				((ParametrosmkdiskContext)_localctx).ENTERO = match(ENTERO);
				 ValSize = (((ParametrosmkdiskContext)_localctx).ENTERO!=null?((ParametrosmkdiskContext)_localctx).ENTERO.getText():null) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				match(FIT);
				setState(52);
				match(IGUAL);
				setState(53);
				match(BF);
				 ValFit = "b" 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				match(FIT);
				setState(56);
				match(IGUAL);
				setState(57);
				match(FF);
				 ValFit = "f" 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(59);
				match(FIT);
				setState(60);
				match(IGUAL);
				setState(61);
				match(WF);
				 ValFit = "w" 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(63);
				match(UNIT);
				setState(64);
				match(IGUAL);
				setState(65);
				((ParametrosmkdiskContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkdiskContext)_localctx).CADENA!=null?((ParametrosmkdiskContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkdiskContext)_localctx).CADENA!=null?((ParametrosmkdiskContext)_localctx).CADENA.getText():null))-1] ValUnit = str 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(67);
				match(PATH);
				setState(68);
				match(IGUAL);
				setState(69);
				((ParametrosmkdiskContext)_localctx).RUTA = match(RUTA);
				 ValPath = ((ParametrosmkdiskContext)_localctx).RUTA 
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

	public static class PrmdiskContext extends ParserRuleContext {
		public Token RUTA;
		public TerminalNode RMDISK() { return getToken(Chems.RMDISK, 0); }
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public PrmdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prmdisk; }
	}

	public final PrmdiskContext prmdisk() throws RecognitionException {
		PrmdiskContext _localctx = new PrmdiskContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_prmdisk);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(RMDISK);
			setState(74);
			match(PATH);
			setState(75);
			match(IGUAL);
			setState(76);
			((PrmdiskContext)_localctx).RUTA = match(RUTA);
			 ValPath = ((PrmdiskContext)_localctx).RUTA 
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

	public static class PfdiskContext extends ParserRuleContext {
		public ParametrosfdiskContext parametrosfdisk() {
			return getRuleContext(ParametrosfdiskContext.class,0);
		}
		public PfdiskContext pfdisk() {
			return getRuleContext(PfdiskContext.class,0);
		}
		public PfdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pfdisk; }
	}

	public final PfdiskContext pfdisk() throws RecognitionException {
		return pfdisk(0);
	}

	private PfdiskContext pfdisk(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PfdiskContext _localctx = new PfdiskContext(_ctx, _parentState);
		PfdiskContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_pfdisk, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(80);
			parametrosfdisk();
			  
			}
			_ctx.stop = _input.LT(-1);
			setState(89);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PfdiskContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pfdisk);
					setState(83);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(84);
					parametrosfdisk();
					  
					}
					} 
				}
				setState(91);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
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

	public static class ParametrosfdiskContext extends ParserRuleContext {
		public Token CADENA;
		public ParametrosmkdiskContext parametrosmkdisk() {
			return getRuleContext(ParametrosmkdiskContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(Chems.TYPE, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public ParametrosfdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosfdisk; }
	}

	public final ParametrosfdiskContext parametrosfdisk() throws RecognitionException {
		ParametrosfdiskContext _localctx = new ParametrosfdiskContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_parametrosfdisk);
		try {
			setState(101);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SIZE:
			case FIT:
			case UNIT:
			case PATH:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				parametrosmkdisk();
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				match(TYPE);
				setState(94);
				match(IGUAL);
				setState(95);
				((ParametrosfdiskContext)_localctx).CADENA = match(CADENA);
				 ValType = (((ParametrosfdiskContext)_localctx).CADENA!=null?((ParametrosfdiskContext)_localctx).CADENA.getText():null) 
				}
				break;
			case NAME:
				enterOuterAlt(_localctx, 3);
				{
				setState(97);
				match(NAME);
				setState(98);
				match(IGUAL);
				setState(99);
				((ParametrosfdiskContext)_localctx).CADENA = match(CADENA);
				 ValName = (((ParametrosfdiskContext)_localctx).CADENA!=null?((ParametrosfdiskContext)_localctx).CADENA.getText():null) 
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

	public static class PmountContext extends ParserRuleContext {
		public ParametrosmountContext parametrosmount() {
			return getRuleContext(ParametrosmountContext.class,0);
		}
		public PmountContext pmount() {
			return getRuleContext(PmountContext.class,0);
		}
		public PmountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmount; }
	}

	public final PmountContext pmount() throws RecognitionException {
		return pmount(0);
	}

	private PmountContext pmount(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PmountContext _localctx = new PmountContext(_ctx, _parentState);
		PmountContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_pmount, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(104);
			parametrosmount();
			  
			}
			_ctx.stop = _input.LT(-1);
			setState(113);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmountContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmount);
					setState(107);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(108);
					parametrosmount();
					  
					}
					} 
				}
				setState(115);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public static class ParametrosmountContext extends ParserRuleContext {
		public Token RUTA;
		public Token CADENA;
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public ParametrosmountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmount; }
	}

	public final ParametrosmountContext parametrosmount() throws RecognitionException {
		ParametrosmountContext _localctx = new ParametrosmountContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_parametrosmount);
		try {
			setState(124);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PATH:
				enterOuterAlt(_localctx, 1);
				{
				setState(116);
				match(PATH);
				setState(117);
				match(IGUAL);
				setState(118);
				((ParametrosmountContext)_localctx).RUTA = match(RUTA);
				 ValPath = (((ParametrosmountContext)_localctx).RUTA!=null?((ParametrosmountContext)_localctx).RUTA.getText():null) 
				}
				break;
			case NAME:
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(NAME);
				setState(121);
				match(IGUAL);
				setState(122);
				((ParametrosmountContext)_localctx).CADENA = match(CADENA);
				 ValName = (((ParametrosmountContext)_localctx).CADENA!=null?((ParametrosmountContext)_localctx).CADENA.getText():null) 
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
			return pmkdisk_sempred((PmkdiskContext)_localctx, predIndex);
		case 5:
			return pfdisk_sempred((PfdiskContext)_localctx, predIndex);
		case 7:
			return pmount_sempred((PmountContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean pmkdisk_sempred(PmkdiskContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean pfdisk_sempred(PfdiskContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean pmount_sempred(PmountContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\60\u0081\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3"+
		"\2\3\2\3\2\5\2\31\n\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3#\n\3\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\7\4-\n\4\f\4\16\4\60\13\4\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\5\5J\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\7\7Z\n\7\f\7\16\7]\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bh\n"+
		"\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\tr\n\t\f\t\16\tu\13\t\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\5\n\177\n\n\3\n\2\5\6\f\20\13\2\4\6\b\n\f\16\20\22"+
		"\2\2\2\u0087\2\30\3\2\2\2\4\"\3\2\2\2\6$\3\2\2\2\bI\3\2\2\2\nK\3\2\2\2"+
		"\fQ\3\2\2\2\16g\3\2\2\2\20i\3\2\2\2\22~\3\2\2\2\24\31\5\4\3\2\25\26\7"+
		"%\2\2\26\31\b\2\1\2\27\31\7\2\2\3\30\24\3\2\2\2\30\25\3\2\2\2\30\27\3"+
		"\2\2\2\31\3\3\2\2\2\32\33\7\3\2\2\33#\5\6\4\2\34\35\7\4\2\2\35#\5\n\6"+
		"\2\36\37\7\5\2\2\37#\5\f\7\2 !\7\6\2\2!#\5\20\t\2\"\32\3\2\2\2\"\34\3"+
		"\2\2\2\"\36\3\2\2\2\" \3\2\2\2#\5\3\2\2\2$%\b\4\1\2%&\5\b\5\2&\'\b\4\1"+
		"\2\'.\3\2\2\2()\f\4\2\2)*\5\b\5\2*+\b\4\1\2+-\3\2\2\2,(\3\2\2\2-\60\3"+
		"\2\2\2.,\3\2\2\2./\3\2\2\2/\7\3\2\2\2\60.\3\2\2\2\61\62\7\22\2\2\62\63"+
		"\7.\2\2\63\64\7*\2\2\64J\b\5\1\2\65\66\7\23\2\2\66\67\7.\2\2\678\7\31"+
		"\2\28J\b\5\1\29:\7\23\2\2:;\7.\2\2;<\7\32\2\2<J\b\5\1\2=>\7\23\2\2>?\7"+
		".\2\2?@\7\33\2\2@J\b\5\1\2AB\7\24\2\2BC\7.\2\2CD\7+\2\2DJ\b\5\1\2EF\7"+
		"\25\2\2FG\7.\2\2GH\7-\2\2HJ\b\5\1\2I\61\3\2\2\2I\65\3\2\2\2I9\3\2\2\2"+
		"I=\3\2\2\2IA\3\2\2\2IE\3\2\2\2J\t\3\2\2\2KL\7\4\2\2LM\7\25\2\2MN\7.\2"+
		"\2NO\7-\2\2OP\b\6\1\2P\13\3\2\2\2QR\b\7\1\2RS\5\16\b\2ST\b\7\1\2T[\3\2"+
		"\2\2UV\f\4\2\2VW\5\16\b\2WX\b\7\1\2XZ\3\2\2\2YU\3\2\2\2Z]\3\2\2\2[Y\3"+
		"\2\2\2[\\\3\2\2\2\\\r\3\2\2\2][\3\2\2\2^h\5\b\5\2_`\7\27\2\2`a\7.\2\2"+
		"ab\7+\2\2bh\b\b\1\2cd\7\26\2\2de\7.\2\2ef\7+\2\2fh\b\b\1\2g^\3\2\2\2g"+
		"_\3\2\2\2gc\3\2\2\2h\17\3\2\2\2ij\b\t\1\2jk\5\22\n\2kl\b\t\1\2ls\3\2\2"+
		"\2mn\f\4\2\2no\5\22\n\2op\b\t\1\2pr\3\2\2\2qm\3\2\2\2ru\3\2\2\2sq\3\2"+
		"\2\2st\3\2\2\2t\21\3\2\2\2us\3\2\2\2vw\7\25\2\2wx\7.\2\2xy\7-\2\2y\177"+
		"\b\n\1\2z{\7\26\2\2{|\7.\2\2|}\7+\2\2}\177\b\n\1\2~v\3\2\2\2~z\3\2\2\2"+
		"\177\23\3\2\2\2\n\30\".I[gs~";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from /home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/MiaProyecto2/Chems.g4 by ANTLR 4.8

  import AD"MiaProyecto2/AdminDisks"
  var ValSize string = ""
  var ValFit  string = ""
  var ValUnit string = ""
  var ValPath string = ""
  var ValType string = ""
  var ValName string = ""
  var ValIdentificador string = ""
  var ValUsuario string = ""
  var ValPassword string = ""
  var ValGrupo string = ""
  var ValR bool = false
  var ValCont string = ""


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
		RULE_start = 0, RULE_adminDiscos = 1, RULE_adminArchivos = 2, RULE_adminUsuarios = 3, 
		RULE_adminCarpetas = 4, RULE_pmkfile = 5, RULE_parametrosmkfile = 6, RULE_pmkdisk = 7, 
		RULE_parametrosmkdisk = 8, RULE_prmdisk = 9, RULE_pmkgrp = 10, RULE_prmkgrp = 11, 
		RULE_prmkusr = 12, RULE_pfdisk = 13, RULE_parametrosfdisk = 14, RULE_pmount = 15, 
		RULE_parametrosmount = 16, RULE_pmkfs = 17, RULE_parametrosmkfs = 18, 
		RULE_plogin = 19, RULE_parametroslogin = 20, RULE_pmkuser = 21, RULE_parametrosmkuser = 22;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "adminDiscos", "adminArchivos", "adminUsuarios", "adminCarpetas", 
			"pmkfile", "parametrosmkfile", "pmkdisk", "parametrosmkdisk", "prmdisk", 
			"pmkgrp", "prmkgrp", "prmkusr", "pfdisk", "parametrosfdisk", "pmount", 
			"parametrosmount", "pmkfs", "parametrosmkfs", "plogin", "parametroslogin", 
			"pmkuser", "parametrosmkuser"
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
		public AdminArchivosContext adminArchivos() {
			return getRuleContext(AdminArchivosContext.class,0);
		}
		public AdminUsuariosContext adminUsuarios() {
			return getRuleContext(AdminUsuariosContext.class,0);
		}
		public AdminCarpetasContext adminCarpetas() {
			return getRuleContext(AdminCarpetasContext.class,0);
		}
		public TerminalNode PAUSA() { return getToken(Chems.PAUSA, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			setState(52);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EOF:
			case MKDISK:
			case RMDISK:
			case FDISK:
			case MOUNT:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				adminDiscos();
				}
				break;
			case MKFS:
				enterOuterAlt(_localctx, 2);
				{
				setState(47);
				adminArchivos();
				}
				break;
			case LOGIN:
			case LOGOUT:
			case MKGRP:
			case RMGRP:
			case MKUSER:
			case RMUSR:
				enterOuterAlt(_localctx, 3);
				{
				setState(48);
				adminUsuarios();
				}
				break;
			case MKFILE:
				enterOuterAlt(_localctx, 4);
				{
				setState(49);
				adminCarpetas();
				}
				break;
			case PAUSA:
				enterOuterAlt(_localctx, 5);
				{
				setState(50);
				match(PAUSA);
				 AD.Pausa() 
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
		public TerminalNode EOF() { return getToken(Chems.EOF, 0); }
		public AdminDiscosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminDiscos; }
	}

	public final AdminDiscosContext adminDiscos() throws RecognitionException {
		AdminDiscosContext _localctx = new AdminDiscosContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_adminDiscos);
		try {
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(54);
				match(MKDISK);
				setState(55);
				pmkdisk(0);
				 AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath); ValSize = ""; ValFit = ""; ValUnit = ""; ValPath = ""; 
				}
				break;
			case RMDISK:
				enterOuterAlt(_localctx, 2);
				{
				setState(58);
				match(RMDISK);
				setState(59);
				prmdisk();
				 AD.Rmdisk(ValPath); ValPath = ""; 
				}
				break;
			case FDISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(62);
				match(FDISK);
				setState(63);
				pfdisk(0);
				 AD.Fdisk(ValSize, ValFit, ValUnit, ValPath, ValType, ValName); ValSize = ""; ValFit = ""; ValUnit = ""; ValPath = ""; ValType = ""; ValName = "";
				}
				break;
			case MOUNT:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				match(MOUNT);
				setState(67);
				pmount(0);
				 AD.Mount(ValPath, ValName); ValPath = ""; ValName = "" 
				}
				break;
			case EOF:
				enterOuterAlt(_localctx, 5);
				{
				setState(70);
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

	public static class AdminArchivosContext extends ParserRuleContext {
		public TerminalNode MKFS() { return getToken(Chems.MKFS, 0); }
		public PmkfsContext pmkfs() {
			return getRuleContext(PmkfsContext.class,0);
		}
		public AdminArchivosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminArchivos; }
	}

	public final AdminArchivosContext adminArchivos() throws RecognitionException {
		AdminArchivosContext _localctx = new AdminArchivosContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_adminArchivos);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(MKFS);
			setState(74);
			pmkfs(0);
			 AD.Mkfs(ValIdentificador, ValType); ValIdentificador = ""; ValType = ""; 
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

	public static class AdminUsuariosContext extends ParserRuleContext {
		public TerminalNode LOGIN() { return getToken(Chems.LOGIN, 0); }
		public PloginContext plogin() {
			return getRuleContext(PloginContext.class,0);
		}
		public TerminalNode LOGOUT() { return getToken(Chems.LOGOUT, 0); }
		public TerminalNode MKGRP() { return getToken(Chems.MKGRP, 0); }
		public PmkgrpContext pmkgrp() {
			return getRuleContext(PmkgrpContext.class,0);
		}
		public TerminalNode RMGRP() { return getToken(Chems.RMGRP, 0); }
		public PrmkgrpContext prmkgrp() {
			return getRuleContext(PrmkgrpContext.class,0);
		}
		public TerminalNode MKUSER() { return getToken(Chems.MKUSER, 0); }
		public PmkuserContext pmkuser() {
			return getRuleContext(PmkuserContext.class,0);
		}
		public TerminalNode RMUSR() { return getToken(Chems.RMUSR, 0); }
		public PrmkusrContext prmkusr() {
			return getRuleContext(PrmkusrContext.class,0);
		}
		public AdminUsuariosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminUsuarios; }
	}

	public final AdminUsuariosContext adminUsuarios() throws RecognitionException {
		AdminUsuariosContext _localctx = new AdminUsuariosContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_adminUsuarios);
		try {
			setState(99);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LOGIN:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				match(LOGIN);
				setState(78);
				plogin(0);
				 AD.Login(ValUsuario, ValPassword, ValIdentificador); ValUsuario = "" ; ValPassword = "" ; ValIdentificador = ""; 
				}
				break;
			case LOGOUT:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				match(LOGOUT);
				 AD.LogOut() 
				}
				break;
			case MKGRP:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				match(MKGRP);
				setState(84);
				pmkgrp();
				 AD.Mkgrp(ValName); ValName = ""; 
				}
				break;
			case RMGRP:
				enterOuterAlt(_localctx, 4);
				{
				setState(87);
				match(RMGRP);
				setState(88);
				prmkgrp();
				 AD.RMkgrp(ValName); ValName = ""; 
				}
				break;
			case MKUSER:
				enterOuterAlt(_localctx, 5);
				{
				setState(91);
				match(MKUSER);
				setState(92);
				pmkuser();
				 AD.Mkusr(ValUsuario, ValPassword, ValGrupo); ValUsuario = "" ; ValPassword = "" ; ValGrupo = ""; 
				}
				break;
			case RMUSR:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				match(RMUSR);
				setState(96);
				prmkusr();
				 AD.RMkusr(ValName); ValName = ""; 
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

	public static class AdminCarpetasContext extends ParserRuleContext {
		public TerminalNode MKFILE() { return getToken(Chems.MKFILE, 0); }
		public PmkfileContext pmkfile() {
			return getRuleContext(PmkfileContext.class,0);
		}
		public AdminCarpetasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_adminCarpetas; }
	}

	public final AdminCarpetasContext adminCarpetas() throws RecognitionException {
		AdminCarpetasContext _localctx = new AdminCarpetasContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_adminCarpetas);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(101);
			match(MKFILE);
			setState(102);
			pmkfile(0);
			 AD.MkFile(ValPath, ValR, ValSize, ValCont); ValPath = ""; ValR = false; ValSize = ""; ValCont = "" 
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

	public static class PmkfileContext extends ParserRuleContext {
		public ParametrosmkfileContext parametrosmkfile() {
			return getRuleContext(ParametrosmkfileContext.class,0);
		}
		public PmkfileContext pmkfile() {
			return getRuleContext(PmkfileContext.class,0);
		}
		public PmkfileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmkfile; }
	}

	public final PmkfileContext pmkfile() throws RecognitionException {
		return pmkfile(0);
	}

	private PmkfileContext pmkfile(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PmkfileContext _localctx = new PmkfileContext(_ctx, _parentState);
		PmkfileContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_pmkfile, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(106);
			parametrosmkfile();
			}
			_ctx.stop = _input.LT(-1);
			setState(112);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmkfileContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmkfile);
					setState(108);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(109);
					parametrosmkfile();
					}
					} 
				}
				setState(114);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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

	public static class ParametrosmkfileContext extends ParserRuleContext {
		public Token ENTERO;
		public Token CADENA;
		public Token RUTA;
		public Token LETRA;
		public TerminalNode SIZE() { return getToken(Chems.SIZE, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode ENTERO() { return getToken(Chems.ENTERO, 0); }
		public TerminalNode CONT() { return getToken(Chems.CONT, 0); }
		public TerminalNode BF() { return getToken(Chems.BF, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public TerminalNode UNIT() { return getToken(Chems.UNIT, 0); }
		public TerminalNode LETRA() { return getToken(Chems.LETRA, 0); }
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode GR() { return getToken(Chems.GR, 0); }
		public ParametrosmkfileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkfile; }
	}

	public final ParametrosmkfileContext parametrosmkfile() throws RecognitionException {
		ParametrosmkfileContext _localctx = new ParametrosmkfileContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_parametrosmkfile);
		try {
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				match(SIZE);
				setState(116);
				match(IGUAL);
				setState(117);
				((ParametrosmkfileContext)_localctx).ENTERO = match(ENTERO);
				 
				    ValSize = (((ParametrosmkfileContext)_localctx).ENTERO!=null?((ParametrosmkfileContext)_localctx).ENTERO.getText():null)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(119);
				match(CONT);
				setState(120);
				match(IGUAL);
				setState(121);
				match(BF);
				 ValCont = "" 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(123);
				match(CONT);
				setState(124);
				match(IGUAL);
				setState(125);
				((ParametrosmkfileContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkfileContext)_localctx).CADENA!=null?((ParametrosmkfileContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkfileContext)_localctx).CADENA!=null?((ParametrosmkfileContext)_localctx).CADENA.getText():null))-1]
				                        ValCont = str 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(127);
				match(CONT);
				setState(128);
				match(IGUAL);
				setState(129);
				((ParametrosmkfileContext)_localctx).RUTA = match(RUTA);
				 ValCont = (((ParametrosmkfileContext)_localctx).RUTA!=null?((ParametrosmkfileContext)_localctx).RUTA.getText():null) 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(131);
				match(UNIT);
				setState(132);
				match(IGUAL);
				setState(133);
				((ParametrosmkfileContext)_localctx).LETRA = match(LETRA);
				 ValUnit = (((ParametrosmkfileContext)_localctx).LETRA!=null?((ParametrosmkfileContext)_localctx).LETRA.getText():null) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(135);
				match(PATH);
				setState(136);
				match(IGUAL);
				setState(137);
				((ParametrosmkfileContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkfileContext)_localctx).CADENA!=null?((ParametrosmkfileContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkfileContext)_localctx).CADENA!=null?((ParametrosmkfileContext)_localctx).CADENA.getText():null))-1]
				                        ValPath = str 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(139);
				match(PATH);
				setState(140);
				match(IGUAL);
				setState(141);
				((ParametrosmkfileContext)_localctx).RUTA = match(RUTA);
				 ValPath = (((ParametrosmkfileContext)_localctx).RUTA!=null?((ParametrosmkfileContext)_localctx).RUTA.getText():null) 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(143);
				match(GR);
				 ValR = true 
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
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_pmkdisk, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(148);
			parametrosmkdisk();
			}
			_ctx.stop = _input.LT(-1);
			setState(154);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmkdiskContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmkdisk);
					setState(150);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(151);
					parametrosmkdisk();
					}
					} 
				}
				setState(156);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		public Token LETRA;
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
		public TerminalNode LETRA() { return getToken(Chems.LETRA, 0); }
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public ParametrosmkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkdisk; }
	}

	public final ParametrosmkdiskContext parametrosmkdisk() throws RecognitionException {
		ParametrosmkdiskContext _localctx = new ParametrosmkdiskContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_parametrosmkdisk);
		try {
			setState(185);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(157);
				match(SIZE);
				setState(158);
				match(IGUAL);
				setState(159);
				((ParametrosmkdiskContext)_localctx).ENTERO = match(ENTERO);
				 ValSize = (((ParametrosmkdiskContext)_localctx).ENTERO!=null?((ParametrosmkdiskContext)_localctx).ENTERO.getText():null) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				match(FIT);
				setState(162);
				match(IGUAL);
				setState(163);
				match(BF);
				 ValFit = "b" 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(165);
				match(FIT);
				setState(166);
				match(IGUAL);
				setState(167);
				match(FF);
				 ValFit = "f" 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(169);
				match(FIT);
				setState(170);
				match(IGUAL);
				setState(171);
				match(WF);
				 ValFit = "w" 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(173);
				match(UNIT);
				setState(174);
				match(IGUAL);
				setState(175);
				((ParametrosmkdiskContext)_localctx).LETRA = match(LETRA);
				 ValUnit = (((ParametrosmkdiskContext)_localctx).LETRA!=null?((ParametrosmkdiskContext)_localctx).LETRA.getText():null) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(177);
				match(PATH);
				setState(178);
				match(IGUAL);
				setState(179);
				((ParametrosmkdiskContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkdiskContext)_localctx).CADENA!=null?((ParametrosmkdiskContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkdiskContext)_localctx).CADENA!=null?((ParametrosmkdiskContext)_localctx).CADENA.getText():null))-1]
				                        ValPath = str 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(181);
				match(PATH);
				setState(182);
				match(IGUAL);
				setState(183);
				((ParametrosmkdiskContext)_localctx).RUTA = match(RUTA);
				 ValPath = (((ParametrosmkdiskContext)_localctx).RUTA!=null?((ParametrosmkdiskContext)_localctx).RUTA.getText():null) 
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
		public Token CADENA;
		public Token RUTA;
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public PrmdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prmdisk; }
	}

	public final PrmdiskContext prmdisk() throws RecognitionException {
		PrmdiskContext _localctx = new PrmdiskContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_prmdisk);
		try {
			setState(195);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(187);
				match(PATH);
				setState(188);
				match(IGUAL);
				setState(189);
				((PrmdiskContext)_localctx).CADENA = match(CADENA);
				  str:= (((PrmdiskContext)_localctx).CADENA!=null?((PrmdiskContext)_localctx).CADENA.getText():null)[1:len((((PrmdiskContext)_localctx).CADENA!=null?((PrmdiskContext)_localctx).CADENA.getText():null))-1]
				                                ValPath = str 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(191);
				match(PATH);
				setState(192);
				match(IGUAL);
				setState(193);
				((PrmdiskContext)_localctx).RUTA = match(RUTA);
				 ValPath = (((PrmdiskContext)_localctx).RUTA!=null?((PrmdiskContext)_localctx).RUTA.getText():null) 
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

	public static class PmkgrpContext extends ParserRuleContext {
		public Token CADENA;
		public Token IDENTIFICADOR;
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public PmkgrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmkgrp; }
	}

	public final PmkgrpContext pmkgrp() throws RecognitionException {
		PmkgrpContext _localctx = new PmkgrpContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_pmkgrp);
		try {
			setState(205);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(197);
				match(NAME);
				setState(198);
				match(IGUAL);
				setState(199);
				((PmkgrpContext)_localctx).CADENA = match(CADENA);
				  str:= (((PmkgrpContext)_localctx).CADENA!=null?((PmkgrpContext)_localctx).CADENA.getText():null)[1:len((((PmkgrpContext)_localctx).CADENA!=null?((PmkgrpContext)_localctx).CADENA.getText():null))-1]
				                                ValName = str 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(201);
				match(NAME);
				setState(202);
				match(IGUAL);
				setState(203);
				((PmkgrpContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValName = (((PmkgrpContext)_localctx).IDENTIFICADOR!=null?((PmkgrpContext)_localctx).IDENTIFICADOR.getText():null) 
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

	public static class PrmkgrpContext extends ParserRuleContext {
		public Token CADENA;
		public Token IDENTIFICADOR;
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public PrmkgrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prmkgrp; }
	}

	public final PrmkgrpContext prmkgrp() throws RecognitionException {
		PrmkgrpContext _localctx = new PrmkgrpContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_prmkgrp);
		try {
			setState(215);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(207);
				match(NAME);
				setState(208);
				match(IGUAL);
				setState(209);
				((PrmkgrpContext)_localctx).CADENA = match(CADENA);
				  str:= (((PrmkgrpContext)_localctx).CADENA!=null?((PrmkgrpContext)_localctx).CADENA.getText():null)[1:len((((PrmkgrpContext)_localctx).CADENA!=null?((PrmkgrpContext)_localctx).CADENA.getText():null))-1]
				                                ValName = str 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(211);
				match(NAME);
				setState(212);
				match(IGUAL);
				setState(213);
				((PrmkgrpContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValName = (((PrmkgrpContext)_localctx).IDENTIFICADOR!=null?((PrmkgrpContext)_localctx).IDENTIFICADOR.getText():null) 
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

	public static class PrmkusrContext extends ParserRuleContext {
		public Token CADENA;
		public Token IDENTIFICADOR;
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public PrmkusrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prmkusr; }
	}

	public final PrmkusrContext prmkusr() throws RecognitionException {
		PrmkusrContext _localctx = new PrmkusrContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_prmkusr);
		try {
			setState(225);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				match(NAME);
				setState(218);
				match(IGUAL);
				setState(219);
				((PrmkusrContext)_localctx).CADENA = match(CADENA);
				  str:= (((PrmkusrContext)_localctx).CADENA!=null?((PrmkusrContext)_localctx).CADENA.getText():null)[1:len((((PrmkusrContext)_localctx).CADENA!=null?((PrmkusrContext)_localctx).CADENA.getText():null))-1]
				                                ValName = str 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(221);
				match(NAME);
				setState(222);
				match(IGUAL);
				setState(223);
				((PrmkusrContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValName = (((PrmkusrContext)_localctx).IDENTIFICADOR!=null?((PrmkusrContext)_localctx).IDENTIFICADOR.getText():null) 
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
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_pfdisk, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(228);
			parametrosfdisk();
			}
			_ctx.stop = _input.LT(-1);
			setState(234);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PfdiskContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pfdisk);
					setState(230);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(231);
					parametrosfdisk();
					}
					} 
				}
				setState(236);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		public Token LETRA;
		public Token CADENA;
		public Token IDENTIFICADOR;
		public ParametrosmkdiskContext parametrosmkdisk() {
			return getRuleContext(ParametrosmkdiskContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(Chems.TYPE, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode LETRA() { return getToken(Chems.LETRA, 0); }
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public ParametrosfdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosfdisk; }
	}

	public final ParametrosfdiskContext parametrosfdisk() throws RecognitionException {
		ParametrosfdiskContext _localctx = new ParametrosfdiskContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_parametrosfdisk);
		try {
			setState(250);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(237);
				parametrosmkdisk();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(238);
				match(TYPE);
				setState(239);
				match(IGUAL);
				setState(240);
				((ParametrosfdiskContext)_localctx).LETRA = match(LETRA);
				 ValType = (((ParametrosfdiskContext)_localctx).LETRA!=null?((ParametrosfdiskContext)_localctx).LETRA.getText():null) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(242);
				match(NAME);
				setState(243);
				match(IGUAL);
				setState(244);
				((ParametrosfdiskContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosfdiskContext)_localctx).CADENA!=null?((ParametrosfdiskContext)_localctx).CADENA.getText():null)[1:len((((ParametrosfdiskContext)_localctx).CADENA!=null?((ParametrosfdiskContext)_localctx).CADENA.getText():null))-1]
				                      ValName = str
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(246);
				match(NAME);
				setState(247);
				match(IGUAL);
				setState(248);
				((ParametrosfdiskContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValName = (((ParametrosfdiskContext)_localctx).IDENTIFICADOR!=null?((ParametrosfdiskContext)_localctx).IDENTIFICADOR.getText():null)
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_pmount, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(253);
			parametrosmount();
			}
			_ctx.stop = _input.LT(-1);
			setState(259);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmountContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmount);
					setState(255);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(256);
					parametrosmount();
					}
					} 
				}
				setState(261);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		public Token CADENA;
		public Token RUTA;
		public Token IDENTIFICADOR;
		public TerminalNode PATH() { return getToken(Chems.PATH, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode RUTA() { return getToken(Chems.RUTA, 0); }
		public TerminalNode NAME() { return getToken(Chems.NAME, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public ParametrosmountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmount; }
	}

	public final ParametrosmountContext parametrosmount() throws RecognitionException {
		ParametrosmountContext _localctx = new ParametrosmountContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_parametrosmount);
		try {
			setState(274);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(262);
				match(PATH);
				setState(263);
				match(IGUAL);
				setState(264);
				((ParametrosmountContext)_localctx).CADENA = match(CADENA);
				  str:= (((ParametrosmountContext)_localctx).CADENA!=null?((ParametrosmountContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmountContext)_localctx).CADENA!=null?((ParametrosmountContext)_localctx).CADENA.getText():null))-1]
				                                ValPath = str 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(266);
				match(PATH);
				setState(267);
				match(IGUAL);
				setState(268);
				((ParametrosmountContext)_localctx).RUTA = match(RUTA);
				 ValPath = (((ParametrosmountContext)_localctx).RUTA!=null?((ParametrosmountContext)_localctx).RUTA.getText():null) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(270);
				match(NAME);
				setState(271);
				match(IGUAL);
				setState(272);
				((ParametrosmountContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValName = (((ParametrosmountContext)_localctx).IDENTIFICADOR!=null?((ParametrosmountContext)_localctx).IDENTIFICADOR.getText():null)
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

	public static class PmkfsContext extends ParserRuleContext {
		public ParametrosmkfsContext parametrosmkfs() {
			return getRuleContext(ParametrosmkfsContext.class,0);
		}
		public PmkfsContext pmkfs() {
			return getRuleContext(PmkfsContext.class,0);
		}
		public PmkfsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmkfs; }
	}

	public final PmkfsContext pmkfs() throws RecognitionException {
		return pmkfs(0);
	}

	private PmkfsContext pmkfs(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PmkfsContext _localctx = new PmkfsContext(_ctx, _parentState);
		PmkfsContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_pmkfs, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(277);
			parametrosmkfs();
			}
			_ctx.stop = _input.LT(-1);
			setState(283);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PmkfsContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pmkfs);
					setState(279);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(280);
					parametrosmkfs();
					}
					} 
				}
				setState(285);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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

	public static class ParametrosmkfsContext extends ParserRuleContext {
		public Token IDENTIFICADOR;
		public Token LETRA;
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public TerminalNode TYPE() { return getToken(Chems.TYPE, 0); }
		public TerminalNode LETRA() { return getToken(Chems.LETRA, 0); }
		public ParametrosmkfsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkfs; }
	}

	public final ParametrosmkfsContext parametrosmkfs() throws RecognitionException {
		ParametrosmkfsContext _localctx = new ParametrosmkfsContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_parametrosmkfs);
		try {
			setState(294);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(286);
				match(ID);
				setState(287);
				match(IGUAL);
				setState(288);
				((ParametrosmkfsContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				  ValIdentificador = (((ParametrosmkfsContext)_localctx).IDENTIFICADOR!=null?((ParametrosmkfsContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 2);
				{
				setState(290);
				match(TYPE);
				setState(291);
				match(IGUAL);
				setState(292);
				((ParametrosmkfsContext)_localctx).LETRA = match(LETRA);
				 ValType = (((ParametrosmkfsContext)_localctx).LETRA!=null?((ParametrosmkfsContext)_localctx).LETRA.getText():null) 
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

	public static class PloginContext extends ParserRuleContext {
		public ParametrosloginContext parametroslogin() {
			return getRuleContext(ParametrosloginContext.class,0);
		}
		public PloginContext plogin() {
			return getRuleContext(PloginContext.class,0);
		}
		public PloginContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_plogin; }
	}

	public final PloginContext plogin() throws RecognitionException {
		return plogin(0);
	}

	private PloginContext plogin(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PloginContext _localctx = new PloginContext(_ctx, _parentState);
		PloginContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_plogin, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(297);
			parametroslogin();
			}
			_ctx.stop = _input.LT(-1);
			setState(303);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PloginContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_plogin);
					setState(299);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(300);
					parametroslogin();
					}
					} 
				}
				setState(305);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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

	public static class ParametrosloginContext extends ParserRuleContext {
		public Token IDENTIFICADOR;
		public Token CADENA;
		public TerminalNode USR() { return getToken(Chems.USR, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode PWD() { return getToken(Chems.PWD, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public ParametrosloginContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametroslogin; }
	}

	public final ParametrosloginContext parametroslogin() throws RecognitionException {
		ParametrosloginContext _localctx = new ParametrosloginContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_parametroslogin);
		try {
			setState(326);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(306);
				match(USR);
				setState(307);
				match(IGUAL);
				setState(308);
				((ParametrosloginContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValUsuario = (((ParametrosloginContext)_localctx).IDENTIFICADOR!=null?((ParametrosloginContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(310);
				match(USR);
				setState(311);
				match(IGUAL);
				setState(312);
				((ParametrosloginContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosloginContext)_localctx).CADENA!=null?((ParametrosloginContext)_localctx).CADENA.getText():null)[1:len((((ParametrosloginContext)_localctx).CADENA!=null?((ParametrosloginContext)_localctx).CADENA.getText():null))-1]
				                      ValUsuario = str 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(314);
				match(PWD);
				setState(315);
				match(IGUAL);
				setState(316);
				((ParametrosloginContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValPassword = (((ParametrosloginContext)_localctx).IDENTIFICADOR!=null?((ParametrosloginContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(318);
				match(PWD);
				setState(319);
				match(IGUAL);
				setState(320);
				((ParametrosloginContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosloginContext)_localctx).CADENA!=null?((ParametrosloginContext)_localctx).CADENA.getText():null)[1:len((((ParametrosloginContext)_localctx).CADENA!=null?((ParametrosloginContext)_localctx).CADENA.getText():null))-1]
				                      ValPassword = str 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(322);
				match(ID);
				setState(323);
				match(IGUAL);
				setState(324);
				((ParametrosloginContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				  ValIdentificador = (((ParametrosloginContext)_localctx).IDENTIFICADOR!=null?((ParametrosloginContext)_localctx).IDENTIFICADOR.getText():null) 
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

	public static class PmkuserContext extends ParserRuleContext {
		public ParametrosmkuserContext parametrosmkuser() {
			return getRuleContext(ParametrosmkuserContext.class,0);
		}
		public PloginContext plogin() {
			return getRuleContext(PloginContext.class,0);
		}
		public PmkuserContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pmkuser; }
	}

	public final PmkuserContext pmkuser() throws RecognitionException {
		PmkuserContext _localctx = new PmkuserContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_pmkuser);
		try {
			setState(332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(328);
				parametrosmkuser();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(329);
				plogin(0);
				setState(330);
				parametrosmkuser();
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

	public static class ParametrosmkuserContext extends ParserRuleContext {
		public Token IDENTIFICADOR;
		public Token CADENA;
		public TerminalNode USR() { return getToken(Chems.USR, 0); }
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public TerminalNode IDENTIFICADOR() { return getToken(Chems.IDENTIFICADOR, 0); }
		public TerminalNode CADENA() { return getToken(Chems.CADENA, 0); }
		public TerminalNode PWD1() { return getToken(Chems.PWD1, 0); }
		public TerminalNode GRP() { return getToken(Chems.GRP, 0); }
		public ParametrosmkuserContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosmkuser; }
	}

	public final ParametrosmkuserContext parametrosmkuser() throws RecognitionException {
		ParametrosmkuserContext _localctx = new ParametrosmkuserContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_parametrosmkuser);
		try {
			setState(358);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(334);
				match(USR);
				setState(335);
				match(IGUAL);
				setState(336);
				((ParametrosmkuserContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValUsuario = (((ParametrosmkuserContext)_localctx).IDENTIFICADOR!=null?((ParametrosmkuserContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(338);
				match(USR);
				setState(339);
				match(IGUAL);
				setState(340);
				((ParametrosmkuserContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null))-1]
				                      ValUsuario = str 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(342);
				match(PWD1);
				setState(343);
				match(IGUAL);
				setState(344);
				((ParametrosmkuserContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				 ValPassword = (((ParametrosmkuserContext)_localctx).IDENTIFICADOR!=null?((ParametrosmkuserContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(346);
				match(PWD1);
				setState(347);
				match(IGUAL);
				setState(348);
				((ParametrosmkuserContext)_localctx).CADENA = match(CADENA);
				 str:= (((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null))-1]
				                      ValPassword = str 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(350);
				match(GRP);
				setState(351);
				match(IGUAL);
				setState(352);
				((ParametrosmkuserContext)_localctx).IDENTIFICADOR = match(IDENTIFICADOR);
				  ValGrupo = (((ParametrosmkuserContext)_localctx).IDENTIFICADOR!=null?((ParametrosmkuserContext)_localctx).IDENTIFICADOR.getText():null) 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(354);
				match(GRP);
				setState(355);
				match(IGUAL);
				setState(356);
				((ParametrosmkuserContext)_localctx).CADENA = match(CADENA);
				  str:= (((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null)[1:len((((ParametrosmkuserContext)_localctx).CADENA!=null?((ParametrosmkuserContext)_localctx).CADENA.getText():null))-1]
				                                ValGrupo = str 
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 5:
			return pmkfile_sempred((PmkfileContext)_localctx, predIndex);
		case 7:
			return pmkdisk_sempred((PmkdiskContext)_localctx, predIndex);
		case 13:
			return pfdisk_sempred((PfdiskContext)_localctx, predIndex);
		case 15:
			return pmount_sempred((PmountContext)_localctx, predIndex);
		case 17:
			return pmkfs_sempred((PmkfsContext)_localctx, predIndex);
		case 19:
			return plogin_sempred((PloginContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean pmkfile_sempred(PmkfileContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pmkdisk_sempred(PmkdiskContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pfdisk_sempred(PfdiskContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pmount_sempred(PmountContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pmkfs_sempred(PmkfsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean plogin_sempred(PloginContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\60\u016b\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\3\2\3\2\3"+
		"\2\3\2\3\2\3\2\5\2\67\n\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\5\3J\n\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5"+
		"\5f\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\7\7q\n\7\f\7\16\7t\13\7\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b\u0094\n\b\3\t\3\t"+
		"\3\t\3\t\3\t\7\t\u009b\n\t\f\t\16\t\u009e\13\t\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\5\n\u00bc\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\5\13\u00c6\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00d0\n\f\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00da\n\r\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\5\16\u00e4\n\16\3\17\3\17\3\17\3\17\3\17\7\17\u00eb\n\17"+
		"\f\17\16\17\u00ee\13\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\5\20\u00fd\n\20\3\21\3\21\3\21\3\21\3\21\7\21\u0104"+
		"\n\21\f\21\16\21\u0107\13\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\5\22\u0115\n\22\3\23\3\23\3\23\3\23\3\23\7\23\u011c"+
		"\n\23\f\23\16\23\u011f\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5"+
		"\24\u0129\n\24\3\25\3\25\3\25\3\25\3\25\7\25\u0130\n\25\f\25\16\25\u0133"+
		"\13\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u0149\n\26\3\27\3\27\3\27\3\27"+
		"\5\27\u014f\n\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\5\30"+
		"\u0169\n\30\3\30\2\b\f\20\34 $(\31\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,.\2\2\2\u0187\2\66\3\2\2\2\4I\3\2\2\2\6K\3\2\2\2\be\3\2\2\2"+
		"\ng\3\2\2\2\fk\3\2\2\2\16\u0093\3\2\2\2\20\u0095\3\2\2\2\22\u00bb\3\2"+
		"\2\2\24\u00c5\3\2\2\2\26\u00cf\3\2\2\2\30\u00d9\3\2\2\2\32\u00e3\3\2\2"+
		"\2\34\u00e5\3\2\2\2\36\u00fc\3\2\2\2 \u00fe\3\2\2\2\"\u0114\3\2\2\2$\u0116"+
		"\3\2\2\2&\u0128\3\2\2\2(\u012a\3\2\2\2*\u0148\3\2\2\2,\u014e\3\2\2\2."+
		"\u0168\3\2\2\2\60\67\5\4\3\2\61\67\5\6\4\2\62\67\5\b\5\2\63\67\5\n\6\2"+
		"\64\65\7%\2\2\65\67\b\2\1\2\66\60\3\2\2\2\66\61\3\2\2\2\66\62\3\2\2\2"+
		"\66\63\3\2\2\2\66\64\3\2\2\2\67\3\3\2\2\289\7\3\2\29:\5\20\t\2:;\b\3\1"+
		"\2;J\3\2\2\2<=\7\4\2\2=>\5\24\13\2>?\b\3\1\2?J\3\2\2\2@A\7\5\2\2AB\5\34"+
		"\17\2BC\b\3\1\2CJ\3\2\2\2DE\7\6\2\2EF\5 \21\2FG\b\3\1\2GJ\3\2\2\2HJ\7"+
		"\2\2\3I8\3\2\2\2I<\3\2\2\2I@\3\2\2\2ID\3\2\2\2IH\3\2\2\2J\5\3\2\2\2KL"+
		"\7\7\2\2LM\5$\23\2MN\b\4\1\2N\7\3\2\2\2OP\7\b\2\2PQ\5(\25\2QR\b\5\1\2"+
		"Rf\3\2\2\2ST\7\t\2\2Tf\b\5\1\2UV\7\n\2\2VW\5\26\f\2WX\b\5\1\2Xf\3\2\2"+
		"\2YZ\7\13\2\2Z[\5\30\r\2[\\\b\5\1\2\\f\3\2\2\2]^\7\f\2\2^_\5,\27\2_`\b"+
		"\5\1\2`f\3\2\2\2ab\7\r\2\2bc\5\32\16\2cd\b\5\1\2df\3\2\2\2eO\3\2\2\2e"+
		"S\3\2\2\2eU\3\2\2\2eY\3\2\2\2e]\3\2\2\2ea\3\2\2\2f\t\3\2\2\2gh\7\16\2"+
		"\2hi\5\f\7\2ij\b\6\1\2j\13\3\2\2\2kl\b\7\1\2lm\5\16\b\2mr\3\2\2\2no\f"+
		"\3\2\2oq\5\16\b\2pn\3\2\2\2qt\3\2\2\2rp\3\2\2\2rs\3\2\2\2s\r\3\2\2\2t"+
		"r\3\2\2\2uv\7\22\2\2vw\7.\2\2wx\7*\2\2x\u0094\b\b\1\2yz\7#\2\2z{\7.\2"+
		"\2{|\7\31\2\2|\u0094\b\b\1\2}~\7#\2\2~\177\7.\2\2\177\u0080\7+\2\2\u0080"+
		"\u0094\b\b\1\2\u0081\u0082\7#\2\2\u0082\u0083\7.\2\2\u0083\u0084\7-\2"+
		"\2\u0084\u0094\b\b\1\2\u0085\u0086\7\24\2\2\u0086\u0087\7.\2\2\u0087\u0088"+
		"\7)\2\2\u0088\u0094\b\b\1\2\u0089\u008a\7\25\2\2\u008a\u008b\7.\2\2\u008b"+
		"\u008c\7+\2\2\u008c\u0094\b\b\1\2\u008d\u008e\7\25\2\2\u008e\u008f\7."+
		"\2\2\u008f\u0090\7-\2\2\u0090\u0094\b\b\1\2\u0091\u0092\7\"\2\2\u0092"+
		"\u0094\b\b\1\2\u0093u\3\2\2\2\u0093y\3\2\2\2\u0093}\3\2\2\2\u0093\u0081"+
		"\3\2\2\2\u0093\u0085\3\2\2\2\u0093\u0089\3\2\2\2\u0093\u008d\3\2\2\2\u0093"+
		"\u0091\3\2\2\2\u0094\17\3\2\2\2\u0095\u0096\b\t\1\2\u0096\u0097\5\22\n"+
		"\2\u0097\u009c\3\2\2\2\u0098\u0099\f\3\2\2\u0099\u009b\5\22\n\2\u009a"+
		"\u0098\3\2\2\2\u009b\u009e\3\2\2\2\u009c\u009a\3\2\2\2\u009c\u009d\3\2"+
		"\2\2\u009d\21\3\2\2\2\u009e\u009c\3\2\2\2\u009f\u00a0\7\22\2\2\u00a0\u00a1"+
		"\7.\2\2\u00a1\u00a2\7*\2\2\u00a2\u00bc\b\n\1\2\u00a3\u00a4\7\23\2\2\u00a4"+
		"\u00a5\7.\2\2\u00a5\u00a6\7\31\2\2\u00a6\u00bc\b\n\1\2\u00a7\u00a8\7\23"+
		"\2\2\u00a8\u00a9\7.\2\2\u00a9\u00aa\7\32\2\2\u00aa\u00bc\b\n\1\2\u00ab"+
		"\u00ac\7\23\2\2\u00ac\u00ad\7.\2\2\u00ad\u00ae\7\33\2\2\u00ae\u00bc\b"+
		"\n\1\2\u00af\u00b0\7\24\2\2\u00b0\u00b1\7.\2\2\u00b1\u00b2\7)\2\2\u00b2"+
		"\u00bc\b\n\1\2\u00b3\u00b4\7\25\2\2\u00b4\u00b5\7.\2\2\u00b5\u00b6\7+"+
		"\2\2\u00b6\u00bc\b\n\1\2\u00b7\u00b8\7\25\2\2\u00b8\u00b9\7.\2\2\u00b9"+
		"\u00ba\7-\2\2\u00ba\u00bc\b\n\1\2\u00bb\u009f\3\2\2\2\u00bb\u00a3\3\2"+
		"\2\2\u00bb\u00a7\3\2\2\2\u00bb\u00ab\3\2\2\2\u00bb\u00af\3\2\2\2\u00bb"+
		"\u00b3\3\2\2\2\u00bb\u00b7\3\2\2\2\u00bc\23\3\2\2\2\u00bd\u00be\7\25\2"+
		"\2\u00be\u00bf\7.\2\2\u00bf\u00c0\7+\2\2\u00c0\u00c6\b\13\1\2\u00c1\u00c2"+
		"\7\25\2\2\u00c2\u00c3\7.\2\2\u00c3\u00c4\7-\2\2\u00c4\u00c6\b\13\1\2\u00c5"+
		"\u00bd\3\2\2\2\u00c5\u00c1\3\2\2\2\u00c6\25\3\2\2\2\u00c7\u00c8\7\26\2"+
		"\2\u00c8\u00c9\7.\2\2\u00c9\u00ca\7+\2\2\u00ca\u00d0\b\f\1\2\u00cb\u00cc"+
		"\7\26\2\2\u00cc\u00cd\7.\2\2\u00cd\u00ce\7,\2\2\u00ce\u00d0\b\f\1\2\u00cf"+
		"\u00c7\3\2\2\2\u00cf\u00cb\3\2\2\2\u00d0\27\3\2\2\2\u00d1\u00d2\7\26\2"+
		"\2\u00d2\u00d3\7.\2\2\u00d3\u00d4\7+\2\2\u00d4\u00da\b\r\1\2\u00d5\u00d6"+
		"\7\26\2\2\u00d6\u00d7\7.\2\2\u00d7\u00d8\7,\2\2\u00d8\u00da\b\r\1\2\u00d9"+
		"\u00d1\3\2\2\2\u00d9\u00d5\3\2\2\2\u00da\31\3\2\2\2\u00db\u00dc\7\26\2"+
		"\2\u00dc\u00dd\7.\2\2\u00dd\u00de\7+\2\2\u00de\u00e4\b\16\1\2\u00df\u00e0"+
		"\7\26\2\2\u00e0\u00e1\7.\2\2\u00e1\u00e2\7,\2\2\u00e2\u00e4\b\16\1\2\u00e3"+
		"\u00db\3\2\2\2\u00e3\u00df\3\2\2\2\u00e4\33\3\2\2\2\u00e5\u00e6\b\17\1"+
		"\2\u00e6\u00e7\5\36\20\2\u00e7\u00ec\3\2\2\2\u00e8\u00e9\f\3\2\2\u00e9"+
		"\u00eb\5\36\20\2\u00ea\u00e8\3\2\2\2\u00eb\u00ee\3\2\2\2\u00ec\u00ea\3"+
		"\2\2\2\u00ec\u00ed\3\2\2\2\u00ed\35\3\2\2\2\u00ee\u00ec\3\2\2\2\u00ef"+
		"\u00fd\5\22\n\2\u00f0\u00f1\7\27\2\2\u00f1\u00f2\7.\2\2\u00f2\u00f3\7"+
		")\2\2\u00f3\u00fd\b\20\1\2\u00f4\u00f5\7\26\2\2\u00f5\u00f6\7.\2\2\u00f6"+
		"\u00f7\7+\2\2\u00f7\u00fd\b\20\1\2\u00f8\u00f9\7\26\2\2\u00f9\u00fa\7"+
		".\2\2\u00fa\u00fb\7,\2\2\u00fb\u00fd\b\20\1\2\u00fc\u00ef\3\2\2\2\u00fc"+
		"\u00f0\3\2\2\2\u00fc\u00f4\3\2\2\2\u00fc\u00f8\3\2\2\2\u00fd\37\3\2\2"+
		"\2\u00fe\u00ff\b\21\1\2\u00ff\u0100\5\"\22\2\u0100\u0105\3\2\2\2\u0101"+
		"\u0102\f\3\2\2\u0102\u0104\5\"\22\2\u0103\u0101\3\2\2\2\u0104\u0107\3"+
		"\2\2\2\u0105\u0103\3\2\2\2\u0105\u0106\3\2\2\2\u0106!\3\2\2\2\u0107\u0105"+
		"\3\2\2\2\u0108\u0109\7\25\2\2\u0109\u010a\7.\2\2\u010a\u010b\7+\2\2\u010b"+
		"\u0115\b\22\1\2\u010c\u010d\7\25\2\2\u010d\u010e\7.\2\2\u010e\u010f\7"+
		"-\2\2\u010f\u0115\b\22\1\2\u0110\u0111\7\26\2\2\u0111\u0112\7.\2\2\u0112"+
		"\u0113\7,\2\2\u0113\u0115\b\22\1\2\u0114\u0108\3\2\2\2\u0114\u010c\3\2"+
		"\2\2\u0114\u0110\3\2\2\2\u0115#\3\2\2\2\u0116\u0117\b\23\1\2\u0117\u0118"+
		"\5&\24\2\u0118\u011d\3\2\2\2\u0119\u011a\f\3\2\2\u011a\u011c\5&\24\2\u011b"+
		"\u0119\3\2\2\2\u011c\u011f\3\2\2\2\u011d\u011b\3\2\2\2\u011d\u011e\3\2"+
		"\2\2\u011e%\3\2\2\2\u011f\u011d\3\2\2\2\u0120\u0121\7\30\2\2\u0121\u0122"+
		"\7.\2\2\u0122\u0123\7,\2\2\u0123\u0129\b\24\1\2\u0124\u0125\7\27\2\2\u0125"+
		"\u0126\7.\2\2\u0126\u0127\7)\2\2\u0127\u0129\b\24\1\2\u0128\u0120\3\2"+
		"\2\2\u0128\u0124\3\2\2\2\u0129\'\3\2\2\2\u012a\u012b\b\25\1\2\u012b\u012c"+
		"\5*\26\2\u012c\u0131\3\2\2\2\u012d\u012e\f\3\2\2\u012e\u0130\5*\26\2\u012f"+
		"\u012d\3\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131\u0132\3\2"+
		"\2\2\u0132)\3\2\2\2\u0133\u0131\3\2\2\2\u0134\u0135\7\36\2\2\u0135\u0136"+
		"\7.\2\2\u0136\u0137\7,\2\2\u0137\u0149\b\26\1\2\u0138\u0139\7\36\2\2\u0139"+
		"\u013a\7.\2\2\u013a\u013b\7+\2\2\u013b\u0149\b\26\1\2\u013c\u013d\7\37"+
		"\2\2\u013d\u013e\7.\2\2\u013e\u013f\7,\2\2\u013f\u0149\b\26\1\2\u0140"+
		"\u0141\7\37\2\2\u0141\u0142\7.\2\2\u0142\u0143\7+\2\2\u0143\u0149\b\26"+
		"\1\2\u0144\u0145\7\30\2\2\u0145\u0146\7.\2\2\u0146\u0147\7,\2\2\u0147"+
		"\u0149\b\26\1\2\u0148\u0134\3\2\2\2\u0148\u0138\3\2\2\2\u0148\u013c\3"+
		"\2\2\2\u0148\u0140\3\2\2\2\u0148\u0144\3\2\2\2\u0149+\3\2\2\2\u014a\u014f"+
		"\5.\30\2\u014b\u014c\5(\25\2\u014c\u014d\5.\30\2\u014d\u014f\3\2\2\2\u014e"+
		"\u014a\3\2\2\2\u014e\u014b\3\2\2\2\u014f-\3\2\2\2\u0150\u0151\7\36\2\2"+
		"\u0151\u0152\7.\2\2\u0152\u0153\7,\2\2\u0153\u0169\b\30\1\2\u0154\u0155"+
		"\7\36\2\2\u0155\u0156\7.\2\2\u0156\u0157\7+\2\2\u0157\u0169\b\30\1\2\u0158"+
		"\u0159\7 \2\2\u0159\u015a\7.\2\2\u015a\u015b\7,\2\2\u015b\u0169\b\30\1"+
		"\2\u015c\u015d\7 \2\2\u015d\u015e\7.\2\2\u015e\u015f\7+\2\2\u015f\u0169"+
		"\b\30\1\2\u0160\u0161\7!\2\2\u0161\u0162\7.\2\2\u0162\u0163\7,\2\2\u0163"+
		"\u0169\b\30\1\2\u0164\u0165\7!\2\2\u0165\u0166\7.\2\2\u0166\u0167\7+\2"+
		"\2\u0167\u0169\b\30\1\2\u0168\u0150\3\2\2\2\u0168\u0154\3\2\2\2\u0168"+
		"\u0158\3\2\2\2\u0168\u015c\3\2\2\2\u0168\u0160\3\2\2\2\u0168\u0164\3\2"+
		"\2\2\u0169/\3\2\2\2\27\66Ier\u0093\u009c\u00bb\u00c5\u00cf\u00d9\u00e3"+
		"\u00ec\u00fc\u0105\u0114\u011d\u0128\u0131\u0148\u014e\u0168";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from /home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/Proyecto2/ChemsLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ChemsLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"MKDISK", "RMDISK", "FDISK", "MOUNT", "MKFS", "LOGIN", "LOGOUT", "MKGRP", 
			"RMGRP", "MKUSER", "RMUSR", "MKFILE", "MKDIR", "EXEC", "REP", "SIZE", 
			"FIT", "UNIT", "PATH", "NAME", "TYPE", "ID", "BF", "FF", "WF", "FAST", 
			"FULL", "USR", "PWD", "PWD1", "GRP", "R", "CONT", "P", "PAUSA", "DISK", 
			"TREE", "FILE", "LETRA", "ENTERO", "CARACTER", "IDENTIFICADOR", "PASSWORD", 
			"CADENA", "EXTENSION", "DIAGONAL", "RUTA", "DIRECTORIO", "IGUAL", "WHITESPACE"
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


	public ChemsLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ChemsLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	@Override
	public void action(RuleContext _localctx, int ruleIndex, int actionIndex) {
		switch (ruleIndex) {
		case 40:
			CARACTER_action((RuleContext)_localctx, actionIndex);
			break;
		case 41:
			IDENTIFICADOR_action((RuleContext)_localctx, actionIndex);
			break;
		case 42:
			PASSWORD_action((RuleContext)_localctx, actionIndex);
			break;
		case 44:
			EXTENSION_action((RuleContext)_localctx, actionIndex);
			break;
		case 46:
			RUTA_action((RuleContext)_localctx, actionIndex);
			break;
		case 47:
			DIRECTORIO_action((RuleContext)_localctx, actionIndex);
			break;
		}
	}
	private void CARACTER_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 0:
			[a-zA-Z]
			break;
		case 1:
			('-')?[0-9]+
			break;
		}
	}
	private void IDENTIFICADOR_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 2:
			[a-zA-Z]
			break;
		case 3:
			[a-zA-Z]
			break;
		case 4:
			('-')?[0-9]+
			break;
		}
	}
	private void PASSWORD_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 5:
			[a-zA-Z]
			break;
		case 6:
			[a-zA-Z]
			break;
		}
	}
	private void EXTENSION_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 7:
			{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*
			break;
		}
	}
	private void RUTA_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 8:
			'/'
			break;
		case 9:
			{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*
			break;
		case 10:
			'/'
			break;
		case 11:
			{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*
			break;
		case 12:
			'.'{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*}
			break;
		}
	}
	private void DIRECTORIO_action(RuleContext _localctx, int actionIndex) {
		switch (actionIndex) {
		case 13:
			'/'
			break;
		case 14:
			{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*
			break;
		}
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0180\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3"+
		"\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3"+
		"\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3"+
		"\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\27\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3"+
		"\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3"+
		"\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3"+
		"\"\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3"+
		"\'\3\'\3\'\3(\3(\3)\5)\u013e\n)\3)\6)\u0141\n)\r)\16)\u0142\3*\3*\5*\u0147"+
		"\n*\3+\3+\3+\3+\7+\u014d\n+\f+\16+\u0150\13+\3,\3,\3,\6,\u0155\n,\r,\16"+
		",\u0156\3-\3-\7-\u015b\n-\f-\16-\u015e\13-\3-\3-\3.\3.\3.\3/\3/\3\60\3"+
		"\60\7\60\u0169\n\60\f\60\16\60\u016c\13\60\3\60\3\60\3\60\3\60\3\61\3"+
		"\61\6\61\u0174\n\61\r\61\16\61\u0175\3\62\3\62\3\63\6\63\u017b\n\63\r"+
		"\63\16\63\u017c\3\63\3\63\2\2\64\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31"+
		"\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60"+
		"_\61a\62c\63e\64\3\2\7\4\2C\\c|\3\2\62;\6\2##&&,-BB\3\2$$\6\2\13\f\17"+
		"\17\"\"^^\2\u018c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2"+
		"!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3"+
		"\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2"+
		"\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E"+
		"\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2"+
		"\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2"+
		"\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\3g\3\2\2\2\5n\3\2\2\2\7u"+
		"\3\2\2\2\t{\3\2\2\2\13\u0081\3\2\2\2\r\u0086\3\2\2\2\17\u008c\3\2\2\2"+
		"\21\u0093\3\2\2\2\23\u0099\3\2\2\2\25\u009f\3\2\2\2\27\u00a6\3\2\2\2\31"+
		"\u00ac\3\2\2\2\33\u00b3\3\2\2\2\35\u00b9\3\2\2\2\37\u00be\3\2\2\2!\u00c2"+
		"\3\2\2\2#\u00c8\3\2\2\2%\u00cd\3\2\2\2\'\u00d3\3\2\2\2)\u00d9\3\2\2\2"+
		"+\u00df\3\2\2\2-\u00e5\3\2\2\2/\u00e9\3\2\2\2\61\u00ec\3\2\2\2\63\u00ef"+
		"\3\2\2\2\65\u00f2\3\2\2\2\67\u00f7\3\2\2\29\u00fc\3\2\2\2;\u0105\3\2\2"+
		"\2=\u010f\3\2\2\2?\u0114\3\2\2\2A\u0119\3\2\2\2C\u011c\3\2\2\2E\u0122"+
		"\3\2\2\2G\u0125\3\2\2\2I\u012b\3\2\2\2K\u0130\3\2\2\2M\u0135\3\2\2\2O"+
		"\u013a\3\2\2\2Q\u013d\3\2\2\2S\u0146\3\2\2\2U\u0148\3\2\2\2W\u0154\3\2"+
		"\2\2Y\u0158\3\2\2\2[\u0161\3\2\2\2]\u0164\3\2\2\2_\u016a\3\2\2\2a\u0173"+
		"\3\2\2\2c\u0177\3\2\2\2e\u017a\3\2\2\2gh\7o\2\2hi\7m\2\2ij\7f\2\2jk\7"+
		"k\2\2kl\7u\2\2lm\7m\2\2m\4\3\2\2\2no\7t\2\2op\7o\2\2pq\7f\2\2qr\7k\2\2"+
		"rs\7u\2\2st\7m\2\2t\6\3\2\2\2uv\7h\2\2vw\7f\2\2wx\7k\2\2xy\7u\2\2yz\7"+
		"m\2\2z\b\3\2\2\2{|\7o\2\2|}\7q\2\2}~\7w\2\2~\177\7p\2\2\177\u0080\7v\2"+
		"\2\u0080\n\3\2\2\2\u0081\u0082\7o\2\2\u0082\u0083\7m\2\2\u0083\u0084\7"+
		"h\2\2\u0084\u0085\7u\2\2\u0085\f\3\2\2\2\u0086\u0087\7n\2\2\u0087\u0088"+
		"\7q\2\2\u0088\u0089\7i\2\2\u0089\u008a\7k\2\2\u008a\u008b\7p\2\2\u008b"+
		"\16\3\2\2\2\u008c\u008d\7n\2\2\u008d\u008e\7q\2\2\u008e\u008f\7i\2\2\u008f"+
		"\u0090\7q\2\2\u0090\u0091\7w\2\2\u0091\u0092\7v\2\2\u0092\20\3\2\2\2\u0093"+
		"\u0094\7o\2\2\u0094\u0095\7m\2\2\u0095\u0096\7i\2\2\u0096\u0097\7t\2\2"+
		"\u0097\u0098\7r\2\2\u0098\22\3\2\2\2\u0099\u009a\7t\2\2\u009a\u009b\7"+
		"o\2\2\u009b\u009c\7i\2\2\u009c\u009d\7t\2\2\u009d\u009e\7r\2\2\u009e\24"+
		"\3\2\2\2\u009f\u00a0\7o\2\2\u00a0\u00a1\7m\2\2\u00a1\u00a2\7w\2\2\u00a2"+
		"\u00a3\7u\2\2\u00a3\u00a4\7g\2\2\u00a4\u00a5\7t\2\2\u00a5\26\3\2\2\2\u00a6"+
		"\u00a7\7t\2\2\u00a7\u00a8\7o\2\2\u00a8\u00a9\7w\2\2\u00a9\u00aa\7u\2\2"+
		"\u00aa\u00ab\7t\2\2\u00ab\30\3\2\2\2\u00ac\u00ad\7o\2\2\u00ad\u00ae\7"+
		"m\2\2\u00ae\u00af\7h\2\2\u00af\u00b0\7k\2\2\u00b0\u00b1\7n\2\2\u00b1\u00b2"+
		"\7g\2\2\u00b2\32\3\2\2\2\u00b3\u00b4\7o\2\2\u00b4\u00b5\7m\2\2\u00b5\u00b6"+
		"\7f\2\2\u00b6\u00b7\7k\2\2\u00b7\u00b8\7t\2\2\u00b8\34\3\2\2\2\u00b9\u00ba"+
		"\7g\2\2\u00ba\u00bb\7z\2\2\u00bb\u00bc\7g\2\2\u00bc\u00bd\7e\2\2\u00bd"+
		"\36\3\2\2\2\u00be\u00bf\7t\2\2\u00bf\u00c0\7g\2\2\u00c0\u00c1\7r\2\2\u00c1"+
		" \3\2\2\2\u00c2\u00c3\7/\2\2\u00c3\u00c4\7u\2\2\u00c4\u00c5\7k\2\2\u00c5"+
		"\u00c6\7|\2\2\u00c6\u00c7\7g\2\2\u00c7\"\3\2\2\2\u00c8\u00c9\7/\2\2\u00c9"+
		"\u00ca\7h\2\2\u00ca\u00cb\7k\2\2\u00cb\u00cc\7v\2\2\u00cc$\3\2\2\2\u00cd"+
		"\u00ce\7/\2\2\u00ce\u00cf\7w\2\2\u00cf\u00d0\7p\2\2\u00d0\u00d1\7k\2\2"+
		"\u00d1\u00d2\7v\2\2\u00d2&\3\2\2\2\u00d3\u00d4\7/\2\2\u00d4\u00d5\7r\2"+
		"\2\u00d5\u00d6\7c\2\2\u00d6\u00d7\7v\2\2\u00d7\u00d8\7j\2\2\u00d8(\3\2"+
		"\2\2\u00d9\u00da\7/\2\2\u00da\u00db\7p\2\2\u00db\u00dc\7c\2\2\u00dc\u00dd"+
		"\7o\2\2\u00dd\u00de\7g\2\2\u00de*\3\2\2\2\u00df\u00e0\7/\2\2\u00e0\u00e1"+
		"\7v\2\2\u00e1\u00e2\7{\2\2\u00e2\u00e3\7r\2\2\u00e3\u00e4\7g\2\2\u00e4"+
		",\3\2\2\2\u00e5\u00e6\7/\2\2\u00e6\u00e7\7k\2\2\u00e7\u00e8\7f\2\2\u00e8"+
		".\3\2\2\2\u00e9\u00ea\7d\2\2\u00ea\u00eb\7h\2\2\u00eb\60\3\2\2\2\u00ec"+
		"\u00ed\7h\2\2\u00ed\u00ee\7h\2\2\u00ee\62\3\2\2\2\u00ef\u00f0\7y\2\2\u00f0"+
		"\u00f1\7h\2\2\u00f1\64\3\2\2\2\u00f2\u00f3\7h\2\2\u00f3\u00f4\7c\2\2\u00f4"+
		"\u00f5\7u\2\2\u00f5\u00f6\7v\2\2\u00f6\66\3\2\2\2\u00f7\u00f8\7H\2\2\u00f8"+
		"\u00f9\7W\2\2\u00f9\u00fa\7N\2\2\u00fa\u00fb\7N\2\2\u00fb8\3\2\2\2\u00fc"+
		"\u00fd\7/\2\2\u00fd\u00fe\7w\2\2\u00fe\u00ff\7u\2\2\u00ff\u0100\7w\2\2"+
		"\u0100\u0101\7c\2\2\u0101\u0102\7t\2\2\u0102\u0103\7k\2\2\u0103\u0104"+
		"\7q\2\2\u0104:\3\2\2\2\u0105\u0106\7/\2\2\u0106\u0107\7r\2\2\u0107\u0108"+
		"\7c\2\2\u0108\u0109\7u\2\2\u0109\u010a\7u\2\2\u010a\u010b\7y\2\2\u010b"+
		"\u010c\7q\2\2\u010c\u010d\7t\2\2\u010d\u010e\7f\2\2\u010e<\3\2\2\2\u010f"+
		"\u0110\7/\2\2\u0110\u0111\7r\2\2\u0111\u0112\7y\2\2\u0112\u0113\7f\2\2"+
		"\u0113>\3\2\2\2\u0114\u0115\7/\2\2\u0115\u0116\7i\2\2\u0116\u0117\7t\2"+
		"\2\u0117\u0118\7r\2\2\u0118@\3\2\2\2\u0119\u011a\7/\2\2\u011a\u011b\7"+
		"t\2\2\u011bB\3\2\2\2\u011c\u011d\7/\2\2\u011d\u011e\7e\2\2\u011e\u011f"+
		"\7q\2\2\u011f\u0120\7p\2\2\u0120\u0121\7v\2\2\u0121D\3\2\2\2\u0122\u0123"+
		"\7/\2\2\u0123\u0124\7r\2\2\u0124F\3\2\2\2\u0125\u0126\7r\2\2\u0126\u0127"+
		"\7c\2\2\u0127\u0128\7w\2\2\u0128\u0129\7u\2\2\u0129\u012a\7g\2\2\u012a"+
		"H\3\2\2\2\u012b\u012c\7f\2\2\u012c\u012d\7k\2\2\u012d\u012e\7u\2\2\u012e"+
		"\u012f\7m\2\2\u012fJ\3\2\2\2\u0130\u0131\7v\2\2\u0131\u0132\7t\2\2\u0132"+
		"\u0133\7g\2\2\u0133\u0134\7g\2\2\u0134L\3\2\2\2\u0135\u0136\7h\2\2\u0136"+
		"\u0137\7k\2\2\u0137\u0138\7n\2\2\u0138\u0139\7g\2\2\u0139N\3\2\2\2\u013a"+
		"\u013b\t\2\2\2\u013bP\3\2\2\2\u013c\u013e\7/\2\2\u013d\u013c\3\2\2\2\u013d"+
		"\u013e\3\2\2\2\u013e\u0140\3\2\2\2\u013f\u0141\t\3\2\2\u0140\u013f\3\2"+
		"\2\2\u0141\u0142\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143\3\2\2\2\u0143"+
		"R\3\2\2\2\u0144\u0147\b*\2\2\u0145\u0147\b*\3\2\u0146\u0144\3\2\2\2\u0146"+
		"\u0145\3\2\2\2\u0147T\3\2\2\2\u0148\u014e\b+\4\2\u0149\u014d\b+\5\2\u014a"+
		"\u014d\b+\6\2\u014b\u014d\7a\2\2\u014c\u0149\3\2\2\2\u014c\u014a\3\2\2"+
		"\2\u014c\u014b\3\2\2\2\u014d\u0150\3\2\2\2\u014e\u014c\3\2\2\2\u014e\u014f"+
		"\3\2\2\2\u014fV\3\2\2\2\u0150\u014e\3\2\2\2\u0151\u0155\b,\7\2\u0152\u0155"+
		"\b,\b\2\u0153\u0155\t\4\2\2\u0154\u0151\3\2\2\2\u0154\u0152\3\2\2\2\u0154"+
		"\u0153\3\2\2\2\u0155\u0156\3\2\2\2\u0156\u0154\3\2\2\2\u0156\u0157\3\2"+
		"\2\2\u0157X\3\2\2\2\u0158\u015c\7$\2\2\u0159\u015b\n\5\2\2\u015a\u0159"+
		"\3\2\2\2\u015b\u015e\3\2\2\2\u015c\u015a\3\2\2\2\u015c\u015d\3\2\2\2\u015d"+
		"\u015f\3\2\2\2\u015e\u015c\3\2\2\2\u015f\u0160\7$\2\2\u0160Z\3\2\2\2\u0161"+
		"\u0162\7\60\2\2\u0162\u0163\b.\t\2\u0163\\\3\2\2\2\u0164\u0165\7\61\2"+
		"\2\u0165^\3\2\2\2\u0166\u0167\b\60\n\2\u0167\u0169\b\60\13\2\u0168\u0166"+
		"\3\2\2\2\u0169\u016c\3\2\2\2\u016a\u0168\3\2\2\2\u016a\u016b\3\2\2\2\u016b"+
		"\u016d\3\2\2\2\u016c\u016a\3\2\2\2\u016d\u016e\b\60\f\2\u016e\u016f\b"+
		"\60\r\2\u016f\u0170\b\60\16\2\u0170`\3\2\2\2\u0171\u0172\b\61\17\2\u0172"+
		"\u0174\b\61\20\2\u0173\u0171\3\2\2\2\u0174\u0175\3\2\2\2\u0175\u0173\3"+
		"\2\2\2\u0175\u0176\3\2\2\2\u0176b\3\2\2\2\u0177\u0178\7?\2\2\u0178d\3"+
		"\2\2\2\u0179\u017b\t\6\2\2\u017a\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017c"+
		"\u017a\3\2\2\2\u017c\u017d\3\2\2\2\u017d\u017e\3\2\2\2\u017e\u017f\b\63"+
		"\21\2\u017ff\3\2\2\2\16\2\u013d\u0142\u0146\u014c\u014e\u0154\u0156\u015c"+
		"\u016a\u0175\u017c\22\3*\2\3*\3\3+\4\3+\5\3+\6\3,\7\3,\b\3.\t\3\60\n\3"+
		"\60\13\3\60\f\3\60\r\3\60\16\3\61\17\3\61\20\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
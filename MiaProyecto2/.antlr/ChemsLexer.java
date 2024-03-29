// Generated from /home/curious1924/Escritorio/GitHub/-MIA_Proyecto1_201901907-/MiaProyecto2/ChemsLexer.g4 by ANTLR 4.8
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
		FAST=26, FULL=27, USR=28, PWD=29, PWD1=30, GRP=31, GR=32, CONT=33, GP=34, 
		PAUSA=35, DISK=36, TREE=37, FILE=38, LETRA=39, ENTERO=40, CADENA=41, IDENTIFICADOR=42, 
		RUTA=43, IGUAL=44, WHITESPACE=45, COMENTARIOS=46;
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
			"FULL", "USR", "PWD", "PWD1", "GRP", "GR", "CONT", "GP", "PAUSA", "DISK", 
			"TREE", "FILE", "LETRA", "ENTERO", "CADENA", "IDENTIFICADOR", "RUTA", 
			"IGUAL", "WHITESPACE", "COMENTARIOS", "ESC_SEQ", "A", "B", "C", "D", 
			"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", 
			"S", "T", "U", "V", "W", "X", "Y", "Z"
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\60\u01d5\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t"+
		"=\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4"+
		"I\tI\4J\tJ\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3"+
		"\32\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3"+
		"\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3"+
		"\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3\"\3\""+
		"\3\"\3\"\3\"\3\"\3#\3#\3#\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3"+
		"&\3&\3\'\3\'\3\'\3\'\3\'\3(\3(\3)\6)\u016c\n)\r)\16)\u016d\3*\3*\7*\u0172"+
		"\n*\f*\16*\u0175\13*\3*\3*\3+\3+\7+\u017b\n+\f+\16+\u017e\13+\3,\3,\7"+
		",\u0182\n,\f,\16,\u0185\13,\3,\5,\u0188\n,\3,\3,\3,\3-\3-\3.\6.\u0190"+
		"\n.\r.\16.\u0191\3.\3.\3/\3/\7/\u0198\n/\f/\16/\u019b\13/\3/\3/\3\60\3"+
		"\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3"+
		"\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3@\3@\3A\3A\3"+
		"B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3G\3G\3H\3H\3I\3I\3J\3J\2\2K\3\3\5\4\7\5"+
		"\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\60_\2a\2c\2e\2g\2i\2k\2m\2o\2q\2s\2u\2w\2y\2{"+
		"\2}\2\177\2\u0081\2\u0083\2\u0085\2\u0087\2\u0089\2\u008b\2\u008d\2\u008f"+
		"\2\u0091\2\u0093\2\3\2#\4\2C\\c|\3\2\62;\3\2$$\6\2\62;C\\aac|\6\2\13\f"+
		"\17\17\"\"^^\3\2%%\t\2\"#%%--/\60<<BB]_\4\2CCcc\4\2DDdd\4\2EEee\4\2FF"+
		"ff\4\2GGgg\4\2HHhh\4\2IIii\4\2JJjj\4\2KKkk\4\2LLll\4\2MMmm\4\2NNnn\4\2"+
		"OOoo\4\2PPpp\4\2QQqq\4\2RRrr\4\2SSss\4\2TTtt\4\2UUuu\4\2VVvv\4\2WWww\4"+
		"\2XXxx\4\2YYyy\4\2ZZzz\4\2[[{{\4\2\\\\||\2\u01c0\2\3\3\2\2\2\2\5\3\2\2"+
		"\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\3\u0095\3\2\2\2\5\u009c\3\2\2\2\7\u00a3"+
		"\3\2\2\2\t\u00a9\3\2\2\2\13\u00af\3\2\2\2\r\u00b4\3\2\2\2\17\u00ba\3\2"+
		"\2\2\21\u00c1\3\2\2\2\23\u00c7\3\2\2\2\25\u00cd\3\2\2\2\27\u00d4\3\2\2"+
		"\2\31\u00da\3\2\2\2\33\u00e1\3\2\2\2\35\u00e7\3\2\2\2\37\u00ec\3\2\2\2"+
		"!\u00f0\3\2\2\2#\u00f6\3\2\2\2%\u00fb\3\2\2\2\'\u0101\3\2\2\2)\u0107\3"+
		"\2\2\2+\u010d\3\2\2\2-\u0113\3\2\2\2/\u0117\3\2\2\2\61\u011a\3\2\2\2\63"+
		"\u011d\3\2\2\2\65\u0120\3\2\2\2\67\u0125\3\2\2\29\u012a\3\2\2\2;\u0133"+
		"\3\2\2\2=\u013d\3\2\2\2?\u0142\3\2\2\2A\u0147\3\2\2\2C\u014a\3\2\2\2E"+
		"\u0150\3\2\2\2G\u0153\3\2\2\2I\u0159\3\2\2\2K\u015e\3\2\2\2M\u0163\3\2"+
		"\2\2O\u0168\3\2\2\2Q\u016b\3\2\2\2S\u016f\3\2\2\2U\u0178\3\2\2\2W\u0183"+
		"\3\2\2\2Y\u018c\3\2\2\2[\u018f\3\2\2\2]\u0195\3\2\2\2_\u019e\3\2\2\2a"+
		"\u01a1\3\2\2\2c\u01a3\3\2\2\2e\u01a5\3\2\2\2g\u01a7\3\2\2\2i\u01a9\3\2"+
		"\2\2k\u01ab\3\2\2\2m\u01ad\3\2\2\2o\u01af\3\2\2\2q\u01b1\3\2\2\2s\u01b3"+
		"\3\2\2\2u\u01b5\3\2\2\2w\u01b7\3\2\2\2y\u01b9\3\2\2\2{\u01bb\3\2\2\2}"+
		"\u01bd\3\2\2\2\177\u01bf\3\2\2\2\u0081\u01c1\3\2\2\2\u0083\u01c3\3\2\2"+
		"\2\u0085\u01c5\3\2\2\2\u0087\u01c7\3\2\2\2\u0089\u01c9\3\2\2\2\u008b\u01cb"+
		"\3\2\2\2\u008d\u01cd\3\2\2\2\u008f\u01cf\3\2\2\2\u0091\u01d1\3\2\2\2\u0093"+
		"\u01d3\3\2\2\2\u0095\u0096\5y=\2\u0096\u0097\5u;\2\u0097\u0098\5g\64\2"+
		"\u0098\u0099\5q9\2\u0099\u009a\5\u0085C\2\u009a\u009b\5u;\2\u009b\4\3"+
		"\2\2\2\u009c\u009d\5\u0083B\2\u009d\u009e\5y=\2\u009e\u009f\5g\64\2\u009f"+
		"\u00a0\5q9\2\u00a0\u00a1\5\u0085C\2\u00a1\u00a2\5u;\2\u00a2\6\3\2\2\2"+
		"\u00a3\u00a4\5k\66\2\u00a4\u00a5\5g\64\2\u00a5\u00a6\5q9\2\u00a6\u00a7"+
		"\5\u0085C\2\u00a7\u00a8\5u;\2\u00a8\b\3\2\2\2\u00a9\u00aa\5y=\2\u00aa"+
		"\u00ab\5}?\2\u00ab\u00ac\5\u0089E\2\u00ac\u00ad\5{>\2\u00ad\u00ae\5\u0087"+
		"D\2\u00ae\n\3\2\2\2\u00af\u00b0\5y=\2\u00b0\u00b1\5u;\2\u00b1\u00b2\5"+
		"k\66\2\u00b2\u00b3\5\u0085C\2\u00b3\f\3\2\2\2\u00b4\u00b5\5w<\2\u00b5"+
		"\u00b6\5}?\2\u00b6\u00b7\5m\67\2\u00b7\u00b8\5q9\2\u00b8\u00b9\5{>\2\u00b9"+
		"\16\3\2\2\2\u00ba\u00bb\5w<\2\u00bb\u00bc\5}?\2\u00bc\u00bd\5m\67\2\u00bd"+
		"\u00be\5}?\2\u00be\u00bf\5\u0089E\2\u00bf\u00c0\5\u0087D\2\u00c0\20\3"+
		"\2\2\2\u00c1\u00c2\5y=\2\u00c2\u00c3\5u;\2\u00c3\u00c4\5m\67\2\u00c4\u00c5"+
		"\5\u0083B\2\u00c5\u00c6\5\177@\2\u00c6\22\3\2\2\2\u00c7\u00c8\5\u0083"+
		"B\2\u00c8\u00c9\5y=\2\u00c9\u00ca\5m\67\2\u00ca\u00cb\5\u0083B\2\u00cb"+
		"\u00cc\5\177@\2\u00cc\24\3\2\2\2\u00cd\u00ce\5y=\2\u00ce\u00cf\5u;\2\u00cf"+
		"\u00d0\5\u0089E\2\u00d0\u00d1\5\u0085C\2\u00d1\u00d2\5i\65\2\u00d2\u00d3"+
		"\5\u0083B\2\u00d3\26\3\2\2\2\u00d4\u00d5\5\u0083B\2\u00d5\u00d6\5y=\2"+
		"\u00d6\u00d7\5\u0089E\2\u00d7\u00d8\5\u0085C\2\u00d8\u00d9\5\u0083B\2"+
		"\u00d9\30\3\2\2\2\u00da\u00db\5y=\2\u00db\u00dc\5u;\2\u00dc\u00dd\5k\66"+
		"\2\u00dd\u00de\5q9\2\u00de\u00df\5w<\2\u00df\u00e0\5i\65\2\u00e0\32\3"+
		"\2\2\2\u00e1\u00e2\5y=\2\u00e2\u00e3\5u;\2\u00e3\u00e4\5g\64\2\u00e4\u00e5"+
		"\5q9\2\u00e5\u00e6\5\u0083B\2\u00e6\34\3\2\2\2\u00e7\u00e8\5i\65\2\u00e8"+
		"\u00e9\5\u008fH\2\u00e9\u00ea\5i\65\2\u00ea\u00eb\5e\63\2\u00eb\36\3\2"+
		"\2\2\u00ec\u00ed\5\u0083B\2\u00ed\u00ee\5i\65\2\u00ee\u00ef\5\177@\2\u00ef"+
		" \3\2\2\2\u00f0\u00f1\7/\2\2\u00f1\u00f2\5\u0085C\2\u00f2\u00f3\5q9\2"+
		"\u00f3\u00f4\5\u0093J\2\u00f4\u00f5\5i\65\2\u00f5\"\3\2\2\2\u00f6\u00f7"+
		"\7/\2\2\u00f7\u00f8\5k\66\2\u00f8\u00f9\5q9\2\u00f9\u00fa\5\u0087D\2\u00fa"+
		"$\3\2\2\2\u00fb\u00fc\7/\2\2\u00fc\u00fd\5\u0089E\2\u00fd\u00fe\5{>\2"+
		"\u00fe\u00ff\5q9\2\u00ff\u0100\5\u0087D\2\u0100&\3\2\2\2\u0101\u0102\7"+
		"/\2\2\u0102\u0103\5\177@\2\u0103\u0104\5a\61\2\u0104\u0105\5\u0087D\2"+
		"\u0105\u0106\5o8\2\u0106(\3\2\2\2\u0107\u0108\7/\2\2\u0108\u0109\5{>\2"+
		"\u0109\u010a\5a\61\2\u010a\u010b\5y=\2\u010b\u010c\5i\65\2\u010c*\3\2"+
		"\2\2\u010d\u010e\7/\2\2\u010e\u010f\5\u0087D\2\u010f\u0110\5\u0091I\2"+
		"\u0110\u0111\5\177@\2\u0111\u0112\5i\65\2\u0112,\3\2\2\2\u0113\u0114\7"+
		"/\2\2\u0114\u0115\5q9\2\u0115\u0116\5g\64\2\u0116.\3\2\2\2\u0117\u0118"+
		"\5c\62\2\u0118\u0119\5k\66\2\u0119\60\3\2\2\2\u011a\u011b\5k\66\2\u011b"+
		"\u011c\5k\66\2\u011c\62\3\2\2\2\u011d\u011e\5\u008dG\2\u011e\u011f\5k"+
		"\66\2\u011f\64\3\2\2\2\u0120\u0121\5k\66\2\u0121\u0122\5a\61\2\u0122\u0123"+
		"\5\u0085C\2\u0123\u0124\5\u0087D\2\u0124\66\3\2\2\2\u0125\u0126\5k\66"+
		"\2\u0126\u0127\5\u0089E\2\u0127\u0128\5w<\2\u0128\u0129\5w<\2\u01298\3"+
		"\2\2\2\u012a\u012b\7/\2\2\u012b\u012c\5\u0089E\2\u012c\u012d\5\u0085C"+
		"\2\u012d\u012e\5\u0089E\2\u012e\u012f\5a\61\2\u012f\u0130\5\u0083B\2\u0130"+
		"\u0131\5q9\2\u0131\u0132\5}?\2\u0132:\3\2\2\2\u0133\u0134\7/\2\2\u0134"+
		"\u0135\5\177@\2\u0135\u0136\5a\61\2\u0136\u0137\5\u0085C\2\u0137\u0138"+
		"\5\u0085C\2\u0138\u0139\5\u008dG\2\u0139\u013a\5}?\2\u013a\u013b\5\u0083"+
		"B\2\u013b\u013c\5g\64\2\u013c<\3\2\2\2\u013d\u013e\7/\2\2\u013e\u013f"+
		"\5\177@\2\u013f\u0140\5\u008dG\2\u0140\u0141\5g\64\2\u0141>\3\2\2\2\u0142"+
		"\u0143\7/\2\2\u0143\u0144\5m\67\2\u0144\u0145\5\u0083B\2\u0145\u0146\5"+
		"\177@\2\u0146@\3\2\2\2\u0147\u0148\7/\2\2\u0148\u0149\5\u0083B\2\u0149"+
		"B\3\2\2\2\u014a\u014b\7/\2\2\u014b\u014c\5e\63\2\u014c\u014d\5}?\2\u014d"+
		"\u014e\5{>\2\u014e\u014f\5\u0087D\2\u014fD\3\2\2\2\u0150\u0151\7/\2\2"+
		"\u0151\u0152\5\177@\2\u0152F\3\2\2\2\u0153\u0154\5\177@\2\u0154\u0155"+
		"\5a\61\2\u0155\u0156\5\u0089E\2\u0156\u0157\5\u0085C\2\u0157\u0158\5i"+
		"\65\2\u0158H\3\2\2\2\u0159\u015a\5g\64\2\u015a\u015b\5q9\2\u015b\u015c"+
		"\5\u0085C\2\u015c\u015d\5u;\2\u015dJ\3\2\2\2\u015e\u015f\5\u0087D\2\u015f"+
		"\u0160\5\u0083B\2\u0160\u0161\5i\65\2\u0161\u0162\5i\65\2\u0162L\3\2\2"+
		"\2\u0163\u0164\5k\66\2\u0164\u0165\5q9\2\u0165\u0166\5w<\2\u0166\u0167"+
		"\5i\65\2\u0167N\3\2\2\2\u0168\u0169\t\2\2\2\u0169P\3\2\2\2\u016a\u016c"+
		"\t\3\2\2\u016b\u016a\3\2\2\2\u016c\u016d\3\2\2\2\u016d\u016b\3\2\2\2\u016d"+
		"\u016e\3\2\2\2\u016eR\3\2\2\2\u016f\u0173\7$\2\2\u0170\u0172\n\4\2\2\u0171"+
		"\u0170\3\2\2\2\u0172\u0175\3\2\2\2\u0173\u0171\3\2\2\2\u0173\u0174\3\2"+
		"\2\2\u0174\u0176\3\2\2\2\u0175\u0173\3\2\2\2\u0176\u0177\7$\2\2\u0177"+
		"T\3\2\2\2\u0178\u017c\t\5\2\2\u0179\u017b\t\5\2\2\u017a\u0179\3\2\2\2"+
		"\u017b\u017e\3\2\2\2\u017c\u017a\3\2\2\2\u017c\u017d\3\2\2\2\u017dV\3"+
		"\2\2\2\u017e\u017c\3\2\2\2\u017f\u0180\7\61\2\2\u0180\u0182\5U+\2\u0181"+
		"\u017f\3\2\2\2\u0182\u0185\3\2\2\2\u0183\u0181\3\2\2\2\u0183\u0184\3\2"+
		"\2\2\u0184\u0187\3\2\2\2\u0185\u0183\3\2\2\2\u0186\u0188\5U+\2\u0187\u0186"+
		"\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u0189\3\2\2\2\u0189\u018a\7\60\2\2"+
		"\u018a\u018b\5U+\2\u018bX\3\2\2\2\u018c\u018d\7?\2\2\u018dZ\3\2\2\2\u018e"+
		"\u0190\t\6\2\2\u018f\u018e\3\2\2\2\u0190\u0191\3\2\2\2\u0191\u018f\3\2"+
		"\2\2\u0191\u0192\3\2\2\2\u0192\u0193\3\2\2\2\u0193\u0194\b.\2\2\u0194"+
		"\\\3\2\2\2\u0195\u0199\t\7\2\2\u0196\u0198\13\2\2\2\u0197\u0196\3\2\2"+
		"\2\u0198\u019b\3\2\2\2\u0199\u0197\3\2\2\2\u0199\u019a\3\2\2\2\u019a\u019c"+
		"\3\2\2\2\u019b\u0199\3\2\2\2\u019c\u019d\b/\2\2\u019d^\3\2\2\2\u019e\u019f"+
		"\7^\2\2\u019f\u01a0\t\b\2\2\u01a0`\3\2\2\2\u01a1\u01a2\t\t\2\2\u01a2b"+
		"\3\2\2\2\u01a3\u01a4\t\n\2\2\u01a4d\3\2\2\2\u01a5\u01a6\t\13\2\2\u01a6"+
		"f\3\2\2\2\u01a7\u01a8\t\f\2\2\u01a8h\3\2\2\2\u01a9\u01aa\t\r\2\2\u01aa"+
		"j\3\2\2\2\u01ab\u01ac\t\16\2\2\u01acl\3\2\2\2\u01ad\u01ae\t\17\2\2\u01ae"+
		"n\3\2\2\2\u01af\u01b0\t\20\2\2\u01b0p\3\2\2\2\u01b1\u01b2\t\21\2\2\u01b2"+
		"r\3\2\2\2\u01b3\u01b4\t\22\2\2\u01b4t\3\2\2\2\u01b5\u01b6\t\23\2\2\u01b6"+
		"v\3\2\2\2\u01b7\u01b8\t\24\2\2\u01b8x\3\2\2\2\u01b9\u01ba\t\25\2\2\u01ba"+
		"z\3\2\2\2\u01bb\u01bc\t\26\2\2\u01bc|\3\2\2\2\u01bd\u01be\t\27\2\2\u01be"+
		"~\3\2\2\2\u01bf\u01c0\t\30\2\2\u01c0\u0080\3\2\2\2\u01c1\u01c2\t\31\2"+
		"\2\u01c2\u0082\3\2\2\2\u01c3\u01c4\t\32\2\2\u01c4\u0084\3\2\2\2\u01c5"+
		"\u01c6\t\33\2\2\u01c6\u0086\3\2\2\2\u01c7\u01c8\t\34\2\2\u01c8\u0088\3"+
		"\2\2\2\u01c9\u01ca\t\35\2\2\u01ca\u008a\3\2\2\2\u01cb\u01cc\t\36\2\2\u01cc"+
		"\u008c\3\2\2\2\u01cd\u01ce\t\37\2\2\u01ce\u008e\3\2\2\2\u01cf\u01d0\t"+
		" \2\2\u01d0\u0090\3\2\2\2\u01d1\u01d2\t!\2\2\u01d2\u0092\3\2\2\2\u01d3"+
		"\u01d4\t\"\2\2\u01d4\u0094\3\2\2\2\n\2\u016d\u0173\u017c\u0183\u0187\u0191"+
		"\u0199\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
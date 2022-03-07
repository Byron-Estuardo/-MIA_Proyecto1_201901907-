/* A Bison parser, made by GNU Bison 3.5.1.  */

/* Bison implementation for Yacc-like parsers in C

   Copyright (C) 1984, 1989-1990, 2000-2015, 2018-2020 Free Software Foundation,
   Inc.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.

   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* C LALR(1) parser skeleton written by Richard Stallman, by
   simplifying the original so-called "semantic" parser.  */

/* All symbols defined below should begin with yy or YY, to avoid
   infringing on user name space.  This should be done even for local
   variables, as they might otherwise be expanded by user macros.
   There are some unavoidable exceptions within include files to
   define necessary library symbols; they are noted "INFRINGES ON
   USER NAME SPACE" below.  */

/* Undocumented macros, especially those whose name start with YY_,
   are private implementation details.  Do not rely on them.  */

/* Identify Bison output.  */
#define YYBISON 1

/* Bison version.  */
#define YYBISON_VERSION "3.5.1"

/* Skeleton name.  */
#define YYSKELETON_NAME "yacc.c"

/* Pure parsers.  */
#define YYPURE 0

/* Push parsers.  */
#define YYPUSH 0

/* Pull parsers.  */
#define YYPULL 1




/* First part of user prologue.  */
#line 1 "Sintactico.y"

    #include "scanner.h"
    #include "nodo.h"
    #include <iostream>

    extern int yylineno;
    extern int columna;
    extern char *yytext;
    extern Nodo *raiz;

    int yyerror(const char* mens){
        std::cout<<mens<<std::endl;
        return 0;
    }


#line 87 "parser.cpp"

# ifndef YY_CAST
#  ifdef __cplusplus
#   define YY_CAST(Type, Val) static_cast<Type> (Val)
#   define YY_REINTERPRET_CAST(Type, Val) reinterpret_cast<Type> (Val)
#  else
#   define YY_CAST(Type, Val) ((Type) (Val))
#   define YY_REINTERPRET_CAST(Type, Val) ((Type) (Val))
#  endif
# endif
# ifndef YY_NULLPTR
#  if defined __cplusplus
#   if 201103L <= __cplusplus
#    define YY_NULLPTR nullptr
#   else
#    define YY_NULLPTR 0
#   endif
#  else
#   define YY_NULLPTR ((void*)0)
#  endif
# endif

/* Enabling verbose error messages.  */
#ifdef YYERROR_VERBOSE
# undef YYERROR_VERBOSE
# define YYERROR_VERBOSE 1
#else
# define YYERROR_VERBOSE 1
#endif

/* Use api.header.include to #include this header
   instead of duplicating it here.  */
#ifndef YY_YY_PARSER_H_INCLUDED
# define YY_YY_PARSER_H_INCLUDED
/* Debug traces.  */
#ifndef YYDEBUG
# define YYDEBUG 0
#endif
#if YYDEBUG
extern int yydebug;
#endif

/* Token type.  */
#ifndef YYTOKENTYPE
# define YYTOKENTYPE
  enum yytokentype
  {
    mkdisk = 258,
    rmdisk = 259,
    fdisk = 260,
    mount = 261,
    unmount = 262,
    rep = 263,
    exec = 264,
    size = 265,
    unit = 266,
    path = 267,
    fit = 268,
    name = 269,
    type = 270,
    del = 271,
    add = 272,
    id = 273,
    bf = 274,
    ff = 275,
    wf = 276,
    fast = 277,
    full = 278,
    mbr = 279,
    disk = 280,
    igual = 281,
    extension = 282,
    num = 283,
    caracter = 284,
    cadena = 285,
    identificador = 286,
    ruta = 287,
    mkfs = 288,
    login = 289,
    logout = 290,
    mkgrp = 291,
    rmgrp = 292,
    mkusr = 293,
    rmusr = 294,
    Rchmod = 295,
    mkfile = 296,
    cat = 297,
    rem = 298,
    edit = 299,
    ren = 300,
    Rmkdir = 301,
    cp = 302,
    mv = 303,
    Rchown = 304,
    chgrp = 305,
    pausa = 306,
    recovery = 307,
    loss = 308,
    fs = 309,
    fs2 = 310,
    fs3 = 311,
    usr = 312,
    pwd = 313,
    pwd1 = 314,
    grp = 315,
    ugo = 316,
    r = 317,
    p = 318,
    cont = 319,
    cont1 = 320,
    file = 321,
    dest = 322,
    rutaRep = 323,
    inode = 324,
    journaling = 325,
    block = 326,
    bm_inode = 327,
    bm_block = 328,
    tree = 329,
    sb = 330,
    fileRep = 331,
    ls = 332,
    password = 333,
    directorio = 334
  };
#endif

/* Value type.  */
#if ! defined YYSTYPE && ! defined YYSTYPE_IS_DECLARED
union YYSTYPE
{
#line 21 "Sintactico.y"

        char text[400];
        class Nodo *nodito;
    

#line 225 "parser.cpp"

};
typedef union YYSTYPE YYSTYPE;
# define YYSTYPE_IS_TRIVIAL 1
# define YYSTYPE_IS_DECLARED 1
#endif


extern YYSTYPE yylval;

int yyparse (void);

#endif /* !YY_YY_PARSER_H_INCLUDED  */



#ifdef short
# undef short
#endif

/* On compilers that do not define __PTRDIFF_MAX__ etc., make sure
   <limits.h> and (if available) <stdint.h> are included
   so that the code can choose integer types of a good width.  */

#ifndef __PTRDIFF_MAX__
# include <limits.h> /* INFRINGES ON USER NAME SPACE */
# if defined __STDC_VERSION__ && 199901 <= __STDC_VERSION__
#  include <stdint.h> /* INFRINGES ON USER NAME SPACE */
#  define YY_STDINT_H
# endif
#endif

/* Narrow types that promote to a signed type and that can represent a
   signed or unsigned integer of at least N bits.  In tables they can
   save space and decrease cache pressure.  Promoting to a signed type
   helps avoid bugs in integer arithmetic.  */

#ifdef __INT_LEAST8_MAX__
typedef __INT_LEAST8_TYPE__ yytype_int8;
#elif defined YY_STDINT_H
typedef int_least8_t yytype_int8;
#else
typedef signed char yytype_int8;
#endif

#ifdef __INT_LEAST16_MAX__
typedef __INT_LEAST16_TYPE__ yytype_int16;
#elif defined YY_STDINT_H
typedef int_least16_t yytype_int16;
#else
typedef short yytype_int16;
#endif

#if defined __UINT_LEAST8_MAX__ && __UINT_LEAST8_MAX__ <= __INT_MAX__
typedef __UINT_LEAST8_TYPE__ yytype_uint8;
#elif (!defined __UINT_LEAST8_MAX__ && defined YY_STDINT_H \
       && UINT_LEAST8_MAX <= INT_MAX)
typedef uint_least8_t yytype_uint8;
#elif !defined __UINT_LEAST8_MAX__ && UCHAR_MAX <= INT_MAX
typedef unsigned char yytype_uint8;
#else
typedef short yytype_uint8;
#endif

#if defined __UINT_LEAST16_MAX__ && __UINT_LEAST16_MAX__ <= __INT_MAX__
typedef __UINT_LEAST16_TYPE__ yytype_uint16;
#elif (!defined __UINT_LEAST16_MAX__ && defined YY_STDINT_H \
       && UINT_LEAST16_MAX <= INT_MAX)
typedef uint_least16_t yytype_uint16;
#elif !defined __UINT_LEAST16_MAX__ && USHRT_MAX <= INT_MAX
typedef unsigned short yytype_uint16;
#else
typedef int yytype_uint16;
#endif

#ifndef YYPTRDIFF_T
# if defined __PTRDIFF_TYPE__ && defined __PTRDIFF_MAX__
#  define YYPTRDIFF_T __PTRDIFF_TYPE__
#  define YYPTRDIFF_MAXIMUM __PTRDIFF_MAX__
# elif defined PTRDIFF_MAX
#  ifndef ptrdiff_t
#   include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  endif
#  define YYPTRDIFF_T ptrdiff_t
#  define YYPTRDIFF_MAXIMUM PTRDIFF_MAX
# else
#  define YYPTRDIFF_T long
#  define YYPTRDIFF_MAXIMUM LONG_MAX
# endif
#endif

#ifndef YYSIZE_T
# ifdef __SIZE_TYPE__
#  define YYSIZE_T __SIZE_TYPE__
# elif defined size_t
#  define YYSIZE_T size_t
# elif defined __STDC_VERSION__ && 199901 <= __STDC_VERSION__
#  include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  define YYSIZE_T size_t
# else
#  define YYSIZE_T unsigned
# endif
#endif

#define YYSIZE_MAXIMUM                                  \
  YY_CAST (YYPTRDIFF_T,                                 \
           (YYPTRDIFF_MAXIMUM < YY_CAST (YYSIZE_T, -1)  \
            ? YYPTRDIFF_MAXIMUM                         \
            : YY_CAST (YYSIZE_T, -1)))

#define YYSIZEOF(X) YY_CAST (YYPTRDIFF_T, sizeof (X))

/* Stored state numbers (used for stacks). */
typedef yytype_int16 yy_state_t;

/* State numbers in computations.  */
typedef int yy_state_fast_t;

#ifndef YY_
# if defined YYENABLE_NLS && YYENABLE_NLS
#  if ENABLE_NLS
#   include <libintl.h> /* INFRINGES ON USER NAME SPACE */
#   define YY_(Msgid) dgettext ("bison-runtime", Msgid)
#  endif
# endif
# ifndef YY_
#  define YY_(Msgid) Msgid
# endif
#endif

#ifndef YY_ATTRIBUTE_PURE
# if defined __GNUC__ && 2 < __GNUC__ + (96 <= __GNUC_MINOR__)
#  define YY_ATTRIBUTE_PURE __attribute__ ((__pure__))
# else
#  define YY_ATTRIBUTE_PURE
# endif
#endif

#ifndef YY_ATTRIBUTE_UNUSED
# if defined __GNUC__ && 2 < __GNUC__ + (7 <= __GNUC_MINOR__)
#  define YY_ATTRIBUTE_UNUSED __attribute__ ((__unused__))
# else
#  define YY_ATTRIBUTE_UNUSED
# endif
#endif

/* Suppress unused-variable warnings by "using" E.  */
#if ! defined lint || defined __GNUC__
# define YYUSE(E) ((void) (E))
#else
# define YYUSE(E) /* empty */
#endif

#if defined __GNUC__ && ! defined __ICC && 407 <= __GNUC__ * 100 + __GNUC_MINOR__
/* Suppress an incorrect diagnostic about yylval being uninitialized.  */
# define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN                            \
    _Pragma ("GCC diagnostic push")                                     \
    _Pragma ("GCC diagnostic ignored \"-Wuninitialized\"")              \
    _Pragma ("GCC diagnostic ignored \"-Wmaybe-uninitialized\"")
# define YY_IGNORE_MAYBE_UNINITIALIZED_END      \
    _Pragma ("GCC diagnostic pop")
#else
# define YY_INITIAL_VALUE(Value) Value
#endif
#ifndef YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_END
#endif
#ifndef YY_INITIAL_VALUE
# define YY_INITIAL_VALUE(Value) /* Nothing. */
#endif

#if defined __cplusplus && defined __GNUC__ && ! defined __ICC && 6 <= __GNUC__
# define YY_IGNORE_USELESS_CAST_BEGIN                          \
    _Pragma ("GCC diagnostic push")                            \
    _Pragma ("GCC diagnostic ignored \"-Wuseless-cast\"")
# define YY_IGNORE_USELESS_CAST_END            \
    _Pragma ("GCC diagnostic pop")
#endif
#ifndef YY_IGNORE_USELESS_CAST_BEGIN
# define YY_IGNORE_USELESS_CAST_BEGIN
# define YY_IGNORE_USELESS_CAST_END
#endif


#define YY_ASSERT(E) ((void) (0 && (E)))

#if ! defined yyoverflow || YYERROR_VERBOSE

/* The parser invokes alloca or malloc; define the necessary symbols.  */

# ifdef YYSTACK_USE_ALLOCA
#  if YYSTACK_USE_ALLOCA
#   ifdef __GNUC__
#    define YYSTACK_ALLOC __builtin_alloca
#   elif defined __BUILTIN_VA_ARG_INCR
#    include <alloca.h> /* INFRINGES ON USER NAME SPACE */
#   elif defined _AIX
#    define YYSTACK_ALLOC __alloca
#   elif defined _MSC_VER
#    include <malloc.h> /* INFRINGES ON USER NAME SPACE */
#    define alloca _alloca
#   else
#    define YYSTACK_ALLOC alloca
#    if ! defined _ALLOCA_H && ! defined EXIT_SUCCESS
#     include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
      /* Use EXIT_SUCCESS as a witness for stdlib.h.  */
#     ifndef EXIT_SUCCESS
#      define EXIT_SUCCESS 0
#     endif
#    endif
#   endif
#  endif
# endif

# ifdef YYSTACK_ALLOC
   /* Pacify GCC's 'empty if-body' warning.  */
#  define YYSTACK_FREE(Ptr) do { /* empty */; } while (0)
#  ifndef YYSTACK_ALLOC_MAXIMUM
    /* The OS might guarantee only one guard page at the bottom of the stack,
       and a page size can be as small as 4096 bytes.  So we cannot safely
       invoke alloca (N) if N exceeds 4096.  Use a slightly smaller number
       to allow for a few compiler-allocated temporary stack slots.  */
#   define YYSTACK_ALLOC_MAXIMUM 4032 /* reasonable circa 2006 */
#  endif
# else
#  define YYSTACK_ALLOC YYMALLOC
#  define YYSTACK_FREE YYFREE
#  ifndef YYSTACK_ALLOC_MAXIMUM
#   define YYSTACK_ALLOC_MAXIMUM YYSIZE_MAXIMUM
#  endif
#  if (defined __cplusplus && ! defined EXIT_SUCCESS \
       && ! ((defined YYMALLOC || defined malloc) \
             && (defined YYFREE || defined free)))
#   include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
#   ifndef EXIT_SUCCESS
#    define EXIT_SUCCESS 0
#   endif
#  endif
#  ifndef YYMALLOC
#   define YYMALLOC malloc
#   if ! defined malloc && ! defined EXIT_SUCCESS
void *malloc (YYSIZE_T); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
#  ifndef YYFREE
#   define YYFREE free
#   if ! defined free && ! defined EXIT_SUCCESS
void free (void *); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
# endif
#endif /* ! defined yyoverflow || YYERROR_VERBOSE */


#if (! defined yyoverflow \
     && (! defined __cplusplus \
         || (defined YYSTYPE_IS_TRIVIAL && YYSTYPE_IS_TRIVIAL)))

/* A type that is properly aligned for any stack member.  */
union yyalloc
{
  yy_state_t yyss_alloc;
  YYSTYPE yyvs_alloc;
};

/* The size of the maximum gap between one aligned stack and the next.  */
# define YYSTACK_GAP_MAXIMUM (YYSIZEOF (union yyalloc) - 1)

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */
# define YYSTACK_BYTES(N) \
     ((N) * (YYSIZEOF (yy_state_t) + YYSIZEOF (YYSTYPE)) \
      + YYSTACK_GAP_MAXIMUM)

# define YYCOPY_NEEDED 1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */
# define YYSTACK_RELOCATE(Stack_alloc, Stack)                           \
    do                                                                  \
      {                                                                 \
        YYPTRDIFF_T yynewbytes;                                         \
        YYCOPY (&yyptr->Stack_alloc, Stack, yysize);                    \
        Stack = &yyptr->Stack_alloc;                                    \
        yynewbytes = yystacksize * YYSIZEOF (*Stack) + YYSTACK_GAP_MAXIMUM; \
        yyptr += yynewbytes / YYSIZEOF (*yyptr);                        \
      }                                                                 \
    while (0)

#endif

#if defined YYCOPY_NEEDED && YYCOPY_NEEDED
/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */
# ifndef YYCOPY
#  if defined __GNUC__ && 1 < __GNUC__
#   define YYCOPY(Dst, Src, Count) \
      __builtin_memcpy (Dst, Src, YY_CAST (YYSIZE_T, (Count)) * sizeof (*(Src)))
#  else
#   define YYCOPY(Dst, Src, Count)              \
      do                                        \
        {                                       \
          YYPTRDIFF_T yyi;                      \
          for (yyi = 0; yyi < (Count); yyi++)   \
            (Dst)[yyi] = (Src)[yyi];            \
        }                                       \
      while (0)
#  endif
# endif
#endif /* !YYCOPY_NEEDED */

/* YYFINAL -- State number of the termination state.  */
#define YYFINAL  128
/* YYLAST -- Last index in YYTABLE.  */
#define YYLAST   300

/* YYNTOKENS -- Number of terminals.  */
#define YYNTOKENS  80
/* YYNNTS -- Number of nonterminals.  */
#define YYNNTS  45
/* YYNRULES -- Number of rules.  */
#define YYNRULES  172
/* YYNSTATES -- Number of states.  */
#define YYNSTATES  301

#define YYUNDEFTOK  2
#define YYMAXUTOK   334


/* YYTRANSLATE(TOKEN-NUM) -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, with out-of-bounds checking.  */
#define YYTRANSLATE(YYX)                                                \
  (0 <= (YYX) && (YYX) <= YYMAXUTOK ? yytranslate[YYX] : YYUNDEFTOK)

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex.  */
static const yytype_int8 yytranslate[] =
{
       0,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     1,     2,     3,     4,
       5,     6,     7,     8,     9,    10,    11,    12,    13,    14,
      15,    16,    17,    18,    19,    20,    21,    22,    23,    24,
      25,    26,    27,    28,    29,    30,    31,    32,    33,    34,
      35,    36,    37,    38,    39,    40,    41,    42,    43,    44,
      45,    46,    47,    48,    49,    50,    51,    52,    53,    54,
      55,    56,    57,    58,    59,    60,    61,    62,    63,    64,
      65,    66,    67,    68,    69,    70,    71,    72,    73,    74,
      75,    76,    77,    78,    79
};

#if YYDEBUG
  /* YYRLINE[YYN] -- Source line where rule number YYN was defined.  */
static const yytype_int16 yyrline[] =
{
       0,   151,   151,   153,   157,   158,   162,   166,   167,   170,
     171,   175,   179,   180,   181,   182,   186,   187,   191,   195,
     196,   197,   201,   205,   209,   213,   217,   221,   225,   226,
     227,   232,   236,   241,   242,   246,   247,   250,   255,   260,
     265,   269,   274,   275,   276,   277,   278,   279,   280,   282,
     286,   291,   292,   293,   294,   296,   302,   303,   304,   307,
     311,   316,   317,   318,   319,   320,   321,   322,   323,   324,
     325,   326,   327,   328,   329,   330,   331,   333,   338,   344,
     348,   353,   354,   355,   356,   357,   359,   363,   368,   369,
     370,   371,   372,   373,   374,   376,   381,   387,   392,   397,
     401,   406,   407,   408,   409,   410,   411,   412,   413,   415,
     420,   426,   430,   435,   436,   437,   438,   440,   444,   450,
     451,   452,   453,   454,   455,   457,   462,   468,   473,   478,
     484,   488,   493,   494,   495,   497,   501,   506,   507,   508,
     513,   515,   519,   524,   525,   526,   527,   529,   533,   538,
     539,   540,   541,   542,   543,   545,   549,   554,   558,   563,
     564,   565,   566,   567,   568,   570,   574,   579,   580,   581,
     582,   584,   590
};
#endif

#if YYDEBUG || YYERROR_VERBOSE || 1
/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */
static const char *const yytname[] =
{
  "$end", "error", "$undefined", "mkdisk", "rmdisk", "fdisk", "mount",
  "unmount", "rep", "exec", "size", "unit", "path", "fit", "name", "type",
  "del", "add", "id", "bf", "ff", "wf", "fast", "full", "mbr", "disk",
  "igual", "extension", "num", "caracter", "cadena", "identificador",
  "ruta", "mkfs", "login", "logout", "mkgrp", "rmgrp", "mkusr", "rmusr",
  "Rchmod", "mkfile", "cat", "rem", "edit", "ren", "Rmkdir", "cp", "mv",
  "Rchown", "chgrp", "pausa", "recovery", "loss", "fs", "fs2", "fs3",
  "usr", "pwd", "pwd1", "grp", "ugo", "r", "p", "cont", "cont1", "file",
  "dest", "rutaRep", "inode", "journaling", "block", "bm_inode",
  "bm_block", "tree", "sb", "fileRep", "ls", "password", "directorio",
  "$accept", "INIT", "COMANDO", "MKDISK", "PARAMETRO_MK", "RMDISK",
  "FDISK", "PARAMETRO_FK", "MOUNT", "PARAMETRO_M", "UNMOUNT", "AJUSTE",
  "REP", "PARAMETRO_R", "SCRIPT", "MKFS", "PARAM_MKFS", "LOGIN",
  "PARAM_LOGIN", "MKGRP", "RMGRP", "MKUSR", "PARAM_MKUSR", "RMUSR",
  "CHMOD", "PARAM_CHMOD", "MKFILE", "PARAM_MKFILE", "CAT", "REM", "EDIT",
  "PARAM_EDIT", "REN", "PARAM_REN", "MKDIR", "PARAM_MKDIR", "CP",
  "PARAM_CP", "MV", "CHOWN", "PARAM_CHOWN", "CHGRP", "PARAM_CHGRP",
  "RECOVERY", "LOSS", YY_NULLPTR
};
#endif

# ifdef YYPRINT
/* YYTOKNUM[NUM] -- (External) token number corresponding to the
   (internal) symbol number NUM (which must be that of a token).  */
static const yytype_int16 yytoknum[] =
{
       0,   256,   257,   258,   259,   260,   261,   262,   263,   264,
     265,   266,   267,   268,   269,   270,   271,   272,   273,   274,
     275,   276,   277,   278,   279,   280,   281,   282,   283,   284,
     285,   286,   287,   288,   289,   290,   291,   292,   293,   294,
     295,   296,   297,   298,   299,   300,   301,   302,   303,   304,
     305,   306,   307,   308,   309,   310,   311,   312,   313,   314,
     315,   316,   317,   318,   319,   320,   321,   322,   323,   324,
     325,   326,   327,   328,   329,   330,   331,   332,   333,   334
};
# endif

#define YYPACT_NINF (-47)

#define yypact_value_is_default(Yyn) \
  ((Yyn) == YYPACT_NINF)

#define YYTABLE_NINF (-1)

#define yytable_value_is_error(Yyn) \
  0

  /* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
     STATE-NUM.  */
static const yytype_int16 yypact[] =
{
      68,   114,     3,     9,    15,    38,   -11,    21,    69,    -6,
     -47,    77,    80,   -46,    43,    24,    -4,    56,   132,   -10,
      25,    -2,    -8,    -8,    26,   -16,   -47,   140,   163,   150,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   156,   157,   158,   159,   114,   -47,   160,   161,   162,
     164,   165,   -47,     9,   -47,   166,   167,    15,   -47,   168,
     169,   170,   171,   172,   -11,   -47,   173,   174,   175,   176,
      69,   -47,   177,   178,   179,    -6,   -47,   180,   181,   182,
     183,   184,   -46,   -47,   185,   186,   187,   -47,    24,   -47,
     188,   189,   -47,   190,    -4,   -47,   191,   192,   193,   194,
     -10,   -47,   195,   196,    25,   -47,   197,   198,   -47,    -2,
     -47,   199,   200,    -8,   -47,   -47,    -8,   201,   202,   -47,
      26,   -47,   203,   204,   -16,   -47,   205,   206,   -47,   207,
     208,    98,    60,   -47,    99,    67,   209,   130,   211,   -47,
     102,   124,   -47,   210,   103,    -7,   212,   106,   -47,   107,
     134,   213,   104,   -47,   214,   131,     0,   -47,   133,   135,
     137,     4,   139,   -47,   141,   -25,   218,   -47,   219,   110,
     111,   -47,   115,    10,   116,   220,   -47,   119,   143,   -47,
      20,   217,   -47,    13,    16,   -47,   -47,    17,   145,   -47,
     147,   149,   -47,   221,   222,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   215,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,   -47,
     -47
};

  /* YYDEFACT[STATE-NUM] -- Default reduction number in state STATE-NUM.
     Performed when YYTABLE does not specify something else to do.  Zero
     means the default is an error.  */
static const yytype_uint8 yydefact[] =
{
       0,     0,     0,     0,     0,     0,     0,     0,     0,     0,
      12,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,     0,     0,    28,     0,     0,     0,
       2,     4,     7,     9,    13,    14,    16,    19,    20,    29,
      30,     0,     0,     0,     0,     3,    32,     0,     0,     0,
       0,     0,    42,     5,    41,     0,     0,     6,    50,     0,
       0,     0,     0,     0,     8,    60,     0,     0,     0,     0,
      10,    80,     0,     0,     0,    11,    87,     0,     0,     0,
       0,     0,    15,   100,     0,     0,     0,   116,    17,   112,
       0,     0,   124,     0,    18,   118,     0,     0,     0,     0,
      21,   131,     0,     0,    22,   136,     0,     0,   146,    23,
     142,     0,     0,    24,   148,   156,    25,     0,     0,   164,
      26,   158,     0,     0,    27,   166,     0,     0,     1,     0,
       0,     0,     0,    31,     0,     0,     0,     0,     0,    40,
       0,     0,    49,     0,     0,     0,     0,     0,    59,     0,
       0,     0,     0,    79,     0,     0,     0,    86,     0,     0,
       0,     0,     0,    99,     0,     0,     0,   111,     0,     0,
       0,   117,     0,     0,     0,     0,   130,     0,     0,   135,
       0,     0,   141,     0,     0,   147,   155,     0,     0,   157,
       0,     0,   165,     0,     0,    33,    35,    36,    37,    56,
      57,    58,    34,    39,    38,    47,    46,    43,    44,    45,
      48,    51,    52,    54,    53,    55,    72,    73,    61,    62,
      63,    64,    65,    66,    67,    68,    69,    70,    71,    74,
      76,    75,    77,    78,    82,    83,    81,    84,    85,    94,
      89,    88,    90,    93,    91,    92,    96,    95,    98,    97,
     102,   101,   103,   106,   104,   105,   108,   107,   110,   109,
     113,   114,   115,   121,   120,   119,   123,   122,   126,   125,
     128,   127,   129,   133,   132,   134,   138,   137,   140,     0,
     144,   143,   145,   150,   149,   151,   153,   152,   154,   160,
     159,   161,   163,   162,   168,   167,   170,   169,   171,   172,
     139
};

  /* YYPGOTO[NTERM-NUM].  */
static const yytype_int16 yypgoto[] =
{
     -47,   -47,   -47,   -47,     8,   -47,   -47,   136,   -47,   216,
     -47,   -47,   -47,   223,   -47,   -47,   224,   -47,   225,   -47,
     -47,   -47,   151,   -47,   -47,   146,   -47,   142,   -47,   -47,
     -47,   154,   -47,   152,   -47,   148,   -47,   -23,   -47,   -47,
     120,   -47,   125,   -47,   -47
};

  /* YYDEFGOTO[NTERM-NUM].  */
static const yytype_int16 yydefgoto[] =
{
      -1,    29,    30,    45,    52,    31,    53,    54,    57,    58,
      32,   202,    64,    65,    33,    70,    71,    75,    76,    34,
      35,    82,    83,    36,    88,    89,    94,    95,    37,    38,
     100,   101,   104,   105,   109,   110,   113,   114,   116,   120,
     121,   124,   125,    39,    40
};

  /* YYTABLE[YYPACT[STATE-NUM]] -- What to do in state STATE-NUM.  If
     positive, shift that token.  If negative, reduce the rule whose
     number is the opposite.  If YYTABLE_NINF, syntax error.  */
static const yytype_int16 yytable[] =
{
     115,    60,    98,    61,   111,   260,    90,    62,    91,    46,
     106,    79,    72,    80,    81,    47,   107,   218,   219,    41,
      42,    43,    44,    48,    49,    50,    51,    55,   242,    56,
     243,   244,   252,    66,   253,   254,    85,   102,   117,   103,
     270,   122,   271,   283,   123,   284,   286,   289,   287,   290,
     280,    73,    74,   133,   261,    99,    59,    63,    92,   112,
      93,   108,   220,   221,   222,   223,   224,   225,   226,   227,
     228,     1,     2,     3,     4,     5,     6,     7,   245,   199,
     200,   201,   255,   118,    67,    86,    87,    68,   119,   272,
     185,    77,   285,   186,    78,   288,   291,   205,   206,   281,
      84,     8,     9,    10,    11,    12,    13,    14,    15,    16,
      17,    18,    19,    20,    21,    22,    23,    24,    25,    26,
      27,    28,    96,    69,    41,    42,    43,    44,   197,   203,
     198,   204,   211,   216,   212,   217,   230,   232,   231,   233,
     264,   266,   265,   267,    97,   268,   273,   269,   274,   276,
     128,   277,   208,   209,   213,   214,   234,   235,   126,   237,
     238,   240,   241,   246,   247,   248,   249,   250,   251,   256,
     257,   258,   259,   278,   279,   292,   293,   294,   295,   296,
     297,   127,   129,   130,   131,   132,   134,   135,   136,   139,
     137,   138,   140,   141,   143,   144,   145,   146,   147,   149,
     150,   151,   152,   154,   155,   156,   158,   159,   160,   161,
     162,   164,   165,   166,   168,   169,   170,   172,   173,   174,
     175,   177,   178,   180,   181,   183,   184,   187,   188,   190,
     191,   193,   194,   163,   167,   195,   171,   196,   207,   210,
     189,   215,   300,   229,   236,   239,   262,   263,   282,   192,
     275,     0,   298,   299,   176,     0,   179,   182,     0,     0,
       0,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,     0,     0,   142,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,     0,     0,     0,   148,     0,     0,
       0,     0,     0,     0,   153,     0,     0,     0,     0,     0,
     157
};

static const yytype_int8 yycheck[] =
{
      23,    12,    12,    14,    12,    30,    10,    18,    12,     1,
      12,    57,    18,    59,    60,    12,    18,    24,    25,    10,
      11,    12,    13,    14,    15,    16,    17,    12,    28,    14,
      30,    31,    28,    12,    30,    31,    12,    12,    12,    14,
      30,    57,    32,    30,    60,    32,    30,    30,    32,    32,
      30,    57,    58,    45,    79,    65,    18,    68,    62,    67,
      64,    63,    69,    70,    71,    72,    73,    74,    75,    76,
      77,     3,     4,     5,     6,     7,     8,     9,    78,    19,
      20,    21,    78,    57,    15,    61,    62,    18,    62,    79,
     113,    14,    79,   116,    14,    79,    79,    30,    31,    79,
      57,    33,    34,    35,    36,    37,    38,    39,    40,    41,
      42,    43,    44,    45,    46,    47,    48,    49,    50,    51,
      52,    53,    66,    54,    10,    11,    12,    13,    30,    30,
      32,    32,    30,    30,    32,    32,    30,    30,    32,    32,
      30,    30,    32,    32,    12,    30,    30,    32,    32,    30,
       0,    32,    22,    23,    30,    31,    22,    23,    18,    55,
      56,    30,    31,    30,    31,    30,    31,    30,    31,    30,
      31,    30,    31,    30,    31,    30,    31,    30,    31,    30,
      31,    18,    26,    26,    26,    26,    26,    26,    26,    53,
      26,    26,    26,    26,    26,    26,    26,    26,    26,    26,
      26,    26,    26,    26,    26,    26,    26,    26,    26,    26,
      26,    26,    26,    26,    26,    26,    26,    26,    26,    26,
      26,    26,    26,    26,    26,    26,    26,    26,    26,    26,
      26,    26,    26,    82,    88,    28,    94,    29,    29,    28,
     120,    31,    27,    31,    31,    31,    28,    28,    31,   124,
      30,    -1,    31,    31,   100,    -1,   104,   109,    -1,    -1,
      -1,    -1,    -1,    -1,    -1,    -1,    -1,    -1,    -1,    -1,
      -1,    -1,    -1,    57,    -1,    -1,    -1,    -1,    -1,    -1,
      -1,    -1,    -1,    -1,    -1,    -1,    -1,    64,    -1,    -1,
      -1,    -1,    -1,    -1,    70,    -1,    -1,    -1,    -1,    -1,
      75
};

  /* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
     symbol of state STATE-NUM.  */
static const yytype_int8 yystos[] =
{
       0,     3,     4,     5,     6,     7,     8,     9,    33,    34,
      35,    36,    37,    38,    39,    40,    41,    42,    43,    44,
      45,    46,    47,    48,    49,    50,    51,    52,    53,    81,
      82,    85,    90,    94,    99,   100,   103,   108,   109,   123,
     124,    10,    11,    12,    13,    83,    84,    12,    14,    15,
      16,    17,    84,    86,    87,    12,    14,    88,    89,    18,
      12,    14,    18,    68,    92,    93,    12,    15,    18,    54,
      95,    96,    18,    57,    58,    97,    98,    14,    14,    57,
      59,    60,   101,   102,    57,    12,    61,    62,   104,   105,
      10,    12,    62,    64,   106,   107,    66,    12,    12,    65,
     110,   111,    12,    14,   112,   113,    12,    18,    63,   114,
     115,    12,    67,   116,   117,   117,   118,    12,    57,    62,
     119,   120,    57,    60,   121,   122,    18,    18,     0,    26,
      26,    26,    26,    84,    26,    26,    26,    26,    26,    87,
      26,    26,    89,    26,    26,    26,    26,    26,    93,    26,
      26,    26,    26,    96,    26,    26,    26,    98,    26,    26,
      26,    26,    26,   102,    26,    26,    26,   105,    26,    26,
      26,   107,    26,    26,    26,    26,   111,    26,    26,   113,
      26,    26,   115,    26,    26,   117,   117,    26,    26,   120,
      26,    26,   122,    26,    26,    28,    29,    30,    32,    19,
      20,    21,    91,    30,    32,    30,    31,    29,    22,    23,
      28,    30,    32,    30,    31,    31,    30,    32,    24,    25,
      69,    70,    71,    72,    73,    74,    75,    76,    77,    31,
      30,    32,    30,    32,    22,    23,    31,    55,    56,    31,
      30,    31,    28,    30,    31,    78,    30,    31,    30,    31,
      30,    31,    28,    30,    31,    78,    30,    31,    30,    31,
      30,    79,    28,    28,    30,    32,    30,    32,    30,    32,
      30,    32,    79,    30,    32,    30,    30,    32,    30,    31,
      30,    79,    31,    30,    32,    79,    30,    32,    79,    30,
      32,    79,    30,    31,    30,    31,    30,    31,    31,    31,
      27
};

  /* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */
static const yytype_int8 yyr1[] =
{
       0,    80,    81,    82,    82,    82,    82,    82,    82,    82,
      82,    82,    82,    82,    82,    82,    82,    82,    82,    82,
      82,    82,    82,    82,    82,    82,    82,    82,    82,    82,
      82,    83,    83,    84,    84,    84,    84,    84,    85,    85,
      86,    86,    87,    87,    87,    87,    87,    87,    87,    88,
      88,    89,    89,    89,    89,    90,    91,    91,    91,    92,
      92,    93,    93,    93,    93,    93,    93,    93,    93,    93,
      93,    93,    93,    93,    93,    93,    93,    94,    94,    95,
      95,    96,    96,    96,    96,    96,    97,    97,    98,    98,
      98,    98,    98,    98,    98,    99,    99,   100,   100,   101,
     101,   102,   102,   102,   102,   102,   102,   102,   102,   103,
     103,   104,   104,   105,   105,   105,   105,   106,   106,   107,
     107,   107,   107,   107,   107,   108,   108,   109,   109,   109,
     110,   110,   111,   111,   111,   112,   112,   113,   113,   113,
     113,   114,   114,   115,   115,   115,   115,   116,   116,   117,
     117,   117,   117,   117,   117,   118,   118,   119,   119,   120,
     120,   120,   120,   120,   120,   121,   121,   122,   122,   122,
     122,   123,   124
};

  /* YYR2[YYN] -- Number of symbols on the right hand side of rule YYN.  */
static const yytype_int8 yyr2[] =
{
       0,     2,     1,     2,     1,     2,     2,     1,     2,     1,
       2,     2,     1,     1,     1,     2,     1,     2,     2,     1,
       1,     2,     2,     2,     2,     2,     2,     2,     1,     1,
       1,     2,     1,     3,     3,     3,     3,     3,     4,     4,
       2,     1,     1,     3,     3,     3,     3,     3,     3,     2,
       1,     3,     3,     3,     3,     4,     1,     1,     1,     2,
       1,     3,     3,     3,     3,     3,     3,     3,     3,     3,
       3,     3,     3,     3,     3,     3,     3,     4,     4,     2,
       1,     3,     3,     3,     3,     3,     2,     1,     3,     3,
       3,     3,     3,     3,     3,     4,     4,     4,     4,     2,
       1,     3,     3,     3,     3,     3,     3,     3,     3,     4,
       4,     2,     1,     3,     3,     3,     1,     2,     1,     3,
       3,     3,     3,     3,     1,     4,     4,     4,     4,     4,
       2,     1,     3,     3,     3,     2,     1,     3,     3,     4,
       3,     2,     1,     3,     3,     3,     1,     2,     1,     3,
       3,     3,     3,     3,     3,     2,     1,     2,     1,     3,
       3,     3,     3,     3,     1,     2,     1,     3,     3,     3,
       3,     4,     4
};


#define yyerrok         (yyerrstatus = 0)
#define yyclearin       (yychar = YYEMPTY)
#define YYEMPTY         (-2)
#define YYEOF           0

#define YYACCEPT        goto yyacceptlab
#define YYABORT         goto yyabortlab
#define YYERROR         goto yyerrorlab


#define YYRECOVERING()  (!!yyerrstatus)

#define YYBACKUP(Token, Value)                                    \
  do                                                              \
    if (yychar == YYEMPTY)                                        \
      {                                                           \
        yychar = (Token);                                         \
        yylval = (Value);                                         \
        YYPOPSTACK (yylen);                                       \
        yystate = *yyssp;                                         \
        goto yybackup;                                            \
      }                                                           \
    else                                                          \
      {                                                           \
        yyerror (YY_("syntax error: cannot back up")); \
        YYERROR;                                                  \
      }                                                           \
  while (0)

/* Error token number */
#define YYTERROR        1
#define YYERRCODE       256



/* Enable debugging if requested.  */
#if YYDEBUG

# ifndef YYFPRINTF
#  include <stdio.h> /* INFRINGES ON USER NAME SPACE */
#  define YYFPRINTF fprintf
# endif

# define YYDPRINTF(Args)                        \
do {                                            \
  if (yydebug)                                  \
    YYFPRINTF Args;                             \
} while (0)

/* This macro is provided for backward compatibility. */
#ifndef YY_LOCATION_PRINT
# define YY_LOCATION_PRINT(File, Loc) ((void) 0)
#endif


# define YY_SYMBOL_PRINT(Title, Type, Value, Location)                    \
do {                                                                      \
  if (yydebug)                                                            \
    {                                                                     \
      YYFPRINTF (stderr, "%s ", Title);                                   \
      yy_symbol_print (stderr,                                            \
                  Type, Value); \
      YYFPRINTF (stderr, "\n");                                           \
    }                                                                     \
} while (0)


/*-----------------------------------.
| Print this symbol's value on YYO.  |
`-----------------------------------*/

static void
yy_symbol_value_print (FILE *yyo, int yytype, YYSTYPE const * const yyvaluep)
{
  FILE *yyoutput = yyo;
  YYUSE (yyoutput);
  if (!yyvaluep)
    return;
# ifdef YYPRINT
  if (yytype < YYNTOKENS)
    YYPRINT (yyo, yytoknum[yytype], *yyvaluep);
# endif
  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  YYUSE (yytype);
  YY_IGNORE_MAYBE_UNINITIALIZED_END
}


/*---------------------------.
| Print this symbol on YYO.  |
`---------------------------*/

static void
yy_symbol_print (FILE *yyo, int yytype, YYSTYPE const * const yyvaluep)
{
  YYFPRINTF (yyo, "%s %s (",
             yytype < YYNTOKENS ? "token" : "nterm", yytname[yytype]);

  yy_symbol_value_print (yyo, yytype, yyvaluep);
  YYFPRINTF (yyo, ")");
}

/*------------------------------------------------------------------.
| yy_stack_print -- Print the state stack from its BOTTOM up to its |
| TOP (included).                                                   |
`------------------------------------------------------------------*/

static void
yy_stack_print (yy_state_t *yybottom, yy_state_t *yytop)
{
  YYFPRINTF (stderr, "Stack now");
  for (; yybottom <= yytop; yybottom++)
    {
      int yybot = *yybottom;
      YYFPRINTF (stderr, " %d", yybot);
    }
  YYFPRINTF (stderr, "\n");
}

# define YY_STACK_PRINT(Bottom, Top)                            \
do {                                                            \
  if (yydebug)                                                  \
    yy_stack_print ((Bottom), (Top));                           \
} while (0)


/*------------------------------------------------.
| Report that the YYRULE is going to be reduced.  |
`------------------------------------------------*/

static void
yy_reduce_print (yy_state_t *yyssp, YYSTYPE *yyvsp, int yyrule)
{
  int yylno = yyrline[yyrule];
  int yynrhs = yyr2[yyrule];
  int yyi;
  YYFPRINTF (stderr, "Reducing stack by rule %d (line %d):\n",
             yyrule - 1, yylno);
  /* The symbols being reduced.  */
  for (yyi = 0; yyi < yynrhs; yyi++)
    {
      YYFPRINTF (stderr, "   $%d = ", yyi + 1);
      yy_symbol_print (stderr,
                       yystos[+yyssp[yyi + 1 - yynrhs]],
                       &yyvsp[(yyi + 1) - (yynrhs)]
                                              );
      YYFPRINTF (stderr, "\n");
    }
}

# define YY_REDUCE_PRINT(Rule)          \
do {                                    \
  if (yydebug)                          \
    yy_reduce_print (yyssp, yyvsp, Rule); \
} while (0)

/* Nonzero means print parse trace.  It is left uninitialized so that
   multiple parsers can coexist.  */
int yydebug;
#else /* !YYDEBUG */
# define YYDPRINTF(Args)
# define YY_SYMBOL_PRINT(Title, Type, Value, Location)
# define YY_STACK_PRINT(Bottom, Top)
# define YY_REDUCE_PRINT(Rule)
#endif /* !YYDEBUG */


/* YYINITDEPTH -- initial size of the parser's stacks.  */
#ifndef YYINITDEPTH
# define YYINITDEPTH 200
#endif

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

#ifndef YYMAXDEPTH
# define YYMAXDEPTH 10000
#endif


#if YYERROR_VERBOSE

# ifndef yystrlen
#  if defined __GLIBC__ && defined _STRING_H
#   define yystrlen(S) (YY_CAST (YYPTRDIFF_T, strlen (S)))
#  else
/* Return the length of YYSTR.  */
static YYPTRDIFF_T
yystrlen (const char *yystr)
{
  YYPTRDIFF_T yylen;
  for (yylen = 0; yystr[yylen]; yylen++)
    continue;
  return yylen;
}
#  endif
# endif

# ifndef yystpcpy
#  if defined __GLIBC__ && defined _STRING_H && defined _GNU_SOURCE
#   define yystpcpy stpcpy
#  else
/* Copy YYSRC to YYDEST, returning the address of the terminating '\0' in
   YYDEST.  */
static char *
yystpcpy (char *yydest, const char *yysrc)
{
  char *yyd = yydest;
  const char *yys = yysrc;

  while ((*yyd++ = *yys++) != '\0')
    continue;

  return yyd - 1;
}
#  endif
# endif

# ifndef yytnamerr
/* Copy to YYRES the contents of YYSTR after stripping away unnecessary
   quotes and backslashes, so that it's suitable for yyerror.  The
   heuristic is that double-quoting is unnecessary unless the string
   contains an apostrophe, a comma, or backslash (other than
   backslash-backslash).  YYSTR is taken from yytname.  If YYRES is
   null, do not copy; instead, return the length of what the result
   would have been.  */
static YYPTRDIFF_T
yytnamerr (char *yyres, const char *yystr)
{
  if (*yystr == '"')
    {
      YYPTRDIFF_T yyn = 0;
      char const *yyp = yystr;

      for (;;)
        switch (*++yyp)
          {
          case '\'':
          case ',':
            goto do_not_strip_quotes;

          case '\\':
            if (*++yyp != '\\')
              goto do_not_strip_quotes;
            else
              goto append;

          append:
          default:
            if (yyres)
              yyres[yyn] = *yyp;
            yyn++;
            break;

          case '"':
            if (yyres)
              yyres[yyn] = '\0';
            return yyn;
          }
    do_not_strip_quotes: ;
    }

  if (yyres)
    return yystpcpy (yyres, yystr) - yyres;
  else
    return yystrlen (yystr);
}
# endif

/* Copy into *YYMSG, which is of size *YYMSG_ALLOC, an error message
   about the unexpected token YYTOKEN for the state stack whose top is
   YYSSP.

   Return 0 if *YYMSG was successfully written.  Return 1 if *YYMSG is
   not large enough to hold the message.  In that case, also set
   *YYMSG_ALLOC to the required number of bytes.  Return 2 if the
   required number of bytes is too large to store.  */
static int
yysyntax_error (YYPTRDIFF_T *yymsg_alloc, char **yymsg,
                yy_state_t *yyssp, int yytoken)
{
  enum { YYERROR_VERBOSE_ARGS_MAXIMUM = 5 };
  /* Internationalized format string. */
  const char *yyformat = YY_NULLPTR;
  /* Arguments of yyformat: reported tokens (one for the "unexpected",
     one per "expected"). */
  char const *yyarg[YYERROR_VERBOSE_ARGS_MAXIMUM];
  /* Actual size of YYARG. */
  int yycount = 0;
  /* Cumulated lengths of YYARG.  */
  YYPTRDIFF_T yysize = 0;

  /* There are many possibilities here to consider:
     - If this state is a consistent state with a default action, then
       the only way this function was invoked is if the default action
       is an error action.  In that case, don't check for expected
       tokens because there are none.
     - The only way there can be no lookahead present (in yychar) is if
       this state is a consistent state with a default action.  Thus,
       detecting the absence of a lookahead is sufficient to determine
       that there is no unexpected or expected token to report.  In that
       case, just report a simple "syntax error".
     - Don't assume there isn't a lookahead just because this state is a
       consistent state with a default action.  There might have been a
       previous inconsistent state, consistent state with a non-default
       action, or user semantic action that manipulated yychar.
     - Of course, the expected token list depends on states to have
       correct lookahead information, and it depends on the parser not
       to perform extra reductions after fetching a lookahead from the
       scanner and before detecting a syntax error.  Thus, state merging
       (from LALR or IELR) and default reductions corrupt the expected
       token list.  However, the list is correct for canonical LR with
       one exception: it will still contain any token that will not be
       accepted due to an error action in a later state.
  */
  if (yytoken != YYEMPTY)
    {
      int yyn = yypact[+*yyssp];
      YYPTRDIFF_T yysize0 = yytnamerr (YY_NULLPTR, yytname[yytoken]);
      yysize = yysize0;
      yyarg[yycount++] = yytname[yytoken];
      if (!yypact_value_is_default (yyn))
        {
          /* Start YYX at -YYN if negative to avoid negative indexes in
             YYCHECK.  In other words, skip the first -YYN actions for
             this state because they are default actions.  */
          int yyxbegin = yyn < 0 ? -yyn : 0;
          /* Stay within bounds of both yycheck and yytname.  */
          int yychecklim = YYLAST - yyn + 1;
          int yyxend = yychecklim < YYNTOKENS ? yychecklim : YYNTOKENS;
          int yyx;

          for (yyx = yyxbegin; yyx < yyxend; ++yyx)
            if (yycheck[yyx + yyn] == yyx && yyx != YYTERROR
                && !yytable_value_is_error (yytable[yyx + yyn]))
              {
                if (yycount == YYERROR_VERBOSE_ARGS_MAXIMUM)
                  {
                    yycount = 1;
                    yysize = yysize0;
                    break;
                  }
                yyarg[yycount++] = yytname[yyx];
                {
                  YYPTRDIFF_T yysize1
                    = yysize + yytnamerr (YY_NULLPTR, yytname[yyx]);
                  if (yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM)
                    yysize = yysize1;
                  else
                    return 2;
                }
              }
        }
    }

  switch (yycount)
    {
# define YYCASE_(N, S)                      \
      case N:                               \
        yyformat = S;                       \
      break
    default: /* Avoid compiler warnings. */
      YYCASE_(0, YY_("syntax error"));
      YYCASE_(1, YY_("syntax error, unexpected %s"));
      YYCASE_(2, YY_("syntax error, unexpected %s, expecting %s"));
      YYCASE_(3, YY_("syntax error, unexpected %s, expecting %s or %s"));
      YYCASE_(4, YY_("syntax error, unexpected %s, expecting %s or %s or %s"));
      YYCASE_(5, YY_("syntax error, unexpected %s, expecting %s or %s or %s or %s"));
# undef YYCASE_
    }

  {
    /* Don't count the "%s"s in the final size, but reserve room for
       the terminator.  */
    YYPTRDIFF_T yysize1 = yysize + (yystrlen (yyformat) - 2 * yycount) + 1;
    if (yysize <= yysize1 && yysize1 <= YYSTACK_ALLOC_MAXIMUM)
      yysize = yysize1;
    else
      return 2;
  }

  if (*yymsg_alloc < yysize)
    {
      *yymsg_alloc = 2 * yysize;
      if (! (yysize <= *yymsg_alloc
             && *yymsg_alloc <= YYSTACK_ALLOC_MAXIMUM))
        *yymsg_alloc = YYSTACK_ALLOC_MAXIMUM;
      return 1;
    }

  /* Avoid sprintf, as that infringes on the user's name space.
     Don't have undefined behavior even if the translation
     produced a string with the wrong number of "%s"s.  */
  {
    char *yyp = *yymsg;
    int yyi = 0;
    while ((*yyp = *yyformat) != '\0')
      if (*yyp == '%' && yyformat[1] == 's' && yyi < yycount)
        {
          yyp += yytnamerr (yyp, yyarg[yyi++]);
          yyformat += 2;
        }
      else
        {
          ++yyp;
          ++yyformat;
        }
  }
  return 0;
}
#endif /* YYERROR_VERBOSE */

/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

static void
yydestruct (const char *yymsg, int yytype, YYSTYPE *yyvaluep)
{
  YYUSE (yyvaluep);
  if (!yymsg)
    yymsg = "Deleting";
  YY_SYMBOL_PRINT (yymsg, yytype, yyvaluep, yylocationp);

  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  YYUSE (yytype);
  YY_IGNORE_MAYBE_UNINITIALIZED_END
}




/* The lookahead symbol.  */
int yychar;

/* The semantic value of the lookahead symbol.  */
YYSTYPE yylval;
/* Number of syntax errors so far.  */
int yynerrs;


/*----------.
| yyparse.  |
`----------*/

int
yyparse (void)
{
    yy_state_fast_t yystate;
    /* Number of tokens to shift before error messages enabled.  */
    int yyerrstatus;

    /* The stacks and their tools:
       'yyss': related to states.
       'yyvs': related to semantic values.

       Refer to the stacks through separate pointers, to allow yyoverflow
       to reallocate them elsewhere.  */

    /* The state stack.  */
    yy_state_t yyssa[YYINITDEPTH];
    yy_state_t *yyss;
    yy_state_t *yyssp;

    /* The semantic value stack.  */
    YYSTYPE yyvsa[YYINITDEPTH];
    YYSTYPE *yyvs;
    YYSTYPE *yyvsp;

    YYPTRDIFF_T yystacksize;

  int yyn;
  int yyresult;
  /* Lookahead token as an internal (translated) token number.  */
  int yytoken = 0;
  /* The variables used to return semantic value and location from the
     action routines.  */
  YYSTYPE yyval;

#if YYERROR_VERBOSE
  /* Buffer for error messages, and its allocated size.  */
  char yymsgbuf[128];
  char *yymsg = yymsgbuf;
  YYPTRDIFF_T yymsg_alloc = sizeof yymsgbuf;
#endif

#define YYPOPSTACK(N)   (yyvsp -= (N), yyssp -= (N))

  /* The number of symbols on the RHS of the reduced rule.
     Keep to zero when no symbol should be popped.  */
  int yylen = 0;

  yyssp = yyss = yyssa;
  yyvsp = yyvs = yyvsa;
  yystacksize = YYINITDEPTH;

  YYDPRINTF ((stderr, "Starting parse\n"));

  yystate = 0;
  yyerrstatus = 0;
  yynerrs = 0;
  yychar = YYEMPTY; /* Cause a token to be read.  */
  goto yysetstate;


/*------------------------------------------------------------.
| yynewstate -- push a new state, which is found in yystate.  |
`------------------------------------------------------------*/
yynewstate:
  /* In all cases, when you get here, the value and location stacks
     have just been pushed.  So pushing a state here evens the stacks.  */
  yyssp++;


/*--------------------------------------------------------------------.
| yysetstate -- set current state (the top of the stack) to yystate.  |
`--------------------------------------------------------------------*/
yysetstate:
  YYDPRINTF ((stderr, "Entering state %d\n", yystate));
  YY_ASSERT (0 <= yystate && yystate < YYNSTATES);
  YY_IGNORE_USELESS_CAST_BEGIN
  *yyssp = YY_CAST (yy_state_t, yystate);
  YY_IGNORE_USELESS_CAST_END

  if (yyss + yystacksize - 1 <= yyssp)
#if !defined yyoverflow && !defined YYSTACK_RELOCATE
    goto yyexhaustedlab;
#else
    {
      /* Get the current used size of the three stacks, in elements.  */
      YYPTRDIFF_T yysize = yyssp - yyss + 1;

# if defined yyoverflow
      {
        /* Give user a chance to reallocate the stack.  Use copies of
           these so that the &'s don't force the real ones into
           memory.  */
        yy_state_t *yyss1 = yyss;
        YYSTYPE *yyvs1 = yyvs;

        /* Each stack pointer address is followed by the size of the
           data in use in that stack, in bytes.  This used to be a
           conditional around just the two extra args, but that might
           be undefined if yyoverflow is a macro.  */
        yyoverflow (YY_("memory exhausted"),
                    &yyss1, yysize * YYSIZEOF (*yyssp),
                    &yyvs1, yysize * YYSIZEOF (*yyvsp),
                    &yystacksize);
        yyss = yyss1;
        yyvs = yyvs1;
      }
# else /* defined YYSTACK_RELOCATE */
      /* Extend the stack our own way.  */
      if (YYMAXDEPTH <= yystacksize)
        goto yyexhaustedlab;
      yystacksize *= 2;
      if (YYMAXDEPTH < yystacksize)
        yystacksize = YYMAXDEPTH;

      {
        yy_state_t *yyss1 = yyss;
        union yyalloc *yyptr =
          YY_CAST (union yyalloc *,
                   YYSTACK_ALLOC (YY_CAST (YYSIZE_T, YYSTACK_BYTES (yystacksize))));
        if (! yyptr)
          goto yyexhaustedlab;
        YYSTACK_RELOCATE (yyss_alloc, yyss);
        YYSTACK_RELOCATE (yyvs_alloc, yyvs);
# undef YYSTACK_RELOCATE
        if (yyss1 != yyssa)
          YYSTACK_FREE (yyss1);
      }
# endif

      yyssp = yyss + yysize - 1;
      yyvsp = yyvs + yysize - 1;

      YY_IGNORE_USELESS_CAST_BEGIN
      YYDPRINTF ((stderr, "Stack size increased to %ld\n",
                  YY_CAST (long, yystacksize)));
      YY_IGNORE_USELESS_CAST_END

      if (yyss + yystacksize - 1 <= yyssp)
        YYABORT;
    }
#endif /* !defined yyoverflow && !defined YYSTACK_RELOCATE */

  if (yystate == YYFINAL)
    YYACCEPT;

  goto yybackup;


/*-----------.
| yybackup.  |
`-----------*/
yybackup:
  /* Do appropriate processing given the current state.  Read a
     lookahead token if we need one and don't already have one.  */

  /* First try to decide what to do without reference to lookahead token.  */
  yyn = yypact[yystate];
  if (yypact_value_is_default (yyn))
    goto yydefault;

  /* Not known => get a lookahead token if don't already have one.  */

  /* YYCHAR is either YYEMPTY or YYEOF or a valid lookahead symbol.  */
  if (yychar == YYEMPTY)
    {
      YYDPRINTF ((stderr, "Reading a token: "));
      yychar = yylex ();
    }

  if (yychar <= YYEOF)
    {
      yychar = yytoken = YYEOF;
      YYDPRINTF ((stderr, "Now at end of input.\n"));
    }
  else
    {
      yytoken = YYTRANSLATE (yychar);
      YY_SYMBOL_PRINT ("Next token is", yytoken, &yylval, &yylloc);
    }

  /* If the proper action on seeing token YYTOKEN is to reduce or to
     detect an error, take that action.  */
  yyn += yytoken;
  if (yyn < 0 || YYLAST < yyn || yycheck[yyn] != yytoken)
    goto yydefault;
  yyn = yytable[yyn];
  if (yyn <= 0)
    {
      if (yytable_value_is_error (yyn))
        goto yyerrlab;
      yyn = -yyn;
      goto yyreduce;
    }

  /* Count tokens shifted since error; after three, turn off error
     status.  */
  if (yyerrstatus)
    yyerrstatus--;

  /* Shift the lookahead token.  */
  YY_SYMBOL_PRINT ("Shifting", yytoken, &yylval, &yylloc);
  yystate = yyn;
  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++yyvsp = yylval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END

  /* Discard the shifted token.  */
  yychar = YYEMPTY;
  goto yynewstate;


/*-----------------------------------------------------------.
| yydefault -- do the default action for the current state.  |
`-----------------------------------------------------------*/
yydefault:
  yyn = yydefact[yystate];
  if (yyn == 0)
    goto yyerrlab;
  goto yyreduce;


/*-----------------------------.
| yyreduce -- do a reduction.  |
`-----------------------------*/
yyreduce:
  /* yyn is the number of a rule to reduce with.  */
  yylen = yyr2[yyn];

  /* If YYLEN is nonzero, implement the default value of the action:
     '$$ = $1'.

     Otherwise, the following line sets YYVAL to garbage.
     This behavior is undocumented and Bison
     users should not rely upon it.  Assigning to YYVAL
     unconditionally makes the parser a bit smaller, and it avoids a
     GCC warning that YYVAL may be used uninitialized.  */
  yyval = yyvsp[1-yylen];


  YY_REDUCE_PRINT (yyn);
  switch (yyn)
    {
  case 2:
#line 151 "Sintactico.y"
               { raiz = (yyval.nodito); }
#line 1627 "parser.cpp"
    break;

  case 3:
#line 153 "Sintactico.y"
                       {
                         (yyval.nodito) = new Nodo("MKDISK","");
                         (yyval.nodito)->add(*(yyvsp[0].nodito));
                       }
#line 1636 "parser.cpp"
    break;

  case 4:
#line 157 "Sintactico.y"
                  { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1642 "parser.cpp"
    break;

  case 5:
#line 158 "Sintactico.y"
                       {
                          (yyval.nodito) = new Nodo("FDISK","");
                          (yyval.nodito)->add(*(yyvsp[0].nodito));
                        }
#line 1651 "parser.cpp"
    break;

  case 6:
#line 162 "Sintactico.y"
                       {
                         (yyval.nodito) = new Nodo("MOUNT", "");
                         (yyval.nodito)->add(*(yyvsp[0].nodito));
                       }
#line 1660 "parser.cpp"
    break;

  case 7:
#line 166 "Sintactico.y"
                   { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1666 "parser.cpp"
    break;

  case 8:
#line 167 "Sintactico.y"
                   { (yyval.nodito) = new Nodo("REP","");
                     (yyval.nodito)->add(*(yyvsp[0].nodito));
                   }
#line 1674 "parser.cpp"
    break;

  case 9:
#line 170 "Sintactico.y"
                  { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1680 "parser.cpp"
    break;

  case 10:
#line 171 "Sintactico.y"
                     {
                        (yyval.nodito) = new Nodo("MKFS","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 1689 "parser.cpp"
    break;

  case 11:
#line 175 "Sintactico.y"
                       {
                         (yyval.nodito) = new Nodo("LOGIN", "");
                         (yyval.nodito)->add(*(yyvsp[0].nodito));
                       }
#line 1698 "parser.cpp"
    break;

  case 12:
#line 179 "Sintactico.y"
                  { (yyval.nodito) = new Nodo("LOGOUT",""); }
#line 1704 "parser.cpp"
    break;

  case 13:
#line 180 "Sintactico.y"
                 { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1710 "parser.cpp"
    break;

  case 14:
#line 181 "Sintactico.y"
                 { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1716 "parser.cpp"
    break;

  case 15:
#line 182 "Sintactico.y"
                       {
                        (yyval.nodito) = new Nodo("MKUSR","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                       }
#line 1725 "parser.cpp"
    break;

  case 16:
#line 186 "Sintactico.y"
                 { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1731 "parser.cpp"
    break;

  case 17:
#line 187 "Sintactico.y"
                        {
                         (yyval.nodito) = new Nodo("CHMOD","");
                         (yyval.nodito)->add(*(yyvsp[0].nodito));
                       }
#line 1740 "parser.cpp"
    break;

  case 18:
#line 191 "Sintactico.y"
                         {
                            (yyval.nodito) = new Nodo("MKFILE","");
                            (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 1749 "parser.cpp"
    break;

  case 19:
#line 195 "Sintactico.y"
               { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1755 "parser.cpp"
    break;

  case 20:
#line 196 "Sintactico.y"
               { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1761 "parser.cpp"
    break;

  case 21:
#line 197 "Sintactico.y"
                     {
                       (yyval.nodito) = new Nodo("EDIT","");
                       (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 1770 "parser.cpp"
    break;

  case 22:
#line 201 "Sintactico.y"
                  {
                    (yyval.nodito) = new Nodo("REN","");
                    (yyval.nodito)->add(*(yyvsp[0].nodito));
                  }
#line 1779 "parser.cpp"
    break;

  case 23:
#line 205 "Sintactico.y"
                       {
                        (yyval.nodito) = new Nodo("MKDIR","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 1788 "parser.cpp"
    break;

  case 24:
#line 209 "Sintactico.y"
                {
                  (yyval.nodito) = new Nodo("CP","");
                  (yyval.nodito)->add(*(yyvsp[0].nodito));
                }
#line 1797 "parser.cpp"
    break;

  case 25:
#line 213 "Sintactico.y"
                {
                  (yyval.nodito) = new Nodo("MV","");
                  (yyval.nodito)->add(*(yyvsp[0].nodito));
                }
#line 1806 "parser.cpp"
    break;

  case 26:
#line 217 "Sintactico.y"
                       {
                        (yyval.nodito) = new Nodo("CHOWN","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 1815 "parser.cpp"
    break;

  case 27:
#line 221 "Sintactico.y"
                      {
                        (yyval.nodito) = new Nodo("CHGRP","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 1824 "parser.cpp"
    break;

  case 28:
#line 225 "Sintactico.y"
                 { (yyval.nodito) = new Nodo("PAUSE",""); }
#line 1830 "parser.cpp"
    break;

  case 29:
#line 226 "Sintactico.y"
                    { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1836 "parser.cpp"
    break;

  case 30:
#line 227 "Sintactico.y"
                { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1842 "parser.cpp"
    break;

  case 31:
#line 232 "Sintactico.y"
                            {
                             (yyval.nodito) = (yyvsp[-1].nodito);
                             (yyval.nodito)->add(*(yyvsp[0].nodito));
                            }
#line 1851 "parser.cpp"
    break;

  case 32:
#line 236 "Sintactico.y"
                      {
                       (yyval.nodito) = new Nodo("PARAMETRO","");
                       (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 1860 "parser.cpp"
    break;

  case 33:
#line 241 "Sintactico.y"
                             { (yyval.nodito)= new Nodo("size",(yyvsp[0].text)); }
#line 1866 "parser.cpp"
    break;

  case 34:
#line 242 "Sintactico.y"
                              {
                                (yyval.nodito) = new Nodo ("fit", "");
                                (yyval.nodito)->add(*(yyvsp[0].nodito));
                               }
#line 1875 "parser.cpp"
    break;

  case 35:
#line 246 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("unit",(yyvsp[0].text)); }
#line 1881 "parser.cpp"
    break;

  case 36:
#line 247 "Sintactico.y"
                               {
                                 (yyval.nodito) = new Nodo("path",(yyvsp[0].text));
                                }
#line 1889 "parser.cpp"
    break;

  case 37:
#line 250 "Sintactico.y"
                             {
                               (yyval.nodito) = new Nodo("path", (yyvsp[0].text));

                             }
#line 1898 "parser.cpp"
    break;

  case 38:
#line 255 "Sintactico.y"
                               {
                                (yyval.nodito) = new Nodo("RMDISK","");
                                Nodo *n = new Nodo("path",(yyvsp[0].text));
                                (yyval.nodito)->add(*n);
                               }
#line 1908 "parser.cpp"
    break;

  case 39:
#line 260 "Sintactico.y"
                                    {
                                      (yyval.nodito) = new Nodo("RMDISK","");
                                      Nodo *ruta = new Nodo("path",(yyvsp[0].text));
                                      (yyval.nodito)->add(*ruta);
                                    }
#line 1918 "parser.cpp"
    break;

  case 40:
#line 265 "Sintactico.y"
                          {
                            (yyval.nodito) = (yyvsp[-1].nodito);
                            (yyval.nodito)->add(*(yyvsp[0].nodito));
                          }
#line 1927 "parser.cpp"
    break;

  case 41:
#line 269 "Sintactico.y"
                       {
                        (yyval.nodito) = new Nodo("PARAMETRO","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 1936 "parser.cpp"
    break;

  case 42:
#line 274 "Sintactico.y"
                           { (yyval.nodito) = (yyvsp[0].nodito); }
#line 1942 "parser.cpp"
    break;

  case 43:
#line 275 "Sintactico.y"
                                    { (yyval.nodito) = new Nodo("type",(yyvsp[0].text)); }
#line 1948 "parser.cpp"
    break;

  case 44:
#line 276 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("delete", "fast"); }
#line 1954 "parser.cpp"
    break;

  case 45:
#line 277 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("delete", "full"); }
#line 1960 "parser.cpp"
    break;

  case 46:
#line 278 "Sintactico.y"
                                         { (yyval.nodito) = new Nodo("name", (yyvsp[0].text)); }
#line 1966 "parser.cpp"
    break;

  case 47:
#line 279 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("name", (yyvsp[0].text)); }
#line 1972 "parser.cpp"
    break;

  case 48:
#line 280 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("add", (yyvsp[0].text)); }
#line 1978 "parser.cpp"
    break;

  case 49:
#line 282 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 1987 "parser.cpp"
    break;

  case 50:
#line 286 "Sintactico.y"
                     {
                        (yyval.nodito) = new Nodo("PARAMETRO","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 1996 "parser.cpp"
    break;

  case 51:
#line 291 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2002 "parser.cpp"
    break;

  case 52:
#line 292 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("path", (yyvsp[0].text)); }
#line 2008 "parser.cpp"
    break;

  case 53:
#line 293 "Sintactico.y"
                                        { (yyval.nodito) = new Nodo("name", (yyvsp[0].text)); }
#line 2014 "parser.cpp"
    break;

  case 54:
#line 294 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("name",(yyvsp[0].text)); }
#line 2020 "parser.cpp"
    break;

  case 55:
#line 296 "Sintactico.y"
                                        {
                                          (yyval.nodito) = new Nodo("UNMOUNT", "");
                                          Nodo *n = new Nodo("id",(yyvsp[0].text));
                                          (yyval.nodito)->add(*n);
                                        }
#line 2030 "parser.cpp"
    break;

  case 56:
#line 302 "Sintactico.y"
           { (yyval.nodito) = new Nodo("ajuste", "bf"); }
#line 2036 "parser.cpp"
    break;

  case 57:
#line 303 "Sintactico.y"
             { (yyval.nodito) = new Nodo("ajuste", "ff"); }
#line 2042 "parser.cpp"
    break;

  case 58:
#line 304 "Sintactico.y"
             { (yyval.nodito) = new Nodo("ajuste", "wf"); }
#line 2048 "parser.cpp"
    break;

  case 59:
#line 307 "Sintactico.y"
                    {
                     (yyval.nodito) = (yyvsp[-1].nodito);
                     (yyval.nodito)->add(*(yyvsp[0].nodito));
                    }
#line 2057 "parser.cpp"
    break;

  case 60:
#line 311 "Sintactico.y"
                  {
                    (yyval.nodito) = new Nodo("PARAMETRO", "");
                    (yyval.nodito)->add(*(yyvsp[0].nodito));
                  }
#line 2066 "parser.cpp"
    break;

  case 61:
#line 316 "Sintactico.y"
                            { (yyval.nodito) = new Nodo("name","mbr"); }
#line 2072 "parser.cpp"
    break;

  case 62:
#line 317 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("name","disk"); }
#line 2078 "parser.cpp"
    break;

  case 63:
#line 318 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("name", "inode"); }
#line 2084 "parser.cpp"
    break;

  case 64:
#line 319 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("name", "journaling"); }
#line 2090 "parser.cpp"
    break;

  case 65:
#line 320 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("name", "block"); }
#line 2096 "parser.cpp"
    break;

  case 66:
#line 321 "Sintactico.y"
                                   { (yyval.nodito) = new Nodo("name", "bm_inode"); }
#line 2102 "parser.cpp"
    break;

  case 67:
#line 322 "Sintactico.y"
                                   { (yyval.nodito) = new Nodo("name", "bm_block"); }
#line 2108 "parser.cpp"
    break;

  case 68:
#line 323 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("name", "tree"); }
#line 2114 "parser.cpp"
    break;

  case 69:
#line 324 "Sintactico.y"
                             { (yyval.nodito) = new Nodo("name", "sb"); }
#line 2120 "parser.cpp"
    break;

  case 70:
#line 325 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("name", "file"); }
#line 2126 "parser.cpp"
    break;

  case 71:
#line 326 "Sintactico.y"
                             { (yyval.nodito) = new Nodo("name", "ls"); }
#line 2132 "parser.cpp"
    break;

  case 72:
#line 327 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("path", (yyvsp[0].text)); }
#line 2138 "parser.cpp"
    break;

  case 73:
#line 328 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2144 "parser.cpp"
    break;

  case 74:
#line 329 "Sintactico.y"
                                      { (yyval.nodito) = new Nodo("id", (yyvsp[0].text)); }
#line 2150 "parser.cpp"
    break;

  case 75:
#line 330 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("ruta", (yyvsp[0].text)); }
#line 2156 "parser.cpp"
    break;

  case 76:
#line 331 "Sintactico.y"
                                    { (yyval.nodito) = new Nodo("ruta", (yyvsp[0].text)); }
#line 2162 "parser.cpp"
    break;

  case 77:
#line 333 "Sintactico.y"
                               {
                                (yyval.nodito) = new Nodo("EXEC","");
                                Nodo *path = new Nodo("path", (yyvsp[0].text));
                                (yyval.nodito)->add(*path);
                               }
#line 2172 "parser.cpp"
    break;

  case 78:
#line 338 "Sintactico.y"
                               {
                                 (yyval.nodito) = new Nodo("EXEC","");
                                 Nodo *n = new Nodo("path", (yyvsp[0].text));
                                 (yyval.nodito)->add(*n);
                               }
#line 2182 "parser.cpp"
    break;

  case 79:
#line 344 "Sintactico.y"
                      {
                        (yyval.nodito) = (yyvsp[-1].nodito);
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 2191 "parser.cpp"
    break;

  case 80:
#line 348 "Sintactico.y"
                   {
                      (yyval.nodito) = new Nodo("PARAMETRO", "");
                      (yyval.nodito)->add(*(yyvsp[0].nodito));
                   }
#line 2200 "parser.cpp"
    break;

  case 81:
#line 353 "Sintactico.y"
                                   { (yyval.nodito) = new Nodo("id",(yyvsp[0].text)); }
#line 2206 "parser.cpp"
    break;

  case 82:
#line 354 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("type", "fast"); }
#line 2212 "parser.cpp"
    break;

  case 83:
#line 355 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("type", "full"); }
#line 2218 "parser.cpp"
    break;

  case 84:
#line 356 "Sintactico.y"
                           { (yyval.nodito) = new Nodo("fs", "2fs"); }
#line 2224 "parser.cpp"
    break;

  case 85:
#line 357 "Sintactico.y"
                           { (yyval.nodito) = new Nodo("fs", "3fs"); }
#line 2230 "parser.cpp"
    break;

  case 86:
#line 359 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2239 "parser.cpp"
    break;

  case 87:
#line 363 "Sintactico.y"
                     {
                        (yyval.nodito) = new Nodo("PARAMETRO","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 2248 "parser.cpp"
    break;

  case 88:
#line 368 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("user", (yyvsp[0].text)); }
#line 2254 "parser.cpp"
    break;

  case 89:
#line 369 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("user", (yyvsp[0].text)); }
#line 2260 "parser.cpp"
    break;

  case 90:
#line 370 "Sintactico.y"
                            { (yyval.nodito) = new Nodo("password",(yyvsp[0].text)); }
#line 2266 "parser.cpp"
    break;

  case 91:
#line 371 "Sintactico.y"
                                      { (yyval.nodito) = new Nodo("password", (yyvsp[0].text)); }
#line 2272 "parser.cpp"
    break;

  case 92:
#line 372 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("password", (yyvsp[0].text)); }
#line 2278 "parser.cpp"
    break;

  case 93:
#line 373 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("password", (yyvsp[0].text)); }
#line 2284 "parser.cpp"
    break;

  case 94:
#line 374 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("id", (yyvsp[0].text)); }
#line 2290 "parser.cpp"
    break;

  case 95:
#line 376 "Sintactico.y"
                                      {
                                        (yyval.nodito) = new Nodo("MKGRP","");
                                        Nodo *n = new Nodo("name",(yyvsp[0].text));
                                        (yyval.nodito)->add(*n);
                                      }
#line 2300 "parser.cpp"
    break;

  case 96:
#line 381 "Sintactico.y"
                                 {
                                    (yyval.nodito) = new Nodo("MKGRP","");
                                    Nodo *n = new Nodo("name",(yyvsp[0].text));
                                    (yyval.nodito)->add(*n);
                                 }
#line 2310 "parser.cpp"
    break;

  case 97:
#line 387 "Sintactico.y"
                                      {
                                        (yyval.nodito) = new Nodo("RMGRP","");
                                        Nodo *n = new Nodo("name", (yyvsp[0].text));
                                        (yyval.nodito)->add(*n);
                                     }
#line 2320 "parser.cpp"
    break;

  case 98:
#line 392 "Sintactico.y"
                                 {
                                    (yyval.nodito) = new Nodo("RMGRP", "");
                                    Nodo *n = new Nodo("name",(yyvsp[0].text));
                                    (yyval.nodito)->add(*n);
                                 }
#line 2330 "parser.cpp"
    break;

  case 99:
#line 397 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2339 "parser.cpp"
    break;

  case 100:
#line 401 "Sintactico.y"
                     {
                        (yyval.nodito) = new Nodo("PARAMETRO", "");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 2348 "parser.cpp"
    break;

  case 101:
#line 406 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2354 "parser.cpp"
    break;

  case 102:
#line 407 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2360 "parser.cpp"
    break;

  case 103:
#line 408 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("password",(yyvsp[0].text)); }
#line 2366 "parser.cpp"
    break;

  case 104:
#line 409 "Sintactico.y"
                                        { (yyval.nodito) = new Nodo("password", (yyvsp[0].text)); }
#line 2372 "parser.cpp"
    break;

  case 105:
#line 410 "Sintactico.y"
                                   { (yyval.nodito) = new Nodo("password", (yyvsp[0].text)); }
#line 2378 "parser.cpp"
    break;

  case 106:
#line 411 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("password",(yyvsp[0].text)); }
#line 2384 "parser.cpp"
    break;

  case 107:
#line 412 "Sintactico.y"
                                       { (yyval.nodito) = new Nodo("group", (yyvsp[0].text)); }
#line 2390 "parser.cpp"
    break;

  case 108:
#line 413 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("group",(yyvsp[0].text)); }
#line 2396 "parser.cpp"
    break;

  case 109:
#line 415 "Sintactico.y"
                                     {
                                        (yyval.nodito) = new Nodo("RMUSR","");
                                        Nodo *n = new Nodo("user",(yyvsp[0].text));
                                        (yyval.nodito)->add(*n);
                                     }
#line 2406 "parser.cpp"
    break;

  case 110:
#line 420 "Sintactico.y"
                               {
                                  (yyval.nodito) = new Nodo("RMUSR", "");
                                  Nodo *n = new Nodo("user",(yyvsp[0].text));
                                  (yyval.nodito)->add(*n);
                               }
#line 2416 "parser.cpp"
    break;

  case 111:
#line 426 "Sintactico.y"
                         {
                          (yyval.nodito) = (yyvsp[-1].nodito);
                          (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2425 "parser.cpp"
    break;

  case 112:
#line 430 "Sintactico.y"
                     {
                       (yyval.nodito) = new Nodo("PARAMETRO","");
                       (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 2434 "parser.cpp"
    break;

  case 113:
#line 435 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2440 "parser.cpp"
    break;

  case 114:
#line 436 "Sintactico.y"
                                    { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2446 "parser.cpp"
    break;

  case 115:
#line 437 "Sintactico.y"
                            { (yyval.nodito) = new Nodo("ugo",(yyvsp[0].text)); }
#line 2452 "parser.cpp"
    break;

  case 116:
#line 438 "Sintactico.y"
                { (yyval.nodito) = new Nodo("r",""); }
#line 2458 "parser.cpp"
    break;

  case 117:
#line 440 "Sintactico.y"
                            {
                              (yyval.nodito) = (yyvsp[-1].nodito);
                              (yyval.nodito)->add(*(yyvsp[0].nodito));
                            }
#line 2467 "parser.cpp"
    break;

  case 118:
#line 444 "Sintactico.y"
                      {
                        (yyval.nodito) = new Nodo("PARAMETRO","");
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 2476 "parser.cpp"
    break;

  case 119:
#line 450 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2482 "parser.cpp"
    break;

  case 120:
#line 451 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2488 "parser.cpp"
    break;

  case 121:
#line 452 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("size",(yyvsp[0].text)); }
#line 2494 "parser.cpp"
    break;

  case 122:
#line 453 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("cont",(yyvsp[0].text)); }
#line 2500 "parser.cpp"
    break;

  case 123:
#line 454 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("cont",(yyvsp[0].text)); }
#line 2506 "parser.cpp"
    break;

  case 124:
#line 455 "Sintactico.y"
                  { (yyval.nodito) = new Nodo("r",""); }
#line 2512 "parser.cpp"
    break;

  case 125:
#line 457 "Sintactico.y"
                        {
                          (yyval.nodito) = new Nodo("CAT","");
                          Nodo *n = new Nodo("file",(yyvsp[0].text));
                          (yyval.nodito)->add(*n);
                        }
#line 2522 "parser.cpp"
    break;

  case 126:
#line 462 "Sintactico.y"
                           {
                             (yyval.nodito) = new Nodo("CAT","");
                             Nodo *n = new Nodo("file",(yyvsp[0].text));
                             (yyval.nodito)->add(*n);
                           }
#line 2532 "parser.cpp"
    break;

  case 127:
#line 468 "Sintactico.y"
                        {
                          (yyval.nodito) = new Nodo("REM","");
                          Nodo *n = new Nodo("path",(yyvsp[0].text));
                          (yyval.nodito)->add(*n);
                        }
#line 2542 "parser.cpp"
    break;

  case 128:
#line 473 "Sintactico.y"
                            {
                               (yyval.nodito) = new Nodo("REM","");
                               Nodo *n = new Nodo("path",(yyvsp[0].text));
                               (yyval.nodito)->add(*n);
                            }
#line 2552 "parser.cpp"
    break;

  case 129:
#line 478 "Sintactico.y"
                                {
                                  (yyval.nodito) = new Nodo("REM","");
                                  Nodo *n = new Nodo("path",(yyvsp[0].text));
                                  (yyval.nodito)->add(*n);
                                }
#line 2562 "parser.cpp"
    break;

  case 130:
#line 484 "Sintactico.y"
                      {
                        (yyval.nodito) = (yyvsp[-1].nodito);
                        (yyval.nodito)->add(*(yyvsp[0].nodito));
                      }
#line 2571 "parser.cpp"
    break;

  case 131:
#line 488 "Sintactico.y"
                   {
                     (yyval.nodito) = new Nodo("PARAMETRO","");
                     (yyval.nodito)->add(*(yyvsp[0].nodito));
                   }
#line 2580 "parser.cpp"
    break;

  case 132:
#line 493 "Sintactico.y"
                            { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2586 "parser.cpp"
    break;

  case 133:
#line 494 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2592 "parser.cpp"
    break;

  case 134:
#line 495 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("cont", (yyvsp[0].text)); }
#line 2598 "parser.cpp"
    break;

  case 135:
#line 497 "Sintactico.y"
                   {
                    (yyval.nodito) = (yyvsp[-1].nodito);
                    (yyval.nodito)->add(*(yyvsp[0].nodito));
                   }
#line 2607 "parser.cpp"
    break;

  case 136:
#line 501 "Sintactico.y"
                 {
                   (yyval.nodito) = new Nodo("PARAMETRO","");
                   (yyval.nodito)->add(*(yyvsp[0].nodito));
                 }
#line 2616 "parser.cpp"
    break;

  case 137:
#line 506 "Sintactico.y"
                           { (yyval.nodito) = new Nodo("path", (yyvsp[0].text)); }
#line 2622 "parser.cpp"
    break;

  case 138:
#line 507 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2628 "parser.cpp"
    break;

  case 139:
#line 508 "Sintactico.y"
                                                {
                                                  (yyval.nodito) = new Nodo("name","");
                                                  Nodo *n = new Nodo((yyvsp[-1].text),(yyvsp[0].text));
                                                  (yyval.nodito)->add(*n);
                                                }
#line 2638 "parser.cpp"
    break;

  case 140:
#line 513 "Sintactico.y"
                               { (yyval.nodito) = new Nodo("name",(yyvsp[0].text)); }
#line 2644 "parser.cpp"
    break;

  case 141:
#line 515 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2653 "parser.cpp"
    break;

  case 142:
#line 519 "Sintactico.y"
                    {
                      (yyval.nodito) = new Nodo("PARAMETRO","");
                      (yyval.nodito)->add(*(yyvsp[0].nodito));
                    }
#line 2662 "parser.cpp"
    break;

  case 143:
#line 524 "Sintactico.y"
                                   { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2668 "parser.cpp"
    break;

  case 144:
#line 525 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2674 "parser.cpp"
    break;

  case 145:
#line 526 "Sintactico.y"
                                      { (yyval.nodito) = new Nodo("id",(yyvsp[0].text));}
#line 2680 "parser.cpp"
    break;

  case 146:
#line 527 "Sintactico.y"
                 { (yyval.nodito) = new Nodo("p",""); }
#line 2686 "parser.cpp"
    break;

  case 147:
#line 529 "Sintactico.y"
               {
                 (yyval.nodito) = (yyvsp[-1].nodito);
                 (yyval.nodito)->add(*(yyvsp[0].nodito));
               }
#line 2695 "parser.cpp"
    break;

  case 148:
#line 533 "Sintactico.y"
              {
                (yyval.nodito) = new Nodo("PARAMETRO","");
                (yyval.nodito)->add(*(yyvsp[0].nodito));
              }
#line 2704 "parser.cpp"
    break;

  case 149:
#line 538 "Sintactico.y"
                          { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2710 "parser.cpp"
    break;

  case 150:
#line 539 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2716 "parser.cpp"
    break;

  case 151:
#line 540 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2722 "parser.cpp"
    break;

  case 152:
#line 541 "Sintactico.y"
                            { (yyval.nodito) = new Nodo("dest",(yyvsp[0].text)); }
#line 2728 "parser.cpp"
    break;

  case 153:
#line 542 "Sintactico.y"
                              { (yyval.nodito) = new Nodo("dest",(yyvsp[0].text)); }
#line 2734 "parser.cpp"
    break;

  case 154:
#line 543 "Sintactico.y"
                                  { (yyval.nodito) = new Nodo("dest",(yyvsp[0].text)); }
#line 2740 "parser.cpp"
    break;

  case 155:
#line 545 "Sintactico.y"
                {
                  (yyval.nodito) = (yyvsp[-1].nodito);
                  (yyval.nodito)->add(*(yyvsp[0].nodito));
                }
#line 2749 "parser.cpp"
    break;

  case 156:
#line 549 "Sintactico.y"
              {
                (yyval.nodito) = new Nodo("PARAMETRO","");
                (yyval.nodito)->add(*(yyvsp[0].nodito));
              }
#line 2758 "parser.cpp"
    break;

  case 157:
#line 554 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2767 "parser.cpp"
    break;

  case 158:
#line 558 "Sintactico.y"
                     {
                       (yyval.nodito) = new Nodo("PARAMETRO","");
                       (yyval.nodito)->add(*(yyvsp[0].nodito));
                     }
#line 2776 "parser.cpp"
    break;

  case 159:
#line 563 "Sintactico.y"
                             { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2782 "parser.cpp"
    break;

  case 160:
#line 564 "Sintactico.y"
                                 { (yyval.nodito) = new Nodo("path", (yyvsp[0].text)); }
#line 2788 "parser.cpp"
    break;

  case 161:
#line 565 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("path",(yyvsp[0].text)); }
#line 2794 "parser.cpp"
    break;

  case 162:
#line 566 "Sintactico.y"
                                       { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2800 "parser.cpp"
    break;

  case 163:
#line 567 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2806 "parser.cpp"
    break;

  case 164:
#line 568 "Sintactico.y"
                 { (yyval.nodito) = new Nodo("r",""); }
#line 2812 "parser.cpp"
    break;

  case 165:
#line 570 "Sintactico.y"
                         {
                           (yyval.nodito) = (yyvsp[-1].nodito);
                           (yyval.nodito)->add(*(yyvsp[0].nodito));
                         }
#line 2821 "parser.cpp"
    break;

  case 166:
#line 574 "Sintactico.y"
                    {
                      (yyval.nodito) = new Nodo("PARAMETRO","");
                      (yyval.nodito)->add(*(yyvsp[0].nodito));
                    }
#line 2830 "parser.cpp"
    break;

  case 167:
#line 579 "Sintactico.y"
                                     { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2836 "parser.cpp"
    break;

  case 168:
#line 580 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("user",(yyvsp[0].text)); }
#line 2842 "parser.cpp"
    break;

  case 169:
#line 581 "Sintactico.y"
                                       { (yyval.nodito) = new Nodo("group",(yyvsp[0].text));}
#line 2848 "parser.cpp"
    break;

  case 170:
#line 582 "Sintactico.y"
                                { (yyval.nodito) = new Nodo("group", (yyvsp[0].text)); }
#line 2854 "parser.cpp"
    break;

  case 171:
#line 584 "Sintactico.y"
                                          {
                                            (yyval.nodito) = new Nodo("RECOVERY","");
                                            Nodo *n = new Nodo("id",(yyvsp[0].text));
                                            (yyval.nodito)->add(*n);
                                          }
#line 2864 "parser.cpp"
    break;

  case 172:
#line 590 "Sintactico.y"
                                 {
                                    (yyval.nodito) = new Nodo("LOSS","");
                                    Nodo *n = new Nodo("id",(yyvsp[0].text));
                                    (yyval.nodito)->add(*n);
                                 }
#line 2874 "parser.cpp"
    break;


#line 2878 "parser.cpp"

      default: break;
    }
  /* User semantic actions sometimes alter yychar, and that requires
     that yytoken be updated with the new translation.  We take the
     approach of translating immediately before every use of yytoken.
     One alternative is translating here after every semantic action,
     but that translation would be missed if the semantic action invokes
     YYABORT, YYACCEPT, or YYERROR immediately after altering yychar or
     if it invokes YYBACKUP.  In the case of YYABORT or YYACCEPT, an
     incorrect destructor might then be invoked immediately.  In the
     case of YYERROR or YYBACKUP, subsequent parser actions might lead
     to an incorrect destructor call or verbose syntax error message
     before the lookahead is translated.  */
  YY_SYMBOL_PRINT ("-> $$ =", yyr1[yyn], &yyval, &yyloc);

  YYPOPSTACK (yylen);
  yylen = 0;
  YY_STACK_PRINT (yyss, yyssp);

  *++yyvsp = yyval;

  /* Now 'shift' the result of the reduction.  Determine what state
     that goes to, based on the state we popped back to and the rule
     number reduced by.  */
  {
    const int yylhs = yyr1[yyn] - YYNTOKENS;
    const int yyi = yypgoto[yylhs] + *yyssp;
    yystate = (0 <= yyi && yyi <= YYLAST && yycheck[yyi] == *yyssp
               ? yytable[yyi]
               : yydefgoto[yylhs]);
  }

  goto yynewstate;


/*--------------------------------------.
| yyerrlab -- here on detecting error.  |
`--------------------------------------*/
yyerrlab:
  /* Make sure we have latest lookahead translation.  See comments at
     user semantic actions for why this is necessary.  */
  yytoken = yychar == YYEMPTY ? YYEMPTY : YYTRANSLATE (yychar);

  /* If not already recovering from an error, report this error.  */
  if (!yyerrstatus)
    {
      ++yynerrs;
#if ! YYERROR_VERBOSE
      yyerror (YY_("syntax error"));
#else
# define YYSYNTAX_ERROR yysyntax_error (&yymsg_alloc, &yymsg, \
                                        yyssp, yytoken)
      {
        char const *yymsgp = YY_("syntax error");
        int yysyntax_error_status;
        yysyntax_error_status = YYSYNTAX_ERROR;
        if (yysyntax_error_status == 0)
          yymsgp = yymsg;
        else if (yysyntax_error_status == 1)
          {
            if (yymsg != yymsgbuf)
              YYSTACK_FREE (yymsg);
            yymsg = YY_CAST (char *, YYSTACK_ALLOC (YY_CAST (YYSIZE_T, yymsg_alloc)));
            if (!yymsg)
              {
                yymsg = yymsgbuf;
                yymsg_alloc = sizeof yymsgbuf;
                yysyntax_error_status = 2;
              }
            else
              {
                yysyntax_error_status = YYSYNTAX_ERROR;
                yymsgp = yymsg;
              }
          }
        yyerror (yymsgp);
        if (yysyntax_error_status == 2)
          goto yyexhaustedlab;
      }
# undef YYSYNTAX_ERROR
#endif
    }



  if (yyerrstatus == 3)
    {
      /* If just tried and failed to reuse lookahead token after an
         error, discard it.  */

      if (yychar <= YYEOF)
        {
          /* Return failure if at end of input.  */
          if (yychar == YYEOF)
            YYABORT;
        }
      else
        {
          yydestruct ("Error: discarding",
                      yytoken, &yylval);
          yychar = YYEMPTY;
        }
    }

  /* Else will try to reuse lookahead token after shifting the error
     token.  */
  goto yyerrlab1;


/*---------------------------------------------------.
| yyerrorlab -- error raised explicitly by YYERROR.  |
`---------------------------------------------------*/
yyerrorlab:
  /* Pacify compilers when the user code never invokes YYERROR and the
     label yyerrorlab therefore never appears in user code.  */
  if (0)
    YYERROR;

  /* Do not reclaim the symbols of the rule whose action triggered
     this YYERROR.  */
  YYPOPSTACK (yylen);
  yylen = 0;
  YY_STACK_PRINT (yyss, yyssp);
  yystate = *yyssp;
  goto yyerrlab1;


/*-------------------------------------------------------------.
| yyerrlab1 -- common code for both syntax error and YYERROR.  |
`-------------------------------------------------------------*/
yyerrlab1:
  yyerrstatus = 3;      /* Each real token shifted decrements this.  */

  for (;;)
    {
      yyn = yypact[yystate];
      if (!yypact_value_is_default (yyn))
        {
          yyn += YYTERROR;
          if (0 <= yyn && yyn <= YYLAST && yycheck[yyn] == YYTERROR)
            {
              yyn = yytable[yyn];
              if (0 < yyn)
                break;
            }
        }

      /* Pop the current state because it cannot handle the error token.  */
      if (yyssp == yyss)
        YYABORT;


      yydestruct ("Error: popping",
                  yystos[yystate], yyvsp);
      YYPOPSTACK (1);
      yystate = *yyssp;
      YY_STACK_PRINT (yyss, yyssp);
    }

  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++yyvsp = yylval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END


  /* Shift the error token.  */
  YY_SYMBOL_PRINT ("Shifting", yystos[yyn], yyvsp, yylsp);

  yystate = yyn;
  goto yynewstate;


/*-------------------------------------.
| yyacceptlab -- YYACCEPT comes here.  |
`-------------------------------------*/
yyacceptlab:
  yyresult = 0;
  goto yyreturn;


/*-----------------------------------.
| yyabortlab -- YYABORT comes here.  |
`-----------------------------------*/
yyabortlab:
  yyresult = 1;
  goto yyreturn;


#if !defined yyoverflow || YYERROR_VERBOSE
/*-------------------------------------------------.
| yyexhaustedlab -- memory exhaustion comes here.  |
`-------------------------------------------------*/
yyexhaustedlab:
  yyerror (YY_("memory exhausted"));
  yyresult = 2;
  /* Fall through.  */
#endif


/*-----------------------------------------------------.
| yyreturn -- parsing is finished, return the result.  |
`-----------------------------------------------------*/
yyreturn:
  if (yychar != YYEMPTY)
    {
      /* Make sure we have latest lookahead translation.  See comments at
         user semantic actions for why this is necessary.  */
      yytoken = YYTRANSLATE (yychar);
      yydestruct ("Cleanup: discarding lookahead",
                  yytoken, &yylval);
    }
  /* Do not reclaim the symbols of the rule whose action triggered
     this YYABORT or YYACCEPT.  */
  YYPOPSTACK (yylen);
  YY_STACK_PRINT (yyss, yyssp);
  while (yyssp != yyss)
    {
      yydestruct ("Cleanup: popping",
                  yystos[+*yyssp], yyvsp);
      YYPOPSTACK (1);
    }
#ifndef yyoverflow
  if (yyss != yyssa)
    YYSTACK_FREE (yyss);
#endif
#if YYERROR_VERBOSE
  if (yymsg != yymsgbuf)
    YYSTACK_FREE (yymsg);
#endif
  return yyresult;
}

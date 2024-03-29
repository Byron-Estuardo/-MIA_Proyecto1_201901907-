/* A Bison parser, made by GNU Bison 3.5.1.  */

/* Bison interface for Yacc-like parsers in C

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

/* Undocumented macros, especially those whose name start with YY_,
   are private implementation details.  Do not rely on them.  */

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
    

#line 143 "parser.h"

};
typedef union YYSTYPE YYSTYPE;
# define YYSTYPE_IS_TRIVIAL 1
# define YYSTYPE_IS_DECLARED 1
#endif


extern YYSTYPE yylval;

int yyparse (void);

#endif /* !YY_YY_PARSER_H_INCLUDED  */

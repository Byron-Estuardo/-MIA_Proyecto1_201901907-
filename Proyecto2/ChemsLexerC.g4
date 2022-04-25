lexer grammar ChemsLexerC;


// Tokens

MKDISK:          M K D I S K;
RMDISK:          R M D I S K;
FDISK:           F D I S K;
MOUNT:           M O U N T;
MKFS:            M K F S;
LOGIN:           L O G I N;
LOGOUT:          L O G O U T;
MKGRP:           M K G R P;
RMGRP:           R M G R P;
MKUSER:          M K U S E R;
RMUSR:           R M U S R;
MKFILE:          M K F I L E;
MKDIR:           M K D I R;
EXEC:            E X E C;
REP:             R E P;
SIZE:            '-' S I Z E;
FIT:             '-' F I T;
UNIT:            '-' U N I T;
PATH:            '-' P A T H;
NAME:            '-' N A M E;
TYPE:            '-' T Y P E;
ID:              '-' I D;
BF:              B F;
FF:              F F;
WF:              W F;
FAST:            F A S T;
FULL:            F U L L ;
USR:             '-' U S U A R I O;
PWD:             '-' P A S S W O R D;
PWD1:            '-' P W D;
GRP:             '-' G R P;
GR:              '-' R;
CONT:            '-' C O N T;
GP:              '-' P;
PAUSA:           P A U S E;
DISK:            D I S K;
TREE:            T R E E;
FILE:            F I L E;

LETRA: [a-zA-Z_];
ENTERO: [0-9]+;
CADENA: '"'~["]*'"';
EXTENSION: '.'{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*};
IDENTIFICADOR: ([a-zA-Z_])[a-zA-Z0-9_]*;
DIAGONAL: '/';
//CARACTER: ({[a-zA-Z]}|{('-')?[0-9]+});
//IDENTIFICADOR: {[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*;
//PASSWORD: ({[a-zA-Z]}|{[a-zA-Z]}|[!$@+*])+;
//RUTA: ({'/'}{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*})*({'/'}{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*}{'.'{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*}});
//DIRECTORIO: ({'/'}{{[a-zA-Z]}({[a-zA-Z]}|{('-')?[0-9]+}|'_')*})+;

IGUAL:           '=';

WHITESPACE: [ \\\r\n\t]+ -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

fragment A:('a'|'A');
fragment B:('b'|'B');
fragment C:('c'|'C');
fragment D:('d'|'D');
fragment E:('e'|'E');
fragment F:('f'|'F');
fragment G:('g'|'G');
fragment H:('h'|'H');
fragment I:('i'|'I');
fragment J:('j'|'J');
fragment K:('k'|'K');
fragment L:('l'|'L');
fragment M:('m'|'M');
fragment N:('n'|'N');
fragment O:('o'|'O');
fragment P:('p'|'P');
fragment Q:('q'|'Q');
fragment R:('r'|'R');
fragment S:('s'|'S');
fragment T:('t'|'T');
fragment U:('u'|'U');
fragment V:('v'|'V');
fragment W:('w'| 'W');
fragment X:('x'| 'X');
fragment Y:('y'| 'Y');
fragment Z:('z'| 'Z');

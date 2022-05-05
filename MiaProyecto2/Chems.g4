parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}

@header {
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

}

start 
: adminDiscos
| adminArchivos
| adminUsuarios
| adminCarpetas
| PAUSA { AD.Pausa() }
;

adminDiscos
: MKDISK pmkdisk { AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath); ValSize = ""; ValFit = ""; ValUnit = ""; ValPath = ""; }
| RMDISK prmdisk { AD.Rmdisk(ValPath); ValPath = ""; }
| FDISK pfdisk  { AD.Fdisk(ValSize, ValFit, ValUnit, ValPath, ValType, ValName); ValSize = ""; ValFit = ""; ValUnit = ""; ValPath = ""; ValType = ""; ValName = "";}
| MOUNT pmount  { AD.Mount(ValPath, ValName); ValPath = ""; ValName = "" }
| EOF
;

adminArchivos
: MKFS pmkfs { AD.Mkfs(ValIdentificador, ValType); ValIdentificador = ""; ValType = ""; }
;

adminUsuarios
: LOGIN plogin { AD.Login(ValUsuario, ValPassword, ValIdentificador); ValUsuario = "" ; ValPassword = "" ; ValIdentificador = ""; }
| LOGOUT { AD.LogOut() }
| MKGRP pmkgrp { AD.Mkgrp(ValName); ValName = ""; }
| RMGRP prmkgrp { AD.RMkgrp(ValName); ValName = ""; }
| MKUSER pmkuser { AD.Mkusr(ValUsuario, ValPassword, ValGrupo); ValUsuario = "" ; ValPassword = "" ; ValGrupo = ""; }
| RMUSR prmkusr { AD.RMkusr(ValName); ValName = ""; }
;

adminCarpetas
: MKFILE pmkfile { AD.MkFile(ValPath, ValR, ValSize, ValCont); ValPath = ""; ValR = false; ValSize = ""; ValCont = "" }
;

pmkfile
: parametrosmkfile
| pmkfile parametrosmkfile
;

parametrosmkfile
: SIZE IGUAL ENTERO { 
    ValSize = $ENTERO.text
}
| CONT IGUAL BF { ValCont = "" }
| CONT IGUAL CADENA   { str:= $CADENA.text[1:len($CADENA.text)-1]
                        ValCont = str }
| CONT IGUAL RUTA   { ValCont = $RUTA.text }
| UNIT IGUAL LETRA { ValUnit = $LETRA.text }
| PATH IGUAL CADENA   { str:= $CADENA.text[1:len($CADENA.text)-1]
                        ValPath = str }
| PATH IGUAL RUTA   { ValPath = $RUTA.text }
| GR { ValR = true }
;

pmkdisk
: parametrosmkdisk
| pmkdisk parametrosmkdisk
;

parametrosmkdisk
: SIZE IGUAL ENTERO { ValSize = $ENTERO.text }
| FIT IGUAL BF { ValFit = "b" }
| FIT IGUAL FF { ValFit = "f" }
| FIT IGUAL WF { ValFit = "w" }
| UNIT IGUAL LETRA { ValUnit = $LETRA.text }
| PATH IGUAL CADENA   { str:= $CADENA.text[1:len($CADENA.text)-1]
                        ValPath = str }
| PATH IGUAL RUTA   { ValPath = $RUTA.text }
;

prmdisk
: PATH IGUAL CADENA   {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValPath = str }
| PATH IGUAL RUTA   { ValPath = $RUTA.text }
;

pmkgrp
: NAME IGUAL CADENA           {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValName = str }
| NAME IGUAL IDENTIFICADOR   { ValName = $IDENTIFICADOR.text }
;

prmkgrp
: NAME IGUAL CADENA           {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValName = str }
| NAME IGUAL IDENTIFICADOR   { ValName = $IDENTIFICADOR.text }
;

prmkusr
: NAME IGUAL CADENA           {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValName = str }
| NAME IGUAL IDENTIFICADOR   { ValName = $IDENTIFICADOR.text }
;

pfdisk
: parametrosfdisk
| pfdisk parametrosfdisk 
;

parametrosfdisk
: parametrosmkdisk
| TYPE IGUAL LETRA { ValType = $LETRA.text } 
| NAME IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1]
                      ValName = str}
| NAME IGUAL IDENTIFICADOR { ValName = $IDENTIFICADOR.text}
;

pmount
: parametrosmount
| pmount parametrosmount 
;

parametrosmount
: PATH IGUAL CADENA   {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValPath = str }
| PATH IGUAL RUTA     { ValPath = $RUTA.text }
| NAME IGUAL IDENTIFICADOR   { ValName = $IDENTIFICADOR.text}
;

pmkfs
: parametrosmkfs
| pmkfs parametrosmkfs
;

parametrosmkfs
: ID IGUAL IDENTIFICADOR {  ValIdentificador = $IDENTIFICADOR.text }
| TYPE IGUAL LETRA { ValType = $LETRA.text } 
;

plogin
: parametroslogin
| plogin parametroslogin
;

parametroslogin
:  USR IGUAL IDENTIFICADOR { ValUsuario = $IDENTIFICADOR.text }
|  USR IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1]
                      ValUsuario = str }
| PWD IGUAL IDENTIFICADOR { ValPassword = $IDENTIFICADOR.text }
| PWD IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1]
                      ValPassword = str }
| ID IGUAL IDENTIFICADOR {  ValIdentificador = $IDENTIFICADOR.text }
;

pmkuser
: parametrosmkuser
| plogin parametrosmkuser
;

parametrosmkuser
:  USR IGUAL IDENTIFICADOR { ValUsuario = $IDENTIFICADOR.text }
|  USR IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1]
                      ValUsuario = str }
| PWD1 IGUAL IDENTIFICADOR { ValPassword = $IDENTIFICADOR.text }
| PWD1 IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1]
                      ValPassword = str }
| GRP IGUAL IDENTIFICADOR {  ValGrupo = $IDENTIFICADOR.text }
| GRP IGUAL CADENA           {  str:= $CADENA.text[1:len($CADENA.text)-1]
                                ValGrupo = str }
;

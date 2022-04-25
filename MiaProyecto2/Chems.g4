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
}

start 
: adminDiscos
| PAUSA { AD.Pausa() }
| EOF
;

adminDiscos
: MKDISK pmkdisk
| RMDISK prmdisk
| FDISK pfdisk
| MOUNT pmount
;

pmkdisk
: pmkdisk parametrosmkdisk {  }
| parametrosmkdisk { AD.Mkdisk(ValSize, ValFit, ValUnit, ValPath); ValSize=""  }
;

parametrosmkdisk
: SIZE IGUAL ENTERO { ValSize = $ENTERO.text } 
| FIT IGUAL BF { ValFit = "b" }
| FIT IGUAL FF { ValFit = "f" }
| FIT IGUAL WF { ValFit = "w" }
| UNIT IGUAL CADENA { str:= $CADENA.text[1:len($CADENA.text)-1] ValUnit = str }
| PATH IGUAL RUTA   { ValPath = $RUTA }
;

prmdisk
: RMDISK PATH IGUAL RUTA { ValPath = $RUTA }
;

pfdisk
: pfdisk parametrosfdisk {  }
| parametrosfdisk {  }
;

parametrosfdisk
: parametrosmkdisk
| TYPE IGUAL CADENA { ValType = $CADENA.text } 
| NAME IGUAL CADENA { ValName = $CADENA.text }
;

pmount
: pmount parametrosmount {  }
| parametrosmount {  }
;

parametrosmount
: PATH IGUAL RUTA   { ValPath = $RUTA.text }
| NAME IGUAL CADENA { ValName = $CADENA.text }
;



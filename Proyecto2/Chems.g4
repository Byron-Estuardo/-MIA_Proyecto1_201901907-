parser grammar Chems;



options {
  tokenVocab = ChemsLexer;
}

@header {
    import "Proyecto2/nodos"
}

start returns[*nodos.Nodo Envio]
: adminDiscos { $Envio = $adminDiscos.Raiz  }
| PAUSA
;

adminDiscos returns[*nodos.Nodo Raiz]
: MKDISK mkdisk{  fmt.Println("hOLA")
                  nodos.NewNodo("MKDISK", "")
                  $Raiz = nodos.Raiz
              }
  | RMDISK 
;

mkdisk
: mkdisk parametrosmkdisk { fmt.Println("hOLA 2")
                            nodos.AddParameters($parametrosmkdisk.NODS) }
| parametrosmkdisk {  fmt.Println("Parametros")
                      nodos.Add(nodos.Nodoo("PARAMETRO", ""))
                      nodos.AddParameters($parametrosmkdisk.NODS) }
;

parametrosmkdisk returns[nodos.Nodo NODS]
: SIZE IGUAL ENTERO { $NODS = nodos.Nodoo("size", $ENTERO.text) } 
| FIT IGUAL ajuste {  $NODS = nodos.Nodoo("fit", "") 
                      nodos.AddParameters($ajuste.AJ) }
|UNIT IGUAL CARACTER { $NODS = nodos.Nodoo("unit", $CARACTER.text) }
|PATH IGUAL RUTA { $NODS = nodos.Nodoo("path", $RUTA.text) }
;

ajuste returns[nodos.Nodo AJ]
: BF { $AJ = nodos.Nodoo("ajuste", "bf") }
| FF { $AJ = nodos.Nodoo("ajuste", "ff") }
| WF { $AJ = nodos.Nodoo("ajuste", "wf") }
;


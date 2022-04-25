package nodos

import "fmt"

type Nodo struct {
	tipo  string
	valor string
	tipo_ int
	hijos []Nodo
}

var Raiz *Nodo = new(Nodo)
var Nuevos []Nodo

func Nodoo(t string, v string) Nodo {
	var nuevo []Nodo
	tipo2 := GetTipo(t)
	Envio := Nodo{t, v, tipo2, nuevo}
	//Nodos(t, v, tipo2)
	return Envio
}

func Pruebas() {
	fmt.Println("A perro")
}

func GetTipo(tipo1 string) int {
	if tipo1 == "MKDISK" {
		return 1
	}
	return 0
}

func NewNodo(t string, v string) *Nodo {
	var nuevo []Nodo
	tipo2 := GetTipo(t)
	Raiz = &Nodo{t, v, tipo2, nuevo}
	return Raiz
}

func Add(n Nodo) {
	Raiz.hijos = append(Raiz.hijos, n)
}

func AddParameters(n Nodo) {
	if len(Raiz.hijos) != 0 {
		Raiz.hijos[0].hijos = append(Raiz.hijos[0].hijos, n)
		fmt.Println("F perro")
	}
}

/*

func main() {
	fmt.Println("Hola mundo ")
	fmt.Println(Raiz)
	Raiz = NewNodo("MKDISK", "")
	Add(Nodoo("Parametros", ""))
	fmt.Println(Raiz)
	AddParameters(Nodoo("size", "5"))
	AddParameters(Nodoo("path", "/home/sdasd/asdasd/dk"))
	fmt.Println(Raiz)
	fmt.Println(Raiz)
}


*/

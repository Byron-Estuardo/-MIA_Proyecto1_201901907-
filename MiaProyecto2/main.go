package main

import (
	"MiaProyecto2/AdminDisks"
	//	"os"
	//	"strings"
	//
	//"github.com/antlr/antlr4/runtime/Go/antlr"
)

/*
type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func AnalizarTexto(a string) {
	// Setup the input
	is := antlr.NewInputStream(a)

	//is, _ := antlr.NewFileStream("entrada.txt")
	// Create the Lexer
	lexer := parser.NewChemsLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	fmt.Println("Hola mundo ")
}

func menu() {
	finalizar := false
	fmt.Println("Ingrese un Comando:")
	reader := bufio.NewReader(os.Stdin)
	//  Ciclo para lectura de multiples comandos
	for !finalizar {
		comando, _ := reader.ReadString('\n')
		if strings.Contains(comando, "pause\n") {
			fmt.Println("Pausado:")
			AdminDisks.Pausa()
			fmt.Println("continua xd")
		} else {
			if comando != "" && comando != "pause\n" {
				ejecucion_comando(comando)
			}
		}
	}
}

func ejecucion_comando(commandArray string) {
	data := strings.ToLower(commandArray)
	if data == "crear_disco" {
		fmt.Println("crear disco")
	} else if data == "escribir" {
		fmt.Println("Escribir")
	} else if data == "mostrar" {
		fmt.Println("mostrar")
	} else {
		fmt.Println("Comando ingresado no es valido")
	}
}
*/
func main() {
	//menu()
	//fmt.Println("Pruebs")
	AdminDisks.Mkdisk("15", "wf", "m", "/home/curious1924/Escritorio/Joder/disco23.dk")
	//fmt.Println("Pruebs1")
	//AnalizarTexto("#Comentario de prueba")
	//crearDisco("/home/curious1924/Escritorio/disco23.dk")
}

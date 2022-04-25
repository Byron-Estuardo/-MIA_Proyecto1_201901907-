package main

import (
	//"Proyecto2/nodos"
	"Proyecto2/parser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("rmdisk")

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
	/*
		fmt.Println("Hola mundo ")
		fmt.Println(nodos.Raiz)
		nodos.NewNodo("MKDISK", "")
		nodos.Add(nodos.Nodoo("Parametros", ""))
		fmt.Println(nodos.Raiz)
		nodos.AddParameters(nodos.Nodoo("size", "5"))
		nodos.AddParameters(nodos.Nodoo("path", "/home/sdasd/asdasd/dk"))
		fmt.Println(nodos.Raiz)
		fmt.Println(nodos.Raiz)
		//Raiz.hijos[0].hijos = append(Raiz.hijos[0].hijos, *newNodo("size", "5"))
		//fmt.Print(Raiz.hijos[0].hijos)
	*/
}

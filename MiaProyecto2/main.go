package main

import (
	"MiaProyecto2/parser"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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
}

var execs = regexp.MustCompile("(?i)exec")

func main() {
	menu :=
		`
------------------------------INGRESE UN COMANDO------------------------------
--------------------------------exit para salir-------------------------------

>`
	finalizar := false
	fmt.Println(menu)
	reader := bufio.NewReader(os.Stdin)
	for !finalizar {
		fmt.Println("Ingrese un Comando:")
		comando, _ := reader.ReadString('\n')
		comando = strings.TrimRight(comando, "\n")
		fmt.Println(comando)
		if comando == "exit" {
			fmt.Println("Adios!")
			break
		}
		if execs.MatchString(comando) {
			a := strings.Split(comando, "=")
			path := a[1]
			if strings.Contains(path, "\"") {
				path = path[1 : len(path)-1]
			}

			readFile, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			var lines []string
			for fileScanner.Scan() {
				lines = append(lines, fileScanner.Text())
			}
			readFile.Close()
			for _, line := range lines {
				AnalizarTexto(line)
			}

		} else {
			AnalizarTexto(comando)
		}
	}
}

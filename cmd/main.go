//Package main
//
//The following lines will trigger the automatic generation
//go:generate rm -f -r $PWD/myparser/
//go:generate antlr -Dlanguage=Go -package myparser -o myparser Expr.g4
package main

import (
	parser "/myparser"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {

	fmt.Println("Launching ...")
	return

	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.NewExprParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&exprListener{}, p.Start())

}

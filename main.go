//Package main
//
//The following lines will trigger the automatic generation
//go:generate rm -f -r $PWD/generated/parser/
//go:generate antlr -Dlanguage=Go -package parser -o generated/parser Expr.g4
package main

import (
	"./generated/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {

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

//Package main
//
//The following lines will trigger the automatic generation
//go:generate rm -f -r internal/pkg/myparser/
//go:generate antlr -Dlanguage=Go -package myparser -o internal/pkg/myparser Expr.g4
package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/myparser"
)

func main() {

	fmt.Println("Launching ...")

	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := myparser.NewExprLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s \t%d \t%q \n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetTokenType(), t.GetText())
	}
	return

}

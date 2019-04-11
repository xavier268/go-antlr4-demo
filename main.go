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

var p *myparser.ExprParser

func main() {

	fmt.Println("Launching ...")

	// Setup the input
	const intxt = "1 + 2 * 3\ntiti = 56 + 68\n"
	is := antlr.NewInputStream(intxt)
	fmt.Println("Preparing to process : \n", intxt)

	// Create the Lexer
	lexer := myparser.NewExprLexer(is)

	// Read all tokens
	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	fmt.Printf("%d \t%s \t%q \n",
	// 		t.GetTokenType(),
	// 		lexer.SymbolicNames[t.GetTokenType()],
	// 		t.GetText())
	// }

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p = myparser.NewExprParser(stream)

	// Build ast, starting with the "prog" rule
	ast := p.Prog()
	fmt.Println("Dumping ast : ", ast.ToStringTree(nil, p))

	// Walk the tree with the dump listener
	listener := new(dumpListener)
	antlr.ParseTreeWalkerDefault.Walk(listener, ast)

	return

}

//dumpListener will dump the rules and context
type dumpListener struct {
	*myparser.BaseExprListener
}

func (dl *dumpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("\n<<<< EXITING rule, context : %q, rule : %d\n", ctx.GetText(), ctx.GetRuleIndex())
	fmt.Println(ctx.ToStringTree(nil, p))
}

func (dl *dumpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("\n>>>> ENTERING rule, context : %q, rule : %d\n", ctx.GetText(), ctx.GetRuleIndex())
	fmt.Println(ctx.ToStringTree(nil, p))
}

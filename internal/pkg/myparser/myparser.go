//Package myparser provides generic parsing, listening and visiting entry points.
package myparser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/auto"
)

//MyParser is the base object for parsing, walking or visiting a parsed tree
type MyParser struct {
	Parser *auto.ExprParser  // The expression parser
	Root   auto.IProgContext // The root of the parsed AST
}

//Parse parse the string, setting the ast (abstract syntax tree)
func (mp *MyParser) Parse(intxt string) {
	fmt.Println("\n\nParsing from string input :\n", intxt)
	is := antlr.NewInputStream(intxt)
	lexer := auto.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	mp.Parser = auto.NewExprParser(stream)
	mp.Root = mp.Parser.Prog()
}

//WalkWith run the listener over the ast (abstract syntax tree)
func (mp *MyParser) WalkWith(list auto.ExprListener) {
	if mp.Root == nil {
		panic("Cannot walk before a parsing was made ?!")
	}
	antlr.ParseTreeWalkerDefault.Walk(list, mp.Root)
}

//Dump the ast in a readable format
func (mp *MyParser) Dump() {
	fmt.Println("Abstract Syntax Tree : ", mp.Root.ToStringTree(nil, mp.Parser))
}

//VisitWith visit root using the provided visitor.
func (mp *MyParser) VisitWith(v antlr.ParseTreeVisitor) interface{} {
	return mp.Root.Accept(v)
}

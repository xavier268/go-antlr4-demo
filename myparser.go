package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/myparser"
)

//MyParser is the base object for parsing and walking a parsed tree
type MyParser struct {
	Parser *myparser.ExprParser
	Root   myparser.IProgContext
}

//Parse parse the string, returning the ast (abstract syntax tree)
func (mp *MyParser) Parse(intxt string) {
	fmt.Println("Parsing from string input :\n", intxt)
	is := antlr.NewInputStream(intxt)
	lexer := myparser.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	mp.Parser = myparser.NewExprParser(stream)
	mp.Root = mp.Parser.Prog()
}

//Walk run the listener over the ast (abstract syntax tree)
func (mp *MyParser) Walk(list myparser.ExprListener) {
	if mp.Root == nil {
		panic("Cannot walk before a parsing was made ?!")
	}
	antlr.ParseTreeWalkerDefault.Walk(list, mp.Root)
}

//Dump the ast in a readable format
func (mp *MyParser) Dump() {
	fmt.Println("Abstract Syntax Tree : ", mp.Root.ToStringTree(nil, mp.Parser))
}

package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/auto"
)

//MyParser is the base object for parsing and walking a parsed tree
type MyParser struct {
	Parser *auto.ExprParser
	Root   auto.IProgContext
}

//Parse parse the string, returning the ast (abstract syntax tree)
func (mp *MyParser) Parse(intxt string) {
	fmt.Println("\n\nParsing from string input :\n", intxt)
	is := antlr.NewInputStream(intxt)
	lexer := auto.NewExprLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	mp.Parser = auto.NewExprParser(stream)
	mp.Root = mp.Parser.Prog()
}

//Walk run the listener over the ast (abstract syntax tree)
func (mp *MyParser) Walk(list auto.ExprListener) {
	if mp.Root == nil {
		panic("Cannot walk before a parsing was made ?!")
	}
	antlr.ParseTreeWalkerDefault.Walk(list, mp.Root)
}

//Dump the ast in a readable format
func (mp *MyParser) Dump() {
	fmt.Println("Abstract Syntax Tree : ", mp.Root.ToStringTree(nil, mp.Parser))
}

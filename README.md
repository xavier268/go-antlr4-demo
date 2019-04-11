# go-antlr4-demo

**Demo of a antlr4 parser generating golang**

This repo is a demo and test of using [antlr4](https://www.antlr.org/) to generate a parser in Golang.
The available Golang doc for [antlr4](https://www.antlr.org/) is still behind the java doc, but [antlr4](https://www.antlr.org/) is 
so much simpler and more powerfull than yacc to generate both parser and lexer that it is clearly worth the effort !

Ideas implemented in this demo include :

* Abstract Syntax Tree (AST) to run multiple sucessive listeners
* dedicated listener to dump/pretty print the ast
* implementation of annotating the ast, using a map on the node payload (similar to the dedicated Map in java)

See also : https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/


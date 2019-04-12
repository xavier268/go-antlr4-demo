//Package main
//
//The following lines will trigger the automatic generation
//go:generate rm -f -r internal/pkg/myparser/
//go:generate antlr -Dlanguage=Go -package myparser -o internal/pkg/myparser Expr.g4
package main

func main() {

	// Setup the input
	const intxt = "1 - 2 * 3\ntiti = 56 + 68\n"

	mp := new(MyParser)
	mp.Parse(intxt)
	mp.Dump()

	cl := NewComputeListener()
	mp.Walk(cl)
	cl.dumpMap()

	dl := NewDumpListener(mp.Parser)
	mp.Walk(dl)

	return
}

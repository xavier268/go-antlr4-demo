//Package main
package main

import (
	"fmt"

	"github.com/xavier268/go-antlr4-demo/internal/pkg/myparser"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/myparser/listeners"
)

func main() {

	fmt.Println("Demo of using  Listener")

	// Setup the input
	const intxt = "1 - 2 * 3\ntiti = 56 + 68\n"

	mp := new(myparser.MyParser)
	mp.Parse(intxt)
	mp.Dump()

	cl := listeners.NewComputeListener()
	mp.WalkWith(cl)
	cl.DumpMaps()

	dl := listeners.NewDumpListener(mp.Parser)
	mp.WalkWith(dl)

	return
}

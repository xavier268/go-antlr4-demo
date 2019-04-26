//Package main
package main

import "fmt"

func main() {

	fmt.Println("Demo of using  Listener")

	// Setup the input
	const intxt = "1 - 2 * 3\ntiti = 56 + 68\n"

	mp := new(MyParser)
	mp.Parse(intxt)
	mp.Dump()

	cl := NewComputeListener()
	mp.Walk(cl)
	cl.dumpMaps()

	dl := NewDumpListener(mp.Parser)
	mp.Walk(dl)

	return
}

package main

import "testing"

func TestWalkBeforeParseShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			panic("The test should have generated a panic")
		}
	}()
	// Code that should panic
	mp := new(MyParser)
	mp.Walk(NewComputeListener())
}

func TestParseInvalidSyntax(t *testing.T) {
	t.Skip()
	mp := new(MyParser)
	mp.Parse("1 + $$ +  \n")
}

func TestWalkDumpListener(t *testing.T) {
	//t.Skip()
	mp := new(MyParser)
	mp.Parse("1+2\n")
	mp.Dump()
	dl := NewDumpListener(mp.Parser)
	mp.Walk(dl)
}

func TestWalkComputeListener(t *testing.T) {
	// maps to the expected value of x after the program was run
	tab := map[string]float64{
		"\n":                0, // x was not allocated
		"1\n":               0, // x was not allocated
		"y = 1\n":           0, // x was not allocated
		"x=1+1\nx=x+1\nx\n": 3,
		"x=-5+8\n":          3.0,
		"x=-x+100\n":        100,
		"x=1+2*3\n":         7,
		"x=1+-2*3\n":        -5,
		"x=-1+2*3\n":        5,
		"x=x+100\n":         100,
		"x=1 + (2*3)\n":     7,
		"x=(1+2)*3)\n":      9,
		"x=-(1+2)*3)\n":     -9,
		"x=---5\n":          -5,
		"x=--5\n":           5,
		"x=-5\n":            -5,
		"x=3/4":             0.75,
		"x=3-4":             -1,
		"x=444444444\n":     444444444,
		"y=45-5\nx=y+4\n":   44,
	}

	for k, v := range tab {
		mp := new(MyParser)
		mp.Parse(k)
		cl := NewComputeListener()
		mp.Walk(cl)
		if len(cl.ids) > 1 {
			// Multiple variables
			cl.dumpMaps()
		}
		if cl.ids["x"] != v {
			mp.Dump()
			cl.dumpMaps()
			t.Error("Expected ", v, " but got ", cl.ids["x"])
			t.FailNow()
		}
	}

}

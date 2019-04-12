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
	t.Skip()
	mp := new(MyParser)
	mp.Parse("titi = 1 +2 *3\n")
	dl := NewDumpListener(mp.Parser)
	mp.Walk(dl)
}

func TestWalkComputeListener(t *testing.T) {
	// maps to the expected value of x
	tab := map[string]float64{
		"1\n":           0, // x was not allocated
		"y = 1\n":       0, // x was not allocated
		"x=1+1\nx\nx\n": 2,
		"x=-5+8\n":      3.0,
		"x=-x+100\n":    100,
		"x=x+100\n":     100,
	}

	for k, v := range tab {
		mp := new(MyParser)
		mp.Parse(k)
		cl := NewComputeListener()
		mp.Walk(cl)
		if cl.ids["x"] != v {
			mp.Dump()
			cl.dumpMaps()
			t.Error("Expecting ", v, " got ", cl.ids["x"])
			t.FailNow()
		}
	}

}

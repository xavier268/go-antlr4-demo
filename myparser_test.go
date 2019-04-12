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
	mp := new(MyParser)
	mp.Parse("1 + $$ +  \n")
}

func TestParseValidSyntax(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("1 + 2 *3 \n")
	mp.Dump()
}

func TestWalkDumpListener(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("titi = 1 +2 *3\n")
	dl := NewDumpListener(mp.Parser)
	mp.Walk(dl)
}

func TestWalkComputeListener1(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("444\n")
	cl := NewComputeListener()
	mp.Walk(cl)
	cl.dumpMaps()
}

func TestWalkComputeListener2(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("titi = -111 + 11\n")
	mp.Dump()
	cl := NewComputeListener()
	mp.Walk(cl)
	cl.dumpMaps()
}

func TestWalkComputeListener3(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("x = 3*(10 + 5/3)\nx+2*x")
	mp.Dump()
	cl := NewComputeListener()
	mp.Walk(cl)
	cl.dumpMaps()
}

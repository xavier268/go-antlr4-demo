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
	mp.Walk(NewDumpListener(mp.Parser))
}

func TestWalkComputeListener(t *testing.T) {
	mp := new(MyParser)
	mp.Parse("titi = 1 +2 *3\n")
	mp.Walk(NewComputeListener())
}

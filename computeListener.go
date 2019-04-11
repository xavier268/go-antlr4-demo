package main

import (
	"fmt"

	"github.com/xavier268/go-antlr4-demo/internal/pkg/myparser"
)

//ComputeListener maintains an annotated tree
type ComputeListener struct {
	*myparser.BaseExprListener
	m     map[interface{}]float64
	count float64
}

// NewComputeListener constructor
func NewComputeListener() *ComputeListener {
	cl := new(ComputeListener)
	cl.m = make(map[interface{}]float64)
	return cl
}

// ExitInt is called when production int is exited.
func (cl *ComputeListener) ExitInt(ctx *myparser.IntContext) {
	fmt.Println(ctx.GetPayload())
	cl.count += 1.0
	cl.m[ctx.GetPayload()] = cl.count
}

func (cl *ComputeListener) dumpMap() {
	for k, v := range cl.m {
		fmt.Printf("\n%p\t->\t%f", k, v)
	}
}

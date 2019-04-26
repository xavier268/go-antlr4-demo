package listeners

import (
	"fmt"
	"strconv"

	"github.com/xavier268/go-antlr4-demo/internal/pkg/auto"
)

//ComputeListener maintains an annotated tree
// (to test and demo the concept, could have a stack instead ...)
type ComputeListener struct {
	*auto.BaseExprListener
	m   map[interface{}]float64 // node -> value
	ids map[string]float64      // id  -> value
}

// NewComputeListener constructor
func NewComputeListener() *ComputeListener {
	cl := new(ComputeListener)
	cl.m = make(map[interface{}]float64)
	cl.ids = make(map[string]float64)
	return cl
}

//DumpMaps will dump and display the tree and the variables map
func (cl *ComputeListener) DumpMaps() {
	fmt.Println("Dumping annotated Tree :")
	for k, v := range cl.m {
		fmt.Printf("%p\t->\t%f\n", k, v)
	}
	fmt.Println("Dumping variables :")
	for k, v := range cl.ids {
		fmt.Printf("%s\t->\t%f\n", k, v)
	}
	fmt.Println()
}

// ExitInt is called when production int is exited.
func (cl *ComputeListener) ExitInt(ctx *auto.IntContext) {
	i, e := strconv.Atoi(ctx.GetText())
	if e != nil {
		panic("Invalid integer format")
	}
	cl.m[ctx.GetPayload()] = float64(i)
}

// ExitPrintExpr is called when production printExpr is exited.
func (cl *ComputeListener) ExitPrintExpr(ctx *auto.PrintExprContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr().GetPayload()]
	fmt.Printf("\nPrinting>%f", cl.m[ctx.GetPayload()])
}

// ExitAssign is called when production assign is exited.
func (cl *ComputeListener) ExitAssign(ctx *auto.AssignContext) {
	val := cl.m[ctx.Expr().GetPayload()]
	cl.ids[ctx.ID().GetText()] = val
	cl.m[ctx.GetPayload()] = val
}

// ExitDiv is called when exiting the Div production.
func (cl *ComputeListener) ExitDiv(ctx *auto.DivContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr(0).GetPayload()] / cl.m[ctx.Expr(1).GetPayload()]
}

// ExitMul is called when exiting the Mul production.
func (cl *ComputeListener) ExitMul(ctx *auto.MulContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr(0).GetPayload()] * cl.m[ctx.Expr(1).GetPayload()]
}

// ExitAdd is called when exiting the Add production.
func (cl *ComputeListener) ExitAdd(ctx *auto.AddContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr(0).GetPayload()] + cl.m[ctx.Expr(1).GetPayload()]
}

// ExitSub is called when exiting the Sub production.
func (cl *ComputeListener) ExitSub(ctx *auto.SubContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr(0).GetPayload()] - cl.m[ctx.Expr(1).GetPayload()]
}

// ExitParenth is called when exiting the () production.
func (cl *ComputeListener) ExitParenth(ctx *auto.ParenthContext) {
	cl.m[ctx.GetPayload()] = cl.m[ctx.Expr().GetPayload()]
}

// ExitId is called when exiting the Id production.
func (cl *ComputeListener) ExitId(ctx *auto.IdContext) { // nolint (should not change spelling !)
	cl.m[ctx.GetPayload()] = cl.ids[ctx.GetText()]
}

// ExitUnary is called when exiting the unary minus production.
func (cl *ComputeListener) ExitUnary(ctx *auto.UnaryContext) {
	cl.m[ctx.GetPayload()] = -cl.m[ctx.Expr().GetPayload()]
}

// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package myparser // Expr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterPrintExpr is called when entering the printExpr production.
	EnterPrintExpr(c *PrintExprContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterBlank is called when entering the blank production.
	EnterBlank(c *BlankContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParenth is called when entering the parenth production.
	EnterParenth(c *ParenthContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitPrintExpr is called when exiting the printExpr production.
	ExitPrintExpr(c *PrintExprContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitBlank is called when exiting the blank production.
	ExitBlank(c *BlankContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParenth is called when exiting the parenth production.
	ExitParenth(c *ParenthContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)
}

// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package auto // Expr
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

	// EnterDiv is called when entering the Div production.
	EnterDiv(c *DivContext)

	// EnterAdd is called when entering the Add production.
	EnterAdd(c *AddContext)

	// EnterSub is called when entering the Sub production.
	EnterSub(c *SubContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterParenth is called when entering the parenth production.
	EnterParenth(c *ParenthContext)

	// EnterUnary is called when entering the Unary production.
	EnterUnary(c *UnaryContext)

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

	// ExitDiv is called when exiting the Div production.
	ExitDiv(c *DivContext)

	// ExitAdd is called when exiting the Add production.
	ExitAdd(c *AddContext)

	// ExitSub is called when exiting the Sub production.
	ExitSub(c *SubContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitParenth is called when exiting the parenth production.
	ExitParenth(c *ParenthContext)

	// ExitUnary is called when exiting the Unary production.
	ExitUnary(c *UnaryContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)
}

// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package auto // Expr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by ExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by ExprParser#Div.
	VisitDiv(ctx *DivContext) interface{}

	// Visit a parse tree produced by ExprParser#Add.
	VisitAdd(ctx *AddContext) interface{}

	// Visit a parse tree produced by ExprParser#Sub.
	VisitSub(ctx *SubContext) interface{}

	// Visit a parse tree produced by ExprParser#Mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by ExprParser#parenth.
	VisitParenth(ctx *ParenthContext) interface{}

	// Visit a parse tree produced by ExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ExprParser#Unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by ExprParser#int.
	VisitInt(ctx *IntContext) interface{}
}

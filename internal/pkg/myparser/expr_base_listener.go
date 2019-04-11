// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package myparser // Expr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseExprListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseExprListener) ExitProg(ctx *ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseExprListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseExprListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseExprListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseExprListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseExprListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseExprListener) ExitBlank(ctx *BlankContext) {}

// EnterDiv is called when production Div is entered.
func (s *BaseExprListener) EnterDiv(ctx *DivContext) {}

// ExitDiv is called when production Div is exited.
func (s *BaseExprListener) ExitDiv(ctx *DivContext) {}

// EnterAdd is called when production Add is entered.
func (s *BaseExprListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production Add is exited.
func (s *BaseExprListener) ExitAdd(ctx *AddContext) {}

// EnterSub is called when production Sub is entered.
func (s *BaseExprListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production Sub is exited.
func (s *BaseExprListener) ExitSub(ctx *SubContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseExprListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseExprListener) ExitMul(ctx *MulContext) {}

// EnterParenth is called when production parenth is entered.
func (s *BaseExprListener) EnterParenth(ctx *ParenthContext) {}

// ExitParenth is called when production parenth is exited.
func (s *BaseExprListener) ExitParenth(ctx *ParenthContext) {}

// EnterId is called when production id is entered.
func (s *BaseExprListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseExprListener) ExitId(ctx *IdContext) {}

// EnterUnary is called when production Unary is entered.
func (s *BaseExprListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production Unary is exited.
func (s *BaseExprListener) ExitUnary(ctx *UnaryContext) {}

// EnterInt is called when production int is entered.
func (s *BaseExprListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseExprListener) ExitInt(ctx *IntContext) {}

package listeners

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xavier268/go-antlr4-demo/internal/pkg/auto"
)

//DumpListener will dump the rules and context
type DumpListener struct {
	*auto.BaseExprListener
	p *auto.ExprParser
}

//NewDumpListener creates a dump listener using the provided parser
func NewDumpListener(p *auto.ExprParser) *DumpListener {
	var d = new(DumpListener)
	d.p = p
	return d
}

//ExitEveryRule dump
func (dl *DumpListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf(">>>> EXITING ...\t")
	dl.dump(ctx)
}

//EnterEveryRule dump
func (dl *DumpListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf(">>>> ENTERING ...\t")
	dl.dump(ctx)
}

func (dl *DumpListener) dump(ctx antlr.ParserRuleContext) {
	fmt.Printf("context : %q, rule : %d\n", ctx.GetText(), ctx.GetRuleIndex())
	fmt.Println(ctx.ToStringTree(nil, dl.p))
}

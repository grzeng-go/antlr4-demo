package labelexpr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//todo 计算的方法还未实现
type evalListener struct {
}

func NewEvalListener() *evalListener {
	return &evalListener{}
}

// VisitTerminal is called when a terminal node is visited.
func (s *evalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *evalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *evalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *evalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *evalListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *evalListener) ExitProg(ctx *ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *evalListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *evalListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *evalListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *evalListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *evalListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *evalListener) ExitBlank(ctx *BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *evalListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *evalListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *evalListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *evalListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *evalListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *evalListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *evalListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *evalListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *evalListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *evalListener) ExitInt(ctx *IntContext) {}

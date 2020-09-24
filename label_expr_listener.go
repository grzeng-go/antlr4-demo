package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grzeng-go/antlr4-demo/antlr/labelexpr"
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
func (s *evalListener) EnterProg(ctx *labelexpr.ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *evalListener) ExitProg(ctx *labelexpr.ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *evalListener) EnterPrintExpr(ctx *labelexpr.PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *evalListener) ExitPrintExpr(ctx *labelexpr.PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *evalListener) EnterAssign(ctx *labelexpr.AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *evalListener) ExitAssign(ctx *labelexpr.AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *evalListener) EnterBlank(ctx *labelexpr.BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *evalListener) ExitBlank(ctx *labelexpr.BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *evalListener) EnterParens(ctx *labelexpr.ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *evalListener) ExitParens(ctx *labelexpr.ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *evalListener) EnterMulDiv(ctx *labelexpr.MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *evalListener) ExitMulDiv(ctx *labelexpr.MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *evalListener) EnterAddSub(ctx *labelexpr.AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *evalListener) ExitAddSub(ctx *labelexpr.AddSubContext) {}

// EnterId is called when production id is entered.
func (s *evalListener) EnterId(ctx *labelexpr.IdContext) {}

// ExitId is called when production id is exited.
func (s *evalListener) ExitId(ctx *labelexpr.IdContext) {}

// EnterInt is called when production int is entered.
func (s *evalListener) EnterInt(ctx *labelexpr.IntContext) {}

// ExitInt is called when production int is exited.
func (s *evalListener) ExitInt(ctx *labelexpr.IntContext) {}

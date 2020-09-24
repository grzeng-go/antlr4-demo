// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/labelexpr\LabelExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package labelexpr // LabelExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLabelExprListener is a complete listener for a parse tree produced by LabelExprParser.
type BaseLabelExprListener struct{}

var _ LabelExprListener = &BaseLabelExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLabelExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLabelExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLabelExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLabelExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseLabelExprListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseLabelExprListener) ExitProg(ctx *ProgContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseLabelExprListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseLabelExprListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseLabelExprListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseLabelExprListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseLabelExprListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseLabelExprListener) ExitBlank(ctx *BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseLabelExprListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseLabelExprListener) ExitParens(ctx *ParensContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseLabelExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseLabelExprListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseLabelExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseLabelExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseLabelExprListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseLabelExprListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *BaseLabelExprListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseLabelExprListener) ExitInt(ctx *IntContext) {}

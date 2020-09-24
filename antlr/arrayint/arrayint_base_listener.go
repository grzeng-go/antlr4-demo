// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/arrayint\ArrayInt.g4 by ANTLR 4.8. DO NOT EDIT.

package arrayint // ArrayInt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArrayIntListener is a complete listener for a parse tree produced by ArrayIntParser.
type BaseArrayIntListener struct{}

var _ ArrayIntListener = &BaseArrayIntListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArrayIntListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArrayIntListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArrayIntListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArrayIntListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInit is called when production init is entered.
func (s *BaseArrayIntListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseArrayIntListener) ExitInit(ctx *InitContext) {}

// EnterValue is called when production value is entered.
func (s *BaseArrayIntListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseArrayIntListener) ExitValue(ctx *ValueContext) {}

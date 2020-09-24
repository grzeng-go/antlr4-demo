// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/labelexpr\LabelExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package labelexpr // LabelExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LabelExprParser.
type LabelExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LabelExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by LabelExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by LabelExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by LabelExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by LabelExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by LabelExprParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by LabelExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by LabelExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by LabelExprParser#int.
	VisitInt(ctx *IntContext) interface{}
}

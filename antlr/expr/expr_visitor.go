// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\Expr.g4 by ANTLR 4.8. DO NOT EDIT.

package expr // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ExprParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by ExprParser#expr.
	VisitExpr(ctx *ExprContext) interface{}
}

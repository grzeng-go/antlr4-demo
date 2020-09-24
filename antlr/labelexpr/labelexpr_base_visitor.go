// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/labelexpr\LabelExpr.g4 by ANTLR 4.8. DO NOT EDIT.

package labelexpr // LabelExpr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLabelExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLabelExprVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLabelExprVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

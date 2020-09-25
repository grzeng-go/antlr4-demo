package labelexpr

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type evalVistor struct {
	LabelExprVisitor
	memory map[string]int
}

func NewEvalVistor() *evalVistor {
	vistor := new(evalVistor)
	vistor.memory = make(map[string]int)
	return vistor
}

func (v *evalVistor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *evalVistor) VisitProg(ctx *ProgContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
		v.Visit(ctx.Stat(i))
	}
	return 0
}

// Visit a parse tree produced by LabelExprParser#printExpr.
func (v *evalVistor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	value := v.Visit(ctx.Expr())
	fmt.Println(value)
	return value
}

// Visit a parse tree produced by LabelExprParser#assign.
func (v *evalVistor) VisitAssign(ctx *AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.memory[id] = value.(int)
	return value
}

// Visit a parse tree produced by LabelExprParser#blank.
func (v *evalVistor) VisitBlank(ctx *BlankContext) interface{} {
	return 0
}

// Visit a parse tree produced by LabelExprParser#parens.
func (v *evalVistor) VisitParens(ctx *ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

// Visit a parse tree produced by LabelExprParser#MulDiv.
func (v *evalVistor) VisitMulDiv(ctx *MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	if ctx.GetOp().GetTokenType() == LabelExprLexerMUL {
		return left.(int) * right.(int)
	} else {
		return left.(int) / right.(int)
	}
}

// Visit a parse tree produced by LabelExprParser#AddSub.
func (v *evalVistor) VisitAddSub(ctx *AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0))
	right := v.Visit(ctx.Expr(1))
	if ctx.GetOp().GetTokenType() == LabelExprLexerADD {
		return left.(int) + right.(int)
	} else {
		return left.(int) - right.(int)
	}
}

// Visit a parse tree produced by LabelExprParser#id.
func (v *evalVistor) VisitId(ctx *IdContext) interface{} {
	id := ctx.ID().GetText()
	if x, ok := v.memory[id]; ok {
		return x
	}
	return 0
}

// Visit a parse tree produced by LabelExprParser#int.
func (v *evalVistor) VisitInt(ctx *IntContext) interface{} {
	i, err := strconv.Atoi(ctx.INT().GetText())
	if err != nil {
		return 0
	}
	return i
}

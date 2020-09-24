// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\Hello.g4 by ANTLR 4.8. DO NOT EDIT.

package hello // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseHelloVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHelloVisitor) VisitR(ctx *RContext) interface{} {
	return v.VisitChildren(ctx)
}

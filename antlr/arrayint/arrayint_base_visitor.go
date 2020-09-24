// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/arrayint\ArrayInt.g4 by ANTLR 4.8. DO NOT EDIT.

package arrayint // ArrayInt

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseArrayIntVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseArrayIntVisitor) VisitInit(ctx *InitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseArrayIntVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

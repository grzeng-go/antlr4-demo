// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/arrayint\ArrayInt.g4 by ANTLR 4.8. DO NOT EDIT.

package arrayint // ArrayInt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ArrayIntParser.
type ArrayIntVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ArrayIntParser#init.
	VisitInit(ctx *InitContext) interface{}

	// Visit a parse tree produced by ArrayIntParser#value.
	VisitValue(ctx *ValueContext) interface{}
}

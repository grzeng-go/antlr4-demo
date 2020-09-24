// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\Hello.g4 by ANTLR 4.8. DO NOT EDIT.

package hello // Hello

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by HelloParser.
type HelloVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HelloParser#r.
	VisitR(ctx *RContext) interface{}
}

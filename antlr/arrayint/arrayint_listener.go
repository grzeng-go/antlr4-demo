// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr/arrayint\ArrayInt.g4 by ANTLR 4.8. DO NOT EDIT.

package arrayint // ArrayInt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ArrayIntListener is a complete listener for a parse tree produced by ArrayIntParser.
type ArrayIntListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}

package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grzeng-go/antlr4-demo/antlr/arrayint"
	"github.com/grzeng-go/antlr4-demo/antlr/expr"
	"github.com/grzeng-go/antlr4-demo/antlr/hello"
	"github.com/grzeng-go/antlr4-demo/antlr/labelexpr"
)

func main() {
	//helloDemo()
	//arrayIntDemo()
	//exprDemo()
	labelExprVistorDemo()
}

func helloDemo() {
	is := antlr.NewInputStream("hello abc xx")
	lexer := hello.NewHelloLexer(is)
	for _, t := range lexer.GetAllTokens() {
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

type arrayIntListener struct {
	*arrayint.BaseArrayIntListener
}

func (s *arrayIntListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	/*text := ctx.GetText()
	fmt.Println(text)*/
}

// EnterInit is called when production init is entered.
func (s *arrayIntListener) EnterInit(ctx *arrayint.InitContext) {
	/*values := ctx.AllValue()
	for _, v := range values {
		valuecontext := (*arrayint.ValueContext)(unsafe.Pointer(reflect.ValueOf(v).Pointer()))
		switch valuecontext.GetParser().GetTokenStream().LA(1) {
		case arrayint.ArrayIntParserINT:
			fmt.Println("int: ", valuecontext.INT().GetText())
		default:
			fmt.Println("init: ", valuecontext.Init().GetText())
		}
	}*/
}

func arrayIntDemo() {
	// 新建一个输入流
	input := antlr.NewInputStream("{99,{2,3},456}")
	// 新建一个词法分析器，处理输入流
	lexer := arrayint.NewArrayIntLexer(input)
	// 新建一个词法符号的缓冲区，用于存储词法分析器将生成的词法符号
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	// 新建一个语法分析器，处理词法符号缓冲区中的内容
	parser := arrayint.NewArrayIntParser(tokenStream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	// 针对init规则，开始语法分析
	tree := parser.Init()
	fmt.Println(tree.ToStringTree(nil, parser))
	// 使用默认的语法分析树遍历器，遍历init规则的词法分析树，并触发自定义的回调监听器
	antlr.ParseTreeWalkerDefault.Walk(new(arrayIntListener), tree)
}

func exprDemo() {
	inputStream := antlr.NewInputStream("a = 1\n b = 2 \n a+b*2\n")
	lexer := expr.NewExprLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := expr.NewExprParser(tokenStream)
	prog := parser.Prog()
	fmt.Println(prog.ToStringTree(nil, parser))
}

func labelExprVistorDemo() {
	inputStream := antlr.NewInputStream("a=1\n b=2\n a+b\n c=a+b\n c+2\n 1+(2-1)*3\n")
	lexer := labelexpr.NewLabelExprLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := labelexpr.NewLabelExprParser(tokenStream)
	prog := parser.Prog()
	fmt.Println(prog.ToStringTree(nil, parser))
	prog.Accept(NewEvalVistor())
}

func labelExprListenerDemo() {
	inputStream := antlr.NewInputStream("a=1\n b=2\n a+b\n c=a+b\n c+2\n 1+(2-1)*3\n")
	lexer := labelexpr.NewLabelExprLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := labelexpr.NewLabelExprParser(tokenStream)
	prog := parser.Prog()
	fmt.Println(prog.ToStringTree(nil, parser))
	antlr.ParseTreeWalkerDefault.Walk(NewEvalListener(), prog)
}

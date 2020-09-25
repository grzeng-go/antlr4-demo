package main

import (
	"fmt"
	antlr_resource "github.com/antlr/antlr4/doc/resources"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grzeng-go/antlr4-demo/antlr/arrayint"
	"github.com/grzeng-go/antlr4-demo/antlr/expr"
	"github.com/grzeng-go/antlr4-demo/antlr/hello"
	"github.com/grzeng-go/antlr4-demo/antlr/labelexpr"
	"github.com/grzeng-go/antlr4-demo/antlr/mysql"
	"github.com/grzeng-go/antlr4-demo/antlr/selectsql"
)

func main() {
	//helloDemo()
	//arrayIntDemo()
	//exprDemo()
	//labelExprVistorDemo()
	//mysqlDemo()
	selectSqlDemo()
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
	prog.Accept(labelexpr.NewEvalVistor())
}

func labelExprListenerDemo() {
	inputStream := antlr.NewInputStream("a=1\n b=2\n a+b\n c=a+b\n c+2\n 1+(2-1)*3\n")
	lexer := labelexpr.NewLabelExprLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := labelexpr.NewLabelExprParser(tokenStream)
	prog := parser.Prog()
	fmt.Println(prog.ToStringTree(nil, parser))
	antlr.ParseTreeWalkerDefault.Walk(labelexpr.NewEvalListener(), prog)
}

func mysqlDemo() {
	//stream := antlr.NewInputStream("select ship_power.gun_power, ship_info.*\nFROM\n\t(\n\t\tselect s.name as ship_name, sum(g.power) as gun_power, max(callibr) as max_callibr\n\t\tfrom\n\t\t\tships s inner join ships_guns sg on s.id = sg.ship_id inner join guns g on g.id = sg.guns_id\n\t\tgroup by s.name\n\t) ship_power\n\tinner join\n\t(\n\t\tselect s.name as ship_name, sc.class_name, sc.tonange, sc.max_length, sc.start_build, sc.max_guns_size\n\t\tfrom\n\t\t\tships s inner join ship_class sc on s.class_id = sc.id\n\t) ship_info using (ship_name);")
	stream := antlr.NewInputStream("select * from abc a where a.name = 'xyz' and exists(select a from user where id = '1')")
	changingStream := antlr_resource.NewCaseChangingStream(stream, true)
	lexer := mysql.NewMySqlLexer(changingStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := mysql.NewMySqlParser(tokenStream)
	tree := parser.SelectStatement()
	fmt.Println(tree.ToStringTree(nil, parser))
	tree.Accept(mysql.NewMysqlVisitor())
}

func selectSqlDemo() {
	//stream := antlr.NewInputStream("select ship_power.gun_power, ship_info.*\nFROM\n\t(\n\t\tselect s.name as ship_name, sum(g.power) as gun_power, max(callibr) as max_callibr\n\t\tfrom\n\t\t\tships s inner join ships_guns sg on s.id = sg.ship_id inner join guns g on g.id = sg.guns_id\n\t\tgroup by s.name\n\t) ship_power\n\tinner join\n\t(\n\t\tselect s.name as ship_name, sc.class_name, sc.tonange, sc.max_length, sc.start_build, sc.max_guns_size\n\t\tfrom\n\t\t\tships s inner join ship_class sc on s.class_id = sc.id\n\t) ship_info using (ship_name);")
	stream := antlr.NewInputStream("select * from abc a where a.name = 'xyz' and exists(select a from user where id = '1')")
	changingStream := antlr_resource.NewCaseChangingStream(stream, true)
	lexer := selectsql.NewMySqlLexer(changingStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := selectsql.NewMySqlParser(tokenStream)
	tree := parser.Root()
	fmt.Println(tree.ToStringTree(nil, parser))
	tree.Accept(selectsql.NewVisitor())
}

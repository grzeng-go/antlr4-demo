package mysql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

type (
	TableInfo struct {
		// 表别名
		Alias string
		// 替换位置
		AddMark string
	}

	SqlHolder struct {
		// sql别名-sql语句
		AliasMap map[string]string
		// sql别名-表名-位置
		Mark map[string]map[string]TableInfo
	}

	MysqlVisitor struct {
		BaseMySqlParserVisitor
		Holder SqlHolder
		Alias  string
		build  strings.Builder
	}
)

func NewMysqlVisitor() *MysqlVisitor {
	return &MysqlVisitor{
		/*Holder: SqlHolder{},
		Alias:  "",
		build:  strings.Builder{},*/
	}
}

func (mv *MysqlVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(mv)
}

func (mv *MysqlVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	return mv.Visit(ctx.QuerySpecification()).(string) + mv.Visit(ctx.LockClause()).(string)
}

func (mv *MysqlVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{} {
	return mv.Visit(ctx)
}

func (mv *MysqlVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{} {
	return ""
}

func (mv *MysqlVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return mv.Visit(ctx)
}

func (mv *MysqlVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{} {
	return mv.Visit(ctx)
}

func (mv *MysqlVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitUnionStatement(ctx *UnionStatementContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectSpec(ctx *SelectSpecContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	fmt.Println(ctx.TableSources().GetText())
	fmt.Println(ctx.GetText())
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MysqlVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return ctx.GetText()
}

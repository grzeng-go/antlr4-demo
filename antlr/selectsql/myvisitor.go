package selectsql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"regexp"
	"strings"
)

const WS = " "

type MyVisitor struct {
	// fromClause-table-alias
	tableAlias map[interface{}]map[string]string
}

func (mv *MyVisitor) VisitQuerySpecification1(ctx *QuerySpecification1Context) interface{} {
	var sb strings.Builder
	sb.WriteString(ctx.SELECT().GetText() + WS)
	selectSpecs := ctx.AllSelectSpec()
	if len(selectSpecs) > 0 {
		for _, selectspec := range selectSpecs {
			sb.WriteString(mv.Visit(selectspec).(string) + WS)
		}
	}
	sb.WriteString(mv.Visit(ctx.SelectElements()).(string) + WS)
	selectIntoExpression := ctx.SelectIntoExpression()
	if selectIntoExpression != nil {
		sb.WriteString(mv.Visit(selectIntoExpression).(string) + WS)
	}
	fromClause := ctx.FromClause()
	if fromClause != nil {
		sb.WriteString(mv.Visit(fromClause).(string) + WS)
	}

	orderByClause := ctx.OrderByClause()
	if orderByClause != nil {
		sb.WriteString(mv.Visit(orderByClause).(string) + WS)
	}

	limitClause := ctx.LimitClause()
	if limitClause != nil {
		sb.WriteString(mv.Visit(limitClause).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitQuerySpecification2(ctx *QuerySpecification2Context) interface{} {
	var sb strings.Builder
	sb.WriteString(ctx.SELECT().GetText() + WS)
	selectSpecs := ctx.AllSelectSpec()
	if len(selectSpecs) > 0 {
		for _, selectspec := range selectSpecs {
			sb.WriteString(mv.Visit(selectspec).(string) + WS)
		}
	}
	sb.WriteString(mv.Visit(ctx.SelectElements()).(string) + WS)

	fromClause := ctx.FromClause()
	if fromClause != nil {
		sb.WriteString(mv.Visit(fromClause).(string) + WS)
	}

	orderByClause := ctx.OrderByClause()
	if orderByClause != nil {
		sb.WriteString(mv.Visit(orderByClause).(string) + WS)
	}

	limitClause := ctx.LimitClause()
	if limitClause != nil {
		sb.WriteString(mv.Visit(limitClause).(string) + WS)
	}

	selectIntoExpression := ctx.SelectIntoExpression()
	if selectIntoExpression != nil {
		sb.WriteString(mv.Visit(selectIntoExpression).(string) + WS)
	}

	return sb.String()
}

func NewVisitor() *MyVisitor {
	visitor := new(MyVisitor)
	visitor.tableAlias = make(map[interface{}]map[string]string)
	return visitor
}

func (mv *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(mv)
}

func (mv *MyVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	return nil
}

func (mv *MyVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (mv *MyVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (mv *MyVisitor) VisitRoot(ctx *RootContext) interface{} {
	sql := mv.Visit(ctx.SelectStatement())
	fmt.Println(sql)
	reg := regexp.MustCompile("@([a-zA-Z_]*)\\.([a-zA-Z_]*)@")
	tables := reg.FindAllStringSubmatch(sql.(string), -1)
	for _, t := range tables {
		for i, tt := range t {
			switch i {
			case 0:
				println("需要替换的字符串：", tt)
			case 1:
				println("表名：", tt)
			case 2:
				println("表别名：", tt)
			}
		}
	}
	return sql
}

func (mv *MyVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	sql := mv.Visit(ctx.QuerySpecification()).(string)
	clause := ctx.LockClause()
	if clause != nil {
		sql += WS + mv.Visit(clause).(string)
	}
	return sql
}

func (mv *MyVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{} {
	// queryExpression lockClause?
	sql := mv.Visit(ctx.QueryExpression()).(string)
	clause := ctx.LockClause()
	if clause != nil {
		sql += WS + mv.Visit(clause).(string)
	}
	return sql
}

func (mv *MyVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	// querySpecificationNointo unionStatement+
	//        unionStatement2?
	//        orderByClause? limitClause? lockClause?
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.QuerySpecificationNointo()).(string) + WS)

	for _, us := range ctx.AllUnionStatement() {
		sb.WriteString(mv.Visit(us).(string) + WS)
	}

	if ctx.UnionStatement2() != nil {
		sb.WriteString(mv.Visit(ctx.UnionStatement2()).(string) + WS)
	}

	if ctx.OrderByClause() != nil {
		sb.WriteString(mv.Visit(ctx.OrderByClause()).(string) + WS)
	}

	if ctx.LimitClause() != nil {
		sb.WriteString(mv.Visit(ctx.LimitClause()).(string) + WS)
	}

	if ctx.LockClause() != nil {
		sb.WriteString(mv.Visit(ctx.LockClause()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{} {
	// queryExpressionNointo unionParenthesis+
	//        unionstatement3?
	//        orderByClause? limitClause? lockClause?
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.QueryExpressionNointo()).(string) + WS)

	for _, us := range ctx.AllUnionParenthesis() {
		sb.WriteString(mv.Visit(us).(string) + WS)
	}

	if ctx.Unionstatement3() != nil {
		sb.WriteString(mv.Visit(ctx.Unionstatement3()).(string) + WS)
	}

	if ctx.OrderByClause() != nil {
		sb.WriteString(mv.Visit(ctx.OrderByClause()).(string) + WS)
	}

	if ctx.LimitClause() != nil {
		sb.WriteString(mv.Visit(ctx.LimitClause()).(string) + WS)
	}

	if ctx.LockClause() != nil {
		sb.WriteString(mv.Visit(ctx.LockClause()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitAssignmentField(ctx *AssignmentFieldContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	if ctx.FOR() != nil {
		return " FOR UPDATE "
	} else {
		return " LOCK IN SHARE MODE "
	}
}

func (mv *MyVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	// ORDER BY orderByExpression (',' orderByExpression)*
	var strs []string
	for _, o := range ctx.AllOrderByExpression() {
		strs = append(strs, mv.Visit(o).(string))
	}
	return ctx.ORDER().GetText() + WS + ctx.BY().GetText() + WS + strings.Join(strs, ", ") + WS
}

func (mv *MyVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	//expression order=(ASC | DESC)?
	var order string
	if ctx.GetOrder() != nil {
		order = ctx.GetOrder().GetText()
	}
	return mv.Visit(ctx.Expression()).(string) + WS + order + WS
}

func (mv *MyVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	// tableSource (',' tableSource)*
	var strs []string
	for _, o := range ctx.AllTableSource() {
		strs = append(strs, mv.Visit(o).(string))
	}
	return strings.Join(strs, ", ") + WS
}

func (mv *MyVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	// tableSourceItem joinPart*
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	for _, join := range ctx.AllJoinPart() {
		sb.WriteString(mv.Visit(join).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	// '(' tableSourceItem joinPart* ')'
	var sb strings.Builder

	sb.WriteString("( ")

	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	for _, join := range ctx.AllJoinPart() {
		sb.WriteString(mv.Visit(join).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	// tableName
	//      (PARTITION '(' uidList ')' )? (AS? alias=uid)?
	//      (indexHint (',' indexHint)* )?
	var sb strings.Builder

	tableName := mv.Visit(ctx.TableName()).(string)
	sb.WriteString(tableName + WS)

	if ctx.PARTITION() != nil {
		sb.WriteString(ctx.PARTITION().GetText() + "(" + mv.Visit(ctx.UidList()).(string) + ") ")
	}

	var alias string
	if ctx.GetAlias() != nil {
		if ctx.AS() != nil {
			sb.WriteString(ctx.AS().GetText() + WS)
		}
		alias = mv.Visit(ctx.GetAlias()).(string)
		sb.WriteString(alias + WS)
	}

	doWithFromClause(ctx, func(node interface{}) {
		n, ok := node.(*FromClauseContext)
		if ok {
			//fmt.Println("---", n.GetText(), "; table: ", tableName, "; alias: ", alias)
			ta, ok := mv.tableAlias[n]
			if ok {
				ta[tableName] = alias
			} else {
				m := make(map[string]string)
				m[tableName] = alias
				mv.tableAlias[n] = m
			}
		}
	})

	hints := ctx.AllIndexHint()
	for _, hint := range hints {
		sb.WriteString(mv.Visit(hint).(string) + WS)
	}

	return sb.String()
}

func doWithFromClause(node antlr.Tree, fn func(node interface{})) {
	if node == nil {
		return
	}

	if isFromClauseNode(node) {
		fn(node)
	} else {
		doWithFromClause(node.GetParent(), fn)
	}
}

func isFromClauseNode(node interface{}) bool {
	if node == nil {
		return false
	}

	rv := reflect.ValueOf(node)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		_, ok := rv.Elem().Interface().(FromClauseContext)
		return ok
	}

	_, ok := rv.Interface().(FromClauseContext)
	return ok
}

func (mv *MyVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	// (
	//      selectStatement
	//      | '(' parenthesisSubquery=selectStatement ')'
	//      )
	//      AS? alias=uid
	var sb strings.Builder
	subquery := ctx.GetParenthesisSubquery()
	if subquery != nil {
		sb.WriteString("( ")
		sb.WriteString(mv.Visit(ctx.SelectStatement()).(string) + WS)
		sb.WriteString(") ")
	} else {
		sb.WriteString(mv.Visit(ctx.SelectStatement()).(string) + WS)
	}
	as := ctx.AS()
	if as != nil {
		sb.WriteString(as.GetText() + WS)
	}

	sb.WriteString(ctx.GetAlias().GetText() + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	// '(' tableSources ')'
	return "( " + mv.Visit(ctx.TableSources()).(string) + ") "
}

func (mv *MyVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	// indexHintAction=(USE | IGNORE | FORCE)
	//      keyFormat=(INDEX|KEY) ( FOR indexHintType)?
	//      '(' uidList ')'
	var sb strings.Builder

	sb.WriteString(ctx.GetIndexHintAction().GetText() + WS)
	sb.WriteString(ctx.GetKeyFormat().GetText() + WS)

	node := ctx.FOR()
	if node != nil {
		sb.WriteString(node.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.IndexHintType()).(string) + WS)
	}

	sb.WriteString("( ")
	sb.WriteString(mv.Visit(ctx.UidList()).(string) + WS)
	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	// JOIN | ORDER BY | GROUP BY
	join := ctx.JOIN()
	if join != nil {
		return join.GetText() + WS
	}

	order := ctx.ORDER()
	if order != nil {
		return order.GetText() + WS + ctx.BY().GetText() + WS
	}

	group := ctx.GROUP()
	if group != nil {
		return group.GetText() + WS + ctx.BY().GetText() + WS
	}

	return ctx.GetText()
}

func (mv *MyVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	// (INNER | CROSS)? JOIN tableSourceItem
	//      (
	//        ON expression
	//        | USING '(' uidList ')'
	//      )?
	var sb strings.Builder

	if ctx.INNER() != nil {
		sb.WriteString(ctx.INNER().GetText() + WS)
	}

	if ctx.CROSS() != nil {
		sb.WriteString(ctx.CROSS().GetText() + WS)
	}

	sb.WriteString(ctx.JOIN().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	if ctx.ON() != nil {
		sb.WriteString(ctx.ON().GetText() + WS)
		sb.WriteString(mv.Visit(ctx.Expression()).(string) + WS)
	}

	if ctx.USING() != nil {
		sb.WriteString(ctx.USING().GetText() + " (")
		sb.WriteString(mv.Visit(ctx.UidList()).(string) + ") ")
	}

	return sb.String()
}

func (mv *MyVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	// STRAIGHT_JOIN tableSourceItem (ON expression)?
	var sb strings.Builder

	sb.WriteString(ctx.STRAIGHT_JOIN().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	if ctx.ON() != nil {
		sb.WriteString(ctx.ON().GetText() + WS)
		sb.WriteString(mv.Visit(ctx.Expression()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	// (LEFT | RIGHT) OUTER? JOIN tableSourceItem
	//        (
	//          ON expression
	//          | USING '(' uidList ')'
	//        )
	var sb strings.Builder

	if ctx.LEFT() != nil {
		sb.WriteString(ctx.LEFT().GetText() + WS)
	}

	if ctx.RIGHT() != nil {
		sb.WriteString(ctx.RIGHT().GetText() + WS)
	}

	if ctx.OUTER() != nil {
		sb.WriteString(ctx.OUTER().GetText() + WS)
	}

	sb.WriteString(ctx.JOIN().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	if ctx.ON() != nil {
		sb.WriteString(ctx.ON().GetText() + WS)
		sb.WriteString(mv.Visit(ctx.Expression()).(string) + WS)
	}

	if ctx.USING() != nil {
		sb.WriteString(ctx.USING().GetText() + " (")
		sb.WriteString(mv.Visit(ctx.UidList()).(string) + ") ")
	}

	return sb.String()
}

func (mv *MyVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	// NATURAL ((LEFT | RIGHT) OUTER?)? JOIN tableSourceItem
	var sb strings.Builder

	sb.WriteString(ctx.NATURAL().GetText() + WS)

	if ctx.LEFT() != nil {
		sb.WriteString(ctx.LEFT().GetText() + WS)
	}

	if ctx.RIGHT() != nil {
		sb.WriteString(ctx.RIGHT().GetText() + WS)
	}

	if ctx.OUTER() != nil {
		sb.WriteString(ctx.OUTER().GetText() + WS)
	}

	sb.WriteString(ctx.JOIN().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.TableSourceItem()).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	// '(' querySpecification ')'
	//    | '(' queryExpression ')'
	var sb strings.Builder

	sb.WriteString("( ")

	if ctx.QuerySpecification() != nil {
		sb.WriteString(mv.Visit(ctx.QuerySpecification()).(string) + WS)
	}

	if ctx.QueryExpression() != nil {
		sb.WriteString(mv.Visit(ctx.QueryExpression()).(string) + WS)
	}

	sb.WriteString(") ")
	return sb.String()
}

func (mv *MyVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{} {
	// '(' querySpecificationNointo ')'
	//    | '(' queryExpressionNointo ')'
	var sb strings.Builder

	sb.WriteString("( ")

	if ctx.QueryExpressionNointo() != nil {
		sb.WriteString(mv.Visit(ctx.QueryExpressionNointo()).(string) + WS)
	}

	if ctx.QuerySpecificationNointo() != nil {
		sb.WriteString(mv.Visit(ctx.QuerySpecificationNointo()).(string) + WS)
	}

	sb.WriteString(") ")
	return sb.String()
}

func (mv *MyVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{} {
	// SELECT selectSpec* selectElements
	//      fromClause? orderByClause? limitClause?
	var sb strings.Builder
	sb.WriteString(ctx.SELECT().GetText() + WS)
	selectSpecs := ctx.AllSelectSpec()
	if len(selectSpecs) > 0 {
		for _, selectspec := range selectSpecs {
			sb.WriteString(mv.Visit(selectspec).(string) + WS)
		}
	}
	sb.WriteString(mv.Visit(ctx.SelectElements()).(string) + WS)

	fromClause := ctx.FromClause()
	if fromClause != nil {
		sb.WriteString(mv.Visit(fromClause).(string) + WS)
	}

	orderByClause := ctx.OrderByClause()
	if orderByClause != nil {
		sb.WriteString(mv.Visit(orderByClause).(string) + WS)
	}

	limitClause := ctx.LimitClause()
	if limitClause != nil {
		sb.WriteString(mv.Visit(limitClause).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{} {
	// UNION unionType=(ALL | DISTINCT)? queryExpressionNointo
	var sb strings.Builder

	sb.WriteString(ctx.UNION().GetText() + WS)

	if ctx.GetUnionType() != nil {
		sb.WriteString(ctx.GetUnionType().GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.QueryExpressionNointo()).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitUnionStatement(ctx *UnionStatementContext) interface{} {
	// UNION unionType=(ALL | DISTINCT)?
	//      (querySpecificationNointo | queryExpressionNointo)
	var sb strings.Builder

	sb.WriteString(ctx.UNION().GetText() + WS)

	if ctx.GetUnionType() != nil {
		sb.WriteString(ctx.GetUnionType().GetText() + WS)
	}

	if ctx.QueryExpressionNointo() != nil {
		sb.WriteString(mv.Visit(ctx.QueryExpressionNointo()).(string) + WS)
	}

	if ctx.QuerySpecificationNointo() != nil {
		sb.WriteString(mv.Visit(ctx.QuerySpecificationNointo()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitUnionStatement2(ctx *UnionStatement2Context) interface{} {
	// (
	//        UNION unionType=(ALL | DISTINCT)?
	//        (querySpecification | queryExpression)
	//      )
	var sb strings.Builder

	sb.WriteString(ctx.UNION().GetText() + WS)

	if ctx.GetUnionType() != nil {
		sb.WriteString(ctx.GetUnionType().GetText() + WS)
	}

	if ctx.QueryExpression() != nil {
		sb.WriteString(mv.Visit(ctx.QueryExpression()).(string) + WS)
	}

	if ctx.QuerySpecification() != nil {
		sb.WriteString(mv.Visit(ctx.QuerySpecification()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitUnionstatement3(ctx *Unionstatement3Context) interface{} {
	//(
	//        UNION unionType=(ALL | DISTINCT)?
	//        queryExpression
	//      )
	var sb strings.Builder

	sb.WriteString(ctx.UNION().GetText() + WS)

	if ctx.GetUnionType() != nil {
		sb.WriteString(ctx.GetUnionType().GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.QueryExpression()).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitSelectSpec(ctx *SelectSpecContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	star := ctx.STAR()
	if star != nil {
		return star.GetText()
	}

	var sb strings.Builder

	elements := ctx.AllSelectElement()
	if len(elements) > 0 {
		var eleStr []string
		for _, e := range elements {
			eleStr = append(eleStr, mv.Visit(e).(string))
		}
		sb.WriteString(strings.Join(eleStr, ", "))
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.FullColumnName()).(string) + WS)

	as := ctx.AS()
	if as != nil {
		sb.WriteString(mv.Visit(as).(string) + WS)
	}

	uid := ctx.Uid()
	if uid != nil {
		sb.WriteString(mv.Visit(uid).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.FunctionCall()).(string) + WS)

	as := ctx.AS()
	if as != nil {
		sb.WriteString(mv.Visit(as).(string) + WS)
	}

	uid := ctx.Uid()
	if uid != nil {
		sb.WriteString(mv.Visit(uid).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{} {
	var sb strings.Builder

	id := ctx.LOCAL_ID()
	if id != nil {
		sb.WriteString(mv.Visit(id).(string) + WS)
	}

	assign := ctx.VAR_ASSIGN()
	if assign != nil {
		sb.WriteString(mv.Visit(assign).(string) + WS)
	}

	sb.WriteString(mv.Visit(ctx.Expression()).(string) + WS)

	as := ctx.AS()
	if as != nil {
		sb.WriteString(mv.Visit(as).(string) + WS)
	}

	uid := ctx.Uid()
	if uid != nil {
		sb.WriteString(mv.Visit(uid).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.INTO().GetText() + WS)

	fields := ctx.AllAssignmentField()
	var fieldStr []string
	for _, field := range fields {
		fieldStr = append(fieldStr, mv.Visit(field).(string))
	}
	sb.WriteString(strings.Join(fieldStr, ", "))
	return sb.String()
}

func (mv *MyVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{} {
	var sb strings.Builder
	sb.WriteString(ctx.INTO().GetText())
	sb.WriteString(ctx.DUMPFILE().GetText())
	sb.WriteString(ctx.STRING_LITERAL().GetText())
	return sb.String()
}

func (mv *MyVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.INTO().GetText() + WS)
	sb.WriteString(ctx.OUTFILE().GetText() + WS)
	sb.WriteString(ctx.GetFilename().GetText() + WS)

	character := ctx.CHARACTER()
	if character != nil {
		sb.WriteString(character.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.SET()).(string) + WS)
		sb.WriteString(mv.Visit(ctx.GetCharset()).(string) + WS)
	}

	format := ctx.GetFieldsFormat()
	if format != nil {
		sb.WriteString(format.GetText() + WS)
		fieldsIntos := ctx.AllSelectFieldsInto()
		for _, fieldsInto := range fieldsIntos {
			sb.WriteString(mv.Visit(fieldsInto).(string) + WS)
		}
	}

	lines := ctx.LINES()
	if lines != nil {
		sb.WriteString(lines.GetText() + WS)
		linesIntos := ctx.AllSelectLinesInto()
		for _, linesInto := range linesIntos {
			sb.WriteString(mv.Visit(linesInto).(string) + WS)
		}
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	var sb strings.Builder

	terminated := ctx.TERMINATED()
	if terminated != nil {
		sb.WriteString(terminated.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		sb.WriteString(ctx.GetTerminationField().GetText() + WS)
	}

	enclosed := ctx.ENCLOSED()
	if enclosed != nil {
		optionally := ctx.OPTIONALLY()
		if optionally != nil {
			sb.WriteString(optionally.GetText() + WS)
		}
		sb.WriteString(enclosed.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		sb.WriteString(ctx.GetEnclosion().GetText() + WS)
	}

	escaped := ctx.ESCAPED()
	if escaped != nil {
		sb.WriteString(escaped.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		sb.WriteString(ctx.GetEscaping().GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	var sb strings.Builder

	starting := ctx.STARTING()
	if starting != nil {
		sb.WriteString(starting.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		sb.WriteString(ctx.GetStarting().GetText() + WS)
	}

	terminated := ctx.TERMINATED()
	if terminated != nil {
		sb.WriteString(terminated.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		sb.WriteString(ctx.GetTerminationLine().GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.FROM().GetText() + WS)

	sb.WriteString(mv.Visit(ctx.TableSources()).(string) + WS)

	where := ctx.WHERE()
	if where != nil {
		sb.WriteString(where.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.GetWhereExpr()).(string) + WS)

	} else {
		sb.WriteString(" WHERE 1 = 1 ")
	}

	tables, ok := mv.tableAlias[ctx]
	if ok {
		for table, alias := range tables {
			sb.WriteString(fmt.Sprintf(" @%v.%v@ ", table, alias))
		}
	}

	group := ctx.GROUP()
	if group != nil {
		sb.WriteString(group.GetText() + WS)
		sb.WriteString(ctx.BY().GetText() + WS)
		items := ctx.AllGroupByItem()
		var itemStr []string
		for _, item := range items {
			itemStr = append(itemStr, mv.Visit(item).(string))
		}
		sb.WriteString(strings.Join(itemStr, ", "))
		with := ctx.WITH()
		if with != nil {
			sb.WriteString(with.GetText() + WS)
			sb.WriteString(ctx.ROLLUP().GetText() + WS)
		}
	}

	having := ctx.HAVING()
	if having != nil {
		sb.WriteString(having.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.GetHavingExpr()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Expression()).(string) + WS)

	order := ctx.GetOrder()
	if order != nil {
		sb.WriteString(order.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.LIMIT().GetText() + WS)

	offset := ctx.OFFSET()
	if offset != nil {
		sb.WriteString(ctx.GetLimit().GetText() + WS + offset.GetText() + WS + ctx.GetOffset().GetText() + WS)
	} else {
		if ctx.GetOffset() != nil {
			sb.WriteString(ctx.GetOffset().GetText() + ", ")
		}
		sb.WriteString(ctx.GetLimit().GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFullId(ctx *FullIdContext) interface{} {
	var sb strings.Builder
	dotId := ctx.DOT_ID()
	if dotId != nil {
		sb.WriteString(ctx.Uid(0).GetText() + WS)
		sb.WriteString(dotId.GetText() + WS)
	} else {
		if len(ctx.AllUid()) > 1 {
			sb.WriteString(ctx.Uid(0).GetText() + "." + ctx.Uid(1).GetText() + WS)
		} else {
			return ctx.GetText()
		}
	}

	return sb.String()
}

func (mv *MyVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return mv.Visit(ctx.FullId())
}

func (mv *MyVisitor) VisitFullColumnName(ctx *FullColumnNameContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Uid()).(string) + WS)

	ids := ctx.AllDottedId()
	if len(ids) > 0 {
		for _, id := range ids {
			sb.WriteString(mv.Visit(id).(string) + WS)
		}
	}

	return sb.String()
}

func (mv *MyVisitor) VisitMysqlVariable(ctx *MysqlVariableContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUid(ctx *UidContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitEngineName(ctx *EngineNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSimpleId(ctx *SimpleIdContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDottedId(ctx *DottedIdContext) interface{} {
	var sb strings.Builder

	dotId := ctx.DOT_ID()
	if dotId != nil {
		sb.WriteString(dotId.GetText() + WS)
	} else {
		sb.WriteString("." + ctx.Uid().GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitStringLiteral1(ctx *StringLiteral1Context) interface{} {
	var sb strings.Builder

	literal := ctx.START_NATIONAL_STRING_LITERAL()
	if literal != nil {
		sb.WriteString(mv.Visit(literal).(string) + WS)
	} else {
		charsetName := ctx.STRING_CHARSET_NAME()
		if charsetName != nil {
			sb.WriteString(mv.Visit(charsetName).(string) + WS)
		}
	}

	stringLiterals := ctx.AllSTRING_LITERAL()
	for _, stringLiteral := range stringLiterals {
		sb.WriteString(stringLiteral.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitStringLiteral2(ctx *StringLiteral2Context) interface{} {
	var sb strings.Builder

	literal := ctx.START_NATIONAL_STRING_LITERAL()
	if literal != nil {
		sb.WriteString(mv.Visit(literal).(string) + WS)
	} else {
		charsetName := ctx.STRING_CHARSET_NAME()
		if charsetName != nil {
			sb.WriteString(charsetName.GetText() + WS)
		}
		sb.WriteString(ctx.STRING_LITERAL().GetText() + WS)
	}

	collate := ctx.COLLATE()
	if collate != nil {
		sb.WriteString(collate.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.CollationName()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitNullNotnull(ctx *NullNotnullContext) interface{} {
	var sb strings.Builder

	not := ctx.NOT()
	if not != nil {
		sb.WriteString(not.GetText() + WS)
	}

	literal := ctx.NULL_LITERAL()
	if literal != nil {
		sb.WriteString(literal.GetText() + WS)
	}

	specLiteral := ctx.NULL_SPEC_LITERAL()
	if specLiteral != nil {
		sb.WriteString(specLiteral.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{} {
	var sb strings.Builder

	charsetName := ctx.STRING_CHARSET_NAME()
	if charsetName != nil {
		sb.WriteString(charsetName.GetText() + WS)
	}

	sb.WriteString(ctx.HEXADECIMAL_LITERAL().GetText() + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	var sb strings.Builder

	stringLiteral := ctx.StringLiteral()
	if stringLiteral != nil {
		sb.WriteString(mv.Visit(stringLiteral).(string) + WS)
	}

	decimalLiteral := ctx.DecimalLiteral()
	booleanLiteral := ctx.BooleanLiteral()
	realLiteral := ctx.REAL_LITERAL()
	bitString := ctx.BIT_STRING()
	if decimalLiteral != nil || booleanLiteral != nil || realLiteral != nil || bitString != nil {
		sb.WriteString(ctx.GetText())
	}

	hexadecimalLiteral := ctx.HexadecimalLiteral()
	if hexadecimalLiteral != nil {
		sb.WriteString(mv.Visit(hexadecimalLiteral).(string) + WS)
	}

	notnull := ctx.NullNotnull()
	if notnull != nil {
		sb.WriteString(mv.Visit(notnull).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitStringDataType(ctx *StringDataTypeContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	binary := ctx.BINARY()
	if binary != nil {
		sb.WriteString(binary.GetText() + WS)
	}

	charsetName := ctx.CharsetName()
	if charsetName != nil {
		charset := ctx.CHARSET()
		if charset != nil {
			sb.WriteString(charset.GetText() + WS)
		}

		set := ctx.SET()
		characters := ctx.AllCHARACTER()
		if set != nil {
			sb.WriteString(characters[len(characters)-1].GetText() + WS)
			sb.WriteString(set.GetText() + WS)
		}

		sb.WriteString(charsetName.GetText() + WS)
	}

	collate := ctx.COLLATE()
	if collate != nil {
		sb.WriteString(collate.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.CollationName()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{} {
	var sb strings.Builder

	if ctx.NATIONAL() != nil {
		sb.WriteString(ctx.NATIONAL().GetText() + WS)
	}

	nchar := ctx.NCHAR()
	if nchar != nil {
		sb.WriteString(nchar.GetText() + WS)
	}

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	binary := ctx.BINARY()
	if binary != nil {
		sb.WriteString(binary.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.NATIONAL().GetText() + WS)

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	sb.WriteString(ctx.VARYING().GetText() + WS)

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	binary := ctx.BINARY()
	if binary != nil {
		sb.WriteString(binary.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{} {
	var sb strings.Builder

	double := ctx.DOUBLE()
	if double != nil {
		sb.WriteString(double.GetText() + WS)
		precision := ctx.PRECISION()
		if precision != nil {
			sb.WriteString(precision.GetText() + WS)
		}
	} else {
		sb.WriteString(ctx.GetTypeName().GetText() + WS)
	}

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	signed := ctx.SIGNED()
	if signed != nil {
		sb.WriteString(signed.GetText() + WS)
	}

	unsigned := ctx.UNSIGNED()
	if unsigned != nil {
		sb.WriteString(unsigned.GetText() + WS)
	}

	zerofill := ctx.ZEROFILL()
	if zerofill != nil {
		sb.WriteString(zerofill.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	collectionOptions := ctx.CollectionOptions()
	if collectionOptions != nil {
		sb.WriteString(mv.Visit(collectionOptions).(string) + WS)
	}

	binary := ctx.BINARY()
	if binary != nil {
		sb.WriteString(binary.GetText() + WS)
	}

	charsetName := ctx.CharsetName()
	if charsetName != nil {
		charset := ctx.CHARSET()
		if charset != nil {
			sb.WriteString(charset.GetText() + WS)
		}

		character := ctx.CHARACTER()
		if character != nil {
			sb.WriteString(character.GetText() + WS)
			sb.WriteString(ctx.SET(1).GetText() + WS)
		}

		sb.WriteString(charsetName.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	varchar := ctx.VARCHAR()
	if varchar != nil {
		sb.WriteString(varchar.GetText() + WS)
	}

	binary := ctx.BINARY()
	if binary != nil {
		sb.WriteString(binary.GetText() + WS)
	}

	charsetName := ctx.CharsetName()
	if charsetName != nil {
		charset := ctx.CHARSET()
		if charset != nil {
			sb.WriteString(charset.GetText() + WS)
		}

		character := ctx.CHARACTER()
		if character != nil {
			sb.WriteString(character.GetText() + WS)
			sb.WriteString(ctx.SET().GetText() + WS)
		}

		sb.WriteString(charsetName.GetText() + WS)
	}

	collate := ctx.COLLATE()
	if collate != nil {
		sb.WriteString(collate.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.CollationName()).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{} {
	return ctx.LONG().GetText() + WS + ctx.VARBINARY().GetText() + WS
}

func (mv *MyVisitor) VisitCollectionOptions(ctx *CollectionOptionsContext) interface{} {
	var sb strings.Builder

	var strs []string
	literals := ctx.AllSTRING_LITERAL()
	for _, literal := range literals {
		strs = append(strs, literal.GetText())
	}

	sb.WriteString("(")
	sb.WriteString(strings.Join(strs, ","))
	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitConvertedDataType1(ctx *ConvertedDataType1Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitConvertedDataType2(ctx *ConvertedDataType2Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	lengthOneDimension := ctx.LengthOneDimension()
	if lengthOneDimension != nil {
		sb.WriteString(mv.Visit(lengthOneDimension).(string) + WS)
	}

	charsetName := ctx.CharsetName()
	if charsetName != nil {
		charset := ctx.CHARSET()
		if charset != nil {
			sb.WriteString(charset.GetText() + WS)
		}

		character := ctx.CHARACTER()
		if character != nil {
			sb.WriteString(character.GetText() + WS)
			sb.WriteString(ctx.SET().GetText() + WS)
		}

		sb.WriteString(charsetName.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitConvertedDataType3(ctx *ConvertedDataType3Context) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitConvertedDataType4(ctx *ConvertedDataType4Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetTypeName().GetText() + WS)

	lengthTwoDimension := ctx.LengthTwoDimension()
	if lengthTwoDimension != nil {
		sb.WriteString(mv.Visit(lengthTwoDimension).(string) + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitConvertedDataType5(ctx *ConvertedDataType5Context) interface{} {
	var sb strings.Builder

	signed := ctx.SIGNED()
	if signed != nil {
		sb.WriteString(signed.GetText() + WS)
	}

	unsigned := ctx.UNSIGNED()
	if unsigned != nil {
		sb.WriteString(unsigned.GetText() + WS)
	}

	integer := ctx.INTEGER()
	if integer != nil {
		sb.WriteString(integer.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{} {
	return "(" + ctx.DecimalLiteral().GetText() + ") "
}

func (mv *MyVisitor) VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{} {
	var strs []string
	literals := ctx.AllDecimalLiteral()
	for _, literal := range literals {
		strs = append(strs, literal.GetText())
	}

	return "(" + strings.Join(strs, ", ") + ") "
}

func (mv *MyVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{} {
	var strs []string
	literals := ctx.AllDecimalLiteral()
	for _, literal := range literals {
		strs = append(strs, literal.GetText())
	}

	return "(" + strings.Join(strs, ", ") + ") "
}

func (mv *MyVisitor) VisitUidList(ctx *UidListContext) interface{} {
	var strs []string
	uids := ctx.AllUid()
	for _, uid := range uids {
		strs = append(strs, uid.GetText())
	}

	return "(" + strings.Join(strs, ", ") + ") "
}

func (mv *MyVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	var strs []string
	expressions := ctx.AllExpression()
	for _, expression := range expressions {
		strs = append(strs, mv.Visit(expression).(string))
	}

	return strings.Join(strs, ", ") + WS
}

func (mv *MyVisitor) VisitIntervalType(ctx *IntervalTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{} {
	return mv.Visit(ctx.SpecificFunction())
}

func (mv *MyVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{} {
	return mv.Visit(ctx.AggregateWindowedFunction())
}

func (mv *MyVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.ScalarFunctionName()).(string) + WS)
	sb.WriteString("(")

	args := ctx.FunctionArgs()
	if args != nil {
		sb.WriteString(mv.Visit(args).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.FullId()).(string) + WS)
	sb.WriteString("(")

	args := ctx.FunctionArgs()
	if args != nil {
		sb.WriteString(mv.Visit(args).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{} {
	return mv.Visit(ctx.PasswordFunctionClause())
}

func (mv *MyVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{} {
	var sb strings.Builder

	if ctx.CONVERT() != nil {
		sb.WriteString(ctx.CONVERT().GetText() + "(")
	}

	if ctx.CAST() != nil {
		sb.WriteString(ctx.CAST().GetText() + "(")
	}

	sb.WriteString(mv.Visit(ctx.Expression()).(string))

	if ctx.GetSeparator() != nil {
		sb.WriteString(ctx.GetSeparator().GetText() + WS)
	}

	if ctx.USING() != nil {
		sb.WriteString(ctx.USING().GetText() + WS)
	}

	if ctx.AS() != nil {
		sb.WriteString(ctx.AS().GetText() + WS)
	}

	if ctx.CharsetName() != nil {
		sb.WriteString(ctx.CharsetName().GetText() + WS)
	}

	if ctx.ConvertedDataType() != nil {
		sb.WriteString(mv.Visit(ctx.ConvertedDataType()).(string) + ") ")
	}

	return sb.String()
}

func (mv *MyVisitor) VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{} {
	return ctx.VALUES().GetText() + "(" + mv.Visit(ctx.FullColumnName()).(string) + ") "
}

func (mv *MyVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.CASE().GetText() + WS)

	expression := ctx.Expression()
	if expression != nil {
		sb.WriteString(mv.Visit(expression).(string) + WS)
	}

	caseFuncAlternatives := ctx.AllCaseFuncAlternative()
	for _, alternative := range caseFuncAlternatives {
		sb.WriteString(mv.Visit(alternative).(string) + WS)
	}

	node := ctx.ELSE()
	if node != nil {
		sb.WriteString(node.GetText() + WS)
		sb.WriteString(mv.Visit(ctx.GetElseArg()).(string) + WS)
	}

	sb.WriteString(ctx.END().GetText() + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.CHAR().GetText() + "(")

	sb.WriteString(mv.Visit(ctx.FunctionArgs()).(string) + WS)

	using := ctx.USING()
	if using != nil {
		sb.WriteString(using.GetText() + WS)
		sb.WriteString(ctx.CharsetName().GetText() + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.POSITION().GetText() + "( ")

	positionString := ctx.GetPositionString()
	if positionString != nil {
		sb.WriteString(mv.Visit(positionString).(string) + WS)
	}

	positionExpression := ctx.GetPositionExpression()
	if positionString != nil {
		sb.WriteString(mv.Visit(positionExpression).(string) + WS)
	}

	sb.WriteString(ctx.IN().GetText() + WS)

	inString := ctx.GetInString()
	if positionString != nil {
		sb.WriteString(mv.Visit(inString).(string) + WS)
	}

	inExpression := ctx.GetInExpression()
	if positionString != nil {
		sb.WriteString(mv.Visit(inExpression).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GetSub().GetText() + "( ")

	sourceString := ctx.GetSourceString()
	if sourceString != nil {
		sb.WriteString(mv.Visit(sourceString).(string) + WS)
	}

	sourceExpression := ctx.GetSourceExpression()
	if sourceExpression != nil {
		sb.WriteString(mv.Visit(sourceExpression).(string) + WS)
	}

	sb.WriteString(ctx.FROM().GetText() + WS)

	fromDecimal := ctx.GetFromDecimal()
	if fromDecimal != nil {
		sb.WriteString(mv.Visit(fromDecimal).(string) + WS)
	}

	fromExpression := ctx.GetFromExpression()
	if fromExpression != nil {
		sb.WriteString(mv.Visit(fromExpression).(string) + WS)
	}

	forNode := ctx.FOR()
	if forNode != nil {
		sb.WriteString(forNode.GetText() + WS)

		forDecimal := ctx.GetForDecimal()
		if forDecimal != nil {
			sb.WriteString(mv.Visit(forDecimal).(string) + WS)
		}

		forExpression := ctx.GetForExpression()
		if forExpression != nil {
			sb.WriteString(mv.Visit(forExpression).(string) + WS)
		}
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.TRIM().GetText() + "( ")

	positionForm := ctx.GetPositioinForm()
	if positionForm != nil {
		sb.WriteString(positionForm.GetText() + WS)
	}

	sourceString := ctx.GetSourceString()
	if sourceString != nil {
		sb.WriteString(mv.Visit(sourceString).(string) + WS)
	}

	sourceExpression := ctx.GetSourceExpression()
	if sourceExpression != nil {
		sb.WriteString(mv.Visit(sourceExpression).(string) + WS)
	}

	sb.WriteString(ctx.FROM().GetText() + WS)

	fromString := ctx.GetFromString()
	if fromString != nil {
		sb.WriteString(mv.Visit(fromString).(string) + WS)
	}

	fromExpression := ctx.GetFromExpression()
	if fromExpression != nil {
		sb.WriteString(mv.Visit(fromExpression).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.WEIGHT_STRING().GetText() + "( ")

	literal := ctx.StringLiteral()
	if literal != nil {
		sb.WriteString(mv.Visit(literal).(string) + WS)
	}

	expression := ctx.Expression()
	if expression != nil {
		sb.WriteString(mv.Visit(expression).(string) + WS)
	}

	asNode := ctx.AS()
	if asNode != nil {
		sb.WriteString(asNode.GetText() + WS)
		sb.WriteString(ctx.GetStringFormat().GetText() + "( ")
		sb.WriteString(ctx.DecimalLiteral().GetText() + ") ")
	}

	weightString := ctx.LevelsInWeightString()
	if weightString != nil {
		sb.WriteString(mv.Visit(weightString).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.EXTRACT().GetText() + "( " + ctx.IntervalType().GetText() + WS + ctx.FROM().GetText() + WS)

	sourceString := ctx.GetSourceString()
	if sourceString != nil {
		sb.WriteString(mv.Visit(sourceString).(string) + WS)
	}

	sourceExpression := ctx.GetSourceExpression()
	if sourceExpression != nil {
		sb.WriteString(mv.Visit(sourceExpression).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{} {
	return ctx.GET_FORMAT().GetText() + "( " + ctx.GetDatetimeFormat().GetText() + ", " + mv.Visit(ctx.StringLiteral()).(string) + ") "
}

func (mv *MyVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{} {
	return ctx.WHEN().GetText() + WS + mv.Visit(ctx.condition).(string) + WS + ctx.THEN().GetText() + mv.Visit(ctx.consequent).(string) + WS
}

func (mv *MyVisitor) VisitLevelWeightList(ctx *LevelWeightListContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.LEVEL().GetText() + WS)

	var strs []string
	elements := ctx.AllLevelInWeightListElement()
	for _, e := range elements {
		strs = append(strs, mv.Visit(e).(string))
	}
	sb.WriteString(strings.Join(strs, ", ") + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{} {
	return ctx.LEVEL().GetText() + WS + ctx.firstLevel.GetText() + " - " + ctx.lastLevel.GetText() + WS
}

func (mv *MyVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.DecimalLiteral().GetText() + WS)

	orderType := ctx.orderType
	if orderType != nil {
		sb.WriteString(orderType.GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction1(ctx *AggregateWindowedFunction1Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.operateName.GetText() + "( ")

	aggregator := ctx.aggregator
	if aggregator != nil {
		sb.WriteString(aggregator.GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.FunctionArg()).(string) + WS)

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction2(ctx *AggregateWindowedFunction2Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.COUNT().GetText() + "( ")

	if ctx.starArg != nil {
		sb.WriteString(ctx.starArg.GetText() + WS)
	} else {
		aggregator := ctx.aggregator
		if aggregator != nil {
			sb.WriteString(aggregator.GetText() + WS)
		}

		sb.WriteString(mv.Visit(ctx.FunctionArg()).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction3(ctx *AggregateWindowedFunction3Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.COUNT().GetText() + "( ")

	sb.WriteString(ctx.aggregator.GetText() + WS)

	sb.WriteString(mv.Visit(ctx.FunctionArgs()).(string) + WS)

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction4(ctx *AggregateWindowedFunction4Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.linkName.GetText() + "( ")

	aggregator := ctx.aggregator
	if aggregator != nil {
		sb.WriteString(aggregator.GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.FunctionArg()).(string) + WS)

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction5(ctx *AggregateWindowedFunction5Context) interface{} {
	var sb strings.Builder

	sb.WriteString(ctx.GROUP_CONCAT().GetText() + "( ")

	aggregator := ctx.aggregator
	if aggregator != nil {
		sb.WriteString(aggregator.GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.FunctionArgs()).(string) + WS)

	order := ctx.ORDER()
	if order != nil {
		sb.WriteString(order.GetText() + WS + ctx.BY().GetText() + WS)
		var strs []string
		expressions := ctx.AllOrderByExpression()
		for _, e := range expressions {
			strs = append(strs, mv.Visit(e).(string))
		}
		sb.WriteString(strings.Join(strs, ", ") + WS)
	}

	separator := ctx.SEPARATOR()
	if separator != nil {
		sb.WriteString(separator.GetText() + WS + ctx.separator.GetText() + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{} {
	return ctx.functionName.GetText() + "( " + mv.Visit(ctx.FunctionArg()).(string) + ") "
}

func (mv *MyVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	var strs []string
	args := ctx.AllFunctionArg()
	for _, arg := range args {
		strs = append(strs, mv.Visit(arg).(string))
	}
	return strings.Join(strs, ", ") + WS
}

func (mv *MyVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	// constant | fullColumnName | functionCall | expression
	constant := ctx.Constant()
	if constant != nil {
		return mv.Visit(constant)
	}

	fullColumnName := ctx.FullColumnName()
	if fullColumnName != nil {
		return mv.Visit(fullColumnName)
	}

	functionCall := ctx.FunctionCall()
	if functionCall != nil {
		return mv.Visit(functionCall)
	}

	expression := ctx.Expression()
	if expression != nil {
		return mv.Visit(expression)
	}

	return ""
}

func (mv *MyVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	//predicate IS NOT? testValue
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Predicate()).(string) + WS + ctx.IS().GetText() + WS)

	if ctx.NOT() != nil {
		sb.WriteString(ctx.NOT().GetText() + WS)
	}

	sb.WriteString(ctx.GetTestValue().GetText() + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return ctx.notOperator.GetText() + WS + mv.Visit(ctx.Expression()).(string) + WS
}

func (mv *MyVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return mv.Visit(ctx.Expression(0)).(string) + WS + ctx.LogicalOperator().GetText() + WS + mv.Visit(ctx.Expression(1)).(string) + WS
}

func (mv *MyVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return mv.Visit(ctx.Predicate())
}

func (mv *MyVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{} {
	return mv.Visit(ctx.Predicate(0)).(string) + WS + ctx.SOUNDS().GetText() + WS + ctx.LIKE().GetText() + WS + mv.Visit(ctx.Predicate(1)).(string) + WS
}

func (mv *MyVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	//(LOCAL_ID VAR_ASSIGN)? expressionAtom
	var sb strings.Builder

	if ctx.LOCAL_ID() != nil {
		sb.WriteString(ctx.LOCAL_ID().GetText() + WS + ctx.VAR_ASSIGN().GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.ExpressionAtom()).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{} {
	// predicate MEMBER OF '(' predicate ')'
	return mv.Visit(ctx.Predicate(0)).(string) + WS + ctx.MEMBER().GetText() + WS + ctx.OF().GetText() + "(" + mv.Visit(ctx.Predicate(1)).(string) + ") "
}

func (mv *MyVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	// predicate NOT? IN '(' (selectStatement | expressions) ')'
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Predicate()).(string) + WS)

	if ctx.NOT() != nil {
		sb.WriteString(ctx.NOT().GetText() + WS)
	}

	sb.WriteString(ctx.IN().GetText() + "(")

	if ctx.SelectStatement() != nil {
		sb.WriteString(mv.Visit(ctx.SelectStatement()).(string) + WS)
	}

	if ctx.Expressions() != nil {
		sb.WriteString(mv.Visit(ctx.Expressions()).(string) + WS)
	}

	sb.WriteString(") ")

	return sb.String()
}

func (mv *MyVisitor) VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{} {
	// predicate comparisonOperator
	//      quantifier=(ALL | ANY | SOME) '(' selectStatement ')'
	return mv.Visit(ctx.Predicate()).(string) + WS + ctx.ComparisonOperator().GetText() + WS + ctx.quantifier.GetText() + "(" + mv.Visit(ctx.SelectStatement()).(string) + ") "
}

func (mv *MyVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	// predicate NOT? BETWEEN predicate AND predicate
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Predicate(0)).(string) + WS)

	if ctx.NOT() != nil {
		sb.WriteString(ctx.NOT().GetText() + WS)
	}

	sb.WriteString(ctx.BETWEEN().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.Predicate(1)).(string) + WS)
	sb.WriteString(ctx.AND().GetText() + WS)
	sb.WriteString(mv.Visit(ctx.Predicate(2)).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	// left=predicate comparisonOperator right=predicate
	return mv.Visit(ctx.GetLeft()).(string) + WS + ctx.ComparisonOperator().GetText() + WS + mv.Visit(ctx.right).(string) + WS
}

func (mv *MyVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	// predicate IS nullNotnull
	return mv.Visit(ctx.Predicate()).(string) + WS + ctx.IS().GetText() + WS + mv.Visit(ctx.NullNotnull()).(string) + WS
}

func (mv *MyVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	// predicate NOT? LIKE predicate (ESCAPE STRING_LITERAL)?
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Predicate(0)).(string) + WS)

	if ctx.NOT() != nil {
		sb.WriteString(ctx.NOT().GetText() + WS)
	}

	sb.WriteString(ctx.LIKE().GetText() + WS)

	sb.WriteString(mv.Visit(ctx.Predicate(1)).(string) + WS)

	if ctx.ESCAPE() != nil {
		sb.WriteString(ctx.ESCAPE().GetText() + WS)
	}

	if ctx.STRING_LITERAL() != nil {
		sb.WriteString(ctx.STRING_LITERAL().GetText() + WS)
	}

	return sb.String()
}

func (mv *MyVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	// predicate NOT? regex=(REGEXP | RLIKE) predicate
	var sb strings.Builder

	sb.WriteString(mv.Visit(ctx.Predicate(0)).(string) + WS)

	if ctx.NOT() != nil {
		sb.WriteString(ctx.NOT().GetText() + WS)
	}

	if ctx.GetRegex() != nil {
		sb.WriteString(ctx.GetRegex().GetText() + WS)
	}

	sb.WriteString(mv.Visit(ctx.Predicate(1)).(string) + WS)

	return sb.String()
}

func (mv *MyVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	// unaryOperator expressionAtom
	return ctx.UnaryOperator().GetText() + WS + mv.Visit(ctx.ExpressionAtom()).(string) + WS
}

func (mv *MyVisitor) VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{} {
	// expressionAtom COLLATE collationName
	return mv.Visit(ctx.ExpressionAtom()).(string) + WS + ctx.COLLATE().GetText() + WS + mv.Visit(ctx.CollationName()).(string) + WS
}

func (mv *MyVisitor) VisitSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) interface{} {
	// '(' selectStatement ')'
	return "( " + mv.Visit(ctx.SelectStatement()).(string) + ") "
}

func (mv *MyVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{} {
	return mv.Visit(ctx.MysqlVariable()).(string) + WS
}

func (mv *MyVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	// '(' expression (',' expression)* ')'
	var strs []string
	for _, e := range ctx.AllExpression() {
		strs = append(strs, mv.Visit(e).(string))
	}
	return "( " + strings.Join(strs, ", ") + ") "
}

func (mv *MyVisitor) VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{} {
	// ROW '(' expression (',' expression)+ ')'
	var strs []string
	for _, e := range ctx.AllExpression() {
		strs = append(strs, mv.Visit(e).(string))
	}
	return ctx.ROW().GetText() + " ( " + strings.Join(strs, ", ") + ") "
}

func (mv *MyVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	// left=expressionAtom mathOperator right=expressionAtom
	return mv.Visit(ctx.GetLeft()).(string) + WS + ctx.MathOperator().GetText() + WS + mv.Visit(ctx.right).(string) + WS
}

func (mv *MyVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{} {
	// INTERVAL expression intervalType
	return ctx.INTERVAL().GetText() + WS + mv.Visit(ctx.Expression()).(string) + WS + ctx.IntervalType().GetText() + WS
}

func (mv *MyVisitor) VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{} {
	// left=expressionAtom jsonOperator right=expressionAtom
	return mv.Visit(ctx.GetLeft()).(string) + WS + ctx.JsonOperator().GetText() + WS + mv.Visit(ctx.right).(string) + WS
}

func (mv *MyVisitor) VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{} {
	// EXISTS '(' selectStatement ')'
	return ctx.EXISTS().GetText() + "(" + mv.Visit(ctx.SelectStatement()).(string) + ") "
}

func (mv *MyVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return mv.Visit(ctx.Constant()).(string) + WS
}

func (mv *MyVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return mv.Visit(ctx.FunctionCall()).(string) + WS
}

func (mv *MyVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{} {
	// BINARY expressionAtom
	return ctx.BINARY().GetText() + WS + mv.Visit(ctx.ExpressionAtom()).(string) + WS
}

func (mv *MyVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return mv.Visit(ctx.FullColumnName()).(string) + WS
}

func (mv *MyVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{} {
	// left=expressionAtom bitOperator right=expressionAtom
	return mv.Visit(ctx.GetLeft()).(string) + WS + ctx.BitOperator().GetText() + WS + mv.Visit(ctx.right).(string) + WS
}

func (mv *MyVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBitOperator(ctx *BitOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitJsonOperator(ctx *JsonOperatorContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCharsetNameBase(ctx *CharsetNameBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPrivilegesBase(ctx *PrivilegesBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDataTypeBase(ctx *DataTypeBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFunctionNameBase(ctx *FunctionNameBaseContext) interface{} {
	return ctx.GetText()
}

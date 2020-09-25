package selectsql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

const WS = " "

type MyVisitor struct{}

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
	return new(MyVisitor)
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
	println(ctx.GetText())
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	println(ctx.GetText())
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{} {
	println(ctx.GetText())
	return ctx.GetText()
}

func (mv *MyVisitor) VisitAssignmentField(ctx *AssignmentFieldContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionStatement(ctx *UnionStatementContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionStatement2(ctx *UnionStatement2Context) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnionstatement3(ctx *Unionstatement3Context) interface{} {
	return ctx.GetText()
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
		for _, e := range elements {
			sb.WriteString(mv.Visit(e).(string) + WS)
		}
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
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFullId(ctx *FullIdContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return ctx.GetText()
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
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNullNotnull(ctx *NullNotnullContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitStringDataType(ctx *StringDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCollectionOptions(ctx *CollectionOptionsContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitConvertedDataType(ctx *ConvertedDataTypeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUidList(ctx *UidListContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return ctx.GetText()
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
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLevelWeightList(ctx *LevelWeightListContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{} {
	return ctx.GetText()
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

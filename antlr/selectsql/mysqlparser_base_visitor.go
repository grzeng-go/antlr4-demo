// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\MySqlParser.g4 by ANTLR 4.8. DO NOT EDIT.

package selectsql // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMySqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySqlParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionSelect(ctx *UnionSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAssignmentField(ctx *AssignmentFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSources(ctx *TableSourcesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInnerJoin(ctx *InnerJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStraightJoin(ctx *StraightJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOuterJoin(ctx *OuterJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecification1(ctx *QuerySpecification1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecification2(ctx *QuerySpecification2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionStatement(ctx *UnionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionStatement2(ctx *UnionStatement2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionstatement3(ctx *Unionstatement3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectSpec(ctx *SelectSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectElements(ctx *SelectElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullId(ctx *FullIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnName(ctx *FullColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysqlVariable(ctx *MysqlVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUid(ctx *UidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEngineName(ctx *EngineNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleId(ctx *SimpleIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDottedId(ctx *DottedIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringLiteral1(ctx *StringLiteral1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringLiteral2(ctx *StringLiteral2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNullNotnull(ctx *NullNotnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringDataType(ctx *StringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollectionOptions(ctx *CollectionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType1(ctx *ConvertedDataType1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType2(ctx *ConvertedDataType2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType3(ctx *ConvertedDataType3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType4(ctx *ConvertedDataType4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType5(ctx *ConvertedDataType5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUidList(ctx *UidListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalType(ctx *IntervalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightList(ctx *LevelWeightListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction1(ctx *AggregateWindowedFunction1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction2(ctx *AggregateWindowedFunction2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction3(ctx *AggregateWindowedFunction3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction4(ctx *AggregateWindowedFunction4Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction5(ctx *AggregateWindowedFunction5Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsExpression(ctx *IsExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInPredicate(ctx *InPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBitOperator(ctx *BitOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathOperator(ctx *MathOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonOperator(ctx *JsonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharsetNameBase(ctx *CharsetNameBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivilegesBase(ctx *PrivilegesBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDataTypeBase(ctx *DataTypeBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionNameBase(ctx *FunctionNameBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\MySqlParser.g4 by ANTLR 4.8. DO NOT EDIT.

package selectsql // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MySqlParser.
type MySqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySqlParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleSelect.
	VisitSimpleSelect(ctx *SimpleSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#parenthesisSelect.
	VisitParenthesisSelect(ctx *ParenthesisSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionParenthesisSelect.
	VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) interface{}

	// Visit a parse tree produced by MySqlParser#assignmentField.
	VisitAssignmentField(ctx *AssignmentFieldContext) interface{}

	// Visit a parse tree produced by MySqlParser#lockClause.
	VisitLockClause(ctx *LockClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSources.
	VisitTableSources(ctx *TableSourcesContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourceBase.
	VisitTableSourceBase(ctx *TableSourceBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourceNested.
	VisitTableSourceNested(ctx *TableSourceNestedContext) interface{}

	// Visit a parse tree produced by MySqlParser#atomTableItem.
	VisitAtomTableItem(ctx *AtomTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryTableItem.
	VisitSubqueryTableItem(ctx *SubqueryTableItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableSourcesItem.
	VisitTableSourcesItem(ctx *TableSourcesItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) interface{}

	// Visit a parse tree produced by MySqlParser#indexHintType.
	VisitIndexHintType(ctx *IndexHintTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#straightJoin.
	VisitStraightJoin(ctx *StraightJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#outerJoin.
	VisitOuterJoin(ctx *OuterJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#naturalJoin.
	VisitNaturalJoin(ctx *NaturalJoinContext) interface{}

	// Visit a parse tree produced by MySqlParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#queryExpressionNointo.
	VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) interface{}

	// Visit a parse tree produced by MySqlParser#querySpecification1.
	VisitQuerySpecification1(ctx *QuerySpecification1Context) interface{}

	// Visit a parse tree produced by MySqlParser#querySpecification2.
	VisitQuerySpecification2(ctx *QuerySpecification2Context) interface{}

	// Visit a parse tree produced by MySqlParser#querySpecificationNointo.
	VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionParenthesis.
	VisitUnionParenthesis(ctx *UnionParenthesisContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionStatement.
	VisitUnionStatement(ctx *UnionStatementContext) interface{}

	// Visit a parse tree produced by MySqlParser#unionStatement2.
	VisitUnionStatement2(ctx *UnionStatement2Context) interface{}

	// Visit a parse tree produced by MySqlParser#unionstatement3.
	VisitUnionstatement3(ctx *Unionstatement3Context) interface{}

	// Visit a parse tree produced by MySqlParser#selectSpec.
	VisitSelectSpec(ctx *SelectSpecContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectStarElement.
	VisitSelectStarElement(ctx *SelectStarElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectColumnElement.
	VisitSelectColumnElement(ctx *SelectColumnElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectFunctionElement.
	VisitSelectFunctionElement(ctx *SelectFunctionElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectExpressionElement.
	VisitSelectExpressionElement(ctx *SelectExpressionElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoVariables.
	VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoDumpFile.
	VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectIntoTextFile.
	VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectFieldsInto.
	VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{}

	// Visit a parse tree produced by MySqlParser#selectLinesInto.
	VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{}

	// Visit a parse tree produced by MySqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) interface{}

	// Visit a parse tree produced by MySqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#limitClauseAtom.
	VisitLimitClauseAtom(ctx *LimitClauseAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullId.
	VisitFullId(ctx *FullIdContext) interface{}

	// Visit a parse tree produced by MySqlParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullColumnName.
	VisitFullColumnName(ctx *FullColumnNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#mysqlVariable.
	VisitMysqlVariable(ctx *MysqlVariableContext) interface{}

	// Visit a parse tree produced by MySqlParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#uid.
	VisitUid(ctx *UidContext) interface{}

	// Visit a parse tree produced by MySqlParser#engineName.
	VisitEngineName(ctx *EngineNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleId.
	VisitSimpleId(ctx *SimpleIdContext) interface{}

	// Visit a parse tree produced by MySqlParser#dottedId.
	VisitDottedId(ctx *DottedIdContext) interface{}

	// Visit a parse tree produced by MySqlParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by MySqlParser#stringLiteral1.
	VisitStringLiteral1(ctx *StringLiteral1Context) interface{}

	// Visit a parse tree produced by MySqlParser#stringLiteral2.
	VisitStringLiteral2(ctx *StringLiteral2Context) interface{}

	// Visit a parse tree produced by MySqlParser#nullNotnull.
	VisitNullNotnull(ctx *NullNotnullContext) interface{}

	// Visit a parse tree produced by MySqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by MySqlParser#hexadecimalLiteral.
	VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) interface{}

	// Visit a parse tree produced by MySqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by MySqlParser#stringDataType.
	VisitStringDataType(ctx *StringDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#nationalStringDataType.
	VisitNationalStringDataType(ctx *NationalStringDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#nationalVaryingStringDataType.
	VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#dimensionDataType.
	VisitDimensionDataType(ctx *DimensionDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleDataType.
	VisitSimpleDataType(ctx *SimpleDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#collectionDataType.
	VisitCollectionDataType(ctx *CollectionDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#spatialDataType.
	VisitSpatialDataType(ctx *SpatialDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#longVarcharDataType.
	VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#longVarbinaryDataType.
	VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#collectionOptions.
	VisitCollectionOptions(ctx *CollectionOptionsContext) interface{}

	// Visit a parse tree produced by MySqlParser#convertedDataType1.
	VisitConvertedDataType1(ctx *ConvertedDataType1Context) interface{}

	// Visit a parse tree produced by MySqlParser#convertedDataType2.
	VisitConvertedDataType2(ctx *ConvertedDataType2Context) interface{}

	// Visit a parse tree produced by MySqlParser#convertedDataType3.
	VisitConvertedDataType3(ctx *ConvertedDataType3Context) interface{}

	// Visit a parse tree produced by MySqlParser#convertedDataType4.
	VisitConvertedDataType4(ctx *ConvertedDataType4Context) interface{}

	// Visit a parse tree produced by MySqlParser#convertedDataType5.
	VisitConvertedDataType5(ctx *ConvertedDataType5Context) interface{}

	// Visit a parse tree produced by MySqlParser#lengthOneDimension.
	VisitLengthOneDimension(ctx *LengthOneDimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#lengthTwoDimension.
	VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#lengthTwoOptionalDimension.
	VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) interface{}

	// Visit a parse tree produced by MySqlParser#uidList.
	VisitUidList(ctx *UidListContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

	// Visit a parse tree produced by MySqlParser#intervalType.
	VisitIntervalType(ctx *IntervalTypeContext) interface{}

	// Visit a parse tree produced by MySqlParser#specificFunctionCall.
	VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateFunctionCall.
	VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#scalarFunctionCall.
	VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#udfFunctionCall.
	VisitUdfFunctionCall(ctx *UdfFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#passwordFunctionCall.
	VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#dataTypeFunctionCall.
	VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#valuesFunctionCall.
	VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseFunctionCall.
	VisitCaseFunctionCall(ctx *CaseFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#charFunctionCall.
	VisitCharFunctionCall(ctx *CharFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#positionFunctionCall.
	VisitPositionFunctionCall(ctx *PositionFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#substrFunctionCall.
	VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#trimFunctionCall.
	VisitTrimFunctionCall(ctx *TrimFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#weightFunctionCall.
	VisitWeightFunctionCall(ctx *WeightFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#extractFunctionCall.
	VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#getFormatFunctionCall.
	VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) interface{}

	// Visit a parse tree produced by MySqlParser#caseFuncAlternative.
	VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) interface{}

	// Visit a parse tree produced by MySqlParser#levelWeightList.
	VisitLevelWeightList(ctx *LevelWeightListContext) interface{}

	// Visit a parse tree produced by MySqlParser#levelWeightRange.
	VisitLevelWeightRange(ctx *LevelWeightRangeContext) interface{}

	// Visit a parse tree produced by MySqlParser#levelInWeightListElement.
	VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction1.
	VisitAggregateWindowedFunction1(ctx *AggregateWindowedFunction1Context) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction2.
	VisitAggregateWindowedFunction2(ctx *AggregateWindowedFunction2Context) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction3.
	VisitAggregateWindowedFunction3(ctx *AggregateWindowedFunction3Context) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction4.
	VisitAggregateWindowedFunction4(ctx *AggregateWindowedFunction4Context) interface{}

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction5.
	VisitAggregateWindowedFunction5(ctx *AggregateWindowedFunction5Context) interface{}

	// Visit a parse tree produced by MySqlParser#scalarFunctionName.
	VisitScalarFunctionName(ctx *ScalarFunctionNameContext) interface{}

	// Visit a parse tree produced by MySqlParser#passwordFunctionClause.
	VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}

	// Visit a parse tree produced by MySqlParser#isExpression.
	VisitIsExpression(ctx *IsExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) interface{}

	// Visit a parse tree produced by MySqlParser#soundsLikePredicate.
	VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#jsonMemberOfPredicate.
	VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryComparasionPredicate.
	VisitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#binaryComparasionPredicate.
	VisitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) interface{}

	// Visit a parse tree produced by MySqlParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#collateExpressionAtom.
	VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#subqueryExpessionAtom.
	VisitSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#mysqlVariableExpressionAtom.
	VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#nestedExpressionAtom.
	VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#nestedRowExpressionAtom.
	VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#intervalExpressionAtom.
	VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#jsonExpressionAtom.
	VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#existsExpessionAtom.
	VisitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#binaryExpressionAtom.
	VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#fullColumnNameExpressionAtom.
	VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#bitExpressionAtom.
	VisitBitExpressionAtom(ctx *BitExpressionAtomContext) interface{}

	// Visit a parse tree produced by MySqlParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#bitOperator.
	VisitBitOperator(ctx *BitOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#jsonOperator.
	VisitJsonOperator(ctx *JsonOperatorContext) interface{}

	// Visit a parse tree produced by MySqlParser#charsetNameBase.
	VisitCharsetNameBase(ctx *CharsetNameBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#transactionLevelBase.
	VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#privilegesBase.
	VisitPrivilegesBase(ctx *PrivilegesBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#intervalTypeBase.
	VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#dataTypeBase.
	VisitDataTypeBase(ctx *DataTypeBaseContext) interface{}

	// Visit a parse tree produced by MySqlParser#keywordsCanBeId.
	VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) interface{}

	// Visit a parse tree produced by MySqlParser#functionNameBase.
	VisitFunctionNameBase(ctx *FunctionNameBaseContext) interface{}
}

// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\MySqlParser.g4 by ANTLR 4.8. DO NOT EDIT.

package selectsql // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type MySqlParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSimpleSelect is called when entering the simpleSelect production.
	EnterSimpleSelect(c *SimpleSelectContext)

	// EnterParenthesisSelect is called when entering the parenthesisSelect production.
	EnterParenthesisSelect(c *ParenthesisSelectContext)

	// EnterUnionSelect is called when entering the unionSelect production.
	EnterUnionSelect(c *UnionSelectContext)

	// EnterUnionParenthesisSelect is called when entering the unionParenthesisSelect production.
	EnterUnionParenthesisSelect(c *UnionParenthesisSelectContext)

	// EnterAssignmentField is called when entering the assignmentField production.
	EnterAssignmentField(c *AssignmentFieldContext)

	// EnterLockClause is called when entering the lockClause production.
	EnterLockClause(c *LockClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByExpression is called when entering the orderByExpression production.
	EnterOrderByExpression(c *OrderByExpressionContext)

	// EnterTableSources is called when entering the tableSources production.
	EnterTableSources(c *TableSourcesContext)

	// EnterTableSourceBase is called when entering the tableSourceBase production.
	EnterTableSourceBase(c *TableSourceBaseContext)

	// EnterTableSourceNested is called when entering the tableSourceNested production.
	EnterTableSourceNested(c *TableSourceNestedContext)

	// EnterAtomTableItem is called when entering the atomTableItem production.
	EnterAtomTableItem(c *AtomTableItemContext)

	// EnterSubqueryTableItem is called when entering the subqueryTableItem production.
	EnterSubqueryTableItem(c *SubqueryTableItemContext)

	// EnterTableSourcesItem is called when entering the tableSourcesItem production.
	EnterTableSourcesItem(c *TableSourcesItemContext)

	// EnterIndexHint is called when entering the indexHint production.
	EnterIndexHint(c *IndexHintContext)

	// EnterIndexHintType is called when entering the indexHintType production.
	EnterIndexHintType(c *IndexHintTypeContext)

	// EnterInnerJoin is called when entering the innerJoin production.
	EnterInnerJoin(c *InnerJoinContext)

	// EnterStraightJoin is called when entering the straightJoin production.
	EnterStraightJoin(c *StraightJoinContext)

	// EnterOuterJoin is called when entering the outerJoin production.
	EnterOuterJoin(c *OuterJoinContext)

	// EnterNaturalJoin is called when entering the naturalJoin production.
	EnterNaturalJoin(c *NaturalJoinContext)

	// EnterQueryExpression is called when entering the queryExpression production.
	EnterQueryExpression(c *QueryExpressionContext)

	// EnterQueryExpressionNointo is called when entering the queryExpressionNointo production.
	EnterQueryExpressionNointo(c *QueryExpressionNointoContext)

	// EnterQuerySpecification1 is called when entering the querySpecification1 production.
	EnterQuerySpecification1(c *QuerySpecification1Context)

	// EnterQuerySpecification2 is called when entering the querySpecification2 production.
	EnterQuerySpecification2(c *QuerySpecification2Context)

	// EnterQuerySpecificationNointo is called when entering the querySpecificationNointo production.
	EnterQuerySpecificationNointo(c *QuerySpecificationNointoContext)

	// EnterUnionParenthesis is called when entering the unionParenthesis production.
	EnterUnionParenthesis(c *UnionParenthesisContext)

	// EnterUnionStatement is called when entering the unionStatement production.
	EnterUnionStatement(c *UnionStatementContext)

	// EnterUnionStatement2 is called when entering the unionStatement2 production.
	EnterUnionStatement2(c *UnionStatement2Context)

	// EnterUnionstatement3 is called when entering the unionstatement3 production.
	EnterUnionstatement3(c *Unionstatement3Context)

	// EnterSelectSpec is called when entering the selectSpec production.
	EnterSelectSpec(c *SelectSpecContext)

	// EnterSelectElements is called when entering the selectElements production.
	EnterSelectElements(c *SelectElementsContext)

	// EnterSelectStarElement is called when entering the selectStarElement production.
	EnterSelectStarElement(c *SelectStarElementContext)

	// EnterSelectColumnElement is called when entering the selectColumnElement production.
	EnterSelectColumnElement(c *SelectColumnElementContext)

	// EnterSelectFunctionElement is called when entering the selectFunctionElement production.
	EnterSelectFunctionElement(c *SelectFunctionElementContext)

	// EnterSelectExpressionElement is called when entering the selectExpressionElement production.
	EnterSelectExpressionElement(c *SelectExpressionElementContext)

	// EnterSelectIntoVariables is called when entering the selectIntoVariables production.
	EnterSelectIntoVariables(c *SelectIntoVariablesContext)

	// EnterSelectIntoDumpFile is called when entering the selectIntoDumpFile production.
	EnterSelectIntoDumpFile(c *SelectIntoDumpFileContext)

	// EnterSelectIntoTextFile is called when entering the selectIntoTextFile production.
	EnterSelectIntoTextFile(c *SelectIntoTextFileContext)

	// EnterSelectFieldsInto is called when entering the selectFieldsInto production.
	EnterSelectFieldsInto(c *SelectFieldsIntoContext)

	// EnterSelectLinesInto is called when entering the selectLinesInto production.
	EnterSelectLinesInto(c *SelectLinesIntoContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterGroupByItem is called when entering the groupByItem production.
	EnterGroupByItem(c *GroupByItemContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterLimitClauseAtom is called when entering the limitClauseAtom production.
	EnterLimitClauseAtom(c *LimitClauseAtomContext)

	// EnterFullId is called when entering the fullId production.
	EnterFullId(c *FullIdContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterFullColumnName is called when entering the fullColumnName production.
	EnterFullColumnName(c *FullColumnNameContext)

	// EnterMysqlVariable is called when entering the mysqlVariable production.
	EnterMysqlVariable(c *MysqlVariableContext)

	// EnterCharsetName is called when entering the charsetName production.
	EnterCharsetName(c *CharsetNameContext)

	// EnterUid is called when entering the uid production.
	EnterUid(c *UidContext)

	// EnterEngineName is called when entering the engineName production.
	EnterEngineName(c *EngineNameContext)

	// EnterSimpleId is called when entering the simpleId production.
	EnterSimpleId(c *SimpleIdContext)

	// EnterDottedId is called when entering the dottedId production.
	EnterDottedId(c *DottedIdContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterStringLiteral1 is called when entering the stringLiteral1 production.
	EnterStringLiteral1(c *StringLiteral1Context)

	// EnterStringLiteral2 is called when entering the stringLiteral2 production.
	EnterStringLiteral2(c *StringLiteral2Context)

	// EnterNullNotnull is called when entering the nullNotnull production.
	EnterNullNotnull(c *NullNotnullContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterHexadecimalLiteral is called when entering the hexadecimalLiteral production.
	EnterHexadecimalLiteral(c *HexadecimalLiteralContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterStringDataType is called when entering the stringDataType production.
	EnterStringDataType(c *StringDataTypeContext)

	// EnterNationalStringDataType is called when entering the nationalStringDataType production.
	EnterNationalStringDataType(c *NationalStringDataTypeContext)

	// EnterNationalVaryingStringDataType is called when entering the nationalVaryingStringDataType production.
	EnterNationalVaryingStringDataType(c *NationalVaryingStringDataTypeContext)

	// EnterDimensionDataType is called when entering the dimensionDataType production.
	EnterDimensionDataType(c *DimensionDataTypeContext)

	// EnterSimpleDataType is called when entering the simpleDataType production.
	EnterSimpleDataType(c *SimpleDataTypeContext)

	// EnterCollectionDataType is called when entering the collectionDataType production.
	EnterCollectionDataType(c *CollectionDataTypeContext)

	// EnterSpatialDataType is called when entering the spatialDataType production.
	EnterSpatialDataType(c *SpatialDataTypeContext)

	// EnterLongVarcharDataType is called when entering the longVarcharDataType production.
	EnterLongVarcharDataType(c *LongVarcharDataTypeContext)

	// EnterLongVarbinaryDataType is called when entering the longVarbinaryDataType production.
	EnterLongVarbinaryDataType(c *LongVarbinaryDataTypeContext)

	// EnterCollectionOptions is called when entering the collectionOptions production.
	EnterCollectionOptions(c *CollectionOptionsContext)

	// EnterConvertedDataType1 is called when entering the convertedDataType1 production.
	EnterConvertedDataType1(c *ConvertedDataType1Context)

	// EnterConvertedDataType2 is called when entering the convertedDataType2 production.
	EnterConvertedDataType2(c *ConvertedDataType2Context)

	// EnterConvertedDataType3 is called when entering the convertedDataType3 production.
	EnterConvertedDataType3(c *ConvertedDataType3Context)

	// EnterConvertedDataType4 is called when entering the convertedDataType4 production.
	EnterConvertedDataType4(c *ConvertedDataType4Context)

	// EnterConvertedDataType5 is called when entering the convertedDataType5 production.
	EnterConvertedDataType5(c *ConvertedDataType5Context)

	// EnterLengthOneDimension is called when entering the lengthOneDimension production.
	EnterLengthOneDimension(c *LengthOneDimensionContext)

	// EnterLengthTwoDimension is called when entering the lengthTwoDimension production.
	EnterLengthTwoDimension(c *LengthTwoDimensionContext)

	// EnterLengthTwoOptionalDimension is called when entering the lengthTwoOptionalDimension production.
	EnterLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// EnterUidList is called when entering the uidList production.
	EnterUidList(c *UidListContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

	// EnterIntervalType is called when entering the intervalType production.
	EnterIntervalType(c *IntervalTypeContext)

	// EnterSpecificFunctionCall is called when entering the specificFunctionCall production.
	EnterSpecificFunctionCall(c *SpecificFunctionCallContext)

	// EnterAggregateFunctionCall is called when entering the aggregateFunctionCall production.
	EnterAggregateFunctionCall(c *AggregateFunctionCallContext)

	// EnterScalarFunctionCall is called when entering the scalarFunctionCall production.
	EnterScalarFunctionCall(c *ScalarFunctionCallContext)

	// EnterUdfFunctionCall is called when entering the udfFunctionCall production.
	EnterUdfFunctionCall(c *UdfFunctionCallContext)

	// EnterPasswordFunctionCall is called when entering the passwordFunctionCall production.
	EnterPasswordFunctionCall(c *PasswordFunctionCallContext)

	// EnterSimpleFunctionCall is called when entering the simpleFunctionCall production.
	EnterSimpleFunctionCall(c *SimpleFunctionCallContext)

	// EnterDataTypeFunctionCall is called when entering the dataTypeFunctionCall production.
	EnterDataTypeFunctionCall(c *DataTypeFunctionCallContext)

	// EnterValuesFunctionCall is called when entering the valuesFunctionCall production.
	EnterValuesFunctionCall(c *ValuesFunctionCallContext)

	// EnterCaseFunctionCall is called when entering the caseFunctionCall production.
	EnterCaseFunctionCall(c *CaseFunctionCallContext)

	// EnterCharFunctionCall is called when entering the charFunctionCall production.
	EnterCharFunctionCall(c *CharFunctionCallContext)

	// EnterPositionFunctionCall is called when entering the positionFunctionCall production.
	EnterPositionFunctionCall(c *PositionFunctionCallContext)

	// EnterSubstrFunctionCall is called when entering the substrFunctionCall production.
	EnterSubstrFunctionCall(c *SubstrFunctionCallContext)

	// EnterTrimFunctionCall is called when entering the trimFunctionCall production.
	EnterTrimFunctionCall(c *TrimFunctionCallContext)

	// EnterWeightFunctionCall is called when entering the weightFunctionCall production.
	EnterWeightFunctionCall(c *WeightFunctionCallContext)

	// EnterExtractFunctionCall is called when entering the extractFunctionCall production.
	EnterExtractFunctionCall(c *ExtractFunctionCallContext)

	// EnterGetFormatFunctionCall is called when entering the getFormatFunctionCall production.
	EnterGetFormatFunctionCall(c *GetFormatFunctionCallContext)

	// EnterCaseFuncAlternative is called when entering the caseFuncAlternative production.
	EnterCaseFuncAlternative(c *CaseFuncAlternativeContext)

	// EnterLevelWeightList is called when entering the levelWeightList production.
	EnterLevelWeightList(c *LevelWeightListContext)

	// EnterLevelWeightRange is called when entering the levelWeightRange production.
	EnterLevelWeightRange(c *LevelWeightRangeContext)

	// EnterLevelInWeightListElement is called when entering the levelInWeightListElement production.
	EnterLevelInWeightListElement(c *LevelInWeightListElementContext)

	// EnterAggregateWindowedFunction1 is called when entering the aggregateWindowedFunction1 production.
	EnterAggregateWindowedFunction1(c *AggregateWindowedFunction1Context)

	// EnterAggregateWindowedFunction2 is called when entering the aggregateWindowedFunction2 production.
	EnterAggregateWindowedFunction2(c *AggregateWindowedFunction2Context)

	// EnterAggregateWindowedFunction3 is called when entering the aggregateWindowedFunction3 production.
	EnterAggregateWindowedFunction3(c *AggregateWindowedFunction3Context)

	// EnterAggregateWindowedFunction4 is called when entering the aggregateWindowedFunction4 production.
	EnterAggregateWindowedFunction4(c *AggregateWindowedFunction4Context)

	// EnterAggregateWindowedFunction5 is called when entering the aggregateWindowedFunction5 production.
	EnterAggregateWindowedFunction5(c *AggregateWindowedFunction5Context)

	// EnterScalarFunctionName is called when entering the scalarFunctionName production.
	EnterScalarFunctionName(c *ScalarFunctionNameContext)

	// EnterPasswordFunctionClause is called when entering the passwordFunctionClause production.
	EnterPasswordFunctionClause(c *PasswordFunctionClauseContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterFunctionArg is called when entering the functionArg production.
	EnterFunctionArg(c *FunctionArgContext)

	// EnterIsExpression is called when entering the isExpression production.
	EnterIsExpression(c *IsExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterSoundsLikePredicate is called when entering the soundsLikePredicate production.
	EnterSoundsLikePredicate(c *SoundsLikePredicateContext)

	// EnterExpressionAtomPredicate is called when entering the expressionAtomPredicate production.
	EnterExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// EnterJsonMemberOfPredicate is called when entering the jsonMemberOfPredicate production.
	EnterJsonMemberOfPredicate(c *JsonMemberOfPredicateContext)

	// EnterInPredicate is called when entering the inPredicate production.
	EnterInPredicate(c *InPredicateContext)

	// EnterSubqueryComparasionPredicate is called when entering the subqueryComparasionPredicate production.
	EnterSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// EnterBetweenPredicate is called when entering the betweenPredicate production.
	EnterBetweenPredicate(c *BetweenPredicateContext)

	// EnterBinaryComparasionPredicate is called when entering the binaryComparasionPredicate production.
	EnterBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// EnterIsNullPredicate is called when entering the isNullPredicate production.
	EnterIsNullPredicate(c *IsNullPredicateContext)

	// EnterLikePredicate is called when entering the likePredicate production.
	EnterLikePredicate(c *LikePredicateContext)

	// EnterRegexpPredicate is called when entering the regexpPredicate production.
	EnterRegexpPredicate(c *RegexpPredicateContext)

	// EnterUnaryExpressionAtom is called when entering the unaryExpressionAtom production.
	EnterUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// EnterCollateExpressionAtom is called when entering the collateExpressionAtom production.
	EnterCollateExpressionAtom(c *CollateExpressionAtomContext)

	// EnterSubqueryExpessionAtom is called when entering the subqueryExpessionAtom production.
	EnterSubqueryExpessionAtom(c *SubqueryExpessionAtomContext)

	// EnterMysqlVariableExpressionAtom is called when entering the mysqlVariableExpressionAtom production.
	EnterMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// EnterNestedExpressionAtom is called when entering the nestedExpressionAtom production.
	EnterNestedExpressionAtom(c *NestedExpressionAtomContext)

	// EnterNestedRowExpressionAtom is called when entering the nestedRowExpressionAtom production.
	EnterNestedRowExpressionAtom(c *NestedRowExpressionAtomContext)

	// EnterMathExpressionAtom is called when entering the mathExpressionAtom production.
	EnterMathExpressionAtom(c *MathExpressionAtomContext)

	// EnterIntervalExpressionAtom is called when entering the intervalExpressionAtom production.
	EnterIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// EnterJsonExpressionAtom is called when entering the jsonExpressionAtom production.
	EnterJsonExpressionAtom(c *JsonExpressionAtomContext)

	// EnterExistsExpessionAtom is called when entering the existsExpessionAtom production.
	EnterExistsExpessionAtom(c *ExistsExpessionAtomContext)

	// EnterConstantExpressionAtom is called when entering the constantExpressionAtom production.
	EnterConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// EnterFunctionCallExpressionAtom is called when entering the functionCallExpressionAtom production.
	EnterFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// EnterBinaryExpressionAtom is called when entering the binaryExpressionAtom production.
	EnterBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// EnterFullColumnNameExpressionAtom is called when entering the fullColumnNameExpressionAtom production.
	EnterFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// EnterBitExpressionAtom is called when entering the bitExpressionAtom production.
	EnterBitExpressionAtom(c *BitExpressionAtomContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterBitOperator is called when entering the bitOperator production.
	EnterBitOperator(c *BitOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterJsonOperator is called when entering the jsonOperator production.
	EnterJsonOperator(c *JsonOperatorContext)

	// EnterCharsetNameBase is called when entering the charsetNameBase production.
	EnterCharsetNameBase(c *CharsetNameBaseContext)

	// EnterTransactionLevelBase is called when entering the transactionLevelBase production.
	EnterTransactionLevelBase(c *TransactionLevelBaseContext)

	// EnterPrivilegesBase is called when entering the privilegesBase production.
	EnterPrivilegesBase(c *PrivilegesBaseContext)

	// EnterIntervalTypeBase is called when entering the intervalTypeBase production.
	EnterIntervalTypeBase(c *IntervalTypeBaseContext)

	// EnterDataTypeBase is called when entering the dataTypeBase production.
	EnterDataTypeBase(c *DataTypeBaseContext)

	// EnterKeywordsCanBeId is called when entering the keywordsCanBeId production.
	EnterKeywordsCanBeId(c *KeywordsCanBeIdContext)

	// EnterFunctionNameBase is called when entering the functionNameBase production.
	EnterFunctionNameBase(c *FunctionNameBaseContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSimpleSelect is called when exiting the simpleSelect production.
	ExitSimpleSelect(c *SimpleSelectContext)

	// ExitParenthesisSelect is called when exiting the parenthesisSelect production.
	ExitParenthesisSelect(c *ParenthesisSelectContext)

	// ExitUnionSelect is called when exiting the unionSelect production.
	ExitUnionSelect(c *UnionSelectContext)

	// ExitUnionParenthesisSelect is called when exiting the unionParenthesisSelect production.
	ExitUnionParenthesisSelect(c *UnionParenthesisSelectContext)

	// ExitAssignmentField is called when exiting the assignmentField production.
	ExitAssignmentField(c *AssignmentFieldContext)

	// ExitLockClause is called when exiting the lockClause production.
	ExitLockClause(c *LockClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByExpression is called when exiting the orderByExpression production.
	ExitOrderByExpression(c *OrderByExpressionContext)

	// ExitTableSources is called when exiting the tableSources production.
	ExitTableSources(c *TableSourcesContext)

	// ExitTableSourceBase is called when exiting the tableSourceBase production.
	ExitTableSourceBase(c *TableSourceBaseContext)

	// ExitTableSourceNested is called when exiting the tableSourceNested production.
	ExitTableSourceNested(c *TableSourceNestedContext)

	// ExitAtomTableItem is called when exiting the atomTableItem production.
	ExitAtomTableItem(c *AtomTableItemContext)

	// ExitSubqueryTableItem is called when exiting the subqueryTableItem production.
	ExitSubqueryTableItem(c *SubqueryTableItemContext)

	// ExitTableSourcesItem is called when exiting the tableSourcesItem production.
	ExitTableSourcesItem(c *TableSourcesItemContext)

	// ExitIndexHint is called when exiting the indexHint production.
	ExitIndexHint(c *IndexHintContext)

	// ExitIndexHintType is called when exiting the indexHintType production.
	ExitIndexHintType(c *IndexHintTypeContext)

	// ExitInnerJoin is called when exiting the innerJoin production.
	ExitInnerJoin(c *InnerJoinContext)

	// ExitStraightJoin is called when exiting the straightJoin production.
	ExitStraightJoin(c *StraightJoinContext)

	// ExitOuterJoin is called when exiting the outerJoin production.
	ExitOuterJoin(c *OuterJoinContext)

	// ExitNaturalJoin is called when exiting the naturalJoin production.
	ExitNaturalJoin(c *NaturalJoinContext)

	// ExitQueryExpression is called when exiting the queryExpression production.
	ExitQueryExpression(c *QueryExpressionContext)

	// ExitQueryExpressionNointo is called when exiting the queryExpressionNointo production.
	ExitQueryExpressionNointo(c *QueryExpressionNointoContext)

	// ExitQuerySpecification1 is called when exiting the querySpecification1 production.
	ExitQuerySpecification1(c *QuerySpecification1Context)

	// ExitQuerySpecification2 is called when exiting the querySpecification2 production.
	ExitQuerySpecification2(c *QuerySpecification2Context)

	// ExitQuerySpecificationNointo is called when exiting the querySpecificationNointo production.
	ExitQuerySpecificationNointo(c *QuerySpecificationNointoContext)

	// ExitUnionParenthesis is called when exiting the unionParenthesis production.
	ExitUnionParenthesis(c *UnionParenthesisContext)

	// ExitUnionStatement is called when exiting the unionStatement production.
	ExitUnionStatement(c *UnionStatementContext)

	// ExitUnionStatement2 is called when exiting the unionStatement2 production.
	ExitUnionStatement2(c *UnionStatement2Context)

	// ExitUnionstatement3 is called when exiting the unionstatement3 production.
	ExitUnionstatement3(c *Unionstatement3Context)

	// ExitSelectSpec is called when exiting the selectSpec production.
	ExitSelectSpec(c *SelectSpecContext)

	// ExitSelectElements is called when exiting the selectElements production.
	ExitSelectElements(c *SelectElementsContext)

	// ExitSelectStarElement is called when exiting the selectStarElement production.
	ExitSelectStarElement(c *SelectStarElementContext)

	// ExitSelectColumnElement is called when exiting the selectColumnElement production.
	ExitSelectColumnElement(c *SelectColumnElementContext)

	// ExitSelectFunctionElement is called when exiting the selectFunctionElement production.
	ExitSelectFunctionElement(c *SelectFunctionElementContext)

	// ExitSelectExpressionElement is called when exiting the selectExpressionElement production.
	ExitSelectExpressionElement(c *SelectExpressionElementContext)

	// ExitSelectIntoVariables is called when exiting the selectIntoVariables production.
	ExitSelectIntoVariables(c *SelectIntoVariablesContext)

	// ExitSelectIntoDumpFile is called when exiting the selectIntoDumpFile production.
	ExitSelectIntoDumpFile(c *SelectIntoDumpFileContext)

	// ExitSelectIntoTextFile is called when exiting the selectIntoTextFile production.
	ExitSelectIntoTextFile(c *SelectIntoTextFileContext)

	// ExitSelectFieldsInto is called when exiting the selectFieldsInto production.
	ExitSelectFieldsInto(c *SelectFieldsIntoContext)

	// ExitSelectLinesInto is called when exiting the selectLinesInto production.
	ExitSelectLinesInto(c *SelectLinesIntoContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitGroupByItem is called when exiting the groupByItem production.
	ExitGroupByItem(c *GroupByItemContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitLimitClauseAtom is called when exiting the limitClauseAtom production.
	ExitLimitClauseAtom(c *LimitClauseAtomContext)

	// ExitFullId is called when exiting the fullId production.
	ExitFullId(c *FullIdContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitFullColumnName is called when exiting the fullColumnName production.
	ExitFullColumnName(c *FullColumnNameContext)

	// ExitMysqlVariable is called when exiting the mysqlVariable production.
	ExitMysqlVariable(c *MysqlVariableContext)

	// ExitCharsetName is called when exiting the charsetName production.
	ExitCharsetName(c *CharsetNameContext)

	// ExitUid is called when exiting the uid production.
	ExitUid(c *UidContext)

	// ExitEngineName is called when exiting the engineName production.
	ExitEngineName(c *EngineNameContext)

	// ExitSimpleId is called when exiting the simpleId production.
	ExitSimpleId(c *SimpleIdContext)

	// ExitDottedId is called when exiting the dottedId production.
	ExitDottedId(c *DottedIdContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitStringLiteral1 is called when exiting the stringLiteral1 production.
	ExitStringLiteral1(c *StringLiteral1Context)

	// ExitStringLiteral2 is called when exiting the stringLiteral2 production.
	ExitStringLiteral2(c *StringLiteral2Context)

	// ExitNullNotnull is called when exiting the nullNotnull production.
	ExitNullNotnull(c *NullNotnullContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitHexadecimalLiteral is called when exiting the hexadecimalLiteral production.
	ExitHexadecimalLiteral(c *HexadecimalLiteralContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitStringDataType is called when exiting the stringDataType production.
	ExitStringDataType(c *StringDataTypeContext)

	// ExitNationalStringDataType is called when exiting the nationalStringDataType production.
	ExitNationalStringDataType(c *NationalStringDataTypeContext)

	// ExitNationalVaryingStringDataType is called when exiting the nationalVaryingStringDataType production.
	ExitNationalVaryingStringDataType(c *NationalVaryingStringDataTypeContext)

	// ExitDimensionDataType is called when exiting the dimensionDataType production.
	ExitDimensionDataType(c *DimensionDataTypeContext)

	// ExitSimpleDataType is called when exiting the simpleDataType production.
	ExitSimpleDataType(c *SimpleDataTypeContext)

	// ExitCollectionDataType is called when exiting the collectionDataType production.
	ExitCollectionDataType(c *CollectionDataTypeContext)

	// ExitSpatialDataType is called when exiting the spatialDataType production.
	ExitSpatialDataType(c *SpatialDataTypeContext)

	// ExitLongVarcharDataType is called when exiting the longVarcharDataType production.
	ExitLongVarcharDataType(c *LongVarcharDataTypeContext)

	// ExitLongVarbinaryDataType is called when exiting the longVarbinaryDataType production.
	ExitLongVarbinaryDataType(c *LongVarbinaryDataTypeContext)

	// ExitCollectionOptions is called when exiting the collectionOptions production.
	ExitCollectionOptions(c *CollectionOptionsContext)

	// ExitConvertedDataType1 is called when exiting the convertedDataType1 production.
	ExitConvertedDataType1(c *ConvertedDataType1Context)

	// ExitConvertedDataType2 is called when exiting the convertedDataType2 production.
	ExitConvertedDataType2(c *ConvertedDataType2Context)

	// ExitConvertedDataType3 is called when exiting the convertedDataType3 production.
	ExitConvertedDataType3(c *ConvertedDataType3Context)

	// ExitConvertedDataType4 is called when exiting the convertedDataType4 production.
	ExitConvertedDataType4(c *ConvertedDataType4Context)

	// ExitConvertedDataType5 is called when exiting the convertedDataType5 production.
	ExitConvertedDataType5(c *ConvertedDataType5Context)

	// ExitLengthOneDimension is called when exiting the lengthOneDimension production.
	ExitLengthOneDimension(c *LengthOneDimensionContext)

	// ExitLengthTwoDimension is called when exiting the lengthTwoDimension production.
	ExitLengthTwoDimension(c *LengthTwoDimensionContext)

	// ExitLengthTwoOptionalDimension is called when exiting the lengthTwoOptionalDimension production.
	ExitLengthTwoOptionalDimension(c *LengthTwoOptionalDimensionContext)

	// ExitUidList is called when exiting the uidList production.
	ExitUidList(c *UidListContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

	// ExitIntervalType is called when exiting the intervalType production.
	ExitIntervalType(c *IntervalTypeContext)

	// ExitSpecificFunctionCall is called when exiting the specificFunctionCall production.
	ExitSpecificFunctionCall(c *SpecificFunctionCallContext)

	// ExitAggregateFunctionCall is called when exiting the aggregateFunctionCall production.
	ExitAggregateFunctionCall(c *AggregateFunctionCallContext)

	// ExitScalarFunctionCall is called when exiting the scalarFunctionCall production.
	ExitScalarFunctionCall(c *ScalarFunctionCallContext)

	// ExitUdfFunctionCall is called when exiting the udfFunctionCall production.
	ExitUdfFunctionCall(c *UdfFunctionCallContext)

	// ExitPasswordFunctionCall is called when exiting the passwordFunctionCall production.
	ExitPasswordFunctionCall(c *PasswordFunctionCallContext)

	// ExitSimpleFunctionCall is called when exiting the simpleFunctionCall production.
	ExitSimpleFunctionCall(c *SimpleFunctionCallContext)

	// ExitDataTypeFunctionCall is called when exiting the dataTypeFunctionCall production.
	ExitDataTypeFunctionCall(c *DataTypeFunctionCallContext)

	// ExitValuesFunctionCall is called when exiting the valuesFunctionCall production.
	ExitValuesFunctionCall(c *ValuesFunctionCallContext)

	// ExitCaseFunctionCall is called when exiting the caseFunctionCall production.
	ExitCaseFunctionCall(c *CaseFunctionCallContext)

	// ExitCharFunctionCall is called when exiting the charFunctionCall production.
	ExitCharFunctionCall(c *CharFunctionCallContext)

	// ExitPositionFunctionCall is called when exiting the positionFunctionCall production.
	ExitPositionFunctionCall(c *PositionFunctionCallContext)

	// ExitSubstrFunctionCall is called when exiting the substrFunctionCall production.
	ExitSubstrFunctionCall(c *SubstrFunctionCallContext)

	// ExitTrimFunctionCall is called when exiting the trimFunctionCall production.
	ExitTrimFunctionCall(c *TrimFunctionCallContext)

	// ExitWeightFunctionCall is called when exiting the weightFunctionCall production.
	ExitWeightFunctionCall(c *WeightFunctionCallContext)

	// ExitExtractFunctionCall is called when exiting the extractFunctionCall production.
	ExitExtractFunctionCall(c *ExtractFunctionCallContext)

	// ExitGetFormatFunctionCall is called when exiting the getFormatFunctionCall production.
	ExitGetFormatFunctionCall(c *GetFormatFunctionCallContext)

	// ExitCaseFuncAlternative is called when exiting the caseFuncAlternative production.
	ExitCaseFuncAlternative(c *CaseFuncAlternativeContext)

	// ExitLevelWeightList is called when exiting the levelWeightList production.
	ExitLevelWeightList(c *LevelWeightListContext)

	// ExitLevelWeightRange is called when exiting the levelWeightRange production.
	ExitLevelWeightRange(c *LevelWeightRangeContext)

	// ExitLevelInWeightListElement is called when exiting the levelInWeightListElement production.
	ExitLevelInWeightListElement(c *LevelInWeightListElementContext)

	// ExitAggregateWindowedFunction1 is called when exiting the aggregateWindowedFunction1 production.
	ExitAggregateWindowedFunction1(c *AggregateWindowedFunction1Context)

	// ExitAggregateWindowedFunction2 is called when exiting the aggregateWindowedFunction2 production.
	ExitAggregateWindowedFunction2(c *AggregateWindowedFunction2Context)

	// ExitAggregateWindowedFunction3 is called when exiting the aggregateWindowedFunction3 production.
	ExitAggregateWindowedFunction3(c *AggregateWindowedFunction3Context)

	// ExitAggregateWindowedFunction4 is called when exiting the aggregateWindowedFunction4 production.
	ExitAggregateWindowedFunction4(c *AggregateWindowedFunction4Context)

	// ExitAggregateWindowedFunction5 is called when exiting the aggregateWindowedFunction5 production.
	ExitAggregateWindowedFunction5(c *AggregateWindowedFunction5Context)

	// ExitScalarFunctionName is called when exiting the scalarFunctionName production.
	ExitScalarFunctionName(c *ScalarFunctionNameContext)

	// ExitPasswordFunctionClause is called when exiting the passwordFunctionClause production.
	ExitPasswordFunctionClause(c *PasswordFunctionClauseContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitFunctionArg is called when exiting the functionArg production.
	ExitFunctionArg(c *FunctionArgContext)

	// ExitIsExpression is called when exiting the isExpression production.
	ExitIsExpression(c *IsExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitSoundsLikePredicate is called when exiting the soundsLikePredicate production.
	ExitSoundsLikePredicate(c *SoundsLikePredicateContext)

	// ExitExpressionAtomPredicate is called when exiting the expressionAtomPredicate production.
	ExitExpressionAtomPredicate(c *ExpressionAtomPredicateContext)

	// ExitJsonMemberOfPredicate is called when exiting the jsonMemberOfPredicate production.
	ExitJsonMemberOfPredicate(c *JsonMemberOfPredicateContext)

	// ExitInPredicate is called when exiting the inPredicate production.
	ExitInPredicate(c *InPredicateContext)

	// ExitSubqueryComparasionPredicate is called when exiting the subqueryComparasionPredicate production.
	ExitSubqueryComparasionPredicate(c *SubqueryComparasionPredicateContext)

	// ExitBetweenPredicate is called when exiting the betweenPredicate production.
	ExitBetweenPredicate(c *BetweenPredicateContext)

	// ExitBinaryComparasionPredicate is called when exiting the binaryComparasionPredicate production.
	ExitBinaryComparasionPredicate(c *BinaryComparasionPredicateContext)

	// ExitIsNullPredicate is called when exiting the isNullPredicate production.
	ExitIsNullPredicate(c *IsNullPredicateContext)

	// ExitLikePredicate is called when exiting the likePredicate production.
	ExitLikePredicate(c *LikePredicateContext)

	// ExitRegexpPredicate is called when exiting the regexpPredicate production.
	ExitRegexpPredicate(c *RegexpPredicateContext)

	// ExitUnaryExpressionAtom is called when exiting the unaryExpressionAtom production.
	ExitUnaryExpressionAtom(c *UnaryExpressionAtomContext)

	// ExitCollateExpressionAtom is called when exiting the collateExpressionAtom production.
	ExitCollateExpressionAtom(c *CollateExpressionAtomContext)

	// ExitSubqueryExpessionAtom is called when exiting the subqueryExpessionAtom production.
	ExitSubqueryExpessionAtom(c *SubqueryExpessionAtomContext)

	// ExitMysqlVariableExpressionAtom is called when exiting the mysqlVariableExpressionAtom production.
	ExitMysqlVariableExpressionAtom(c *MysqlVariableExpressionAtomContext)

	// ExitNestedExpressionAtom is called when exiting the nestedExpressionAtom production.
	ExitNestedExpressionAtom(c *NestedExpressionAtomContext)

	// ExitNestedRowExpressionAtom is called when exiting the nestedRowExpressionAtom production.
	ExitNestedRowExpressionAtom(c *NestedRowExpressionAtomContext)

	// ExitMathExpressionAtom is called when exiting the mathExpressionAtom production.
	ExitMathExpressionAtom(c *MathExpressionAtomContext)

	// ExitIntervalExpressionAtom is called when exiting the intervalExpressionAtom production.
	ExitIntervalExpressionAtom(c *IntervalExpressionAtomContext)

	// ExitJsonExpressionAtom is called when exiting the jsonExpressionAtom production.
	ExitJsonExpressionAtom(c *JsonExpressionAtomContext)

	// ExitExistsExpessionAtom is called when exiting the existsExpessionAtom production.
	ExitExistsExpessionAtom(c *ExistsExpessionAtomContext)

	// ExitConstantExpressionAtom is called when exiting the constantExpressionAtom production.
	ExitConstantExpressionAtom(c *ConstantExpressionAtomContext)

	// ExitFunctionCallExpressionAtom is called when exiting the functionCallExpressionAtom production.
	ExitFunctionCallExpressionAtom(c *FunctionCallExpressionAtomContext)

	// ExitBinaryExpressionAtom is called when exiting the binaryExpressionAtom production.
	ExitBinaryExpressionAtom(c *BinaryExpressionAtomContext)

	// ExitFullColumnNameExpressionAtom is called when exiting the fullColumnNameExpressionAtom production.
	ExitFullColumnNameExpressionAtom(c *FullColumnNameExpressionAtomContext)

	// ExitBitExpressionAtom is called when exiting the bitExpressionAtom production.
	ExitBitExpressionAtom(c *BitExpressionAtomContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitBitOperator is called when exiting the bitOperator production.
	ExitBitOperator(c *BitOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitJsonOperator is called when exiting the jsonOperator production.
	ExitJsonOperator(c *JsonOperatorContext)

	// ExitCharsetNameBase is called when exiting the charsetNameBase production.
	ExitCharsetNameBase(c *CharsetNameBaseContext)

	// ExitTransactionLevelBase is called when exiting the transactionLevelBase production.
	ExitTransactionLevelBase(c *TransactionLevelBaseContext)

	// ExitPrivilegesBase is called when exiting the privilegesBase production.
	ExitPrivilegesBase(c *PrivilegesBaseContext)

	// ExitIntervalTypeBase is called when exiting the intervalTypeBase production.
	ExitIntervalTypeBase(c *IntervalTypeBaseContext)

	// ExitDataTypeBase is called when exiting the dataTypeBase production.
	ExitDataTypeBase(c *DataTypeBaseContext)

	// ExitKeywordsCanBeId is called when exiting the keywordsCanBeId production.
	ExitKeywordsCanBeId(c *KeywordsCanBeIdContext)

	// ExitFunctionNameBase is called when exiting the functionNameBase production.
	ExitFunctionNameBase(c *FunctionNameBaseContext)
}

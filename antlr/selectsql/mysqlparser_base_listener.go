// Code generated from D:/1_workspace/go/grzeng/antlr4-demo/antlr\MySqlParser.g4 by ANTLR 4.8. DO NOT EDIT.

package selectsql // MySqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMySqlParserListener is a complete listener for a parse tree produced by MySqlParser.
type BaseMySqlParserListener struct{}

var _ MySqlParserListener = &BaseMySqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMySqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMySqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterSimpleSelect is called when production simpleSelect is entered.
func (s *BaseMySqlParserListener) EnterSimpleSelect(ctx *SimpleSelectContext) {}

// ExitSimpleSelect is called when production simpleSelect is exited.
func (s *BaseMySqlParserListener) ExitSimpleSelect(ctx *SimpleSelectContext) {}

// EnterParenthesisSelect is called when production parenthesisSelect is entered.
func (s *BaseMySqlParserListener) EnterParenthesisSelect(ctx *ParenthesisSelectContext) {}

// ExitParenthesisSelect is called when production parenthesisSelect is exited.
func (s *BaseMySqlParserListener) ExitParenthesisSelect(ctx *ParenthesisSelectContext) {}

// EnterUnionSelect is called when production unionSelect is entered.
func (s *BaseMySqlParserListener) EnterUnionSelect(ctx *UnionSelectContext) {}

// ExitUnionSelect is called when production unionSelect is exited.
func (s *BaseMySqlParserListener) ExitUnionSelect(ctx *UnionSelectContext) {}

// EnterUnionParenthesisSelect is called when production unionParenthesisSelect is entered.
func (s *BaseMySqlParserListener) EnterUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) {}

// ExitUnionParenthesisSelect is called when production unionParenthesisSelect is exited.
func (s *BaseMySqlParserListener) ExitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) {}

// EnterAssignmentField is called when production assignmentField is entered.
func (s *BaseMySqlParserListener) EnterAssignmentField(ctx *AssignmentFieldContext) {}

// ExitAssignmentField is called when production assignmentField is exited.
func (s *BaseMySqlParserListener) ExitAssignmentField(ctx *AssignmentFieldContext) {}

// EnterLockClause is called when production lockClause is entered.
func (s *BaseMySqlParserListener) EnterLockClause(ctx *LockClauseContext) {}

// ExitLockClause is called when production lockClause is exited.
func (s *BaseMySqlParserListener) ExitLockClause(ctx *LockClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseMySqlParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseMySqlParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByExpression is called when production orderByExpression is entered.
func (s *BaseMySqlParserListener) EnterOrderByExpression(ctx *OrderByExpressionContext) {}

// ExitOrderByExpression is called when production orderByExpression is exited.
func (s *BaseMySqlParserListener) ExitOrderByExpression(ctx *OrderByExpressionContext) {}

// EnterTableSources is called when production tableSources is entered.
func (s *BaseMySqlParserListener) EnterTableSources(ctx *TableSourcesContext) {}

// ExitTableSources is called when production tableSources is exited.
func (s *BaseMySqlParserListener) ExitTableSources(ctx *TableSourcesContext) {}

// EnterTableSourceBase is called when production tableSourceBase is entered.
func (s *BaseMySqlParserListener) EnterTableSourceBase(ctx *TableSourceBaseContext) {}

// ExitTableSourceBase is called when production tableSourceBase is exited.
func (s *BaseMySqlParserListener) ExitTableSourceBase(ctx *TableSourceBaseContext) {}

// EnterTableSourceNested is called when production tableSourceNested is entered.
func (s *BaseMySqlParserListener) EnterTableSourceNested(ctx *TableSourceNestedContext) {}

// ExitTableSourceNested is called when production tableSourceNested is exited.
func (s *BaseMySqlParserListener) ExitTableSourceNested(ctx *TableSourceNestedContext) {}

// EnterAtomTableItem is called when production atomTableItem is entered.
func (s *BaseMySqlParserListener) EnterAtomTableItem(ctx *AtomTableItemContext) {}

// ExitAtomTableItem is called when production atomTableItem is exited.
func (s *BaseMySqlParserListener) ExitAtomTableItem(ctx *AtomTableItemContext) {}

// EnterSubqueryTableItem is called when production subqueryTableItem is entered.
func (s *BaseMySqlParserListener) EnterSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// ExitSubqueryTableItem is called when production subqueryTableItem is exited.
func (s *BaseMySqlParserListener) ExitSubqueryTableItem(ctx *SubqueryTableItemContext) {}

// EnterTableSourcesItem is called when production tableSourcesItem is entered.
func (s *BaseMySqlParserListener) EnterTableSourcesItem(ctx *TableSourcesItemContext) {}

// ExitTableSourcesItem is called when production tableSourcesItem is exited.
func (s *BaseMySqlParserListener) ExitTableSourcesItem(ctx *TableSourcesItemContext) {}

// EnterIndexHint is called when production indexHint is entered.
func (s *BaseMySqlParserListener) EnterIndexHint(ctx *IndexHintContext) {}

// ExitIndexHint is called when production indexHint is exited.
func (s *BaseMySqlParserListener) ExitIndexHint(ctx *IndexHintContext) {}

// EnterIndexHintType is called when production indexHintType is entered.
func (s *BaseMySqlParserListener) EnterIndexHintType(ctx *IndexHintTypeContext) {}

// ExitIndexHintType is called when production indexHintType is exited.
func (s *BaseMySqlParserListener) ExitIndexHintType(ctx *IndexHintTypeContext) {}

// EnterInnerJoin is called when production innerJoin is entered.
func (s *BaseMySqlParserListener) EnterInnerJoin(ctx *InnerJoinContext) {}

// ExitInnerJoin is called when production innerJoin is exited.
func (s *BaseMySqlParserListener) ExitInnerJoin(ctx *InnerJoinContext) {}

// EnterStraightJoin is called when production straightJoin is entered.
func (s *BaseMySqlParserListener) EnterStraightJoin(ctx *StraightJoinContext) {}

// ExitStraightJoin is called when production straightJoin is exited.
func (s *BaseMySqlParserListener) ExitStraightJoin(ctx *StraightJoinContext) {}

// EnterOuterJoin is called when production outerJoin is entered.
func (s *BaseMySqlParserListener) EnterOuterJoin(ctx *OuterJoinContext) {}

// ExitOuterJoin is called when production outerJoin is exited.
func (s *BaseMySqlParserListener) ExitOuterJoin(ctx *OuterJoinContext) {}

// EnterNaturalJoin is called when production naturalJoin is entered.
func (s *BaseMySqlParserListener) EnterNaturalJoin(ctx *NaturalJoinContext) {}

// ExitNaturalJoin is called when production naturalJoin is exited.
func (s *BaseMySqlParserListener) ExitNaturalJoin(ctx *NaturalJoinContext) {}

// EnterQueryExpression is called when production queryExpression is entered.
func (s *BaseMySqlParserListener) EnterQueryExpression(ctx *QueryExpressionContext) {}

// ExitQueryExpression is called when production queryExpression is exited.
func (s *BaseMySqlParserListener) ExitQueryExpression(ctx *QueryExpressionContext) {}

// EnterQueryExpressionNointo is called when production queryExpressionNointo is entered.
func (s *BaseMySqlParserListener) EnterQueryExpressionNointo(ctx *QueryExpressionNointoContext) {}

// ExitQueryExpressionNointo is called when production queryExpressionNointo is exited.
func (s *BaseMySqlParserListener) ExitQueryExpressionNointo(ctx *QueryExpressionNointoContext) {}

// EnterQuerySpecification1 is called when production querySpecification1 is entered.
func (s *BaseMySqlParserListener) EnterQuerySpecification1(ctx *QuerySpecification1Context) {}

// ExitQuerySpecification1 is called when production querySpecification1 is exited.
func (s *BaseMySqlParserListener) ExitQuerySpecification1(ctx *QuerySpecification1Context) {}

// EnterQuerySpecification2 is called when production querySpecification2 is entered.
func (s *BaseMySqlParserListener) EnterQuerySpecification2(ctx *QuerySpecification2Context) {}

// ExitQuerySpecification2 is called when production querySpecification2 is exited.
func (s *BaseMySqlParserListener) ExitQuerySpecification2(ctx *QuerySpecification2Context) {}

// EnterQuerySpecificationNointo is called when production querySpecificationNointo is entered.
func (s *BaseMySqlParserListener) EnterQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) {
}

// ExitQuerySpecificationNointo is called when production querySpecificationNointo is exited.
func (s *BaseMySqlParserListener) ExitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) {
}

// EnterUnionParenthesis is called when production unionParenthesis is entered.
func (s *BaseMySqlParserListener) EnterUnionParenthesis(ctx *UnionParenthesisContext) {}

// ExitUnionParenthesis is called when production unionParenthesis is exited.
func (s *BaseMySqlParserListener) ExitUnionParenthesis(ctx *UnionParenthesisContext) {}

// EnterUnionStatement is called when production unionStatement is entered.
func (s *BaseMySqlParserListener) EnterUnionStatement(ctx *UnionStatementContext) {}

// ExitUnionStatement is called when production unionStatement is exited.
func (s *BaseMySqlParserListener) ExitUnionStatement(ctx *UnionStatementContext) {}

// EnterUnionStatement2 is called when production unionStatement2 is entered.
func (s *BaseMySqlParserListener) EnterUnionStatement2(ctx *UnionStatement2Context) {}

// ExitUnionStatement2 is called when production unionStatement2 is exited.
func (s *BaseMySqlParserListener) ExitUnionStatement2(ctx *UnionStatement2Context) {}

// EnterUnionstatement3 is called when production unionstatement3 is entered.
func (s *BaseMySqlParserListener) EnterUnionstatement3(ctx *Unionstatement3Context) {}

// ExitUnionstatement3 is called when production unionstatement3 is exited.
func (s *BaseMySqlParserListener) ExitUnionstatement3(ctx *Unionstatement3Context) {}

// EnterSelectSpec is called when production selectSpec is entered.
func (s *BaseMySqlParserListener) EnterSelectSpec(ctx *SelectSpecContext) {}

// ExitSelectSpec is called when production selectSpec is exited.
func (s *BaseMySqlParserListener) ExitSelectSpec(ctx *SelectSpecContext) {}

// EnterSelectElements is called when production selectElements is entered.
func (s *BaseMySqlParserListener) EnterSelectElements(ctx *SelectElementsContext) {}

// ExitSelectElements is called when production selectElements is exited.
func (s *BaseMySqlParserListener) ExitSelectElements(ctx *SelectElementsContext) {}

// EnterSelectStarElement is called when production selectStarElement is entered.
func (s *BaseMySqlParserListener) EnterSelectStarElement(ctx *SelectStarElementContext) {}

// ExitSelectStarElement is called when production selectStarElement is exited.
func (s *BaseMySqlParserListener) ExitSelectStarElement(ctx *SelectStarElementContext) {}

// EnterSelectColumnElement is called when production selectColumnElement is entered.
func (s *BaseMySqlParserListener) EnterSelectColumnElement(ctx *SelectColumnElementContext) {}

// ExitSelectColumnElement is called when production selectColumnElement is exited.
func (s *BaseMySqlParserListener) ExitSelectColumnElement(ctx *SelectColumnElementContext) {}

// EnterSelectFunctionElement is called when production selectFunctionElement is entered.
func (s *BaseMySqlParserListener) EnterSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// ExitSelectFunctionElement is called when production selectFunctionElement is exited.
func (s *BaseMySqlParserListener) ExitSelectFunctionElement(ctx *SelectFunctionElementContext) {}

// EnterSelectExpressionElement is called when production selectExpressionElement is entered.
func (s *BaseMySqlParserListener) EnterSelectExpressionElement(ctx *SelectExpressionElementContext) {}

// ExitSelectExpressionElement is called when production selectExpressionElement is exited.
func (s *BaseMySqlParserListener) ExitSelectExpressionElement(ctx *SelectExpressionElementContext) {}

// EnterSelectIntoVariables is called when production selectIntoVariables is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoVariables(ctx *SelectIntoVariablesContext) {}

// ExitSelectIntoVariables is called when production selectIntoVariables is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoVariables(ctx *SelectIntoVariablesContext) {}

// EnterSelectIntoDumpFile is called when production selectIntoDumpFile is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) {}

// ExitSelectIntoDumpFile is called when production selectIntoDumpFile is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) {}

// EnterSelectIntoTextFile is called when production selectIntoTextFile is entered.
func (s *BaseMySqlParserListener) EnterSelectIntoTextFile(ctx *SelectIntoTextFileContext) {}

// ExitSelectIntoTextFile is called when production selectIntoTextFile is exited.
func (s *BaseMySqlParserListener) ExitSelectIntoTextFile(ctx *SelectIntoTextFileContext) {}

// EnterSelectFieldsInto is called when production selectFieldsInto is entered.
func (s *BaseMySqlParserListener) EnterSelectFieldsInto(ctx *SelectFieldsIntoContext) {}

// ExitSelectFieldsInto is called when production selectFieldsInto is exited.
func (s *BaseMySqlParserListener) ExitSelectFieldsInto(ctx *SelectFieldsIntoContext) {}

// EnterSelectLinesInto is called when production selectLinesInto is entered.
func (s *BaseMySqlParserListener) EnterSelectLinesInto(ctx *SelectLinesIntoContext) {}

// ExitSelectLinesInto is called when production selectLinesInto is exited.
func (s *BaseMySqlParserListener) ExitSelectLinesInto(ctx *SelectLinesIntoContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseMySqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseMySqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterGroupByItem is called when production groupByItem is entered.
func (s *BaseMySqlParserListener) EnterGroupByItem(ctx *GroupByItemContext) {}

// ExitGroupByItem is called when production groupByItem is exited.
func (s *BaseMySqlParserListener) ExitGroupByItem(ctx *GroupByItemContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseMySqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseMySqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitClauseAtom is called when production limitClauseAtom is entered.
func (s *BaseMySqlParserListener) EnterLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// ExitLimitClauseAtom is called when production limitClauseAtom is exited.
func (s *BaseMySqlParserListener) ExitLimitClauseAtom(ctx *LimitClauseAtomContext) {}

// EnterFullId is called when production fullId is entered.
func (s *BaseMySqlParserListener) EnterFullId(ctx *FullIdContext) {}

// ExitFullId is called when production fullId is exited.
func (s *BaseMySqlParserListener) ExitFullId(ctx *FullIdContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseMySqlParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseMySqlParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterFullColumnName is called when production fullColumnName is entered.
func (s *BaseMySqlParserListener) EnterFullColumnName(ctx *FullColumnNameContext) {}

// ExitFullColumnName is called when production fullColumnName is exited.
func (s *BaseMySqlParserListener) ExitFullColumnName(ctx *FullColumnNameContext) {}

// EnterMysqlVariable is called when production mysqlVariable is entered.
func (s *BaseMySqlParserListener) EnterMysqlVariable(ctx *MysqlVariableContext) {}

// ExitMysqlVariable is called when production mysqlVariable is exited.
func (s *BaseMySqlParserListener) ExitMysqlVariable(ctx *MysqlVariableContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseMySqlParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseMySqlParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterUid is called when production uid is entered.
func (s *BaseMySqlParserListener) EnterUid(ctx *UidContext) {}

// ExitUid is called when production uid is exited.
func (s *BaseMySqlParserListener) ExitUid(ctx *UidContext) {}

// EnterEngineName is called when production engineName is entered.
func (s *BaseMySqlParserListener) EnterEngineName(ctx *EngineNameContext) {}

// ExitEngineName is called when production engineName is exited.
func (s *BaseMySqlParserListener) ExitEngineName(ctx *EngineNameContext) {}

// EnterSimpleId is called when production simpleId is entered.
func (s *BaseMySqlParserListener) EnterSimpleId(ctx *SimpleIdContext) {}

// ExitSimpleId is called when production simpleId is exited.
func (s *BaseMySqlParserListener) ExitSimpleId(ctx *SimpleIdContext) {}

// EnterDottedId is called when production dottedId is entered.
func (s *BaseMySqlParserListener) EnterDottedId(ctx *DottedIdContext) {}

// ExitDottedId is called when production dottedId is exited.
func (s *BaseMySqlParserListener) ExitDottedId(ctx *DottedIdContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseMySqlParserListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseMySqlParserListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseMySqlParserListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseMySqlParserListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseMySqlParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseMySqlParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNullNotnull is called when production nullNotnull is entered.
func (s *BaseMySqlParserListener) EnterNullNotnull(ctx *NullNotnullContext) {}

// ExitNullNotnull is called when production nullNotnull is exited.
func (s *BaseMySqlParserListener) ExitNullNotnull(ctx *NullNotnullContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseMySqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseMySqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterHexadecimalLiteral is called when production hexadecimalLiteral is entered.
func (s *BaseMySqlParserListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production hexadecimalLiteral is exited.
func (s *BaseMySqlParserListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseMySqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseMySqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterStringDataType is called when production stringDataType is entered.
func (s *BaseMySqlParserListener) EnterStringDataType(ctx *StringDataTypeContext) {}

// ExitStringDataType is called when production stringDataType is exited.
func (s *BaseMySqlParserListener) ExitStringDataType(ctx *StringDataTypeContext) {}

// EnterNationalStringDataType is called when production nationalStringDataType is entered.
func (s *BaseMySqlParserListener) EnterNationalStringDataType(ctx *NationalStringDataTypeContext) {}

// ExitNationalStringDataType is called when production nationalStringDataType is exited.
func (s *BaseMySqlParserListener) ExitNationalStringDataType(ctx *NationalStringDataTypeContext) {}

// EnterNationalVaryingStringDataType is called when production nationalVaryingStringDataType is entered.
func (s *BaseMySqlParserListener) EnterNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) {
}

// ExitNationalVaryingStringDataType is called when production nationalVaryingStringDataType is exited.
func (s *BaseMySqlParserListener) ExitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) {
}

// EnterDimensionDataType is called when production dimensionDataType is entered.
func (s *BaseMySqlParserListener) EnterDimensionDataType(ctx *DimensionDataTypeContext) {}

// ExitDimensionDataType is called when production dimensionDataType is exited.
func (s *BaseMySqlParserListener) ExitDimensionDataType(ctx *DimensionDataTypeContext) {}

// EnterSimpleDataType is called when production simpleDataType is entered.
func (s *BaseMySqlParserListener) EnterSimpleDataType(ctx *SimpleDataTypeContext) {}

// ExitSimpleDataType is called when production simpleDataType is exited.
func (s *BaseMySqlParserListener) ExitSimpleDataType(ctx *SimpleDataTypeContext) {}

// EnterCollectionDataType is called when production collectionDataType is entered.
func (s *BaseMySqlParserListener) EnterCollectionDataType(ctx *CollectionDataTypeContext) {}

// ExitCollectionDataType is called when production collectionDataType is exited.
func (s *BaseMySqlParserListener) ExitCollectionDataType(ctx *CollectionDataTypeContext) {}

// EnterSpatialDataType is called when production spatialDataType is entered.
func (s *BaseMySqlParserListener) EnterSpatialDataType(ctx *SpatialDataTypeContext) {}

// ExitSpatialDataType is called when production spatialDataType is exited.
func (s *BaseMySqlParserListener) ExitSpatialDataType(ctx *SpatialDataTypeContext) {}

// EnterLongVarcharDataType is called when production longVarcharDataType is entered.
func (s *BaseMySqlParserListener) EnterLongVarcharDataType(ctx *LongVarcharDataTypeContext) {}

// ExitLongVarcharDataType is called when production longVarcharDataType is exited.
func (s *BaseMySqlParserListener) ExitLongVarcharDataType(ctx *LongVarcharDataTypeContext) {}

// EnterLongVarbinaryDataType is called when production longVarbinaryDataType is entered.
func (s *BaseMySqlParserListener) EnterLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) {}

// ExitLongVarbinaryDataType is called when production longVarbinaryDataType is exited.
func (s *BaseMySqlParserListener) ExitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) {}

// EnterCollectionOptions is called when production collectionOptions is entered.
func (s *BaseMySqlParserListener) EnterCollectionOptions(ctx *CollectionOptionsContext) {}

// ExitCollectionOptions is called when production collectionOptions is exited.
func (s *BaseMySqlParserListener) ExitCollectionOptions(ctx *CollectionOptionsContext) {}

// EnterConvertedDataType is called when production convertedDataType is entered.
func (s *BaseMySqlParserListener) EnterConvertedDataType(ctx *ConvertedDataTypeContext) {}

// ExitConvertedDataType is called when production convertedDataType is exited.
func (s *BaseMySqlParserListener) ExitConvertedDataType(ctx *ConvertedDataTypeContext) {}

// EnterLengthOneDimension is called when production lengthOneDimension is entered.
func (s *BaseMySqlParserListener) EnterLengthOneDimension(ctx *LengthOneDimensionContext) {}

// ExitLengthOneDimension is called when production lengthOneDimension is exited.
func (s *BaseMySqlParserListener) ExitLengthOneDimension(ctx *LengthOneDimensionContext) {}

// EnterLengthTwoDimension is called when production lengthTwoDimension is entered.
func (s *BaseMySqlParserListener) EnterLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// ExitLengthTwoDimension is called when production lengthTwoDimension is exited.
func (s *BaseMySqlParserListener) ExitLengthTwoDimension(ctx *LengthTwoDimensionContext) {}

// EnterLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is entered.
func (s *BaseMySqlParserListener) EnterLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// ExitLengthTwoOptionalDimension is called when production lengthTwoOptionalDimension is exited.
func (s *BaseMySqlParserListener) ExitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) {
}

// EnterUidList is called when production uidList is entered.
func (s *BaseMySqlParserListener) EnterUidList(ctx *UidListContext) {}

// ExitUidList is called when production uidList is exited.
func (s *BaseMySqlParserListener) ExitUidList(ctx *UidListContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseMySqlParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseMySqlParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterIntervalType is called when production intervalType is entered.
func (s *BaseMySqlParserListener) EnterIntervalType(ctx *IntervalTypeContext) {}

// ExitIntervalType is called when production intervalType is exited.
func (s *BaseMySqlParserListener) ExitIntervalType(ctx *IntervalTypeContext) {}

// EnterSpecificFunctionCall is called when production specificFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// ExitSpecificFunctionCall is called when production specificFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitSpecificFunctionCall(ctx *SpecificFunctionCallContext) {}

// EnterAggregateFunctionCall is called when production aggregateFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// ExitAggregateFunctionCall is called when production aggregateFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitAggregateFunctionCall(ctx *AggregateFunctionCallContext) {}

// EnterScalarFunctionCall is called when production scalarFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// ExitScalarFunctionCall is called when production scalarFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitScalarFunctionCall(ctx *ScalarFunctionCallContext) {}

// EnterUdfFunctionCall is called when production udfFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// ExitUdfFunctionCall is called when production udfFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitUdfFunctionCall(ctx *UdfFunctionCallContext) {}

// EnterPasswordFunctionCall is called when production passwordFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterPasswordFunctionCall(ctx *PasswordFunctionCallContext) {}

// ExitPasswordFunctionCall is called when production passwordFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitPasswordFunctionCall(ctx *PasswordFunctionCallContext) {}

// EnterSimpleFunctionCall is called when production simpleFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// ExitSimpleFunctionCall is called when production simpleFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// EnterDataTypeFunctionCall is called when production dataTypeFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) {}

// ExitDataTypeFunctionCall is called when production dataTypeFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) {}

// EnterValuesFunctionCall is called when production valuesFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterValuesFunctionCall(ctx *ValuesFunctionCallContext) {}

// ExitValuesFunctionCall is called when production valuesFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitValuesFunctionCall(ctx *ValuesFunctionCallContext) {}

// EnterCaseFunctionCall is called when production caseFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterCaseFunctionCall(ctx *CaseFunctionCallContext) {}

// ExitCaseFunctionCall is called when production caseFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitCaseFunctionCall(ctx *CaseFunctionCallContext) {}

// EnterCharFunctionCall is called when production charFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterCharFunctionCall(ctx *CharFunctionCallContext) {}

// ExitCharFunctionCall is called when production charFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitCharFunctionCall(ctx *CharFunctionCallContext) {}

// EnterPositionFunctionCall is called when production positionFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterPositionFunctionCall(ctx *PositionFunctionCallContext) {}

// ExitPositionFunctionCall is called when production positionFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitPositionFunctionCall(ctx *PositionFunctionCallContext) {}

// EnterSubstrFunctionCall is called when production substrFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterSubstrFunctionCall(ctx *SubstrFunctionCallContext) {}

// ExitSubstrFunctionCall is called when production substrFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitSubstrFunctionCall(ctx *SubstrFunctionCallContext) {}

// EnterTrimFunctionCall is called when production trimFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterTrimFunctionCall(ctx *TrimFunctionCallContext) {}

// ExitTrimFunctionCall is called when production trimFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitTrimFunctionCall(ctx *TrimFunctionCallContext) {}

// EnterWeightFunctionCall is called when production weightFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterWeightFunctionCall(ctx *WeightFunctionCallContext) {}

// ExitWeightFunctionCall is called when production weightFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitWeightFunctionCall(ctx *WeightFunctionCallContext) {}

// EnterExtractFunctionCall is called when production extractFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterExtractFunctionCall(ctx *ExtractFunctionCallContext) {}

// ExitExtractFunctionCall is called when production extractFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitExtractFunctionCall(ctx *ExtractFunctionCallContext) {}

// EnterGetFormatFunctionCall is called when production getFormatFunctionCall is entered.
func (s *BaseMySqlParserListener) EnterGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) {}

// ExitGetFormatFunctionCall is called when production getFormatFunctionCall is exited.
func (s *BaseMySqlParserListener) ExitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) {}

// EnterCaseFuncAlternative is called when production caseFuncAlternative is entered.
func (s *BaseMySqlParserListener) EnterCaseFuncAlternative(ctx *CaseFuncAlternativeContext) {}

// ExitCaseFuncAlternative is called when production caseFuncAlternative is exited.
func (s *BaseMySqlParserListener) ExitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) {}

// EnterLevelWeightList is called when production levelWeightList is entered.
func (s *BaseMySqlParserListener) EnterLevelWeightList(ctx *LevelWeightListContext) {}

// ExitLevelWeightList is called when production levelWeightList is exited.
func (s *BaseMySqlParserListener) ExitLevelWeightList(ctx *LevelWeightListContext) {}

// EnterLevelWeightRange is called when production levelWeightRange is entered.
func (s *BaseMySqlParserListener) EnterLevelWeightRange(ctx *LevelWeightRangeContext) {}

// ExitLevelWeightRange is called when production levelWeightRange is exited.
func (s *BaseMySqlParserListener) ExitLevelWeightRange(ctx *LevelWeightRangeContext) {}

// EnterLevelInWeightListElement is called when production levelInWeightListElement is entered.
func (s *BaseMySqlParserListener) EnterLevelInWeightListElement(ctx *LevelInWeightListElementContext) {
}

// ExitLevelInWeightListElement is called when production levelInWeightListElement is exited.
func (s *BaseMySqlParserListener) ExitLevelInWeightListElement(ctx *LevelInWeightListElementContext) {
}

// EnterAggregateWindowedFunction is called when production aggregateWindowedFunction is entered.
func (s *BaseMySqlParserListener) EnterAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) {
}

// ExitAggregateWindowedFunction is called when production aggregateWindowedFunction is exited.
func (s *BaseMySqlParserListener) ExitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) {
}

// EnterScalarFunctionName is called when production scalarFunctionName is entered.
func (s *BaseMySqlParserListener) EnterScalarFunctionName(ctx *ScalarFunctionNameContext) {}

// ExitScalarFunctionName is called when production scalarFunctionName is exited.
func (s *BaseMySqlParserListener) ExitScalarFunctionName(ctx *ScalarFunctionNameContext) {}

// EnterPasswordFunctionClause is called when production passwordFunctionClause is entered.
func (s *BaseMySqlParserListener) EnterPasswordFunctionClause(ctx *PasswordFunctionClauseContext) {}

// ExitPasswordFunctionClause is called when production passwordFunctionClause is exited.
func (s *BaseMySqlParserListener) ExitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseMySqlParserListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseMySqlParserListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterFunctionArg is called when production functionArg is entered.
func (s *BaseMySqlParserListener) EnterFunctionArg(ctx *FunctionArgContext) {}

// ExitFunctionArg is called when production functionArg is exited.
func (s *BaseMySqlParserListener) ExitFunctionArg(ctx *FunctionArgContext) {}

// EnterIsExpression is called when production isExpression is entered.
func (s *BaseMySqlParserListener) EnterIsExpression(ctx *IsExpressionContext) {}

// ExitIsExpression is called when production isExpression is exited.
func (s *BaseMySqlParserListener) ExitIsExpression(ctx *IsExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseMySqlParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseMySqlParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseMySqlParserListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseMySqlParserListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseMySqlParserListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseMySqlParserListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterSoundsLikePredicate is called when production soundsLikePredicate is entered.
func (s *BaseMySqlParserListener) EnterSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// ExitSoundsLikePredicate is called when production soundsLikePredicate is exited.
func (s *BaseMySqlParserListener) ExitSoundsLikePredicate(ctx *SoundsLikePredicateContext) {}

// EnterExpressionAtomPredicate is called when production expressionAtomPredicate is entered.
func (s *BaseMySqlParserListener) EnterExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// ExitExpressionAtomPredicate is called when production expressionAtomPredicate is exited.
func (s *BaseMySqlParserListener) ExitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) {}

// EnterJsonMemberOfPredicate is called when production jsonMemberOfPredicate is entered.
func (s *BaseMySqlParserListener) EnterJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) {}

// ExitJsonMemberOfPredicate is called when production jsonMemberOfPredicate is exited.
func (s *BaseMySqlParserListener) ExitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) {}

// EnterInPredicate is called when production inPredicate is entered.
func (s *BaseMySqlParserListener) EnterInPredicate(ctx *InPredicateContext) {}

// ExitInPredicate is called when production inPredicate is exited.
func (s *BaseMySqlParserListener) ExitInPredicate(ctx *InPredicateContext) {}

// EnterSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {
}

// ExitSubqueryComparasionPredicate is called when production subqueryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitSubqueryComparasionPredicate(ctx *SubqueryComparasionPredicateContext) {
}

// EnterBetweenPredicate is called when production betweenPredicate is entered.
func (s *BaseMySqlParserListener) EnterBetweenPredicate(ctx *BetweenPredicateContext) {}

// ExitBetweenPredicate is called when production betweenPredicate is exited.
func (s *BaseMySqlParserListener) ExitBetweenPredicate(ctx *BetweenPredicateContext) {}

// EnterBinaryComparasionPredicate is called when production binaryComparasionPredicate is entered.
func (s *BaseMySqlParserListener) EnterBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// ExitBinaryComparasionPredicate is called when production binaryComparasionPredicate is exited.
func (s *BaseMySqlParserListener) ExitBinaryComparasionPredicate(ctx *BinaryComparasionPredicateContext) {
}

// EnterIsNullPredicate is called when production isNullPredicate is entered.
func (s *BaseMySqlParserListener) EnterIsNullPredicate(ctx *IsNullPredicateContext) {}

// ExitIsNullPredicate is called when production isNullPredicate is exited.
func (s *BaseMySqlParserListener) ExitIsNullPredicate(ctx *IsNullPredicateContext) {}

// EnterLikePredicate is called when production likePredicate is entered.
func (s *BaseMySqlParserListener) EnterLikePredicate(ctx *LikePredicateContext) {}

// ExitLikePredicate is called when production likePredicate is exited.
func (s *BaseMySqlParserListener) ExitLikePredicate(ctx *LikePredicateContext) {}

// EnterRegexpPredicate is called when production regexpPredicate is entered.
func (s *BaseMySqlParserListener) EnterRegexpPredicate(ctx *RegexpPredicateContext) {}

// ExitRegexpPredicate is called when production regexpPredicate is exited.
func (s *BaseMySqlParserListener) ExitRegexpPredicate(ctx *RegexpPredicateContext) {}

// EnterUnaryExpressionAtom is called when production unaryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// ExitUnaryExpressionAtom is called when production unaryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) {}

// EnterCollateExpressionAtom is called when production collateExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterCollateExpressionAtom(ctx *CollateExpressionAtomContext) {}

// ExitCollateExpressionAtom is called when production collateExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitCollateExpressionAtom(ctx *CollateExpressionAtomContext) {}

// EnterSubqueryExpessionAtom is called when production subqueryExpessionAtom is entered.
func (s *BaseMySqlParserListener) EnterSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) {}

// ExitSubqueryExpessionAtom is called when production subqueryExpessionAtom is exited.
func (s *BaseMySqlParserListener) ExitSubqueryExpessionAtom(ctx *SubqueryExpessionAtomContext) {}

// EnterMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// ExitMysqlVariableExpressionAtom is called when production mysqlVariableExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) {
}

// EnterNestedExpressionAtom is called when production nestedExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// ExitNestedExpressionAtom is called when production nestedExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitNestedExpressionAtom(ctx *NestedExpressionAtomContext) {}

// EnterNestedRowExpressionAtom is called when production nestedRowExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) {}

// ExitNestedRowExpressionAtom is called when production nestedRowExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) {}

// EnterMathExpressionAtom is called when production mathExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// ExitMathExpressionAtom is called when production mathExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitMathExpressionAtom(ctx *MathExpressionAtomContext) {}

// EnterIntervalExpressionAtom is called when production intervalExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// ExitIntervalExpressionAtom is called when production intervalExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) {}

// EnterJsonExpressionAtom is called when production jsonExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterJsonExpressionAtom(ctx *JsonExpressionAtomContext) {}

// ExitJsonExpressionAtom is called when production jsonExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitJsonExpressionAtom(ctx *JsonExpressionAtomContext) {}

// EnterExistsExpessionAtom is called when production existsExpessionAtom is entered.
func (s *BaseMySqlParserListener) EnterExistsExpessionAtom(ctx *ExistsExpessionAtomContext) {}

// ExitExistsExpessionAtom is called when production existsExpessionAtom is exited.
func (s *BaseMySqlParserListener) ExitExistsExpessionAtom(ctx *ExistsExpessionAtomContext) {}

// EnterConstantExpressionAtom is called when production constantExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// ExitConstantExpressionAtom is called when production constantExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) {}

// EnterFunctionCallExpressionAtom is called when production functionCallExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// ExitFunctionCallExpressionAtom is called when production functionCallExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) {
}

// EnterBinaryExpressionAtom is called when production binaryExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// ExitBinaryExpressionAtom is called when production binaryExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) {}

// EnterFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// ExitFullColumnNameExpressionAtom is called when production fullColumnNameExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) {
}

// EnterBitExpressionAtom is called when production bitExpressionAtom is entered.
func (s *BaseMySqlParserListener) EnterBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// ExitBitExpressionAtom is called when production bitExpressionAtom is exited.
func (s *BaseMySqlParserListener) ExitBitExpressionAtom(ctx *BitExpressionAtomContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseMySqlParserListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseMySqlParserListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseMySqlParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseMySqlParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseMySqlParserListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseMySqlParserListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterBitOperator is called when production bitOperator is entered.
func (s *BaseMySqlParserListener) EnterBitOperator(ctx *BitOperatorContext) {}

// ExitBitOperator is called when production bitOperator is exited.
func (s *BaseMySqlParserListener) ExitBitOperator(ctx *BitOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseMySqlParserListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseMySqlParserListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterJsonOperator is called when production jsonOperator is entered.
func (s *BaseMySqlParserListener) EnterJsonOperator(ctx *JsonOperatorContext) {}

// ExitJsonOperator is called when production jsonOperator is exited.
func (s *BaseMySqlParserListener) ExitJsonOperator(ctx *JsonOperatorContext) {}

// EnterCharsetNameBase is called when production charsetNameBase is entered.
func (s *BaseMySqlParserListener) EnterCharsetNameBase(ctx *CharsetNameBaseContext) {}

// ExitCharsetNameBase is called when production charsetNameBase is exited.
func (s *BaseMySqlParserListener) ExitCharsetNameBase(ctx *CharsetNameBaseContext) {}

// EnterTransactionLevelBase is called when production transactionLevelBase is entered.
func (s *BaseMySqlParserListener) EnterTransactionLevelBase(ctx *TransactionLevelBaseContext) {}

// ExitTransactionLevelBase is called when production transactionLevelBase is exited.
func (s *BaseMySqlParserListener) ExitTransactionLevelBase(ctx *TransactionLevelBaseContext) {}

// EnterPrivilegesBase is called when production privilegesBase is entered.
func (s *BaseMySqlParserListener) EnterPrivilegesBase(ctx *PrivilegesBaseContext) {}

// ExitPrivilegesBase is called when production privilegesBase is exited.
func (s *BaseMySqlParserListener) ExitPrivilegesBase(ctx *PrivilegesBaseContext) {}

// EnterIntervalTypeBase is called when production intervalTypeBase is entered.
func (s *BaseMySqlParserListener) EnterIntervalTypeBase(ctx *IntervalTypeBaseContext) {}

// ExitIntervalTypeBase is called when production intervalTypeBase is exited.
func (s *BaseMySqlParserListener) ExitIntervalTypeBase(ctx *IntervalTypeBaseContext) {}

// EnterDataTypeBase is called when production dataTypeBase is entered.
func (s *BaseMySqlParserListener) EnterDataTypeBase(ctx *DataTypeBaseContext) {}

// ExitDataTypeBase is called when production dataTypeBase is exited.
func (s *BaseMySqlParserListener) ExitDataTypeBase(ctx *DataTypeBaseContext) {}

// EnterKeywordsCanBeId is called when production keywordsCanBeId is entered.
func (s *BaseMySqlParserListener) EnterKeywordsCanBeId(ctx *KeywordsCanBeIdContext) {}

// ExitKeywordsCanBeId is called when production keywordsCanBeId is exited.
func (s *BaseMySqlParserListener) ExitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) {}

// EnterFunctionNameBase is called when production functionNameBase is entered.
func (s *BaseMySqlParserListener) EnterFunctionNameBase(ctx *FunctionNameBaseContext) {}

// ExitFunctionNameBase is called when production functionNameBase is exited.
func (s *BaseMySqlParserListener) ExitFunctionNameBase(ctx *FunctionNameBaseContext) {}

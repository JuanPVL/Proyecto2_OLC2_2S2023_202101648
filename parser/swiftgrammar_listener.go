// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterStructCreation is called when entering the structCreation production.
	EnterStructCreation(c *StructCreationContext)

	// EnterListStructDec is called when entering the listStructDec production.
	EnterListStructDec(c *ListStructDecContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterBlockelsif is called when entering the blockelsif production.
	EnterBlockelsif(c *BlockelsifContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterListcases is called when entering the listcases production.
	EnterListcases(c *ListcasesContext)

	// EnterCasestmt is called when entering the casestmt production.
	EnterCasestmt(c *CasestmtContext)

	// EnterInstdefault is called when entering the instdefault production.
	EnterInstdefault(c *InstdefaultContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterDeclarationstmt is called when entering the declarationstmt production.
	EnterDeclarationstmt(c *DeclarationstmtContext)

	// EnterAsignationstmt is called when entering the asignationstmt production.
	EnterAsignationstmt(c *AsignationstmtContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterListParamsFunc is called when entering the listParamsFunc production.
	EnterListParamsFunc(c *ListParamsFuncContext)

	// EnterCallFuncionIns is called when entering the callFuncionIns production.
	EnterCallFuncionIns(c *CallFuncionInsContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTypesmatriz is called when entering the typesmatriz production.
	EnterTypesmatriz(c *TypesmatrizContext)

	// EnterExprFor is called when entering the exprFor production.
	EnterExprFor(c *ExprForContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterConversionstmt is called when entering the conversionstmt production.
	EnterConversionstmt(c *ConversionstmtContext)

	// EnterExprvector is called when entering the exprvector production.
	EnterExprvector(c *ExprvectorContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

	// EnterListArray is called when entering the listArray production.
	EnterListArray(c *ListArrayContext)

	// EnterListAccessArray is called when entering the listAccessArray production.
	EnterListAccessArray(c *ListAccessArrayContext)

	// EnterCallFuncion is called when entering the callFuncion production.
	EnterCallFuncion(c *CallFuncionContext)

	// EnterListParamsCall is called when entering the listParamsCall production.
	EnterListParamsCall(c *ListParamsCallContext)

	// EnterListStructExp is called when entering the listStructExp production.
	EnterListStructExp(c *ListStructExpContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitStructCreation is called when exiting the structCreation production.
	ExitStructCreation(c *StructCreationContext)

	// ExitListStructDec is called when exiting the listStructDec production.
	ExitListStructDec(c *ListStructDecContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitBlockelsif is called when exiting the blockelsif production.
	ExitBlockelsif(c *BlockelsifContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitListcases is called when exiting the listcases production.
	ExitListcases(c *ListcasesContext)

	// ExitCasestmt is called when exiting the casestmt production.
	ExitCasestmt(c *CasestmtContext)

	// ExitInstdefault is called when exiting the instdefault production.
	ExitInstdefault(c *InstdefaultContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitDeclarationstmt is called when exiting the declarationstmt production.
	ExitDeclarationstmt(c *DeclarationstmtContext)

	// ExitAsignationstmt is called when exiting the asignationstmt production.
	ExitAsignationstmt(c *AsignationstmtContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitListParamsFunc is called when exiting the listParamsFunc production.
	ExitListParamsFunc(c *ListParamsFuncContext)

	// ExitCallFuncionIns is called when exiting the callFuncionIns production.
	ExitCallFuncionIns(c *CallFuncionInsContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTypesmatriz is called when exiting the typesmatriz production.
	ExitTypesmatriz(c *TypesmatrizContext)

	// ExitExprFor is called when exiting the exprFor production.
	ExitExprFor(c *ExprForContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitConversionstmt is called when exiting the conversionstmt production.
	ExitConversionstmt(c *ConversionstmtContext)

	// ExitExprvector is called when exiting the exprvector production.
	ExitExprvector(c *ExprvectorContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

	// ExitListArray is called when exiting the listArray production.
	ExitListArray(c *ListArrayContext)

	// ExitListAccessArray is called when exiting the listAccessArray production.
	ExitListAccessArray(c *ListAccessArrayContext)

	// ExitCallFuncion is called when exiting the callFuncion production.
	ExitCallFuncion(c *CallFuncionContext)

	// ExitListParamsCall is called when exiting the listParamsCall production.
	ExitListParamsCall(c *ListParamsCallContext)

	// ExitListStructExp is called when exiting the listStructExp production.
	ExitListStructExp(c *ListStructExpContext)
}

// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterStructCreation is called when production structCreation is entered.
func (s *BaseSwiftGrammarListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production structCreation is exited.
func (s *BaseSwiftGrammarListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterListStructDec is called when production listStructDec is entered.
func (s *BaseSwiftGrammarListener) EnterListStructDec(ctx *ListStructDecContext) {}

// ExitListStructDec is called when production listStructDec is exited.
func (s *BaseSwiftGrammarListener) ExitListStructDec(ctx *ListStructDecContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterBlockelsif is called when production blockelsif is entered.
func (s *BaseSwiftGrammarListener) EnterBlockelsif(ctx *BlockelsifContext) {}

// ExitBlockelsif is called when production blockelsif is exited.
func (s *BaseSwiftGrammarListener) ExitBlockelsif(ctx *BlockelsifContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterDeclarationstmt is called when production declarationstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDeclarationstmt(ctx *DeclarationstmtContext) {}

// ExitDeclarationstmt is called when production declarationstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDeclarationstmt(ctx *DeclarationstmtContext) {}

// EnterAsignationstmt is called when production asignationstmt is entered.
func (s *BaseSwiftGrammarListener) EnterAsignationstmt(ctx *AsignationstmtContext) {}

// ExitAsignationstmt is called when production asignationstmt is exited.
func (s *BaseSwiftGrammarListener) ExitAsignationstmt(ctx *AsignationstmtContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseSwiftGrammarListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseSwiftGrammarListener) ExitFunction(ctx *FunctionContext) {}

// EnterListParamsFunc is called when production listParamsFunc is entered.
func (s *BaseSwiftGrammarListener) EnterListParamsFunc(ctx *ListParamsFuncContext) {}

// ExitListParamsFunc is called when production listParamsFunc is exited.
func (s *BaseSwiftGrammarListener) ExitListParamsFunc(ctx *ListParamsFuncContext) {}

// EnterCallFuncionIns is called when production callFuncionIns is entered.
func (s *BaseSwiftGrammarListener) EnterCallFuncionIns(ctx *CallFuncionInsContext) {}

// ExitCallFuncionIns is called when production callFuncionIns is exited.
func (s *BaseSwiftGrammarListener) ExitCallFuncionIns(ctx *CallFuncionInsContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseSwiftGrammarListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseSwiftGrammarListener) ExitTypes(ctx *TypesContext) {}

// EnterTypesmatriz is called when production typesmatriz is entered.
func (s *BaseSwiftGrammarListener) EnterTypesmatriz(ctx *TypesmatrizContext) {}

// ExitTypesmatriz is called when production typesmatriz is exited.
func (s *BaseSwiftGrammarListener) ExitTypesmatriz(ctx *TypesmatrizContext) {}

// EnterExprFor is called when production exprFor is entered.
func (s *BaseSwiftGrammarListener) EnterExprFor(ctx *ExprForContext) {}

// ExitExprFor is called when production exprFor is exited.
func (s *BaseSwiftGrammarListener) ExitExprFor(ctx *ExprForContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterConversionstmt is called when production conversionstmt is entered.
func (s *BaseSwiftGrammarListener) EnterConversionstmt(ctx *ConversionstmtContext) {}

// ExitConversionstmt is called when production conversionstmt is exited.
func (s *BaseSwiftGrammarListener) ExitConversionstmt(ctx *ConversionstmtContext) {}

// EnterExprvector is called when production exprvector is entered.
func (s *BaseSwiftGrammarListener) EnterExprvector(ctx *ExprvectorContext) {}

// ExitExprvector is called when production exprvector is exited.
func (s *BaseSwiftGrammarListener) ExitExprvector(ctx *ExprvectorContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseSwiftGrammarListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseSwiftGrammarListener) ExitListParams(ctx *ListParamsContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseSwiftGrammarListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseSwiftGrammarListener) ExitListArray(ctx *ListArrayContext) {}

// EnterListAccessArray is called when production listAccessArray is entered.
func (s *BaseSwiftGrammarListener) EnterListAccessArray(ctx *ListAccessArrayContext) {}

// ExitListAccessArray is called when production listAccessArray is exited.
func (s *BaseSwiftGrammarListener) ExitListAccessArray(ctx *ListAccessArrayContext) {}

// EnterCallFuncion is called when production callFuncion is entered.
func (s *BaseSwiftGrammarListener) EnterCallFuncion(ctx *CallFuncionContext) {}

// ExitCallFuncion is called when production callFuncion is exited.
func (s *BaseSwiftGrammarListener) ExitCallFuncion(ctx *CallFuncionContext) {}

// EnterListParamsCall is called when production listParamsCall is entered.
func (s *BaseSwiftGrammarListener) EnterListParamsCall(ctx *ListParamsCallContext) {}

// ExitListParamsCall is called when production listParamsCall is exited.
func (s *BaseSwiftGrammarListener) ExitListParamsCall(ctx *ListParamsCallContext) {}

// EnterListStructExp is called when production listStructExp is entered.
func (s *BaseSwiftGrammarListener) EnterListStructExp(ctx *ListStructExpContext) {}

// ExitListStructExp is called when production listStructExp is exited.
func (s *BaseSwiftGrammarListener) ExitListStructExp(ctx *ListStructExpContext) {}

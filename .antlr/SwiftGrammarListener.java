// Generated from c:/Users/USUARIO/Desktop/Proyecto2_OLC2_2S2023_202101648/SwiftGrammar.g4 by ANTLR 4.13.1

    import "Proyecto1_OLC2_2S2023_202101648/interfaces"
    import "Proyecto1_OLC2_2S2023_202101648/Environment"
    import "Proyecto1_OLC2_2S2023_202101648/expressions"
    import "Proyecto1_OLC2_2S2023_202101648/instructions"
    import "strings"

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftGrammarParser}.
 */
public interface SwiftGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftGrammarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftGrammarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#structCreation}.
	 * @param ctx the parse tree
	 */
	void enterStructCreation(SwiftGrammarParser.StructCreationContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#structCreation}.
	 * @param ctx the parse tree
	 */
	void exitStructCreation(SwiftGrammarParser.StructCreationContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listStructDec}.
	 * @param ctx the parse tree
	 */
	void enterListStructDec(SwiftGrammarParser.ListStructDecContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listStructDec}.
	 * @param ctx the parse tree
	 */
	void exitListStructDec(SwiftGrammarParser.ListStructDecContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#blockelsif}.
	 * @param ctx the parse tree
	 */
	void enterBlockelsif(SwiftGrammarParser.BlockelsifContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#blockelsif}.
	 * @param ctx the parse tree
	 */
	void exitBlockelsif(SwiftGrammarParser.BlockelsifContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void enterIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void exitIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void enterWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void exitWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void enterGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void exitGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void enterForstmt(SwiftGrammarParser.ForstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void exitForstmt(SwiftGrammarParser.ForstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#declarationstmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclarationstmt(SwiftGrammarParser.DeclarationstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#declarationstmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclarationstmt(SwiftGrammarParser.DeclarationstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#asignationstmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignationstmt(SwiftGrammarParser.AsignationstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#asignationstmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignationstmt(SwiftGrammarParser.AsignationstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#function}.
	 * @param ctx the parse tree
	 */
	void enterFunction(SwiftGrammarParser.FunctionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#function}.
	 * @param ctx the parse tree
	 */
	void exitFunction(SwiftGrammarParser.FunctionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParamsFunc}.
	 * @param ctx the parse tree
	 */
	void enterListParamsFunc(SwiftGrammarParser.ListParamsFuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParamsFunc}.
	 * @param ctx the parse tree
	 */
	void exitListParamsFunc(SwiftGrammarParser.ListParamsFuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callFuncionIns}.
	 * @param ctx the parse tree
	 */
	void enterCallFuncionIns(SwiftGrammarParser.CallFuncionInsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callFuncionIns}.
	 * @param ctx the parse tree
	 */
	void exitCallFuncionIns(SwiftGrammarParser.CallFuncionInsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#types}.
	 * @param ctx the parse tree
	 */
	void enterTypes(SwiftGrammarParser.TypesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#types}.
	 * @param ctx the parse tree
	 */
	void exitTypes(SwiftGrammarParser.TypesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#typesmatriz}.
	 * @param ctx the parse tree
	 */
	void enterTypesmatriz(SwiftGrammarParser.TypesmatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#typesmatriz}.
	 * @param ctx the parse tree
	 */
	void exitTypesmatriz(SwiftGrammarParser.TypesmatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#exprFor}.
	 * @param ctx the parse tree
	 */
	void enterExprFor(SwiftGrammarParser.ExprForContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#exprFor}.
	 * @param ctx the parse tree
	 */
	void exitExprFor(SwiftGrammarParser.ExprForContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#conversionstmt}.
	 * @param ctx the parse tree
	 */
	void enterConversionstmt(SwiftGrammarParser.ConversionstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#conversionstmt}.
	 * @param ctx the parse tree
	 */
	void exitConversionstmt(SwiftGrammarParser.ConversionstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#exprvector}.
	 * @param ctx the parse tree
	 */
	void enterExprvector(SwiftGrammarParser.ExprvectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#exprvector}.
	 * @param ctx the parse tree
	 */
	void exitExprvector(SwiftGrammarParser.ExprvectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParams}.
	 * @param ctx the parse tree
	 */
	void enterListParams(SwiftGrammarParser.ListParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParams}.
	 * @param ctx the parse tree
	 */
	void exitListParams(SwiftGrammarParser.ListParamsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listArray}.
	 * @param ctx the parse tree
	 */
	void enterListArray(SwiftGrammarParser.ListArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listArray}.
	 * @param ctx the parse tree
	 */
	void exitListArray(SwiftGrammarParser.ListArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callFuncion}.
	 * @param ctx the parse tree
	 */
	void enterCallFuncion(SwiftGrammarParser.CallFuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callFuncion}.
	 * @param ctx the parse tree
	 */
	void exitCallFuncion(SwiftGrammarParser.CallFuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listParamsCall}.
	 * @param ctx the parse tree
	 */
	void enterListParamsCall(SwiftGrammarParser.ListParamsCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listParamsCall}.
	 * @param ctx the parse tree
	 */
	void exitListParamsCall(SwiftGrammarParser.ListParamsCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listStructExp}.
	 * @param ctx the parse tree
	 */
	void enterListStructExp(SwiftGrammarParser.ListStructExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listStructExp}.
	 * @param ctx the parse tree
	 */
	void exitListStructExp(SwiftGrammarParser.ListStructExpContext ctx);
}
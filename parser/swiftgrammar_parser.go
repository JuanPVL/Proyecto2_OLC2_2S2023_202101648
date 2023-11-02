// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Proyecto2_OLC2_2S2023_202101648/interfaces"
import "Proyecto2_OLC2_2S2023_202101648/Environment"
import "Proyecto2_OLC2_2S2023_202101648/expressions"
import "Proyecto2_OLC2_2S2023_202101648/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", "'print'",
		"'if'", "'else'", "'while'", "'for'", "'in'", "'switch'", "'case'",
		"'default'", "'var'", "'let'", "'nil'", "'break'", "'continue'", "'append'",
		"'removeLast'", "'remove'", "'at'", "'isEmpty'", "'count'", "'array'",
		"'return'", "'func'", "'struct'", "'guard'", "", "", "", "", "'!='",
		"'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'",
		"'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'",
		"'['", "']'", "','", "'?'", "';'", "'.'", "'->'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "IN", "SWITCH", "CASE", "DEFAULT", "VAR", "LET", "NIL",
		"BREAK", "CONTINUE", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY",
		"COUNT", "ARRAY", "RETURN", "FUNC", "STRUCT", "GUARD", "NUMBER", "STRING",
		"CHAR", "ID", "DIFE", "IG_IG", "NOT", "OR", "AND", "IGUAL", "MAYIG",
		"MENIG", "MAYOR", "MENOR", "MULT", "DIV", "SUM", "RES", "MOD", "PAR_IZQ",
		"PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "DOSPUNTOS", "COR_IZQ", "COR_DER",
		"COMA", "CIERRAPREGUNTA", "PUNTOCOMA", "PUNTO", "FLECHA", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "structCreation", "listStructDec", "printstmt",
		"blockelsif", "ifstmt", "whilestmt", "guardstmt", "listcases", "casestmt",
		"instdefault", "switchstmt", "forstmt", "declarationstmt", "asignationstmt",
		"function", "listParamsFunc", "callFuncionIns", "types", "typesmatriz",
		"exprFor", "expr", "conversionstmt", "exprvector", "listParams", "listArray",
		"listAccessArray", "callFuncion", "listParamsCall", "listStructExp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 65, 741, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 70, 8, 1, 11, 1, 12, 1, 71, 1,
		1, 1, 1, 1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 87, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 93, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 117, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 128, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 138, 8, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 148, 8, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 160, 8, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 167, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		173, 8, 2, 1, 2, 3, 2, 176, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 198, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 215, 8, 4, 10, 4, 12,
		4, 218, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 227, 8, 6,
		11, 6, 12, 6, 228, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 260, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 4, 10, 278, 8, 10, 11, 10, 12, 10, 279, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 310, 8, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 382, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 403, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 441, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 457,
		8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 475, 8, 18, 10, 18,
		12, 18, 478, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20,
		497, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3,
		21, 507, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 519, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 573, 8, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 5, 23, 610, 8, 23, 10, 23, 12, 23, 613, 9, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 633, 8, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 645, 8, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 656,
		8, 26, 10, 26, 12, 26, 659, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 673, 8, 27, 10, 27,
		12, 27, 676, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 690, 8, 28, 10, 28, 12, 28, 693,
		9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 3, 30, 706, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30,
		713, 8, 30, 10, 30, 12, 30, 716, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 3, 31, 725, 8, 31, 1, 31, 1, 31, 3, 31, 729, 8, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 736, 8, 31, 10, 31, 12, 31, 739,
		9, 31, 1, 31, 0, 8, 8, 36, 46, 52, 54, 56, 60, 62, 32, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 62, 0, 5, 1, 0, 48, 49, 2, 0, 46, 47, 50, 50,
		2, 0, 42, 42, 44, 44, 2, 0, 43, 43, 45, 45, 1, 0, 36, 37, 801, 0, 64, 1,
		0, 0, 0, 2, 69, 1, 0, 0, 0, 4, 175, 1, 0, 0, 0, 6, 177, 1, 0, 0, 0, 8,
		197, 1, 0, 0, 0, 10, 219, 1, 0, 0, 0, 12, 226, 1, 0, 0, 0, 14, 259, 1,
		0, 0, 0, 16, 261, 1, 0, 0, 0, 18, 268, 1, 0, 0, 0, 20, 277, 1, 0, 0, 0,
		22, 283, 1, 0, 0, 0, 24, 289, 1, 0, 0, 0, 26, 309, 1, 0, 0, 0, 28, 311,
		1, 0, 0, 0, 30, 381, 1, 0, 0, 0, 32, 402, 1, 0, 0, 0, 34, 440, 1, 0, 0,
		0, 36, 456, 1, 0, 0, 0, 38, 479, 1, 0, 0, 0, 40, 496, 1, 0, 0, 0, 42, 506,
		1, 0, 0, 0, 44, 518, 1, 0, 0, 0, 46, 572, 1, 0, 0, 0, 48, 632, 1, 0, 0,
		0, 50, 644, 1, 0, 0, 0, 52, 646, 1, 0, 0, 0, 54, 660, 1, 0, 0, 0, 56, 677,
		1, 0, 0, 0, 58, 694, 1, 0, 0, 0, 60, 705, 1, 0, 0, 0, 62, 724, 1, 0, 0,
		0, 64, 65, 3, 2, 1, 0, 65, 66, 5, 0, 0, 1, 66, 67, 6, 0, -1, 0, 67, 1,
		1, 0, 0, 0, 68, 70, 3, 4, 2, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 6,
		1, -1, 0, 74, 3, 1, 0, 0, 0, 75, 77, 3, 10, 5, 0, 76, 78, 5, 60, 0, 0,
		77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 6,
		2, -1, 0, 80, 176, 1, 0, 0, 0, 81, 82, 3, 14, 7, 0, 82, 83, 6, 2, -1, 0,
		83, 176, 1, 0, 0, 0, 84, 86, 3, 30, 15, 0, 85, 87, 5, 60, 0, 0, 86, 85,
		1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 6, 2, -1, 0,
		89, 176, 1, 0, 0, 0, 90, 92, 3, 32, 16, 0, 91, 93, 5, 60, 0, 0, 92, 91,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 6, 2, -1, 0,
		95, 176, 1, 0, 0, 0, 96, 97, 3, 16, 8, 0, 97, 98, 6, 2, -1, 0, 98, 176,
		1, 0, 0, 0, 99, 100, 3, 28, 14, 0, 100, 101, 6, 2, -1, 0, 101, 176, 1,
		0, 0, 0, 102, 103, 3, 18, 9, 0, 103, 104, 6, 2, -1, 0, 104, 176, 1, 0,
		0, 0, 105, 106, 3, 26, 13, 0, 106, 107, 6, 2, -1, 0, 107, 176, 1, 0, 0,
		0, 108, 109, 3, 34, 17, 0, 109, 110, 6, 2, -1, 0, 110, 176, 1, 0, 0, 0,
		111, 112, 3, 6, 3, 0, 112, 113, 6, 2, -1, 0, 113, 176, 1, 0, 0, 0, 114,
		116, 3, 38, 19, 0, 115, 117, 5, 60, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 6, 2, -1, 0, 119, 176, 1, 0,
		0, 0, 120, 122, 5, 19, 0, 0, 121, 123, 5, 60, 0, 0, 122, 121, 1, 0, 0,
		0, 122, 123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 176, 6, 2, -1, 0, 125,
		127, 5, 20, 0, 0, 126, 128, 5, 60, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 176, 6, 2, -1, 0, 130, 131, 5, 35,
		0, 0, 131, 132, 5, 61, 0, 0, 132, 133, 5, 21, 0, 0, 133, 134, 5, 51, 0,
		0, 134, 135, 3, 46, 23, 0, 135, 137, 5, 52, 0, 0, 136, 138, 5, 60, 0, 0,
		137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		140, 6, 2, -1, 0, 140, 176, 1, 0, 0, 0, 141, 142, 5, 35, 0, 0, 142, 143,
		5, 61, 0, 0, 143, 144, 5, 22, 0, 0, 144, 145, 5, 51, 0, 0, 145, 147, 5,
		52, 0, 0, 146, 148, 5, 60, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0,
		0, 0, 148, 149, 1, 0, 0, 0, 149, 176, 6, 2, -1, 0, 150, 151, 5, 35, 0,
		0, 151, 152, 5, 61, 0, 0, 152, 153, 5, 23, 0, 0, 153, 154, 5, 51, 0, 0,
		154, 155, 5, 24, 0, 0, 155, 156, 5, 55, 0, 0, 156, 157, 3, 46, 23, 0, 157,
		159, 5, 52, 0, 0, 158, 160, 5, 60, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 6, 2, -1, 0, 162, 176, 1, 0,
		0, 0, 163, 164, 5, 28, 0, 0, 164, 166, 3, 46, 23, 0, 165, 167, 5, 60, 0,
		0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		169, 6, 2, -1, 0, 169, 176, 1, 0, 0, 0, 170, 172, 5, 28, 0, 0, 171, 173,
		5, 60, 0, 0, 172, 171, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1, 0,
		0, 0, 174, 176, 6, 2, -1, 0, 175, 75, 1, 0, 0, 0, 175, 81, 1, 0, 0, 0,
		175, 84, 1, 0, 0, 0, 175, 90, 1, 0, 0, 0, 175, 96, 1, 0, 0, 0, 175, 99,
		1, 0, 0, 0, 175, 102, 1, 0, 0, 0, 175, 105, 1, 0, 0, 0, 175, 108, 1, 0,
		0, 0, 175, 111, 1, 0, 0, 0, 175, 114, 1, 0, 0, 0, 175, 120, 1, 0, 0, 0,
		175, 125, 1, 0, 0, 0, 175, 130, 1, 0, 0, 0, 175, 141, 1, 0, 0, 0, 175,
		150, 1, 0, 0, 0, 175, 163, 1, 0, 0, 0, 175, 170, 1, 0, 0, 0, 176, 5, 1,
		0, 0, 0, 177, 178, 5, 30, 0, 0, 178, 179, 5, 35, 0, 0, 179, 180, 5, 53,
		0, 0, 180, 181, 3, 8, 4, 0, 181, 182, 5, 54, 0, 0, 182, 183, 6, 3, -1,
		0, 183, 7, 1, 0, 0, 0, 184, 185, 6, 4, -1, 0, 185, 186, 5, 16, 0, 0, 186,
		187, 5, 35, 0, 0, 187, 188, 5, 55, 0, 0, 188, 189, 3, 40, 20, 0, 189, 190,
		6, 4, -1, 0, 190, 198, 1, 0, 0, 0, 191, 192, 5, 16, 0, 0, 192, 193, 5,
		35, 0, 0, 193, 194, 5, 55, 0, 0, 194, 195, 5, 35, 0, 0, 195, 198, 6, 4,
		-1, 0, 196, 198, 6, 4, -1, 0, 197, 184, 1, 0, 0, 0, 197, 191, 1, 0, 0,
		0, 197, 196, 1, 0, 0, 0, 198, 216, 1, 0, 0, 0, 199, 200, 10, 5, 0, 0, 200,
		201, 5, 58, 0, 0, 201, 202, 5, 16, 0, 0, 202, 203, 5, 35, 0, 0, 203, 204,
		5, 55, 0, 0, 204, 205, 3, 40, 20, 0, 205, 206, 6, 4, -1, 0, 206, 215, 1,
		0, 0, 0, 207, 208, 10, 4, 0, 0, 208, 209, 5, 58, 0, 0, 209, 210, 5, 16,
		0, 0, 210, 211, 5, 35, 0, 0, 211, 212, 5, 55, 0, 0, 212, 213, 5, 35, 0,
		0, 213, 215, 6, 4, -1, 0, 214, 199, 1, 0, 0, 0, 214, 207, 1, 0, 0, 0, 215,
		218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 9, 1,
		0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220, 5, 7, 0, 0, 220, 221, 5, 51, 0,
		0, 221, 222, 3, 52, 26, 0, 222, 223, 5, 52, 0, 0, 223, 224, 6, 5, -1, 0,
		224, 11, 1, 0, 0, 0, 225, 227, 3, 14, 7, 0, 226, 225, 1, 0, 0, 0, 227,
		228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 231, 6, 6, -1, 0, 231, 13, 1, 0, 0, 0, 232, 233, 5, 8,
		0, 0, 233, 234, 3, 46, 23, 0, 234, 235, 5, 53, 0, 0, 235, 236, 3, 2, 1,
		0, 236, 237, 5, 54, 0, 0, 237, 238, 6, 7, -1, 0, 238, 260, 1, 0, 0, 0,
		239, 240, 5, 8, 0, 0, 240, 241, 3, 46, 23, 0, 241, 242, 5, 53, 0, 0, 242,
		243, 3, 2, 1, 0, 243, 244, 5, 54, 0, 0, 244, 245, 5, 9, 0, 0, 245, 246,
		5, 53, 0, 0, 246, 247, 3, 2, 1, 0, 247, 248, 5, 54, 0, 0, 248, 249, 6,
		7, -1, 0, 249, 260, 1, 0, 0, 0, 250, 251, 5, 8, 0, 0, 251, 252, 3, 46,
		23, 0, 252, 253, 5, 53, 0, 0, 253, 254, 3, 2, 1, 0, 254, 255, 5, 54, 0,
		0, 255, 256, 5, 9, 0, 0, 256, 257, 3, 12, 6, 0, 257, 258, 6, 7, -1, 0,
		258, 260, 1, 0, 0, 0, 259, 232, 1, 0, 0, 0, 259, 239, 1, 0, 0, 0, 259,
		250, 1, 0, 0, 0, 260, 15, 1, 0, 0, 0, 261, 262, 5, 10, 0, 0, 262, 263,
		3, 46, 23, 0, 263, 264, 5, 53, 0, 0, 264, 265, 3, 2, 1, 0, 265, 266, 5,
		54, 0, 0, 266, 267, 6, 8, -1, 0, 267, 17, 1, 0, 0, 0, 268, 269, 5, 31,
		0, 0, 269, 270, 3, 46, 23, 0, 270, 271, 5, 9, 0, 0, 271, 272, 5, 53, 0,
		0, 272, 273, 3, 2, 1, 0, 273, 274, 5, 54, 0, 0, 274, 275, 6, 9, -1, 0,
		275, 19, 1, 0, 0, 0, 276, 278, 3, 22, 11, 0, 277, 276, 1, 0, 0, 0, 278,
		279, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 282, 6, 10, -1, 0, 282, 21, 1, 0, 0, 0, 283, 284, 5, 14,
		0, 0, 284, 285, 3, 46, 23, 0, 285, 286, 5, 55, 0, 0, 286, 287, 3, 2, 1,
		0, 287, 288, 6, 11, -1, 0, 288, 23, 1, 0, 0, 0, 289, 290, 5, 15, 0, 0,
		290, 291, 5, 55, 0, 0, 291, 292, 3, 2, 1, 0, 292, 293, 6, 12, -1, 0, 293,
		25, 1, 0, 0, 0, 294, 295, 5, 13, 0, 0, 295, 296, 3, 46, 23, 0, 296, 297,
		5, 53, 0, 0, 297, 298, 3, 20, 10, 0, 298, 299, 3, 24, 12, 0, 299, 300,
		5, 54, 0, 0, 300, 301, 6, 13, -1, 0, 301, 310, 1, 0, 0, 0, 302, 303, 5,
		13, 0, 0, 303, 304, 3, 46, 23, 0, 304, 305, 5, 53, 0, 0, 305, 306, 3, 20,
		10, 0, 306, 307, 5, 54, 0, 0, 307, 308, 6, 13, -1, 0, 308, 310, 1, 0, 0,
		0, 309, 294, 1, 0, 0, 0, 309, 302, 1, 0, 0, 0, 310, 27, 1, 0, 0, 0, 311,
		312, 5, 11, 0, 0, 312, 313, 5, 35, 0, 0, 313, 314, 5, 12, 0, 0, 314, 315,
		3, 44, 22, 0, 315, 316, 5, 53, 0, 0, 316, 317, 3, 2, 1, 0, 317, 318, 5,
		54, 0, 0, 318, 319, 6, 14, -1, 0, 319, 29, 1, 0, 0, 0, 320, 321, 5, 16,
		0, 0, 321, 322, 5, 35, 0, 0, 322, 323, 5, 55, 0, 0, 323, 324, 3, 40, 20,
		0, 324, 325, 5, 41, 0, 0, 325, 326, 3, 46, 23, 0, 326, 327, 6, 15, -1,
		0, 327, 382, 1, 0, 0, 0, 328, 329, 5, 16, 0, 0, 329, 330, 5, 35, 0, 0,
		330, 331, 5, 41, 0, 0, 331, 332, 3, 46, 23, 0, 332, 333, 6, 15, -1, 0,
		333, 382, 1, 0, 0, 0, 334, 335, 5, 16, 0, 0, 335, 336, 5, 35, 0, 0, 336,
		337, 5, 55, 0, 0, 337, 338, 3, 40, 20, 0, 338, 339, 5, 59, 0, 0, 339, 340,
		6, 15, -1, 0, 340, 382, 1, 0, 0, 0, 341, 342, 5, 16, 0, 0, 342, 343, 5,
		35, 0, 0, 343, 344, 5, 55, 0, 0, 344, 345, 5, 56, 0, 0, 345, 346, 3, 40,
		20, 0, 346, 347, 5, 57, 0, 0, 347, 348, 5, 41, 0, 0, 348, 349, 3, 50, 25,
		0, 349, 350, 6, 15, -1, 0, 350, 382, 1, 0, 0, 0, 351, 352, 5, 16, 0, 0,
		352, 353, 5, 35, 0, 0, 353, 354, 5, 55, 0, 0, 354, 355, 3, 42, 21, 0, 355,
		356, 5, 41, 0, 0, 356, 357, 3, 46, 23, 0, 357, 358, 6, 15, -1, 0, 358,
		382, 1, 0, 0, 0, 359, 360, 5, 17, 0, 0, 360, 361, 5, 35, 0, 0, 361, 362,
		5, 55, 0, 0, 362, 363, 3, 42, 21, 0, 363, 364, 5, 41, 0, 0, 364, 365, 3,
		46, 23, 0, 365, 366, 6, 15, -1, 0, 366, 382, 1, 0, 0, 0, 367, 368, 5, 17,
		0, 0, 368, 369, 5, 35, 0, 0, 369, 370, 5, 55, 0, 0, 370, 371, 3, 40, 20,
		0, 371, 372, 5, 41, 0, 0, 372, 373, 3, 46, 23, 0, 373, 374, 6, 15, -1,
		0, 374, 382, 1, 0, 0, 0, 375, 376, 5, 17, 0, 0, 376, 377, 5, 35, 0, 0,
		377, 378, 5, 41, 0, 0, 378, 379, 3, 46, 23, 0, 379, 380, 6, 15, -1, 0,
		380, 382, 1, 0, 0, 0, 381, 320, 1, 0, 0, 0, 381, 328, 1, 0, 0, 0, 381,
		334, 1, 0, 0, 0, 381, 341, 1, 0, 0, 0, 381, 351, 1, 0, 0, 0, 381, 359,
		1, 0, 0, 0, 381, 367, 1, 0, 0, 0, 381, 375, 1, 0, 0, 0, 382, 31, 1, 0,
		0, 0, 383, 384, 5, 35, 0, 0, 384, 385, 5, 41, 0, 0, 385, 386, 3, 46, 23,
		0, 386, 387, 6, 16, -1, 0, 387, 403, 1, 0, 0, 0, 388, 389, 5, 35, 0, 0,
		389, 390, 5, 56, 0, 0, 390, 391, 3, 46, 23, 0, 391, 392, 5, 57, 0, 0, 392,
		393, 5, 41, 0, 0, 393, 394, 3, 46, 23, 0, 394, 395, 6, 16, -1, 0, 395,
		403, 1, 0, 0, 0, 396, 397, 5, 35, 0, 0, 397, 398, 7, 0, 0, 0, 398, 399,
		5, 41, 0, 0, 399, 400, 3, 46, 23, 0, 400, 401, 6, 16, -1, 0, 401, 403,
		1, 0, 0, 0, 402, 383, 1, 0, 0, 0, 402, 388, 1, 0, 0, 0, 402, 396, 1, 0,
		0, 0, 403, 33, 1, 0, 0, 0, 404, 405, 5, 29, 0, 0, 405, 406, 5, 35, 0, 0,
		406, 407, 5, 51, 0, 0, 407, 408, 3, 36, 18, 0, 408, 409, 5, 52, 0, 0, 409,
		410, 5, 53, 0, 0, 410, 411, 3, 2, 1, 0, 411, 412, 5, 54, 0, 0, 412, 413,
		6, 17, -1, 0, 413, 441, 1, 0, 0, 0, 414, 415, 5, 29, 0, 0, 415, 416, 5,
		35, 0, 0, 416, 417, 5, 51, 0, 0, 417, 418, 3, 36, 18, 0, 418, 419, 5, 52,
		0, 0, 419, 420, 5, 62, 0, 0, 420, 421, 3, 40, 20, 0, 421, 422, 5, 53, 0,
		0, 422, 423, 3, 2, 1, 0, 423, 424, 5, 54, 0, 0, 424, 425, 6, 17, -1, 0,
		425, 441, 1, 0, 0, 0, 426, 427, 5, 29, 0, 0, 427, 428, 5, 35, 0, 0, 428,
		429, 5, 51, 0, 0, 429, 430, 3, 36, 18, 0, 430, 431, 5, 52, 0, 0, 431, 432,
		5, 62, 0, 0, 432, 433, 5, 56, 0, 0, 433, 434, 3, 40, 20, 0, 434, 435, 5,
		57, 0, 0, 435, 436, 5, 53, 0, 0, 436, 437, 3, 2, 1, 0, 437, 438, 5, 54,
		0, 0, 438, 439, 6, 17, -1, 0, 439, 441, 1, 0, 0, 0, 440, 404, 1, 0, 0,
		0, 440, 414, 1, 0, 0, 0, 440, 426, 1, 0, 0, 0, 441, 35, 1, 0, 0, 0, 442,
		443, 6, 18, -1, 0, 443, 444, 5, 35, 0, 0, 444, 445, 5, 55, 0, 0, 445, 446,
		3, 40, 20, 0, 446, 447, 6, 18, -1, 0, 447, 457, 1, 0, 0, 0, 448, 449, 5,
		35, 0, 0, 449, 450, 5, 55, 0, 0, 450, 451, 5, 56, 0, 0, 451, 452, 3, 40,
		20, 0, 452, 453, 5, 57, 0, 0, 453, 454, 6, 18, -1, 0, 454, 457, 1, 0, 0,
		0, 455, 457, 6, 18, -1, 0, 456, 442, 1, 0, 0, 0, 456, 448, 1, 0, 0, 0,
		456, 455, 1, 0, 0, 0, 457, 476, 1, 0, 0, 0, 458, 459, 10, 5, 0, 0, 459,
		460, 5, 58, 0, 0, 460, 461, 5, 35, 0, 0, 461, 462, 5, 55, 0, 0, 462, 463,
		3, 40, 20, 0, 463, 464, 6, 18, -1, 0, 464, 475, 1, 0, 0, 0, 465, 466, 10,
		4, 0, 0, 466, 467, 5, 58, 0, 0, 467, 468, 5, 35, 0, 0, 468, 469, 5, 55,
		0, 0, 469, 470, 5, 56, 0, 0, 470, 471, 3, 40, 20, 0, 471, 472, 5, 57, 0,
		0, 472, 473, 6, 18, -1, 0, 473, 475, 1, 0, 0, 0, 474, 458, 1, 0, 0, 0,
		474, 465, 1, 0, 0, 0, 475, 478, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476,
		477, 1, 0, 0, 0, 477, 37, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 479, 480, 5,
		35, 0, 0, 480, 481, 5, 51, 0, 0, 481, 482, 3, 60, 30, 0, 482, 483, 5, 52,
		0, 0, 483, 484, 6, 19, -1, 0, 484, 39, 1, 0, 0, 0, 485, 486, 5, 1, 0, 0,
		486, 497, 6, 20, -1, 0, 487, 488, 5, 2, 0, 0, 488, 497, 6, 20, -1, 0, 489,
		490, 5, 4, 0, 0, 490, 497, 6, 20, -1, 0, 491, 492, 5, 3, 0, 0, 492, 497,
		6, 20, -1, 0, 493, 494, 5, 56, 0, 0, 494, 495, 5, 57, 0, 0, 495, 497, 6,
		20, -1, 0, 496, 485, 1, 0, 0, 0, 496, 487, 1, 0, 0, 0, 496, 489, 1, 0,
		0, 0, 496, 491, 1, 0, 0, 0, 496, 493, 1, 0, 0, 0, 497, 41, 1, 0, 0, 0,
		498, 499, 5, 56, 0, 0, 499, 500, 3, 42, 21, 0, 500, 501, 5, 57, 0, 0, 501,
		502, 6, 21, -1, 0, 502, 507, 1, 0, 0, 0, 503, 504, 3, 40, 20, 0, 504, 505,
		6, 21, -1, 0, 505, 507, 1, 0, 0, 0, 506, 498, 1, 0, 0, 0, 506, 503, 1,
		0, 0, 0, 507, 43, 1, 0, 0, 0, 508, 509, 3, 46, 23, 0, 509, 510, 5, 61,
		0, 0, 510, 511, 5, 61, 0, 0, 511, 512, 5, 61, 0, 0, 512, 513, 3, 46, 23,
		0, 513, 514, 6, 22, -1, 0, 514, 519, 1, 0, 0, 0, 515, 516, 3, 46, 23, 0,
		516, 517, 6, 22, -1, 0, 517, 519, 1, 0, 0, 0, 518, 508, 1, 0, 0, 0, 518,
		515, 1, 0, 0, 0, 519, 45, 1, 0, 0, 0, 520, 521, 6, 23, -1, 0, 521, 522,
		5, 49, 0, 0, 522, 523, 3, 46, 23, 22, 523, 524, 6, 23, -1, 0, 524, 573,
		1, 0, 0, 0, 525, 526, 5, 38, 0, 0, 526, 527, 3, 46, 23, 16, 527, 528, 6,
		23, -1, 0, 528, 573, 1, 0, 0, 0, 529, 530, 5, 35, 0, 0, 530, 531, 5, 51,
		0, 0, 531, 532, 3, 62, 31, 0, 532, 533, 5, 52, 0, 0, 533, 534, 6, 23, -1,
		0, 534, 573, 1, 0, 0, 0, 535, 536, 3, 58, 29, 0, 536, 537, 6, 23, -1, 0,
		537, 573, 1, 0, 0, 0, 538, 539, 5, 51, 0, 0, 539, 540, 3, 46, 23, 0, 540,
		541, 5, 52, 0, 0, 541, 542, 6, 23, -1, 0, 542, 573, 1, 0, 0, 0, 543, 544,
		3, 48, 24, 0, 544, 545, 6, 23, -1, 0, 545, 573, 1, 0, 0, 0, 546, 547, 5,
		35, 0, 0, 547, 548, 5, 61, 0, 0, 548, 549, 5, 26, 0, 0, 549, 573, 6, 23,
		-1, 0, 550, 551, 5, 35, 0, 0, 551, 552, 5, 61, 0, 0, 552, 553, 5, 25, 0,
		0, 553, 573, 6, 23, -1, 0, 554, 555, 3, 54, 27, 0, 555, 556, 6, 23, -1,
		0, 556, 573, 1, 0, 0, 0, 557, 558, 5, 56, 0, 0, 558, 559, 3, 52, 26, 0,
		559, 560, 5, 57, 0, 0, 560, 561, 6, 23, -1, 0, 561, 573, 1, 0, 0, 0, 562,
		563, 5, 32, 0, 0, 563, 573, 6, 23, -1, 0, 564, 565, 5, 33, 0, 0, 565, 573,
		6, 23, -1, 0, 566, 567, 5, 5, 0, 0, 567, 573, 6, 23, -1, 0, 568, 569, 5,
		6, 0, 0, 569, 573, 6, 23, -1, 0, 570, 571, 5, 18, 0, 0, 571, 573, 6, 23,
		-1, 0, 572, 520, 1, 0, 0, 0, 572, 525, 1, 0, 0, 0, 572, 529, 1, 0, 0, 0,
		572, 535, 1, 0, 0, 0, 572, 538, 1, 0, 0, 0, 572, 543, 1, 0, 0, 0, 572,
		546, 1, 0, 0, 0, 572, 550, 1, 0, 0, 0, 572, 554, 1, 0, 0, 0, 572, 557,
		1, 0, 0, 0, 572, 562, 1, 0, 0, 0, 572, 564, 1, 0, 0, 0, 572, 566, 1, 0,
		0, 0, 572, 568, 1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 573, 611, 1, 0, 0, 0,
		574, 575, 10, 21, 0, 0, 575, 576, 7, 1, 0, 0, 576, 577, 3, 46, 23, 22,
		577, 578, 6, 23, -1, 0, 578, 610, 1, 0, 0, 0, 579, 580, 10, 20, 0, 0, 580,
		581, 7, 0, 0, 0, 581, 582, 3, 46, 23, 21, 582, 583, 6, 23, -1, 0, 583,
		610, 1, 0, 0, 0, 584, 585, 10, 19, 0, 0, 585, 586, 7, 2, 0, 0, 586, 587,
		3, 46, 23, 20, 587, 588, 6, 23, -1, 0, 588, 610, 1, 0, 0, 0, 589, 590,
		10, 18, 0, 0, 590, 591, 7, 3, 0, 0, 591, 592, 3, 46, 23, 19, 592, 593,
		6, 23, -1, 0, 593, 610, 1, 0, 0, 0, 594, 595, 10, 17, 0, 0, 595, 596, 7,
		4, 0, 0, 596, 597, 3, 46, 23, 18, 597, 598, 6, 23, -1, 0, 598, 610, 1,
		0, 0, 0, 599, 600, 10, 15, 0, 0, 600, 601, 5, 40, 0, 0, 601, 602, 3, 46,
		23, 16, 602, 603, 6, 23, -1, 0, 603, 610, 1, 0, 0, 0, 604, 605, 10, 14,
		0, 0, 605, 606, 5, 39, 0, 0, 606, 607, 3, 46, 23, 15, 607, 608, 6, 23,
		-1, 0, 608, 610, 1, 0, 0, 0, 609, 574, 1, 0, 0, 0, 609, 579, 1, 0, 0, 0,
		609, 584, 1, 0, 0, 0, 609, 589, 1, 0, 0, 0, 609, 594, 1, 0, 0, 0, 609,
		599, 1, 0, 0, 0, 609, 604, 1, 0, 0, 0, 610, 613, 1, 0, 0, 0, 611, 609,
		1, 0, 0, 0, 611, 612, 1, 0, 0, 0, 612, 47, 1, 0, 0, 0, 613, 611, 1, 0,
		0, 0, 614, 615, 5, 1, 0, 0, 615, 616, 5, 51, 0, 0, 616, 617, 3, 46, 23,
		0, 617, 618, 5, 52, 0, 0, 618, 619, 6, 24, -1, 0, 619, 633, 1, 0, 0, 0,
		620, 621, 5, 2, 0, 0, 621, 622, 5, 51, 0, 0, 622, 623, 3, 46, 23, 0, 623,
		624, 5, 52, 0, 0, 624, 625, 6, 24, -1, 0, 625, 633, 1, 0, 0, 0, 626, 627,
		5, 4, 0, 0, 627, 628, 5, 51, 0, 0, 628, 629, 3, 46, 23, 0, 629, 630, 5,
		52, 0, 0, 630, 631, 6, 24, -1, 0, 631, 633, 1, 0, 0, 0, 632, 614, 1, 0,
		0, 0, 632, 620, 1, 0, 0, 0, 632, 626, 1, 0, 0, 0, 633, 49, 1, 0, 0, 0,
		634, 635, 5, 56, 0, 0, 635, 636, 3, 52, 26, 0, 636, 637, 5, 57, 0, 0, 637,
		638, 6, 25, -1, 0, 638, 645, 1, 0, 0, 0, 639, 640, 5, 56, 0, 0, 640, 641,
		5, 57, 0, 0, 641, 645, 6, 25, -1, 0, 642, 643, 5, 35, 0, 0, 643, 645, 6,
		25, -1, 0, 644, 634, 1, 0, 0, 0, 644, 639, 1, 0, 0, 0, 644, 642, 1, 0,
		0, 0, 645, 51, 1, 0, 0, 0, 646, 647, 6, 26, -1, 0, 647, 648, 3, 46, 23,
		0, 648, 649, 6, 26, -1, 0, 649, 657, 1, 0, 0, 0, 650, 651, 10, 2, 0, 0,
		651, 652, 5, 58, 0, 0, 652, 653, 3, 46, 23, 0, 653, 654, 6, 26, -1, 0,
		654, 656, 1, 0, 0, 0, 655, 650, 1, 0, 0, 0, 656, 659, 1, 0, 0, 0, 657,
		655, 1, 0, 0, 0, 657, 658, 1, 0, 0, 0, 658, 53, 1, 0, 0, 0, 659, 657, 1,
		0, 0, 0, 660, 661, 6, 27, -1, 0, 661, 662, 5, 35, 0, 0, 662, 663, 6, 27,
		-1, 0, 663, 674, 1, 0, 0, 0, 664, 665, 10, 3, 0, 0, 665, 666, 3, 56, 28,
		0, 666, 667, 6, 27, -1, 0, 667, 673, 1, 0, 0, 0, 668, 669, 10, 2, 0, 0,
		669, 670, 5, 61, 0, 0, 670, 671, 5, 35, 0, 0, 671, 673, 6, 27, -1, 0, 672,
		664, 1, 0, 0, 0, 672, 668, 1, 0, 0, 0, 673, 676, 1, 0, 0, 0, 674, 672,
		1, 0, 0, 0, 674, 675, 1, 0, 0, 0, 675, 55, 1, 0, 0, 0, 676, 674, 1, 0,
		0, 0, 677, 678, 6, 28, -1, 0, 678, 679, 5, 56, 0, 0, 679, 680, 3, 46, 23,
		0, 680, 681, 5, 57, 0, 0, 681, 682, 6, 28, -1, 0, 682, 691, 1, 0, 0, 0,
		683, 684, 10, 2, 0, 0, 684, 685, 5, 56, 0, 0, 685, 686, 3, 46, 23, 0, 686,
		687, 5, 57, 0, 0, 687, 688, 6, 28, -1, 0, 688, 690, 1, 0, 0, 0, 689, 683,
		1, 0, 0, 0, 690, 693, 1, 0, 0, 0, 691, 689, 1, 0, 0, 0, 691, 692, 1, 0,
		0, 0, 692, 57, 1, 0, 0, 0, 693, 691, 1, 0, 0, 0, 694, 695, 5, 35, 0, 0,
		695, 696, 5, 51, 0, 0, 696, 697, 3, 60, 30, 0, 697, 698, 5, 52, 0, 0, 698,
		699, 6, 29, -1, 0, 699, 59, 1, 0, 0, 0, 700, 701, 6, 30, -1, 0, 701, 702,
		3, 46, 23, 0, 702, 703, 6, 30, -1, 0, 703, 706, 1, 0, 0, 0, 704, 706, 6,
		30, -1, 0, 705, 700, 1, 0, 0, 0, 705, 704, 1, 0, 0, 0, 706, 714, 1, 0,
		0, 0, 707, 708, 10, 3, 0, 0, 708, 709, 5, 58, 0, 0, 709, 710, 3, 46, 23,
		0, 710, 711, 6, 30, -1, 0, 711, 713, 1, 0, 0, 0, 712, 707, 1, 0, 0, 0,
		713, 716, 1, 0, 0, 0, 714, 712, 1, 0, 0, 0, 714, 715, 1, 0, 0, 0, 715,
		61, 1, 0, 0, 0, 716, 714, 1, 0, 0, 0, 717, 718, 6, 31, -1, 0, 718, 719,
		5, 35, 0, 0, 719, 720, 5, 55, 0, 0, 720, 721, 3, 46, 23, 0, 721, 722, 6,
		31, -1, 0, 722, 725, 1, 0, 0, 0, 723, 725, 6, 31, -1, 0, 724, 717, 1, 0,
		0, 0, 724, 723, 1, 0, 0, 0, 725, 737, 1, 0, 0, 0, 726, 728, 10, 3, 0, 0,
		727, 729, 5, 58, 0, 0, 728, 727, 1, 0, 0, 0, 728, 729, 1, 0, 0, 0, 729,
		730, 1, 0, 0, 0, 730, 731, 5, 35, 0, 0, 731, 732, 5, 55, 0, 0, 732, 733,
		3, 46, 23, 0, 733, 734, 6, 31, -1, 0, 734, 736, 1, 0, 0, 0, 735, 726, 1,
		0, 0, 0, 736, 739, 1, 0, 0, 0, 737, 735, 1, 0, 0, 0, 737, 738, 1, 0, 0,
		0, 738, 63, 1, 0, 0, 0, 739, 737, 1, 0, 0, 0, 43, 71, 77, 86, 92, 116,
		122, 127, 137, 147, 159, 166, 172, 175, 197, 214, 216, 228, 259, 279, 309,
		381, 402, 440, 456, 474, 476, 496, 506, 518, 572, 609, 611, 632, 644, 657,
		672, 674, 691, 705, 714, 724, 728, 737,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF            = antlr.TokenEOF
	SwiftGrammarParserINT            = 1
	SwiftGrammarParserFLOAT          = 2
	SwiftGrammarParserBOOL           = 3
	SwiftGrammarParserSTR            = 4
	SwiftGrammarParserTRU            = 5
	SwiftGrammarParserFAL            = 6
	SwiftGrammarParserPRINT          = 7
	SwiftGrammarParserIF             = 8
	SwiftGrammarParserELSE           = 9
	SwiftGrammarParserWHILE          = 10
	SwiftGrammarParserFOR            = 11
	SwiftGrammarParserIN             = 12
	SwiftGrammarParserSWITCH         = 13
	SwiftGrammarParserCASE           = 14
	SwiftGrammarParserDEFAULT        = 15
	SwiftGrammarParserVAR            = 16
	SwiftGrammarParserLET            = 17
	SwiftGrammarParserNIL            = 18
	SwiftGrammarParserBREAK          = 19
	SwiftGrammarParserCONTINUE       = 20
	SwiftGrammarParserAPPEND         = 21
	SwiftGrammarParserREMOVELAST     = 22
	SwiftGrammarParserREMOVE         = 23
	SwiftGrammarParserAT             = 24
	SwiftGrammarParserISEMPTY        = 25
	SwiftGrammarParserCOUNT          = 26
	SwiftGrammarParserARRAY          = 27
	SwiftGrammarParserRETURN         = 28
	SwiftGrammarParserFUNC           = 29
	SwiftGrammarParserSTRUCT         = 30
	SwiftGrammarParserGUARD          = 31
	SwiftGrammarParserNUMBER         = 32
	SwiftGrammarParserSTRING         = 33
	SwiftGrammarParserCHAR           = 34
	SwiftGrammarParserID             = 35
	SwiftGrammarParserDIFE           = 36
	SwiftGrammarParserIG_IG          = 37
	SwiftGrammarParserNOT            = 38
	SwiftGrammarParserOR             = 39
	SwiftGrammarParserAND            = 40
	SwiftGrammarParserIGUAL          = 41
	SwiftGrammarParserMAYIG          = 42
	SwiftGrammarParserMENIG          = 43
	SwiftGrammarParserMAYOR          = 44
	SwiftGrammarParserMENOR          = 45
	SwiftGrammarParserMULT           = 46
	SwiftGrammarParserDIV            = 47
	SwiftGrammarParserSUM            = 48
	SwiftGrammarParserRES            = 49
	SwiftGrammarParserMOD            = 50
	SwiftGrammarParserPAR_IZQ        = 51
	SwiftGrammarParserPAR_DER        = 52
	SwiftGrammarParserLLAVE_IZQ      = 53
	SwiftGrammarParserLLAVE_DER      = 54
	SwiftGrammarParserDOSPUNTOS      = 55
	SwiftGrammarParserCOR_IZQ        = 56
	SwiftGrammarParserCOR_DER        = 57
	SwiftGrammarParserCOMA           = 58
	SwiftGrammarParserCIERRAPREGUNTA = 59
	SwiftGrammarParserPUNTOCOMA      = 60
	SwiftGrammarParserPUNTO          = 61
	SwiftGrammarParserFLECHA         = 62
	SwiftGrammarParserWHITESPACE     = 63
	SwiftGrammarParserCOMMENT        = 64
	SwiftGrammarParserLINE_COMMENT   = 65
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s               = 0
	SwiftGrammarParserRULE_block           = 1
	SwiftGrammarParserRULE_instruction     = 2
	SwiftGrammarParserRULE_structCreation  = 3
	SwiftGrammarParserRULE_listStructDec   = 4
	SwiftGrammarParserRULE_printstmt       = 5
	SwiftGrammarParserRULE_blockelsif      = 6
	SwiftGrammarParserRULE_ifstmt          = 7
	SwiftGrammarParserRULE_whilestmt       = 8
	SwiftGrammarParserRULE_guardstmt       = 9
	SwiftGrammarParserRULE_listcases       = 10
	SwiftGrammarParserRULE_casestmt        = 11
	SwiftGrammarParserRULE_instdefault     = 12
	SwiftGrammarParserRULE_switchstmt      = 13
	SwiftGrammarParserRULE_forstmt         = 14
	SwiftGrammarParserRULE_declarationstmt = 15
	SwiftGrammarParserRULE_asignationstmt  = 16
	SwiftGrammarParserRULE_function        = 17
	SwiftGrammarParserRULE_listParamsFunc  = 18
	SwiftGrammarParserRULE_callFuncionIns  = 19
	SwiftGrammarParserRULE_types           = 20
	SwiftGrammarParserRULE_typesmatriz     = 21
	SwiftGrammarParserRULE_exprFor         = 22
	SwiftGrammarParserRULE_expr            = 23
	SwiftGrammarParserRULE_conversionstmt  = 24
	SwiftGrammarParserRULE_exprvector      = 25
	SwiftGrammarParserRULE_listParams      = 26
	SwiftGrammarParserRULE_listArray       = 27
	SwiftGrammarParserRULE_listAccessArray = 28
	SwiftGrammarParserRULE_callFuncion     = 29
	SwiftGrammarParserRULE_listParamsCall  = 30
	SwiftGrammarParserRULE_listStructExp   = 31
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(65)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&38388051328) != 0) {
		{
			p.SetState(68)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_declarationstmt returns the _declarationstmt rule contexts.
	Get_declarationstmt() IDeclarationstmtContext

	// Get_asignationstmt returns the _asignationstmt rule contexts.
	Get_asignationstmt() IAsignationstmtContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_guardstmt returns the _guardstmt rule contexts.
	Get_guardstmt() IGuardstmtContext

	// Get_switchstmt returns the _switchstmt rule contexts.
	Get_switchstmt() ISwitchstmtContext

	// Get_function returns the _function rule contexts.
	Get_function() IFunctionContext

	// Get_structCreation returns the _structCreation rule contexts.
	Get_structCreation() IStructCreationContext

	// Get_callFuncionIns returns the _callFuncionIns rule contexts.
	Get_callFuncionIns() ICallFuncionInsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_declarationstmt sets the _declarationstmt rule contexts.
	Set_declarationstmt(IDeclarationstmtContext)

	// Set_asignationstmt sets the _asignationstmt rule contexts.
	Set_asignationstmt(IAsignationstmtContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_guardstmt sets the _guardstmt rule contexts.
	Set_guardstmt(IGuardstmtContext)

	// Set_switchstmt sets the _switchstmt rule contexts.
	Set_switchstmt(ISwitchstmtContext)

	// Set_function sets the _function rule contexts.
	Set_function(IFunctionContext)

	// Set_structCreation sets the _structCreation rule contexts.
	Set_structCreation(IStructCreationContext)

	// Set_callFuncionIns sets the _callFuncionIns rule contexts.
	Set_callFuncionIns(ICallFuncionInsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PUNTOCOMA() antlr.TerminalNode
	Ifstmt() IIfstmtContext
	Declarationstmt() IDeclarationstmtContext
	Asignationstmt() IAsignationstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Switchstmt() ISwitchstmtContext
	Function() IFunctionContext
	StructCreation() IStructCreationContext
	CallFuncionIns() ICallFuncionInsContext
	BREAK() antlr.TerminalNode
	CONTINUE() antlr.TerminalNode
	ID() antlr.TerminalNode
	PUNTO() antlr.TerminalNode
	APPEND() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	Expr() IExprContext
	PAR_DER() antlr.TerminalNode
	REMOVELAST() antlr.TerminalNode
	REMOVE() antlr.TerminalNode
	AT() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	RETURN() antlr.TerminalNode

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	inst             interfaces.Instruction
	_printstmt       IPrintstmtContext
	_ifstmt          IIfstmtContext
	_declarationstmt IDeclarationstmtContext
	_asignationstmt  IAsignationstmtContext
	_whilestmt       IWhilestmtContext
	_forstmt         IForstmtContext
	_guardstmt       IGuardstmtContext
	_switchstmt      ISwitchstmtContext
	_function        IFunctionContext
	_structCreation  IStructCreationContext
	_callFuncionIns  ICallFuncionInsContext
	_BREAK           antlr.Token
	_CONTINUE        antlr.Token
	_ID              antlr.Token
	_expr            IExprContext
	_RETURN          antlr.Token
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *InstructionContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *InstructionContext) Get_ID() antlr.Token { return s._ID }

func (s *InstructionContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *InstructionContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *InstructionContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *InstructionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *InstructionContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_declarationstmt() IDeclarationstmtContext { return s._declarationstmt }

func (s *InstructionContext) Get_asignationstmt() IAsignationstmtContext { return s._asignationstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_switchstmt() ISwitchstmtContext { return s._switchstmt }

func (s *InstructionContext) Get_function() IFunctionContext { return s._function }

func (s *InstructionContext) Get_structCreation() IStructCreationContext { return s._structCreation }

func (s *InstructionContext) Get_callFuncionIns() ICallFuncionInsContext { return s._callFuncionIns }

func (s *InstructionContext) Get_expr() IExprContext { return s._expr }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_declarationstmt(v IDeclarationstmtContext) { s._declarationstmt = v }

func (s *InstructionContext) Set_asignationstmt(v IAsignationstmtContext) { s._asignationstmt = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_guardstmt(v IGuardstmtContext) { s._guardstmt = v }

func (s *InstructionContext) Set_switchstmt(v ISwitchstmtContext) { s._switchstmt = v }

func (s *InstructionContext) Set_function(v IFunctionContext) { s._function = v }

func (s *InstructionContext) Set_structCreation(v IStructCreationContext) { s._structCreation = v }

func (s *InstructionContext) Set_callFuncionIns(v ICallFuncionInsContext) { s._callFuncionIns = v }

func (s *InstructionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTOCOMA, 0)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Declarationstmt() IDeclarationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationstmtContext)
}

func (s *InstructionContext) Asignationstmt() IAsignationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignationstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *InstructionContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *InstructionContext) StructCreation() IStructCreationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCreationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCreationContext)
}

func (s *InstructionContext) CallFuncionIns() ICallFuncionInsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFuncionInsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFuncionInsContext)
}

func (s *InstructionContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *InstructionContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *InstructionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *InstructionContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *InstructionContext) APPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAPPEND, 0)
}

func (s *InstructionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *InstructionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InstructionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *InstructionContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVELAST, 0)
}

func (s *InstructionContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserREMOVE, 0)
}

func (s *InstructionContext) AT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAT, 0)
}

func (s *InstructionContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *InstructionContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRETURN, 0)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	var _la int

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(76)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(85)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)

			var _x = p.Asignationstmt()

			localctx.(*InstructionContext)._asignationstmt = _x
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(91)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignationstmt().GetAsig()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhileinst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetForinst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGd()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(105)

			var _x = p.Switchstmt()

			localctx.(*InstructionContext)._switchstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstmt().GetSw()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(108)

			var _x = p.Function()

			localctx.(*InstructionContext)._function = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_function().GetFun()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(111)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(114)

			var _x = p.CallFuncionIns()

			localctx.(*InstructionContext)._callFuncionIns = _x
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(115)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_callFuncionIns().GetCf()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(120)

			var _m = p.Match(SwiftGrammarParserBREAK)

			localctx.(*InstructionContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(121)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewBreak((func() int {
			if localctx.(*InstructionContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_BREAK().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_BREAK().GetColumn()
			}
		}()))

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(125)

			var _m = p.Match(SwiftGrammarParserCONTINUE)

			localctx.(*InstructionContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(126)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewContinue((func() int {
			if localctx.(*InstructionContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_CONTINUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_CONTINUE() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_CONTINUE().GetColumn()
			}
		}()))

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(130)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*InstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(SwiftGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		{
			p.SetState(135)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(136)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewAppend((func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*InstructionContext).Get_ID().GetText()
			}
		}()), localctx.(*InstructionContext).Get_expr().GetE())

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(141)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*InstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(SwiftGrammarParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(146)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewRemoveLast((func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*InstructionContext).Get_ID().GetText()
			}
		}()))

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(150)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*InstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(SwiftGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(SwiftGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		{
			p.SetState(157)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(158)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewRemoveAt((func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*InstructionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*InstructionContext).Get_ID().GetText()
			}
		}()), localctx.(*InstructionContext).Get_expr().GetE())

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(163)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*InstructionContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(165)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewReturn((func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetColumn()
			}
		}()), localctx.(*InstructionContext).Get_expr().GetE())

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(170)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*InstructionContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(171)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = instructions.NewReturn((func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetLine()
			}
		}()), (func() int {
			if localctx.(*InstructionContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*InstructionContext).Get_RETURN().GetColumn()
			}
		}()), nil)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructCreationContext is an interface to support dynamic dispatch.
type IStructCreationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRUCT returns the _STRUCT token.
	Get_STRUCT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_STRUCT sets the _STRUCT token.
	Set_STRUCT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listStructDec returns the _listStructDec rule contexts.
	Get_listStructDec() IListStructDecContext

	// Set_listStructDec sets the _listStructDec rule contexts.
	Set_listStructDec(IListStructDecContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVE_IZQ() antlr.TerminalNode
	ListStructDec() IListStructDecContext
	LLAVE_DER() antlr.TerminalNode

	// IsStructCreationContext differentiates from other interfaces.
	IsStructCreationContext()
}

type StructCreationContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	dec            interfaces.Instruction
	_STRUCT        antlr.Token
	_ID            antlr.Token
	_listStructDec IListStructDecContext
}

func NewEmptyStructCreationContext() *StructCreationContext {
	var p = new(StructCreationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
	return p
}

func InitEmptyStructCreationContext(p *StructCreationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
}

func (*StructCreationContext) IsStructCreationContext() {}

func NewStructCreationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCreationContext {
	var p = new(StructCreationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structCreation

	return p
}

func (s *StructCreationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCreationContext) Get_STRUCT() antlr.Token { return s._STRUCT }

func (s *StructCreationContext) Get_ID() antlr.Token { return s._ID }

func (s *StructCreationContext) Set_STRUCT(v antlr.Token) { s._STRUCT = v }

func (s *StructCreationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructCreationContext) Get_listStructDec() IListStructDecContext { return s._listStructDec }

func (s *StructCreationContext) Set_listStructDec(v IListStructDecContext) { s._listStructDec = v }

func (s *StructCreationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *StructCreationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *StructCreationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *StructCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructCreationContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *StructCreationContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *StructCreationContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (p *SwiftGrammarParser) StructCreation() (localctx IStructCreationContext) {
	localctx = NewStructCreationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_structCreation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)

		var _m = p.Match(SwiftGrammarParserSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructCreationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(181)
		p.Match(SwiftGrammarParserLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructCreationContext).dec = instructions.NewStruct((func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructCreationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructCreationContext).Get_ID().GetText()
		}
	}()), localctx.(*StructCreationContext).Get_listStructDec().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructDecContext is an interface to support dynamic dispatch.
type IListStructDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetIdp returns the idp token.
	GetIdp() antlr.Token

	// GetIds returns the ids token.
	GetIds() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetIdp sets the idp token.
	SetIdp(antlr.Token)

	// SetIds sets the ids token.
	SetIds(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructDecContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListStructDecContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	VAR() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Types() ITypesContext
	COMA() antlr.TerminalNode
	ListStructDec() IListStructDecContext

	// IsListStructDecContext differentiates from other interfaces.
	IsListStructDecContext()
}

type ListStructDecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructDecContext
	_ID    antlr.Token
	_types ITypesContext
	idp    antlr.Token
	ids    antlr.Token
}

func NewEmptyListStructDecContext() *ListStructDecContext {
	var p = new(ListStructDecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
	return p
}

func InitEmptyListStructDecContext(p *ListStructDecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
}

func (*ListStructDecContext) IsListStructDecContext() {}

func NewListStructDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructDecContext {
	var p = new(ListStructDecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec

	return p
}

func (s *ListStructDecContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructDecContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructDecContext) GetIdp() antlr.Token { return s.idp }

func (s *ListStructDecContext) GetIds() antlr.Token { return s.ids }

func (s *ListStructDecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructDecContext) SetIdp(v antlr.Token) { s.idp = v }

func (s *ListStructDecContext) SetIds(v antlr.Token) { s.ids = v }

func (s *ListStructDecContext) GetList() IListStructDecContext { return s.list }

func (s *ListStructDecContext) Get_types() ITypesContext { return s._types }

func (s *ListStructDecContext) SetList(v IListStructDecContext) { s.list = v }

func (s *ListStructDecContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListStructDecContext) GetL() []interface{} { return s.l }

func (s *ListStructDecContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *ListStructDecContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ListStructDecContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ListStructDecContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *ListStructDecContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListStructDecContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructDecContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *ListStructDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructDec(s)
	}
}

func (s *ListStructDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructDec(s)
	}
}

func (p *SwiftGrammarParser) ListStructDec() (localctx IListStructDecContext) {
	return p.listStructDec(0)
}

func (p *SwiftGrammarParser) listStructDec(_p int) (localctx IListStructDecContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructDecContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructDecContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, SwiftGrammarParserRULE_listStructDec, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(185)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Types()

			localctx.(*ListStructDecContext)._types = _x
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructDecContext).Get_types().GetTy(), "")
		arr = append(arr, newParams)
		localctx.(*ListStructDecContext).l = arr

	case 2:
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext).idp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext).ids = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).GetIdp() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).GetIdp().GetText()
			}
		}()), environment.DEPENDIENTE, (func() string {
			if localctx.(*ListStructDecContext).GetIds() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).GetIds().GetText()
			}
		}()))
		arr = append(arr, newParams)
		localctx.(*ListStructDecContext).l = arr

	case 3:
		localctx.(*ListStructDecContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(200)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(201)
					p.Match(SwiftGrammarParserVAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(202)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(203)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(204)

					var _x = p.Types()

					localctx.(*ListStructDecContext)._types = _x
				}

				var arr []interface{}
				newParams := environment.NewStructType((func() string {
					if localctx.(*ListStructDecContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).Get_ID().GetText()
					}
				}()), localctx.(*ListStructDecContext).Get_types().GetTy(), "")
				arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
				localctx.(*ListStructDecContext).l = arr

			case 2:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(208)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(209)
					p.Match(SwiftGrammarParserVAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(210)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).idp = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(211)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(212)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).ids = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				var arr []interface{}
				newParams := environment.NewStructType((func() string {
					if localctx.(*ListStructDecContext).GetIdp() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).GetIdp().GetText()
					}
				}()), environment.DEPENDIENTE, (func() string {
					if localctx.(*ListStructDecContext).GetIds() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).GetIds().GetText()
					}
				}()))
				arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
				localctx.(*ListStructDecContext).l = arr

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	PAR_DER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	prnt        interfaces.Instruction
	_PRINT      antlr.Token
	_listParams IListParamsContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *PrintstmtContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *PrintstmtContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *PrintstmtContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)

		var _x = p.listParams(0)

		localctx.(*PrintstmtContext)._listParams = _x
	}
	{
		p.SetState(222)
		p.Match(SwiftGrammarParserPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_listParams().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockelsifContext is an interface to support dynamic dispatch.
type IBlockelsifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// GetElseif returns the elseif rule context list.
	GetElseif() []IIfstmtContext

	// SetElseif sets the elseif rule context list.
	SetElseif([]IIfstmtContext)

	// GetBlkif returns the blkif attribute.
	GetBlkif() []interface{}

	// SetBlkif sets the blkif attribute.
	SetBlkif([]interface{})

	// Getter signatures
	AllIfstmt() []IIfstmtContext
	Ifstmt(i int) IIfstmtContext

	// IsBlockelsifContext differentiates from other interfaces.
	IsBlockelsifContext()
}

type BlockelsifContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	blkif   []interface{}
	_ifstmt IIfstmtContext
	elseif  []IIfstmtContext
}

func NewEmptyBlockelsifContext() *BlockelsifContext {
	var p = new(BlockelsifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockelsif
	return p
}

func InitEmptyBlockelsifContext(p *BlockelsifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_blockelsif
}

func (*BlockelsifContext) IsBlockelsifContext() {}

func NewBlockelsifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockelsifContext {
	var p = new(BlockelsifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_blockelsif

	return p
}

func (s *BlockelsifContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockelsifContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *BlockelsifContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *BlockelsifContext) GetElseif() []IIfstmtContext { return s.elseif }

func (s *BlockelsifContext) SetElseif(v []IIfstmtContext) { s.elseif = v }

func (s *BlockelsifContext) GetBlkif() []interface{} { return s.blkif }

func (s *BlockelsifContext) SetBlkif(v []interface{}) { s.blkif = v }

func (s *BlockelsifContext) AllIfstmt() []IIfstmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfstmtContext); ok {
			len++
		}
	}

	tst := make([]IIfstmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfstmtContext); ok {
			tst[i] = t.(IIfstmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockelsifContext) Ifstmt(i int) IIfstmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *BlockelsifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockelsifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockelsifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlockelsif(s)
	}
}

func (s *BlockelsifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlockelsif(s)
	}
}

func (p *SwiftGrammarParser) Blockelsif() (localctx IBlockelsifContext) {
	localctx = NewBlockelsifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_blockelsif)

	localctx.(*BlockelsifContext).blkif = []interface{}{}
	var listIfs []IIfstmtContext

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(225)

				var _x = p.Ifstmt()

				localctx.(*BlockelsifContext)._ifstmt = _x
			}
			localctx.(*BlockelsifContext).elseif = append(localctx.(*BlockelsifContext).elseif, localctx.(*BlockelsifContext)._ifstmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	listIfs = localctx.(*BlockelsifContext).GetElseif()
	for _, e := range listIfs {
		localctx.(*BlockelsifContext).blkif = append(localctx.(*BlockelsifContext).blkif, e.GetIfinst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetIfblck returns the ifblck rule contexts.
	GetIfblck() IBlockContext

	// GetElseblck returns the elseblck rule contexts.
	GetElseblck() IBlockContext

	// Get_blockelsif returns the _blockelsif rule contexts.
	Get_blockelsif() IBlockelsifContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetIfblck sets the ifblck rule contexts.
	SetIfblck(IBlockContext)

	// SetElseblck sets the elseblck rule contexts.
	SetElseblck(IBlockContext)

	// Set_blockelsif sets the _blockelsif rule contexts.
	Set_blockelsif(IBlockelsifContext)

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVE_IZQ() []antlr.TerminalNode
	LLAVE_IZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVE_DER() []antlr.TerminalNode
	LLAVE_DER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Blockelsif() IBlockelsifContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	ifinst      interfaces.Instruction
	_IF         antlr.Token
	_expr       IExprContext
	_block      IBlockContext
	ifblck      IBlockContext
	elseblck    IBlockContext
	_blockelsif IBlockelsifContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) Get_block() IBlockContext { return s._block }

func (s *IfstmtContext) GetIfblck() IBlockContext { return s.ifblck }

func (s *IfstmtContext) GetElseblck() IBlockContext { return s.elseblck }

func (s *IfstmtContext) Get_blockelsif() IBlockelsifContext { return s._blockelsif }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) SetIfblck(v IBlockContext) { s.ifblck = v }

func (s *IfstmtContext) SetElseblck(v IBlockContext) { s.elseblck = v }

func (s *IfstmtContext) Set_blockelsif(v IBlockelsifContext) { s._blockelsif = v }

func (s *IfstmtContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *IfstmtContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVE_IZQ)
}

func (s *IfstmtContext) LLAVE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfstmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) AllLLAVE_DER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVE_DER)
}

func (s *IfstmtContext) LLAVE_DER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, i)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) Blockelsif() IBlockelsifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockelsifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockelsifContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_ifstmt)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(234)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(236)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(241)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifblck = _x
		}
		{
			p.SetState(243)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)

			var _x = p.Block()

			localctx.(*IfstmtContext).elseblck = _x
		}
		{
			p.SetState(247)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetIfblck().GetBlk(), localctx.(*IfstmtContext).GetElseblck().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(252)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifblck = _x
		}
		{
			p.SetState(254)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)

			var _x = p.Blockelsif()

			localctx.(*IfstmtContext)._blockelsif = _x
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetIfblck().GetBlk(), localctx.(*IfstmtContext).Get_blockelsif().GetBlkif())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhileinst returns the whileinst attribute.
	GetWhileinst() interfaces.Instruction

	// SetWhileinst sets the whileinst attribute.
	SetWhileinst(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVE_IZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVE_DER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	whileinst interfaces.Instruction
	_WHILE    antlr.Token
	_expr     IExprContext
	_block    IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhileinst() interfaces.Instruction { return s.whileinst }

func (s *WhilestmtContext) SetWhileinst(v interfaces.Instruction) { s.whileinst = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(263)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(265)
		p.Match(SwiftGrammarParserLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).whileinst = instructions.NewWhile((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGd returns the gd attribute.
	GetGd() interfaces.Instruction

	// SetGd sets the gd attribute.
	SetGd(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVE_IZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVE_DER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	gd     interfaces.Instruction
	_GUARD antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardstmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardstmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardstmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardstmtContext) GetGd() interfaces.Instruction { return s.gd }

func (s *GuardstmtContext) SetGd(v interfaces.Instruction) { s.gd = v }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardstmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(270)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(273)
		p.Match(SwiftGrammarParserLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardstmtContext).gd = instructions.NewGuard((func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardstmtContext).Get_expr().GetE(), localctx.(*GuardstmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListcasesContext is an interface to support dynamic dispatch.
type IListcasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_casestmt returns the _casestmt rule contexts.
	Get_casestmt() ICasestmtContext

	// Set_casestmt sets the _casestmt rule contexts.
	Set_casestmt(ICasestmtContext)

	// GetCases returns the cases rule context list.
	GetCases() []ICasestmtContext

	// SetCases sets the cases rule context list.
	SetCases([]ICasestmtContext)

	// GetLcas returns the lcas attribute.
	GetLcas() []interface{}

	// SetLcas sets the lcas attribute.
	SetLcas([]interface{})

	// Getter signatures
	AllCasestmt() []ICasestmtContext
	Casestmt(i int) ICasestmtContext

	// IsListcasesContext differentiates from other interfaces.
	IsListcasesContext()
}

type ListcasesContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	lcas      []interface{}
	_casestmt ICasestmtContext
	cases     []ICasestmtContext
}

func NewEmptyListcasesContext() *ListcasesContext {
	var p = new(ListcasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listcases
	return p
}

func InitEmptyListcasesContext(p *ListcasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listcases
}

func (*ListcasesContext) IsListcasesContext() {}

func NewListcasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListcasesContext {
	var p = new(ListcasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listcases

	return p
}

func (s *ListcasesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListcasesContext) Get_casestmt() ICasestmtContext { return s._casestmt }

func (s *ListcasesContext) Set_casestmt(v ICasestmtContext) { s._casestmt = v }

func (s *ListcasesContext) GetCases() []ICasestmtContext { return s.cases }

func (s *ListcasesContext) SetCases(v []ICasestmtContext) { s.cases = v }

func (s *ListcasesContext) GetLcas() []interface{} { return s.lcas }

func (s *ListcasesContext) SetLcas(v []interface{}) { s.lcas = v }

func (s *ListcasesContext) AllCasestmt() []ICasestmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICasestmtContext); ok {
			len++
		}
	}

	tst := make([]ICasestmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICasestmtContext); ok {
			tst[i] = t.(ICasestmtContext)
			i++
		}
	}

	return tst
}

func (s *ListcasesContext) Casestmt(i int) ICasestmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasestmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasestmtContext)
}

func (s *ListcasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListcasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListcasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListcases(s)
	}
}

func (s *ListcasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListcases(s)
	}
}

func (p *SwiftGrammarParser) Listcases() (localctx IListcasesContext) {
	localctx = NewListcasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_listcases)

	localctx.(*ListcasesContext).lcas = []interface{}{}
	var listCase []ICasestmtContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SwiftGrammarParserCASE {
		{
			p.SetState(276)

			var _x = p.Casestmt()

			localctx.(*ListcasesContext)._casestmt = _x
		}
		localctx.(*ListcasesContext).cases = append(localctx.(*ListcasesContext).cases, localctx.(*ListcasesContext)._casestmt)

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listCase = localctx.(*ListcasesContext).GetCases()
	for _, e := range listCase {
		localctx.(*ListcasesContext).lcas = append(localctx.(*ListcasesContext).lcas, e.GetCaseinst())
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasestmtContext is an interface to support dynamic dispatch.
type ICasestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCaseinst returns the caseinst attribute.
	GetCaseinst() interfaces.Instruction

	// SetCaseinst sets the caseinst attribute.
	SetCaseinst(interfaces.Instruction)

	// Getter signatures
	CASE() antlr.TerminalNode
	Expr() IExprContext
	DOSPUNTOS() antlr.TerminalNode
	Block() IBlockContext

	// IsCasestmtContext differentiates from other interfaces.
	IsCasestmtContext()
}

type CasestmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	caseinst interfaces.Instruction
	_CASE    antlr.Token
	_expr    IExprContext
	_block   IBlockContext
}

func NewEmptyCasestmtContext() *CasestmtContext {
	var p = new(CasestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_casestmt
	return p
}

func InitEmptyCasestmtContext(p *CasestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_casestmt
}

func (*CasestmtContext) IsCasestmtContext() {}

func NewCasestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasestmtContext {
	var p = new(CasestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_casestmt

	return p
}

func (s *CasestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *CasestmtContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasestmtContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasestmtContext) Get_expr() IExprContext { return s._expr }

func (s *CasestmtContext) Get_block() IBlockContext { return s._block }

func (s *CasestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *CasestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *CasestmtContext) GetCaseinst() interfaces.Instruction { return s.caseinst }

func (s *CasestmtContext) SetCaseinst(v interfaces.Instruction) { s.caseinst = v }

func (s *CasestmtContext) CASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCASE, 0)
}

func (s *CasestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CasestmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *CasestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CasestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCasestmt(s)
	}
}

func (s *CasestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCasestmt(s)
	}
}

func (p *SwiftGrammarParser) Casestmt() (localctx ICasestmtContext) {
	localctx = NewCasestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_casestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)

		var _m = p.Match(SwiftGrammarParserCASE)

		localctx.(*CasestmtContext)._CASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)

		var _x = p.expr(0)

		localctx.(*CasestmtContext)._expr = _x
	}
	{
		p.SetState(285)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)

		var _x = p.Block()

		localctx.(*CasestmtContext)._block = _x
	}
	localctx.(*CasestmtContext).caseinst = instructions.NewCase((func() int {
		if localctx.(*CasestmtContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasestmtContext).Get_CASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CasestmtContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasestmtContext).Get_CASE().GetColumn()
		}
	}()), localctx.(*CasestmtContext).Get_expr().GetE(), localctx.(*CasestmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstdefaultContext is an interface to support dynamic dispatch.
type IInstdefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_DEFAULT returns the _DEFAULT token.
	Get_DEFAULT() antlr.Token

	// Set_DEFAULT sets the _DEFAULT token.
	Set_DEFAULT(antlr.Token)

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetInstdef returns the instdef attribute.
	GetInstdef() interfaces.Instruction

	// SetInstdef sets the instdef attribute.
	SetInstdef(interfaces.Instruction)

	// Getter signatures
	DEFAULT() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Block() IBlockContext

	// IsInstdefaultContext differentiates from other interfaces.
	IsInstdefaultContext()
}

type InstdefaultContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	instdef  interfaces.Instruction
	_DEFAULT antlr.Token
	_block   IBlockContext
}

func NewEmptyInstdefaultContext() *InstdefaultContext {
	var p = new(InstdefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instdefault
	return p
}

func InitEmptyInstdefaultContext(p *InstdefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instdefault
}

func (*InstdefaultContext) IsInstdefaultContext() {}

func NewInstdefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstdefaultContext {
	var p = new(InstdefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instdefault

	return p
}

func (s *InstdefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *InstdefaultContext) Get_DEFAULT() antlr.Token { return s._DEFAULT }

func (s *InstdefaultContext) Set_DEFAULT(v antlr.Token) { s._DEFAULT = v }

func (s *InstdefaultContext) Get_block() IBlockContext { return s._block }

func (s *InstdefaultContext) Set_block(v IBlockContext) { s._block = v }

func (s *InstdefaultContext) GetInstdef() interfaces.Instruction { return s.instdef }

func (s *InstdefaultContext) SetInstdef(v interfaces.Instruction) { s.instdef = v }

func (s *InstdefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDEFAULT, 0)
}

func (s *InstdefaultContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *InstdefaultContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *InstdefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstdefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstdefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstdefault(s)
	}
}

func (s *InstdefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstdefault(s)
	}
}

func (p *SwiftGrammarParser) Instdefault() (localctx IInstdefaultContext) {
	localctx = NewInstdefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_instdefault)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)

		var _m = p.Match(SwiftGrammarParserDEFAULT)

		localctx.(*InstdefaultContext)._DEFAULT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(SwiftGrammarParserDOSPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)

		var _x = p.Block()

		localctx.(*InstdefaultContext)._block = _x
	}
	localctx.(*InstdefaultContext).instdef = instructions.NewDefault((func() int {
		if localctx.(*InstdefaultContext).Get_DEFAULT() == nil {
			return 0
		} else {
			return localctx.(*InstdefaultContext).Get_DEFAULT().GetLine()
		}
	}()), (func() int {
		if localctx.(*InstdefaultContext).Get_DEFAULT() == nil {
			return 0
		} else {
			return localctx.(*InstdefaultContext).Get_DEFAULT().GetColumn()
		}
	}()), localctx.(*InstdefaultContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SWITCH returns the _SWITCH token.
	Get_SWITCH() antlr.Token

	// Set_SWITCH sets the _SWITCH token.
	Set_SWITCH(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_listcases returns the _listcases rule contexts.
	Get_listcases() IListcasesContext

	// Get_instdefault returns the _instdefault rule contexts.
	Get_instdefault() IInstdefaultContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_listcases sets the _listcases rule contexts.
	Set_listcases(IListcasesContext)

	// Set_instdefault sets the _instdefault rule contexts.
	Set_instdefault(IInstdefaultContext)

	// GetSw returns the sw attribute.
	GetSw() interfaces.Instruction

	// SetSw sets the sw attribute.
	SetSw(interfaces.Instruction)

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVE_IZQ() antlr.TerminalNode
	Listcases() IListcasesContext
	Instdefault() IInstdefaultContext
	LLAVE_DER() antlr.TerminalNode

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	sw           interfaces.Instruction
	_SWITCH      antlr.Token
	_expr        IExprContext
	_listcases   IListcasesContext
	_instdefault IInstdefaultContext
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) Get_SWITCH() antlr.Token { return s._SWITCH }

func (s *SwitchstmtContext) Set_SWITCH(v antlr.Token) { s._SWITCH = v }

func (s *SwitchstmtContext) Get_expr() IExprContext { return s._expr }

func (s *SwitchstmtContext) Get_listcases() IListcasesContext { return s._listcases }

func (s *SwitchstmtContext) Get_instdefault() IInstdefaultContext { return s._instdefault }

func (s *SwitchstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *SwitchstmtContext) Set_listcases(v IListcasesContext) { s._listcases = v }

func (s *SwitchstmtContext) Set_instdefault(v IInstdefaultContext) { s._instdefault = v }

func (s *SwitchstmtContext) GetSw() interfaces.Instruction { return s.sw }

func (s *SwitchstmtContext) SetSw(v interfaces.Instruction) { s.sw = v }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *SwitchstmtContext) Listcases() IListcasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListcasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListcasesContext)
}

func (s *SwitchstmtContext) Instdefault() IInstdefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstdefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstdefaultContext)
}

func (s *SwitchstmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstmt(s)
	}
}

func (s *SwitchstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstmt(s)
	}
}

func (p *SwiftGrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_switchstmt)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)

			var _m = p.Match(SwiftGrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(296)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _x = p.Listcases()

			localctx.(*SwitchstmtContext)._listcases = _x
		}
		{
			p.SetState(298)

			var _x = p.Instdefault()

			localctx.(*SwitchstmtContext)._instdefault = _x
		}
		{
			p.SetState(299)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).sw = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), localctx.(*SwitchstmtContext).Get_listcases().GetLcas(), localctx.(*SwitchstmtContext).Get_instdefault().GetInstdef())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)

			var _m = p.Match(SwiftGrammarParserSWITCH)

			localctx.(*SwitchstmtContext)._SWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _x = p.expr(0)

			localctx.(*SwitchstmtContext)._expr = _x
		}
		{
			p.SetState(304)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)

			var _x = p.Listcases()

			localctx.(*SwitchstmtContext)._listcases = _x
		}
		{
			p.SetState(306)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstmtContext).sw = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstmtContext).Get_SWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstmtContext).Get_SWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstmtContext).Get_expr().GetE(), localctx.(*SwitchstmtContext).Get_listcases().GetLcas(), nil)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_exprFor returns the _exprFor rule contexts.
	Get_exprFor() IExprForContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_exprFor sets the _exprFor rule contexts.
	Set_exprFor(IExprForContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetForinst returns the forinst attribute.
	GetForinst() interfaces.Instruction

	// SetForinst sets the forinst attribute.
	SetForinst(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	ExprFor() IExprForContext
	LLAVE_IZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVE_DER() antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	forinst  interfaces.Instruction
	_FOR     antlr.Token
	_ID      antlr.Token
	_exprFor IExprForContext
	_block   IBlockContext
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Get_exprFor() IExprForContext { return s._exprFor }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) Set_exprFor(v IExprForContext) { s._exprFor = v }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) GetForinst() interfaces.Instruction { return s.forinst }

func (s *ForstmtContext) SetForinst(v interfaces.Instruction) { s.forinst = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) ExprFor() IExprForContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprForContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprForContext)
}

func (s *ForstmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForstmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)

		var _m = p.Match(SwiftGrammarParserFOR)

		localctx.(*ForstmtContext)._FOR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ForstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(SwiftGrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)

		var _x = p.ExprFor()

		localctx.(*ForstmtContext)._exprFor = _x
	}
	{
		p.SetState(315)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)

		var _x = p.Block()

		localctx.(*ForstmtContext)._block = _x
	}
	{
		p.SetState(317)
		p.Match(SwiftGrammarParserLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ForstmtContext).forinst = instructions.NewFor((func() int {
		if localctx.(*ForstmtContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*ForstmtContext).Get_FOR().GetLine()
		}
	}()), (func() int {
		if localctx.(*ForstmtContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*ForstmtContext).Get_FOR().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ForstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ForstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*ForstmtContext).Get_exprFor().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclarationstmtContext is an interface to support dynamic dispatch.
type IDeclarationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_exprvector returns the _exprvector rule contexts.
	Get_exprvector() IExprvectorContext

	// Get_typesmatriz returns the _typesmatriz rule contexts.
	Get_typesmatriz() ITypesmatrizContext

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_exprvector sets the _exprvector rule contexts.
	Set_exprvector(IExprvectorContext)

	// Set_typesmatriz sets the _typesmatriz rule contexts.
	Set_typesmatriz(ITypesmatrizContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Types() ITypesContext
	IGUAL() antlr.TerminalNode
	Expr() IExprContext
	CIERRAPREGUNTA() antlr.TerminalNode
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode
	Exprvector() IExprvectorContext
	Typesmatriz() ITypesmatrizContext
	LET() antlr.TerminalNode

	// IsDeclarationstmtContext differentiates from other interfaces.
	IsDeclarationstmtContext()
}

type DeclarationstmtContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	dec          interfaces.Instruction
	_VAR         antlr.Token
	_ID          antlr.Token
	_types       ITypesContext
	_expr        IExprContext
	_exprvector  IExprvectorContext
	_typesmatriz ITypesmatrizContext
	_LET         antlr.Token
}

func NewEmptyDeclarationstmtContext() *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
	return p
}

func InitEmptyDeclarationstmtContext(p *DeclarationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
}

func (*DeclarationstmtContext) IsDeclarationstmtContext() {}

func NewDeclarationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt

	return p
}

func (s *DeclarationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclarationstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarationstmtContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclarationstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarationstmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclarationstmtContext) Get_types() ITypesContext { return s._types }

func (s *DeclarationstmtContext) Get_expr() IExprContext { return s._expr }

func (s *DeclarationstmtContext) Get_exprvector() IExprvectorContext { return s._exprvector }

func (s *DeclarationstmtContext) Get_typesmatriz() ITypesmatrizContext { return s._typesmatriz }

func (s *DeclarationstmtContext) Set_types(v ITypesContext) { s._types = v }

func (s *DeclarationstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclarationstmtContext) Set_exprvector(v IExprvectorContext) { s._exprvector = v }

func (s *DeclarationstmtContext) Set_typesmatriz(v ITypesmatrizContext) { s._typesmatriz = v }

func (s *DeclarationstmtContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarationstmtContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarationstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *DeclarationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DeclarationstmtContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *DeclarationstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *DeclarationstmtContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIGUAL, 0)
}

func (s *DeclarationstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationstmtContext) CIERRAPREGUNTA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCIERRAPREGUNTA, 0)
}

func (s *DeclarationstmtContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *DeclarationstmtContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *DeclarationstmtContext) Exprvector() IExprvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprvectorContext)
}

func (s *DeclarationstmtContext) Typesmatriz() ITypesmatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesmatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesmatrizContext)
}

func (s *DeclarationstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *DeclarationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclarationstmt(s)
	}
}

func (s *DeclarationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclarationstmt(s)
	}
}

func (p *SwiftGrammarParser) Declarationstmt() (localctx IDeclarationstmtContext) {
	localctx = NewDeclarationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(324)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracion((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), true, localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracion((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), true, environment.DEPENDIENTE, localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(338)
			p.Match(SwiftGrammarParserCIERRAPREGUNTA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracion((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), true, localctx.(*DeclarationstmtContext).Get_types().GetTy(), nil)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(341)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(346)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)

			var _x = p.Exprvector()

			localctx.(*DeclarationstmtContext)._exprvector = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracionVector((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), true, localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_exprvector().GetExprv())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(351)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)

			var _x = p.Typesmatriz()

			localctx.(*DeclarationstmtContext)._typesmatriz = _x
		}
		{
			p.SetState(355)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracionMatriz((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), true, localctx.(*DeclarationstmtContext).Get_typesmatriz().GetTm(), localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(359)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)

			var _x = p.Typesmatriz()

			localctx.(*DeclarationstmtContext)._typesmatriz = _x
		}
		{
			p.SetState(363)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracionMatriz((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), false, localctx.(*DeclarationstmtContext).Get_typesmatriz().GetTm(), localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(367)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(371)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracion((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), false, localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(375)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaracion((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), false, environment.DEPENDIENTE, localctx.(*DeclarationstmtContext).Get_expr().GetE())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignationstmtContext is an interface to support dynamic dispatch.
type IAsignationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetIndex returns the index rule contexts.
	GetIndex() IExprContext

	// GetListan returns the listan rule contexts.
	GetListan() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetIndex sets the index rule contexts.
	SetIndex(IExprContext)

	// SetListan sets the listan rule contexts.
	SetListan(IExprContext)

	// GetAsig returns the asig attribute.
	GetAsig() interfaces.Instruction

	// SetAsig sets the asig attribute.
	SetAsig(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IGUAL() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode
	SUM() antlr.TerminalNode
	RES() antlr.TerminalNode

	// IsAsignationstmtContext differentiates from other interfaces.
	IsAsignationstmtContext()
}

type AsignationstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	asig   interfaces.Instruction
	_ID    antlr.Token
	_expr  IExprContext
	index  IExprContext
	listan IExprContext
	op     antlr.Token
}

func NewEmptyAsignationstmtContext() *AsignationstmtContext {
	var p = new(AsignationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignationstmt
	return p
}

func InitEmptyAsignationstmtContext(p *AsignationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignationstmt
}

func (*AsignationstmtContext) IsAsignationstmtContext() {}

func NewAsignationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignationstmtContext {
	var p = new(AsignationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignationstmt

	return p
}

func (s *AsignationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignationstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignationstmtContext) GetOp() antlr.Token { return s.op }

func (s *AsignationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignationstmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignationstmtContext) Get_expr() IExprContext { return s._expr }

func (s *AsignationstmtContext) GetIndex() IExprContext { return s.index }

func (s *AsignationstmtContext) GetListan() IExprContext { return s.listan }

func (s *AsignationstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AsignationstmtContext) SetIndex(v IExprContext) { s.index = v }

func (s *AsignationstmtContext) SetListan(v IExprContext) { s.listan = v }

func (s *AsignationstmtContext) GetAsig() interfaces.Instruction { return s.asig }

func (s *AsignationstmtContext) SetAsig(v interfaces.Instruction) { s.asig = v }

func (s *AsignationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignationstmtContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIGUAL, 0)
}

func (s *AsignationstmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AsignationstmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignationstmtContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *AsignationstmtContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *AsignationstmtContext) SUM() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUM, 0)
}

func (s *AsignationstmtContext) RES() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRES, 0)
}

func (s *AsignationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignationstmt(s)
	}
}

func (s *AsignationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignationstmt(s)
	}
}

func (p *SwiftGrammarParser) Asignationstmt() (localctx IAsignationstmtContext) {
	localctx = NewAsignationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_asignationstmt)
	var _la int

	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(384)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)

			var _x = p.expr(0)

			localctx.(*AsignationstmtContext)._expr = _x
		}
		localctx.(*AsignationstmtContext).asig = instructions.NewAsignacion((func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignationstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(389)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _x = p.expr(0)

			localctx.(*AsignationstmtContext).index = _x
		}
		{
			p.SetState(391)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)

			var _x = p.expr(0)

			localctx.(*AsignationstmtContext).listan = _x
		}
		localctx.(*AsignationstmtContext).asig = instructions.NewAsignacionIndexVector((func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignationstmtContext).GetIndex().GetE(), localctx.(*AsignationstmtContext).GetListan().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignationstmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserSUM || _la == SwiftGrammarParserRES) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignationstmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(398)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)

			var _x = p.expr(0)

			localctx.(*AsignationstmtContext)._expr = _x
		}
		localctx.(*AsignationstmtContext).asig = instructions.NewOperacionAsignacion((func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AsignationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignationstmtContext).Get_expr().GetE(), (func() string {
			if localctx.(*AsignationstmtContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*AsignationstmtContext).GetOp().GetText()
			}
		}()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsFunc returns the _listParamsFunc rule contexts.
	Get_listParamsFunc() IListParamsFuncContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Set_listParamsFunc sets the _listParamsFunc rule contexts.
	Set_listParamsFunc(IListParamsFuncContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetFun returns the fun attribute.
	GetFun() interfaces.Instruction

	// SetFun sets the fun attribute.
	SetFun(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext
	PAR_DER() antlr.TerminalNode
	LLAVE_IZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVE_DER() antlr.TerminalNode
	FLECHA() antlr.TerminalNode
	Types() ITypesContext
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	fun             interfaces.Instruction
	_FUNC           antlr.Token
	_ID             antlr.Token
	_listParamsFunc IListParamsFuncContext
	_block          IBlockContext
	_types          ITypesContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_listParamsFunc() IListParamsFuncContext { return s._listParamsFunc }

func (s *FunctionContext) Get_block() IBlockContext { return s._block }

func (s *FunctionContext) Get_types() ITypesContext { return s._types }

func (s *FunctionContext) Set_listParamsFunc(v IListParamsFuncContext) { s._listParamsFunc = v }

func (s *FunctionContext) Set_block(v IBlockContext) { s._block = v }

func (s *FunctionContext) Set_types(v ITypesContext) { s._types = v }

func (s *FunctionContext) GetFun() interfaces.Instruction { return s.fun }

func (s *FunctionContext) SetFun(v interfaces.Instruction) { s.fun = v }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FunctionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *FunctionContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *FunctionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *FunctionContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_IZQ, 0)
}

func (s *FunctionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVE_DER, 0)
}

func (s *FunctionContext) FLECHA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLECHA, 0)
}

func (s *FunctionContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *FunctionContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *FunctionContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *SwiftGrammarParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_function)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(408)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(411)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunctionContext).fun = instructions.NewFuncion((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), environment.NULL, localctx.(*FunctionContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(414)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(416)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(418)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)

			var _x = p.Types()

			localctx.(*FunctionContext)._types = _x
		}
		{
			p.SetState(421)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(423)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunctionContext).fun = instructions.NewFuncion((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), localctx.(*FunctionContext).Get_types().GetTy(), localctx.(*FunctionContext).Get_block().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(426)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(430)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Types()
		}
		{
			p.SetState(434)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(437)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FunctionContext).fun = instructions.NewFuncion((func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FunctionContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParamsFunc().GetLpf(), environment.VECTOR, localctx.(*FunctionContext).Get_block().GetBlk())

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsFuncContext is an interface to support dynamic dispatch.
type IListParamsFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListParamsFuncContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListParamsFuncContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetLpf returns the lpf attribute.
	GetLpf() []interface{}

	// SetLpf sets the lpf attribute.
	SetLpf([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Types() ITypesContext
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode
	COMA() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext

	// IsListParamsFuncContext differentiates from other interfaces.
	IsListParamsFuncContext()
}

type ListParamsFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lpf    []interface{}
	list   IListParamsFuncContext
	_ID    antlr.Token
	_types ITypesContext
}

func NewEmptyListParamsFuncContext() *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
	return p
}

func InitEmptyListParamsFuncContext(p *ListParamsFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
}

func (*ListParamsFuncContext) IsListParamsFuncContext() {}

func NewListParamsFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc

	return p
}

func (s *ListParamsFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsFuncContext) Get_ID() antlr.Token { return s._ID }

func (s *ListParamsFuncContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListParamsFuncContext) GetList() IListParamsFuncContext { return s.list }

func (s *ListParamsFuncContext) Get_types() ITypesContext { return s._types }

func (s *ListParamsFuncContext) SetList(v IListParamsFuncContext) { s.list = v }

func (s *ListParamsFuncContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListParamsFuncContext) GetLpf() []interface{} { return s.lpf }

func (s *ListParamsFuncContext) SetLpf(v []interface{}) { s.lpf = v }

func (s *ListParamsFuncContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListParamsFuncContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *ListParamsFuncContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListParamsFuncContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *ListParamsFuncContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *ListParamsFuncContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsFuncContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *ListParamsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsFunc(s)
	}
}

func (s *ListParamsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsFunc(s)
	}
}

func (p *SwiftGrammarParser) ListParamsFunc() (localctx IListParamsFuncContext) {
	return p.listParamsFunc(0)
}

func (p *SwiftGrammarParser) listParamsFunc(_p int) (localctx IListParamsFuncContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsFuncContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsFuncContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SwiftGrammarParserRULE_listParamsFunc, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(443)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)

			var _x = p.Types()

			localctx.(*ListParamsFuncContext)._types = _x
		}

		localctx.(*ListParamsFuncContext).lpf = []interface{}{}
		newParam := instructions.NewDeclaracionParametros((func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetText()
			}
		}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
		localctx.(*ListParamsFuncContext).lpf = append(localctx.(*ListParamsFuncContext).lpf, newParam)

	case 2:
		{
			p.SetState(448)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(451)

			var _x = p.Types()

			localctx.(*ListParamsFuncContext)._types = _x
		}
		{
			p.SetState(452)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*ListParamsFuncContext).lpf = []interface{}{}
		newParam := instructions.NewDeclaracionParametros((func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ListParamsFuncContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListParamsFuncContext).Get_ID().GetText()
			}
		}()), environment.VECTOR)
		localctx.(*ListParamsFuncContext).lpf = append(localctx.(*ListParamsFuncContext).lpf, newParam)

	case 3:
		localctx.(*ListParamsFuncContext).lpf = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(474)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
				localctx.(*ListParamsFuncContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
				p.SetState(458)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(459)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(460)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListParamsFuncContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(461)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(462)

					var _x = p.Types()

					localctx.(*ListParamsFuncContext)._types = _x
				}

				var arr []interface{}
				newParam := instructions.NewDeclaracionParametros((func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
					}
				}()), (func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
					}
				}()), (func() string {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetText()
					}
				}()), localctx.(*ListParamsFuncContext).Get_types().GetTy())
				arr = append(localctx.(*ListParamsFuncContext).GetList().GetLpf(), newParam)
				localctx.(*ListParamsFuncContext).lpf = arr

			case 2:
				localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
				localctx.(*ListParamsFuncContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
				p.SetState(465)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(466)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(467)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListParamsFuncContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(468)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(469)
					p.Match(SwiftGrammarParserCOR_IZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(470)

					var _x = p.Types()

					localctx.(*ListParamsFuncContext)._types = _x
				}
				{
					p.SetState(471)
					p.Match(SwiftGrammarParserCOR_DER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				var arr []interface{}
				newParam := instructions.NewDeclaracionParametros((func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetLine()
					}
				}()), (func() int {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetColumn()
					}
				}()), (func() string {
					if localctx.(*ListParamsFuncContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListParamsFuncContext).Get_ID().GetText()
					}
				}()), environment.VECTOR)
				arr = append(localctx.(*ListParamsFuncContext).GetList().GetLpf(), newParam)
				localctx.(*ListParamsFuncContext).lpf = arr

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFuncionInsContext is an interface to support dynamic dispatch.
type ICallFuncionInsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCf returns the cf attribute.
	GetCf() interfaces.Expression

	// SetCf sets the cf attribute.
	SetCf(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext
	PAR_DER() antlr.TerminalNode

	// IsCallFuncionInsContext differentiates from other interfaces.
	IsCallFuncionInsContext()
}

type CallFuncionInsContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	cf              interfaces.Expression
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallFuncionInsContext() *CallFuncionInsContext {
	var p = new(CallFuncionInsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFuncionIns
	return p
}

func InitEmptyCallFuncionInsContext(p *CallFuncionInsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFuncionIns
}

func (*CallFuncionInsContext) IsCallFuncionInsContext() {}

func NewCallFuncionInsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFuncionInsContext {
	var p = new(CallFuncionInsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callFuncionIns

	return p
}

func (s *CallFuncionInsContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFuncionInsContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFuncionInsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFuncionInsContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallFuncionInsContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallFuncionInsContext) GetCf() interfaces.Expression { return s.cf }

func (s *CallFuncionInsContext) SetCf(v interfaces.Expression) { s.cf = v }

func (s *CallFuncionInsContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallFuncionInsContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *CallFuncionInsContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallFuncionInsContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *CallFuncionInsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFuncionInsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFuncionInsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallFuncionIns(s)
	}
}

func (s *CallFuncionInsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallFuncionIns(s)
	}
}

func (p *SwiftGrammarParser) CallFuncionIns() (localctx ICallFuncionInsContext) {
	localctx = NewCallFuncionInsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_callFuncionIns)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFuncionInsContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(481)

		var _x = p.listParamsCall(0)

		localctx.(*CallFuncionInsContext)._listParamsCall = _x
	}
	{
		p.SetState(482)
		p.Match(SwiftGrammarParserPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallFuncionInsContext).cf = expressions.NewLlamadoFuncion((func() int {
		if localctx.(*CallFuncionInsContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFuncionInsContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallFuncionInsContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFuncionInsContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallFuncionInsContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallFuncionInsContext).Get_ID().GetText()
		}
	}()), localctx.(*CallFuncionInsContext).Get_listParamsCall().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty attribute.
	GetTy() environment.TipoExpresion

	// SetTy sets the ty attribute.
	SetTy(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     environment.TipoExpresion
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) GetTy() environment.TipoExpresion { return s.ty }

func (s *TypesContext) SetTy(v environment.TipoExpresion) { s.ty = v }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TypesContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TypesContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *TypesContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *SwiftGrammarParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_types)
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(485)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(489)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(491)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case SwiftGrammarParserCOR_IZQ:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(493)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesmatrizContext is an interface to support dynamic dispatch.
type ITypesmatrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() ITypesmatrizContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(ITypesmatrizContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetTm returns the tm attribute.
	GetTm() []interface{}

	// SetTm sets the tm attribute.
	SetTm([]interface{})

	// Getter signatures
	COR_IZQ() antlr.TerminalNode
	COR_DER() antlr.TerminalNode
	Typesmatriz() ITypesmatrizContext
	Types() ITypesContext

	// IsTypesmatrizContext differentiates from other interfaces.
	IsTypesmatrizContext()
}

type TypesmatrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tm     []interface{}
	list   ITypesmatrizContext
	_types ITypesContext
}

func NewEmptyTypesmatrizContext() *TypesmatrizContext {
	var p = new(TypesmatrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typesmatriz
	return p
}

func InitEmptyTypesmatrizContext(p *TypesmatrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_typesmatriz
}

func (*TypesmatrizContext) IsTypesmatrizContext() {}

func NewTypesmatrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesmatrizContext {
	var p = new(TypesmatrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_typesmatriz

	return p
}

func (s *TypesmatrizContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesmatrizContext) GetList() ITypesmatrizContext { return s.list }

func (s *TypesmatrizContext) Get_types() ITypesContext { return s._types }

func (s *TypesmatrizContext) SetList(v ITypesmatrizContext) { s.list = v }

func (s *TypesmatrizContext) Set_types(v ITypesContext) { s._types = v }

func (s *TypesmatrizContext) GetTm() []interface{} { return s.tm }

func (s *TypesmatrizContext) SetTm(v []interface{}) { s.tm = v }

func (s *TypesmatrizContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *TypesmatrizContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *TypesmatrizContext) Typesmatriz() ITypesmatrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesmatrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesmatrizContext)
}

func (s *TypesmatrizContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TypesmatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesmatrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesmatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypesmatriz(s)
	}
}

func (s *TypesmatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypesmatriz(s)
	}
}

func (p *SwiftGrammarParser) Typesmatriz() (localctx ITypesmatrizContext) {
	localctx = NewTypesmatrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_typesmatriz)
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(498)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(499)

			var _x = p.Typesmatriz()

			localctx.(*TypesmatrizContext).list = _x
		}
		{
			p.SetState(500)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		var arr []interface{}
		newTipo := environment.NewTipoArray(environment.ARRAY)
		arr = append(localctx.(*TypesmatrizContext).GetList().GetTm(), newTipo)
		localctx.(*TypesmatrizContext).tm = arr

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)

			var _x = p.Types()

			localctx.(*TypesmatrizContext)._types = _x
		}

		localctx.(*TypesmatrizContext).tm = []interface{}{}
		newTipo := environment.NewTipoArray(localctx.(*TypesmatrizContext).Get_types().GetTy())
		localctx.(*TypesmatrizContext).tm = append(localctx.(*TypesmatrizContext).tm, newTipo)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprForContext is an interface to support dynamic dispatch.
type IExprForContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRange1 returns the range1 rule contexts.
	GetRange1() IExprContext

	// GetRange2 returns the range2 rule contexts.
	GetRange2() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetRange1 sets the range1 rule contexts.
	SetRange1(IExprContext)

	// SetRange2 sets the range2 rule contexts.
	SetRange2(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	AllPUNTO() []antlr.TerminalNode
	PUNTO(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprForContext differentiates from other interfaces.
	IsExprForContext()
}

type ExprForContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	e      interfaces.Expression
	range1 IExprContext
	range2 IExprContext
	_expr  IExprContext
}

func NewEmptyExprForContext() *ExprForContext {
	var p = new(ExprForContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprFor
	return p
}

func InitEmptyExprForContext(p *ExprForContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprFor
}

func (*ExprForContext) IsExprForContext() {}

func NewExprForContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprForContext {
	var p = new(ExprForContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_exprFor

	return p
}

func (s *ExprForContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprForContext) GetRange1() IExprContext { return s.range1 }

func (s *ExprForContext) GetRange2() IExprContext { return s.range2 }

func (s *ExprForContext) Get_expr() IExprContext { return s._expr }

func (s *ExprForContext) SetRange1(v IExprContext) { s.range1 = v }

func (s *ExprForContext) SetRange2(v IExprContext) { s.range2 = v }

func (s *ExprForContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprForContext) GetE() interfaces.Expression { return s.e }

func (s *ExprForContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprForContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPUNTO)
}

func (s *ExprForContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, i)
}

func (s *ExprForContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprForContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprForContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExprFor(s)
	}
}

func (s *ExprForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExprFor(s)
	}
}

func (p *SwiftGrammarParser) ExprFor() (localctx IExprForContext) {
	localctx = NewExprForContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_exprFor)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(508)

			var _x = p.expr(0)

			localctx.(*ExprForContext).range1 = _x
		}
		{
			p.SetState(509)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(512)

			var _x = p.expr(0)

			localctx.(*ExprForContext).range2 = _x
		}
		localctx.(*ExprForContext).e = expressions.NewForRange((func() antlr.Token {
			if localctx.(*ExprForContext).GetRange1() == nil {
				return nil
			} else {
				return localctx.(*ExprForContext).GetRange1().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprForContext).GetRange1() == nil {
				return nil
			} else {
				return localctx.(*ExprForContext).GetRange1().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprForContext).GetRange1().GetE(), localctx.(*ExprForContext).GetRange2().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(515)

			var _x = p.expr(0)

			localctx.(*ExprForContext)._expr = _x
		}
		localctx.(*ExprForContext).e = localctx.(*ExprForContext).Get_expr().GetE()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RES returns the _RES token.
	Get_RES() antlr.Token

	// Get_NOT returns the _NOT token.
	Get_NOT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_COR_IZQ returns the _COR_IZQ token.
	Get_COR_IZQ() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_RES sets the _RES token.
	Set_RES(antlr.Token)

	// Set_NOT sets the _NOT token.
	Set_NOT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_COR_IZQ sets the _COR_IZQ token.
	Set_COR_IZQ(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_listStructExp returns the _listStructExp rule contexts.
	Get_listStructExp() IListStructExpContext

	// Get_callFuncion returns the _callFuncion rule contexts.
	Get_callFuncion() ICallFuncionContext

	// Get_conversionstmt returns the _conversionstmt rule contexts.
	Get_conversionstmt() IConversionstmtContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_listStructExp sets the _listStructExp rule contexts.
	Set_listStructExp(IListStructExpContext)

	// Set_callFuncion sets the _callFuncion rule contexts.
	Set_callFuncion(ICallFuncionContext)

	// Set_conversionstmt sets the _conversionstmt rule contexts.
	Set_conversionstmt(IConversionstmtContext)

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	RES() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	ID() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	ListStructExp() IListStructExpContext
	PAR_DER() antlr.TerminalNode
	CallFuncion() ICallFuncionContext
	Conversionstmt() IConversionstmtContext
	PUNTO() antlr.TerminalNode
	COUNT() antlr.TerminalNode
	ISEMPTY() antlr.TerminalNode
	ListArray() IListArrayContext
	COR_IZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	COR_DER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	NIL() antlr.TerminalNode
	MULT() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	SUM() antlr.TerminalNode
	MAYIG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MENIG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIFE() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	e               interfaces.Expression
	left            IExprContext
	_RES            antlr.Token
	_expr           IExprContext
	_NOT            antlr.Token
	_ID             antlr.Token
	_listStructExp  IListStructExpContext
	_callFuncion    ICallFuncionContext
	_conversionstmt IConversionstmtContext
	list            IListArrayContext
	_COR_IZQ        antlr.Token
	_listParams     IListParamsContext
	_NUMBER         antlr.Token
	_STRING         antlr.Token
	_TRU            antlr.Token
	_FAL            antlr.Token
	_NIL            antlr.Token
	op              antlr.Token
	right           IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Get_RES() antlr.Token { return s._RES }

func (s *ExprContext) Get_NOT() antlr.Token { return s._NOT }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_COR_IZQ() antlr.Token { return s._COR_IZQ }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) Get_NIL() antlr.Token { return s._NIL }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Set_RES(v antlr.Token) { s._RES = v }

func (s *ExprContext) Set_NOT(v antlr.Token) { s._NOT = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_COR_IZQ(v antlr.Token) { s._COR_IZQ = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_listStructExp() IListStructExpContext { return s._listStructExp }

func (s *ExprContext) Get_callFuncion() ICallFuncionContext { return s._callFuncion }

func (s *ExprContext) Get_conversionstmt() IConversionstmtContext { return s._conversionstmt }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_listStructExp(v IListStructExpContext) { s._listStructExp = v }

func (s *ExprContext) Set_callFuncion(v ICallFuncionContext) { s._callFuncion = v }

func (s *ExprContext) Set_conversionstmt(v IConversionstmtContext) { s._conversionstmt = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) RES() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRES, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *ExprContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ExprContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *ExprContext) CallFuncion() ICallFuncionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFuncionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFuncionContext)
}

func (s *ExprContext) Conversionstmt() IConversionstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConversionstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConversionstmtContext)
}

func (s *ExprContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *ExprContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserISEMPTY, 0)
}

func (s *ExprContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ExprContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *ExprContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
}

func (s *ExprContext) MULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMULT, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *ExprContext) SUM() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUM, 0)
}

func (s *ExprContext) MAYIG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYIG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MENIG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENIG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIFE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIFE, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(521)

			var _m = p.Match(SwiftGrammarParserRES)

			localctx.(*ExprContext)._RES = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)

			var _x = p.expr(22)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() int {
			if localctx.(*ExprContext).Get_RES() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RES().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_RES() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RES().GetColumn()
			}
		}()), localctx.(*ExprContext).GetLeft().GetE(), "UNARIO", nil)

	case 2:
		{
			p.SetState(525)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext)._NOT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)

			var _x = p.expr(16)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() int {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NOT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NOT().GetColumn()
			}
		}()), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NOT().GetText()
			}
		}()), nil)

	case 3:
		{
			p.SetState(529)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(530)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(531)

			var _x = p.listStructExp(0)

			localctx.(*ExprContext)._listStructExp = _x
		}
		{
			p.SetState(532)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewStructExp((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()), localctx.(*ExprContext).Get_listStructExp().GetL())

	case 4:
		{
			p.SetState(535)

			var _x = p.CallFuncion()

			localctx.(*ExprContext)._callFuncion = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_callFuncion().GetCf()

	case 5:
		{
			p.SetState(538)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(540)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 6:
		{
			p.SetState(543)

			var _x = p.Conversionstmt()

			localctx.(*ExprContext)._conversionstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_conversionstmt().GetConv()

	case 7:
		{
			p.SetState(546)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCount((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 8:
		{
			p.SetState(550)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Match(SwiftGrammarParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewIsEmpty((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 9:
		{
			p.SetState(554)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case 10:
		{
			p.SetState(557)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(559)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_COR_IZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_COR_IZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listParams().GetL())

	case 11:
		{
			p.SetState(562)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 12:
		{
			p.SetState(564)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 13:
		{
			p.SetState(566)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 14:
		{
			p.SetState(568)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case 15:
		{
			p.SetState(570)

			var _m = p.Match(SwiftGrammarParserNIL)

			localctx.(*ExprContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetColumn()
			}
		}()), nil, environment.NULL)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(609)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(574)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(575)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1337006139375616) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(576)

					var _x = p.expr(22)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(579)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(580)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserSUM || _la == SwiftGrammarParserRES) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(581)

					var _x = p.expr(21)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(584)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(585)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAYIG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(586)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(589)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(590)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMENIG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(591)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(594)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(595)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIFE || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(596)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(599)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(600)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(601)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(604)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(605)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(606)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(613)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConversionstmtContext is an interface to support dynamic dispatch.
type IConversionstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_STR returns the _STR token.
	Get_STR() antlr.Token

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_STR sets the _STR token.
	Set_STR(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetConv returns the conv attribute.
	GetConv() interfaces.Expression

	// SetConv sets the conv attribute.
	SetConv(interfaces.Expression)

	// Getter signatures
	INT() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	Expr() IExprContext
	PAR_DER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode

	// IsConversionstmtContext differentiates from other interfaces.
	IsConversionstmtContext()
}

type ConversionstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	conv   interfaces.Expression
	_INT   antlr.Token
	_expr  IExprContext
	_FLOAT antlr.Token
	_STR   antlr.Token
}

func NewEmptyConversionstmtContext() *ConversionstmtContext {
	var p = new(ConversionstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_conversionstmt
	return p
}

func InitEmptyConversionstmtContext(p *ConversionstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_conversionstmt
}

func (*ConversionstmtContext) IsConversionstmtContext() {}

func NewConversionstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConversionstmtContext {
	var p = new(ConversionstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_conversionstmt

	return p
}

func (s *ConversionstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ConversionstmtContext) Get_INT() antlr.Token { return s._INT }

func (s *ConversionstmtContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *ConversionstmtContext) Get_STR() antlr.Token { return s._STR }

func (s *ConversionstmtContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *ConversionstmtContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *ConversionstmtContext) Set_STR(v antlr.Token) { s._STR = v }

func (s *ConversionstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ConversionstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ConversionstmtContext) GetConv() interfaces.Expression { return s.conv }

func (s *ConversionstmtContext) SetConv(v interfaces.Expression) { s.conv = v }

func (s *ConversionstmtContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *ConversionstmtContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *ConversionstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConversionstmtContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *ConversionstmtContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *ConversionstmtContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *ConversionstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConversionstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConversionstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterConversionstmt(s)
	}
}

func (s *ConversionstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitConversionstmt(s)
	}
}

func (p *SwiftGrammarParser) Conversionstmt() (localctx IConversionstmtContext) {
	localctx = NewConversionstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_conversionstmt)
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(614)

			var _m = p.Match(SwiftGrammarParserINT)

			localctx.(*ConversionstmtContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(615)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(617)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ConversionstmtContext).conv = expressions.NewToInt((func() int {
			if localctx.(*ConversionstmtContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_INT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConversionstmtContext).Get_INT() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_INT().GetColumn()
			}
		}()), localctx.(*ConversionstmtContext).Get_expr().GetE())

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(620)

			var _m = p.Match(SwiftGrammarParserFLOAT)

			localctx.(*ConversionstmtContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(623)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ConversionstmtContext).conv = expressions.NewToFloat((func() int {
			if localctx.(*ConversionstmtContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_FLOAT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConversionstmtContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_FLOAT().GetColumn()
			}
		}()), localctx.(*ConversionstmtContext).Get_expr().GetE())

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(626)

			var _m = p.Match(SwiftGrammarParserSTR)

			localctx.(*ConversionstmtContext)._STR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(627)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(628)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(629)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ConversionstmtContext).conv = expressions.NewToString((func() int {
			if localctx.(*ConversionstmtContext).Get_STR() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_STR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConversionstmtContext).Get_STR() == nil {
				return 0
			} else {
				return localctx.(*ConversionstmtContext).Get_STR().GetColumn()
			}
		}()), localctx.(*ConversionstmtContext).Get_expr().GetE())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprvectorContext is an interface to support dynamic dispatch.
type IExprvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_COR_IZQ returns the _COR_IZQ token.
	Get_COR_IZQ() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_COR_IZQ sets the _COR_IZQ token.
	Set_COR_IZQ(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetExprv returns the exprv attribute.
	GetExprv() interfaces.Expression

	// SetExprv sets the exprv attribute.
	SetExprv(interfaces.Expression)

	// Getter signatures
	COR_IZQ() antlr.TerminalNode
	ListParams() IListParamsContext
	COR_DER() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsExprvectorContext differentiates from other interfaces.
	IsExprvectorContext()
}

type ExprvectorContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	exprv       interfaces.Expression
	_COR_IZQ    antlr.Token
	_listParams IListParamsContext
	_ID         antlr.Token
}

func NewEmptyExprvectorContext() *ExprvectorContext {
	var p = new(ExprvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprvector
	return p
}

func InitEmptyExprvectorContext(p *ExprvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprvector
}

func (*ExprvectorContext) IsExprvectorContext() {}

func NewExprvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprvectorContext {
	var p = new(ExprvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_exprvector

	return p
}

func (s *ExprvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprvectorContext) Get_COR_IZQ() antlr.Token { return s._COR_IZQ }

func (s *ExprvectorContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprvectorContext) Set_COR_IZQ(v antlr.Token) { s._COR_IZQ = v }

func (s *ExprvectorContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprvectorContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprvectorContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprvectorContext) GetExprv() interfaces.Expression { return s.exprv }

func (s *ExprvectorContext) SetExprv(v interfaces.Expression) { s.exprv = v }

func (s *ExprvectorContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *ExprvectorContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprvectorContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *ExprvectorContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExprvector(s)
	}
}

func (s *ExprvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExprvector(s)
	}
}

func (p *SwiftGrammarParser) Exprvector() (localctx IExprvectorContext) {
	localctx = NewExprvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_exprvector)
	p.SetState(644)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(634)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprvectorContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(635)

			var _x = p.listParams(0)

			localctx.(*ExprvectorContext)._listParams = _x
		}
		{
			p.SetState(636)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprvectorContext).exprv = expressions.NewVector((func() int {
			if localctx.(*ExprvectorContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_COR_IZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprvectorContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_COR_IZQ().GetColumn()
			}
		}()), localctx.(*ExprvectorContext).Get_listParams().GetL())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(639)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprvectorContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(640)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprvectorContext).exprv = expressions.NewVector((func() int {
			if localctx.(*ExprvectorContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_COR_IZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprvectorContext).Get_COR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_COR_IZQ().GetColumn()
			}
		}()), nil)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(642)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprvectorContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprvectorContext).exprv = expressions.NewLlamadoVar((func() int {
			if localctx.(*ExprvectorContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprvectorContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprvectorContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprvectorContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprvectorContext).Get_ID().GetText()
			}
		}()))

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *SwiftGrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *SwiftGrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(657)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(650)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(651)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(652)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(659)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// GetArr returns the arr rule contexts.
	GetArr() IListAccessArrayContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// SetArr sets the arr rule contexts.
	SetArr(IListAccessArrayContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	ListArray() IListArrayContext
	ListAccessArray() IListAccessArrayContext
	PUNTO() antlr.TerminalNode

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expression
	list   IListArrayContext
	_ID    antlr.Token
	arr    IListAccessArrayContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
	return p
}

func InitEmptyListArrayContext(p *ListArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) GetArr() IListAccessArrayContext { return s.arr }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) SetArr(v IListAccessArrayContext) { s.arr = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) ListAccessArray() IListAccessArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAccessArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAccessArrayContext)
}

func (s *ListArrayContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, 0)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *SwiftGrammarParser) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *SwiftGrammarParser) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(661)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ListArrayContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ListArrayContext).p = expressions.NewLlamadoVar((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(672)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(664)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(665)

					var _x = p.listAccessArray(0)

					localctx.(*ListArrayContext).arr = _x
				}
				localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() int {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetLine()
					}
				}()), (func() int {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return 0
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetColumn()
					}
				}()), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).GetArr().GetL())

			case 2:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(668)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(669)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(670)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListArrayContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewStructAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), (func() string {
					if localctx.(*ListArrayContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListArrayContext).Get_ID().GetText()
					}
				}()))

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(676)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListAccessArrayContext is an interface to support dynamic dispatch.
type IListAccessArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListAccessArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListAccessArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	COR_IZQ() antlr.TerminalNode
	Expr() IExprContext
	COR_DER() antlr.TerminalNode
	ListAccessArray() IListAccessArrayContext

	// IsListAccessArrayContext differentiates from other interfaces.
	IsListAccessArrayContext()
}

type ListAccessArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListAccessArrayContext
	_expr  IExprContext
}

func NewEmptyListAccessArrayContext() *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray
	return p
}

func InitEmptyListAccessArrayContext(p *ListAccessArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray
}

func (*ListAccessArrayContext) IsListAccessArrayContext() {}

func NewListAccessArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray

	return p
}

func (s *ListAccessArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAccessArrayContext) GetList() IListAccessArrayContext { return s.list }

func (s *ListAccessArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListAccessArrayContext) SetList(v IListAccessArrayContext) { s.list = v }

func (s *ListAccessArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListAccessArrayContext) GetL() []interface{} { return s.l }

func (s *ListAccessArrayContext) SetL(v []interface{}) { s.l = v }

func (s *ListAccessArrayContext) COR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_IZQ, 0)
}

func (s *ListAccessArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListAccessArrayContext) COR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOR_DER, 0)
}

func (s *ListAccessArrayContext) ListAccessArray() IListAccessArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAccessArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAccessArrayContext)
}

func (s *ListAccessArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAccessArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListAccessArray(s)
	}
}

func (s *ListAccessArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListAccessArray(s)
	}
}

func (p *SwiftGrammarParser) ListAccessArray() (localctx IListAccessArrayContext) {
	return p.listAccessArray(0)
}

func (p *SwiftGrammarParser) listAccessArray(_p int) (localctx IListAccessArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListAccessArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListAccessArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, SwiftGrammarParserRULE_listAccessArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(678)
		p.Match(SwiftGrammarParserCOR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(679)

		var _x = p.expr(0)

		localctx.(*ListAccessArrayContext)._expr = _x
	}
	{
		p.SetState(680)
		p.Match(SwiftGrammarParserCOR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ListAccessArrayContext).l = []interface{}{}
	localctx.(*ListAccessArrayContext).l = append(localctx.(*ListAccessArrayContext).l, localctx.(*ListAccessArrayContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(691)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAccessArrayContext(p, _parentctx, _parentState)
			localctx.(*ListAccessArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listAccessArray)
			p.SetState(683)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(684)
				p.Match(SwiftGrammarParserCOR_IZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(685)

				var _x = p.expr(0)

				localctx.(*ListAccessArrayContext)._expr = _x
			}
			{
				p.SetState(686)
				p.Match(SwiftGrammarParserCOR_DER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			var arr []interface{}
			arr = append(localctx.(*ListAccessArrayContext).GetList().GetL(), localctx.(*ListAccessArrayContext).Get_expr().GetE())
			localctx.(*ListAccessArrayContext).l = arr

		}
		p.SetState(693)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFuncionContext is an interface to support dynamic dispatch.
type ICallFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCf returns the cf attribute.
	GetCf() interfaces.Expression

	// SetCf sets the cf attribute.
	SetCf(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext
	PAR_DER() antlr.TerminalNode

	// IsCallFuncionContext differentiates from other interfaces.
	IsCallFuncionContext()
}

type CallFuncionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	cf              interfaces.Expression
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallFuncionContext() *CallFuncionContext {
	var p = new(CallFuncionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFuncion
	return p
}

func InitEmptyCallFuncionContext(p *CallFuncionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFuncion
}

func (*CallFuncionContext) IsCallFuncionContext() {}

func NewCallFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFuncionContext {
	var p = new(CallFuncionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callFuncion

	return p
}

func (s *CallFuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFuncionContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFuncionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFuncionContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallFuncionContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallFuncionContext) GetCf() interfaces.Expression { return s.cf }

func (s *CallFuncionContext) SetCf(v interfaces.Expression) { s.cf = v }

func (s *CallFuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallFuncionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_IZQ, 0)
}

func (s *CallFuncionContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallFuncionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPAR_DER, 0)
}

func (s *CallFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallFuncion(s)
	}
}

func (s *CallFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallFuncion(s)
	}
}

func (p *SwiftGrammarParser) CallFuncion() (localctx ICallFuncionContext) {
	localctx = NewCallFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_callFuncion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(694)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFuncionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(695)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(696)

		var _x = p.listParamsCall(0)

		localctx.(*CallFuncionContext)._listParamsCall = _x
	}
	{
		p.SetState(697)
		p.Match(SwiftGrammarParserPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallFuncionContext).cf = expressions.NewLlamadoFuncion((func() int {
		if localctx.(*CallFuncionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFuncionContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallFuncionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFuncionContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallFuncionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallFuncionContext).Get_ID().GetText()
		}
	}()), localctx.(*CallFuncionContext).Get_listParamsCall().GetL())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListParamsCallContext is an interface to support dynamic dispatch.
type IListParamsCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsCallContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsCallContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext

	// IsListParamsCallContext differentiates from other interfaces.
	IsListParamsCallContext()
}

type ListParamsCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsCallContext
	_expr  IExprContext
}

func NewEmptyListParamsCallContext() *ListParamsCallContext {
	var p = new(ListParamsCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
	return p
}

func InitEmptyListParamsCallContext(p *ListParamsCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
}

func (*ListParamsCallContext) IsListParamsCallContext() {}

func NewListParamsCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsCallContext {
	var p = new(ListParamsCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall

	return p
}

func (s *ListParamsCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsCallContext) GetList() IListParamsCallContext { return s.list }

func (s *ListParamsCallContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsCallContext) SetList(v IListParamsCallContext) { s.list = v }

func (s *ListParamsCallContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsCallContext) GetL() []interface{} { return s.l }

func (s *ListParamsCallContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsCallContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsCallContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *ListParamsCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsCall(s)
	}
}

func (s *ListParamsCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsCall(s)
	}
}

func (p *SwiftGrammarParser) ListParamsCall() (localctx IListParamsCallContext) {
	return p.listParamsCall(0)
}

func (p *SwiftGrammarParser) listParamsCall(_p int) (localctx IListParamsCallContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsCallContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsCallContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, SwiftGrammarParserRULE_listParamsCall, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(705)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(701)

			var _x = p.expr(0)

			localctx.(*ListParamsCallContext)._expr = _x
		}

		localctx.(*ListParamsCallContext).l = []interface{}{}
		localctx.(*ListParamsCallContext).l = append(localctx.(*ListParamsCallContext).l, localctx.(*ListParamsCallContext).Get_expr().GetE())

	case 2:

		localctx.(*ListParamsCallContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(714)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsCallContext(p, _parentctx, _parentState)
			localctx.(*ListParamsCallContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsCall)
			p.SetState(707)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(708)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(709)

				var _x = p.expr(0)

				localctx.(*ListParamsCallContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsCallContext).GetList().GetL(), localctx.(*ListParamsCallContext).Get_expr().GetE())
			localctx.(*ListParamsCallContext).l = arr

		}
		p.SetState(716)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListStructExpContext is an interface to support dynamic dispatch.
type IListStructExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructExpContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListStructExpContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	ID() antlr.TerminalNode
	DOSPUNTOS() antlr.TerminalNode
	Expr() IExprContext
	ListStructExp() IListStructExpContext
	COMA() antlr.TerminalNode

	// IsListStructExpContext differentiates from other interfaces.
	IsListStructExpContext()
}

type ListStructExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructExpContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListStructExpContext() *ListStructExpContext {
	var p = new(ListStructExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
	return p
}

func InitEmptyListStructExpContext(p *ListStructExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp
}

func (*ListStructExpContext) IsListStructExpContext() {}

func NewListStructExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructExpContext {
	var p = new(ListStructExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructExp

	return p
}

func (s *ListStructExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructExpContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructExpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructExpContext) GetList() IListStructExpContext { return s.list }

func (s *ListStructExpContext) Get_expr() IExprContext { return s._expr }

func (s *ListStructExpContext) SetList(v IListStructExpContext) { s.list = v }

func (s *ListStructExpContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListStructExpContext) GetL() []interface{} { return s.l }

func (s *ListStructExpContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructExpContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListStructExpContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPUNTOS, 0)
}

func (s *ListStructExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListStructExpContext) ListStructExp() IListStructExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructExpContext)
}

func (s *ListStructExpContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructExp(s)
	}
}

func (s *ListStructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructExp(s)
	}
}

func (p *SwiftGrammarParser) ListStructExp() (localctx IListStructExpContext) {
	return p.listStructExp(0)
}

func (p *SwiftGrammarParser) listStructExp(_p int) (localctx IListStructExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, SwiftGrammarParserRULE_listStructExp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(724)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(718)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructExpContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(719)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(720)

			var _x = p.expr(0)

			localctx.(*ListStructExpContext)._expr = _x
		}

		var arr []interface{}
		StrExp := environment.NewStructContent((func() string {
			if localctx.(*ListStructExpContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructExpContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructExpContext).Get_expr().GetE())
		arr = append(arr, StrExp)
		localctx.(*ListStructExpContext).l = arr

	case 2:

		localctx.(*ListStructExpContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(737)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListStructExpContext(p, _parentctx, _parentState)
			localctx.(*ListStructExpContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructExp)
			p.SetState(726)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			p.SetState(728)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SwiftGrammarParserCOMA {
				{
					p.SetState(727)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(730)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructExpContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(731)
				p.Match(SwiftGrammarParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(732)

				var _x = p.expr(0)

				localctx.(*ListStructExpContext)._expr = _x
			}

			var arr []interface{}
			StrExp := environment.NewStructContent((func() string {
				if localctx.(*ListStructExpContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListStructExpContext).Get_ID().GetText()
				}
			}()), localctx.(*ListStructExpContext).Get_expr().GetE())
			arr = append(localctx.(*ListStructExpContext).GetList().GetL(), StrExp)
			localctx.(*ListStructExpContext).l = arr

		}
		p.SetState(739)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 18:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 23:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 26:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 27:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 28:
		var t *ListAccessArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessArrayContext)
		}
		return p.ListAccessArray_Sempred(t, predIndex)

	case 30:
		var t *ListParamsCallContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsCallContext)
		}
		return p.ListParamsCall_Sempred(t, predIndex)

	case 31:
		var t *ListStructExpContext = nil
		if localctx != nil {
			t = localctx.(*ListStructExpContext)
		}
		return p.ListStructExp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsFunc_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 14)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListAccessArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsCall_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListStructExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

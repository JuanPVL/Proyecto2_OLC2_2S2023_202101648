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
		"'if'", "'else'", "'while'", "'for'", "'in'", "'var'", "'let'", "'nil'",
		"'break'", "'continue'", "'append'", "'removeLast'", "'remove'", "'at'",
		"'isEmpty'", "'count'", "'array'", "'return'", "'func'", "'struct'",
		"'guard'", "", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='",
		"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('",
		"')'", "'{'", "'}'", "':'", "'['", "']'", "','", "'?'", "';'", "'.'",
		"'->'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE",
		"WHILE", "FOR", "IN", "VAR", "LET", "NIL", "BREAK", "CONTINUE", "APPEND",
		"REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "ARRAY", "RETURN",
		"FUNC", "STRUCT", "GUARD", "NUMBER", "STRING", "CHAR", "ID", "DIFE",
		"IG_IG", "NOT", "OR", "AND", "IGUAL", "MAYIG", "MENIG", "MAYOR", "MENOR",
		"MULT", "DIV", "SUM", "RES", "MOD", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ",
		"LLAVE_DER", "DOSPUNTOS", "COR_IZQ", "COR_DER", "COMA", "CIERRAPREGUNTA",
		"PUNTOCOMA", "PUNTO", "FLECHA", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "structCreation", "listStructDec", "printstmt",
		"blockelsif", "ifstmt", "whilestmt", "guardstmt", "forstmt", "declarationstmt",
		"asignationstmt", "function", "listParamsFunc", "callFuncionIns", "types",
		"typesmatriz", "exprFor", "expr", "conversionstmt", "exprvector", "listParams",
		"listArray", "listAccessArray", "callFuncion", "listParamsCall", "listStructExp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 687, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 62, 8, 1, 11,
		1, 12, 1, 63, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 79, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 85,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 106, 8, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 112, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 117, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 127, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 137, 8, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 149, 8, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 3, 2, 156, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 162,
		8, 2, 1, 2, 3, 2, 165, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 187, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 204, 8, 4, 10, 4, 12, 4,
		207, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 4, 6, 216, 8, 6, 11,
		6, 12, 6, 217, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 249, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 328, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 349, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 387,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 403, 8, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 421, 8, 14, 10, 14, 12, 14, 424, 9, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 443, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 453, 8, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 465,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 3, 19, 519, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 556, 8, 19, 10,
		19, 12, 19, 559, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 579, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 3, 21, 591, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 602, 8, 22, 10, 22, 12, 22, 605,
		9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 5, 23, 619, 8, 23, 10, 23, 12, 23, 622, 9, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 636, 8, 24, 10, 24, 12, 24, 639, 9, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 652, 8,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 659, 8, 26, 10, 26, 12, 26,
		662, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 671,
		8, 27, 1, 27, 1, 27, 3, 27, 675, 8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 5, 27, 682, 8, 27, 10, 27, 12, 27, 685, 9, 27, 1, 27, 0, 8, 8, 28,
		38, 44, 46, 48, 52, 54, 28, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 0, 5, 1,
		0, 45, 46, 2, 0, 43, 44, 47, 47, 2, 0, 39, 39, 41, 41, 2, 0, 40, 40, 42,
		42, 1, 0, 33, 34, 747, 0, 56, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 164, 1,
		0, 0, 0, 6, 166, 1, 0, 0, 0, 8, 186, 1, 0, 0, 0, 10, 208, 1, 0, 0, 0, 12,
		215, 1, 0, 0, 0, 14, 248, 1, 0, 0, 0, 16, 250, 1, 0, 0, 0, 18, 257, 1,
		0, 0, 0, 20, 265, 1, 0, 0, 0, 22, 327, 1, 0, 0, 0, 24, 348, 1, 0, 0, 0,
		26, 386, 1, 0, 0, 0, 28, 402, 1, 0, 0, 0, 30, 425, 1, 0, 0, 0, 32, 442,
		1, 0, 0, 0, 34, 452, 1, 0, 0, 0, 36, 464, 1, 0, 0, 0, 38, 518, 1, 0, 0,
		0, 40, 578, 1, 0, 0, 0, 42, 590, 1, 0, 0, 0, 44, 592, 1, 0, 0, 0, 46, 606,
		1, 0, 0, 0, 48, 623, 1, 0, 0, 0, 50, 640, 1, 0, 0, 0, 52, 651, 1, 0, 0,
		0, 54, 670, 1, 0, 0, 0, 56, 57, 3, 2, 1, 0, 57, 58, 5, 0, 0, 1, 58, 59,
		6, 0, -1, 0, 59, 1, 1, 0, 0, 0, 60, 62, 3, 4, 2, 0, 61, 60, 1, 0, 0, 0,
		62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1,
		0, 0, 0, 65, 66, 6, 1, -1, 0, 66, 3, 1, 0, 0, 0, 67, 69, 3, 10, 5, 0, 68,
		70, 5, 57, 0, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0,
		0, 0, 71, 72, 6, 2, -1, 0, 72, 165, 1, 0, 0, 0, 73, 74, 3, 14, 7, 0, 74,
		75, 6, 2, -1, 0, 75, 165, 1, 0, 0, 0, 76, 78, 3, 22, 11, 0, 77, 79, 5,
		57, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80,
		81, 6, 2, -1, 0, 81, 165, 1, 0, 0, 0, 82, 84, 3, 24, 12, 0, 83, 85, 5,
		57, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86,
		87, 6, 2, -1, 0, 87, 165, 1, 0, 0, 0, 88, 89, 3, 16, 8, 0, 89, 90, 6, 2,
		-1, 0, 90, 165, 1, 0, 0, 0, 91, 92, 3, 20, 10, 0, 92, 93, 6, 2, -1, 0,
		93, 165, 1, 0, 0, 0, 94, 95, 3, 18, 9, 0, 95, 96, 6, 2, -1, 0, 96, 165,
		1, 0, 0, 0, 97, 98, 3, 26, 13, 0, 98, 99, 6, 2, -1, 0, 99, 165, 1, 0, 0,
		0, 100, 101, 3, 6, 3, 0, 101, 102, 6, 2, -1, 0, 102, 165, 1, 0, 0, 0, 103,
		105, 3, 30, 15, 0, 104, 106, 5, 57, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108, 6, 2, -1, 0, 108, 165, 1, 0,
		0, 0, 109, 111, 5, 16, 0, 0, 110, 112, 5, 57, 0, 0, 111, 110, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 165, 6, 2, -1, 0, 114,
		116, 5, 17, 0, 0, 115, 117, 5, 57, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 165, 6, 2, -1, 0, 119, 120, 5, 32,
		0, 0, 120, 121, 5, 58, 0, 0, 121, 122, 5, 18, 0, 0, 122, 123, 5, 48, 0,
		0, 123, 124, 3, 38, 19, 0, 124, 126, 5, 49, 0, 0, 125, 127, 5, 57, 0, 0,
		126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 6, 2, -1, 0, 129, 165, 1, 0, 0, 0, 130, 131, 5, 32, 0, 0, 131, 132,
		5, 58, 0, 0, 132, 133, 5, 19, 0, 0, 133, 134, 5, 48, 0, 0, 134, 136, 5,
		49, 0, 0, 135, 137, 5, 57, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0,
		0, 0, 137, 138, 1, 0, 0, 0, 138, 165, 6, 2, -1, 0, 139, 140, 5, 32, 0,
		0, 140, 141, 5, 58, 0, 0, 141, 142, 5, 20, 0, 0, 142, 143, 5, 48, 0, 0,
		143, 144, 5, 21, 0, 0, 144, 145, 5, 52, 0, 0, 145, 146, 3, 38, 19, 0, 146,
		148, 5, 49, 0, 0, 147, 149, 5, 57, 0, 0, 148, 147, 1, 0, 0, 0, 148, 149,
		1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 6, 2, -1, 0, 151, 165, 1, 0,
		0, 0, 152, 153, 5, 25, 0, 0, 153, 155, 3, 38, 19, 0, 154, 156, 5, 57, 0,
		0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157,
		158, 6, 2, -1, 0, 158, 165, 1, 0, 0, 0, 159, 161, 5, 25, 0, 0, 160, 162,
		5, 57, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0,
		0, 0, 163, 165, 6, 2, -1, 0, 164, 67, 1, 0, 0, 0, 164, 73, 1, 0, 0, 0,
		164, 76, 1, 0, 0, 0, 164, 82, 1, 0, 0, 0, 164, 88, 1, 0, 0, 0, 164, 91,
		1, 0, 0, 0, 164, 94, 1, 0, 0, 0, 164, 97, 1, 0, 0, 0, 164, 100, 1, 0, 0,
		0, 164, 103, 1, 0, 0, 0, 164, 109, 1, 0, 0, 0, 164, 114, 1, 0, 0, 0, 164,
		119, 1, 0, 0, 0, 164, 130, 1, 0, 0, 0, 164, 139, 1, 0, 0, 0, 164, 152,
		1, 0, 0, 0, 164, 159, 1, 0, 0, 0, 165, 5, 1, 0, 0, 0, 166, 167, 5, 27,
		0, 0, 167, 168, 5, 32, 0, 0, 168, 169, 5, 50, 0, 0, 169, 170, 3, 8, 4,
		0, 170, 171, 5, 51, 0, 0, 171, 172, 6, 3, -1, 0, 172, 7, 1, 0, 0, 0, 173,
		174, 6, 4, -1, 0, 174, 175, 5, 13, 0, 0, 175, 176, 5, 32, 0, 0, 176, 177,
		5, 52, 0, 0, 177, 178, 3, 32, 16, 0, 178, 179, 6, 4, -1, 0, 179, 187, 1,
		0, 0, 0, 180, 181, 5, 13, 0, 0, 181, 182, 5, 32, 0, 0, 182, 183, 5, 52,
		0, 0, 183, 184, 5, 32, 0, 0, 184, 187, 6, 4, -1, 0, 185, 187, 6, 4, -1,
		0, 186, 173, 1, 0, 0, 0, 186, 180, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187,
		205, 1, 0, 0, 0, 188, 189, 10, 5, 0, 0, 189, 190, 5, 55, 0, 0, 190, 191,
		5, 13, 0, 0, 191, 192, 5, 32, 0, 0, 192, 193, 5, 52, 0, 0, 193, 194, 3,
		32, 16, 0, 194, 195, 6, 4, -1, 0, 195, 204, 1, 0, 0, 0, 196, 197, 10, 4,
		0, 0, 197, 198, 5, 55, 0, 0, 198, 199, 5, 13, 0, 0, 199, 200, 5, 32, 0,
		0, 200, 201, 5, 52, 0, 0, 201, 202, 5, 32, 0, 0, 202, 204, 6, 4, -1, 0,
		203, 188, 1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205,
		203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 9, 1, 0, 0, 0, 207, 205, 1,
		0, 0, 0, 208, 209, 5, 7, 0, 0, 209, 210, 5, 48, 0, 0, 210, 211, 3, 44,
		22, 0, 211, 212, 5, 49, 0, 0, 212, 213, 6, 5, -1, 0, 213, 11, 1, 0, 0,
		0, 214, 216, 3, 14, 7, 0, 215, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217,
		215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 220,
		6, 6, -1, 0, 220, 13, 1, 0, 0, 0, 221, 222, 5, 8, 0, 0, 222, 223, 3, 38,
		19, 0, 223, 224, 5, 50, 0, 0, 224, 225, 3, 2, 1, 0, 225, 226, 5, 51, 0,
		0, 226, 227, 6, 7, -1, 0, 227, 249, 1, 0, 0, 0, 228, 229, 5, 8, 0, 0, 229,
		230, 3, 38, 19, 0, 230, 231, 5, 50, 0, 0, 231, 232, 3, 2, 1, 0, 232, 233,
		5, 51, 0, 0, 233, 234, 5, 9, 0, 0, 234, 235, 5, 50, 0, 0, 235, 236, 3,
		2, 1, 0, 236, 237, 5, 51, 0, 0, 237, 238, 6, 7, -1, 0, 238, 249, 1, 0,
		0, 0, 239, 240, 5, 8, 0, 0, 240, 241, 3, 38, 19, 0, 241, 242, 5, 50, 0,
		0, 242, 243, 3, 2, 1, 0, 243, 244, 5, 51, 0, 0, 244, 245, 5, 9, 0, 0, 245,
		246, 3, 12, 6, 0, 246, 247, 6, 7, -1, 0, 247, 249, 1, 0, 0, 0, 248, 221,
		1, 0, 0, 0, 248, 228, 1, 0, 0, 0, 248, 239, 1, 0, 0, 0, 249, 15, 1, 0,
		0, 0, 250, 251, 5, 10, 0, 0, 251, 252, 3, 38, 19, 0, 252, 253, 5, 50, 0,
		0, 253, 254, 3, 2, 1, 0, 254, 255, 5, 51, 0, 0, 255, 256, 6, 8, -1, 0,
		256, 17, 1, 0, 0, 0, 257, 258, 5, 28, 0, 0, 258, 259, 3, 38, 19, 0, 259,
		260, 5, 9, 0, 0, 260, 261, 5, 50, 0, 0, 261, 262, 3, 2, 1, 0, 262, 263,
		5, 51, 0, 0, 263, 264, 6, 9, -1, 0, 264, 19, 1, 0, 0, 0, 265, 266, 5, 11,
		0, 0, 266, 267, 5, 32, 0, 0, 267, 268, 5, 12, 0, 0, 268, 269, 3, 36, 18,
		0, 269, 270, 5, 50, 0, 0, 270, 271, 3, 2, 1, 0, 271, 272, 5, 51, 0, 0,
		272, 273, 6, 10, -1, 0, 273, 21, 1, 0, 0, 0, 274, 275, 5, 13, 0, 0, 275,
		276, 5, 32, 0, 0, 276, 277, 5, 52, 0, 0, 277, 278, 3, 32, 16, 0, 278, 279,
		5, 38, 0, 0, 279, 280, 3, 38, 19, 0, 280, 281, 6, 11, -1, 0, 281, 328,
		1, 0, 0, 0, 282, 283, 5, 13, 0, 0, 283, 284, 5, 32, 0, 0, 284, 285, 5,
		38, 0, 0, 285, 286, 3, 38, 19, 0, 286, 287, 6, 11, -1, 0, 287, 328, 1,
		0, 0, 0, 288, 289, 5, 13, 0, 0, 289, 290, 5, 32, 0, 0, 290, 291, 5, 52,
		0, 0, 291, 292, 3, 32, 16, 0, 292, 293, 5, 56, 0, 0, 293, 294, 6, 11, -1,
		0, 294, 328, 1, 0, 0, 0, 295, 296, 5, 13, 0, 0, 296, 297, 5, 32, 0, 0,
		297, 298, 5, 52, 0, 0, 298, 299, 5, 53, 0, 0, 299, 300, 3, 32, 16, 0, 300,
		301, 5, 54, 0, 0, 301, 302, 5, 38, 0, 0, 302, 303, 3, 42, 21, 0, 303, 304,
		6, 11, -1, 0, 304, 328, 1, 0, 0, 0, 305, 306, 5, 13, 0, 0, 306, 307, 5,
		32, 0, 0, 307, 308, 5, 52, 0, 0, 308, 309, 3, 34, 17, 0, 309, 310, 5, 38,
		0, 0, 310, 311, 3, 38, 19, 0, 311, 312, 6, 11, -1, 0, 312, 328, 1, 0, 0,
		0, 313, 314, 5, 14, 0, 0, 314, 315, 5, 32, 0, 0, 315, 316, 5, 52, 0, 0,
		316, 317, 3, 32, 16, 0, 317, 318, 5, 38, 0, 0, 318, 319, 3, 38, 19, 0,
		319, 320, 6, 11, -1, 0, 320, 328, 1, 0, 0, 0, 321, 322, 5, 14, 0, 0, 322,
		323, 5, 32, 0, 0, 323, 324, 5, 38, 0, 0, 324, 325, 3, 38, 19, 0, 325, 326,
		6, 11, -1, 0, 326, 328, 1, 0, 0, 0, 327, 274, 1, 0, 0, 0, 327, 282, 1,
		0, 0, 0, 327, 288, 1, 0, 0, 0, 327, 295, 1, 0, 0, 0, 327, 305, 1, 0, 0,
		0, 327, 313, 1, 0, 0, 0, 327, 321, 1, 0, 0, 0, 328, 23, 1, 0, 0, 0, 329,
		330, 5, 32, 0, 0, 330, 331, 5, 38, 0, 0, 331, 332, 3, 38, 19, 0, 332, 333,
		6, 12, -1, 0, 333, 349, 1, 0, 0, 0, 334, 335, 5, 32, 0, 0, 335, 336, 5,
		53, 0, 0, 336, 337, 3, 38, 19, 0, 337, 338, 5, 54, 0, 0, 338, 339, 5, 38,
		0, 0, 339, 340, 3, 38, 19, 0, 340, 341, 6, 12, -1, 0, 341, 349, 1, 0, 0,
		0, 342, 343, 5, 32, 0, 0, 343, 344, 7, 0, 0, 0, 344, 345, 5, 38, 0, 0,
		345, 346, 3, 38, 19, 0, 346, 347, 6, 12, -1, 0, 347, 349, 1, 0, 0, 0, 348,
		329, 1, 0, 0, 0, 348, 334, 1, 0, 0, 0, 348, 342, 1, 0, 0, 0, 349, 25, 1,
		0, 0, 0, 350, 351, 5, 26, 0, 0, 351, 352, 5, 32, 0, 0, 352, 353, 5, 48,
		0, 0, 353, 354, 3, 28, 14, 0, 354, 355, 5, 49, 0, 0, 355, 356, 5, 50, 0,
		0, 356, 357, 3, 2, 1, 0, 357, 358, 5, 51, 0, 0, 358, 359, 6, 13, -1, 0,
		359, 387, 1, 0, 0, 0, 360, 361, 5, 26, 0, 0, 361, 362, 5, 32, 0, 0, 362,
		363, 5, 48, 0, 0, 363, 364, 3, 28, 14, 0, 364, 365, 5, 49, 0, 0, 365, 366,
		5, 59, 0, 0, 366, 367, 3, 32, 16, 0, 367, 368, 5, 50, 0, 0, 368, 369, 3,
		2, 1, 0, 369, 370, 5, 51, 0, 0, 370, 371, 6, 13, -1, 0, 371, 387, 1, 0,
		0, 0, 372, 373, 5, 26, 0, 0, 373, 374, 5, 32, 0, 0, 374, 375, 5, 48, 0,
		0, 375, 376, 3, 28, 14, 0, 376, 377, 5, 49, 0, 0, 377, 378, 5, 59, 0, 0,
		378, 379, 5, 53, 0, 0, 379, 380, 3, 32, 16, 0, 380, 381, 5, 54, 0, 0, 381,
		382, 5, 50, 0, 0, 382, 383, 3, 2, 1, 0, 383, 384, 5, 51, 0, 0, 384, 385,
		6, 13, -1, 0, 385, 387, 1, 0, 0, 0, 386, 350, 1, 0, 0, 0, 386, 360, 1,
		0, 0, 0, 386, 372, 1, 0, 0, 0, 387, 27, 1, 0, 0, 0, 388, 389, 6, 14, -1,
		0, 389, 390, 5, 32, 0, 0, 390, 391, 5, 52, 0, 0, 391, 392, 3, 32, 16, 0,
		392, 393, 6, 14, -1, 0, 393, 403, 1, 0, 0, 0, 394, 395, 5, 32, 0, 0, 395,
		396, 5, 52, 0, 0, 396, 397, 5, 53, 0, 0, 397, 398, 3, 32, 16, 0, 398, 399,
		5, 54, 0, 0, 399, 400, 6, 14, -1, 0, 400, 403, 1, 0, 0, 0, 401, 403, 6,
		14, -1, 0, 402, 388, 1, 0, 0, 0, 402, 394, 1, 0, 0, 0, 402, 401, 1, 0,
		0, 0, 403, 422, 1, 0, 0, 0, 404, 405, 10, 5, 0, 0, 405, 406, 5, 55, 0,
		0, 406, 407, 5, 32, 0, 0, 407, 408, 5, 52, 0, 0, 408, 409, 3, 32, 16, 0,
		409, 410, 6, 14, -1, 0, 410, 421, 1, 0, 0, 0, 411, 412, 10, 4, 0, 0, 412,
		413, 5, 55, 0, 0, 413, 414, 5, 32, 0, 0, 414, 415, 5, 52, 0, 0, 415, 416,
		5, 53, 0, 0, 416, 417, 3, 32, 16, 0, 417, 418, 5, 54, 0, 0, 418, 419, 6,
		14, -1, 0, 419, 421, 1, 0, 0, 0, 420, 404, 1, 0, 0, 0, 420, 411, 1, 0,
		0, 0, 421, 424, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0,
		423, 29, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 425, 426, 5, 32, 0, 0, 426,
		427, 5, 48, 0, 0, 427, 428, 3, 52, 26, 0, 428, 429, 5, 49, 0, 0, 429, 430,
		6, 15, -1, 0, 430, 31, 1, 0, 0, 0, 431, 432, 5, 1, 0, 0, 432, 443, 6, 16,
		-1, 0, 433, 434, 5, 2, 0, 0, 434, 443, 6, 16, -1, 0, 435, 436, 5, 4, 0,
		0, 436, 443, 6, 16, -1, 0, 437, 438, 5, 3, 0, 0, 438, 443, 6, 16, -1, 0,
		439, 440, 5, 53, 0, 0, 440, 441, 5, 54, 0, 0, 441, 443, 6, 16, -1, 0, 442,
		431, 1, 0, 0, 0, 442, 433, 1, 0, 0, 0, 442, 435, 1, 0, 0, 0, 442, 437,
		1, 0, 0, 0, 442, 439, 1, 0, 0, 0, 443, 33, 1, 0, 0, 0, 444, 445, 5, 53,
		0, 0, 445, 446, 3, 34, 17, 0, 446, 447, 5, 54, 0, 0, 447, 448, 6, 17, -1,
		0, 448, 453, 1, 0, 0, 0, 449, 450, 3, 32, 16, 0, 450, 451, 6, 17, -1, 0,
		451, 453, 1, 0, 0, 0, 452, 444, 1, 0, 0, 0, 452, 449, 1, 0, 0, 0, 453,
		35, 1, 0, 0, 0, 454, 455, 3, 38, 19, 0, 455, 456, 5, 58, 0, 0, 456, 457,
		5, 58, 0, 0, 457, 458, 5, 58, 0, 0, 458, 459, 3, 38, 19, 0, 459, 460, 6,
		18, -1, 0, 460, 465, 1, 0, 0, 0, 461, 462, 3, 38, 19, 0, 462, 463, 6, 18,
		-1, 0, 463, 465, 1, 0, 0, 0, 464, 454, 1, 0, 0, 0, 464, 461, 1, 0, 0, 0,
		465, 37, 1, 0, 0, 0, 466, 467, 6, 19, -1, 0, 467, 468, 5, 46, 0, 0, 468,
		469, 3, 38, 19, 22, 469, 470, 6, 19, -1, 0, 470, 519, 1, 0, 0, 0, 471,
		472, 5, 35, 0, 0, 472, 473, 3, 38, 19, 16, 473, 474, 6, 19, -1, 0, 474,
		519, 1, 0, 0, 0, 475, 476, 5, 32, 0, 0, 476, 477, 5, 48, 0, 0, 477, 478,
		3, 54, 27, 0, 478, 479, 5, 49, 0, 0, 479, 480, 6, 19, -1, 0, 480, 519,
		1, 0, 0, 0, 481, 482, 3, 50, 25, 0, 482, 483, 6, 19, -1, 0, 483, 519, 1,
		0, 0, 0, 484, 485, 5, 48, 0, 0, 485, 486, 3, 38, 19, 0, 486, 487, 5, 49,
		0, 0, 487, 488, 6, 19, -1, 0, 488, 519, 1, 0, 0, 0, 489, 490, 3, 40, 20,
		0, 490, 491, 6, 19, -1, 0, 491, 519, 1, 0, 0, 0, 492, 493, 5, 32, 0, 0,
		493, 494, 5, 58, 0, 0, 494, 495, 5, 23, 0, 0, 495, 519, 6, 19, -1, 0, 496,
		497, 5, 32, 0, 0, 497, 498, 5, 58, 0, 0, 498, 499, 5, 22, 0, 0, 499, 519,
		6, 19, -1, 0, 500, 501, 3, 46, 23, 0, 501, 502, 6, 19, -1, 0, 502, 519,
		1, 0, 0, 0, 503, 504, 5, 53, 0, 0, 504, 505, 3, 44, 22, 0, 505, 506, 5,
		54, 0, 0, 506, 507, 6, 19, -1, 0, 507, 519, 1, 0, 0, 0, 508, 509, 5, 29,
		0, 0, 509, 519, 6, 19, -1, 0, 510, 511, 5, 30, 0, 0, 511, 519, 6, 19, -1,
		0, 512, 513, 5, 5, 0, 0, 513, 519, 6, 19, -1, 0, 514, 515, 5, 6, 0, 0,
		515, 519, 6, 19, -1, 0, 516, 517, 5, 15, 0, 0, 517, 519, 6, 19, -1, 0,
		518, 466, 1, 0, 0, 0, 518, 471, 1, 0, 0, 0, 518, 475, 1, 0, 0, 0, 518,
		481, 1, 0, 0, 0, 518, 484, 1, 0, 0, 0, 518, 489, 1, 0, 0, 0, 518, 492,
		1, 0, 0, 0, 518, 496, 1, 0, 0, 0, 518, 500, 1, 0, 0, 0, 518, 503, 1, 0,
		0, 0, 518, 508, 1, 0, 0, 0, 518, 510, 1, 0, 0, 0, 518, 512, 1, 0, 0, 0,
		518, 514, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 519, 557, 1, 0, 0, 0, 520,
		521, 10, 21, 0, 0, 521, 522, 7, 1, 0, 0, 522, 523, 3, 38, 19, 22, 523,
		524, 6, 19, -1, 0, 524, 556, 1, 0, 0, 0, 525, 526, 10, 20, 0, 0, 526, 527,
		7, 0, 0, 0, 527, 528, 3, 38, 19, 21, 528, 529, 6, 19, -1, 0, 529, 556,
		1, 0, 0, 0, 530, 531, 10, 19, 0, 0, 531, 532, 7, 2, 0, 0, 532, 533, 3,
		38, 19, 20, 533, 534, 6, 19, -1, 0, 534, 556, 1, 0, 0, 0, 535, 536, 10,
		18, 0, 0, 536, 537, 7, 3, 0, 0, 537, 538, 3, 38, 19, 19, 538, 539, 6, 19,
		-1, 0, 539, 556, 1, 0, 0, 0, 540, 541, 10, 17, 0, 0, 541, 542, 7, 4, 0,
		0, 542, 543, 3, 38, 19, 18, 543, 544, 6, 19, -1, 0, 544, 556, 1, 0, 0,
		0, 545, 546, 10, 15, 0, 0, 546, 547, 5, 37, 0, 0, 547, 548, 3, 38, 19,
		16, 548, 549, 6, 19, -1, 0, 549, 556, 1, 0, 0, 0, 550, 551, 10, 14, 0,
		0, 551, 552, 5, 36, 0, 0, 552, 553, 3, 38, 19, 15, 553, 554, 6, 19, -1,
		0, 554, 556, 1, 0, 0, 0, 555, 520, 1, 0, 0, 0, 555, 525, 1, 0, 0, 0, 555,
		530, 1, 0, 0, 0, 555, 535, 1, 0, 0, 0, 555, 540, 1, 0, 0, 0, 555, 545,
		1, 0, 0, 0, 555, 550, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0,
		0, 0, 557, 558, 1, 0, 0, 0, 558, 39, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0,
		560, 561, 5, 1, 0, 0, 561, 562, 5, 48, 0, 0, 562, 563, 3, 38, 19, 0, 563,
		564, 5, 49, 0, 0, 564, 565, 6, 20, -1, 0, 565, 579, 1, 0, 0, 0, 566, 567,
		5, 2, 0, 0, 567, 568, 5, 48, 0, 0, 568, 569, 3, 38, 19, 0, 569, 570, 5,
		49, 0, 0, 570, 571, 6, 20, -1, 0, 571, 579, 1, 0, 0, 0, 572, 573, 5, 4,
		0, 0, 573, 574, 5, 48, 0, 0, 574, 575, 3, 38, 19, 0, 575, 576, 5, 49, 0,
		0, 576, 577, 6, 20, -1, 0, 577, 579, 1, 0, 0, 0, 578, 560, 1, 0, 0, 0,
		578, 566, 1, 0, 0, 0, 578, 572, 1, 0, 0, 0, 579, 41, 1, 0, 0, 0, 580, 581,
		5, 53, 0, 0, 581, 582, 3, 44, 22, 0, 582, 583, 5, 54, 0, 0, 583, 584, 6,
		21, -1, 0, 584, 591, 1, 0, 0, 0, 585, 586, 5, 53, 0, 0, 586, 587, 5, 54,
		0, 0, 587, 591, 6, 21, -1, 0, 588, 589, 5, 32, 0, 0, 589, 591, 6, 21, -1,
		0, 590, 580, 1, 0, 0, 0, 590, 585, 1, 0, 0, 0, 590, 588, 1, 0, 0, 0, 591,
		43, 1, 0, 0, 0, 592, 593, 6, 22, -1, 0, 593, 594, 3, 38, 19, 0, 594, 595,
		6, 22, -1, 0, 595, 603, 1, 0, 0, 0, 596, 597, 10, 2, 0, 0, 597, 598, 5,
		55, 0, 0, 598, 599, 3, 38, 19, 0, 599, 600, 6, 22, -1, 0, 600, 602, 1,
		0, 0, 0, 601, 596, 1, 0, 0, 0, 602, 605, 1, 0, 0, 0, 603, 601, 1, 0, 0,
		0, 603, 604, 1, 0, 0, 0, 604, 45, 1, 0, 0, 0, 605, 603, 1, 0, 0, 0, 606,
		607, 6, 23, -1, 0, 607, 608, 5, 32, 0, 0, 608, 609, 6, 23, -1, 0, 609,
		620, 1, 0, 0, 0, 610, 611, 10, 3, 0, 0, 611, 612, 3, 48, 24, 0, 612, 613,
		6, 23, -1, 0, 613, 619, 1, 0, 0, 0, 614, 615, 10, 2, 0, 0, 615, 616, 5,
		58, 0, 0, 616, 617, 5, 32, 0, 0, 617, 619, 6, 23, -1, 0, 618, 610, 1, 0,
		0, 0, 618, 614, 1, 0, 0, 0, 619, 622, 1, 0, 0, 0, 620, 618, 1, 0, 0, 0,
		620, 621, 1, 0, 0, 0, 621, 47, 1, 0, 0, 0, 622, 620, 1, 0, 0, 0, 623, 624,
		6, 24, -1, 0, 624, 625, 5, 53, 0, 0, 625, 626, 3, 38, 19, 0, 626, 627,
		5, 54, 0, 0, 627, 628, 6, 24, -1, 0, 628, 637, 1, 0, 0, 0, 629, 630, 10,
		2, 0, 0, 630, 631, 5, 53, 0, 0, 631, 632, 3, 38, 19, 0, 632, 633, 5, 54,
		0, 0, 633, 634, 6, 24, -1, 0, 634, 636, 1, 0, 0, 0, 635, 629, 1, 0, 0,
		0, 636, 639, 1, 0, 0, 0, 637, 635, 1, 0, 0, 0, 637, 638, 1, 0, 0, 0, 638,
		49, 1, 0, 0, 0, 639, 637, 1, 0, 0, 0, 640, 641, 5, 32, 0, 0, 641, 642,
		5, 48, 0, 0, 642, 643, 3, 52, 26, 0, 643, 644, 5, 49, 0, 0, 644, 645, 6,
		25, -1, 0, 645, 51, 1, 0, 0, 0, 646, 647, 6, 26, -1, 0, 647, 648, 3, 38,
		19, 0, 648, 649, 6, 26, -1, 0, 649, 652, 1, 0, 0, 0, 650, 652, 6, 26, -1,
		0, 651, 646, 1, 0, 0, 0, 651, 650, 1, 0, 0, 0, 652, 660, 1, 0, 0, 0, 653,
		654, 10, 3, 0, 0, 654, 655, 5, 55, 0, 0, 655, 656, 3, 38, 19, 0, 656, 657,
		6, 26, -1, 0, 657, 659, 1, 0, 0, 0, 658, 653, 1, 0, 0, 0, 659, 662, 1,
		0, 0, 0, 660, 658, 1, 0, 0, 0, 660, 661, 1, 0, 0, 0, 661, 53, 1, 0, 0,
		0, 662, 660, 1, 0, 0, 0, 663, 664, 6, 27, -1, 0, 664, 665, 5, 32, 0, 0,
		665, 666, 5, 52, 0, 0, 666, 667, 3, 38, 19, 0, 667, 668, 6, 27, -1, 0,
		668, 671, 1, 0, 0, 0, 669, 671, 6, 27, -1, 0, 670, 663, 1, 0, 0, 0, 670,
		669, 1, 0, 0, 0, 671, 683, 1, 0, 0, 0, 672, 674, 10, 3, 0, 0, 673, 675,
		5, 55, 0, 0, 674, 673, 1, 0, 0, 0, 674, 675, 1, 0, 0, 0, 675, 676, 1, 0,
		0, 0, 676, 677, 5, 32, 0, 0, 677, 678, 5, 52, 0, 0, 678, 679, 3, 38, 19,
		0, 679, 680, 6, 27, -1, 0, 680, 682, 1, 0, 0, 0, 681, 672, 1, 0, 0, 0,
		682, 685, 1, 0, 0, 0, 683, 681, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684,
		55, 1, 0, 0, 0, 685, 683, 1, 0, 0, 0, 41, 63, 69, 78, 84, 105, 111, 116,
		126, 136, 148, 155, 161, 164, 186, 203, 205, 217, 248, 327, 348, 386, 402,
		420, 422, 442, 452, 464, 518, 555, 557, 578, 590, 603, 618, 620, 637, 651,
		660, 670, 674, 683,
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
	SwiftGrammarParserVAR            = 13
	SwiftGrammarParserLET            = 14
	SwiftGrammarParserNIL            = 15
	SwiftGrammarParserBREAK          = 16
	SwiftGrammarParserCONTINUE       = 17
	SwiftGrammarParserAPPEND         = 18
	SwiftGrammarParserREMOVELAST     = 19
	SwiftGrammarParserREMOVE         = 20
	SwiftGrammarParserAT             = 21
	SwiftGrammarParserISEMPTY        = 22
	SwiftGrammarParserCOUNT          = 23
	SwiftGrammarParserARRAY          = 24
	SwiftGrammarParserRETURN         = 25
	SwiftGrammarParserFUNC           = 26
	SwiftGrammarParserSTRUCT         = 27
	SwiftGrammarParserGUARD          = 28
	SwiftGrammarParserNUMBER         = 29
	SwiftGrammarParserSTRING         = 30
	SwiftGrammarParserCHAR           = 31
	SwiftGrammarParserID             = 32
	SwiftGrammarParserDIFE           = 33
	SwiftGrammarParserIG_IG          = 34
	SwiftGrammarParserNOT            = 35
	SwiftGrammarParserOR             = 36
	SwiftGrammarParserAND            = 37
	SwiftGrammarParserIGUAL          = 38
	SwiftGrammarParserMAYIG          = 39
	SwiftGrammarParserMENIG          = 40
	SwiftGrammarParserMAYOR          = 41
	SwiftGrammarParserMENOR          = 42
	SwiftGrammarParserMULT           = 43
	SwiftGrammarParserDIV            = 44
	SwiftGrammarParserSUM            = 45
	SwiftGrammarParserRES            = 46
	SwiftGrammarParserMOD            = 47
	SwiftGrammarParserPAR_IZQ        = 48
	SwiftGrammarParserPAR_DER        = 49
	SwiftGrammarParserLLAVE_IZQ      = 50
	SwiftGrammarParserLLAVE_DER      = 51
	SwiftGrammarParserDOSPUNTOS      = 52
	SwiftGrammarParserCOR_IZQ        = 53
	SwiftGrammarParserCOR_DER        = 54
	SwiftGrammarParserCOMA           = 55
	SwiftGrammarParserCIERRAPREGUNTA = 56
	SwiftGrammarParserPUNTOCOMA      = 57
	SwiftGrammarParserPUNTO          = 58
	SwiftGrammarParserFLECHA         = 59
	SwiftGrammarParserWHITESPACE     = 60
	SwiftGrammarParserCOMMENT        = 61
	SwiftGrammarParserLINE_COMMENT   = 62
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
	SwiftGrammarParserRULE_forstmt         = 10
	SwiftGrammarParserRULE_declarationstmt = 11
	SwiftGrammarParserRULE_asignationstmt  = 12
	SwiftGrammarParserRULE_function        = 13
	SwiftGrammarParserRULE_listParamsFunc  = 14
	SwiftGrammarParserRULE_callFuncionIns  = 15
	SwiftGrammarParserRULE_types           = 16
	SwiftGrammarParserRULE_typesmatriz     = 17
	SwiftGrammarParserRULE_exprFor         = 18
	SwiftGrammarParserRULE_expr            = 19
	SwiftGrammarParserRULE_conversionstmt  = 20
	SwiftGrammarParserRULE_exprvector      = 21
	SwiftGrammarParserRULE_listParams      = 22
	SwiftGrammarParserRULE_listArray       = 23
	SwiftGrammarParserRULE_listAccessArray = 24
	SwiftGrammarParserRULE_callFuncion     = 25
	SwiftGrammarParserRULE_listParamsCall  = 26
	SwiftGrammarParserRULE_listStructExp   = 27
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
		p.SetState(56)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(57)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4798508416) != 0) {
		{
			p.SetState(60)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(63)
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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(68)
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
			p.SetState(73)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(77)
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
			p.SetState(82)

			var _x = p.Asignationstmt()

			localctx.(*InstructionContext)._asignationstmt = _x
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(83)
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
			p.SetState(88)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhileinst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetForinst()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(94)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGd()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(97)

			var _x = p.Function()

			localctx.(*InstructionContext)._function = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_function().GetFun()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(103)

			var _x = p.CallFuncionIns()

			localctx.(*InstructionContext)._callFuncionIns = _x
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(104)
				p.Match(SwiftGrammarParserPUNTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_callFuncionIns().GetCf()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(109)

			var _m = p.Match(SwiftGrammarParserBREAK)

			localctx.(*InstructionContext)._BREAK = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(110)
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

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(114)

			var _m = p.Match(SwiftGrammarParserCONTINUE)

			localctx.(*InstructionContext)._CONTINUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(119)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*InstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(SwiftGrammarParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		{
			p.SetState(124)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(125)
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
			p.Match(SwiftGrammarParserREMOVELAST)
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
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(135)
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

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(139)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*InstructionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(SwiftGrammarParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(SwiftGrammarParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		{
			p.SetState(146)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(147)
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

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(152)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*InstructionContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)

			var _x = p.expr(0)

			localctx.(*InstructionContext)._expr = _x
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(154)
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

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(159)

			var _m = p.Match(SwiftGrammarParserRETURN)

			localctx.(*InstructionContext)._RETURN = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPUNTOCOMA {
			{
				p.SetState(160)
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
		p.SetState(166)

		var _m = p.Match(SwiftGrammarParserSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructCreationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(170)
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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(174)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)

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
			p.SetState(180)
			p.Match(SwiftGrammarParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext).idp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)

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
	p.SetState(205)
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
			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(189)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(190)
					p.Match(SwiftGrammarParserVAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(191)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(192)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(193)

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
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(197)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(198)
					p.Match(SwiftGrammarParserVAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(199)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).idp = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(200)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(201)

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
		p.SetState(207)
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
		p.SetState(208)

		var _m = p.Match(SwiftGrammarParserPRINT)

		localctx.(*PrintstmtContext)._PRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)

		var _x = p.listParams(0)

		localctx.(*PrintstmtContext)._listParams = _x
	}
	{
		p.SetState(211)
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
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(214)

				var _x = p.Ifstmt()

				localctx.(*BlockelsifContext)._ifstmt = _x
			}
			localctx.(*BlockelsifContext).elseif = append(localctx.(*BlockelsifContext).elseif, localctx.(*BlockelsifContext)._ifstmt)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(217)
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
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(223)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(225)
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
			p.SetState(228)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(230)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)

			var _x = p.Block()

			localctx.(*IfstmtContext).ifblck = _x
		}
		{
			p.SetState(232)
			p.Match(SwiftGrammarParserLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

			localctx.(*IfstmtContext).elseblck = _x
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
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetIfblck().GetBlk(), localctx.(*IfstmtContext).GetElseblck().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
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
		p.SetState(250)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
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

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(254)
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
		p.SetState(257)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(259)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(262)
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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)

		var _m = p.Match(SwiftGrammarParserFOR)

		localctx.(*ForstmtContext)._FOR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ForstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(SwiftGrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)

		var _x = p.ExprFor()

		localctx.(*ForstmtContext)._exprFor = _x
	}
	{
		p.SetState(269)
		p.Match(SwiftGrammarParserLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)

		var _x = p.Block()

		localctx.(*ForstmtContext)._block = _x
	}
	{
		p.SetState(271)
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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(278)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

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
			p.SetState(282)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)

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
			p.SetState(288)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
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

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(292)
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
			p.SetState(295)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(300)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)

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
			p.SetState(305)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)

			var _x = p.Typesmatriz()

			localctx.(*DeclarationstmtContext)._typesmatriz = _x
		}
		{
			p.SetState(309)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)

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
			p.SetState(313)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(317)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)

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

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(321)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)

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
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_asignationstmt)
	var _la int

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
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
			p.SetState(334)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)

			var _x = p.expr(0)

			localctx.(*AsignationstmtContext).index = _x
		}
		{
			p.SetState(337)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)

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
			p.SetState(342)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AsignationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)

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
			p.SetState(344)
			p.Match(SwiftGrammarParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)

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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_function)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(354)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(357)
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
			p.SetState(360)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(364)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(366)

			var _x = p.Types()

			localctx.(*FunctionContext)._types = _x
		}
		{
			p.SetState(367)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(369)
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
			p.SetState(372)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FunctionContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FunctionContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)

			var _x = p.listParamsFunc(0)

			localctx.(*FunctionContext)._listParamsFunc = _x
		}
		{
			p.SetState(376)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Types()
		}
		{
			p.SetState(380)
			p.Match(SwiftGrammarParserCOR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(SwiftGrammarParserLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)

			var _x = p.Block()

			localctx.(*FunctionContext)._block = _x
		}
		{
			p.SetState(383)
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, SwiftGrammarParserRULE_listParamsFunc, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(389)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)

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
			p.SetState(394)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListParamsFuncContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _x = p.Types()

			localctx.(*ListParamsFuncContext)._types = _x
		}
		{
			p.SetState(398)
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
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
				localctx.(*ListParamsFuncContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
				p.SetState(404)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(405)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(406)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListParamsFuncContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(407)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(408)

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
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(412)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(413)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListParamsFuncContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(414)
					p.Match(SwiftGrammarParserDOSPUNTOS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(415)
					p.Match(SwiftGrammarParserCOR_IZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(416)

					var _x = p.Types()

					localctx.(*ListParamsFuncContext)._types = _x
				}
				{
					p.SetState(417)
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
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_callFuncionIns)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFuncionInsContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)

		var _x = p.listParamsCall(0)

		localctx.(*CallFuncionInsContext)._listParamsCall = _x
	}
	{
		p.SetState(428)
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
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_types)
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(431)
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
			p.SetState(433)
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
			p.SetState(435)
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
			p.SetState(437)
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
			p.SetState(439)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
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
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_typesmatriz)
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(444)
			p.Match(SwiftGrammarParserCOR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)

			var _x = p.Typesmatriz()

			localctx.(*TypesmatrizContext).list = _x
		}
		{
			p.SetState(446)
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
			p.SetState(449)

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
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_exprFor)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)

			var _x = p.expr(0)

			localctx.(*ExprForContext).range1 = _x
		}
		{
			p.SetState(455)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(458)

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
			p.SetState(461)

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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(467)

			var _m = p.Match(SwiftGrammarParserRES)

			localctx.(*ExprContext)._RES = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)

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
			p.SetState(471)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext)._NOT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)

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
			p.SetState(475)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)

			var _x = p.listStructExp(0)

			localctx.(*ExprContext)._listStructExp = _x
		}
		{
			p.SetState(478)
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
			p.SetState(481)

			var _x = p.CallFuncion()

			localctx.(*ExprContext)._callFuncion = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_callFuncion().GetCf()

	case 5:
		{
			p.SetState(484)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(486)
			p.Match(SwiftGrammarParserPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 6:
		{
			p.SetState(489)

			var _x = p.Conversionstmt()

			localctx.(*ExprContext)._conversionstmt = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_conversionstmt().GetConv()

	case 7:
		{
			p.SetState(492)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
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
			p.SetState(496)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
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
			p.SetState(500)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case 10:
		{
			p.SetState(503)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(505)
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
			p.SetState(508)

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
			p.SetState(510)

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
			p.SetState(512)

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
			p.SetState(514)

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
			p.SetState(516)

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
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(520)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(521)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&167125767421952) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(522)

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
				p.SetState(525)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(526)

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
					p.SetState(527)

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
				p.SetState(530)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(531)

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
					p.SetState(532)

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
				p.SetState(535)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(536)

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
					p.SetState(537)

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
				p.SetState(540)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(541)

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
					p.SetState(542)

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
				p.SetState(545)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(546)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(547)

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
				p.SetState(550)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(551)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(552)

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
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_conversionstmt)
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(560)

			var _m = p.Match(SwiftGrammarParserINT)

			localctx.(*ConversionstmtContext)._INT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(563)
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
			p.SetState(566)

			var _m = p.Match(SwiftGrammarParserFLOAT)

			localctx.(*ConversionstmtContext)._FLOAT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(568)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(569)
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
			p.SetState(572)

			var _m = p.Match(SwiftGrammarParserSTR)

			localctx.(*ConversionstmtContext)._STR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(573)
			p.Match(SwiftGrammarParserPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)

			var _x = p.expr(0)

			localctx.(*ConversionstmtContext)._expr = _x
		}
		{
			p.SetState(575)
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
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_exprvector)
	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(580)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprvectorContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(581)

			var _x = p.listParams(0)

			localctx.(*ExprvectorContext)._listParams = _x
		}
		{
			p.SetState(582)
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
			p.SetState(585)

			var _m = p.Match(SwiftGrammarParserCOR_IZQ)

			localctx.(*ExprvectorContext)._COR_IZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
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
			p.SetState(588)

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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(593)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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
			p.SetState(596)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(597)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(598)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(607)

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
	p.SetState(620)
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
			p.SetState(618)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(610)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(611)

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
				p.SetState(614)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(615)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(616)

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
		p.SetState(622)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, SwiftGrammarParserRULE_listAccessArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(624)
		p.Match(SwiftGrammarParserCOR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(625)

		var _x = p.expr(0)

		localctx.(*ListAccessArrayContext)._expr = _x
	}
	{
		p.SetState(626)
		p.Match(SwiftGrammarParserCOR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ListAccessArrayContext).l = []interface{}{}
	localctx.(*ListAccessArrayContext).l = append(localctx.(*ListAccessArrayContext).l, localctx.(*ListAccessArrayContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
			p.SetState(629)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(630)
				p.Match(SwiftGrammarParserCOR_IZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(631)

				var _x = p.expr(0)

				localctx.(*ListAccessArrayContext)._expr = _x
			}
			{
				p.SetState(632)
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
		p.SetState(639)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_callFuncion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFuncionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(641)
		p.Match(SwiftGrammarParserPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(642)

		var _x = p.listParamsCall(0)

		localctx.(*CallFuncionContext)._listParamsCall = _x
	}
	{
		p.SetState(643)
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
	_startState := 52
	p.EnterRecursionRule(localctx, 52, SwiftGrammarParserRULE_listParamsCall, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(651)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(647)

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
	p.SetState(660)
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
			localctx = NewListParamsCallContext(p, _parentctx, _parentState)
			localctx.(*ListParamsCallContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsCall)
			p.SetState(653)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(654)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(655)

				var _x = p.expr(0)

				localctx.(*ListParamsCallContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsCallContext).GetList().GetL(), localctx.(*ListParamsCallContext).Get_expr().GetE())
			localctx.(*ListParamsCallContext).l = arr

		}
		p.SetState(662)
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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, SwiftGrammarParserRULE_listStructExp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(664)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructExpContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(665)
			p.Match(SwiftGrammarParserDOSPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(666)

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
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
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
			p.SetState(672)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			p.SetState(674)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == SwiftGrammarParserCOMA {
				{
					p.SetState(673)
					p.Match(SwiftGrammarParserCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(676)

				var _m = p.Match(SwiftGrammarParserID)

				localctx.(*ListStructExpContext)._ID = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(677)
				p.Match(SwiftGrammarParserDOSPUNTOS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(678)

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
		p.SetState(685)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
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

	case 14:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 19:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 22:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 23:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 24:
		var t *ListAccessArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessArrayContext)
		}
		return p.ListAccessArray_Sempred(t, predIndex)

	case 26:
		var t *ListParamsCallContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsCallContext)
		}
		return p.ListParamsCall_Sempred(t, predIndex)

	case 27:
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

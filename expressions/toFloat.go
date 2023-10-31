package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	// "fmt"
	// "strconv"
	// // "math"
	// "strings"
)

type toFloat struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToFloat(lin int, col int, exp interfaces.Expression) toFloat {
	return toFloat{lin,col,exp}
}

func (p toFloat) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

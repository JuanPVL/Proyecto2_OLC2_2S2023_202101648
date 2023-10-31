package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "fmt"
	// "strconv"
	// "math"
	// "strings"
)

type toInt struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToInt(lin int, col int, exp interfaces.Expression) toInt {
	return toInt{lin,col,exp}
}

func (p toInt) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
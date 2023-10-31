package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "fmt"
	// "strconv"
	// "math"
	//"strings"
)

type toString struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToString(lin int, col int, exp interfaces.Expression) toString {
	return toString{lin,col,exp}
}

func (p toString) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
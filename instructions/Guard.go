package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"strconv"
)

type Guard struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
}

func NewGuard(lin int, col int, condicion interfaces.Expression, bloque []interface{}) Guard {
	return Guard{lin,col,condicion,bloque}
}

func (g Guard) Ejecutar(ast *environment.AST,env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
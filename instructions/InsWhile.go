package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type While struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque []interface{}) While {
	whileInstr := While{lin, col, condition, bloque}
	return whileInstr
}

func (p While) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
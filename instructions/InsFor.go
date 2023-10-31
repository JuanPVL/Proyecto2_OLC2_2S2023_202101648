package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type For struct {
	Lin 		int
	Col 		int
	Id 			string
	Rango 		interfaces.Expression
	Bloque 		[]interface{}
}

func NewFor(lin int, col int, id string, rango interfaces.Expression, bloque []interface{}) For {
	return For{lin,col,id,rango,bloque}
}

func (p For) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
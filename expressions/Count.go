package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type Count struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewCount(lin int, col int, id string) Count {
	exp := Count{lin,col,id}
	return exp
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
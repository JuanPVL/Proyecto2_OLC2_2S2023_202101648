package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type IsEmpty struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewIsEmpty(lin int, col int, id string) IsEmpty {
	exp := IsEmpty{lin,col,id}
	return exp
}

func (p IsEmpty) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
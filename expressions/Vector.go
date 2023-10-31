package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "Proyecto2_OLC2_2S2023_202101648/interfaces"
	// "strconv"
)

type Vector struct {
	Lin 		int
	Col 		int
	ListExp		[]interface{}
}

func NewVector(lin int, col int, list []interface{}) Vector {
	exp := Vector{lin,col,list}
	return exp
}

func (p Vector) Ejecutar(ast *environment.AST,env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

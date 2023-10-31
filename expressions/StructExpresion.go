package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "Proyecto2_OLC2_2S2023_202101648/interfaces"
	// "strconv"
)

type StructExp struct {
	Lin     int
	Col     int
	Id      string
	ListExp []interface{}
}

func NewStructExp(lin int, col int, id string, list []interface{}) StructExp {
	exp := StructExp{lin, col, id, list}
	return exp
}

func (p StructExp) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
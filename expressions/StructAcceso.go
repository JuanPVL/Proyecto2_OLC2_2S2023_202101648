package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"strconv"
)

type StructAccess struct {
	Lin    int
	Col    int
	Struct interfaces.Expression
	Id     string
}

func NewStructAccess(lin int, col int, str interfaces.Expression, id string) StructAccess {
	exp := StructAccess{lin, col, str, id}
	return exp
}

func (p StructAccess) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var result environment.Value
	return result
}
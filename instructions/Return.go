package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	//"fmt"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
)

type Return struct {
	Lin 		int
	Col 		int
	Expression 	interfaces.Expression
}

func NewReturn(lin int, col int, exp interfaces.Expression) Return {
	return Return{lin,col,exp}
}

func (p Return) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type Append struct {
	Lin 		int
	Col 		int
	Id 			string
	Expresion 	interfaces.Expression
}

func NewAppend(lin int, col int,id string, val interfaces.Expression) Append{
	instr := Append{lin, col,id, val}
	return instr
}

func (va Append) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
	//"fmt"
	//"strconv"
)

type RemoveLast struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewRemoveLast(lin int, col int,id string) RemoveLast{
	instr := RemoveLast{lin, col,id}
	return instr
}

func (va RemoveLast) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

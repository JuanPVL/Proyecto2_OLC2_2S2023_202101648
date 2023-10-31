package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"strconv"
)

type DeclaracionVector struct {
	Lin 		int
	Col 		int
	Id 			string
	Mutable		bool
	Tipo 		environment.TipoExpresion
	Expresion 	interfaces.Expression
}

func NewDeclaracionVector(lin int, col int,id string, mut bool, tipo environment.TipoExpresion, val interfaces.Expression) DeclaracionVector{
	instr := DeclaracionVector{lin, col,id,mut, tipo, val}
	return instr
}

func (va DeclaracionVector) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
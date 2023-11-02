package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"Proyecto2_OLC2_2S2023_202101648/interfaces"
)

type Default struct {
	Lin    int
	Col    int
	Bloque []interface{}
}

func NewDefault(lin int, col int, bloque []interface{}) Default {
	defaultInstr := Default{lin, col, bloque}
	return defaultInstr
}

func (p Default) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

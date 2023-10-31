package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
)

type Break struct {
	Lin 		int
	Col 		int
}

func NewBreak(lin int, col int) Break {
	return Break{lin,col}
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value{
	var resultado environment.Value
	
	resultado = environment.NewValue("",false,environment.NULL)

	return resultado
}

package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
)

type Continue struct {
	Lin 		int
	Col 		int
}

func NewContinue(lin int, col int) Continue {
	return Continue{lin,col}
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value{
	var resultado environment.Value

	resultado = environment.NewValue("",false,environment.NULL,false,false,true)
	return resultado
}

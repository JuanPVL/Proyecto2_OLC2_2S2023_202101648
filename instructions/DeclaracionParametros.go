package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
)

type DeclaracionParametros struct {
	Lin 		int
	Col 		int
	Id 			string
	Tipo 		environment.TipoExpresion
}

func NewDeclaracionParametros(lin int, col int, id string, tipo environment.TipoExpresion) DeclaracionParametros {
	return DeclaracionParametros{lin,col,id,tipo}
}

func(p DeclaracionParametros) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	result = environment.NewValue(p.Id, false, p.Tipo, false, false, false)
	return result
}
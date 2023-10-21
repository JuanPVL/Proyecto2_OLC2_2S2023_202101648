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

func(p DeclaracionParametros) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var resultado environment.Symbol
	resultado = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: p.Tipo, Valor: 0, Mutable: true}
	env.(environment.Environment).SaveVariable(p.Id,resultado,ast)
	return resultado
}
package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
)

type Funcion struct {
	Lin 		int
	Col 		int
	Id			string
	ListaParametros []interface{}
	Tipo 		environment.TipoExpresion
	Bloque 		[]interface{}
}

func NewFuncion(lin int, col int, id string, listaparametros []interface{}, tipo environment.TipoExpresion, bloque []interface{}) Funcion {
	return Funcion{lin,col,id,listaparametros,tipo,bloque}
}

func (f Funcion) Ejecutar(ast *environment.AST,env interface{},gen *generator.Generator) environment.Symbol {
	var resultado environment.Symbol
	var funcion environment.FunctionSymbol

	funcion = environment.FunctionSymbol{Lin: f.Lin, Col: f.Col, Id: f.Id, ListaParametros: f.ListaParametros, Bloque: f.Bloque, TipoRetorno: f.Tipo}
	env.(environment.Environment).SaveFunction(f.Id,funcion,ast)
	return resultado
}
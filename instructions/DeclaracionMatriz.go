package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	"strconv"
)

type DeclaracionMatriz struct {
	Lin 		int
	Col 		int
	Id 			string
	Mutable		bool
	Tipo 		[]interface{}
	Expresion 	interfaces.Expression
}

func NewDeclaracionMatriz(linea int, columna int,id string, mut bool, tipo []interface{}, val interfaces.Expression) DeclaracionMatriz{
	instr := DeclaracionMatriz{linea, columna,id,mut, tipo, val}
	return instr
}

func (p DeclaracionMatriz) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var retorno environment.Symbol
	retorno = p.Expresion.Ejecutar(ast,env)
	if p.ValidarArray(retorno) {
		env.(environment.Environment).SaveVariable(p.Id, retorno,ast)
	} else {
		ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(p.Lin), Col: strconv.Itoa(p.Col), Descripcion: "Tipos de datos no coinciden", Ambito: env.(environment.Environment).Id})
	}
	return retorno
}

func (p DeclaracionMatriz) ValidarArray(result environment.Symbol) bool {
	return true
}
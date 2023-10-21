package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	//"fmt"
	"strconv"
)

type Count struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewCount(lin int, col int, id string) Count {
	exp := Count{lin,col,id}
	return exp
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var temporal environment.Symbol
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	temporal = env.(environment.Environment).GetVariable(p.Id,ast,linea,columna)
	if temporal.Tipo == environment.VECTOR {
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: len(temporal.Valor.([]interface{})), Mutable: true}
	} else {
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no es un vector", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: 0, Mutable: true}
	}
}
package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	//"fmt"
	"strconv"
)

type IsEmpty struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewIsEmpty(lin int, col int, id string) IsEmpty {
	exp := IsEmpty{lin,col,id}
	return exp
}

func (p IsEmpty) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var temporal environment.Symbol
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	temporal = env.(environment.Environment).GetVariable(p.Id,ast,linea,columna)
	if temporal.Tipo == environment.VECTOR {
		arr:=temporal.Valor.([]interface{})
		if len(arr)>0{
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: false, Mutable: true}
		} else {
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: true, Mutable: true}
		}
	} else {
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no es un vector", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: false, Mutable: true}
	}
}
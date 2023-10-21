package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"strconv"
)

type StructAccess struct {
	Lin    int
	Col    int
	Struct interfaces.Expression
	Id     string
}

func NewStructAccess(lin int, col int, str interfaces.Expression, id string) StructAccess {
	exp := StructAccess{lin, col, str, id}
	return exp
}

func (p StructAccess) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result, tempStruct environment.Symbol
	tempStruct = p.Struct.Ejecutar(ast, env) //obtengo el struct

	if tempStruct.Tipo == environment.STRUCT {

		if variable, ok := tempStruct.Valor.(map[string]environment.Symbol)[p.Id]; ok {
			return variable
		}
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no existe en el struct", Ambito: env.(environment.Environment).Id})
		return result
	}
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no es un struct", Ambito: env.(environment.Environment).Id})
	return result
}
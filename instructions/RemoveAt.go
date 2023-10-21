package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	"strconv"
)

type RemoveAt struct {
	Lin 		int
	Col 		int
	Id 			string
	Expresion 	interfaces.Expression
}

func NewRemoveAt(lin int, col int,id string, val interfaces.Expression) RemoveAt{
	instr := RemoveAt{lin, col,id, val}
	return instr
}

func (va RemoveAt) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var result environment.Symbol
	linea := strconv.Itoa(va.Lin)
	columna := strconv.Itoa(va.Col)
	result = env.(environment.Environment).GetVariable(va.Id,ast,linea,columna)
	if result.Tipo == environment.VECTOR{
		if result.Valor != nil{
			temparr := result.Valor.([]interface{})
			//verificar que se tengan elementos
			if len(temparr) > 0{
				index := va.Expresion.Ejecutar(ast,env)
				//verificar que el indice pueda ser accedido
				if index.Tipo == environment.INTEGER{
					//verificar que no sea nulo
					if index.Valor != nil{
						i := index.Valor.(int)
						//verificar que el indice este dentro del rango
						if i >= 0 && i < len(temparr){
							temparr = append(temparr[:i], temparr[i+1:]...)
							result.Valor = temparr
							env.(environment.Environment).SetVariable(va.Id, result,ast)
						} else {
							ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El indice esta fuera de rango", Ambito: env.(environment.Environment).Id})
						}
					} else {
						ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El indice no puede ser nulo", Ambito: env.(environment.Environment).Id})
					}
				} else {
					ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El indice no es un entero", Ambito: env.(environment.Environment).Id})
				}
			} else {
				ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El arreglo no tiene elementos", Ambito: env.(environment.Environment).Id})
			}
		} else {
			ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El arreglo no puede ser nulo", Ambito: env.(environment.Environment).Id})
		}
	}
	return result
}
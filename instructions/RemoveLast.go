package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
	//"fmt"
	"strconv"
)

type RemoveLast struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewRemoveLast(lin int, col int,id string) RemoveLast{
	instr := RemoveLast{lin, col,id}
	return instr
}

func (va RemoveLast) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var result environment.Symbol
	linea := strconv.Itoa(va.Lin)
	columna := strconv.Itoa(va.Col)
	result = env.(environment.Environment).GetVariable(va.Id,ast,linea,columna)
	if result.Tipo == environment.VECTOR{
		if result.Valor != nil{
			arr := result.Valor.([]interface{})
			if len(arr) > 0{
				arr = arr[:len(arr)-1]
				result.Valor = arr
				env.(environment.Environment).SetVariable(va.Id, result,ast)
			}
		}
	}
	return result
}

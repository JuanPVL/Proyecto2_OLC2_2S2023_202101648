package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"Proyecto1_OLC2_2S2023_202101648/instructions"
	"strconv"
	"strings"
	"fmt"
)

type LlamadoFuncionIns struct {
	Lin 		int
	Col 		int
	Id 			string
	Parametros 	[]interface{}
}

func NewLlamadoFuncion(lin int, col int, id string, parametros []interface{}) LlamadoFuncionIns {
	return LlamadoFuncionIns{lin,col,id,parametros}
}

func (lf LlamadoFuncionIns) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value{
	var result environment.Value
	linea := strconv.Itoa(lf.Lin)
	columna := strconv.Itoa(lf.Col)
	size := env.(environment.Environment).Size["size"]
	gen.AddComment("Inicio de llamada")
	if len(lf.Parametros) > 0 {
		tmp1 := gen.NewTemp()
		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(lf.Parametros); i++ {

			//ejecutando parametros
			if strings.Contains(fmt.Sprintf("%T", lf.Parametros[i]), "expressions") {
				result = lf.Parametros[i].(interfaces.Expression).Ejecutar(ast, env, gen)
			} else {
				ast.SetErrors(environment.ErrorS{Lin: linea,Col: columna,Descripcion: "Error en parametros de llamada a funcion",Ambito: "bloque"})
				fmt.Println("Error en bloque")
			}
			//guardando
			gen.AddSetStack("(int)"+tmp1, result.Value)
			if (len(lf.Parametros) - 1) != i {
				gen.AddExpression(tmp1, tmp1, "1", "+")
			}
		}

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(lf.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	} else {
		tmp1 := gen.NewTemp()

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(lf.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	}
	gen.AddComment("Final de llamada")
	return result
}

func (lf LlamadoFuncionIns) GetGlobal(env interface{}) environment.Environment {
	var tmpGlobal environment.Environment
	tmpGlobal = env.(environment.Environment)
	for {
		if tmpGlobal.Id == "GLOBAL" {
			return tmpGlobal
		}
		if tmpGlobal.Anterior == nil {
			break
		} else {
			tmpGlobal = tmpGlobal.Anterior.(environment.Environment)
		}
	}
	return tmpGlobal
}
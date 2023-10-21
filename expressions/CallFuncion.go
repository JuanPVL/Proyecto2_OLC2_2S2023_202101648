package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/instructions"
	"strconv"
)

type LlamadoFuncion struct {
	Lin 		int
	Col 		int
	Id 			string
	Parametros 	[]interface{}
}

func NewLlamadoFuncion(lin int, col int, id string, parametros []interface{}) LlamadoFuncion {
	return LlamadoFuncion{lin,col,id,parametros}
}

func (lf LlamadoFuncion) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var resultado environment.Symbol
	var funcion environment.FunctionSymbol
	linea := strconv.Itoa(lf.Lin)
	columna := strconv.Itoa(lf.Col)
	funcion = env.(environment.Environment).GetFunction(lf.Id,ast,linea,columna)
	var envFuncion environment.Environment
	envFuncion = environment.NewEnvironment(lf.GetGlobal(env),lf.Id+"(Funcion)")

	if len(lf.Parametros) == len(funcion.ListaParametros) {
		for i:=0; i < len(lf.Parametros); i++ {
			var param environment.Symbol
			param = lf.Parametros[i].(interfaces.Expression).Ejecutar(ast,env)
			if param.Tipo == funcion.ListaParametros[i].(instructions.DeclaracionParametros).Tipo {
				envFuncion.SaveVariable(funcion.ListaParametros[i].(instructions.DeclaracionParametros).Id,param,ast)
			} else {
				ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna,Descripcion: "El tipo de dato del parámetro no coincide", Ambito: envFuncion.Id})
				return resultado
			}
		}
	} else {
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna,Descripcion: "La cantidad de parámetros no coincide", Ambito: envFuncion.Id})
	}

	for _,inst := range funcion.Bloque {
		val := inst.(interfaces.Instruction).Ejecutar(ast,envFuncion)
		if val.ReturnFlag == true {
			if funcion.TipoRetorno == val.Tipo {
				return val
			} else {
				ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna,Descripcion: "El tipo de dato del retorno no coincide", Ambito: envFuncion.Id})
				return environment.Symbol{Lin: lf.Lin, Col: lf.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
			}
		} else if val.BreakFlag == true || val.Valor == "break"{
			return val
		} else if val.ContinueFlag == true || val.Valor == "continue"{
			return val
		}
	}
	return resultado
}

func (lf LlamadoFuncion) GetGlobal(env interface{}) environment.Environment {
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
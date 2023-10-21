package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"strconv"
)

type Guard struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
}

func NewGuard(lin int, col int, condicion interfaces.Expression, bloque []interface{}) Guard {
	return Guard{lin,col,condicion,bloque}
}

func (g Guard) Ejecutar(ast *environment.AST,env interface{},gen *generator.Generator) environment.Symbol {
	var condicion environment.Symbol
	condicion = g.Condicion.Ejecutar(ast,env)

	if condicion.Tipo != environment.BOOLEAN {
		linea := strconv.Itoa(g.Lin)
		columna := strconv.Itoa(g.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La condicion no es valida", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: g.Lin, Col: g.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}

	if condicion.Valor == true {
		return environment.Symbol{Lin: g.Lin, Col: g.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	} else {
		var guardEnv environment.Environment
		guardEnv = environment.NewEnvironment(env.(environment.Environment),"GUARD")
		for _, inst := range g.Bloque {
			val := inst.(interfaces.Instruction).Ejecutar(ast,guardEnv)
			if val.BreakFlag == true || val.Valor == "break"{ 
				return val
			} else if val.ContinueFlag == true || val.Valor == "continue"{
				return val
			} else if val.ReturnFlag == true {
				return val
			}
		}
		return environment.Symbol{Lin: g.Lin, Col: g.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}
}
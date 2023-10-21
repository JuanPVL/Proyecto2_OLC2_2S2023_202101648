package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"strconv"
)

type If struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
	ElseBloque 	[]interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque []interface{}, elsebloque []interface{}) If {
	ifInstr := If{lin, col, condition, bloque, elsebloque}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var condicion environment.Symbol
	condicion = p.Condicion.Ejecutar(ast,env)

	if condicion.Tipo != environment.BOOLEAN {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La condicion no es valida", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}

	if condicion.Valor == true {
		var ifEnv environment.Environment
		ifEnv = environment.NewEnvironment(env.(environment.Environment),"IF")
		for _, inst := range p.Bloque {
			val:=inst.(interfaces.Instruction).Ejecutar(ast,ifEnv)
			if val.BreakFlag == true || val.Valor == "break"{
				return val
			} else if val.ContinueFlag == true || val.Valor == "continue"{
				return val
			} else if val.ReturnFlag == true {
				return val
			}
		}
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	} else {
		var elseEnv environment.Environment
		elseEnv = environment.NewEnvironment(env.(environment.Environment),"ELSE")
		for _, inst := range p.ElseBloque {
			val:=inst.(interfaces.Instruction).Ejecutar(ast,elseEnv)
			if val.BreakFlag == true || val.Valor == "break"{
				return val
			} else if val.ContinueFlag == true || val.Valor == "continue"{
				return val
			} else if val.ReturnFlag == true {
				return val
			}
		}
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
}



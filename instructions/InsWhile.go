package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"fmt"
	"strconv"
)

type While struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque []interface{}) While {
	whileInstr := While{lin, col, condition, bloque}
	return whileInstr
}

func (p While) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var condicion environment.Symbol
	//breakBandera := false
	firstLoop:for{
		condicion = p.Condicion.Ejecutar(ast,env)

		if condicion.Tipo != environment.BOOLEAN {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
			ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La condicion no es valida", Ambito: env.(environment.Environment).Id})
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
		}
		if condicion.Valor == true {
			var whileEnv environment.Environment
			whileEnv = environment.NewEnvironment(env.(environment.Environment),"WHILE")
	
			 for _, instr := range p.Bloque {
				val := instr.(interfaces.Instruction).Ejecutar(ast,whileEnv)
				//fmt.Println("VALOR DE Val: ",val.Valor)
				if val.BreakFlag == true || val.Valor == "break"{
					fmt.Print("ENCONTRE BREAK en while")
					break firstLoop
				}else if val.ContinueFlag == true || val.Valor == "continue"{
					fmt.Print("ENCONTRE CONTINUE en while")
					continue firstLoop
				} else if val.ReturnFlag == true {
					return val
				}
			}
		} else {
			break
		}	
	}
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
}
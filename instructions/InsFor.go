package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
	"strconv"
)

type For struct {
	Lin 		int
	Col 		int
	Id 			string
	Rango 		interfaces.Expression
	Bloque 		[]interface{}
}

func NewFor(lin int, col int, id string, rango interfaces.Expression, bloque []interface{}) For {
	return For{lin,col,id,rango,bloque}
}

func (p For) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {

	var rango []interface{}

	if p.Rango.Ejecutar(ast,env).Tipo == environment.ARRAY {
		rango = p.Rango.Ejecutar(ast,env).Valor.([]interface{})
		primar:for _, valor := range rango {
			var ForEnv environment.Environment
			ForEnv = environment.NewEnvironment(env.(environment.Environment),"FOR")
			ForEnv.SaveVariable(p.Id,valor.(environment.Symbol),ast)
			
			for _, inst := range p.Bloque {
				val:=inst.(interfaces.Instruction).Ejecutar(ast,ForEnv)
				if val.BreakFlag == true || val.Valor == "break"{
					break primar
				} else if val.ContinueFlag == true || val.Valor == "continue"{
					continue primar
				} else if val.ReturnFlag == true {
					return val
				}
			}
		}
	} else if p.Rango.Ejecutar(ast,env).Tipo == environment.STRING {
		frase := p.Rango.Ejecutar(ast,env).Valor.(string)
		longitud := len(frase)
		princ:for i := 0; i < longitud; i++ {
			var ForEnv environment.Environment
			ForEnv = environment.NewEnvironment(env.(environment.Environment),"FOR")
			ForEnv.SaveVariable(p.Id,environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.STRING, Valor: string(frase[i]), Mutable: true},ast)
			for _, inst := range p.Bloque {
				val := inst.(interfaces.Instruction).Ejecutar(ast,ForEnv)
				if val.BreakFlag == true || val.Valor == "break"{
					break princ
				} else if val.ContinueFlag == true || val.Valor == "continue"{
					continue princ
				} else if val.ReturnFlag == true {
					return val
				}
			}
		}
	}else if p.Rango.Ejecutar(ast,env).Tipo == environment.VECTOR {
		rango = p.Rango.Ejecutar(ast,env).Valor.([]interface{})
		sec:for _, valor := range rango {
			var ForEnv environment.Environment
			ForEnv = environment.NewEnvironment(env.(environment.Environment),"FOR")
			ForEnv.SaveVariable(p.Id,valor.(environment.Symbol),ast)
			
			for _, inst := range p.Bloque {
				val:=inst.(interfaces.Instruction).Ejecutar(ast,ForEnv)
				if val.BreakFlag == true || val.Valor == "break"{
					break sec
				} else if val.ContinueFlag == true || val.Valor == "continue"{
					continue sec
				} else if val.ReturnFlag == true {
					return val
				}
			}
		}
	} else {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El rango no es valido", Ambito: "FOR"})
	}
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
}
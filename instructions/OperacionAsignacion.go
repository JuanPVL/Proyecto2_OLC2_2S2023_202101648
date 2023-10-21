package instructions
import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"fmt"
	"strconv"
)

type OperacionAsignacion struct {
	Lin 		int
	Col 		int
	Id 			string
	Valor 		interfaces.Expression
	Operador 	string
}

func NewOperacionAsignacion(lin int, col int, id string, val interfaces.Expression, op string) OperacionAsignacion {
	return OperacionAsignacion{lin,col,id,val,op}
}

func (p OperacionAsignacion) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var dominante environment.TipoExpresion

	tabla_dominante := [5][5]environment.TipoExpresion{
		// INT						FLOAT				STRING			BOOLEAN					NULL	
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		//FlOAT
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		//STRING
		{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL},
		//BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.BOOLEAN, environment.NULL},
		//NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	resultado := env.(environment.Environment).GetVariable(p.Id,ast,linea,columna)
	var op1 environment.Symbol
	op1 = p.Valor.Ejecutar(ast,env)
	switch p.Operador {
		case "+":
			{
				dominante = tabla_dominante[resultado.Tipo][op1.Tipo]
				if dominante == environment.INTEGER {
					resultado.Valor = resultado.Valor.(int) + op1.Valor.(int)
					resultado.Mutable = true
					env.(environment.Environment).SetVariable(p.Id,resultado,ast)
					return resultado
				} else if dominante == environment.FLOAT {
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
					resultado.Valor = val1 + val2
					resultado.Mutable = true
					resultado.Tipo = environment.FLOAT
					env.(environment.Environment).SetVariable(p.Id,resultado,ast)
					return resultado
				} else if dominante == environment.STRING {
					r1 := fmt.Sprintf("%v", resultado.Valor)
					r2 := fmt.Sprintf("%v", op1.Valor)
					resultado.Valor = r1 + r2
					resultado.Mutable = true
					resultado.Tipo = environment.STRING
					env.(environment.Environment).SetVariable(p.Id,resultado,ast)
				} else {
					linea := strconv.Itoa(p.Lin)
					columna := strconv.Itoa(p.Col)
					ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "Error de tipos en la suma", Ambito: env.(environment.Environment).Id})
					return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
				}
			}
		case "-":
			{
				dominante = tabla_dominante[resultado.Tipo][op1.Tipo]
				if dominante == environment.INTEGER {
					resultado.Valor = resultado.Valor.(int) - op1.Valor.(int)
					resultado.Mutable = true
					resultado.Tipo = environment.INTEGER
					env.(environment.Environment).SetVariable(p.Id,resultado,ast)
					return resultado
				} else if dominante == environment.FLOAT{
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", resultado.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
					resultado.Valor = val1 - val2
					resultado.Mutable = true
					resultado.Tipo = environment.FLOAT
					env.(environment.Environment).SetVariable(p.Id,resultado,ast)
					return resultado
				} else {
					linea := strconv.Itoa(p.Lin)
					columna := strconv.Itoa(p.Col)
					ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "Error de tipos en la resta", Ambito: env.(environment.Environment).Id})
					return environment.Symbol{Lin: 0, Col: 0, Tipo: environment.NULL, Valor: 0}
				}
			}
	}
	var result interface{}
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: result,Mutable: true}
}
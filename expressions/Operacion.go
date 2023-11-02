package expressions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"

	//"fmt"
	"strconv"
)

type Operacion struct {
	Lin          int
	Col          int
	Operador_izq interfaces.Expression
	Operador     string
	Operador_der interfaces.Expression
}

func NewOperation(lin int, col int, op1 interfaces.Expression, operador string, op2 interfaces.Expression) Operacion {
	exp := Operacion{Lin: lin, Col: col, Operador_izq: op1, Operador: operador, Operador_der: op2}
	return exp
}

func (o Operacion) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
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

	var op1, op2, result environment.Value

	newTemp := gen.NewTemp()

	switch o.Operador {
	case "+":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				gen.GenerateConcatString()
				gen.AddComment("Concatenacion de string")
				envSize := strconv.Itoa(env.(environment.Environment).Size["size"])
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("swift_concatString")
				gen.AddSetStack("(int)"+tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				result = environment.NewValue(tmp2, true, dominante, false, false, false)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la suma", Ambito: env.(environment.Environment).Id})
			}
		}
	case "-":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue - op2.IntValue
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la resta", Ambito: env.(environment.Environment).Id})
			}
		}
	case "*":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue * op2.IntValue
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la multiplicacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case "/":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la division", Ambito: env.(environment.Environment).Id})
			}

		}

	case "%":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, "(int)"+op1.Value, "(int)"+op2.Value, "%")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la division", Ambito: env.(environment.Environment).Id})
			}
		}
	case "UNARIO":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			if op1.Type == environment.INTEGER {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.INTEGER, false, false, false)
				return result
			} else if op1.Type == environment.FLOAT {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.FLOAT, false, false, false)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en el unario", Ambito: env.(environment.Environment).Id})
			}
		}
	case "<":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case ">":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case "<=":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case ">=":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case "==":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "==", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case "!=":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			op2 = o.Operador_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "!=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue(newTemp, true, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la comparacion", Ambito: env.(environment.Environment).Id})
			}
		}
	case "&&":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)

			for _, lvl := range op1.TrueLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.Operador_der.Ejecutar(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		}
	case "||":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)

			for _, lvl := range op1.TrueLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.Operador_der.Ejecutar(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result

		}
	case "!":
		{
			op1 = o.Operador_izq.Ejecutar(ast, env, gen)
			if op1.Type == environment.BOOLEAN {
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)

				for _, lvl := range op1.TrueLabel {
					result.FalseLabel = append(result.FalseLabel, lvl)
				}
				for _, lvl := range op1.FalseLabel {
					result.TrueLabel = append(result.TrueLabel, lvl)
				}
				return result
			} else {
				ast.SetErrors(environment.ErrorS{Lin: strconv.Itoa(o.Lin), Col: strconv.Itoa(o.Col), Descripcion: "Error de tipos en la negacion", Ambito: env.(environment.Environment).Id})
			}
		}
	}

	gen.AddBr()
	return environment.Value{}
}

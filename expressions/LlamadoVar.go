package expressions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"strconv"
)

type LlamadoVar struct {
	Lin int
	Col int
	Id  string
}

func NewLlamadoVar(lin int, col int, id string) LlamadoVar {
	exp := LlamadoVar{lin, col, id}
	return exp
}

func (p LlamadoVar) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	gen.AddComment("Llamando una variable")
	var result environment.Value
	retSym := env.(environment.Environment).GetVariable(p.Id, ast, linea, columna)
	newTemp := gen.NewTemp()
	newTemp2 := gen.NewTemp()
	if gen.MainCode {
		gen.AddGetStack(newTemp2, strconv.Itoa(retSym.Posicion))
	} else {
		gen.AddExpression(newTemp, "P", strconv.Itoa(retSym.Posicion), "+")
		gen.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	if retSym.Tipo == environment.BOOLEAN {
		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		gen.AddIf(newTemp2, "1", "==", trueLabel)
		gen.AddGoto(falseLabel)
		result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else if retSym.Tipo == environment.ARRAY {
		gen.AddExpression(newTemp2, newTemp2, strconv.Itoa(retSym.Posicion), "+")
		result = environment.NewValue(newTemp2, true, retSym.Tipo, false, false, false)
		result.ArrSize = retSym.ArrSize
	} else {
		result = environment.NewValue(newTemp2, true, retSym.Tipo, false, false, false)
		result.ArrSize = retSym.ArrSize
	}
	return result
}

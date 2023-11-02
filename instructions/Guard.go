package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"strconv"
)

type Guard struct {
	Lin       int
	Col       int
	Condicion interfaces.Expression
	Bloque    []interface{}
}

func NewGuard(lin int, col int, condicion interfaces.Expression, bloque []interface{}) Guard {
	return Guard{lin, col, condicion, bloque}
}

func (g Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result, condicion environment.Value
	gen.AddComment("Guard")
	var outLabel []interface{}
	condicion = g.Condicion.Ejecutar(ast, env, gen)

	if condicion.Type != environment.BOOLEAN {
		linea := strconv.Itoa(g.Lin)
		columna := strconv.Itoa(g.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La condicion no es valida", Ambito: env.(environment.Environment).Id})
		var result environment.Value
		return result
	}

	newLabel := gen.NewLabel()

	//guard ejecuta lo del bloque si la condicion es falsa
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	for _, s := range g.Bloque {
		resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
		if resInst.BreakFlag {
			gen.AddGoto(gen.BreakLabel)
			resInst.BreakFlag = false
		}
		if resInst.ContinueFlag {
			gen.AddGoto(gen.ContinueLabel)
			resInst.ContinueFlag = false
		}
		for _, lvl := range resInst.OutLabel {
			outLabel = append(outLabel, lvl)
		}
	}

	gen.AddGoto(newLabel)

	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	outLabel = append(outLabel, newLabel)

	copiedSlice := make([]interface{}, len(outLabel))
	for i, item := range outLabel {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	return result
}

package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"strings"
	//"fmt"
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

func (p If) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	gen.AddComment("****** Generando If")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Condicion.Ejecutar(ast, env, gen)

	if condicion.Type != environment.BOOLEAN {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La condicion no es valida", Ambito: env.(environment.Environment).Id})
		var result environment.Value
		return result
	}
	newLabel := gen.NewLabel()
	//*****************************************add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, s := range p.Bloque {
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
				OutLvls = append(OutLvls, lvl)
			}
	}
	//*****************************************add out labels
	gen.AddGoto(newLabel)
	//*****************************************add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	if len(p.ElseBloque) > 0 {
		for _, s := range p.ElseBloque {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)

			if resInst.BreakFlag {
				gen.AddGoto(gen.BreakLabel)
				resInst.BreakFlag = false
			}
			if resInst.ContinueFlag {
				gen.AddGoto(gen.ContinueLabel)
				resInst.ContinueFlag = false
			}
				//agregando etiquetas de salida
				for _, lvl := range resInst.OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
		}
	}
	OutLvls = append(OutLvls, newLabel)

	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	return result
}



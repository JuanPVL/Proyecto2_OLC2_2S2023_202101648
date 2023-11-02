package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	//"fmt"
	//"strconv"
)

type While struct {
	Lin       int
	Col       int
	Condicion interfaces.Expression
	Bloque    []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque []interface{}) While {
	whileInstr := While{lin, col, condition, bloque}
	return whileInstr
}

func (p While) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("While Loop")
	var result, condicion environment.Value

	RetL := gen.NewLabel()

	gen.AddLabel(RetL)

	condicion = p.Condicion.Ejecutar(ast, env, gen)

	gen.AddContinueLabel(RetL)
	gen.AddBreakLabel(condicion.FalseLabel[0].(string))

	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}

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
			gen.AddLabel(lvl.(string))
		}
	}

	gen.AddGoto(RetL)

	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	return result
}

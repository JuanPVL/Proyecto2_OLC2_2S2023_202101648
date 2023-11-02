package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"strconv"
	//"fmt"
	//"strconv"
)

type For struct {
	Lin    int
	Col    int
	Id     string
	Rango  interfaces.Expression
	Bloque []interface{}
}

func NewFor(lin int, col int, id string, rango interfaces.Expression, bloque []interface{}) For {
	return For{lin, col, id, rango, bloque}
}

func (p For) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	gen.AddComment("For Loop")

	ArregloTmp := p.Rango.Ejecutar(ast, env, gen)

	tmp1 := gen.NewTemp()
	tmp2 := gen.NewTemp()
	tmp3 := gen.NewTemp()
	tmp4 := gen.NewTemp()
	tmp5 := gen.NewTemp()
	tmp6 := gen.NewTemp()
	lvl1 := gen.NewLabel()
	lvl2 := gen.NewLabel()
	lvl3 := gen.NewLabel()

	gen.AddGetHeap(tmp1, "(int)"+ArregloTmp.Value)
	gen.AddExpression(tmp2, "0", "", "")
	gen.AddExpression(tmp3, tmp2, "1", "+")
	gen.AddExpression(tmp4, ArregloTmp.Value, tmp3, "+")
	gen.AddGetHeap(tmp5, "(int)"+tmp4)
	resp := env.(environment.Environment).SaveVariable(p.Id, linea, columna, environment.INTEGER, ast, true)
	gen.AddExpression(tmp6, "P", strconv.Itoa(resp.Posicion), "+")
	gen.AddSetStack("(int)"+tmp6, tmp5)
	gen.AddLabel(lvl1)
	gen.AddIf(tmp2, tmp1, "==", lvl3)

	for _, s := range p.Bloque {
		resInstr := s.(interfaces.Instruction).Ejecutar(ast, env, gen)

		if resInstr.BreakFlag {
			gen.AddGoto(gen.BreakLabel)
			resInstr.BreakFlag = false
		}

		if resInstr.ContinueFlag {
			gen.AddGoto(gen.ContinueLabel)
			resInstr.ContinueFlag = false
		}

		for _, lvl := range resInstr.OutLabel {
			gen.AddLabel(lvl.(string))
		}
	}

	gen.AddLabel(lvl2)
	gen.AddExpression(tmp2, tmp2, "1", "+")
	gen.AddExpression(tmp3, tmp2, "1", "+")
	gen.AddExpression(tmp4, ArregloTmp.Value, tmp3, "+")
	gen.AddGetHeap(tmp5, "(int)"+tmp4)
	gen.AddSetStack("(int)"+tmp6, tmp5)
	gen.AddGoto(lvl1)
	gen.AddLabel(lvl3)

	return result
}

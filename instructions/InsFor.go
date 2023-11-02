package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/expressions"
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
	Rango1 interfaces.Expression
	Rango2 interfaces.Expression
	Bloque []interface{}
}

func NewFor(lin int, col int, id string, rango1, rango2 interfaces.Expression, bloque []interface{}) For {
	return For{lin, col, id, rango1, rango2, bloque}
}

func (p For) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result, rango2 environment.Value
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)

	rango2 = p.Rango2.Ejecutar(ast, env, gen)

	respVar := NewDeclaracion(p.Lin, p.Col, p.Id, true, environment.INTEGER, p.Rango1)
	respVar.Ejecutar(ast, env, gen)

	llamar := expressions.NewLlamadoVar(p.Lin, p.Col, p.Id)

	rValue := llamar.Ejecutar(ast, env, gen)
	newValue := rValue.Value

	lvl1 := gen.NewLabel()

	gen.AddComment("For Loop")

	gen.AddLabel(lvl1)

	lvl2 := gen.NewLabel()
	lvl3 := gen.NewLabel()

	gen.AddIf(newValue, rango2.Value, "<=", lvl2)
	gen.AddGoto(lvl3)
	gen.AddLabel(lvl2)

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

	gen.AddExpression(newValue, newValue, "1", "+")

	vari := env.(environment.Environment).GetVariable(p.Id, ast, linea, columna)

	gen.AddSetStack(strconv.Itoa(vari.Posicion), newValue)

	gen.AddGoto(lvl1)
	gen.AddLabel(lvl3)

	return result
}

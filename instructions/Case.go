package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
)

type Case struct {
	Lin 		int
	Col 		int
	Condicion 	interfaces.Expression
	Bloque 		[]interface{}
}

func NewCase(lin int, col int, condition interfaces.Expression, bloque []interface{}) Case {
	caseInstr := Case{lin, col, condition, bloque}
	return caseInstr
}

func (p Case) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value

	if p.Condicion != nil {
		result = p.Condicion.Ejecutar(ast, env, gen)
		return result
	}
	return result
}

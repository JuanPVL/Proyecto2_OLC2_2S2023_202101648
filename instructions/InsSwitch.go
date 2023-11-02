package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	//"fmt"
	//"strconv"
)

type Switch struct {
	Lin       int
	Col       int
	Condicion interfaces.Expression
	Block     []interface{}
	Default   interface{}
}

func NewSwitch(lin int, col int, condition interfaces.Expression, block []interface{}, def interface{}) Switch {
	switchInstr := Switch{lin, col, condition, block, def}
	return switchInstr
}

func (p Switch) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var condicion, result environment.Value

	newLabel := gen.NewLabel()
	LEnd := gen.NewLabel()

	gen.AddComment("Switch")

	condicion = p.Condicion.Ejecutar(ast, env, gen)

	if len(p.Block) != 0 {
		gen.AddComment("Case")

		for _, s := range p.Block {
			if s.(Case).Condicion != nil {
				Clvl := gen.NewLabel()
				ResInst := s.(Case).Condicion.Ejecutar(ast, env, gen)
				gen.AddIf(condicion.Value, ResInst.Value, "==", Clvl)
			}
		}
		for _, s := range p.Block {
			if s.(Case).Condicion != nil {
				gen.ReducLabel()
			}
		}
		if p.Default != nil {
			gen.AddComment("Default")
			for _, s := range p.Default.(Default).Bloque {
				s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			}
			gen.AddGoto(LEnd)
		}
		for _, s := range p.Block {
			gen.AddComment("Instrucciones Case")
			Clvl := gen.NewLabel()
			gen.AddLabel(Clvl)
			for _, s := range s.(Case).Bloque {
				s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			}
			gen.AddGoto(newLabel)
		}
	}
	gen.AddLabel(LEnd)
	gen.AddLabel(newLabel)
	return result
}

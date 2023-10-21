package instructions	

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "strconv"
	// "fmt"
)

type AsignacionIndexVector struct {
	Lin 		int
	Col 		int
	Id 			string
	Index 	interfaces.Expression
	Expresion2 	interfaces.Expression
}

func NewAsignacionIndexVector(lin int, col int,id string, val1 interfaces.Expression,val2 interfaces.Expression) AsignacionIndexVector{
	instr := AsignacionIndexVector{lin, col,id, val1,val2}
	return instr
}

func (p AsignacionIndexVector) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}


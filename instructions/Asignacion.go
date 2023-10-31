package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"strconv"
)

type Asignacion struct {
	Lin 		int
	Col 		int
	Id 			string
	Expresion 	interfaces.Expression
}

func NewAsignacion(lin int, col int,id string, val interfaces.Expression) Asignacion{
	instr := Asignacion{lin, col,id, val}
	return instr
}

func (va Asignacion) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	linea := strconv.Itoa(va.Lin)
	columna := strconv.Itoa(va.Col)
	var result environment.Value
	gen.AddComment("Generando Asignacion")
	//verificar si existe la variable
	variable := env.(*environment.Environment).GetVariable(va.Id,ast,linea,columna)

	result = va.Expresion.Ejecutar(ast,env,gen)

	//realizando asignacion
	gen.AddSetStack(strconv.Itoa(variable.Posicion),result.Value)
	gen.AddBr()
	return result
}
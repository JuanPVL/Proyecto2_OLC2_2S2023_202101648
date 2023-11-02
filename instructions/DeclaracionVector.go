package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"fmt"
	"strconv"
)

type DeclaracionVector struct {
	Lin 		int
	Col 		int
	Id 			string
	Mutable		bool
	Tipo 		environment.TipoExpresion
	Expresion 	interfaces.Expression
}

func NewDeclaracionVector(lin int, col int,id string, mut bool, tipo environment.TipoExpresion, val interfaces.Expression) DeclaracionVector{
	instr := DeclaracionVector{lin, col,id,mut, tipo, val}
	return instr
}

func (p DeclaracionVector) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	var newVar environment.Symbol
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)

	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	if result.Type == environment.VECTOR {
		newVar = env.(environment.Environment).SaveArrayVariable(p.Id,linea,columna, environment.VECTOR, len(result.ArrValue),ast)
		fmt.Println("newVar: ", newVar)
		gen.AddComment("Iniciando la declaración de un Vector")
		p.ArrayValidation(ast, env, gen, result.ArrValue)
		gen.AddComment("Se finalizó la declaración de un Vector")
	}
	return result
}

func (p DeclaracionVector) ArrayValidation(ast *environment.AST, env interface{}, gen *generator.Generator, arr []interface{}) {
	for _, val := range arr {
		if val.(environment.Value).Type == environment.VECTOR {
			p.ArrayValidation(ast, env, gen, val.(environment.Value).ArrValue)
		} else {
			envSize := env.(environment.Environment).NewVariable()
			gen.AddSetStack(strconv.Itoa(envSize.Posicion), val.(environment.Value).Value)
			gen.AddBr()
		}
	}
}
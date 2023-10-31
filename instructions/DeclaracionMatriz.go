package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"fmt"
	"strconv"
)

type DeclaracionMatriz struct {
	Lin 		int
	Col 		int
	Id 			string
	Mutable		bool
	Tipo 		[]interface{}
	Expresion 	interfaces.Expression
}

func NewDeclaracionMatriz(linea int, columna int,id string, mut bool, tipo []interface{}, val interfaces.Expression) DeclaracionMatriz{
	instr := DeclaracionMatriz{linea, columna,id,mut, tipo, val}
	return instr
}

func (p DeclaracionMatriz) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	var newVar environment.Symbol
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)

	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	if result.Type == environment.ARRAY {
		newVar = env.(environment.Environment).SaveArrayVariable(p.Id,linea,columna, environment.INTEGER, len(result.ArrValue),ast)
		fmt.Println("newVar: ", newVar)
		gen.AddComment("Iniciando la declaración de un ARRAY")
		p.ArrayValidation(ast, env, gen, result.ArrValue)
		gen.AddComment("Se finalizó la declaración de un ARRAY")
	}
	return result
}

func (p DeclaracionMatriz) ArrayValidation(ast *environment.AST, env interface{}, gen *generator.Generator, arr []interface{}) {
	for _, val := range arr {
		if val.(environment.Value).Type == environment.ARRAY {
			p.ArrayValidation(ast, env, gen, val.(environment.Value).ArrValue)
		} else {
			envSize := env.(environment.Environment).NewVariable()
			gen.AddSetStack(strconv.Itoa(envSize.Posicion), val.(environment.Value).Value)
			gen.AddBr()
		}
	}
}
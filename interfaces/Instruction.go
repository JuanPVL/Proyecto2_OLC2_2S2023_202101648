package interfaces

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
)


type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{}
}
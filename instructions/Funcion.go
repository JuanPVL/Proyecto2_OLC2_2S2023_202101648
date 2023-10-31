package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strings"
	"strconv"
	//"Proyecto1_OLC2_2S2023_202101648/interfaces"
)

type Funcion struct {
	Lin 		int
	Col 		int
	Id			string
	ListaParametros []interface{}
	Tipo 		environment.TipoExpresion
	Bloque 		[]interface{}
}

func NewFuncion(lin int, col int, id string, listaparametros []interface{}, tipo environment.TipoExpresion, bloque []interface{}) Funcion {
	return Funcion{lin,col,id,listaparametros,tipo,bloque}
}

func (f Funcion) Ejecutar(ast *environment.AST,env interface{},gen *generator.Generator) environment.Value {
	linea := strconv.Itoa(f.Lin)
	columna := strconv.Itoa(f.Col)
	var result environment.Value
	gen.SetMainFlag(false)
	gen.AddComment("******** Funcion " + f.Id)
	gen.AddTittle(f.Id)
	//entorno
	var envFunc environment.Environment
	envFunc = environment.NewEnvironment(env.(environment.Environment), f.Id)
	envFunc.Size["size"] = envFunc.Size["size"] + 1
	//variables
	for _, s := range f.ListaParametros {
		res := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
		envFunc.SaveVariable(res.Value,linea,columna, res.Type,ast,true)
	}
	//instrucciones func
	for _, s := range f.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, envFunc, gen)
				//agregando etiquetas de salida
				for _, lvl := range resInst.OutLabel {
					gen.AddLabel(lvl.(string))
				}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, envFunc, gen)
			//agregando etiquetas de salida
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}
	gen.AddEnd()
	gen.SetMainFlag(true)
	return result
}
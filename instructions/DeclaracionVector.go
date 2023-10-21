package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	//"fmt"
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

func (va DeclaracionVector) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var retorno environment.Symbol
	expresiontmp := va.Expresion.Ejecutar(ast,env)
	if expresiontmp.Tipo == environment.VECTOR {
		for _, valor := range expresiontmp.Valor.([]interface{}) {
			if va.Tipo == environment.DEPENDIENTE {
				if len(valor.(environment.Symbol).Valor.([]interface{})) == 0 {
					continue
				}
			}
			if valor.(environment.Symbol).Tipo != va.Tipo {
				linea := strconv.Itoa(va.Lin)
				columna := strconv.Itoa(va.Col)
				ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El tipo de dato no es valido", Ambito: env.(environment.Environment).Id})
				return retorno
			}
		}
		retorno = environment.Symbol{Lin: va.Lin, Col: va.Col, Tipo: environment.VECTOR, Valor: expresiontmp.Valor, Mutable: va.Mutable, VectorTipo: va.Tipo}
	} else {
		linea := strconv.Itoa(va.Lin)
		columna := strconv.Itoa(va.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El tipo de dato no es valido", Ambito: env.(environment.Environment).Id})
		return retorno
	}
	env.(environment.Environment).SaveVariable(va.Id, retorno,ast)
	return retorno
}
package expressions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"strconv"
)

type Vector struct {
	Lin 		int
	Col 		int
	ListExp		[]interface{}
}

func NewVector(lin int, col int, list []interface{}) Vector {
	exp := Vector{lin,col,list}
	return exp
}

func (p Vector) Ejecutar(ast *environment.AST,env interface{}) environment.Symbol {
	var retorno environment.Symbol
	var tempExp []interface{}

	if len(p.ListExp) > 0 {
		tempTipo := p.ListExp[0].(interfaces.Expression).Ejecutar(ast,env).Tipo
		var tempVectorTipo environment.TipoExpresion
		for _, s := range p.ListExp {
			simboloTmp := s.(interfaces.Expression).Ejecutar(ast,env)
			if simboloTmp.Tipo == tempTipo {
				tempExp = append(tempExp,simboloTmp)
				tempVectorTipo = tempTipo
			} else {
				linea := strconv.Itoa(p.Lin)
				columna := strconv.Itoa(p.Col)
				ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "Error en tipos", Ambito: env.(environment.Environment).Id})
			}
		}
		retorno = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.VECTOR, Valor: tempExp, Mutable: true, VectorTipo: tempVectorTipo}
	} else {
		retorno = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.VECTOR, Valor: tempExp, Mutable: true}
	}
	return retorno
}

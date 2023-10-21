package instructions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	//"fmt"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
)

type Return struct {
	Lin 		int
	Col 		int
	Expression 	interfaces.Expression
}

func NewReturn(lin int, col int, exp interfaces.Expression) Return {
	return Return{lin,col,exp}
}

func (p Return) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var resultado environment.Symbol
	if p.Expression != nil {
		tmpvalor := p.Expression.Ejecutar(ast,env)
		resultado = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: tmpvalor.Tipo, Valor: tmpvalor.Valor, Mutable: true,ReturnFlag: true, VectorTipo: tmpvalor.VectorTipo}
	} else {
		resultado = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true,ReturnFlag: true}
	}
	return resultado
}

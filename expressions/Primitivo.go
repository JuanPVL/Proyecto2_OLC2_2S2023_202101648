package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
)

type Primitivo struct {
	Lin 	int
	Col 	int
	Valor 	interface{}
	Tipo 	environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitivo {
	exp := Primitivo{Lin: lin, Col: col, Valor: valor, Tipo: tipo}
	return exp
}

func (p Primitivo) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{
		Lin: p.Lin,
		Col: p.Col,
		Tipo: p.Tipo,
		Valor: p.Valor,
		Mutable: true,
	}
}
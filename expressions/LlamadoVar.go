package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"strconv"
)

type LlamadoVar struct {
	Lin 		int
	Col 		int
	Id 			string
}

func NewLlamadoVar(lin int, col int, id string) LlamadoVar {
	exp := LlamadoVar{lin,col,id}
	return exp
}

func (p LlamadoVar) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	resultado := env.(environment.Environment).GetVariable(p.Id,ast,linea,columna)
	return resultado
}
package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	//"fmt"
	"strconv"
)

type ArrayAccess struct {
	Lin 		int
	Col 		int
	Array		interfaces.Expression
	Index		interfaces.Expression
}

func NewArrayAccess(lin int, col int, array interfaces.Expression, index interfaces.Expression) ArrayAccess {
	exp := ArrayAccess{lin,col,array,index}
	return exp
}

func (p ArrayAccess) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var tempArray environment.Symbol
	tempArray = p.Array.Ejecutar(ast,env)

	if tempArray.Tipo == environment.ARRAY || tempArray.Tipo == environment.VECTOR {
		var tempIndex environment.Symbol
		tempIndex = p.Index.Ejecutar(ast,env)
		var tempValor interface{}
		tempValor = tempArray.Valor
		
		if tempIndex.Valor.(int) >= 0 && tempIndex.Valor.(int) < len(tempValor.([]interface{})) {
			valorRetorno := tempValor.([]interface{})[(tempIndex.Valor.(int))].(environment.Symbol)
			return valorRetorno
		} else {
			linea := strconv.Itoa(p.Lin)
			columna := strconv.Itoa(p.Col)
			ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "El indice no es valido", Ambito: env.(environment.Environment).Id})
			//fmt.Println("Indice: ", tempIndex.Valor.(int))
			//fmt.Println("Tamaño: ", len(tempValor.([]interface{})))
		}
	} else {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no es un arreglo", Ambito: env.(environment.Environment).Id})
	}
	return environment.Symbol{
		Lin: p.Lin,
		Col: p.Col,
		Tipo: environment.NULL,
		Valor: 0,
		Mutable: true,
	}
}
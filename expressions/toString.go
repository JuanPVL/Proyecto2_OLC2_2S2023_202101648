package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strconv"
	// "math"
	//"strings"
)

type toString struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToString(lin int, col int, exp interfaces.Expression) toString {
	return toString{lin,col,exp}
}

func (p toString) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var exp environment.Symbol
	exp = p.Expresion.Ejecutar(ast,env)

	if exp.Tipo == environment.INTEGER{
		valor := strconv.Itoa(exp.Valor.(int))
		exp.Valor = valor
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.STRING, Valor: exp.Valor, Mutable: true}
	} else if exp.Tipo == environment.FLOAT {
		valor := strconv.FormatFloat(exp.Valor.(float64), 'f', 3, 64)
		exp.Valor = valor
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.STRING, Valor: exp.Valor, Mutable: true}
	} else if exp.Tipo == environment.BOOLEAN {
		valor := fmt.Sprintf("%v", exp.Valor)
		exp.Valor = valor
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.STRING, Valor: exp.Valor, Mutable: true}
	} else {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no puede convertirse a string", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}
}
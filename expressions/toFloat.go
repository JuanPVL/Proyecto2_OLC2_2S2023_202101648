package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strconv"
	// "math"
	"strings"
)

type toFloat struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToFloat(lin int, col int, exp interfaces.Expression) toFloat {
	return toFloat{lin,col,exp}
}

func (p toFloat) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var exp environment.Symbol
	exp = p.Expresion.Ejecutar(ast,env)

	if exp.Tipo == environment.STRING {
		numero:= fmt.Sprintf("%v", exp.Valor)
		if (strings.Contains(numero,".")){
			num,err := strconv.ParseFloat(numero, 64);
			if err!= nil{
				fmt.Println(err)
				return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: num, Mutable: true}
		} else if !(strings.Contains(numero,".")) {
			num,err1 := strconv.ParseFloat(numero, 64);
			num2 := strconv.FormatFloat(num, 'f', 2, 64)
			num3,err2 := strconv.ParseFloat(num2, 64);
			//fmt.Printf("num2 :%.2f\n",num3)
			if err1!= nil && err2!=nil{
				fmt.Println(err1)
				fmt.Println(err2)
				return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: num3, Mutable: true}
		}
	} else {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no es un string", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
}

package expressions

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strconv"
	"math"
	"strings"
)

type toInt struct {
	Lin 			int
	Col 			int
	Expresion 		interfaces.Expression
}

func NewToInt(lin int, col int, exp interfaces.Expression) toInt {
	return toInt{lin,col,exp}
}

func (p toInt) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var exp environment.Symbol
	exp = p.Expresion.Ejecutar(ast,env)
	
	if exp.Tipo == environment.STRING {
		numero:= fmt.Sprintf("%v", exp.Valor)
		if (strings.Contains(numero,".")){
			num,err := strconv.ParseFloat(numero, 64);
            if err!= nil{
                fmt.Println(err)
            }
			valor := math.Trunc(num)
			valorInt := int(valor)
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: valorInt, Mutable: true}
		} else {
			valor, _ := strconv.Atoi(numero)
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: valor, Mutable: true}
		}
	} else if exp.Tipo == environment.FLOAT {
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", exp.Valor), 64)
		//fmt.Print("valor1:",val1)
		valor := math.Trunc(val1)
		//fmt.Print("valor:",valor)
		valorInt := int(valor)
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: valorInt, Mutable: true}
	} else {
		linea := strconv.Itoa(p.Lin)
		columna := strconv.Itoa(p.Col)
		ast.SetErrors(environment.ErrorS{Lin: linea, Col: columna, Descripcion: "La variable no pudo ser convertida a int", Ambito: env.(environment.Environment).Id})
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true}
	}
}
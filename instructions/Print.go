package instructions	

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"fmt"
)

type Print struct {
	Lin 		int
	Col 		int
	Valor 		[]interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin,col,val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Symbol {
	var consoleOut string
	for _, valor := range p.Valor {
		value := valor.(interfaces.Expression).Ejecutar(ast,env)
		if value.Tipo == environment.FLOAT{
			consoleOut = consoleOut + " " + fmt.Sprintf("%.4f",value.Valor)
		} else {
			consoleOut = consoleOut + " "+fmt.Sprintf("%v",value.Valor)
		}
	}
	ast.SetPrint(consoleOut + "\n")
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.NULL, Valor: nil, Mutable: true} 
}


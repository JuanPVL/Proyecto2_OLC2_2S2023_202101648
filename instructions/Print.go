package instructions

import (
	environment "Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strconv"
)

type Print struct {
	Lin   int
	Col   int
	Valor []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	for _, valor := range p.Valor {
		result = valor.(interfaces.Expression).Ejecutar(ast, env, gen)
		if result.Type == environment.INTEGER || result.Type == environment.FLOAT || result.Type == environment.ARRAY {
			gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			gen.AddPrintf("c", "32")
			gen.AddBr()
		} else if result.Type == environment.BOOLEAN {
			if result.IsTemp {
				//cuando es variable
			}
			newLabel := gen.NewLabel()
			//add labels
			for _, lvl := range result.TrueLabel {
				gen.AddLabel(lvl.(string))
			}
			gen.AddPrintf("c", "(char)116")
			gen.AddPrintf("c", "(char)114")
			gen.AddPrintf("c", "(char)117")
			gen.AddPrintf("c", "(char)101")
			gen.AddGoto(newLabel)
			//add labels
			for _, lvl := range result.FalseLabel {
				gen.AddLabel(lvl.(string))
			}
			gen.AddPrintf("c", "(char)102")
			gen.AddPrintf("c", "(char)97")
			gen.AddPrintf("c", "(char)108")
			gen.AddPrintf("c", "(char)115")
			gen.AddPrintf("c", "(char)101")
			gen.AddLabel(newLabel)
			gen.AddPrintf("c", "32")
			gen.AddBr()
		} else if result.Type == environment.STRING {
			//llamar a generar printstring
			gen.GeneratePrintString()
			//agregar codigo en el main
			newTemp1 := gen.NewTemp()
			newTemp2 := gen.NewTemp()
			size := strconv.Itoa(env.(environment.Environment).Size["size"])
			gen.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
			gen.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
			gen.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			gen.AddExpression("P", "P", size, "+")          // cambio de entorno
			gen.AddCall("swift_printString")                //Llamada
			gen.AddGetStack(newTemp2, "(int)P")             //obtencion retorno
			gen.AddExpression("P", "P", size, "-")          //regreso del entorno
			gen.AddPrintf("c", "32")                        //salto de linea
			gen.AddBr()
		} else if result.Type == environment.NULL {
			gen.AddPrintf("c", "(char)110")
			gen.AddPrintf("c", "(char)105")
			gen.AddPrintf("c", "(char)108")
			gen.AddPrintf("c", "32")
			gen.AddBr()
		}
	}
	gen.AddPrintf("c", "10")
	return environment.Value{}
}

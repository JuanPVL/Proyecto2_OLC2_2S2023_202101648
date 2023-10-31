package instructions
import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	// "fmt"
	// "strconv"
)

type OperacionAsignacion struct {
	Lin 		int
	Col 		int
	Id 			string
	Valor 		interfaces.Expression
	Operador 	string
}

func NewOperacionAsignacion(lin int, col int, id string, val interfaces.Expression, op string) OperacionAsignacion {
	return OperacionAsignacion{lin,col,id,val,op}
}

func (p OperacionAsignacion) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	// var dominante environment.TipoExpresion

	// tabla_dominante := [5][5]environment.TipoExpresion{
	// 	// INT						FLOAT				STRING			BOOLEAN					NULL	
	// 	{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
	// 	//FlOAT
	// 	{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
	// 	//STRING
	// 	{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL},
	// 	//BOOLEAN
	// 	{environment.NULL, environment.NULL, environment.NULL, environment.BOOLEAN, environment.NULL},
	// 	//NULL
	// 	{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	// }
	var result environment.Value
	return result
}
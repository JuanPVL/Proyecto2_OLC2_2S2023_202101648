package instructions	

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"strconv"
	//"fmt"
)

type Declaracion struct {
	Lin 		int
	Col 		int
	Id 			string
	Mutable		bool
	Tipo 		environment.TipoExpresion
	Expresion 	interfaces.Expression
}

func NewDeclaracion(lin int, col int,id string, mut bool, tipo environment.TipoExpresion, val interfaces.Expression) Declaracion{
	instr := Declaracion{lin, col,id,mut, tipo, val}
	return instr
}

func (p Declaracion) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	linea := strconv.Itoa(p.Lin)
	columna := strconv.Itoa(p.Col)
	var result environment.Value
	var newVar environment.Symbol
	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	if result.Type == environment.BOOLEAN {
		newVar = env.(environment.Environment).SaveVariable(p.Id,linea,columna, p.Tipo,ast, p.Mutable)
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else if result.Type == environment.ARRAY {
		newVar = env.(environment.Environment).SaveArrayVariable(p.Id,linea,columna, environment.INTEGER, len(result.ArrValue),ast)
		//fmt.Println("newVar: ", newVar)
		gen.AddComment("Iniciando la declaración de un ARRAY")
		p.ArrayValidation(ast, env, gen, result.ArrValue)
		gen.AddComment("Se finalizó la declaración de un ARRAY")
	} else {
		//si es temp (num,string,etc)
		newVar = env.(environment.Environment).SaveVariable(p.Id,linea,columna, p.Tipo,ast, p.Mutable)
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
		gen.AddBr()
	}

	return result
}

func (p Declaracion) ArrayValidation(ast *environment.AST, env interface{}, gen *generator.Generator, arr []interface{}) {
	for _, val := range arr {
		if val.(environment.Value).Type == environment.ARRAY {
			p.ArrayValidation(ast, env, gen, val.(environment.Value).ArrValue)
		} else {
			envSize := env.(environment.Environment).NewVariable()
			gen.AddSetStack(strconv.Itoa(envSize.Posicion), val.(environment.Value).Value)
			gen.AddBr()
		}
	}
}
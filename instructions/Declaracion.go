package instructions	

import(
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"strconv"
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

func (va Declaracion) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) interface{} {

	var result environment.Value
	var newVar environment.Symbol
	var resultin environment.Symbol
	result = va.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	newVar = env.(environment.Environment).SaveVariable(va.Id,resultin,va.Tipo,ast)

	if result.Type == environment.BOOLEAN {
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
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
		gen.AddBr()
	}

	return result
} 

func (va Declaracion) ArrayValidation(result environment.Symbol) bool {
	return true
}
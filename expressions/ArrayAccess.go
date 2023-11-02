package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"

	//"fmt"

	//"fmt"
	"strconv"
)

type ArrayAccess struct {
	Lin 		int
	Col 		int
	Array		interfaces.Expression
	Index		[]interface{}
}

func NewArrayAccess(lin int, col int, array interfaces.Expression, index []interface{}) ArrayAccess {
	exp := ArrayAccess{lin,col,array,index}
	return exp
}

func (p ArrayAccess) Ejecutar(ast *environment.AST, env interface{},gen *generator.Generator) environment.Value {
	var result environment.Value
	result = p.Array.Ejecutar(ast, env, gen) //t1, true, ARRAY, size
	//variables RM
	In := 0
	Nn := result.ArrSize


	var FinalIndex int
	//validacion array
	if result.Type == environment.ARRAY || result.Type == environment.VECTOR {
		//agregando primera dimension
		firstIndex := p.Index[0].(interfaces.Expression).Ejecutar(ast, env, gen)
		i, _ := strconv.Atoi(firstIndex.Value)
		//           (i-I1)
		FinalIndex = i - In
		//copiando arreglo de indices
		copiedSlice := make([]interface{}, len(p.Index))
		for i, item := range p.Index {
			if i != 0 {
				copiedSlice[i] = item
			}
		}
		//eliminando primer valor
		copiedSlice = copiedSlice[1:]
		//agregando N dimension
		for _, ind := range copiedSlice {
			firstIndex := ind.(interfaces.Expression).Ejecutar(ast, env, gen)
			j, _ := strconv.Atoi(firstIndex.Value)
			//              (i-I1) * N2 + j - I2
			FinalIndex = FinalIndex*Nn + j - In
		}
		// FinalIndex += 1
		fmt.Println("FinalIndex: ", FinalIndex)

		//validacion array
		lvl1 := gen.NewLabel()
		lvl2 := gen.NewLabel()
		lvl3 := gen.NewLabel()

		gen.AddIf(strconv.Itoa(FinalIndex), "0", ">=", lvl1)
		gen.AddLabel(lvl1)
		gen.AddIf(strconv.Itoa(FinalIndex), strconv.Itoa(Nn), "<=", lvl2)
		gen.AddPrintf("c","66")
		gen.AddPrintf("c","111")
		gen.AddPrintf("c","117")
		gen.AddPrintf("c","110")
		gen.AddPrintf("c","100")
		gen.AddPrintf("c","115")
		gen.AddPrintf("c","69")
		gen.AddPrintf("c","114")
		gen.AddPrintf("c","114")
		gen.AddPrintf("c","111")
		gen.AddPrintf("c","114")
		gen.AddPrintf("c","10")
		gen.AddGoto(lvl3)
		gen.AddLabel(lvl2)
		
		//accediendo al arreglo
		newTmp := gen.NewTemp()
		//se obtiene el array
		tmp := gen.NewTemp()
		gen.AddGetStack(tmp, "(int)"+result.Value)
		//accediendo al arreglo
		gen.AddExpression(newTmp, result.Value, strconv.Itoa(FinalIndex), "+")
		gen.AddExpression(newTmp, newTmp, "1", "+")
		//accediendo al indice
		newTmp2 := gen.NewTemp()
		gen.AddGetStack(newTmp2, "(int)"+newTmp)

		//ToDo: to change
		result = environment.NewValue(newTmp2, true, environment.INTEGER, false, false, false)
		// result.ArrType = tempArray.ArrType
		gen.AddLabel(lvl3)
	}

	return result
}
package expressions

import (
	"Proyecto2_OLC2_2S2023_202101648/Environment"
	"Proyecto2_OLC2_2S2023_202101648/generator"
	"Proyecto2_OLC2_2S2023_202101648/interfaces"
	"fmt"
	"strconv"
	//"strconv"
	//"fmt"
)

type ForRange struct {
	Lin 		int
	Col 		int
	range1 		interfaces.Expression
	range2 		interfaces.Expression
}

func NewForRange(lin int, col int, r1 interfaces.Expression, r2 interfaces.Expression) ForRange {
	return ForRange{lin,col,r1,r2}
}

func (p ForRange) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	var TipoArr environment.TipoExpresion
	var tmpList []interface{}

	rango1 := p.range1.Ejecutar(ast, env, gen)
	rango2 := p.range2.Ejecutar(ast, env, gen)

	if rango1.Type == environment.INTEGER && rango2.Type == environment.INTEGER {
		r1 := rango1.IntValue
		r2 := rango2.IntValue
		tamanio := r2 - r1 + 1

		for i:=r1; i<=r2; i++ {
			rs:= environment.NewValue(strconv.Itoa(i),false,environment.INTEGER,false,false,false)
			tmpList = append(tmpList, rs)
		}
		
		for _, s := range tmpList {
			fmt.Println("VALOR S",s.(environment.Value).Value)
		}

		gen.AddComment("Iniciando arreglo FOR")
		newTmp1 := gen.NewTemp()
		newTmp2 := gen.NewTemp()

		gen.AddAssign(newTmp1, "H")
		gen.AddExpression(newTmp2,newTmp1,"1","+")
		gen.AddSetHeap("(int)H",strconv.Itoa(tamanio))
		gen.AddExpression("H","H",strconv.Itoa(tamanio+1),"+")

		for _ , s := range tmpList {
			TipoArr = s.(environment.Value).Type
			gen.AddSetHeap("(int)"+newTmp2,s.(environment.Value).Value)
			gen.AddExpression(newTmp2,newTmp2,"1","+")
		}
		result = environment.NewValue(newTmp1,true,environment.INTEGER,false,false,false)
		result.ArrType = TipoArr
	}

	return result
}
package generator

import (
	"fmt"
)

type Generator struct {
	Temporal         int
	Label            int
	Code             []interface{}
	FinalCode        []interface{}
	Natives          []interface{}
	FuncCode         []interface{}
	TempList         []interface{}
	PrintStringFlag  bool
	ConcatStringFlag bool
	BreakLabel       string
	ContinueLabel    string
	MainCode         bool
}

func NewGenerator() Generator {
	generator := Generator{
		Temporal:         0,
		Label:            0,
		BreakLabel:       "",
		ContinueLabel:    "",
		PrintStringFlag:  true,
		ConcatStringFlag: true,
		MainCode:         false,
	}
	return generator
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

func (g Generator) GetTemps() []interface{} {
	return g.TempList
}

func (g *Generator) SetMainFlag(newVal bool) {
	g.MainCode = newVal
}

// add break lvl
func (g *Generator) AddBreakLabel(lvl string) {
	g.BreakLabel = lvl
}

// add continue lvl
func (g *Generator) AddContinueLabel(lvl string) {
	g.ContinueLabel = lvl
}

// generar nuevo tmp
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temporal)
	g.Temporal = g.Temporal + 1
	//Lo guardamos para declararlo
	g.TempList = append(g.TempList, temp)
	return temp
}

// Generador de Label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

//Reductor Label para switch
func (g *Generator) ReducLabel() {
	g.Label = g.Label - 1
}

// generar nuevo label
func (g *Generator) AddLabel(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, Label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, Label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+Label+";\n")
	}
}

func (g *Generator) AddGoto(Label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+Label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+Label+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "heap["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "heap["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = heap["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = heap["+index+"];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "stack["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "stack["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = stack["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = stack["+index+"];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}

func (g *Generator) AddTittle(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "void "+target+"() {\n")
	} else {
		g.FuncCode = append(g.FuncCode, "void "+target+"() {\n")
	}
}

func (g *Generator) AddEnd() {
	if g.MainCode {
		g.Code = append(g.Code, "\treturn;\n")
		g.Code = append(g.Code, "}\n\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\treturn;\n")
		g.FuncCode = append(g.FuncCode, "}\n\n")
	}
}

// agregar headers
func (g *Generator) GenerateFinalCode() {
	//****************** add head
	g.FinalCode = append(g.FinalCode, "/*------HEADER------*/\n")
	g.FinalCode = append(g.FinalCode, "#include <stdio.h>\n")
	g.FinalCode = append(g.FinalCode, "#include <math.h>\n")
	g.FinalCode = append(g.FinalCode, "double heap[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double stack[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double P;\n")
	g.FinalCode = append(g.FinalCode, "double H;\n")
	g.FinalCode = append(g.FinalCode, "double ")
	//****************** add temporal declaration
	tempArr := g.GetTemps()
	if len(tempArr) > 0 {
		tmpDec := fmt.Sprintf("%v", tempArr[0])
		tempArr = tempArr[1:]
		for _, s := range tempArr {
			tmpDec += ", "
			tmpDec += fmt.Sprintf("%v", s)
		}
		tmpDec += ";\n\n"
		g.FinalCode = append(g.FinalCode, tmpDec)
	}
	//****************** add natives functions
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------NATIVES------*/\n")
		for _, s := range g.Natives {
			g.FinalCode = append(g.FinalCode, s)
		}
	}
	//****************** add functions
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------FUNCTIONS------*/\n")
		for _, s := range g.FuncCode {
			g.FinalCode = append(g.FinalCode, s)
		}
	}
	//****************** add main
	g.FinalCode = append(g.FinalCode, "/*------MAIN------*/\n")
	g.FinalCode = append(g.FinalCode, "void main() {\n")
	g.FinalCode = append(g.FinalCode, "\tP = 0; H = 0;\n\n")
	for _, s := range g.Code {
		g.FinalCode = append(g.FinalCode, "\t"+s.(string))
	}
	g.FinalCode = append(g.FinalCode, "\n\treturn;\n}\n")
}

func (g *Generator) GeneratePrintString() {
	if g.PrintStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void swift_printString() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = stack[(int)"+newTemp1+"];\n")
		g.Natives = append(g.Natives, "\t"+newLvl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = heap[(int)"+newTemp2+"];\n")
		g.Natives = append(g.Natives, "\tif("+newTemp3+" == -1) goto "+newLvl1+";\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char)"+newTemp3+");\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = "+newTemp2+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl1+":\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.PrintStringFlag = false
	}
}

func (g *Generator) GenerateConcatString() {
	if g.ConcatStringFlag {
		//generando temporales y etiquetas
		newTemp1 := g.NewTemp()
		newTemp2 := g.NewTemp()
		newTemp3 := g.NewTemp()
		newTemp4 := g.NewTemp()
		newTemp5 := g.NewTemp()
		newTemp6 := g.NewTemp()
		newLvl1 := g.NewLabel()
		newLvl2 := g.NewLabel()
		newLvl3 := g.NewLabel()
		newLvl4 := g.NewLabel()
		newLvl5 := g.NewLabel()
		//se genera la funcion printstring
		g.Natives = append(g.Natives, "void swift_concatString() {\n")
		g.Natives = append(g.Natives, "\t"+newTemp1+" = H;\n")
		g.Natives = append(g.Natives, "\t"+newTemp2+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = stack[(int)"+newTemp2+"];\n")
		g.Natives = append(g.Natives, "\t"+newTemp4+" = P + 2;\n")
		g.Natives = append(g.Natives, "\t"+newTemp5+"= stack[(int)"+newTemp4+"];\n")
		g.Natives = append(g.Natives, "\t"+newLvl3+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp6+" = heap[(int)" + newTemp3 + "];\n")
		g.Natives = append(g.Natives, "\t"+"if("+newTemp6+" != -1) goto "+newLvl1+";\n")
		g.Natives = append(g.Natives, "\t"+"goto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl1+":\n")
		g.Natives = append(g.Natives, "\t"+ "heap[(int)H] = "+newTemp6+";\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp3+" = "+newTemp3+" + 1;\n")
		g.Natives = append(g.Natives, "\t"+"goto "+newLvl3+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTemp6+" = heap[(int)" + newTemp5 + "];\n")
		g.Natives = append(g.Natives, "\t"+"if("+newTemp6+" != -1) goto "+newLvl4+";\n")
		g.Natives = append(g.Natives, "\t"+"goto "+newLvl5+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl4+":\n")
		g.Natives = append(g.Natives, "\t"+ "heap[(int)H] = "+newTemp6+";\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTemp5+" = "+newTemp5+" + 1;\n")
		g.Natives = append(g.Natives, "\t"+"goto "+newLvl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLvl5+":\n")
		g.Natives = append(g.Natives, "\t"+"heap[(int)H] = -1;\n")
		g.Natives = append(g.Natives, "\t"+"H = H + 1;\n")
		g.Natives = append(g.Natives, "\t"+"stack[(int)P] = "+newTemp1+";\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.ConcatStringFlag = false
	}
}

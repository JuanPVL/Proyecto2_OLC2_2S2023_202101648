package environment

type AST struct {
	Instructions []interface{}
	Print 	  	 string
	Errors 	  	 []interface{}
	TablaSimbolos []interface{}
}

func NewAST(inst []interface{},print string, errors []interface{}) *AST {
	ast := AST{Instructions: inst, Print: print, Errors: errors}
	return &ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() []interface{} {
	return a.Errors
}

func (a *AST) SetErrors(ToError ErrorS) {
	a.Errors = append(a.Errors, ToError)
}

func (a *AST) GetTablaSimbolos() []interface{} {
	return a.TablaSimbolos
}

func (a *AST) SetTablaSimbolos(ToTablaSimbolos SimbolTabla) {
	a.TablaSimbolos = append(a.TablaSimbolos, ToTablaSimbolos)
}
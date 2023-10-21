package environment

type FunctionSymbol struct {
	Lin 		int
	Col 		int
	Id 			string
	ListaParametros []interface{}
	Bloque 		[]interface{}
	TipoRetorno TipoExpresion
}
package environment

type Symbol struct {
	Lin        int
	Col        int
	Id         string
	Tipo       TipoExpresion
	Valor      interface{}
	Mutable    bool
	ArrSize    int
	Posicion   int
}

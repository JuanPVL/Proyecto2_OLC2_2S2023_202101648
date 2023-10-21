package environment

type StructType struct {
	Id   string
	Tipo TipoExpresion
	StructID string
}

func NewStructType(id string, tip TipoExpresion, id2 string) StructType {
	exp := StructType{id, tip, id2}
	return exp
}
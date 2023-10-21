package environment

type TipoArray struct {
	Tipo TipoExpresion
}

func NewTipoArray(tipo TipoExpresion) TipoArray {
	return TipoArray{tipo}
}
package environment

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	ArrType      TipoExpresion
	TrueLabel    []interface{}
	FalseLabel   []interface{}
	OutLabel     []interface{}
	ArrValue     []interface{}
	ArrSize      int
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
	IntValue     int
}

func NewValue(Val string, tmp bool, typ TipoExpresion, retu, brk, cntinue bool) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
		ArrType:      NULL,
		ArrValue:     []interface{}{},
		ArrSize:      0,
		ReturnFlag:   retu,
		BreakFlag:    brk,
		ContinueFlag: cntinue,
		IntValue:     0,
	}
	return result
}

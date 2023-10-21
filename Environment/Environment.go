package environment

import (
	"fmt"
	"strconv"
)

type Environment struct {
	Anterior interface{}
	Tabla	 map[string]Symbol
	Structs   map[string]Symbol
	Functions map[string]FunctionSymbol
	Id	     string
	Size	 map[string]int
}

func NewEnvironment(anterior interface{},id string) Environment{
	env := Environment{
		Anterior: anterior,
		Tabla: make(map[string]Symbol),
		Structs:   make(map[string]Symbol),
		Functions: make(map[string]FunctionSymbol),
		Id: id,
		Size: make(map[string]int),
	}
	env.Size["size"] = 0
	return env
}

func (env Environment) SaveVariable(id string,value Symbol,tipon TipoExpresion,ast *AST) Symbol {
	var tipo = ""
	linea := strconv.Itoa(value.Lin)
	columna := strconv.Itoa(value.Col)
	if variable, ok := env.Tabla[id]; ok {
		ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Variable ya declarada " + id, Ambito: env.Id})
		fmt.Println("Variable ya declarada: ",variable)
		return env.Tabla[id]
	}
	if value.Tipo == INTEGER {
		tipo = "Int"
	} else if value.Tipo == FLOAT {
		tipo = "Float"
	} else if value.Tipo == STRING {
		tipo = "String"
	} else if value.Tipo == BOOLEAN {
		tipo = "Bool"
	} else if value.Tipo == VECTOR {
		tipo = "Vector"
	} else if value.Tipo == STRUCT {
		tipo = "Struct"
	}
	ast.SetTablaSimbolos(SimbolTabla{Lin: linea, Col: columna, TipoSimbolo: "Variable", TipoDato: tipo, Ambito: env.Id,Id: id})
	env.Tabla[id] = Symbol{Lin: 0, Col: 0, Tipo: tipon, Posicion: env.Size["size"]}
	env.Size["size"] = env.Size["size"] + 1
	return env.Tabla[id]
}

func (env Environment) GetVariable(id string,ast *AST,linea string, columna string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("Variable no declarada: ",id)
	ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Variable no declarada " + id, Ambito: env.Id})
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}

func (env Environment) SetVariable(id string, value Symbol,ast *AST) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			if tmpEnv.Tabla[id].Mutable == true{
				if tmpEnv.Tabla[id].Tipo == value.Tipo {
					tmpEnv.Tabla[id] = value
					return variable
				} else {
					fmt.Println("Tipo de dato incorrecto: ")
					linea := strconv.Itoa(value.Lin)
					columna := strconv.Itoa(value.Col)
					ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Tipo de dato incorrecto" , Ambito: env.Id})
					return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
				}
			} else {
				fmt.Println("Variable no mutable: " , tmpEnv.Tabla[id].Valor)
				linea := strconv.Itoa(value.Lin)
				columna := strconv.Itoa(value.Col)
				ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Variable no mutable" , Ambito: env.Id})
				return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
			}
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("Variable no declarada: ",id)
	linea := strconv.Itoa(value.Lin)
	columna := strconv.Itoa(value.Col)
	ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Variable no declarada" , Ambito: env.Id})
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}

func (env Environment) SaveFunction(id string, value FunctionSymbol,ast *AST) {
	var tipo = ""
	linea := strconv.Itoa(value.Lin)
	columna := strconv.Itoa(value.Col)
	if variable, ok := env.Functions[id]; ok {
		ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Funcion ya existe " + id, Ambito: env.Id})
		fmt.Println("La funcion " + variable.Id + " ya existe")
		return
	}
	if value.TipoRetorno == INTEGER {
		tipo = "Int"
	} else if value.TipoRetorno == FLOAT {
		tipo = "Float"
	} else if value.TipoRetorno == STRING {
		tipo = "String"
	} else if value.TipoRetorno == BOOLEAN {
		tipo = "Bool"
	} else if value.TipoRetorno == VECTOR {
		tipo = "Vector"
	}
	ast.SetTablaSimbolos(SimbolTabla{Lin: linea, Col: columna, TipoSimbolo: "Funcion", TipoDato: tipo, Ambito: env.Id,Id: id})
	env.Functions[id] = value
}

func (env Environment) GetFunction(id string,ast *AST,linea string, columna string) FunctionSymbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Functions[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Funcion no existe " + id, Ambito: env.Id})
	fmt.Println("La funcion ", id, " no existe")
	return FunctionSymbol{TipoRetorno: NULL}
}

func (env Environment) SaveStruct(id string, list []interface{},ast *AST,linea string, columna string) {
	if _, ok := env.Structs[id]; ok {
		ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Struct ya existe " + id, Ambito: env.Id})
		//fmt.Println("El struct " + id + " ya existe")
		return
	}
	env.Structs[id] = Symbol{Lin: 0, Col: 0, Tipo: STRUCT, Valor: list}
	ast.SetTablaSimbolos(SimbolTabla{Lin: linea, Col: columna, TipoSimbolo: "Struct", TipoDato: "STRUCT", Ambito: env.Id,Id: id})
}

func (env Environment) GetStruct(id string,ast *AST,linea string, columna string) Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if tmpStruct, ok := tmpEnv.Structs[id]; ok {
			return tmpStruct
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}

	ast.SetErrors(ErrorS{Lin: linea, Col: columna, Descripcion: "Struct no existe " + id, Ambito: env.Id})
	fmt.Println("El struct ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NULL, Valor: 0}
}
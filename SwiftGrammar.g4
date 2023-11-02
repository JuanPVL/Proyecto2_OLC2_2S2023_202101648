grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Proyecto2_OLC2_2S2023_202101648/interfaces"
    import "Proyecto2_OLC2_2S2023_202101648/Environment"
    import "Proyecto2_OLC2_2S2023_202101648/expressions"
    import "Proyecto2_OLC2_2S2023_202101648/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: printstmt (PUNTOCOMA)? { $inst = $printstmt.prnt}
| ifstmt { $inst = $ifstmt.ifinst }
| declarationstmt (PUNTOCOMA)? { $inst = $declarationstmt.dec }
| asignationstmt (PUNTOCOMA)? { $inst = $asignationstmt.asig }
| whilestmt { $inst = $whilestmt.whileinst }
| forstmt { $inst = $forstmt.forinst }
| guardstmt { $inst = $guardstmt.gd }
| switchstmt { $inst = $switchstmt.sw }
| function {$inst = $function.fun}
| structCreation { $inst = $structCreation.dec }
| callFuncionIns (PUNTOCOMA)? {$inst = $callFuncionIns.cf}
| BREAK (PUNTOCOMA)? {$inst = instructions.NewBreak($BREAK.line, $BREAK.pos)}
| CONTINUE (PUNTOCOMA)? {$inst = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos)}
| ID PUNTO APPEND PAR_IZQ expr PAR_DER (PUNTOCOMA)? {$inst = instructions.NewAppend($ID.line, $ID.pos, $ID.text, $expr.e)}
| ID PUNTO REMOVELAST PAR_IZQ PAR_DER (PUNTOCOMA)? {$inst = instructions.NewRemoveLast($ID.line, $ID.pos, $ID.text)}
| ID PUNTO REMOVE PAR_IZQ AT DOSPUNTOS expr PAR_DER (PUNTOCOMA)? {$inst = instructions.NewRemoveAt($ID.line, $ID.pos, $ID.text, $expr.e)}
| RETURN expr (PUNTOCOMA)? {$inst = instructions.NewReturn($RETURN.line, $RETURN.pos, $expr.e)}
| RETURN (PUNTOCOMA)? {$inst = instructions.NewReturn($RETURN.line, $RETURN.pos, nil)}
;

structCreation returns[interfaces.Instruction dec]
: STRUCT ID LLAVE_IZQ listStructDec LLAVE_DER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
;

listStructDec returns[[]interface{} l]
: list=listStructDec COMA VAR ID DOSPUNTOS types {
                                                var arr []interface{}
                                                newParams := environment.NewStructType($ID.text, $types.ty,"")
                                                arr = append($list.l, newParams)
                                                $l = arr
                                            }
| list=listStructDec COMA VAR idp=ID DOSPUNTOS ids=ID {
                                                var arr []interface{}
                                                newParams := environment.NewStructType($idp.text, environment.DEPENDIENTE,$ids.text)
                                                arr = append($list.l, newParams)
                                                $l = arr
                                            }
| VAR ID DOSPUNTOS types {
                        var arr []interface{}
                        newParams := environment.NewStructType($ID.text, $types.ty,"")
                        arr = append(arr, newParams)
                        $l = arr
                    }
| VAR idp=ID DOSPUNTOS ids=ID {
                        var arr []interface{}
                        newParams := environment.NewStructType($idp.text, environment.DEPENDIENTE,$ids.text)
                        arr = append(arr, newParams)
                        $l = arr
                    }
|  { $l = []interface{}{} }
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PAR_IZQ listParams PAR_DER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$listParams.l)}
;

blockelsif returns [[]interface{} blkif]
@init{
    $blkif = []interface{}{}
    var listIfs []IIfstmtContext
    }
: elseif+=ifstmt+ 
    {
        listIfs = localctx.(*BlockelsifContext).GetElseif()
        for _, e := range listIfs {
            $blkif = append($blkif, e.GetIfinst())
        }
    }
;

ifstmt returns [interfaces.Instruction ifinst]
: IF expr LLAVE_IZQ block LLAVE_DER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil) }
| IF expr LLAVE_IZQ ifblck=block LLAVE_DER ELSE LLAVE_IZQ elseblck=block LLAVE_DER {$ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $ifblck.blk, $elseblck.blk)}
| IF expr LLAVE_IZQ ifblck=block LLAVE_DER ELSE blockelsif {$ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $ifblck.blk, $blockelsif.blkif)}
;

whilestmt returns [interfaces.Instruction whileinst]
: WHILE expr LLAVE_IZQ block LLAVE_DER { $whileinst = instructions.NewWhile($WHILE.line, $WHILE.pos, $expr.e, $block.blk) }
;

guardstmt returns [interfaces.Instruction gd]
: GUARD expr ELSE LLAVE_IZQ block LLAVE_DER { $gd = instructions.NewGuard($GUARD.line, $GUARD.pos, $expr.e, $block.blk) }
;

listcases returns [[]interface{} lcas]
@init{
    $lcas = []interface{}{}
    var listCase []ICasestmtContext
    }
: cases+=casestmt+ 
    {
        listCase = localctx.(*ListcasesContext).GetCases()
        for _, e := range listCase {
            $lcas = append($lcas, e.GetCaseinst())
        }
    }
;

casestmt returns [interfaces.Instruction caseinst]
:CASE expr DOSPUNTOS block { $caseinst = instructions.NewCase($CASE.line, $CASE.pos, $expr.e, $block.blk) }
;

instdefault returns [interfaces.Instruction instdef]
:DEFAULT DOSPUNTOS block { $instdef = instructions.NewDefault($DEFAULT.line, $DEFAULT.pos,$block.blk) }
;

switchstmt returns [interfaces.Instruction sw]
: SWITCH expr LLAVE_IZQ listcases instdefault LLAVE_DER { $sw = instructions.NewSwitch($SWITCH.line, $SWITCH.pos, $expr.e, $listcases.lcas,$instdefault.instdef) }
| SWITCH expr LLAVE_IZQ listcases LLAVE_DER { $sw = instructions.NewSwitch($SWITCH.line, $SWITCH.pos, $expr.e,$listcases.lcas, nil)}
;

forstmt returns [interfaces.Instruction forinst]
: FOR ID IN exprFor LLAVE_IZQ block LLAVE_DER {$forinst = instructions.NewFor($FOR.line, $FOR.pos, $ID.text, $exprFor.e, $block.blk)}
;

declarationstmt returns [interfaces.Instruction dec]
: VAR ID DOSPUNTOS types IGUAL expr  { $dec = instructions.NewDeclaracion($VAR.line, $VAR.pos, $ID.text,true, $types.ty, $expr.e) }
| VAR ID IGUAL expr { $dec = instructions.NewDeclaracion($VAR.line, $VAR.pos, $ID.text,true,environment.DEPENDIENTE, $expr.e) }
| VAR ID DOSPUNTOS types CIERRAPREGUNTA { $dec = instructions.NewDeclaracion($VAR.line, $VAR.pos, $ID.text,true, $types.ty, nil) }
| VAR ID DOSPUNTOS COR_IZQ types COR_DER IGUAL exprvector { $dec = instructions.NewDeclaracionVector($VAR.line, $VAR.pos, $ID.text,true, $types.ty, $exprvector.exprv) }
|VAR ID DOSPUNTOS typesmatriz IGUAL expr  { $dec = instructions.NewDeclaracionMatriz($VAR.line, $VAR.pos, $ID.text,true, $typesmatriz.tm, $expr.e) }
|LET ID DOSPUNTOS typesmatriz IGUAL expr  { $dec = instructions.NewDeclaracionMatriz($LET.line, $LET.pos, $ID.text,false, $typesmatriz.tm, $expr.e) }
| LET ID DOSPUNTOS types IGUAL expr { $dec = instructions.NewDeclaracion($LET.line, $LET.pos, $ID.text,false, $types.ty, $expr.e) }
| LET ID IGUAL expr { $dec = instructions.NewDeclaracion($LET.line, $LET.pos, $ID.text,false,environment.DEPENDIENTE, $expr.e) }
;

asignationstmt returns [interfaces.Instruction asig]
: ID IGUAL expr { $asig = instructions.NewAsignacion($ID.line, $ID.pos, $ID.text, $expr.e) }
| ID COR_IZQ index=expr COR_DER IGUAL listan=expr { $asig = instructions.NewAsignacionIndexVector($ID.line, $ID.pos, $ID.text, $index.e, $listan.e) }
| ID op=(SUM|RES) IGUAL expr {$asig = instructions.NewOperacionAsignacion($ID.line, $ID.pos, $ID.text, $expr.e, $op.text)}
;


function returns [interfaces.Instruction fun]
: FUNC ID PAR_IZQ listParamsFunc PAR_DER LLAVE_IZQ block LLAVE_DER {$fun = instructions.NewFuncion($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf,environment.NULL, $block.blk)}
| FUNC ID PAR_IZQ listParamsFunc PAR_DER FLECHA types LLAVE_IZQ block LLAVE_DER {$fun = instructions.NewFuncion($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, $types.ty, $block.blk)}
| FUNC ID PAR_IZQ listParamsFunc PAR_DER FLECHA COR_IZQ types COR_DER LLAVE_IZQ block LLAVE_DER {$fun = instructions.NewFuncion($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, environment.VECTOR, $block.blk)}
;

listParamsFunc returns[[]interface{} lpf]
: list=listParamsFunc COMA ID DOSPUNTOS types {
    var arr []interface{}
    newParam := instructions.NewDeclaracionParametros($ID.line, $ID.pos, $ID.text, $types.ty)
    arr = append($list.lpf, newParam)
    $lpf = arr
    }
|list=listParamsFunc COMA ID DOSPUNTOS COR_IZQ types COR_DER {
    var arr []interface{}
    newParam := instructions.NewDeclaracionParametros($ID.line, $ID.pos, $ID.text, environment.VECTOR)
    arr = append($list.lpf, newParam)
    $lpf = arr
    }
| ID DOSPUNTOS types {
    $lpf = []interface{}{}
    newParam := instructions.NewDeclaracionParametros($ID.line, $ID.pos, $ID.text, $types.ty)
    $lpf = append($lpf, newParam)
    }
| ID DOSPUNTOS COR_IZQ types COR_DER {
    $lpf = []interface{}{}
    newParam := instructions.NewDeclaracionParametros($ID.line, $ID.pos, $ID.text, environment.VECTOR)
    $lpf = append($lpf, newParam)
    }
|  { $lpf = []interface{}{} }
;

callFuncionIns returns[interfaces.Expression cf]
: ID PAR_IZQ listParamsCall PAR_DER { $cf = expressions.NewLlamadoFuncion($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;

types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| STR { $ty = environment.STRING }
| BOOL { $ty = environment.BOOLEAN }
| COR_IZQ COR_DER { $ty = environment.ARRAY}
;

typesmatriz returns[[]interface{} tm]
:COR_IZQ list=typesmatriz COR_DER {
                                var arr []interface{}
                                newTipo := environment.NewTipoArray(environment.ARRAY)
                                arr = append($list.tm, newTipo)
                                $tm = arr
                            }   
| types {
            $tm = []interface{}{}
            newTipo := environment.NewTipoArray($types.ty)
            $tm = append($tm, newTipo)
        }
;

exprFor returns[interfaces.Expression e]
:range1=expr PUNTO PUNTO PUNTO range2=expr {$e = expressions.NewForRange($range1.start.GetLine(), $range1.start.GetColumn(), $range1.e, $range2.e)}
|expr {$e = $expr.e}
;

expr returns [interfaces.Expression e]
: RES left=expr { $e = expressions.NewOperation($RES.line, $RES.pos, $left.e, "UNARIO", nil) }
| left=expr op=(MULT|DIV|MOD) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(SUM|RES) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAYIG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MENIG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIFE) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| NOT left=expr {$e = expressions.NewOperation($NOT.line, $NOT.pos, $left.e, $NOT.text, nil)}
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| ID PAR_IZQ listStructExp PAR_DER { $e = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
| callFuncion {$e = $callFuncion.cf}
| PAR_IZQ expr PAR_DER { $e = $expr.e }
| conversionstmt { $e = $conversionstmt.conv }
|ID PUNTO COUNT { $e = expressions.NewCount($ID.line, $ID.pos, $ID.text)}
|ID PUNTO ISEMPTY { $e = expressions.NewIsEmpty($ID.line, $ID.pos, $ID.text)}
| list=listArray { $e = $list.p}
| COR_IZQ listParams COR_DER { $e = expressions.NewArray($COR_IZQ.line, $COR_IZQ.pos, $listParams.l) }
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
| NIL { $e = expressions.NewPrimitive($NIL.line, $NIL.pos, nil, environment.NULL) }
;

conversionstmt returns [interfaces.Expression conv]
: INT PAR_IZQ expr PAR_DER { $conv = expressions.NewToInt($INT.line, $INT.pos, $expr.e) }
| FLOAT PAR_IZQ expr PAR_DER { $conv = expressions.NewToFloat($FLOAT.line, $FLOAT.pos, $expr.e) }
| STR PAR_IZQ expr PAR_DER { $conv = expressions.NewToString($STR.line, $STR.pos, $expr.e) }
;

exprvector returns [interfaces.Expression exprv]
: COR_IZQ listParams COR_DER { $exprv = expressions.NewVector($COR_IZQ.line, $COR_IZQ.pos, $listParams.l) }
| COR_IZQ COR_DER { $exprv = expressions.NewVector($COR_IZQ.line, $COR_IZQ.pos, nil) }
| ID { $exprv = expressions.NewLlamadoVar($ID.line, $ID.pos, $ID.text)}
;

listParams returns[[]interface{} l]
: list=listParams COMA expr {
                                var arr []interface{}
                                arr = append($list.l, $expr.e)
                                $l = arr
                            }   
| expr {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
;

listArray returns[interfaces.Expression p]
: list = listArray arr = listAccessArray { $p = expressions.NewArrayAccess($ID.line, $ID.pos, $list.p, $arr.l) }
|list = listArray PUNTO ID { $p = expressions.NewStructAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $ID.text)  }
| ID { $p = expressions.NewLlamadoVar($ID.line, $ID.pos, $ID.text)}
;

listAccessArray returns[[]interface{} l]
: list = listAccessArray COR_IZQ expr COR_DER {
                                                var arr []interface{}
                                                arr = append($list.l, $expr.e)
                                                $l = arr
                                            } 
| COR_IZQ expr COR_DER    {
                            $l = []interface{}{}
                            $l = append($l, $expr.e)
                        }
;

callFuncion returns[interfaces.Expression cf]
: ID PAR_IZQ listParamsCall PAR_DER { $cf = expressions.NewLlamadoFuncion($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;

listParamsCall returns[[]interface{} l]
: list=listParamsCall COMA expr {
                                    var arr []interface{}
                                    arr = append($list.l, $expr.e)
                                    $l = arr
                                }
| expr  {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
|   {
        $l = []interface{}{}
    }
;

listStructExp returns[[]interface{} l]
: list=listStructExp (COMA)? ID DOSPUNTOS expr {
                                            var arr []interface{}
                                            StrExp := environment.NewStructContent($ID.text, $expr.e)
                                            arr = append($list.l, StrExp)
                                            $l = arr
                                        }
| ID DOSPUNTOS expr{
                    var arr []interface{}
                    StrExp := environment.NewStructContent($ID.text, $expr.e)
                    arr = append(arr, StrExp)
                    $l = arr
                }
|   {
        $l = []interface{}{}
    }
;
package lang

import (
	"sync"
	"strconv"
	"strings"
	"encoding/hex"
	"bytes"


	"fmt"
)


type Mtoken struct {
	Tp 		uint8
	Val		interface{}
	Mutex	sync.RWMutex
}

func (tk *Mtoken) Bool() bool{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(bool)
}

func (tk *Mtoken) Byte() byte{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(byte)
}

func (tk *Mtoken) Char() rune{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(rune)
}

func (tk *Mtoken) Int() int{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(int)
}

func (tk *Mtoken) Float() float64{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(float64)
}

func (tk *Mtoken) Str() string{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(string)
}

func (tk *Mtoken) Table() *Mtable {
	// tk.Mutex.RLock()
	// defer tk.Mutex.RUnlock()
	return tk.Val.(*Mtable)
}

func (tk *Mtoken) Tks() []*Mtoken {
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(*MtokenList).List()
}

func (tk *Mtoken) Mcall() *Mcall{
	tk.Mutex.RLock()
	defer tk.Mutex.RUnlock()
	return tk.Val.(*Mcall)
}




func (tk *Mtoken) ToStr() string{
	switch tk.Tp {
	case TNIL:
		return "nil"
	case TNONE:
		return "none"
	case TERR:
		return "ERROR: " + tk.Val.(string)
	case TDATATYPE:
		return Tp2Str(tk.Tp)
	case TBOOL:
		return strconv.FormatBool(tk.Bool())
	case TBYTE:
		return hex.EncodeToString([]byte{tk.Byte()})
	case TCHAR:
		return "#'" + string(tk.Char()) + "'"
	case TINT:
		return strconv.Itoa(tk.Int())
	case TFLOAT:
		var result = strconv.FormatFloat(tk.Float(), 'f', -1, 64)
		if strings.IndexByte(result, '.') < 0 {
			result += ".0"
		}
		return result
	case TSTRING:
		return "\"" + tk.Str() + "\""
	
	case TPAREN:
		var buffer bytes.Buffer
		buffer.WriteString("(")
		for _, item := range tk.Tks(){
			buffer.WriteString(item.ToStr())
			buffer.WriteString(" ")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ')'
		}else{
			buffer.WriteString(")")
		}
		return buffer.String()
	case TBLOCK:
		var buffer bytes.Buffer
		buffer.WriteString("[")
		for _, item := range tk.Tks(){
			buffer.WriteString(item.ToStr())
			buffer.WriteString(" ")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ']'
		}else{
			buffer.WriteString("]")
		}
		return buffer.String()


	case TWORD:
		return tk.Str()
	case TCALL:
		var call = tk.Val.(*Mcall)
		var result = call.Name + "("
		var buffer bytes.Buffer
		for _, item := range call.Args{
			buffer.WriteString(item.ToStr())
			buffer.WriteString(" ")
		}
		if len(buffer.Bytes()) > 1 {
			buffer.Bytes()[len(buffer.Bytes())-1] = ')'
		}else{
			buffer.WriteString(")")
		}
		result = result + buffer.String()
		return result


	case TLIT_WORD:
		return "'" + tk.Str()
	case TSET_WORD:
		return tk.Str() + ":"

	case TNATIVE:
		return "native"
	case TFUNC:
		return "func"
	case TOP:
		return "op"


	default:
		return "undefined"
	}
}


func (tk *Mtoken) OutputStr() string {
	var result = tk.ToStr()
	if tk.Tp == TSTRING {
		result = result[1:len(result)-1]
	}
	return result
}

func (tk *Mtoken) Echo(){
	fmt.Println(tk.ToStr())
}


func (tk *Mtoken) Copy(source *Mtoken){
	tk.Tp = source.Tp
	tk.Val = source.Val
}

func (tk *Mtoken) Dup() *Mtoken{
	return &Mtoken{tk.Tp, tk.Val, sync.RWMutex{}}
}

func (tk *Mtoken) Clone() *Mtoken{
	var result = &Mtoken{tk.Tp, tk.Val, sync.RWMutex{}}
	return result
	// switch tk.Tp {
	// case TBLOCK, TPAREN:
	// 	result.Val = NewTks(8)
	// 	result.List().AddAll(tK.List())
		
	// 	return result
	// case OBJECT:
	// 	result.Val = &BindMap{make(map[string]*Token), t.Ctx().Father, t.Ctx().Tp, sync.RWMutex{}}
	// 	for k, v := range(t.Ctx().Table) {
	// 		result.Ctx().PutNow(k, v.Clone())
	// 	}
	// 	return result
	// case MAP:
	// 	result.Val = t.Map().Clone()
	// 	return result
	// default:
	// 	return result
	// }
}

func (tk *Mtoken) CloneDeep() *Mtoken{
	var result = &Mtoken{tk.Tp, tk.Val, sync.RWMutex{}}
	return result

	// switch t.Tp {
	// case BLOCK, PAREN, PATH:
	// 	result.Val = NewTks(8)
	// 	for _, item := range(t.Tks()){
	// 		result.List().Add(item.CloneDeep())
	// 		// result.Val = append(result.Tks(), item.CloneDeep())
	// 	}
	// 	return result
	// case OBJECT:
	// 	result.Val = &BindMap{make(map[string]*Token), t.Ctx().Father, t.Ctx().Tp, sync.RWMutex{}}
	// 	for k, v := range(t.Ctx().Table) {
	// 		result.Ctx().PutNow(k, v.CloneDeep())
	// 	}
	// 	return result
	// case MAP:  
	// 	result.Val = t.Map().CloneDeep()
	// 	return result
	// default:
	// 	return result
	// }
}




func (tk *Mtoken) GetVal(ctx *Mtoken) *Mtoken {
	var result = &Mtoken{Mutex: sync.RWMutex{}}

	switch tk.Tp {
	case TWORD:
		ctx.Mutex.RLock()
		defer ctx.Mutex.RUnlock()
		return ctx.Table().Get(tk.Str())

	case TLIT_WORD:
		result.Tp = TWORD
		result.Val = tk.Str()
		return result

	case TCALL:
		return tk.Val.(*Mcall).Exec(ctx)
	

	default:
		return tk
	}

}







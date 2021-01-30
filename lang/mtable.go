package lang


import (
	"runtime"
	"bytes"
	"fmt"
	"sync"
)

const (
	SYS_CTX = iota
	USR_CTX
	TMP_CTX
)
	


type Mtable struct{
	Table 	map[string]*Mtoken
	Father 	*Mtoken
	Tp 		int
}

func (mt *Mtable)GetNow(key string) *Mtoken{
	var tk *Mtoken
	var ok bool

	tk, ok = mt.Table[key]

	if ok {
		return tk
	}else{
		return &Mtoken{TNONE, "", sync.RWMutex{}}
	}
}

func (mt *Mtable) Get(key string) *Mtoken{

	var ctx = mt
	var prev = ctx

	var tk *Mtoken
	var ok bool
	if(ctx.Table != nil){
	
		tk, ok = ctx.Table[key]
	
		if(ok){
			return tk
		}
	}

	for !ok && ctx.Father != nil {
		prev = ctx
		ctx.Father.Mutex.RLock()
		defer ctx.Father.Mutex.RUnlock()
		ctx = ctx.Father.Table()
		
		if(ctx.Table != nil){
		
			tk, ok = ctx.Table[key]
		
		}
		
	}

	if tk != nil {
		if ctx.Father == nil {
			prev.Table[key] = tk
		}
		return tk
	}else{
		return &Mtoken{Tp: TNONE, Mutex: sync.RWMutex{}}
	}
}


func (mt *Mtable)PutNow(key string, val *Mtoken){
	mt.Table[key] = val
}


func (mt *Mtable)Put(key string, val *Mtoken){

	var ctx = mt
	var inserted = false
	var ok = false

	if(ctx.Table != nil){
		_, ok = ctx.Table[key]
	}

	if(ok){
		mt.Table[key] = val.Clone()
		inserted = true
	}else{
		for !inserted && !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				_, ok = ctx.Table[key]
			}
			if(ok){
				ctx.Table[key] = val.Clone()

				inserted = true
				break
			}
			ctx.Father.Mutex.Lock()
			defer ctx.Father.Mutex.Unlock()
			ctx = ctx.Father.Table()
		}
	}
	if(!inserted){
		mt.PutLocal(key, val)
	}
}


func (mt *Mtable)PutLocal(key string, val *Mtoken){
	var ctx = mt

	for ctx.Tp != USR_CTX && ctx.Father != nil {
		ctx.Father.Mutex.RLock()
		ctx = ctx.Father.Table()
		ctx.Father.Mutex.RUnlock()
	}

	ctx.Table[key] = val.Dup()

}


func (mt *Mtable)Unset(key string){

	var ctx = mt
	var ok = false

	if(ctx.Table != nil){
		_, ok = ctx.Table[key]
	}

	if(ok){
		delete(ctx.Table, key)
		runtime.GC()
	}else{
		for !ok && ctx.Father != nil {
			if(ctx.Table != nil){
				_, ok = ctx.Table[key]
			}
			if(ok){
				delete(ctx.Table, key)
				runtime.GC()
				break
			}
			ctx.Father.Mutex.Lock()
			defer ctx.Father.Mutex.Unlock()
			ctx = ctx.Father.Table()
		}
	}
	
}

func (mt *Mtable) Echo() {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	for k, v := range mt.Table {
		buffer.WriteString(k)
		buffer.WriteString(": ")
		buffer.WriteString(v.ToStr())
		buffer.WriteString(" ")
	}
	if len(buffer.Bytes()) > 1 {
		buffer.Bytes()[len(buffer.Bytes())-1] = '}'
	}else{
		buffer.WriteString("}")
	}
	fmt.Println(buffer.String())
}

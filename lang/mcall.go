package lang

import (
	"sync"
)

type Mcall struct {
	Name string
	Args []*Mtoken
}

func (mc *Mcall) Exec(ctx *Mtoken) *Mtoken {
	ctx.Mutex.RLock()
	var tk = ctx.Table().Get(mc.Name)
	ctx.Mutex.RUnlock()

	if tk.Tp == TNATIVE {
		return tk.Val.(*Mnative).Exec(mc.Args, ctx)
	} else if tk.Tp == TFUNC {
		return tk.Val.(*Mfunc).Exec(mc.Args, ctx)
	} else {
		var result = Mtoken{Tp: TERR, Val: "No such native! or func! for " + mc.Name}
		return &result
	}

}

func (mc *Mcall) DimFunc(ctx *Mtoken, code string) *Mtoken {
	ctx.Mutex.RLock()
	var f = ctx.Table().Get(mc.Name)
	ctx.Mutex.RUnlock()

	if f.Tp != TFUNC {
		f = &Mtoken{Tp: TFUNC, Mutex: sync.RWMutex{}}
		f.Val = &Mfunc{make(map[int]MfuncExec, 4), make(map[string]MfuncExec, 4)}
	}

	code = code[1:len(code)-1]

	
	var isAllWord = true
	for _, item := range mc.Args {
		if item.Tp != TWORD {
			isAllWord = false
		}
	}

	if isAllWord {
		var funcExec = MfuncExec{mc.Args, code}
		f.Mutex.Lock()
		f.Val.(*Mfunc).CountMap[len(mc.Args)] = funcExec
		f.Mutex.Unlock()
	}else{
		var dimWithTypes = true
		if len(mc.Args) % 2 != 0 {
			return &Mtoken{TERR, "Worng args for dim function!", sync.RWMutex{}}
		}

		var typesStr = ""
		var args []*Mtoken
		var idx = 0
		for idx < len(mc.Args) {
			if mc.Args[idx].Tp != TSET_WORD || mc.Args[idx+1].Tp != TDATATYPE {
				dimWithTypes = false
			}
			var arg = mc.Args[idx].Dup()
			arg.Tp = TWORD
			args = append(args, arg)
			typesStr += Tp2Str(mc.Args[idx+1].Val.(uint8))
			idx += 2
		}

		if !dimWithTypes {
			return &Mtoken{TERR, "Worng args for dim function!", sync.RWMutex{}}
		}

		var funcExec = MfuncExec{args, code}

		f.Mutex.Lock()
		f.Val.(*Mfunc).TypesMap[typesStr] = funcExec
		f.Mutex.Unlock()
	}

	

	ctx.Mutex.Lock()
	ctx.Table().Put(mc.Name, f)
	ctx.Mutex.Unlock()
	return f
}

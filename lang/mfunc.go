package lang

import (
	"sync"
)

type MfuncExec struct{
	Args 	[]*Mtoken
	Codes 	string
}

type Mfunc struct{
	CountMap 	map[int]MfuncExec
	TypesMap 	map[string]MfuncExec
}



func (mf *Mfunc) Exec(args []*Mtoken, ctx *Mtoken) *Mtoken{
	var ms = Msolver{}
	ms.Clear()

	var argsAct = ms.ReduceBlk(args, ctx)

	var exec MfuncExec
	var found = false

	if len(mf.TypesMap) > 0 {
		var typesStr = ""
		for _, item := range argsAct {
			typesStr += Tp2Str(item.Tp)
		}

		exec, found = mf.TypesMap[typesStr]
	}

	if !found {
		var argsLen = len(argsAct)
		exec, found = mf.CountMap[argsLen]
	}

	if !found {
		var result = Mtoken{Tp: TERR, Val: "No such native!"}
		return &result
	}
	
	var funcTable = Mtable{make(map[string]*Mtoken, 128), ctx, USR_CTX}
	var funcCtx = Mtoken{TTABLE, &funcTable, sync.RWMutex{}}

	for idx, item := range exec.Args {
		funcTable.PutNow(item.Str(), argsAct[idx])
	}

	
	return ms.EvalStr(exec.Codes, &funcCtx)
	
}













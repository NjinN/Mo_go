package lang

import "sync"

// import (
// 	"fmt"
// )

const (
	MODEL_BLK = iota
	MODEL_STR
)


type Msolver struct{
	InpStr 		[]string
	InpBlk 		[]*Mtoken
	InpLen 		int
	Idx			int
	NowTk		*Mtoken
	NextTk		*Mtoken
	NextTkVal 	*Mtoken
	Model 		uint8
}

func (ms *Msolver) Clear(){
	ms.InpStr = []string{}
	ms.InpBlk = []*Mtoken{}
	ms.InpLen = 0
	ms.Idx = 0
	ms.NowTk = nil
	ms.NextTk = nil
	ms.NextTkVal = nil
}

func (ms *Msolver) EvalStr(inpStr string, ctx *Mtoken) *Mtoken{
	ms.Clear()
	ms.Model = MODEL_STR
	var strs = StrCut(inpStr)
	ms.InpStr = strs
	ms.InpLen = len(strs)
	var result *Mtoken
	for ms.Idx < ms.InpLen {
		result = ms.EvalOne(ctx, true)
	}
	return result
}

func (ms *Msolver) ReduceStr(inpStr string, ctx *Mtoken) []*Mtoken{
	ms.Clear()
	ms.Model = MODEL_STR
	var strs = StrCut(inpStr)
	ms.InpStr = strs
	ms.InpLen = len(strs)
	var result []*Mtoken
	for ms.Idx < ms.InpLen {
		result = append(result, ms.EvalOne(ctx, true))
	}
	return result
}


func (ms *Msolver) EvalBlk(inpBlk []*Mtoken, ctx *Mtoken) *Mtoken{
	ms.Clear()
	ms.Model = MODEL_BLK
	ms.InpBlk = inpBlk
	ms.InpLen = len(inpBlk)
	var result *Mtoken
	for ms.Idx < ms.InpLen {
		result = ms.EvalOne(ctx, true)
	}
	return result
}

func (ms *Msolver) ReduceBlk(inpBlk []*Mtoken, ctx *Mtoken) []*Mtoken{
	ms.Clear()
	ms.Model = MODEL_BLK
	ms.InpBlk = inpBlk
	ms.InpLen = len(inpBlk)
	var result []*Mtoken
	for ms.Idx < ms.InpLen {
		result = append(result, ms.EvalOne(ctx, true))
	}
	return result
}


func (ms *Msolver) PreRead(ctx *Mtoken){
	if ms.Idx >= ms.InpLen {
		return
	}

	if ms.NextTk == nil {
		if ms.Model == MODEL_STR {
			ms.NowTk = MakeMtoken(ms.InpStr[ms.Idx], ctx)
		}else{
			ms.NowTk = ms.InpBlk[ms.Idx]
		}
		
	}else{
		ms.NowTk = ms.NextTk
	}

	if ms.Idx < ms.InpLen - 1 {
		if ms.Model == MODEL_STR {
			ms.NextTk = MakeMtoken(ms.InpStr[ms.Idx + 1], ctx)
		}else{
			ms.NextTk = ms.InpBlk[ms.Idx + 1]
		}
		
		ms.NextTkVal = ms.NextTk.GetVal(ctx)
	}

	ms.Idx++
}



func (ms *Msolver) EvalOne(ctx *Mtoken, preRead bool) *Mtoken {
	if preRead {
		ms.PreRead(ctx)
	}

	ms.NowTk = ms.NowTk.GetVal(ctx)

	if ms.NowTk.Tp == TSET_WORD {
		var k = ms.NowTk.Str()
		var v = ms.EvalOne(ctx, true)
		var table = ctx.Table()
		ctx.Mutex.Lock()
		defer ctx.Mutex.Unlock()
		table.Put(k, v)
		return v
	}else if ms.NowTk.Tp == TSET_FUNC {
		if ms.NextTk.Tp != TBLOCK && ms.NextTk.Tp != TSTRING {
			return &Mtoken{TERR, "function! body must be block! or string!", sync.RWMutex{}}
		}
		var call = ms.NowTk.Mcall()
		ms.Idx += 1
		ms.NowTk = nil
		ms.NextTk = nil
		ms.NextTkVal = nil
		return call.DimFunc(ctx, ms.InpStr[ms.Idx - 1])

	}else if ms.NextTkVal != nil && ms.NextTkVal.Tp == TOP {
		var argL = ms.NowTk
		var op = ms.NextTkVal
		ms.PreRead(ctx)
		ms.PreRead(ctx)
		var argR = ms.NowTk.GetVal(ctx)

		var temp = op.Val.(*Mnative).Exec([]*Mtoken{argL, argR}, ctx)

		if ms.NextTkVal.Tp == TOP {
			ms.NowTk = temp
			return ms.EvalOne(ctx, false)
		}else{
			return temp
		}

	}

	return ms.NowTk
}

























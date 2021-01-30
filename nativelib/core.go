package nativelib

import (
	"sync"

	. "github.com/NjinN/Mo_go/lang"
)

func initReduce() *Mtoken{
	var native = Mnative{
		Name: "add",
		CountMap: make(map[int]MnativeExec, 4),
		TypesMap: make(map[string]MnativeExec, 4),
	}

	native.TypesMap["block!"] = reduce_block
	native.TypesMap["string!"] = reduce_string


	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func reduce_block(args []*Mtoken, ctx *Mtoken) *Mtoken {
	var ms = Msolver{}
	ms.Clear()

	var result = Mtoken{Tp: TBLOCK, Mutex: sync.RWMutex{}}
	var list = NewTks(8)
	list.AddArr(ms.ReduceBlk(args[0].Tks(), ctx))
	result.Val = list
	return &result
}

func reduce_string(args []*Mtoken, ctx *Mtoken) *Mtoken {
	var ms = Msolver{}
	ms.Clear()

	var result = Mtoken{Tp: TBLOCK, Mutex: sync.RWMutex{}}
	var list = NewTks(8)
	list.AddArr(ms.ReduceStr(args[0].Str(), ctx))
	result.Val = list
	return &result
}




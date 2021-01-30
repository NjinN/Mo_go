package nativelib

import (
	"sync"

	. "github.com/NjinN/Mo_go/lang"
)


func initIf() *Mtoken{
	var native = Mnative{
		Name: "if",
		// CountMap: make(map[int]MnativeExec, 4),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["bool!block!"] = if_block



	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func if_block(args []*Mtoken, ctx *Mtoken) *Mtoken {
	var ms = Msolver{}
	ms.Clear()

	if args[0].Bool() {
		return ms.EvalBlk(args[1].Clone().Tks(), ctx)
	}
	
	return &Mtoken{TNIL, nil, sync.RWMutex{}}
}



func initEither() *Mtoken{
	var native = Mnative{
		Name: "either",
		// CountMap: make(map[int]MnativeExec, 4),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["bool!block!block!"] = either_block_block



	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func either_block_block(args []*Mtoken, ctx *Mtoken) *Mtoken {
	var ms = Msolver{}
	ms.Clear()

	if args[0].Bool() {
		return ms.EvalBlk(args[1].Clone().Tks(), ctx)
	}else{
		return ms.EvalBlk(args[2].Clone().Tks(), ctx)
	}
	
}






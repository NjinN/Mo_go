package lang

type MnativeExec func(args []*Mtoken, ctx *Mtoken) *Mtoken

type Mnative struct {
	Name 		string
	CountMap 	map[int]MnativeExec
	TypesMap 	map[string]MnativeExec
}


func (mn *Mnative) Exec(args []*Mtoken, ctx *Mtoken) *Mtoken{
	var ms = Msolver{}
	ms.Clear()

	var argsAct = ms.ReduceBlk(args, ctx)

	if len(mn.TypesMap) > 0 {
		var typesStr = ""
		for _, item := range argsAct {
			typesStr += Tp2Str(item.Tp)
		}

		exec, ok := mn.TypesMap[typesStr]
		if ok {
			return exec(argsAct, ctx)
		}

	}

	var argsLen = len(argsAct)
	exec, ok := mn.CountMap[argsLen]
	if ok {
		return exec(argsAct, ctx)
	}

	var result = Mtoken{Tp: TERR, Val: "No such native! for " + mn.Name}
	return &result
}





















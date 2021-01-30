package nativelib

import (
	"sync"

	. "github.com/NjinN/Mo_go/lang"
)

func InitEq() *Mtoken{
	var native = Mnative{
		Name: "eq",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = eq_int_int
	native.TypesMap["int!float!"] = eq_int_float
	native.TypesMap["float!int!"] = eq_float_int
	native.TypesMap["float!float!"] = eq_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func eq_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Int() == args[1].Int(),
		sync.RWMutex{},
	}
}

func eq_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		float64(args[0].Int()) == args[1].Float(),
		sync.RWMutex{},
	}
}

func eq_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() == float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func eq_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() == args[1].Float(),
		sync.RWMutex{},
	}
}





func InitLt() *Mtoken{
	var native = Mnative{
		Name: "lt",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = lt_int_int
	native.TypesMap["int!float!"] = lt_int_float
	native.TypesMap["float!int!"] = lt_float_int
	native.TypesMap["float!float!"] = lt_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func lt_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Int() < args[1].Int(),
		sync.RWMutex{},
	}
}

func lt_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		float64(args[0].Int()) < args[1].Float(),
		sync.RWMutex{},
	}
}

func lt_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() < float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func lt_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() < args[1].Float(),
		sync.RWMutex{},
	}
}


func InitGt() *Mtoken{
	var native = Mnative{
		Name: "gt",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = gt_int_int
	native.TypesMap["int!float!"] = gt_int_float
	native.TypesMap["float!int!"] = gt_float_int
	native.TypesMap["float!float!"] = gt_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func gt_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Int() > args[1].Int(),
		sync.RWMutex{},
	}
}

func gt_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		float64(args[0].Int()) > args[1].Float(),
		sync.RWMutex{},
	}
}

func gt_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() > float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func gt_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() > args[1].Float(),
		sync.RWMutex{},
	}
}




func InitLe() *Mtoken{
	var native = Mnative{
		Name: "le",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = le_int_int
	native.TypesMap["int!float!"] = le_int_float
	native.TypesMap["float!int!"] = le_float_int
	native.TypesMap["float!float!"] = le_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func le_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Int() <= args[1].Int(),
		sync.RWMutex{},
	}
}

func le_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		float64(args[0].Int()) <= args[1].Float(),
		sync.RWMutex{},
	}
}

func le_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() <= float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func le_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() <= args[1].Float(),
		sync.RWMutex{},
	}
}

func InitGe() *Mtoken{
	var native = Mnative{
		Name: "ge",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = ge_int_int
	native.TypesMap["int!float!"] = ge_int_float
	native.TypesMap["float!int!"] = ge_float_int
	native.TypesMap["float!float!"] = ge_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func ge_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Int() >= args[1].Int(),
		sync.RWMutex{},
	}
}

func ge_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		float64(args[0].Int()) >= args[1].Float(),
		sync.RWMutex{},
	}
}

func ge_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() >= float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func ge_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TBOOL,
		args[0].Float() >= args[1].Float(),
		sync.RWMutex{},
	}
}



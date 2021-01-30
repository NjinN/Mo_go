package nativelib

import (
	"sync"

	. "github.com/NjinN/Mo_go/lang"
)

func InitAdd() *Mtoken{
	var native = Mnative{
		Name: "add",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = add_int_int
	native.TypesMap["int!float!"] = add_int_float
	native.TypesMap["float!int!"] = add_float_int
	native.TypesMap["float!float!"] = add_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}


func add_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TINT,
		args[0].Int() + args[1].Int(),
		sync.RWMutex{},
	}
}

func add_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		float64(args[0].Int()) + args[1].Float(),
		sync.RWMutex{},
	}
}

func add_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() + float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func add_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() + args[1].Float(),
		sync.RWMutex{},
	}
}


func InitSub() *Mtoken{
	var native = Mnative{
		Name: "sub",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = sub_int_int
	native.TypesMap["int!float!"] = sub_int_float
	native.TypesMap["float!int!"] = sub_float_int
	native.TypesMap["float!float!"] = sub_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}


func sub_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TINT,
		args[0].Int() - args[1].Int(),
		sync.RWMutex{},
	}
}

func sub_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		float64(args[0].Int()) - args[1].Float(),
		sync.RWMutex{},
	}
}

func sub_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() - float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func sub_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() - args[1].Float(),
		sync.RWMutex{},
	}
}


func InitMul() *Mtoken{
	var native = Mnative{
		Name: "mul",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = mul_int_int
	native.TypesMap["int!float!"] = mul_int_float
	native.TypesMap["float!int!"] = mul_float_int
	native.TypesMap["float!float!"] = mul_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}


func mul_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TINT,
		args[0].Int() * args[1].Int(),
		sync.RWMutex{},
	}
}

func mul_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		float64(args[0].Int()) * args[1].Float(),
		sync.RWMutex{},
	}
}

func mul_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() * float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func mul_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() * args[1].Float(),
		sync.RWMutex{},
	}
}


func InitDiv() *Mtoken{
	var native = Mnative{
		Name: "div",
		CountMap: make(map[int]MnativeExec, 8),
		TypesMap: make(map[string]MnativeExec, 8),
	}

	native.TypesMap["int!int!"] = div_int_int
	native.TypesMap["int!float!"] = div_int_float
	native.TypesMap["float!int!"] = div_float_int
	native.TypesMap["float!float!"] = div_float_float

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}


func div_int_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		float64(args[0].Int()) / float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func div_int_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		float64(args[0].Int()) / args[1].Float(),
		sync.RWMutex{},
	}
}

func div_float_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() / float64(args[1].Int()),
		sync.RWMutex{},
	}
}

func div_float_float(args []*Mtoken, ctx *Mtoken) *Mtoken {
	return &Mtoken{
		TFLOAT,
		args[0].Float() / args[1].Float(),
		sync.RWMutex{},
	}
}




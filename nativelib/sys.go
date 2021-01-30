package nativelib

import (
	"fmt"
	"sync"
	"os"
	"os/exec"

	. "github.com/NjinN/Mo_go/lang"
)


func initQuit() *Mtoken{
	var native = Mnative{
		Name: "quit",
		CountMap: make(map[int]MnativeExec, 2),
		TypesMap: make(map[string]MnativeExec, 2),
	}

	native.CountMap[0] = quit_zero
	native.TypesMap["int!"] = quit_int

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func quit_zero(args []*Mtoken, ctx *Mtoken) *Mtoken {
	os.Exit(0)
	return nil
}

func quit_int(args []*Mtoken, ctx *Mtoken) *Mtoken {
	os.Exit(args[0].Int())
	return nil
}




func initClear() *Mtoken {
	var native = Mnative{
		Name: "clear",
		CountMap: make(map[int]MnativeExec, 2),
	}

	native.CountMap[0] = clear_zero

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func clear_zero(args []*Mtoken, ctx *Mtoken) *Mtoken {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return &Mtoken{TNIL, nil, sync.RWMutex{}}
}



func initPrint() *Mtoken {
	var native = Mnative{
		Name: "print",
		CountMap: make(map[int]MnativeExec, 2),
	}

	native.CountMap[1] = print_one

	var result = Mtoken{
		TNATIVE,
		&native,
		sync.RWMutex{},
	}

	return &result
}

func print_one(args []*Mtoken, ctx *Mtoken) *Mtoken {
	fmt.Println(args[0].OutputStr())
	return &Mtoken{TNIL, nil, sync.RWMutex{}}
}



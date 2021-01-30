package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"sync"
	

	. "github.com/NjinN/Mo_go/lang"
	. "github.com/NjinN/Mo_go/nativelib"
	. "github.com/NjinN/Mo_go/oplib"
)


func main(){
	// fmt.Println("Hello world")
	// fmt.Println(StrCut("1 2 3 \"456\" [7 \"]89\" [0 1] ] "))

	var libTable = Mtable{make(map[string]*Mtoken, 256), nil, SYS_CTX}
	var libCtx = Mtoken{TTABLE, &libTable, sync.RWMutex{}}

	InitNative(&libCtx)
	InitOp(&libCtx)

	// libTable.Echo()

	var usrTable = Mtable{make(map[string]*Mtoken, 128), &libCtx, USR_CTX}
	var usrCtx = Mtoken{TTABLE, &usrTable, sync.RWMutex{}}

	var ms = Msolver{}

	/** 获取控制台输入并执行 **/
	var reader = bufio.NewReader(os.Stdin)
	var inp string
	for {
		fmt.Print(">> ")

		temp, _ := reader.ReadString('\n')
		// temp = strings.Replace(temp, "\r\n", "", -1)
		temp = strings.TrimSpace(temp)
		if temp == "" {
			continue
		}

		if temp[len(temp)-1] == '~' {
			inp += temp[0 : len(temp)-1]
			continue
		} else {
			if len(inp) > 0 {
				inp += temp
			} else {
				inp = temp
			}
		}

		inp = strings.TrimSpace(inp)
		
		ms.Clear()
		
		var ans = ms.EvalStr(inp, &usrCtx)

		if ans != nil && ans.Tp != TNIL {
			fmt.Println(ans.ToStr())
		}else{
			fmt.Println("")
		}
		
		inp = ""
	}

}


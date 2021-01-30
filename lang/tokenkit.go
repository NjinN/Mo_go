package lang

import (

	"sync"
	"strings"
	"strconv"
	"regexp"

	// "fmt"
)


func MakeMtoken(s string, ctx *Mtoken) *Mtoken{
	var result Mtoken
	result.Mutex = sync.RWMutex{}

	var str = strings.TrimSpace(s)
	// fmt.Println(str)
	var runes = []rune(str)

	if strings.ToLower(str) == "nil" {
		result.Tp = TNIL
		return &result
	}

	if strings.ToLower(str) == "none" {
		result.Tp = TNONE
		result.Val = "none"
		return &result
	}

	if strings.ToLower(str) == "true" {
		result.Tp = TBOOL
		result.Val = true
		return &result
	}

	if strings.ToLower(str) == "false" {
		result.Tp = TBOOL
		result.Val = false
		return &result
	}

	if len(runes) == 4 && str[0:2] == "#'" && runes[3] == rune('"') {
		result.Tp = TCHAR
		result.Val = runes[2]
		return &result
	}

	if str[0] == '"' {
		result.Tp = TSTRING
		result.Val = str[1:len(str)-1]
		return &result
	}

	if str[len(str)-1] == '!' {
		result.Tp = TDATATYPE
		result.Val = Str2Tp(str)
		return &result
	}

	if IsNumberStr(str) == 0 {
		result.Tp = TINT
		i, err := strconv.Atoi(str)

		if err != nil {
			panic(err)
		}else{
			result.Val = i
		}
		return &result
	}

	if IsNumberStr(str) == 1 {
		result.Tp = TFLOAT
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}else{
			result.Val = f
		}
		return &result
	}

	if(str[0] == '['){
		result.Tp = TBLOCK
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == ']'){
				break
			}
		}
		result.Val = NewTks(8)
		result.Val.(*MtokenList).AddArr(MakeMtokens(str[1 : endIdx], ctx))
		return &result
	}

	if(str[0] == '('){
		result.Tp = TPAREN
		var endIdx int
		for endIdx=len(str)-1; endIdx>=0; endIdx-- {
			if(str[endIdx] == ')'){
				break
			}
		}
		result.Val = NewTks(8)
		result.Val.(*MtokenList).AddArr(MakeMtokens(str[1 : endIdx], ctx))
		return &result
	}



	matched, err := regexp.MatchString("^.+\\(.*\\):$", str)
	if err != nil {
		result.Tp = TERR
		result.Val = err.Error()
		return &result
	}
	if matched {
		// fmt.Println("is set-func")
		var parenIdx = strings.Index(str, "(")
		result.Tp = TSET_FUNC
		var call = Mcall{}
		call.Name = str[0:parenIdx]
		call.Args = MakeMtokens(str[parenIdx+1:len(str)-2], ctx)
		result.Val = &call
		return &result
	}


	matched, err = regexp.MatchString("^.+\\(.*\\)$", str)
	if err != nil {
		result.Tp = TERR
		result.Val = err.Error()
		return &result
	}
	if matched {
		// fmt.Println("is call")
		var parenIdx = strings.Index(str, "(")
		result.Tp = TCALL
		var call = Mcall{}
		call.Name = str[0:parenIdx]
		call.Args = MakeMtokens(str[parenIdx+1:len(str)-1], ctx)
		result.Val = &call
		return &result
	}





	if str[0] == '\'' {
		result.Tp = TLIT_WORD
		result.Val = str[1:]
		return &result
	}

	if str[len(str)-1] == ':' {
		result.Tp = TSET_WORD
		result.Val = str[0:len(str)-1]
		return &result
	}

	result.Tp = TWORD
	result.Val = str
	return &result
}



func MakeMtokens(str string, ctx *Mtoken) []*Mtoken{
	var result []*Mtoken
	var strs = StrCut(str)
	for _, item := range strs {
		result = append(result, MakeMtoken(item, ctx))
	}
	return result
}







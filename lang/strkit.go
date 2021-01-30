package lang

import (
	"unicode"
	"strings"
	// "fmt"
)




func StrCut(str string) []string{
	var result []string
	var runes = []rune(strings.TrimSpace(str))
	var idx = 0

	for idx < len(runes){
		if runes[idx] == rune('"'){
			result = append(result, TakeOneStr(runes, &idx))
		}else if runes[idx] == rune('['){
			result = append(result, TakeOneBlock(runes, &idx))
		}else if runes[idx] == rune('('){
			result = append(result, TakeOneParen(runes, &idx))
		}else if runes[idx] == rune('{'){
			result = append(result, TakeOneObject(runes, &idx))
		}else if !unicode.IsSpace(runes[idx]){
			result = append(result, TakeOneWord(runes, &idx))
		}

		idx++
	}

	return result
}


func TakeOneWord(runes []rune, idx *int) string{
	var start = *idx
	for *idx < len(runes) {
		if unicode.IsSpace(runes[*idx]) {
			break
		}else if runes[*idx] == rune('(') {
			TakeOneParen(runes, idx)
			if *idx < len(runes) && runes[*idx] == rune(':'){
				(*idx)++
			}
			break
		}
		(*idx)++
	}

	return string(runes[start:*idx])
}

func TakeOneStr(runes []rune, idx *int) string{
	var start = *idx
	(*idx)++

	for *idx < len(runes) {
		if runes[*idx] == rune('^'){
			(*idx)++
		}else if runes[*idx] == rune('"'){
			(*idx)++
			break
		}
		(*idx)++
	}

	return string(runes[start:*idx])
}


func TakeOneBlock(runes []rune, idx *int) string{
	var start = *idx
	var bfloor = 0

	for *idx < len(runes) {
		if runes[*idx] == rune('[') {
			bfloor++
		}else if runes[*idx] == rune(']') {
			bfloor--
		}else if runes[*idx] == rune('"') {
			(*idx)++
			TakeOneStr(runes, idx)
			continue
		}

		(*idx)++

		if bfloor == 0 {
			break
		}
	}

	return string(runes[start:*idx])
}


func TakeOneParen(runes []rune, idx *int) string{
	var start = *idx
	var pfloor = 0

	for *idx < len(runes) {
		if runes[*idx] == rune('(') {
			pfloor++
		}else if runes[*idx] == rune(')') {
			pfloor--
		}else if runes[*idx] == rune('"') {
			(*idx)++
			TakeOneStr(runes, idx)
			continue
		}

		(*idx)++

		if pfloor == 0 {
			break
		}
	}

	return string(runes[start:*idx])
}


func TakeOneObject(runes []rune, idx *int) string{
	var start = *idx
	var ofloor = 0

	for *idx < len(runes) {
		if runes[*idx] == rune('{') {
			ofloor++
		}else if runes[*idx] == rune('}') {
			ofloor--
		}else if runes[*idx] == rune('"') {
			(*idx)++
			TakeOneStr(runes, idx)
			continue
		}

		(*idx)++

		if ofloor == 0 {
			break
		}
	}

	return string(runes[start:*idx])
}


func IsNumber(c byte) bool{
	if(c >= 48 && c <= 57){
		return true
	}
	return false
}

func IsNumberStr(s string) int{
	if(len(s) == 0){
		return -1
	}
	if(s[0] != '-' && !IsNumber(s[0]) || s== "-"){
		return -1
	}

	var dot = 0
	for idx:=1; idx<len(s); idx++ {
		if(!IsNumber(s[idx]) && s[idx] != '.'){
			return -1
		}
		if(s[idx] == '.'){
			dot += 1
		}
	}
	return dot
}

func StartWith(source string, target string) bool{
	if len(source) == 0 {
		return false
	}
	if len(target) == 0 {
		return true
	}
	if len(target) > len(source){
		return false
	}
	return source[0:len(target)] == target
}

func EndWith(source string, target string) bool{
	if len(source) == 0 {
		return false
	}
	if len(target) == 0 {
		return true
	}
	if len(target) > len(source){
		return false
	}
	return source[len(source) - len(target):] == target
}





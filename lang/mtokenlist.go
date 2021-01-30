package lang

import "bytes"
import "fmt"

type MtokenList struct {
	Room		uint
	EndIdx		uint
	Line		[]*Mtoken
}

func NewTks(size int) *MtokenList {
	var tks = &MtokenList{}
	tks.Room = uint(size) + 1
	tks.EndIdx = 0
	tks.Line = make([]*Mtoken, uint(size) + 1)
	return tks
}

func (tks *MtokenList) Init() {
	tks.EndIdx = 0
	tks.Room = 9
	tks.Line = make([]*Mtoken, uint(9))
}

func (tks *MtokenList) List() []*Mtoken{
	return tks.Line[0:tks.EndIdx]
}

func (tks *MtokenList) Get(idx int) *Mtoken{
	if idx < 0 || uint(idx) >= tks.EndIdx {
		return nil
	}
	return tks.Line[idx]
}

func (tks *MtokenList) Size() int{
	return int(tks.Room)
}

func (tks *MtokenList) Len() int{
	return int(tks.EndIdx)
}

func (tks *MtokenList) Resize(size uint) {
	tks.Room = size
	temp := make([]*Mtoken, size + 1)

	if tks.EndIdx > size {
		copy(temp, tks.Line[0 : size])
		tks.EndIdx = size
	}else{
		copy(temp, tks.Line[0 : tks.EndIdx])
	}
	tks.Line = temp
	
}

func (tks *MtokenList) Add(t *Mtoken) {
	if tks.Room == 0 || tks.EndIdx >= tks.Room - 1 {
		tks.Resize((tks.Room + 1) * 2)
	}
	tks.Line[tks.EndIdx] = t
	tks.EndIdx++
}

func (tks *MtokenList) Pop() {
	if tks.EndIdx > 0 {
		tks.EndIdx--
	}
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}
}

func (tks *MtokenList) Put(idx uint, t *Mtoken) {
	for idx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	if tks.Room > uint(len(tks.Line)) {
		tks.Resize(tks.Room)
	}

	tks.Line[idx] = t
	if idx >= tks.EndIdx {
		tks.EndIdx = idx + 1
	}
}

func (tks *MtokenList) Insert(idx int, t *Mtoken) {
	if tks.EndIdx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Mtoken, tks.Room + 1)

	if tks.EndIdx >= uint(idx) {
		copy(temp, tks.Line[0:idx])
		temp[idx] = t
		copy(temp[idx+1:], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx++
	}else{
		tks.Put(uint(idx), t)
	}
}

func (tks *MtokenList) InsertAll(idx int, list *MtokenList) {
	if tks.EndIdx + list.EndIdx >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Mtoken, tks.Room + 1)

	if tks.EndIdx >= uint(idx) {
		copy(temp, tks.Line[0:idx])
		copy(temp[idx:], list.Line[0:list.EndIdx])
		copy(temp[uint(idx)+list.EndIdx:], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx += list.EndIdx
	}else{
		copy(temp, tks.Line)
		copy(temp[idx:], list.Line[0:list.EndIdx])
		tks.Line = temp
		tks.EndIdx = uint(idx) + list.EndIdx
	}
}

func (tks *MtokenList) InsertArr(idx int, arr []*Mtoken) {
	if tks.EndIdx + uint(len(arr)) >= tks.Room - 1 {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Mtoken, tks.Room + 1)

	if tks.EndIdx > uint(idx) {
		copy(temp, tks.Line[0:idx])
		copy(temp[idx:], arr)
		copy(temp[idx+len(arr):], tks.Line[idx : tks.EndIdx])
		tks.Line = temp
		tks.EndIdx += uint(len(arr))
	}else{
		copy(temp, tks.Line)
		copy(temp[idx:], arr)
		tks.Line = temp
		tks.EndIdx = uint(idx) + uint(len(arr))
	}
}

func (tks *MtokenList) First() *Mtoken{
	if tks.EndIdx > 0 {
		return tks.Line[0]
	}
	return nil
}

func (tks *MtokenList) Last() *Mtoken{
	if tks.EndIdx > 0 {
		return tks.Line[tks.EndIdx - 1]
	}
	return nil
}

func (tks *MtokenList) AddAll(list *MtokenList) {
	for tks.Room < tks.EndIdx + list.EndIdx {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Mtoken, tks.Room + 1)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], list.Line[0:list.EndIdx])

	tks.EndIdx = tks.EndIdx + list.EndIdx
	tks.Line = temp
} 

func (tks *MtokenList) AddArr(arr []*Mtoken) {
	for tks.Room < tks.EndIdx + uint(len(arr)) {
		tks.Room = tks.Room * 2 + 1
	}
	temp := make([]*Mtoken, tks.Room + 1)
	copy(temp, tks.Line[0:tks.EndIdx])
	copy(temp[tks.EndIdx:], arr)

	tks.EndIdx = tks.EndIdx + uint(len(arr))
	tks.Line = temp
}

func (tks *MtokenList) PopFirst() {
	if tks.EndIdx <= 0 {
		return
	}
	temp := make([]*Mtoken, tks.Room)
	copy(temp, tks.Line[1:])
	tks.Line = temp
	tks.EndIdx--
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}
}

func (tks *MtokenList) ToString() string{
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i:=0; i<int(tks.EndIdx); i++ {
		buffer.WriteString(tks.Line[i].ToStr())
		buffer.WriteString(" ")
	}
	if len(buffer.Bytes()) > 1 {
		buffer.Bytes()[len(buffer.Bytes())-1] = ']'
	}else{
		buffer.WriteString("]")
	}
	return buffer.String()
}

func (tks *MtokenList) Echo() {
	fmt.Println(tks.ToString())
}

func (tks *MtokenList) Clone() *MtokenList{
	var result MtokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Mtoken, result.Room + 1)
	copy(result.Line, tks.Line)
	return &result
}

func (tks *MtokenList) CloneDeep() *MtokenList{
	var result MtokenList
	result.Room = tks.Room
	result.EndIdx = tks.EndIdx
	result.Line = make([]*Mtoken, result.Room + 1)
	for i:=0; i< int(tks.EndIdx); i++ {
		result.Line[i] = tks.Line[i].CloneDeep()
	}
	return &result
}


func (tks *MtokenList) Take(startIdx int, endIdx int) *MtokenList{
	var result = NewTks(8)
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > int(tks.EndIdx) {
		endIdx = int(tks.EndIdx)
	}
	result.AddArr(tks.Line[startIdx:endIdx])
	
	var part = endIdx - startIdx
	temp := make([]*Mtoken, tks.Room + 1)
	copy(temp, tks.Line[0:startIdx])
	copy(temp[startIdx:], tks.Line[endIdx:])
	tks.Line = temp
	tks.EndIdx -= uint(part)
	if tks.EndIdx < tks.Room / 5 - 1 {
		tks.Resize(tks.Room / 5)
	}

	return result
}

func (tks *MtokenList) Replace(old *Mtoken, new *Mtoken, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		if tks.Line[i].Tp == old.Tp && tks.Line[i].Val == old.Val && i >= 0 {
			tks.Line[i].Copy(new)
			amount--
		}
	}
}

func (tks *MtokenList) ReplacePart(old *MtokenList, new *MtokenList, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		var eq = true
		for j:=0; j<old.Len(); j++ {
			if tks.Line[i+j].Tp != old.Line[j].Tp || tks.Line[i+j].Val != old.Line[j].Val {
				eq = false
			}
		}
		if eq {
			tks.Take(i, i + old.Len())
			tks.InsertAll(i, new)
			amount--
		}
	}
}

func (tks *MtokenList) ReplacePartToOne(old *MtokenList, new *Mtoken, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		var eq = true
		for j:=0; j<old.Len(); j++ {
			if tks.Line[i+j].Tp != old.Line[j].Tp || tks.Line[i+j].Val != old.Line[j].Val {
				eq = false
			}
		}
		if eq {
			tks.Take(i, i + old.Len())
			tks.Insert(i, new)
			amount--
		}
	}
}

func (tks *MtokenList) ReplaceOneToPart(old *Mtoken, new *MtokenList, at int, amount int){
	if amount < 0 {
		amount = int(^uint(0) >> 1) //取有符号int最大值
	}

	for i:=at; i<tks.Len() && amount>0; i++ {
		if tks.Line[i].Tp == old.Tp && tks.Line[i].Val == old.Val {
			tks.Take(i, i + 1)
			tks.InsertAll(i, new)
			amount--
		}
	}
}


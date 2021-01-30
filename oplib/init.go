package oplib

import (
	. "github.com/NjinN/Mo_go/lang"
	. "github.com/NjinN/Mo_go/nativelib"
)

func InitOp(ctx *Mtoken){
	var table = ctx.Table()

	var addOp = InitAdd()
	addOp.Tp = TOP
	table.PutNow("+", addOp)

	var subOp = InitSub()
	subOp.Tp = TOP
	table.PutNow("-", subOp)

	var mulOp = InitMul()
	mulOp.Tp = TOP
	table.PutNow("*", mulOp)

	var divOp = InitDiv()
	divOp.Tp = TOP
	table.PutNow("/", divOp)

	var eqOp = InitEq()
	eqOp.Tp = TOP
	table.PutNow("=", eqOp)

	var ltOp = InitLt()
	ltOp.Tp = TOP
	table.PutNow("<", ltOp)

	var gtOp = InitGt()
	gtOp.Tp = TOP
	table.PutNow(">", gtOp)

	var leOp = InitLe()
	leOp.Tp = TOP
	table.PutNow("<=", leOp)

	var geOp = InitGe()
	geOp.Tp = TOP
	table.PutNow(">=", geOp)

}
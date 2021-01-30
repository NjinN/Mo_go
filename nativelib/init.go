package nativelib

import (
	. "github.com/NjinN/Mo_go/lang"
)

func InitNative(ctx *Mtoken){
	var table = ctx.Table()

	table.PutNow("add", InitAdd())
	table.PutNow("sub", InitSub())
	table.PutNow("mul", InitMul())
	table.PutNow("div", InitDiv())

	table.PutNow("eq", InitEq())
	table.PutNow("lt", InitLt())
	table.PutNow("gt", InitGt())
	table.PutNow("le", InitLe())
	table.PutNow("ge", InitGe())

	table.PutNow("quit", initQuit())
	table.PutNow("q", initQuit())
	table.PutNow("clear", initClear())

	table.PutNow("print", initPrint())

	table.PutNow("reduce", initReduce())

	table.PutNow("if", initIf())
	table.PutNow("either", initEither())
}







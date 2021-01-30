package lang

const (
	TNIL 		= iota
	TNONE
	TERR
	TDATATYPE
	TBOOL
	TBYTE
	TCHAR
	TINT
	TFLOAT
	TSTRING
	TBLOCK
	TPAREN
	TOBJECT
	TTABLE
	TWORD
	TCALL
	TLIT_WORD
	TSET_WORD
	TSET_FUNC
	TNATIVE
	TFUNC
	TOP


	

	TUNDEFINED
)


func Tp2Str(n uint8) string{
	switch n {
	case TNIL:
		return "nil!"
	case TNONE:
		return "none!"
	case TERR:
		return "err!"
	case TDATATYPE:
		return "datetype!"
	case TBOOL:
		return "bool!"
	case TBYTE:
		return "byte!"
	case TCHAR:
		return "char!"
	case TINT:
		return "int!"
	case TFLOAT:
		return "float!"
	case TSTRING:
		return "string!"
	case TBLOCK:
		return "block!"
	case TPAREN:
		return "paren!"
	case TOBJECT:
		return "object!"
	case TTABLE:
		return "table!"
	case TWORD:
		return "word!"
	case TCALL:
		return "call!"
	case TLIT_WORD:
		return "lit_word!"
	case TSET_WORD:
		return "set_word!"
	case TSET_FUNC:
		return "set_func!"
	case TNATIVE:
		return "native!"
	case TFUNC:
		return "func!"
	case TOP:
		return "op!"


	default:
		return "undefined!"
	}
}



func Str2Tp(s string) uint8{
	switch s {
	case "nil!":
		return TNIL
	case "none!":
		return TNONE
	case "err!":
		return TERR
	case "datetype!":
		return TDATATYPE
	case "bool!":
		return TBOOL
	case "byte!":
		return TBYTE
	case "char!":
		return TCHAR
	case "int!":
		return TINT
	case "float!":
		return TFLOAT
	case "string!":
		return TSTRING
	case "block!":
		return TBLOCK
	case "paren!":
		return TPAREN
	case "object!":
		return TOBJECT
	case "table!":
		return TTABLE
	case "word!":
		return TWORD
	case "call!":
		return TCALL
	case "lit_word!":
		return TLIT_WORD
	case "set_word!":
		return TSET_WORD
	case "set_func!":
		return TSET_FUNC
	case "native!":
		return TNATIVE
	case "func!":
		return TFUNC
	case "op!":
		return TOP


	default:
		return TUNDEFINED
	}
}





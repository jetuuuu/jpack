package types

type FieldType int

const (
	Builtin FieldType = iota
	Float32
	Float32_P
	Int
	Int_P
	Bool
	Bool_P
	Byte
	Byte_P
	Float64
	Float64_P
	Int16
	Int16_P
	Int32
	Int32_P
	Int64
	Int64_P
	Int8
	Int8_P
	String
	String_P
	Uint
	Uint_P
	Uint16
	Uint16_P
	Uint32
	Uint32_P
	Uint64
	Uint64_P
	Uint8
	Uint8_P
	Time
	Time_P
	Map
	Map_P

	Internal
)

func (f FieldType) String() string {
	switch f {
	case Builtin:
		return "builtin"
	case Float32:
		return "float32"
	case Int:
		return "int"
	case Bool:
		return "bool"
	case Byte:
		return "byte"
	case Float64:
		return "float64"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Int8:
		return "int8"
	case String:
		return "string"
	case Uint:
		return "uint"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	case Uint8:
		return "uint8"
	case Internal:
		return "internal"
	default:
		return ""
	}
}
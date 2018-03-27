package pack

import (
	"github.com/jetuuuu/jpack/types"
	"bytes"
	"text/template"
	"github.com/jetuuuu/jpack/field"
)

type F struct {
	encode func(v field.FieldInfo) string
	decode func(v field.FieldInfo) string
	size uint
}


var typeToFunc = map[types.FieldType]F{
	types.Bool: F{packBool, unpackBool, 1},
	types.Bool_P: F{packPointerToBool, unpackPointerToBool, 1},

	types.Int8: {pack8BitsNum, unpack8BitsNum, 1},
	types.Int8_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum, 1},
	types.Uint8: {pack8BitsNum, unpack8BitsNum, 1},
	types.Uint8_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum, 1},
	types.Byte: {pack8BitsNum, unpack8BitsNum, 1},
	types.Byte_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum, 1},

	types.Int16: {pack16BitsNum, unpack16BitsNum, 2},
	types.Int16_P: {packPointerTo16BitsNum, unpackPointerTo16BitsNum, 2},
	types.Uint16: {pack16BitsNum, unpack16BitsNum, 2},
	types.Uint16_P: {packPointerTo16BitsNum, unpackPointerTo16BitsNum, 2},

	types.Int32: {pack32BitsNum, unpack32BitsNum, 4},
	types.Int32_P: {packPointerTo32BitsNum, unpackPointerTo32BitsNum, 4},
	types.Uint32: {pack32BitsNum, unpack32BitsNum, 4},
	types.Uint32_P: {packPointerTo32BitsNum, unpackPointerTo32BitsNum, 4},

	types.Int64: {pack64BitsNum, unpack64BitsNum, 8},
	types.Int64_P: {packPointerTo64BitsNum, unpackPointerTo64BitsNum, 8},
	types.Uint64: {pack64BitsNum, unpack32BitsNum, 8},
	types.Uint64_P: {packPointerTo64BitsNum, unpackPointerTo64BitsNum, 8},

	types.Float32: {packFloat32, unpackFloat32, 4},
	types.Float32_P: {packPointerToFloat32, unpackPointerToFloat32, 4},

	types.Float64: {packFloat64, unpackFloat64, 8},
	types.Float64_P: {packPointerToFloat64, unpackPointerToFloat64, 8},

	types.String: {packString, unpackString, 4},
	types.Time: {packTime, unpackTime, 15},
}


func useTmpl(tmpl string, v field.FieldInfo) string {
	w := bytes.NewBuffer(nil)
	t := template.Must(template.New("template").Parse(tmpl))
	t.Execute(w, v)

	return w.String()
}

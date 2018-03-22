package pack

import (
	"github.com/jetuuuu/jpack/types"
	"bytes"
	"text/template"
)

type F struct {
	encode func(v FieldInfo) string
	decode func(v FieldInfo) string
}


var typeToFunc = map[types.FieldType]F{
	types.Bool: F{packBool, unpackBool},
	types.Bool_P: F{packPointerToBool, unpackPointerToBool},

	types.Int8: {pack8BitsNum, unpack8BitsNum},
	types.Int8_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum},
	types.Uint8: {pack8BitsNum, unpack8BitsNum},
	types.Uint8_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum},
	types.Byte: {pack8BitsNum, unpack8BitsNum},
	types.Byte_P: {packPointerTo8BitsNum, unpackPointerTo8BitsNum},

	types.Int16: {pack16BitsNum, unpack16BitsNum},
	types.Int16_P: {packPointerTo16BitsNum, unpackPointerTo16BitsNum},
	types.Uint16: {pack16BitsNum, unpack16BitsNum},
	types.Uint16_P: {packPointerTo16BitsNum, unpackPointerTo16BitsNum},

	types.Int32: {pack32BitsNum, unpack32BitsNum},
	types.Int32_P: {packPointerTo32BitsNum, unpackPointerTo32BitsNum},
	types.Uint32: {pack32BitsNum, unpack32BitsNum},
	types.Uint32_P: {packPointerTo32BitsNum, unpackPointerTo32BitsNum},

	types.Int64: {pack64BitsNum, unpack64BitsNum},
	types.Int64_P: {packPointerTo64BitsNum, unpackPointerTo64BitsNum},
	types.Uint64: {pack64BitsNum, unpack32BitsNum},
	types.Uint64_P: {packPointerTo64BitsNum, unpackPointerTo64BitsNum},

	types.Float32: {packFloat32, unpackFloat32},
	types.Float32_P: {packPointerToFloat32, unpackPointerToFloat32},

	types.Float64: {packFloat64, unpackFloat64},
	types.Float64_P: {packPointerToFloat64, unpackPointerToFloat64},

	types.String: {packString, unpackString},
}


func useTmpl(tmpl string, v FieldInfo) string {
	w := bytes.NewBuffer(nil)
	t := template.Must(template.New("template").Parse(tmpl))
	t.Execute(w, v)

	return w.String()
}

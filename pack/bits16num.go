package pack

import "github.com/jetuuuu/jpack/field"

const bits16TmplW = `
{
	b[offset] = byte({{.Name}})
	b[offset + 1] = byte({{.Name}} >> 8)

	offset += 2
}
`

const bits16TmplR = `
{
	{{.Name}} = {{.Type}}(uint16(b[offset]) | uint16(b[offset + 1])<<8)
	offset += 2
}
`

const bits16PTmplR = `
{
	value := {{.Type}}(uint16(b[offset]) | uint16(b[offset + 1])<<8)
	{{.Name}} = &value
	offset += 2
}
`

func pack16BitsNum(f field.FieldInfo) string {
	return useTmpl(bits16TmplW, f)
}

func packPointerTo16BitsNum(f field.FieldInfo) string {
	return pack16BitsNum(f)
}

func unpack16BitsNum(f field.FieldInfo) string {
	return useTmpl(bits16TmplR, f)
}

func unpackPointerTo16BitsNum(f field.FieldInfo) string {
	return useTmpl(bits16PTmplR, f)
}
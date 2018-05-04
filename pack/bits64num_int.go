package pack

import "github.com/jetuuuu/jpack/field"

const bits64TmplW = `
{
	b[offset] = byte({{.Name}})
	b[offset+1] = byte({{.Name}} >> 8)
	b[offset+2] = byte({{.Name}} >> 16)
	b[offset+3] = byte({{.Name}} >> 24)
	b[offset+4] = byte({{.Name}} >> 32)
	b[offset+5] = byte({{.Name}} >> 40)
	b[offset+6] = byte({{.Name}} >> 48)
	b[offset+7] = byte({{.Name}} >> 56)

	offset += 8
}
`

const bits64TmplR = `
{
	{{.Name}} = {{.Type}}(uint64(b[offset]) | uint64(b[offset+1])<<8 | uint64(b[offset+2])<<16 | uint64(b[offset+3])<<24 | uint64(b[offset+4])<<32 | uint64(b[offset+5])<<40 | uint64(b[offset+6])<<48 | uint64(b[offset+7])<<56)
	offset += 8
}
`

const bits64PTmplR = `
{
	value := {{.Type}}(uint64(b[offset]) | uint64(b[offset+1])<<8 | uint64(b[offset+2])<<16 | uint64(b[offset+3])<<24 | uint64(b[offset+4])<<32 | uint64(b[offset+5])<<40 | uint64(b[offset+6])<<48 | uint64(b[offset+7])<<56)
	{{.Name}} = &value
	offset += 8
}
`

func pack64BitsNum(f field.FieldInfo) string {
	return useTmpl(bits64TmplW, f)
}

func packPointerTo64BitsNum(f field.FieldInfo) string {
	return pack64BitsNum(f)
}

func unpack64BitsNum(f field.FieldInfo) string {
	return useTmpl(bits64TmplR, f)
}

func unpackPointerTo64BitsNum(f field.FieldInfo) string {
	return useTmpl(bits64PTmplR, f)
}

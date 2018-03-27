package pack

import "github.com/jetuuuu/jpack/field"

const bits32TmplW = `
{
    b[offset] = byte({{.Name}})
	b[offset + 1] = byte({{.Name}} >> 8)
	b[offset + 2] = byte({{.Name}} >> 16)
	b[offset + 3] = byte({{.Name}} >> 24)

	offset += 4
}
`

const bits32TmplR = `
{
	{{.Name}} = {{.Type}}(uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24)
	offset += 4
}
`

const bits32PTmplR = `
{
	value := {{.Type}}(uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24)
	{{.Name}} = &value
	offset += 4
}
`

func pack32BitsNum(f field.FieldInfo) string {
	return useTmpl(bits32TmplW, f)
}

func packPointerTo32BitsNum(f field.FieldInfo) string {
	return pack32BitsNum(f)
}

func unpack32BitsNum(f field.FieldInfo) string {
	return useTmpl(bits32TmplR, f)
}

func unpackPointerTo32BitsNum(f field.FieldInfo) string {
	return useTmpl(bits32PTmplR, f)
}

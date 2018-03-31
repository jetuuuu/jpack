package pack

import "github.com/jetuuuu/jpack/field"

const float64W = `
{
	f := math.Float64bits({{.Name}})

	b[offset] = byte(f)
	b[offset+1] = byte(f >> 8)
	b[offset+2] = byte(f >> 16)
	b[offset+3] = byte(f >> 24)
	b[offset+4] = byte(f >> 32)
	b[offset+5] = byte(f >> 40)
	b[offset+6] = byte(f >> 48)
	b[offset+7] = byte(f >> 56)

	offset += 8
}
`

const float64R = `
{
	f := uint64(b[offset]) | uint64(b[offset+1])<<8 | uint64(b[offset+2])<<16 | uint64(b[offset+3])<<24 | uint64(b[offset+4])<<32 | uint64(b[offset+5])<<40 | uint64(b[offset+6])<<48 | uint64(b[offset+7])<<56
	offset += 8

	{{.Name}} = math.Float64frombits(f)
}
`

const float64PR = `
{
	f := uint64(b[offset]) | uint64(b[offset+1])<<8 | uint64(b[offset+2])<<16 | uint64(b[offset+3])<<24 | uint64(b[offset+4])<<32 | uint64(b[offset+5])<<40 | uint64(b[offset+6])<<48 | uint64(b[offset+7])<<56
	offset += 8

	value := math.Float64frombits(f)
	{{.Name}} = &value
}
`

func packFloat64(f field.FieldInfo) string {
	return useTmpl(float64W, f)
}

func packPointerToFloat64(f field.FieldInfo) string {
	return packFloat64(f)
}

func unpackFloat64(f field.FieldInfo) string {
	return useTmpl(float64R, f)
}

func unpackPointerToFloat64(f field.FieldInfo) string {
	return useTmpl(float64PR, f)
}
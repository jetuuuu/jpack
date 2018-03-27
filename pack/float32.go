package pack

import "github.com/jetuuuu/jpack/field"

const float32W = `
{
	f := math.Float32bits({{.Name}})

    b[offset] = byte(f)
	b[offset + 1] = byte(f >> 8)
	b[offset + 2] = byte(f >> 16)
	b[offset + 3] = byte(f >> 24)

	offset += 4
}
`

const float32R = `
{
	f := uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24
	offset += 4

	{{.Name}} = math.Float32frombits(f)
}
`

const float32PR = `
{
	f := uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24
	offset += 4

	value := math.Float32frombits(f)
	{{.Name}} = &value
}
`

func packFloat32(f field.FieldInfo) string {
	return useTmpl(float32W, f)
}

func packPointerToFloat32(f field.FieldInfo) string {
	return packFloat32(f)
}

func unpackFloat32(f field.FieldInfo) string {
	return useTmpl(float32R, f)
}

func unpackPointerToFloat32(f field.FieldInfo) string {
	return useTmpl(float32PR, f)
}
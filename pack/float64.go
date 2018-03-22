package pack

const importFloat = "import \"math\""

const float64W = `
{
	f := math.Float64bits({{.V}})

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
	f := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	offset += 8

	{{.V}} := math.Float64frombits(f)
}
`

const float64PR = `
{
	f := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	offset += 8

	value := math.Float64frombits(f)
	{{.V}} = &value
}
`

func packFloat64(f FieldInfo) string {
	return useTmpl(float64W, f)
}

func packPointerToFloat64(f FieldInfo) string {
	return packFloat64(f)
}

func unpackFloat64(f FieldInfo) string {
	return useTmpl(float64R, f)
}

func unpackPointerToFloat64(f FieldInfo) string {
	return useTmpl(float64PR, f)
}
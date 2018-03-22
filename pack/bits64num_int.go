package pack

const bits64TmplW = `
{
	b[offset] = byte({{.V}})
	b[offset+1] = byte({{.V}} >> 8)
	b[offset+2] = byte({{.V}} >> 16)
	b[offset+3] = byte({{.V}} >> 24)
	b[offset+4] = byte({{.V}} >> 32)
	b[offset+5] = byte({{.V}} >> 40)
	b[offset+6] = byte({{.V}} >> 48)
	b[offset+7] = byte({{.V}} >> 56)

	offset += 8
}
`

const bits64TmplR = `
{
	{{.V}} = {{.Type}}(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
	offset += 8
}
`

const bits64PTmplR = `
{
	value := {{.Type}}(uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 | uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56)
	{{.V}} = &value
	offset += 8
}
`

func pack64BitsNum(f FieldInfo) string {
	return useTmpl(bits64TmplW, f)
}

func packPointerTo64BitsNum(f FieldInfo) string {
	return pack64BitsNum(f)
}

func unpack64BitsNum(f FieldInfo) string {
	return useTmpl(bits64TmplR, f)
}

func unpackPointerTo64BitsNum(f FieldInfo) string {
	return useTmpl(bits64PTmplR, f)
}

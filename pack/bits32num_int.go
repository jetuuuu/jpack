package pack

const bits32TmplW = `
{
    b[offset] = byte({{.V}})
	b[offset + 1] = byte({{.V}} >> 8)
	b[offset + 2] = byte({{.V}} >> 16)
	b[offset + 3] = byte({{.V}} >> 24)

	offset += 4
}
`

const bits32TmplR = `
{
	{{.V}} = {{.Type}}(uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24)
	offset += 4
}
`

const bits32PTmplR = `
{
	value := {{.Type}}(uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24)
	{{.V}} = &value
	offset += 4
}
`

func pack32BitsNum(f FieldInfo) string {
	return useTmpl(bits32TmplW, f)
}

func packPointerTo32BitsNum(f FieldInfo) string {
	return pack32BitsNum(f)
}

func unpack32BitsNum(f FieldInfo) string {
	return useTmpl(bits32TmplR, f)
}

func unpackPointerTo32BitsNum(f FieldInfo) string {
	return useTmpl(bits32PTmplR, f)
}

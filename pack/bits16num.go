package pack

const bits16TmplW = `
{
	b[offset] = byte({{.V}})
	b[offset + 1] = byte({{.V}} >> 8)

	offset += 2
}
`

const bits16TmplR = `
{
	{{.V}} = {{.Type}}(uint16(b[offset]) | uint16(b[offset + 1])<<8)
	offset += 2
}
`

const bits16PTmplR = `
{
	value := {{.Type}}(uint16(b[offset]) | uint16(b[offset + 1])<<8)
	{{.V}} = &value
	offset += 2
}
`

func pack16BitsNum(f FieldInfo) string {
	return useTmpl(bits16TmplW, f)
}

func packPointerTo16BitsNum(f FieldInfo) string {
	return pack16BitsNum(f)
}

func unpack16BitsNum(f FieldInfo) string {
	return useTmpl(bits16TmplR, f)
}

func unpackPointerTo16BitsNum(f FieldInfo) string {
	return useTmpl(bits16PTmplR, f)
}
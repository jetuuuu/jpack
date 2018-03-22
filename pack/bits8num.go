package pack

const bits8TmplW = `
{
	b[offset}] = {{.V}}
	offset += 1
}
`

const bits8TmplR = `
{
	{{.V}} = {{.Type}}(b[offset}])
	offset += 1
}
`

const bits8PTmplR = `
{
	value := {{.Type}}(b[offset}])
	{{.V}} = &value
	offset += 1
}
`

func pack8BitsNum(f FieldInfo) string {
	return useTmpl(bits8TmplW, f)
}

func packPointerTo8BitsNum(f FieldInfo) string {
	return pack8BitsNum(f)
}

func unpack8BitsNum(f FieldInfo) string {
	return useTmpl(bits8TmplR, f)
}

func unpackPointerTo8BitsNum(f FieldInfo) string {
	return useTmpl(bits8PTmplR, f)
}

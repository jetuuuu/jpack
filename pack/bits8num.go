package pack

import "github.com/jetuuuu/jpack/field"

const bits8TmplW = `
{
	b[offset}] = {{.Name}}
	offset += 1
}
`

const bits8TmplR = `
{
	{{.Name}} = {{.Type}}(b[offset}])
	offset += 1
}
`

const bits8PTmplR = `
{
	value := {{.Type}}(b[offset}])
	{{.Name}} = &value
	offset += 1
}
`

func pack8BitsNum(f field.FieldInfo) string {
	return useTmpl(bits8TmplW, f)
}

func packPointerTo8BitsNum(f field.FieldInfo) string {
	return pack8BitsNum(f)
}

func unpack8BitsNum(f field.FieldInfo) string {
	return useTmpl(bits8TmplR, f)
}

func unpackPointerTo8BitsNum(f field.FieldInfo) string {
	return useTmpl(bits8PTmplR, f)
}

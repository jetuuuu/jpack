package pack

import "github.com/jetuuuu/jpack/field"

const boolTmplW = `
{
	if {{.Name}} {
		b[offset] = 1
	} else {
		b[offset] = 0
	}

	offset += 1
}
`

const boolTmplR = `
{
	{{.Name}} = b[offset] == 1

	offset += 1
}
`

const boolPTmplR = `
{
	flag := b[offset] == 1
	{{.Name}} = &flag

	offset += 1
}
`
func packBool(f field.FieldInfo) string {
	return useTmpl(boolTmplW, f)
}

func unpackBool(f field.FieldInfo) string {
	return useTmpl(boolTmplR, f)
}

func packPointerToBool(f field.FieldInfo) string {
	f.Name = "*" + f.Name
	return packBool(f)
}

func unpackPointerToBool(f field.FieldInfo) string {
	return useTmpl(boolPTmplR, f)
}

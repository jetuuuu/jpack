package pack

const boolTmplW = `
{
	if {{.V}} {
		b[offset] = 1
	} else {
		b[offset] = 0
	}

	offset += 1
}
`

const boolTmplR = `
{
	{{.V}} = b[offset] == 1

	offset += 1
}
`

const boolPTmplR = `
{
	flag := b[offset] == 1
	{{.V}} = &flag

	offset += 1
}
`
func packBool(f FieldInfo) string {
	return useTmpl(boolTmplW, f)
}

func unpackBool(f FieldInfo) string {
	return useTmpl(boolTmplR, f)
}

func packPointerToBool(f FieldInfo) string {
	f.Name = "*" + f.Name
	return packBool(f)
}

func unpackPointerToBool(f FieldInfo) string {
	return useTmpl(boolPTmplR, f)
}

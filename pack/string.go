package pack

const stringTmplW = `
{
    _string_as_byes_jpack := []byte({{.V}})
	_len_string_as_byes_jpack := len(_string_as_byes_jpack)

    b[offset] = byte(_len_string_as_byes_jpack)
	b[offset + 1] = byte(_len_string_as_byes_jpack >> 8)
	b[offset + 2] = byte(_len_string_as_byes_jpack >> 16)
	b[offset + 3] = byte(_len_string_as_byes_jpack >> 24)


	copy(b[offset + 4:], _string_as_byes_jpack)
	offset += 4 + _len_string_as_byes_jpack
}
`

const stringTmplR = `
{
    _len_string_as_byes_jpack = int(uint32(b[offset]) | uint32(b[offset+1])<<8 | uint32(b[offset+2])<<16 | uint32(b[offset+3])<<24)
	offset += 4

	{{.V}} = string(b[offset:offset+_len_string_as_byes_jpack])
	offset += _len_string_as_byes_jpack
}
`

func packString(f FieldInfo) string {
	return useTmpl(stringTmplW, f)
}

func unpackString(f FieldInfo) string {
	return useTmpl(stringTmplR, f)
}

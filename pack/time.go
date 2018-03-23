package pack

const timeTmplW = `
{
	var timeBinary []byte
	timeBinary, err = {{.V}}.MarshalBinary()
	if err != nil {
		return nil, err
	}

	copy(b[offset:], timeBinary)
	offset += 15
}
`

const timeTmplR = `
{
	err = {{.V}}.UnmarshalBinary(b[offset:offset + 15])
	if err != nil {
		return nil, err
	}
	offset += 15
}
`

func packTime(f FieldInfo) string {
	return useTmpl(timeTmplW, f)
}

func unpackTime(f FieldInfo) string {
	return useTmpl(timeTmplR, f)
}
package pack

import "github.com/jetuuuu/jpack/field"

const timeTmplW = `
{
	var timeBinary []byte
	timeBinary, err := {{.Name}}.MarshalBinary()
	if err != nil {
		return nil, err
	}

	copy(b[offset:], timeBinary)
	offset += 15
}
`

const timeTmplR = `
{
	err := {{.Name}}.UnmarshalBinary(b[offset:offset + 15])
	if err != nil {
		return err
	}
	offset += 15
}
`

func packTime(f field.FieldInfo) string {
	return useTmpl(timeTmplW, f)
}

func unpackTime(f field.FieldInfo) string {
	return useTmpl(timeTmplR, f)
}
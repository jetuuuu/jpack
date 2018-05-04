package pack

import "github.com/jetuuuu/jpack/field"

const timeDurationTmplW = `
{
	nanoseconds := {{.Name}}.Nanoseconds()
	b[offset] = byte(nanoseconds)
	b[offset+1] = byte(nanoseconds >> 8)
	b[offset+2] = byte(nanoseconds >> 16)
	b[offset+3] = byte(nanoseconds >> 24)
	b[offset+4] = byte(nanoseconds >> 32)
	b[offset+5] = byte(nanoseconds >> 40)
	b[offset+6] = byte(nanoseconds >> 48)
	b[offset+7] = byte(nanoseconds >> 56)

	offset += 8
}
`

const timeDurationTmplR = `
{
	{{.Name}} = time.Duration(uint64(b[offset]) | uint64(b[offset+1])<<8 | uint64(b[offset+2])<<16 | uint64(b[offset+3])<<24 | uint64(b[offset+4])<<32 | uint64(b[offset+5])<<40 | uint64(b[offset+6])<<48 | uint64(b[offset+7])<<56)
	offset += 8
}
`

func packTimeDuration(f field.FieldInfo) string {
	return useTmpl(timeDurationTmplW, f)
}

func unpackTimeDuration(f field.FieldInfo) string {
	return useTmpl(timeDurationTmplR, f)
}

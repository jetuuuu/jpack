package pack

import (
	"text/template"
	"bytes"
	"fmt"
	"github.com/jetuuuu/jpack/types"
)

type FieldInfo struct {
	Name string
	Type types.FieldType
}

type Pack struct {
	structName string
	objectName string
	fieldToType []FieldInfo
	b *bytes.Buffer
}

func New(sName string, field2Type []FieldInfo) Pack {
	return Pack{
		structName: sName,
		objectName: "_jpack_obj_" + sName,
		fieldToType: field2Type,
		b: bytes.NewBuffer(nil),
	}
}

func (p Pack) Generate() string {
	p.generateSizeFunction()
	p.generateMarshalFunction()
	p.generateUnmarshalFunction()
	return p.b.String()
}

func (p Pack) generateSizeFunction() {
	fmt.Fprintln(p.b, "//calcualte structure size")
	fmt.Fprintf(p.b, "func (%s *%s) Size() uint64 {\n", p.objectName, p.structName)

	fmt.Fprintln(p.b, "var size uint64")
	for _, pair := range p.fieldToType {
		fmt.Fprintf(p.b, "//size for \"%s\" field\n", pair.Name)
		switch pair.Type {
		case types.String:
			fmt.Fprintln(p.b, "size += 4")
			fmt.Fprintf(p.b, "size += uint64(len(%s.%s))\n", p.objectName, pair.Name)
		case types.Int8, types.Bool, types.Uint8, types.Byte, types.Int8_P, types.Uint8_P, types.Byte_P, types.Bool_P:
			fmt.Fprintln(p.b, "size += 1")
		case types.Int32, types.Uint32, types.Float32, types.Int32_P, types.Uint32_P, types.Float32_P:
			fmt.Fprintln(p.b, "size += 4")
		}
	}

	fmt.Fprintln(p.b, "return size")
	fmt.Fprintln(p.b, "}")
	fmt.Fprintln(p.b, "\n")
}

func (p Pack) generateMarshalFunction() {
	fmt.Fprintln(p.b, "//marshal struct " + p.structName + " to bytes")
	fmt.Fprintf(p.b, "func (%s *%s) Marshal() ([]byte, error) {\n", p.objectName, p.structName)
	fmt.Fprintf(p.b, "b := make([]byte, %s.Size())\n", p.objectName)

	fmt.Fprintln(p.b, "offset := 0")
	for _, pair := range p.fieldToType {
		v := p.objectName + "." + pair.Name

		fmt.Fprintln(p.b, typeToFunc[pair.Type].encode(pair))
	}
	fmt.Fprintln(p.b, "return b, nil")
	fmt.Fprintln(p.b, "}")
	fmt.Fprintln(p.b, "\n")
}

func (p Pack) generateUnmarshalFunction() {
	fmt.Fprintln(p.b, "//unmarshal struct " + p.structName + " to bytes")
	fmt.Fprintf(p.b, "func (%s *%s) Unmarshal(b []byte) {\n", p.objectName, p.structName)

	fmt.Fprintln(p.b, "offset := 0")
	for _, pair := range p.fieldToType {
		v := p.objectName + "." + pair.Name

		fmt.Fprintln(p.b, typeToFunc[pair.Type].decode(pair))
	}

	fmt.Fprintln(p.b, "}")
	fmt.Fprintln(p.b, "\n")
}
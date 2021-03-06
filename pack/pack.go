package pack

import (
	"bytes"
	"fmt"
	"github.com/jetuuuu/jpack/types"
	"github.com/jetuuuu/jpack/field"
)

type Pack struct {
	pkg string
	structName string
	objectName string
	fieldToType []field.FieldInfo
	b *bytes.Buffer
}

func New(pkg, sName string, field2Type []field.FieldInfo) Pack {
	return Pack{
		pkg: pkg,
		structName: sName,
		objectName: "_jpack_obj_" + sName,
		fieldToType: field2Type,
		b: bytes.NewBuffer(nil),
	}
}

func (p Pack) Generate() string {
	p.generateHeader()
	p.generateSizeFunction()
	p.generateMarshalFunction()
	p.generateUnmarshalFunction()

	return p.b.String()
}

func (p Pack) generateHeader() {
	fmt.Fprintln(p.b, "package " + p.pkg)
	fmt.Fprintln(p.b, "import \"math\"")
	fmt.Fprintln(p.b, "var _ = math.NaN()\n")
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
		default:
			fmt.Fprintf(p.b, "size += %d\n", typeToFunc[pair.Type].size)
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
		pair.Name = p.objectName + "." + pair.Name
		fmt.Println(pair)
		f, ok := typeToFunc[pair.Type]
		if ok {
			fmt.Fprintln(p.b, f.encode(pair))
		}

	}
	fmt.Fprintln(p.b, "return b, nil")
	fmt.Fprintln(p.b, "}")
	fmt.Fprintln(p.b, "\n")
}

func (p Pack) generateUnmarshalFunction() {
	fmt.Fprintln(p.b, "//unmarshal struct " + p.structName + " to bytes")
	fmt.Fprintf(p.b, "func (%s *%s) Unmarshal(b []byte) error {\n", p.objectName, p.structName)

	fmt.Fprintln(p.b, "offset := 0")
	for _, pair := range p.fieldToType {
		pair.Name = p.objectName + "." + pair.Name
		f, ok := typeToFunc[pair.Type]
		if ok {
			fmt.Fprintln(p.b, f.decode(pair))
		}
	}

	fmt.Fprintln(p.b, "return nil")
	fmt.Fprintln(p.b, "}")
	fmt.Fprintln(p.b, "\n")
}
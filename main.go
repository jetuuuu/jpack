package main

import (
	"fmt"
	"go/token"
	"go/ast"
	"go/parser"
	"github.com/jetuuuu/jpack/field"
	"github.com/jetuuuu/jpack/types"
)

func main() {
	fset := token.NewFileSet()
	path := "task.go"
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}
	v := &visitor{structures: make(map[string][]field.FieldInfo)}
	ast.Walk(v, f)
	fmt.Println(v.structures)
}

type visitor struct {
	imports []string
	structures map[string][]field.FieldInfo
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return v
	}

	switch n := node.(type) {
	case *ast.ImportSpec:
		v.imports = append(v.imports, n.Path.Value)
	case *ast.GenDecl:
		if n.Tok != token.TYPE {
			break
		}
		ts := n.Specs[0].(*ast.TypeSpec)
		fields := ts.Type.(*ast.StructType).Fields
		for _, l := range fields.List {
			lt := l.Type
			indent, ok := lt.(*ast.Ident)
			if !ok {
				continue
			}
			for _, n := range l.Names {
				v.structures[ts.Name.Name] = append(v.structures[ts.Name.Name], field.FieldInfo{Name:n.Name, Type: types.FromString(indent.Name)})
			}
		}
	case *ast.Comment:
		fmt.Println("Comm: ", n.Text)
	}

	return v
}
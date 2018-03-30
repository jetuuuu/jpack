package main

import (
	"fmt"
	"go/token"
	"go/ast"
	"go/parser"
	"github.com/jetuuuu/jpack/field"
	"github.com/jetuuuu/jpack/types"
	"github.com/jetuuuu/jpack/pack"
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
	p := pack.New("A", v.structures["A"])
	fmt.Println(p.Generate())
}

type visitor struct {
	imports []Import
	structures map[string][]field.FieldInfo
}

type Import struct {
	Path string
	Alias string
}

func findImport(is []Import, alias string) (Import, bool) {
	for _, i := range is {
		if i.Alias == alias {
			return i, true
		}
	}

	return Import{}, false
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return v
	}

	switch n := node.(type) {
	case *ast.ImportSpec:
		name := ""
		if n.Name != nil {
			name = n.Name.Name
		}
		v.imports = append(v.imports, Import{Path: n.Path.Value, Alias: name})
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
				selExpt, ok  := lt.(*ast.SelectorExpr)
				if ok {
					indent, ok = selExpt.X.(*ast.Ident)
					if ok {
						if selExpt.Sel.Name == "Time" {
							if _, ok := findImport(v.imports, indent.Name); ok || indent.Name == "time" {
								indent.Name = "time"
							}
						}
					}
				}
				if indent == nil {
					continue
				}
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
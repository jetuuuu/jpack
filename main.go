package main

import (
	"fmt"
	"go/token"
	"go/ast"
	"go/parser"
	"github.com/jetuuuu/jpack/field"
	"github.com/jetuuuu/jpack/types"
	"github.com/jetuuuu/jpack/pack"
	"os"
	"flag"
	"path"
)


func main() {
	pkgPtr := flag.String("pkg", "", "pkg name")
	flag.Parse()

	if pkgPtr == nil || *pkgPtr == "" {
		fmt.Println("Pkg must be not empty string")
		return
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, *pkgPtr, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for name, pkg := range pkgs {
		v := &visitor{structures: make(map[string][]field.FieldInfo)}
		ast.Walk(v, pkg)
		fmt.Println(v.structures)
		for structName, s := range v.structures {
			p := pack.New(name, structName, s)
			file, _ := os.Create(  path.Join(*pkgPtr, name + "_" + structName + "_jpack_generated.go"))
			defer file.Close()
			file.WriteString(p.Generate())
		}
	}
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
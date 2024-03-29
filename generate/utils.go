package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RemoveGeneratedCode(dir string) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.Type().IsRegular() {
			if strings.Contains(path, ".gen.") {
				os.Remove(path)
			}
		}
		return err
	})
}

func FormatCode(dir string) {
	//log.Println("formating go code...")
	cmd := exec.Command("goimports", "-w", dir)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("format code error:", err.Error())
		fmt.Println(string(out))
		return
	}

	fmt.Println(string(out))
}

func exportedSourceSymbols(filename string) (types []string, methods []string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return
	}

	// Traverse the AST and print type names
	ast.Inspect(node, func(n ast.Node) bool {
		// Check if the node is a type specification
		if typeSpec, ok := n.(*ast.TypeSpec); ok && typeSpec.Name.IsExported() {
			types = append(types, typeSpec.Name.Name)
		}
		if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
			var recvTypeName string
			switch t := funcDecl.Recv.List[0].Type.(type) {
			case *ast.StarExpr: // Pointer receiver
				if ident, ok := t.X.(*ast.Ident); ok && ident.IsExported() {
					recvTypeName = ident.Name
				}
			case *ast.Ident: // Value receiver
				if t.IsExported() {
					recvTypeName = t.Name
				}
			}
			if recvTypeName != "" && funcDecl.Name.IsExported() {
				methods = append(methods, fmt.Sprintf("%s#%s", recvTypeName, funcDecl.Name.Name))

			}
		}
		return true
	})

	return
}

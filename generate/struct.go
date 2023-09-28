package generate

import (
	"log"

	"github.com/progrium/macdriver/generate/codegen"
	"github.com/progrium/macdriver/generate/modules"
)

func (db *Generator) ToStruct(fw string, sym Symbol) *codegen.Struct {
	knownIssues := map[string]bool{}
	if knownIssues[sym.Name] {
		_, err := sym.Parse("macos") // TODO(tmc): think about this
		log.Printf("skipping struct %s %s because of known issue: decl='%s' err='%v'\n", fw, sym.Name, sym.Declaration, err)
		return nil
	}
	typ := db.TypeFromSymbol(sym)
	s := &codegen.Struct{
		Name:        sym.Name,
		GoName:      modules.TrimPrefix(sym.Name),
		Description: sym.Description,
		DocURL:      sym.DocURL(),
		Type:        typ,
	}

	return s

}

/*
 */

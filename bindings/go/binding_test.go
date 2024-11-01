package tree_sitter_gotemplate_test

import (
	"context"
	"testing"

	"github.com/qvalentin/tree-sitter-go-template/bindings/go"
	sitter "github.com/smacker/go-tree-sitter"
)

func TestGrammar(t *testing.T) {
	n, _ := sitter.ParseCtx(context.Background(), []byte("{{ nil }}"), tree_sitter_gotemplate.GetLanguage())

	if n.String() != "(template (nil))" {
		t.Errorf("Parsing did not work")
	}
}

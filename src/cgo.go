// Package src is a dummy package to make sure c files are included when vendoring
package src

// #cgo CFLAGS: -std=c11 -fPIC
import "C"

import _ "github.com/qvalentin/tree-sitter-go-template/src/tree_sitter"

package c

import "go-core/3-syntax-two/cyclic_deps/pkg/b"

var C int = b.B + 10

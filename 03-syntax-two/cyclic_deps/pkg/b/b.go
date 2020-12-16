package b

import "go-core/3-syntax-two/cyclic_deps/pkg/a"

var B = a.A + 10

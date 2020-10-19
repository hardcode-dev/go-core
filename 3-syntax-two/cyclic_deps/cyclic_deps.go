package a

import "go-core/3-syntax-two/snippets/pkg/c"

var A = 10
var C = c.C + 10
//===========================================
package b

import "go-core/3-syntax-two/snippets/pkg/a"	// где здесь будет ошибка

var B = a.A + 10
//===========================================
package c

import "go-core/3-syntax-two/snippets/pkg/b"

var C int = b.B + 10
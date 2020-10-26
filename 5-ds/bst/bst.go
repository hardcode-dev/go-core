// Пример реализации структуры данных "Двоичное дерево поиска"
// Пример для простоты приведён для дерева, содержащего в качетве значений целые числа.
// Можно по аналогии со стандартной библиотекой с помощью интерфейса обобщить на произвольные типы данных.
// Для этого потребуется контракт на функцию сравнения элементов.
//
// Википедия: Двоичное дерево поиска — это двоичное дерево, для которого выполняются следующие дополнительные условия (свойства дерева поиска):
// - Оба поддерева — левое и правое — являются двоичными деревьями поиска.
// - У всех узлов левого поддерева произвольного узла X значения ключей данных меньше, нежели значение ключа данных самого узла X.
// - У всех узлов правого поддерева произвольного узла X значения ключей данных больше либо равны, нежели значение ключа данных самого узла X.
package main

import "fmt"

// Tree - Двоичное дерево поиска
type Tree struct {
	root *Element
}

// Element - элемент дерева
type Element struct {
	left, right *Element
	Value       int
}

// Insert - вставка элемента в дерево
func (t *Tree) Insert(x int) {
	e := &Element{Value: x}
	if t.root == nil {
		t.root = e
		return
	}
	insert(t.root, e)
}

// inset рекурсивно вставляет элемент в нужный уровень дерева.
func insert(node, new *Element) {
	if new.Value < node.Value {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Value >= node.Value {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search - поиск значения в дереве, выдаёт true если найдено, иначе false
func (t *Tree) Search(x int) bool {
	return search(t.root, x)
}
func search(el *Element, x int) bool {
	if el == nil {
		return false
	}
	if el.Value == x {
		return true
	}
	if el.Value < x {
		return search(el.right, x)
	}
	return search(el.left, x)
}

// String - реализуем интерфейс Stringer для функций печати пакета fmt
func (t Tree) String() string {
	return prettyPrint(t.root, 0)
}

// prettyPrint печатает дерево в виде дерева :)
func prettyPrint(e *Element, spaces int) (res string) {
	if e == nil {
		return res
	}

	spaces++
	res += prettyPrint(e.right, spaces)
	for i := 0; i < spaces; i++ {
		res += fmt.Sprint("\t")
	}
	res += fmt.Sprintf("%d\n", e.Value)
	res += prettyPrint(e.left, spaces)

	return res
}

func initTree() *Tree {
	var t Tree
	t.Insert(10)
	t.Insert(5)
	t.Insert(20)
	t.Insert(15)
	t.Insert(25)
	t.Insert(30)
	t.Insert(35)
	t.Insert(1)
	t.Insert(2)
	t.Insert(6)
	return &t
}

func main() {
	t := initTree()
	fmt.Println(t)

	// 	Output:
	// 					35
	// 				30
	// 			25
	// 		20
	// 			15
	// 	10
	//			6
	// 		5
	// 				2
	// 			1

}

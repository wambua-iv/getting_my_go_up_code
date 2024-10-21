package main

import (
	"cmp"
	"fmt"
)

type Pair[T fmt.Stringer] struct {
	val1 T
	val2 T
}

type Pair2D struct {
	x, y int
}

func (p2 Pair2D) String() string {
	return fmt.Sprintf("values of 2D is %v, %v", p2.x, p2.y)
}

func main() {
	pair1 := Pair[Pair2D]{Pair2D{1, 2}, Pair2D{5, 6}}
	fmt.Println(pair1)

	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(40)
	fmt.Print(t1.Contains(15))
	
	p := NewTree(Person.Order)
	p.Add(Person{"Maria", 50})
	p.Add(Person{"Jane", 30})
	fmt.Println(p.Contains(Person{"Bob", 35}))
}

//generic data and generic functions

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{
			val: v,
		}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

// perform comparison on structs

type Person struct {
	name string
	age  int
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.name, p.name)
	if out == 0 {
		out = cmp.Compare(p.age, p.age)
	}
	return out
}

package main

type IntTree struct {
	val         int
	left, right *IntTree
}

func (iT *IntTree) Insert(val int) *IntTree {
	if iT == nil {
		return &IntTree{val: val}
	}
	if val < iT.val {
		iT.left = iT.left.Insert(val)
	} else if val > iT.val {
		iT.right.Insert(val)
	}
	return iT
}

func (iT *IntTree) Contains(val int) bool {
	switch {
	case iT == nil:
		return false
	case val < iT.val:
		return iT.left.Contains(val)
	case val > iT.val:
		return iT.right.Contains(val)
	default:
		return true
	}
}

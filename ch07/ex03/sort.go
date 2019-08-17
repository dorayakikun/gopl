package treesort

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	return branchString(t, 0)
}

func branchString(t *tree, depth int) string {
	if t == nil {
		return ""
	}

	var sb strings.Builder
	if t.right != nil {
		sb.WriteString(branchString(t.right, depth + 1))
	}

	for i := 0; i < depth; i ++ {
		sb.WriteString("\t")
	}
	sb.WriteString(fmt.Sprintf("%d\n", t.value))

	if t.left != nil {
		sb.WriteString(branchString(t.left, depth + 1))
	}
	return sb.String()
}
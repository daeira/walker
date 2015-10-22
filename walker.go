package walker

import (
	"errors"

	"github.com/coreos/go-etcd/etcd"
)

var SkipNode = errors.New("skip this Entry")

// WalkFunc is the type of the function called for each Node
// visited by Walk.

// If the function returns SkipNode when invoked on a directory node,
// Walk skips the directory's child nodes entirely.
// If the function returns SkipNode when invoked on a non-directory key,
// Walk skips the remaining nodes in the containing directory node.
type WalkFunc func(n *etcd.Node) error

// Walk walks the the response, calling walkFn recursively for each Node
// in the response.
// The files are walked in lexical order, which makes the output deterministic
// but means that for very large responses, Walk can be inefficient.
func Walk(r *etcd.Response, walkFn WalkFunc) error {
	return walk(r.Node, walkFn)
}

// walk recursively descends nodes, calling walkFn
func walk(node *etcd.Node, walkFn WalkFunc) error {
	err := walkFn(node)
	if err != nil {
		if node.Dir && err == SkipNode {
			return nil
		}
		return err
	}

	if !node.Dir {
		return nil
	}

	for _, n := range node.Nodes {
		err = walk(n, walkFn)
		if err != nil {
			if !n.Dir || err != SkipNode {
				return err
			}
		}
	}
	return nil
}

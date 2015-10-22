package walker

import (
	"testing"

	"github.com/coreos/go-etcd/etcd"
)

type walker struct {
	solution string
}

func (w *walker) appendWalkFn(n *etcd.Node) error {
	w.solution = w.solution + n.Key
	return nil
}

var r = &etcd.Response{
	Node: &etcd.Node{
		Key: "h",
		Dir: true,
		Nodes: etcd.Nodes{
			&etcd.Node{Key: "e", Dir: false},
			&etcd.Node{Key: "l", Dir: true},
			&etcd.Node{Key: "l", Dir: false},
			&etcd.Node{
				Key: "o",
				Dir: true,
				Nodes: etcd.Nodes{
					&etcd.Node{Key: " ", Dir: false},
					&etcd.Node{Key: "e", Dir: true},
					&etcd.Node{
						Key: "t",
						Dir: true,
						Nodes: etcd.Nodes{
							&etcd.Node{Key: "c", Dir: false},
							&etcd.Node{Key: "d", Dir: false},
						},
					},
				},
			},
		},
	},
}

func TestWalk(t *testing.T) {
	var solution = "hello etcd"
	w := new(walker)
	if err := Walk(r, w.appendWalkFn); err != nil {
		t.Error(err)
	}
	if w.solution != solution {
		t.Errorf("walk produced wrong solution, got: %s, wanted: %s", w.solution, solution)
	}
}

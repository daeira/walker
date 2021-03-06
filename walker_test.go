package walker

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/coreos/go-etcd/etcd"
	"github.com/daeira/walker/merger"
	"gopkg.in/yaml.v2"
)

type walker struct {
	solution   string
	skipLetter string
}

func (w *walker) appendWalkFn(n *etcd.Node) error {
	if n.Key == w.skipLetter {
		return SkipNode
	}
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
	var tests = []struct {
		name       string
		skipLetter string
		solution   string
	}{
		{"no skip letter", "", "hello etcd"},
		{"skip letter o", "o", "hell"},
	}

	for _, test := range tests {
		w := &walker{skipLetter: test.skipLetter}
		if err := Walk(r, w.appendWalkFn); err != nil {
			t.Error(err)
		}
		if w.solution != test.solution {
			t.Errorf("walk produced wrong solution for test '%s', got: '%s', wanted: '%s'", test.name, w.solution, test.solution)
		}
	}
}

func TestRead(t *testing.T) {
	path := "/home/rz/repos/puppet/puppet-foreman/enc/puppet"
	if err := filepath.Walk(path, readYAMLWalkFn); err != nil {
		t.Error(err)
	}
	dst := make(map[interface{}]interface{})

	for _, s := range sources {
		rel, _ := filepath.Rel(path, s.Origin)
		merger.Origin(s.Data, s.Plugin+":"+rel)
		merger.Merge(dst, s.Data)
	}

	data, err := yaml.Marshal(dst)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
}

package walker

import (
	"fmt"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

var files = `/aaa/eee/iii/kkkk
/aaa/eee/iii/llll
/aaa/eee/iii/mmmm
/aaa/eee/jjj/kkkk
/aaa/eee/jjj/llll
/aaa/eee/jjj/mmmm
/aaa/fff/iii/kkkk
/aaa/fff/iii/llll
/aaa/fff/iii/mmmm
/aaa/fff/jjj/kkkk
/aaa/fff/jjj/llll
/aaa/fff/jjj/mmmm
/aaa/ggg/iii/kkkk
/aaa/ggg/iii/llll
/aaa/ggg/iii/mmmm
/aaa/ggg/jjj/kkkk
/aaa/ggg/jjj/llll
/aaa/ggg/jjj/mmmm
/aaa/hhh/iii/kkkk
/aaa/hhh/iii/llll
/aaa/hhh/iii/mmmm
/aaa/hhh/jjj/kkkk
/aaa/hhh/jjj/llll
/aaa/hhh/jjj/mmmm
/bbb/eee/iii/kkkk
/bbb/eee/iii/llll
/bbb/eee/iii/mmmm
/bbb/eee/jjj/kkkk
/bbb/eee/jjj/llll
/bbb/eee/jjj/mmmm
/bbb/fff/iii/kkkk
/bbb/fff/iii/llll
/bbb/fff/iii/mmmm
/bbb/fff/jjj/kkkk
/bbb/fff/jjj/llll
/bbb/fff/jjj/mmmm
/bbb/ggg/iii/kkkk
/bbb/ggg/iii/llll
/bbb/ggg/iii/mmmm
/bbb/ggg/jjj/kkkk
/bbb/ggg/jjj/llll
/bbb/ggg/jjj/mmmm
/bbb/hhh/iii/kkkk
/bbb/hhh/iii/llll
/bbb/hhh/iii/mmmm
/bbb/hhh/jjj/kkkk
/bbb/hhh/jjj/llll
/bbb/hhh/jjj/mmmm
/ccc/eee/iii/kkkk
/ccc/eee/iii/llll
/ccc/eee/iii/mmmm
/ccc/eee/jjj/kkkk
/ccc/eee/jjj/llll
/ccc/eee/jjj/mmmm
/ccc/fff/iii/kkkk
/ccc/fff/iii/llll
/ccc/fff/iii/mmmm
/ccc/fff/jjj/kkkk
/ccc/fff/jjj/llll
/ccc/fff/jjj/mmmm
/ccc/ggg/iii/kkkk
/ccc/ggg/iii/llll
/ccc/ggg/iii/mmmm
/ccc/ggg/jjj/kkkk
/ccc/ggg/jjj/llll
/ccc/ggg/jjj/mmmm
/ccc/hhh/iii/kkkk
/ccc/hhh/iii/llll
/ccc/hhh/iii/mmmm
/ccc/hhh/jjj/kkkk
/ccc/hhh/jjj/llll
/ccc/hhh/jjj/mmmm
/ddd/eee/iii/kkkk
/ddd/eee/iii/llll
/ddd/eee/iii/mmmm
/ddd/eee/jjj/kkkk
/ddd/eee/jjj/llll
/ddd/eee/jjj/mmmm
/ddd/fff/iii/kkkk
/ddd/fff/iii/llll
/ddd/fff/iii/mmmm
/ddd/fff/jjj/kkkk
/ddd/fff/jjj/llll
/ddd/fff/jjj/mmmm
/ddd/ggg/iii/kkkk
/ddd/ggg/iii/llll
/ddd/ggg/iii/mmmm
/ddd/ggg/jjj/kkkk
/ddd/ggg/jjj/llll
/ddd/ggg/jjj/mmmm
/ddd/hhh/iii/kkkk
/ddd/hhh/iii/llll
/ddd/hhh/iii/mmmm
/ddd/hhh/jjj/kkkk
/ddd/hhh/jjj/llll
/ddd/hhh/jjj/mmmm
`

func TestCreate(t *testing.T) {
	list := strings.Split(files, "\n")

	deps := []Dependency{
		{
			Dir: "/ddd/fff/iii",
			Dep: "/ccc/eee/iii",
		},
		{
			Dir: "/ddd/hhh/jjj/llll",
			Dep: "/ddd/hhh/jjj/mmmm",
		},
	}
	for _, f := range list {
		for i, _ := range deps {
			deps[i].match(f)
		}
	}

	g := Create(list, deps)
	data, err := yaml.Marshal(&g)
	if err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "src", err)
	}
	fmt.Println(string(data))
}

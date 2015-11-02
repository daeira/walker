package walker

import (
	"fmt"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

var files = `/application/orbit
/application/deva
/application/edm
/application/oipf
/application/adwh_tool
/application/bmc_bppm_integration
/db/oracle_ogg
/db/sqoop
/db/postgresql
/db/postgresql_smdb
/db/oracle_ggs
/db/hive
/group/accesszone_dualhomed
/group/kfzv
/group/user_playground
/group/kfzv_db
/group/cardservices_app
/group/kfzv_app
/middleware/middleware
/middleware/wls
/middleware/play
/middleware/tomcat
/middleware/oep
/middleware/osb
/os/i18n
/os/pf_systeminfo
/os/kvm
/os/hardening
/os/services
/os/kickstart
/tools/keyadm
/tools/puppet_enc_demo
/tools/splunk_masternode
/tools/keybox
/tools/splunkforwarder
/tools/ansible
`

func TestCreate(t *testing.T) {
	list := strings.Split(files, "\n")

	deps := []Dependency{
	// {
	// Dir: "/application",
	// Dep: "/os/",
	// },

	// {
	// Dir: "/ddd/hhh/jjj/llll",
	// Dep: "/ddd/hhh/jjj/mmmm",
	// },
	}

	g := Create(list, deps)
	data, err := yaml.Marshal(&g)
	if err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "src", err)
	}

	order, cyclic := topSortDFS(*g)
	fmt.Println(string(data))
	fmt.Println(order)
	fmt.Println(cyclic)
}

func TestBla(t *testing.T) {
	data := `---
- name: 2
  edges: []
- name: 5
  edges: [11]
- name: 11
  edges: [2, 9, 10]
- name: 7
  edges: [11, 8]
- name: 9
  edges: []
- name: 10
  edges: []
- name: 8
  edges: [9]
- name: 3
  edges: [10, 8, 11]
`
	nodes := []Node{}
	if err := yaml.Unmarshal([]byte(data), &nodes); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "src", err)
	}
	fmt.Println(nodes)
	l, err := bla(nodes)
	fmt.Println(err)
	fmt.Println(l)
}

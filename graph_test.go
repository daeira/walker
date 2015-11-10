package walker

import (
	"fmt"
	"reflect"
	"regexp"
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
/tools/ansible`

func TestCreate(t *testing.T) {
	list := strings.Split(files, "\n")

	deps := []Dependency{
		{
			DirRegexp: regexp.MustCompile(`/os/.*`),
			DepRegexp: regexp.MustCompile(`/tools/.*`),
		},
		{
			DirRegexp: regexp.MustCompile(`/tools/.*`),
			DepRegexp: regexp.MustCompile(`/db/.*`),
		},
		{
			DirRegexp: regexp.MustCompile(`/db/.*`),
			DepRegexp: regexp.MustCompile(`/group/.*`),
		},
		{
			DirRegexp: regexp.MustCompile(`/db/oracle_ggs`),
			DepRegexp: regexp.MustCompile(`/db/sqoop`),
		},
	}

	g := Create(list, deps)
	fmt.Println(g)
	l, err := topoSort(g)
	if err != nil {
		t.Fatal(err)
	}
	data, err := yaml.Marshal(g)
	if err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "src", err)
	}

	fmt.Println(string(data))
	fmt.Println(l)
}

func TestTopoSort(t *testing.T) {
	g := Graph{
		Node{
			Name: "2",
		},
		Node{
			Name:  "5",
			Edges: []string{"11"},
		},
		Node{
			Name:  "11",
			Edges: []string{"2", "9", "10"},
		},
		Node{
			Name:  "7",
			Edges: []string{"11", "8"},
		},
		Node{
			Name: "9",
		},
		Node{
			Name: "10",
		},
		Node{
			Name:  "8",
			Edges: []string{"9"},
		},
		Node{
			Name:  "3",
			Edges: []string{"10", "8", "11"},
		},
	}

	wanted := []string{"2", "9", "10", "8", "11", "7", "3", "5"}

	sorted, err := topoSort(g)
	if err != nil {
		t.Errorf("got err %s", err)
	}
	if !reflect.DeepEqual(wanted, sorted) {
		t.Errorf("got sorted list %v, wanted %v", sorted, wanted)
	}
}

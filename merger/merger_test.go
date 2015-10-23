package merger

import (
	"testing"

	"github.com/juju/testing/checkers"

	"gopkg.in/yaml.v2"
)

var (
	dst = `---
a11: a
b11:
  b21:
    b31: b31
    b32: b32
    b33: b33
  b22: {}
c11:
  c21: c21
d11: d11
e11:
  - e11
  - e12
  - e13
`

	src = `---
b11:
  b21:
    b31: B31
    b34: B34
c11: {}
e11:
  - E11
f11: []
`

	merged = `---
a11: a
b11:
  b21:
    b31: B31
    b32: b32
    b33: b33
    b34: B34
  b22: {}
c11:
  c21: c21
d11: d11
e11:
  - E11
f11: []
`

	withOrigin = `---
a11: a [origin]
b11:
  b21:
    b31: B31 [origin]
    b32: b32 [origin]
    b33: b33 [origin]
    b34: B34 [origin]
  b22: '{} [origin]'
c11:
  c21: c21 [origin]
d11: d11 [origin]
e11:
  - E11 [origin]
f11: '[] [origin]'
`
)

func TestMerge(t *testing.T) {
	srcMap := make(map[interface{}]interface{})
	dstMap := make(map[interface{}]interface{})
	mergedMap := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(src), srcMap); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "src", err)
	}
	if err := yaml.Unmarshal([]byte(dst), dstMap); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "dst", err)
	}
	if err := yaml.Unmarshal([]byte(merged), mergedMap); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "merged", err)
	}
	Merge(dstMap, srcMap)
	if _, err := checkers.DeepEqual(mergedMap, dstMap); err != nil {
		t.Error(err)
	}
}

func TestOrigin(t *testing.T) {
	srcMap := make(map[interface{}]interface{})
	withOriginMap := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(merged), srcMap); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "merged", err)
	}
	if err := yaml.Unmarshal([]byte(withOrigin), withOriginMap); err != nil {
		t.Fatalf("failed to unmarshal %s: %s", "withOrigin", err)
	}
	Origin(srcMap, "origin")
	if _, err := checkers.DeepEqual(withOriginMap, srcMap); err != nil {
		t.Error(err)
	}
}

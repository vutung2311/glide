package strip

import "testing"

func TestRewriteGodepImport(t *testing.T) {
	tests := map[string]string{
		"github.com/vutung2311/glide/action":                            "github.com/vutung2311/glide/action",
		"github.com/tools/godep/Godeps/_workspace/src/github.com/kr/fs": "github.com/kr/fs",
	}

	for k, v := range tests {
		o := rewriteGodepImport(k)
		if o != v {
			t.Errorf("Incorrect Godep import path rewritten %s: %s", v, o)
		}
	}
}

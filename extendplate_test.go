package extendplate

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestNesting(t *testing.T) {
	set, err := ParseDir("testdata", "*.html", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{
		"base/dashboard.html",
		"base/billing.html",
		"base/docs/release.html",
		"/base/dashboard.html",
		"/base/billing.html",
		"/base/docs/release.html",
	} {
		t.Run(filepath.Base(path), func(t *testing.T) {
			path := path
			tmpl := set.Lookup(path)
			if tmpl == nil {
				t.Fatalf("no template found for %s", path)
			}
			var w bytes.Buffer
			if err := tmpl.Execute(&w, nil); err != nil {
				t.Error(err)
			}
		})
	}
}

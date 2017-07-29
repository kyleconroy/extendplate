package extendplate

import (
	"html/template"
	"path"
	"path/filepath"
	"strings"
)

type TemplateSet struct {
	root      string
	templates map[string]*template.Template
}

func (ts *TemplateSet) Lookup(p string) *template.Template {
	fp := filepath.Join(ts.root, filepath.FromSlash(path.Clean("/"+p)))
	tmpl, _ := ts.templates[fp]
	return tmpl
}

func (ts *TemplateSet) parse(parent *template.Template, dir, pattern string) error {
	files, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		return err
	}
	for _, file := range files {
		var tmpl *template.Template
		if parent == nil {
			tmpl, err = template.ParseFiles(file)
		} else {
			tmpl, err = parent.Clone()
			if err != nil {
				return err
			}
			tmpl, err = tmpl.ParseFiles(file)
		}
		if err != nil {
			return err
		}
		ts.templates[file] = tmpl
		if err := ts.parse(tmpl, strings.Replace(file, filepath.Ext(file), "", 1), pattern); err != nil {
			return err
		}
	}
	return nil
}

// Assumes templates have a specific file extension
// TODO Figure out how to pass a partial template
func ParseDir(dir, pattern string, parent *template.Template) (*TemplateSet, error) {
	if dir == "" {
		dir = "."
	}
	set := TemplateSet{templates: map[string]*template.Template{}, root: dir}
	return &set, set.parse(parent, dir, pattern)
}

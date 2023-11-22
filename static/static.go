package static

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var TemplatesFS embed.FS

//go:embed static/*
var StaticFS embed.FS

func ListLicenses() (licenses []string, err error) {
	e, err := TemplatesFS.ReadDir("templates")
	if err != nil {
		return
	}
	for _, f := range e {
		fn := f.Name()
		splittedFn := strings.Split(fn, ".")
		if splittedFn[0] == "LICENSE" {
			licenses = append(licenses, splittedFn[1])
		}
	}
	return
}

func copyStatic(path string) (err error) {
	entries, err := StaticFS.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		p := filepath.Join(path, entry.Name())
		withoutPrefix := p[len("static/"):]
		if entry.IsDir() {
			if err = os.Mkdir(withoutPrefix, 0755); err != nil {
				return
			}
			if err = copyStatic(p); err != nil {
				return
			}
		} else {
			f, err := os.Create(withoutPrefix)
			if err != nil {
				return err
			}
			defer f.Close()
			b, err := StaticFS.ReadFile(p)
			if err != nil {
				return err
			}
			io.Copy(f, bytes.NewReader(b))
		}
	}
	return
}

func CopyStatic() (err error) {
	return copyStatic("static")
}

type Copy struct {
	Name    string
	License string
	Author  string
	Year    string
	HasMain bool
}

func copyTemplate(c Copy, src string, to string) (err error) {
	b, err := TemplatesFS.ReadFile(src)
	if err != nil {
		return err
	}
	tmpl := string(b)
	t, err := template.New(to).Parse(tmpl)
	if err != nil {
		return err
	}
	f, err := os.Create(to)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = t.Execute(f, c); err != nil {
		return err
	}
	return
}

func CopyTemplates(c Copy) (err error) {
	if err = copyTemplate(c, fmt.Sprintf("templates/LICENSE.%s.template", c.License), "LICENSE"); err != nil {
		return
	}
	if err = copyTemplate(c, "templates/README.md.template", "README.md"); err != nil {
		return
	}
	if c.HasMain {
		if err = copyTemplate(c, "templates/main.go.template", "main.go"); err != nil {
			return
		}
	} else {
		if err = copyTemplate(c, "templates/lib.go.template", fmt.Sprintf("%s.go", c.Name)); err != nil {
			return
		}
	}
	return
}

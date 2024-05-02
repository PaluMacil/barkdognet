package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"os"
	"path"
	"sync"

	"github.com/Masterminds/sprig/v3"
)

//go:embed *
var embeddedFiles embed.FS

type TemplateProvider interface {
	Get(route string) (*template.Template, error)
	Execute(route string, data interface{}) error
}

func NewTemplateSource(reloadTemplates bool) *TemplateSource {
	if reloadTemplates {
		return &TemplateSource{
			FS:        os.DirFS("templates"),
			Reload:    true,
			templates: make(map[string]*template.Template),
		}
	} else {
		subFS, _ := fs.Sub(embeddedFiles, "templates")
		return &TemplateSource{
			FS:        subFS,
			Reload:    false,
			templates: make(map[string]*template.Template),
		}
	}
}

type TemplateSource struct {
	sync.RWMutex
	FS        fs.FS
	Reload    bool
	templates map[string]*template.Template
}

func (ts *TemplateSource) Get(route string) (*template.Template, error) {
	if !ts.Reload {
		ts.RLock()
		tmpl, ok := ts.templates[route]
		ts.RUnlock()
		if ok {
			return tmpl, nil
		}
	}

	tmpl, err := ts.loadTemplate(route)
	if err != nil {
		return nil, err
	}

	if !ts.Reload {
		ts.Lock()
		ts.templates[route] = tmpl
		ts.Unlock()
	}

	return tmpl, nil
}

func (ts *TemplateSource) loadTemplate(route string) (*template.Template, error) {
	templateFile := path.Join(route, "index.gohtml")
	tmpl, err := template.New("layout").Funcs(ts.getFuncMap()).ParseFS(ts.FS, templateFile, "layout.gohtml")
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func (ts *TemplateSource) getFuncMap() template.FuncMap {
	var funcMap = sprig.HtmlFuncMap()
	funcMap["siteName"] = func() string {
		return "My Site"
	}
	return funcMap
}

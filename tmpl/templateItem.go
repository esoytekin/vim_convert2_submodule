package tmpl

import (
	"html/template"
	"os"
	"path"

	"github.com/esoytekin/vim_convert2_submodule/helper"
	"github.com/esoytekin/vim_convert2_submodule/model"
)

//TemplItem is item for template
type TemplItem struct {
	Items []model.SubM
}

func (t TemplItem) WriteTemplate(config model.Conf) {
	vimPath := path.Join(config.VimPath())
	templateName := "submodule.tmpl"
	fPath := path.Join("tmpl", templateName)
	var tp = template.Must(template.New(templateName).ParseFiles(fPath))
	f, err := os.Create(path.Join(vimPath, ".gitmodules"))

	if err != nil {
		panic(err)
	}

	if err := tp.Execute(f, t); err != nil {
		panic(err)
	}

}
func (t TemplItem) SubModules(config model.Conf) {
	for _, item := range t.Items {
		helper.AddSubModule(item, config.VimPath())
	}
}

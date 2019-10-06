package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/esoytekin/vim_convert2_submodule/helper"
	"github.com/esoytekin/vim_convert2_submodule/model"
	"github.com/esoytekin/vim_convert2_submodule/tmpl"
)

func processFile(out chan<- model.SubM, config model.Conf, f os.FileInfo, masterGitURL string) {

	pluginPath := path.Join(config.BundlePath(), f.Name())
	by, err := helper.GetGitURL(pluginPath)

	if err == nil {
		url := strings.TrimSpace(string(by))
		// if it is not master url
		dotGitPath := path.Join(pluginPath, ".git")
		if url != masterGitURL && !helper.IsSubModule(dotGitPath) {
			// remove plugin forlder
			if err := os.RemoveAll(pluginPath); err != nil {
				log.Println(err)

			}
			// add submodule name and url to the list
			out <- model.SubM{Name: f.Name(), URL: url}
		}
	} else {
		log.Println(err)
	}
}

func produce(out chan<- model.SubM, config model.Conf) {

	var wg sync.WaitGroup

	helper.CheckGitModulesFile(config.VimPath())

	files, err := ioutil.ReadDir(config.BundlePath())

	if err != nil {
		panic(err)
	}

	masterGitURL := helper.GetVimFolderGitURL(config.VimPath())
	log.Printf("masterGitURL: %q\n", masterGitURL)
	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			go func(f os.FileInfo) {
				defer wg.Done()
				processFile(out, config, f, masterGitURL)
			}(f)

		}

	}

	wg.Wait()
	close(out)
}

func consume(in <-chan model.SubM) tmpl.TemplItem {
	var subModuleList []model.SubM
	for s := range in {
		subModuleList = append(subModuleList, s)
	}

	return tmpl.TemplItem{Items: subModuleList}
}

// git submodule update --init --recursive
func main() {

	config := model.NewConfig()

	workElem := make(chan model.SubM)

	go produce(workElem, config)

	templItem := consume(workElem)

	templItem.SubModules(config)

}

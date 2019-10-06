package helper

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/esoytekin/vim_convert2_submodule/model"
)

// GetVimFolderGitURL returns git url for vim folder
func GetVimFolderGitURL(vimPath string) string {
	by, err := GetGitURL(vimPath)

	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(by))
}

// GetGitURL returns git url for given path
func GetGitURL(path string) ([]byte, error) {
	// git command to get remote url
	cmd := exec.Command("git", "remote", "get-url", "origin")

	cmd.Dir = path
	return cmd.CombinedOutput()

}

// IsSubModule returns true if path is a git submodule
func IsSubModule(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return false
	}

	return !fi.IsDir()

}

func UpdateGitModules(path string) {
	cmd := exec.Command("git", "submodule", "update", "--init", "--recursive")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// CheckGitModulesFile checks for .gitmodules file under vim folder
// throws exception if it exists
func CheckGitModulesFile(vimPath string) {

	gitmodulesPath := path.Join(vimPath, ".gitmodules")

	if _, err := os.Stat(gitmodulesPath); err == nil || !os.IsNotExist(err) {
		panic(errors.New(".gitmodules file exists in vim folder"))
	}

}

func AddSubModule(subModule model.SubM, path string) {

	cmd := exec.Command("git", "submodule", "add", subModule.URL, "bundle/"+subModule.Name)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

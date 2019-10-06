package model

import (
	"flag"
	"os/user"
	"path"
	"runtime"
)

//Conf configuration specitic entities
type Conf struct {
	currentUser   *user.User // current user
	vimFolderName string     // os specific vim folder name
	privateFolder string
}

func (c Conf) VimPath() string {

	if c.privateFolder != "" {
		return c.privateFolder
	}

	return path.Join(c.currentUser.HomeDir, c.vimFolderName)
}

func (c Conf) BundlePath() string {
	return path.Join(c.VimPath(), "bundle")
}

func NewConfig() Conf {

	currentUser, err := user.Current()

	if err != nil {
		panic(err)
	}

	vimFolderName := getVimFolderName()

	pf := flag.String("path", "", "folder to check for bundle plugin folder")
	flag.Parse()

	//return *pf

	return Conf{
		currentUser:   currentUser,
		vimFolderName: vimFolderName,
		privateFolder: *pf,
	}

}

func getVimFolderName() string {
	if runtime.GOOS == "windows" {
		return "vimfiles"
	}

	return ".vim"

}

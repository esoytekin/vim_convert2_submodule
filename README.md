# Vim Convert To Submodule


Converts plugins under vim bundle folder to git submodule folders


## who should use it?

- if you are using pathogen as your vim plugin manager
- if you are keeping your vim configuration under git

## why should i use it?

- it more managable to keep plugin as submodules. this way you can update your modules with single git command
```git
git submodule update --init --recursive
```

## usage

by default executes in `~/.vim` or `%USERPROFILE%/vimfolder` depending on the operating system

you can give custom path using `--path` flag

package main

import (
	"fmt"
	"os"
)

func main() {
	/*data := make([]byte, 0)
	for j := 0; j < 100; j++ {
		MkDir("test/" + strconv.Itoa(j))
		for i := 0; i < 100; i++ {
			os.WriteFile("test/"+strconv.Itoa(j)+"/"+strconv.Itoa(i), data, os.FileMode(0666))
		}
	}*/
	out := ScanDirTree("./")

	for i := 0; i < len(out); i++ {
		fmt.Printf("%d %s < Name | IsDir > %t \n", i+1, out[i].Name, out[i].IsDir)
	}

	fmt.Printf("\n\n---  Dirs Only  ---\n\n")
	dirs := ListDirs(out)
	for i := 0; i < len(dirs); i++ {
		fmt.Printf("%d %s\n", i+1, dirs[i])
	}

	fmt.Printf("\n\n---  Files Only  ---\n\n")
	files := ListFiles(out)
	for i := 0; i < len(files); i++ {
		fmt.Printf("%d %s\n", i+1, files[i])
	}

	fmt.Printf("\nTotal items: %d\n\n", len(out))
}

func Copy(orgPath string, dstPath string) (err string) {
	orgFile, fail := os.ReadFile(orgPath)
	if fail != nil {
		err = "you don't have the permission to read the orig file"
	} else {
		os.WriteFile(dstPath, orgFile, os.FileMode(0666))
	}
	return err
}

func MkDir(path string) (err string) {
	fail := os.Mkdir(path, os.FileMode(0777))
	if fail != nil {
		err = "error directory can't be created"
	}
	return err
}

func Remove(path string) (err string) {
	fail := os.RemoveAll(path)
	if fail != nil {
		err = "file: \"" + path + "\" can't be removed"
	}
	return
}

func ScanDirTree(dir string) []FileTree {
	if byte(dir[len(dir)-1]) != '/' {
		dir += "/"
	}
	que := make([]string, 1)
	que[0] = dir
	out := make([]FileTree, 1)
	out[0] = FileTree{Name: shrinkLeft(dir, len(dir)-1), IsDir: true}
	for j := 0; j < len(que); j++ {
		directory, _ := os.ReadDir(que[j])
		for i := 0; i < len(directory); i++ {
			if directory[i].IsDir() {
				que = growList(que, que[j]+directory[i].Name()+"/")
			}
			out = growFileTree(out, FileTree{Name: shrinkLeft(que[j]+directory[i].Name(), len(dir)-1), IsDir: directory[i].IsDir()})
			//fmt.Printf("%d %s < Name | IsDir > %t \n", len(out), out[len(out)-1].Name, out[len(out)-1].IsDir)
		}
	}
	return out
}

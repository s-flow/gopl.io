package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const curDir = "./"

var studyMembers = []string{
	"jeonghyunseok",
	"parkjinhyung",
}

var ignoreDirNames = []string{
	".git",
	"gen-study-folder",
}

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

func makeExampleDir(memberName string, exDirs []string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("fail to get working directory: %v\n", err)
		return
	}

	_ = os.Mkdir("_"+memberName, 0755)

	err = os.Chdir(filepath.Join(wd, "_"+memberName))
	if err != nil {
		fmt.Printf("fail to change directory: %v", err)
		return
	}

	for _, exDir := range exDirs {
		_ = os.Mkdir(exDir, 0755)
	}

	// upperDir := filepath.Join(wd, "..")
	err = os.Chdir(wd)
	if err != nil {
		fmt.Printf("fail to change directory: %v\n", err)
		return
	}
}

func makeMemberDir(chapterName string) {
	// 1. go into chapter directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("fail to get working directory: %v\n", err)
		return
	}
	newDir := filepath.Join(wd, chapterName)
	err = os.Chdir(newDir)

	tmpDir, _ := os.Getwd()
	fmt.Println("makeMemberDir", "Chdir", tmpDir)
	if err != nil {
		fmt.Printf("fail to change directory: %v\n", err)
		return
	}

	// 2. get example directory names
	dirNames, err := getDirNames()
	fmt.Println("dirName of Examples", dirNames)
	if err != nil {
		fmt.Printf("fail to get dir names: %v\n", err)
		return
	}

	// 3. run makeExampleDir()
	for _, sm := range studyMembers {
		makeExampleDir(sm, dirNames)
	}
	// upperDir := filepath.Join(wd, "..")
	err = os.Chdir(wd)
	if err != nil {
		fmt.Printf("fail to change directory: %v\n", err)
		return
	}
}

func needIgnore(name string, igList []string) bool {
	for _, ig := range igList {
		if name == ig {
			return true
		}
	}
	return false
}

// https://stackoverflow.com/a/14668907/6513756
func getDirNames() ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(curDir)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		if file.IsDir() && !needIgnore(file.Name(), ignoreDirNames) {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

func main() {
	// 1. get all the folder name except Ignore
	chapterNames, err := getDirNames()
	fmt.Println("chapterNames", chapterNames)
	if err != nil {
		fmt.Printf("failed to get dir names: %v\n", err)
	}

	for _, cn := range chapterNames {
		makeMemberDir(cn)
	}
}

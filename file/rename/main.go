package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//dirList := getDirList("/home/liyang/Share/盘龙_大灰狼/")
	//for _, dirPath := range dirList {
	//fmt.Println(dirPath)
	//wg.Add(1)
	//go renameFile(dirPath)
	//}
	//wg.Wait()
}

func getDirList(dir string) []string {
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var dirs []string
	for _, v := range dirList {
		if v.IsDir() {
			dirPath := dir + v.Name()
			dirs = append(dirs, dirPath)
		}
	}
	return dirs
}

func renameFile(file string) {
	if !strings.HasSuffix(file, "/") {
		file = file + "/"
	}
	dirList, err := ioutil.ReadDir(file)
	if err != nil {
		panic(err)
	}
	for _, v := range dirList {
		filePath := file + v.Name()
		renamePath := file + v.Name()[0:3] + path.Ext(v.Name())
		err := os.Rename(filePath, renamePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(filePath)
		fmt.Println(renamePath)
	}
	defer wg.Done()
}

package main

import (
	"fmt"
	"io/ioutil"
)

func getFiles (path string)([]string, error){
	var files []string
	dirFiles,err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}
	for _,i := range dirFiles {
		files = append(files, i.Name())
	}
	return files, nil

}

func main() {
	f, _ := getFiles(".")
	fmt.Println(f)
}

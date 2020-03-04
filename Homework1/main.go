package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dir := "files"
	util(dir)
}

func util(dir string){
	x := make(map[string][]string)
	files, _ := ioutil.ReadDir(dir)

	for _, file := range files {
		file, err := os.Open("files/" + file.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		data := make([]byte, 64)
		for {
			n, err := file.Read(data)
			if err == io.EOF {
				break
			}
			text := string(data[:n])
			arr := strings.Split(text, " ")
			for i := 0; i < len(arr); i++ {
				now := x[arr[i]]
				now = append(now, file.Name())
				x[arr[i]] = now
			}
		}
	}

	file, err := os.Create("answer.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}

	defer file.Close()

	for key := range x {
		file.WriteString("'" + key + "'" + ":" + " " + "{")
		for i := 0; i < len(x[key]); i++ {
			if i == len(x[key])-1 {
				file.WriteString(x[key][i][len(dir)+1:] + "}")
				file.WriteString("\n")
			} else {
				file.WriteString(x[key][i][len(dir)+1:] + ", ")
			}
		}
	}
}
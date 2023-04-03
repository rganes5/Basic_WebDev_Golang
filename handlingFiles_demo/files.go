package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "This is the content inside the file"
	//creating a file
	file, err := os.Create("./filesSample.txt")
	checkError(err)
	// if err != nil {
	// 	panic(err)
	// }

	//writing data into the file
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("length is :", length)
	defer file.Close()
	readfile("./filesSample.txt")
}

// reading the file
func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("The data using the readfile is :", string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

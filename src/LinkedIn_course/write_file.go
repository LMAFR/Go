package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// This time let's define the text in format byte, because, some times, we could want to read/write texts in a format which is not text:
	contenido := []byte("Hello world!")

	data := ioutil.WriteFile("hello_world.txt", contenido, 0755) // 0755 means 0 111 101 101 and is a way to define what could do each user with
	// this file. (Permissions)
	show_error(data)

	fmt.Println(data)

}

func show_error(e error) {

	if e != nil {
		panic(e)
	}

}

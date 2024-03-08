package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("C:/dev/projects/go/service-helper/test.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	runInstaller(f, "C:/dev/projects/go/fake-installer/")
}

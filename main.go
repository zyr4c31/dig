package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[:]
	if len(args) < 1 {
		log.Fatal("no args")
	}

	url := fmt.Sprintf("https://pkg.go.dev/search?q=%v", args[len(args)-1])
	fmt.Printf("url: %v\n", url)

	curl := exec.Command("curl", url)
	out, err := curl.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("out: %v\n", string(out))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// get input file location
	var dir string
	fmt.Print("Directory of your input file (drag & drop may help) -> ")
	fmt.Scanf("%s", &dir)

	file, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	content, err := os.Open("teste.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

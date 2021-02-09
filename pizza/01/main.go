package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("a_example")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var n, t2, t3, t4 int
	scanner.Scan()
	firstLine := scanner.Text()
	arrStr := strings.Split(firstLine, " ")
	n, err = strconv.Atoi(arrStr[0])
	Handle(err)
	t2, err = strconv.Atoi(arrStr[1])
	Handle(err)
	t3, err = strconv.Atoi(arrStr[2])
	Handle(err)
	t3, err = strconv.Atoi(arrStr[3])
	Handle(err)

	fmt.Println(">>>", n, t2, t3, t4)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("vim-go")
}

func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

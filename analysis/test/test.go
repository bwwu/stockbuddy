package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var tflag = flag.Bool("test_arg", false, "help")

func main() {
	flag.Parse()
	fmt.Printf("The flag val is %v\n", *tflag)
	f, err := os.Open("../../sp500.txt")
	if err != nil {
		log.Panic(err)
	}

	reader := bufio.NewReader(f)
	cnt := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(strings.TrimSpace(line))
		cnt++
	}

	fmt.Printf("Num lines = %v\n", cnt)
}

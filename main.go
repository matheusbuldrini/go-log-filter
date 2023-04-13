package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "myapp.log", "Path to the log to filter")
	level := flag.String("level", "ERROR", "DEBUG/INFO/ERROR/CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}

}

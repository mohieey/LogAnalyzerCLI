package main //all go source files starting with package declaration

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log file that will be analyzed")
	level := flag.String("level", "ERROR", "The level of the logs to analyze, default is ERROR, but you can change it to INFO, DEBUG or WARNING")
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

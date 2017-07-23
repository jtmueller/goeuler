package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	problems := []func(){
		run001, run002, run003,
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nEnter an Euler problem number, 1 to %d: ", len(problems))
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if input == "exit" {
			break
		}

		num, err := strconv.Atoi(input)
		if err == nil && num > 0 && num <= len(problems) {
			start := time.Now()
			problems[num-1]()
			elapsed := time.Since(start)
			fmt.Printf("Execution time: %s\n", elapsed)
		}
	}
}

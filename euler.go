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
		run001, run002, run003, run004, run005,
		run006, run007, run008, run009, run010,
		run011, run012, run013, run014, run015,
		run016, run017, run018, run019,
	}
	reader := bufio.NewReader(os.Stdin)

Loop:
	for {
		fmt.Printf("\nEnter an Euler problem number, 1 to %d or 'all' or 'exit': ", len(problems))
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")

		switch input {
		case "exit":
			break Loop
		case "all":
			start := time.Now()
			for _, p := range problems {
				p()
				fmt.Println()
			}
			elapsed := time.Since(start)
			fmt.Printf("Execution time: %s\n", elapsed)
		default:
			num, err := strconv.Atoi(input)
			if err == nil && num > 0 && num <= len(problems) {
				start := time.Now()
				problems[num-1]()
				fmt.Println()
				elapsed := time.Since(start)
				fmt.Printf("Execution time: %s\n", elapsed)
			}
		}
	}
}

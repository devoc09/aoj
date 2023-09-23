package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		if scanner.Text() == "0" {
			return
		}

		s := scanner.Text()
		var sum int
		for _, n := range s {
			n, err := strconv.Atoi(string(n))
			if err != nil {
				panic(err)
			}
			sum += n
		}
		fmt.Println(sum)
	}
}

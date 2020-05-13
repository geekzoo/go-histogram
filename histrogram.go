package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	red    = "\033[91m"
	yellow = "\033[93m"
	green  = "\033[92m"
	blue   = "\033[94m"
	white  = "\033[97m"
	reset  = "\033[0m"

	graf = '∎'
	//graf = '🆇'
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {

		data := stdin.Text()
		out := strings.Fields(data)
		counts, _ := strconv.Atoi(out[1])

		fmt.Printf(blue)
		fmt.Printf("%v [%4v]", out[0], out[1])
		fmt.Printf(reset)

		for row := 1; row <= counts; row++ {
			fmt.Printf(red)
			fmt.Printf("%c", graf)
			fmt.Printf(reset)
		}
		fmt.Println()
	}
}

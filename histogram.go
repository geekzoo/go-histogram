package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const (
	red    = "\033[91m"
	yellow = "\033[93m"
	green  = "\033[92m"
	blue   = "\033[94m"
	white  = "\033[97m"
	reset  = "\033[0m"

	graf = '∎'
	//graf = '>'
)

var version = "development" //Set by Makefile at build time for production
// histogram

func main() {
	cleanclose()
	menu()
	buildhist()
}

func cleanclose() {
	catch := make(chan os.Signal)
	signal.Notify(catch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-catch
		fmt.Printf("\n\n")
		fmt.Printf(reset)
		fmt.Println(reset, "\t [Ctrl+C pressed in Terminal reseting]")
		fmt.Printf(reset)
		os.Exit(0)
	}()
}

func menu() {
	FlagVers := flag.Bool("v", false, "")
	FlagHelp := flag.Bool("h", false, "")
	flag.Parse()
	if *FlagVers {
		fmt.Printf("code version: %v\n", version)
		os.Exit(0)
	}
	if *FlagHelp {
		fmt.Printf("Usage: cat mydata.txt | ./go-histrogram\n")
		fmt.Printf("-v Version:\n")
		fmt.Printf("-h Help:\n")
		os.Exit(0)
	}
}

func buildhist() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {

		data := stdin.Text()
		if data == "" {
			log.Println(yellow + "LINE: NULL" + reset)
			os.Exit(1)
		}
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
		fmt.Printf("\n")
	}
}

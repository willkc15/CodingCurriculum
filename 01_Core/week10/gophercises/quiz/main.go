package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var fnPtr *string
var timePtr *int

func init() {
	//Set up flags
	fnPtr = flag.String("file", "problems.csv", "path to file to read from")
	timePtr = flag.Int("time", 30, "how long user has to take quiz")
	flag.Parse()
}

func main() {
	f, err := os.Open(*fnPtr)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	r := csv.NewReader(f)                 //New reader to read csv
	scanner := bufio.NewScanner(os.Stdin) //Scanner to read input from stdin
	go quiz(r, scanner)
	time.Sleep(time.Duration(*timePtr) * time.Second)

}

func quiz(r *csv.Reader, scanner *bufio.Scanner) {
	var correct, total int
	for {
		//Read the csv file, if we've reached EOF break loop
		ln, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		q, ans := ln[0], ln[1]
		total++

		fmt.Println(q)
		fmt.Println("Please type your answer now.")

		//Scan for input from stdin
		scanner.Scan()
		input := scanner.Text()

		if input == ans {
			fmt.Println("Correct")
			correct++
		} else {
			fmt.Println("WRONG!, Answer was " + ans)
		}
	}
	fmt.Printf("You answered %d questions correctly out of %d\n", correct, total)
}

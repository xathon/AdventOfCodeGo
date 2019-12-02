package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputfile = "input.txt"

func main() {
	t1Calc()
	t2Calc()
}

func t1Calc() {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var summe int
	var zahl string

	for scanner.Scan() {
		if zahl, err := strconv.Atoi(scanner.Text()); err == nil {
			summe += t1GetFuel(zahl)
		}
	}
	_ = zahl

	file.Close()
	fmt.Println("The answer to Puzzle 1 is: " + strconv.Itoa(summe))
}

func t1GetFuel(mass int) int {
	var fuel int
	fuel = mass / 3
	fuel -= 2

	return fuel

}

func t2Calc() {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var summe int
	var zahl string

	for scanner.Scan() {
		if zahl, err := strconv.Atoi(scanner.Text()); err == nil {
			summe += t2GetFuel(zahl)
		}
	}
	_ = zahl

	file.Close()
	fmt.Println("The answer to Puzzle 2 is: " + strconv.Itoa(summe))
}

func t2GetFuel(mass int) int {
	var fuel int
	fuel = mass / 3
	fuel -= 2

	if fuel > 0 {
		fuel += t2GetFuel(fuel)
	} else {
		fuel = 0
	}

	return fuel

}

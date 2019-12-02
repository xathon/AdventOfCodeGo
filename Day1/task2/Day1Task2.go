package task2

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strconv"
)

func Calc() {
	file, err := os.Open("./Day1/input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file);
	scanner.Split(bufio.ScanLines);
	var summe int
	var zahl string
	
	for scanner.Scan() {
		if zahl, err := strconv.Atoi(scanner.Text()); err == nil {
			summe += getFuel(zahl);
		}
	}
	_ = zahl;

	file.Close();
	fmt.Println("The answer to Puzzle 2 is: " + strconv.Itoa(summe));
}

func getFuel(mass int) int {
	var fuel int
	fuel = mass / 3;
	fuel -= 2;
	
	if fuel > 0 {
		fuel += getFuel(fuel)
	} else {
		fuel = 0
	}
	
	return fuel
	
}


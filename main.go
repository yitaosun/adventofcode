package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type solver func(input []string) int

var solvers = map[string]solver{
	"01A": solve01A,
	"01B": solve01B,
	"02A": solve02A,
	"02B": solve02B,
	"03A": solve03A,
	"03B": solve03B,
	"04A": solve04A,
	"04B": solve04B,
	"05A": solve05A,
	"05B": solve05B,
	"06A": solve06A,
	"06B": solve06B,
	"07A": solve07A,
	"07B": solve07B,
}

func runSolver(id, env string) {
	solve, ok := solvers[id]
	if !ok {
		log.Fatal("No solver for", id)
	}
	f, err := os.Open(fmt.Sprintf("%s/%s.txt", id[:2], env))
	if err != nil {
		log.Fatalf("No %s input for %s", env, id)
	}
	var input []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}
	log.Printf("%s: %d", strings.Title(env), solve(input))
}

func main() {
	log.SetFlags(log.Lshortfile)
	for _, id := range os.Args[1:] {
		runSolver(id, "test")
		runSolver(id, "prod")
	}
}
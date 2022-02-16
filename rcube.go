package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	Moves = "BFURLD"
)

var (
	MoveNum = flag.Int("m", 30, "The number of moves which will be in the scramble")
)

// determine a chance of 1/n
func oneInN(n uint64) bool {
	return rand.Intn(int(n)) == 1
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	var scramble []string
	var forbiddenMove string
	for i := 0; i < *MoveNum; i++ {
		// Get the last move in the scramble
		var last string
		if len(scramble) != 0 {
			last = scramble[len(scramble)-1]
		}

		// Get a new move
		n := rand.Intn(len(Moves))
		m := string(Moves[n])

		// Check to see if a move is repeated more than 2 times
		if m == forbiddenMove {
			continue
		}

		// If the next move is the same as the last move, remove the last
		// move and make it {move}2 (e.g., U2, B2, etc) and make it so that
		// the next one cannot be the same one
		if last == m {
			scramble = scramble[:len(scramble)-1]
			scramble = append(scramble, fmt.Sprintf("%s2", m))
			forbiddenMove = m
		} else if oneInN(3) {
			// This statement is for a random chance of being a prime move
			// Chance of being prime: 1/3
			scramble = append(scramble, fmt.Sprintf("%s'", m))
		} else {
			// It is a unique move, and is by chance, not prime
			scramble = append(scramble, m)
		}
	}
	fmt.Println(strings.Join(scramble, " "))
}

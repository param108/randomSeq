package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type State struct {
	Pattern string
	Index   int
}

func generateString(length int) string {
	bitstring := ""
	for i := 0; i < length; i++ {
		if rand.Int()%2 == 0 {
			bitstring = bitstring + "0"
		} else {
			bitstring = bitstring + "1"
		}
	}
	return bitstring
}

func runPattern(state *State, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(string(state.Pattern[state.Index]))
		state.Index++
		if state.Index >= len(state.Pattern) {
			state.Index = 0
		}
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	length := (rand.Int() % 145) + 5
	offset := (rand.Int() % length)
	repeat, _ := strconv.Atoi(os.Args[1])
	state := &State{generateString(length), offset}
	runPattern(state, repeat)
}


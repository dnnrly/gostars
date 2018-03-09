package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/unixpickle/markovchain"
)

func main() {
	numWords := flag.Int("words", 10, "maximum number of words to print")
	minLength := flag.Int("min", 5, "minimum length of words")
	prefixLen := flag.Int("prefix", 8, "prefix length in words")
	source := flag.String("source", "", "file with source text")
	flag.Parse()

	if *source == "" {
		fmt.Fprintf(os.Stderr, "Must specify input data\n")
		flag.Usage()
		os.Exit(1)
	}

	historySize := *prefixLen

	body, err := ioutil.ReadFile(*source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read text:", err)
		os.Exit(1)
	}

	fields := strings.Split(string(body), "")
	fieldChan := make(chan string, 10)
	go func() {
		for _, field := range fields {
			fieldChan <- field
		}
		close(fieldChan)
	}()

	// Avoid clashing with known names
	existing := make(map[string]bool)
	for _, v := range strings.Split(string(body), "\n") {
		existing[strings.ToLower(v)] = true
	}

	rand.Seed(time.Now().UnixNano())

	chain := markovchain.NewChainText(fieldChan, historySize)
	state := randomStart(chain)
	word := ""
	names := make([]string, 0, *numWords)
	for len(names) < *numWords {
		ts := state.(markovchain.TextState)
		ch := ts[len(ts)-1]

		if ch == "\n" {
			word = strings.Trim(word, " \n-_")
			if len(word) >= *minLength && !existing[strings.ToLower(word)] {
				names = append(names, word)
			}
			word = ""
			state = randomStart(chain)
		} else {
			word += ch
			state = randomTransition(chain, state)
		}
	}

	for _, v := range names {
		name := strings.ToUpper(v[0:1]) + v[1:]
		fmt.Printf("%s\n", name)
	}
}

func randomStart(ch *markovchain.Chain) markovchain.State {
	var allStates []markovchain.State
	ch.Iterate(func(s *markovchain.StateTransitions) bool {
		allStates = append(allStates, s.State)
		return true
	})
	state := allStates[rand.Intn(len(allStates))]

	// Run through the markov chain to land at a more
	// likely state.
	for i := 0; i < 10; i++ {
		newState := randomTransition(ch, state)
		if newState == nil {
			break
		}
		state = newState
	}

	return state
}

func randomTransition(ch *markovchain.Chain, state markovchain.State) markovchain.State {
	entry := ch.Lookup(state)
	if entry == nil || len(entry.Targets) == 0 {
		return nil
	}

	prob := rand.Float64()
	var curProb float64
	for i, x := range entry.Probabilities {
		curProb += x
		if curProb > prob {
			return entry.Targets[i]
		}
	}

	return entry.Targets[len(entry.Targets)-1]
}

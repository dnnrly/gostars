package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz")
)

func main() {
	words := flag.Uint("words", 10, "number of words to generate")
	min := flag.Uint("min", 4, "minimum numbers of letters per word")
	max := flag.Uint("max", 10, "maxumum numbers of letters per word")

	flag.Parse()

	valid := true

	if *words < 3 {
		fmt.Fprintf(os.Stderr, "Minimum number of words to generate is 3\n")
		valid = false
	}

	if *min < 3 {
		fmt.Fprintf(os.Stderr, "Minimum number of letters per word is 3\n")
		valid = false
	}

	if *max > 15 {
		fmt.Fprintf(os.Stderr, "Maximum number of letters per word is 15\n")
		valid = false
	}

	if *max <= *min {
		fmt.Fprintf(os.Stderr, "Maximum number of letters per word must be greater than the minimum\n")
		valid = false
	}

	if !valid {
		flag.Usage()
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	for i := uint(0); i < *words; i++ {
		r := int64(*max - *min)
		len := int64(*min) + rand.Int63n(r) - 1
		w := strings.ToUpper(randLetter())
		for l := 0; l < int(len); l++ {
			w += randLetter()
		}

		fmt.Println(w)
	}
}

func randLetter() string {
	l := int64(len(letters))
	p := rand.Int63n(l)
	return string(letters[p])
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type graphemes []string

func (g graphemes) Next() string {
	l := int64(len(g))
	p := rand.Int63n(l)
	return string(g[p])
}

func main() {
	words := flag.Uint("words", 10, "number of words to generate")
	min := flag.Uint("min", 4, "minimum numbers of letters per word")
	max := flag.Uint("max", 10, "maxumum numbers of letters per word")
	source := flag.String("source", "", "source file for graphemes")

	flag.Parse()

	valid := true

	if *source == "" {
		fmt.Fprintf(os.Stderr, "Must specify source of grapheme data")
		valid = false
	}

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

	body, err := ioutil.ReadFile(*source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read text:", err)
		os.Exit(1)
	}

	g := make(graphemes, 0)
	for _, l := range strings.Split(string(body), "\n") {
		g = append(g, l)
	}

	rand.Seed(time.Now().UnixNano())

	for i := uint(0); i < *words; i++ {
		r := int64(*max - *min)
		len := int64(*min) + rand.Int63n(r) - 1
		w := g.Next()
		for l := 0; l < int(len); l++ {
			w += g.Next()
		}

		fmt.Println(w)
	}
}

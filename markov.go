package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type Markov struct {
	words []string
	cache map[string][]string
}

// Returns a pointer to a new Markov instance.
func NewMarkov() *Markov {
	return &Markov{
		words: make([]string, 0),
		cache: make(map[string][]string),
	}
}

// Takes a slice of slice of words and seeds internal word list to choose from.
func (m *Markov) Seed(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	m.words = strings.Split(string(dat), " ")

	for i := range m.words {
		triple := m.words[i : i+3]
		if i >= len(m.words)-3 {
			return nil
		}

		one := triple[0]
		two := triple[1]
		key := one + two
		_, ok := m.cache[key]
		if !ok {
			m.cache[key] = []string{}
		}

		m.cache[key] = append(m.cache[key], triple[2])
	}

	return nil
}

// Generate random text from our cache.
func (m *Markov) GenerateText(size int) string {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(len(m.words) - 3)

	// Select words randomly.
	genWords := []string{}
	w1 := m.words[seed]
	w2 := m.words[seed+1]

	for i := 0; i < size; i++ {
		w1 = strings.TrimSpace(w1)
		genWords = append(genWords, w1)
		key := w1 + w2
		res, ok := m.cache[key]
		if !ok {
			continue
		}
		tempw2 := res[rand.Intn(len(res))]
		w1 = w2
		w2 = tempw2
	}
	genWords = append(genWords, w2)
	return strings.Join(genWords, " ")
}

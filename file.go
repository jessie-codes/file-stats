package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/montanaflynn/stats"
)

type statistics struct {
	NumDupes  int
	MedLength float64
	StdLength float64
	MedTokens float64
	StdTokens float64
	Keywords  map[string]int
}

var (
	lineLengths  = []float64{}
	tokenCounts  = []float64{}
	lines        = map[string]int{}
	s            = statistics{}
	wg           = sync.WaitGroup{}
	countMutex   = sync.Mutex{}
	keywordMutex = sync.Mutex{}
)

func loadKeywords(path string) {
	s.Keywords = make(map[string]int)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.Keywords[scanner.Text()] = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readFiles(pattern string) {
	files, _ := filepath.Glob(pattern)
	for _, v := range files {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			wg.Add(1)
			go lineStats(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	wg.Wait()
}

func lineStats(line string) {
	countMutex.Lock()
	tokens := strings.Fields(line)
	lineLengths = append(lineLengths, float64(len(line)))
	tokenCounts = append(tokenCounts, float64(len(tokens)))
	lines[line]++

	if lines[line] > 1 {
		s.NumDupes++
	}
	countMutex.Unlock()

	keywordMutex.Lock()
	for _, v := range tokens {
		for i := range s.Keywords {
			if i == v {
				s.Keywords[i]++
			}
		}
	}
	keywordMutex.Unlock()
	wg.Done()
}

func logResults(path string) {
	s.MedLength, _ = stats.Median(lineLengths)
	s.StdLength, _ = stats.StandardDeviation(lineLengths)
	s.MedTokens, _ = stats.Median(tokenCounts)
	s.StdTokens, _ = stats.StandardDeviation(tokenCounts)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(fmt.Sprintf("num dupes\t%d\n", s.NumDupes))
	file.WriteString(fmt.Sprintf("med length\t%f\n", s.MedLength))
	file.WriteString(fmt.Sprintf("std length\t%f\n", s.StdLength))
	file.WriteString(fmt.Sprintf("med tokens\t%f\n", s.MedTokens))
	file.WriteString(fmt.Sprintf("std tokens\t%f\n", s.StdTokens))

	for i, v := range s.Keywords {
		file.WriteString(fmt.Sprintf("keyword_%s\t%d\n", i, v))
	}
}

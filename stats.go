package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ivpusic/grpool"
	"github.com/montanaflynn/stats"
)

type statistics struct {
	numDupes    int
	medLength   float64
	stdLength   float64
	medTokens   float64
	stdTokens   float64
	keywords    map[string]int
	lineLengths []float64
	tokenCounts []float64
	lines       map[string]int
}

var (
	countMutex   = sync.Mutex{}
	keywordMutex = sync.Mutex{}
	wg           = sync.WaitGroup{}
)

func newStatistics() *statistics {
	return &statistics{
		keywords: make(map[string]int),
		lines:    make(map[string]int),
	}
}

func (s *statistics) loadKeywords(path string) *statistics {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s.keywords[scanner.Text()] = 0
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *statistics) readFiles(pattern string) *statistics {
	pool := grpool.NewPool(100, 50)
	defer pool.Release()

	files, _ := filepath.Glob(pattern)
	for _, v := range files {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			wg.Add(1)
			pool.JobQueue <- func() {
				defer wg.Done()
				s.lineStats(line)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	wg.Wait()
	return s
}

func (s *statistics) calc() *statistics {
	s.medLength, _ = stats.Median(s.lineLengths)
	s.stdLength, _ = stats.StandardDeviation(s.lineLengths)
	s.medTokens, _ = stats.Median(s.tokenCounts)
	s.stdTokens, _ = stats.StandardDeviation(s.tokenCounts)

	return s
}

func (s *statistics) output(path string) *statistics {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(fmt.Sprintf("num dupes\t%d\n", s.numDupes))
	file.WriteString(fmt.Sprintf("med length\t%f\n", s.medLength))
	file.WriteString(fmt.Sprintf("std length\t%f\n", s.stdLength))
	file.WriteString(fmt.Sprintf("med tokens\t%f\n", s.medTokens))
	file.WriteString(fmt.Sprintf("std tokens\t%f\n", s.stdTokens))

	for i, v := range s.keywords {
		file.WriteString(fmt.Sprintf("keyword_%s\t%d\n", i, v))
	}

	return s
}

func (s *statistics) lineStats(line string) *statistics {
	countMutex.Lock()
	tokens := strings.Fields(line)
	s.lineLengths = append(s.lineLengths, float64(len(line)))
	s.tokenCounts = append(s.tokenCounts, float64(len(tokens)))
	s.lines[line]++

	if s.lines[line] > 1 {
		s.numDupes++
	}
	countMutex.Unlock()

	keywordMutex.Lock()
	for _, v := range tokens {
		for i := range s.keywords {
			if i == v {
				s.keywords[i]++
			}
		}
	}
	keywordMutex.Unlock()

	return s
}

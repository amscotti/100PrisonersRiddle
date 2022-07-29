package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/amscotti/100PrisonersRiddle/strategy"
	"github.com/schollz/progressbar/v3"
)

const NUMBER_OF_PRISONERS = 100
const NUMBER_OF_SESSIONS = 1000000

func main() {

	var numberOfSessions int
	var strategyName string

	flag.IntVar(&numberOfSessions, "n", NUMBER_OF_SESSIONS, "Number of sessions to run")
	flag.StringVar(&strategyName, "s", "loop", "Type of strategy to use, 'loop', 'random', 'true', 'false'")
	flag.Parse()

	sessions := make(chan int, numberOfSessions)
	results := make(chan bool, numberOfSessions)
	countOfSuccess := 0

	bar := progressbar.Default(int64(numberOfSessions))

	runtime.GOMAXPROCS(runtime.NumCPU())
	for g := 0; g < runtime.NumCPU(); g++ {
		go session(sessions, results, strategy.Get(strategyName))
	}

	go func() {
		for s := 0; s < NUMBER_OF_SESSIONS; s++ {
			sessions <- s
		}
	}()

	for r := 0; r < NUMBER_OF_SESSIONS; r++ {
		if <-results {
			countOfSuccess++
		}
		_ = bar.Add(1)
	}

	close(sessions)
	close(results)

	f := float64(countOfSuccess) / float64(NUMBER_OF_SESSIONS) * 100
	fmt.Printf("Count of success %d - Percentage %.2F%%\n", countOfSuccess, f)
}

func session(sessions <-chan int, results chan<- bool, strategy strategy.Strategy) {
	for range sessions {
		rand.Seed(time.Now().UnixNano())
		boxes := rand.Perm(NUMBER_OF_PRISONERS)

		found := false
		for i := 0; i < NUMBER_OF_PRISONERS; i++ {
			found = strategy.Search(boxes, i)
			if !found {
				break
			}
		}

		results <- found
	}
}

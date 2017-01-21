package warmer

import (
	"github.com/ahmetalpbalkan/go-linq"
	"sync"
	"time"
)

type Warmer struct {
	Purge       bool
	Concurrency uint
	Break string
}

func (warmer Warmer) Run(urls []string) {
	var waitGroup sync.WaitGroup

	var locker = make(chan uint, warmer.Concurrency)
	defer close(locker)

	linq.From(urls).
		Distinct().
		ToSlice(&urls)

	for _, url := range urls {
		locker <- 1
		waitGroup.Add(1)

		go warmer.consume(url, locker, &waitGroup)
	}

	waitGroup.Wait()
}

func (warmer Warmer) consume(url string, locker <-chan uint, waitGroup *sync.WaitGroup) {
	if warmer.Purge {
		Purge(url)
	}

	WarmUp(url)

	duration, _  := time.ParseDuration(warmer.Break)
	time.Sleep(duration)

	<-locker

	waitGroup.Done()
}

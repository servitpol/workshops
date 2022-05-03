package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, ch chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		ch <- tweet
	}
	close(ch)
}

func consumer(tweets chan *Tweet) {
	// TODO: use channel here
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
			continue
		}

		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Modification starts from here
	// Hint: this can be resolved via channels
	var ch = make(chan *Tweet)

	// Producer
	go producer(stream, ch)
	// Consumer
	consumer(ch)

	fmt.Printf("Process took %s\n", time.Since(start))
}

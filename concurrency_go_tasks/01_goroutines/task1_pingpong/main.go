package main

import (
	"io"
	"os"
	"sync"
)

func PingPong(w io.Writer) {
	chPing := make(chan struct{})
	chPong := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			w.Write([]byte("ping\n"))
			chPing <- struct{}{}
			<-chPong
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chPing
			w.Write([]byte("pong\n"))
			chPong <- struct{}{}
		}
	}()

	wg.Wait()
}

func main() {
	PingPong(os.Stdout)
}

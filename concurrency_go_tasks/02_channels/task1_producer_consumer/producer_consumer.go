package producerconsumer

import (
	"fmt"
	"io"
	"sync"
)

var wg sync.WaitGroup

func Run(w io.Writer) {
	chConsumer := make(chan int) //
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			chConsumer <- i
		}
		close(chConsumer)

	}()

	go func() {
		defer wg.Done()
		for value := range chConsumer {
			fmt.Fprintln(w, value)

		}
	}()

	wg.Wait()

}

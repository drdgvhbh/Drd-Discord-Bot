package main

/* package main

import (
	"fmt"

	"github.com/ReactiveX/RxGo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	watcher := observer.Observer{

		// Register a handler function for every next available item.
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},

		// Register a handler for any emitted error.
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},

		// Register a handler when a stream is completed.
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	input := make(chan interface{})
	go func() {
		input <- "hello world"
		input <- "LOL"
		close(input)
	}()

	<-observable.Observable(input).Subscribe(watcher)
}
*/

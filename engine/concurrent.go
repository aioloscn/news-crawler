package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	in := make(chan Request)
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got Item: %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {

}

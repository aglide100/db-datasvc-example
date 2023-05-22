package worker

import (
	"log"

	"github.com/aglide100/db-datasvc-snippet/pkg/request"
)


type Worker struct {
	Queue *request.Queue
	MaxConcurrent int
	RunningJob chan struct{}
	CompletedSignal chan bool
}

func NewWorker(queue *request.Queue, maxConcurrent int) *Worker {
	return &Worker{
		Queue: queue,
		MaxConcurrent: maxConcurrent,
		RunningJob: make(chan struct{}, maxConcurrent),
		CompletedSignal: make(chan bool),
	}
}

func (w *Worker) DoWork() {
	for {
		select {
		case job := <-w.Queue.Jobs:
			if len(w.RunningJob) < w.MaxConcurrent{
				switch job.Type {
				default:
					go func(job request.Job) {
						defer func() {
							w.RunningJob <- struct{}{}
						}()
			
						err := job.Run()
						if err != nil {
							log.Printf(err.Error())
						}
					}(job)
				}
			} else {
				// TODO
				log.Printf("Too many concurrent jobs. Job %s is queued.", job.Name)
			}
			
		}
	}
}

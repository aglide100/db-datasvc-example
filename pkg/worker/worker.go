package worker

import (
	"log"
	"sync"

	"github.com/aglide100/db-datasvc-example/pkg/request"
)


type Worker struct {
	Queue *request.Queue
	MaxConcurrent int
	RunningJob chan struct{}
}

func NewWorker(queue *request.Queue, maxConcurrent int) *Worker {
	return &Worker{
		Queue: queue,
		MaxConcurrent: maxConcurrent,
		RunningJob: make(chan struct{}, maxConcurrent),
	}
}

func (w *Worker) DoWork() {
	wg := new(sync.WaitGroup)
	
	wg.Add(w.MaxConcurrent+1) 

    // completedSignal := make(chan request.Job)

	go func()  {
		wg.Wait()
	}()

	for {
		log.Printf("current running job's count : %d", len(w.RunningJob))
		select {
		case job := <-w.Queue.Jobs:
			select {
				default:
					log.Printf("waiting...")
					// TODO add primary something...
					// w.Queue.AddJob(job)
				case w.RunningJob <- struct{}{}:
					go func(job request.Job) {
						defer func() {
							if err := job.Run(); err != nil {
								log.Printf("%v", err)
							} else {
								<-w.RunningJob
								// wg.Done()
							}
						}()
						
						// completedSignal <- job
					}(job)
			}
		}
		
	}
	
}

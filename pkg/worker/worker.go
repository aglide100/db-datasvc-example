package worker

import (
	"context"
	"log"

	"github.com/aglide100/db-datasvc-snippet/pkg/request"
)


type Worker struct {
	Queue *request.Queue
}

func NewWorker(queue *request.Queue,) *Worker {
	return &Worker{
		Queue: queue,
	}
}

func (w *Worker) DoWork() bool {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case job := <-w.Queue.Jobs:
		go func(job request.Job) {
			err := job.Run()
			if err != nil {
				log.Printf(err.Error())
			}
			cancel()
		}(job)
		}
	}
	// for {
	// 	// time.Sleep(time.Second*3)

	// 	select {
	// 	case <-w.Queue.Ctx.Done():
	// 		log.Printf("Work done in queue %s", w.Queue.Name, w.Queue.Ctx.Err())
	// 		return true
	// 	case job := <-w.Queue.Jobs:
	// 		err := job.Run()
	// 		if err != nil {
	// 			log.Printf(err.Error())
	// 			continue
	// 		} 
	// 	}
	// }
}

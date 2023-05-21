package request

import (
	"context"
	"log"
	"sync"
)

type Queue struct {
	Name string
	Ctx context.Context
	cancel context.CancelFunc
	Jobs chan Job
}

type Job struct {
	Name string
	Action func() error
}

func NewQueue(name string) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		Jobs: make(chan Job),
		Name: name,
		Ctx: ctx,
		cancel: cancel,
	}
}

func (q *Queue) AddJobs(jobs []Job) {
	var wg sync.WaitGroup

	wg.Add(len(jobs))

	for _, job := range jobs {
		go func(job Job) {
			q.AddJob(job)
			wg.Done()	
		}(job)
	}

	go func() {
		wg.Wait()

		q.cancel()
	}()
}

func (q *Queue) AddJob(job Job) {
	q.Jobs <- job
	log.Printf("New job %s added to %s queue", job.Name, q.Name)
}

func (j Job) Run() error {
	log.Printf("Job running: %s", j.Name)
	err := j.Action()
	if err != nil {
		return err
	}

	return nil
}
package request

import (
	"context"
	"log"
)

type Queue struct {
	Name string
	Ctx context.Context
	cancel context.CancelFunc
	Jobs chan Job
	Num int
}

type Job struct {
	Name string
	// TODO change type
	Type string
	Action func() error
}

func NewQueue(name string) *Queue {
	_, cancel := context.WithCancel(context.Background())

	return &Queue{
		Jobs: make(chan Job),
		Name: name,
		cancel: cancel,
		Num: 0,
	}
}

func (q *Queue) AddJobs(jobs []Job) {

	for _, job := range jobs {
		go func(job Job) {
			q.AddJob(job)
		}(job)
	}
}

func (q *Queue) AddJob(job Job) {
	q.Jobs <- job

	q.Num = q.Num+1;
	log.Printf("New job %s added to %s queue, %d", job.Name, q.Name, q.Num)
}

func (j Job) Run() error {
	log.Printf("Job running: %s", j.Name)
	
	errChan := make(chan error)
	go func() {
		err := j.Action()
		errChan <- err
	}()
	
	err := <-errChan
	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}
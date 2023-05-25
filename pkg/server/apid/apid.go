package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pb_svc_apid "github.com/aglide100/db-datasvc-example/pb/svc/apid"
	"github.com/aglide100/db-datasvc-example/pkg/request"
)

type ApidSrv struct {
	pb_svc_apid.ApidServiceServer
	queue *request.Queue
}

func NewApidServiceServer(queue *request.Queue) *ApidSrv {
	return &ApidSrv{
		queue: queue,
	}
}

func (s *ApidSrv)GetArticleList(ctx context.Context, in *pb_svc_apid.GetArticleListReq) (*pb_svc_apid.GetArticleListRes, error) {
	log.Println("GetArticleList")

	newJob := request.Job{
		Name: in.Id,
		Action: func() error {
			// log.Printf("job start! in  %s", in.Id)
			waitTime := rand.Int31n(10)
            fmt.Println("job:", in.Id, "wait time:", waitTime, "second")
            time.Sleep(time.Duration(waitTime) * time.Second)
			log.Printf("job done! in %s", in.Id)
			
			return nil
		},
	}

	s.queue.AddJob(newJob)
	
	return &pb_svc_apid.GetArticleListRes{
	
	}, nil
}

func (s *ApidSrv) AddArticle(ctx context.Context, in *pb_svc_apid.AddArticleReq) (*pb_svc_apid.AddArticleRes, error) {
	newJob := request.Job{
		Name: "add",
		Action: func() error {
			log.Printf("Wrote in here, db func")
			return nil
		},
	}

	s.queue.AddJob(newJob)

	return &pb_svc_apid.AddArticleRes{

	}, nil
}
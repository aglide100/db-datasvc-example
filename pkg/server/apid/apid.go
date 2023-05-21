package server

import (
	"context"
	"log"
	"time"

	pb_svc_apid "github.com/aglide100/db-datasvc-snippet/pb/svc/apid"
	"github.com/aglide100/db-datasvc-snippet/pkg/request"
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
		Name: "test",
		Action: func() error {
			log.Printf("test func! in %s", in.Id)
			time.Sleep(time.Second*5)

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
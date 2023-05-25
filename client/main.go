package main

import (
	"context"
	"flag"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb_svc_apid "github.com/aglide100/db-datasvc-example/pb/svc/apid"
	"github.com/gofrs/uuid"
)

var (
	addr = flag.String("server-addr", "0.0.0.0:49999", "")
)

func CallGetArticle() error {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("can't connect grpc server : %v", err)
	}


	client := pb_svc_apid.NewApidServiceClient(conn)


	id, err := uuid.NewV4()
	if err != nil {
		log.Printf("Can't make uuid %v", err)
	}
	in := &pb_svc_apid.GetArticleListReq{
		Id: id.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	err = nil
	res, err := client.GetArticleList(ctx, in)
    if err != nil {
        log.Printf("failed to create scraper: %v", err)
		return err
    }

	log.Printf("%v", res)

	return nil
}

func main() {
	wg, ctx := errgroup.WithContext(context.Background())
	_ = ctx

	wg.Go(func() error {
		for {
			log.Printf("A")
			time.Sleep(time.Second * 1)
			CallGetArticle()
		}
	})

	// wg.Go(func() error {
	// 	for {

	// 		log.Printf("B")
	// 		time.Sleep(time.Second * 2)
	// 	}
	// 	return nil
	// })

	wg.Wait()
}
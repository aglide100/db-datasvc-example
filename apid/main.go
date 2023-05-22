package main

import (
	"context"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	pb_svc_apid "github.com/aglide100/db-datasvc-snippet/pb/svc/apid"
	"github.com/aglide100/db-datasvc-snippet/pkg/request"
	apidServer "github.com/aglide100/db-datasvc-snippet/pkg/server/apid"
	"github.com/aglide100/db-datasvc-snippet/pkg/worker"
	"golang.org/x/sync/errgroup"
)

var (
	grpcAddr = flag.String("grpc-addr", "0.0.0.0:49999", "used for grpc addr")
)

func main() {
	gRPCL, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		panic(err)
	}
	defer gRPCL.Close()
	
	wg, ctx := errgroup.WithContext(context.Background())
	_ = ctx
	
	grpcServer := grpc.NewServer()

	// using queue channel for svc
	requestQueue := request.NewQueue("requestQueue")

	pb_svc_apid.RegisterApidServiceServer(grpcServer, apidServer.NewApidServiceServer(requestQueue))
	worker := worker.NewWorker(requestQueue, 3)
	wg.Go(func() error {
		// processing job in queue
		// for {
		// 	select {
		// 	case job := <-requestQueue
		// 	}
		// }
		worker.DoWork()
		return nil
	})
	
	wg.Go(func() error {
		log.Printf("Starting grpcServer at: %s", *grpcAddr)
		err := grpcServer.Serve(gRPCL)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	

	wg.Wait()

}
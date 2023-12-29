package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/htetmyomyint-kmp/grpc-server-stream/data"
	"github.com/htetmyomyint-kmp/grpc-server-stream/proto/checker"
	"github.com/htetmyomyint-kmp/grpc-server-stream/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Panicln(err)
	}

	productClient := data.NewProductClient()
	pcServer := server.NewPCServer(productClient)

	checker.RegisterPriceCheckerServer(gs, pcServer)

	go func() {
		time.Sleep(time.Second)

		conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("dial err", err)
			return
		}
		client := checker.NewPriceCheckerClient(conn)
		// ctx, cancelFunc := context.WithTimeout(context.Background(), time.Hour)
		// defer cancelFunc()
		stream, err := client.CheckPrice(context.Background(), &checker.PCRequest{ProductId: "xxx"})
		if err != nil {
			log.Println("client call err", err)
			return
		}
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Println("io err", err)
				return
			}
			if err != nil {
				log.Println("dial err", err)
				return
			}
			log.Println("received ", resp)
		}
	}()

	if err := gs.Serve(l); err != nil {
		log.Panicln(err)
	}
}

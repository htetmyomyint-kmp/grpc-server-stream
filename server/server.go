package server

import (
	"log"
	"time"

	"github.com/htetmyomyint-kmp/grpc-server-stream/data"
	"github.com/htetmyomyint-kmp/grpc-server-stream/proto/checker"
)

type PCServer struct {
	ProductClient *data.ProductClient
}

func NewPCServer(pClient *data.ProductClient) *PCServer {
	return &PCServer{ProductClient: pClient}
}

func (p *PCServer) CheckPrice(req *checker.PCRequest, checkerServer checker.PriceChecker_CheckPriceServer) error {
	for {
		time.Sleep(5 * time.Second)
		if p.ProductClient.IsPriceChanged(req.ProductId) {
			log.Println("price changed")
			checkerServer.Send(&checker.PCResponse{ProductId: req.ProductId, Price: 0.505})
		} else {
			log.Println("same price: do nothing")
		}
	}

}

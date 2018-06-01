package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/magaldima/bizday/pkg"

	"google.golang.org/grpc"
)

var port int64

func init() {
	flag.Int64Var(&port, "port", 8080, "the server port")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pkg.RegisterDateCalcServer(grpcServer, &pkg.Server{})
	// use TLS
	log.Printf("starting bizday server on port %d", port)
	grpcServer.Serve(lis)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"

	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/dcb"
	"github.com/magaldima/bizday/holiday"
	"github.com/magaldima/bizday/pkg"

	"google.golang.org/grpc"
)

var port int64

func init() {
	flag.Int64Var(&port, "port", 8080, "the server port")
}

func main() {
	flag.Parse()

	dcbClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: dcb.Handshake,
		Plugins:         dcb.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("DCB_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})
	defer dcbClient.Kill()

	dayCountBasisClient, err := dcbClient.Client()
	if err != nil {
		panic(err)
	}

	hClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: holiday.Handshake,
		Plugins:         holiday.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("HOLIDAY_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})
	defer hClient.Kill()

	holidayClient, err := hClient.Client()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterDateCalcServer(grpcServer, pkg.New(dayCountBasisClient, holidayClient))
	// use TLS
	log.Printf("starting bizday server on port %d", port)
	grpcServer.Serve(lis)
}

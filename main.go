package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/magaldima/bizday/calendar"

	"github.com/magaldima/bizday/holidays/shared"

	"github.com/hashicorp/go-plugin"

	api "github.com/magaldima/bizday/api"
	"github.com/magaldima/bizday/pkg"

	"google.golang.org/grpc"
)

var port int64

func init() {
	flag.Int64Var(&port, "port", 8080, "the server port")
}

func main() {
	flag.Parse()

	cClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: calendar.Handshake,
		Plugins:         calendar.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("CALENDAR_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})
	defer cClient.Kill()

	calendarClient, err := cClient.Client()
	if err != nil {
		panic(err)
	}

	hClient := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
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
	api.RegisterDateCalcServer(grpcServer, pkg.New(calendarClient, holidayClient))
	// use TLS
	log.Printf("starting bizday server on port %d", port)
	grpcServer.Serve(lis)
}

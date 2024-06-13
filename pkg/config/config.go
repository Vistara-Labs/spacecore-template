package config

import (
	"context"
	"fmt"
	"os/exec"

	"spaceco/pkg/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type MySpacecore struct{}

// HandshakeConfig is the handshake config for the plugin.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SPACECORE_PLUGIN",
	MagicCookieValue: "v1",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"spacecore": &SpacecorePluginImpl{Impl: &MySpacecore{}},
}

// Start implements proto.KSpacecoreServer.
func (g *GRPCServer) Start(ctx context.Context, req *proto.StartRequest) (*proto.StartResponse, error) {
	fmt.Printf("Starting Spacecore...\n")

	msg, err := g.Impl.Start(context.Background())
	if err != nil {
		return nil, err
	}
	// return &proto.StartResponse{Message: msg}, nil
	return &proto.StartResponse{Status: msg}, nil
}

func (g *GRPCServer) Status(context.Context, *proto.StatusRequest) (*proto.StatusResponse, error) {
	fmt.Printf("Checking Spacecore status...\n")
	status, err := g.Impl.Status(context.Background())
	return &proto.StatusResponse{Status: status}, err
}

func (g *GRPCServer) Stop(context.Context, *proto.StopRequest) (*proto.StopResponse, error) {
	fmt.Printf("Stopping Spacecore...\n")
	msg, err := g.Impl.Stop(context.Background())
	// msg, err := g.Impl.Stop(context.Background())
	if err != nil {
		return nil, err
	}
	return &proto.StopResponse{Status: msg}, nil
}

func (g *GRPCServer) Logs(context.Context, *proto.LogsRequest) (*proto.LogsResponse, error) {
	msg, err := g.Impl.Logs(context.Background())
	fmt.Printf("Checking Spacecore logs...\n %s", msg)
	if err != nil {
		return nil, err
	}
	return &proto.LogsResponse{}, nil
}

type SpacecorePluginImpl struct {
	plugin.GRPCPlugin
	plugin.Plugin
	// concrete implementation, written in Go. This is only used for plugins
	// that are written in Go. If the plugin is written in another language,
	// then this will be nil.
	Impl *MySpacecore
}

type GRPCServer struct {
	Impl MySpacecore // Interface for Spacecore services
}

// GRPCClient is an implementation of proto.KSpacecoreClient that talks over RPC.
type GRPCClient struct{ client proto.ScPluginClient }

func (p *SpacecorePluginImpl) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	// proto.RegisterKSpacecoreServer(s, &GRPCServer{Impl: *p.Impl})
	proto.RegisterScPluginServer(s, &GRPCServer{Impl: *p.Impl})
	return nil
}

func (p *SpacecorePluginImpl) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewScPluginClient(c)}, nil
}

func (g *MySpacecore) Start(ctx context.Context) (string, error) {
	// I want this start to be a long running process
	// like running a binary that exposes a REST API
	// or a long running process that listens to a port
	// run vimana run celestia light-node command

	binPath := "/Users/mayurchougule/development/vistara/vimana-cli/tmp/vimana/vimana-darwin-arm64/vimana"
	c := exec.Command(binPath, "run", "celestia", "light-node")
	err := c.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return c.Path, nil

	// return "Statred updated", nil
}

func (g *MySpacecore) Stop(ctx context.Context) (string, error) {
	fmt.Println("Spacecore stopping...")
	return "Stopped updated", nil
}

func (g *MySpacecore) Status(ctx context.Context) (string, error) {
	return "Running blah blah \n nows updated \n", nil
}

func (g *MySpacecore) Logs(ctx context.Context) (string, error) {
	return "Logs updated", nil
}

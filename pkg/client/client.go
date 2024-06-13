package client

import (

	// "proto"

	"spaceco/pkg/proto"
)

type GRPCClient struct {
	client proto.ScPluginClient
}

// implement grpcclient function
// func (p *config.SpacecorePluginImpl) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
// 	return &GRPCClient{client: proto.NewScPluginClient(c)}, nil
// }

// func (p *config.SpacecorePluginImpl) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
// 	return &GRPCClient{client: proto.NewScPluginClient(c)}, nil
// }

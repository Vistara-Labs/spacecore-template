package main

import (
	"log"

	"github.com/hashicorp/go-plugin"
)

func main() {
	// Serve the plugin
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         PluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
	log.Println("Plugin server started successfully.")
	select {}
}

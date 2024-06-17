package main

import (
	"context"
	"log"
)

func (s *SpacecorePluginImpl) Start(ctx context.Context) (string, error) {
	// Implement the Start method
	return runDemoCommand(ctx)
}

func (s *SpacecorePluginImpl) Stop(ctx context.Context) (string, error) {
	return "stopped", nil

}

func (s *SpacecorePluginImpl) Status(ctx context.Context) (string, error) {
	return "running", nil
}

func (s *SpacecorePluginImpl) Logs(ctx context.Context) (string, error) {
	log.Println("Fetching spacecore logs...")
	return runDemoLogs()
}

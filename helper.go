package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func setupSpacecoreLogger() (*logrus.Entry, *os.File, error) {
	logFile, err := os.OpenFile("/tmp/spacecore.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening log file: %v", err)
	}

	logger := logrus.New()
	logger.SetOutput(logFile)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logrus.Fields{"plugin": "spacecore"}), logFile, nil
}

func runDemoCommand(ctx context.Context) (string, error) {
	logger, logFile, err := setupSpacecoreLogger()
	if err != nil {
		return "", err
	}
	defer logFile.Close()

	// Execute a random spawn process for demo purposes
	cmd := exec.CommandContext(ctx, "sh", "-c", "echo 'Spacecore is running'; sleep 1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Error starting Spacecore: %v\n", err)
		return "", err
	}
	logger.Info(string(output))

	// Return the command output as a string
	return string(output), nil
}

func runDemoLogs() (string, error) {
	logFile := "/tmp/spacecore.log"
	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		return "", err
	}
	fmt.Println(string(content))

	return string(content), err
}

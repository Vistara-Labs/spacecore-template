package main

import (
	"context"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

//go:embed hac.toml
var configData embed.FS

type HACConfig struct {
	Spacecore struct {
		ContainerImage string `toml:"container_image"`
		Backend        string `toml:"backend"`
	} `toml:"spacecore"`
	Commands struct {
		Start  string
		Stop   string
		Status string
		Logs   string
	}
}

func loadConfig() (HACConfig, error) {
	var config HACConfig

	cData, err := configData.ReadFile("hac.toml")
	if err != nil {
		return HACConfig{}, err
	}

	if _, err := toml.Decode(string(cData), &config); err != nil {
		log.Printf("Error loading hac.toml: %v\n", err)
		return HACConfig{}, err
	}
	log.Printf("Loaded config : %v\n", config)
	return config, nil
}

func executeCommand(cmdStr string) (string, error) {
	cmd := exec.Command("sh", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

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

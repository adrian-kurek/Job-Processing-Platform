package main

import "github.com/adrian-kurek/Job-Processing-Platform/common/logger"

func main() {
	logger := logger.Setup(logger.APILogSource)
	logger.Info("Application started")
}

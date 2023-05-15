package main

import "github.com/begenov/helpers/pkg/logger"

func main() {
	logger := logger.NewLogger()
	logger.MissingArg("error")
}

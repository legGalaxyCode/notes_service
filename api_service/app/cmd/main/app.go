package main

import "github.com/leggalaxycode/notes_service/api_service/pkg/logging"

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger inited")
}

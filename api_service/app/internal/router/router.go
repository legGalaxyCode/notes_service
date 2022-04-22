package router

import (
	"github.com/julienschmidt/httprouter"
)

func Init() {
	router := httprouter.New()
	logger := logging.New()
}

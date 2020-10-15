package main

import (
	"github.com/ibreakthecloud/infra-mon/pkg/store"
	inmemory "github.com/ibreakthecloud/infra-mon/pkg/store/in-memory"
	"github.com/ibreakthecloud/infra-mon/router"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	port = "8080"
)

func main() {

	// if PORT is provided in environment
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	r := router.New()

	// initialize the data store
	store.NewStore =  inmemory.New()

	logrus.Info("LISTENING ON PORT: ", port)
	logrus.Fatal(r.Run(":"+port))
}
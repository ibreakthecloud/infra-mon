package main

import (
	"github.com/ibreakthecloud/infra-mon/router"
	"github.com/ibreakthecloud/infra-mon/store"
	"github.com/sirupsen/logrus"
)

const (
	port = ":8082"
)

func main() {
	r := router.New()

	// initialize the data store
	store.DataStore = store.New()

	logrus.Info("LISTENING ON PORT: ", port)
	logrus.Fatal(r.Run(port))
}
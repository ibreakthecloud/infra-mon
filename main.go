package main

import (
	"github.com/ibreakthecloud/infra-mon/router"
	"github.com/sirupsen/logrus"
)

const (
	port = ":8082"
)

func main() {
	r := router.New()

	logrus.Info("LISTENING ON PORT: ", port)
	logrus.Fatal(r.Run(port))
}
package main

import (
	"flag"
	"loan/stub"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	var stubPath string
	flag.StringVar(&stubPath, "path", "./", "path to config file")
	flag.Parse()

	route := gin.Default()
	stub.Route(route, path.Join(stubPath, "stub/loanorigination.json"))
	logrus.Fatal(route.Run(":8882"))
}

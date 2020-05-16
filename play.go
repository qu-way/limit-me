package main

import (
	"os"

	"github.com/logcfg/getzap"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
)

func init() {
	log = getzap.GetDevelopmentLogger("", "").Sugar()
}

func main() {
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	log.Infow("hello world",
		"host", host,
		"pwd", pwd,
	)
}

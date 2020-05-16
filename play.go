package main

import (
	"os"
	"time"

	wl "github.com/kimsudo/wulimt"
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
	log.Debugw("hello world",
		"host", host,
		"pwd", pwd,
	)

	r := wl.GetOrParseNew("hello", "2-5s")
	last := time.Now()
	for i := 1; i <= 10; i++ {
		r.Wait()

		cur := time.Now()
		log.Infow("approved to run", "index", i, "interval", cur.Sub(last))
		last = cur
	}
}

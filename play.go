package main

import (
	"context"
	"os"
	"time"

	wl "github.com/kimsudo/wulimt"
	"github.com/logcfg/getzap"
	"go.uber.org/zap"
	xr "golang.org/x/time/rate"
)

var (
	log *zap.SugaredLogger
)

func init() {
	log = getzap.GetDevelopmentLogger("", "").Sugar()
}

func main() {
	defer log.Sync()

	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	log.Debugw("hello world",
		"host", host,
		"pwd", pwd,
	)

	// test1()
	// test2()
	test3()
}

func test1() {
	r := wl.GetOrParseNew("hello", "2-5s")
	last := time.Now()
	for i := 1; i <= 10; i++ {
		// r.Wait()
		r.WaitMore()

		cur := time.Now()
		log.Infow("approved to run", "index", i, "interval", cur.Sub(last))
		last = cur
	}
}

func test2() {
	dataC := make(chan int, 10)
	doneC := make(chan bool)

	start := time.Now()

	last := time.Now()
	go func(c <-chan int) {
		for {
			d := <-c
			cur := time.Now()
			if d <= 0 {
				log.Warnw("got signal to stop", "data", d)
				break
			}

			log.Infow("task to run", "data", d, "interval", cur.Sub(last))
			last = cur
		}
		doneC <- true

	}(dataC)

	r := wl.GetOrParseNew("hello", "2-5s")

	for i := 1; i <= 10; i++ {
		r.Wait()
		log.Debugw("approved to send", "index", i)
		dataC <- i
	}

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

func test3() {
	dataC := make(chan int, 10)
	doneC := make(chan bool)

	start := time.Now()

	last := time.Now()
	go func(c <-chan int) {
		for {
			d := <-c
			cur := time.Now()
			if d <= 0 {
				log.Warnw("got signal to stop", "data", d)
				break
			}

			log.Infow("task to run", "data", d, "interval", cur.Sub(last))
			last = cur
		}
		doneC <- true

	}(dataC)

	r := xr.NewLimiter(xr.Every(time.Duration(float64(time.Second)*2.5)), 1)
	// ctx := context.TODO()
	ctx, _ := context.WithTimeout(context.Background(), 11*time.Second)

	for i := 1; i <= 10; i++ {
		if err := r.Wait(ctx); err != nil {
			log.Warnw("got limiter error", "index", i, zap.Error(err))
		} else {
			log.Debugw("approved to send", "index", i)
			dataC <- i
		}
	}

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

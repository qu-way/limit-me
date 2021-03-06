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
	// test3()
	// test4()
	// test5()
	// test6()
	test7()
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
	ctx := context.TODO()
	// ctx, _ := context.WithTimeout(context.Background(), 11*time.Second)

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

func test4() {
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

	r := xr.NewLimiter(xr.Every(time.Duration(float64(time.Second)*2.5)), 2)
	ctx := context.TODO()

	for i := 1; i <= 10; i += 2 {
		if err := r.WaitN(ctx, 2); err != nil {
			log.Warnw("got limiter error", "index", i, zap.Error(err))
		} else {
			log.Debugw("approved to send", "index", i)
			dataC <- i
			dataC <- i + 1
		}
	}

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

func test5() {
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
	// ctx, _ := context.WithTimeout(context.Background(), 11*time.Second)

	for i := 1; i <= 10; i++ {
		for !r.Allow() {
			log.Debugw("not allowed, waiting", "index", i)
			time.Sleep(250 * time.Millisecond)
		}
		log.Debugw("approved to send", "index", i)
		dataC <- i

		// if err := r.Wait(ctx); err != nil {
		// 	log.Warnw("got limiter error", "index", i, zap.Error(err))
		// } else {
		// 	log.Debugw("approved to send", "index", i)
		// 	dataC <- i
		// }
	}

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

func test6() {
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
	// ctx, _ := context.WithTimeout(context.Background(), 11*time.Second)

	for i := 1; i <= 10; i++ {
		rev := r.Reserve()
		if !rev.OK() {
			log.Warnw("got limiter error: reserve not ok")
			continue
		}

		delay := rev.Delay()
		if delay > 0 {
			log.Debugw("not allowed, waiting", "index", i, "interval", delay)
			<-time.After(delay)
			log.Debugw("time is up, and let's do it")
		}

		log.Debugw("approved to send", "index", i)
		dataC <- i
	}

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

func test7() {
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
	ctx := context.TODO()
	// ctx, _ := context.WithTimeout(context.Background(), 11*time.Second)

	for j := 100; j <= 500; j += 100 {
		go func(j int) {
			for i := 1; i <= 10; i++ {
				if err := r.Wait(ctx); err != nil {
					log.Warnw("got limiter error", "worker", j, "index", i, zap.Error(err))
				} else {
					log.Debugw("approved to send", "worker", j, "index", i)
					dataC <- i + j
				}
			}
		}(j)
	}

	log.Info("begin waiting to finish")
	time.Sleep(30 * time.Second)
	log.Info("end waiting to finish")

	dataC <- 0
	<-doneC

	end := time.Now()
	log.Debugw("all is done", "time_cost", end.Sub(start))
}

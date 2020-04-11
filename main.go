package main

import (
	"fmt"
	"github.com/Dadard29/go-core-job/connector"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	podIp := os.Getenv("API_HOST")
	podPortStr := os.Getenv("API_PORT")
	podPort, err := strconv.Atoi(podPortStr)
	if err != nil {
		panic(err)
	}

	periodStr := os.Getenv("PERIOD")
	period, err := strconv.Atoi(periodStr)
	if err != nil {
		panic(err)
	}

	protectedToken := os.Getenv("PROTECTED_TOKEN")

	c := connector.NewCoreConnector(podIp, podPort, protectedToken)

	fmt.Println("starting ticker with period", periodStr)

	tick := time.NewTicker(time.Second * time.Duration(period))
	done := make(chan bool)
	go func(tick *time.Ticker, done chan bool) {
		fmt.Println("first tick")
		err := c.Job()
		if err != nil {
			fmt.Printf("[ERROR] %v\n", err.Error())
		}

		for {
			select {
			case t := <-tick.C:
				fmt.Println("tick", t)
				fmt.Println("executing job")
				err := c.Job()
				if err != nil {
					fmt.Printf("[ERROR] %v\n", err)
				}

			case <-done:
				return
			}
		}
	}(tick, done)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	done <- true
}

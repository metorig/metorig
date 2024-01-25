package main

import (
	"fmt"
	"github.com/metorig/metorig/model"
	"github.com/metorig/metorig/source"
	"os"
	"syscall"
	"time"
)

var exit chan bool

func track(dataSource source.DataSource) {
	var info syscall.Sysinfo_t
	for {
		err := syscall.Sysinfo(&info)
		if err == nil {
			m := model.Metrics{
				TotalMem: info.Totalram * uint64(info.Unit),
				FreeMem:  info.Freeram * uint64(info.Unit),
			}
			m.UsedMem = m.TotalMem - m.FreeMem
			_ = dataSource.Store(&m)
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("Starting instance resource tracker")

	token := os.Getenv("INFLUX_TOKEN")
	url := os.Getenv("INFLUX_HOST")
	org := os.Getenv("INFLUX_ORG")
	bucket := os.Getenv("INFLUX_BUCKET")

	go track(source.NewInflux(token, url, org, bucket))
	<-exit
	fmt.Println("Work done")
}

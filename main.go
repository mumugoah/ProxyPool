package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/mumugoah/ProxyPool/api"
	"github.com/mumugoah/ProxyPool/getter"
	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/storage"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ipChan := make(chan *models.IP, 2000)
	conn := storage.NewStorage()

	// Start HTTP
	go func() {
		api.Run()
	}()

	// Check the IPs in DB
	// add 定时
	go func() {
		for {
			storage.CheckProxyDB()
			time.Sleep(5 * time.Minute)
		}

	}()

	// Check the IPs in channel
	for i := 0; i < 100; i++ {
		go func() {
			for {
				storage.CheckProxy(<-ipChan)
			}
		}()
	}

	// Start getters to scraper IP and put it in channel
	for {
		x := conn.Count()
		log.Printf("Chan: %v, IP: %v\n", len(ipChan), x)
		if len(ipChan) < 100 {
			go run(ipChan)
		}
		// 5 Minutes Loop Get
		time.Sleep(5 * time.Minute)
	}
}

func run(ipChan chan<- *models.IP) {
	var wg sync.WaitGroup
	funs := []func() []*models.IP{
		getter.IP66,
		getter.Au2,
		getter.CoolProxy,
		getter.Data5u,
		getter.Kuai,
		getter.Xici,
	}
	for _, f := range funs {
		wg.Add(1)
		go func(f func() []*models.IP) {
			temp := f()
			for _, v := range temp {
				ipChan <- v
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
}

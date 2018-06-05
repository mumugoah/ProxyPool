package storage

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"strings"

	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/util"
	"github.com/parnurzeal/gorequest"
)

// CheckProxy .
func CheckProxy(ip *models.IP) {
	if CheckIP(ip) {
		ProxyAdd(ip)
	}
}

// CheckIP is to check the ip work or not
func CheckIP(ip *models.IP) bool {
	//三次验证如果又一次通过都可以
	for i := 0; i < 3; i++ {
		resp, body, errs := gorequest.New().Proxy(ip.Type + "://" + ip.Data).Timeout(5 * time.Second).Get(util.NewConfig().CheckURL).End()
		if errs != nil {
			if i < 2 {
				time.Sleep(1 * time.Second)
				continue
			} else {
				break
			}

		}
		if resp.StatusCode == 200 {
			// 判断结果是否含有值
			if strings.Contains(body, util.NewConfig().CheckString) {
				log.Printf("Check Proxy %s success", ip.Data)
				return true
			}
		}
		if i < 2 {
			time.Sleep(1 * time.Second)
			continue
		} else {
			break
		}
	}
	log.Printf("Check Proxy %s fail", ip.Data)
	return false
}

// CheckProxyDB to check the ip in DB
func CheckProxyDB() {
	conn := NewStorage()
	x := conn.Count()
	log.Println("Before check, DB has:", x, "records.")
	ips, err := conn.GetAll()
	if err != nil {
		log.Println(err.Error())
		return
	}
	var wg sync.WaitGroup
	for _, v := range ips {
		wg.Add(1)
		go func(v *models.IP) {
			if !CheckIP(v) {
				ProxyDel(v)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	x = conn.Count()
	log.Println("After check, DB has:", x, "records.")
}

// ProxyRandom .
func ProxyRandom() (ip *models.IP) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	conn := NewStorage()
	ips, _ := conn.GetAll()
	x := len(ips)

	return ips[r.Intn(x)]
}

// ProxyFind .
func ProxyFind(value string) (ip *models.IP) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	conn := NewStorage()
	ips, _ := conn.FindAll(value)
	x := len(ips)

	return ips[r.Intn(x)]
}

// ProxyAdd .
func ProxyAdd(ip *models.IP) {
	conn := NewStorage()
	_, err := conn.GetOne(ip.Data)
	if err != nil {
		conn.Create(ip)
	}
}

// ProxyDel .
func ProxyDel(ip *models.IP) {
	conn := NewStorage()
	if err := conn.Delete(ip); err != nil {
		log.Println(err.Error())
	}
}

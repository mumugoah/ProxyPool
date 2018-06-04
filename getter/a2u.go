package getter

import (
	"log"
	"strings"

	"github.com/mumugoah/ProxyPool/models"
	"github.com/parnurzeal/gorequest"
)

// Au2 get ip from 66ip.cn
func Au2() (result []*models.IP) {
	pollURL := "https://raw.githubusercontent.com/a2u/free-proxy-list/master/free-proxy-list.txt"
	_, body, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	proxies := strings.Split(body, "\n")

	for _, proxy := range proxies {
		ip := models.NewIP()
		ip.Data = proxy
		ip.Type = "http"
		//log.Println(ip.Data)
		result = append(result, ip)
	}
	log.Printf("Fetch Au2 done. Get %d IPs", len(result))

	return result
}

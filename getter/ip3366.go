package getter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/util"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html"
)

func IP3366() (result []*models.IP) {
	baseUrl := "http://www.ip3366.net/free"

	for i := 1; i < 5; i++ {
		url := fmt.Sprintf("%s/?page=%d/", baseUrl, i)

		resp, _, errs := gorequest.New().
			Get(url).
			Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36").
			End()

		if errs != nil {
			log.Println(errs)
			return
		}

		doc, err := htmlquery.Parse(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		htmlquery.FindEach(doc, "//table/tbody/tr", func(i int, n *html.Node) {
			ip := models.NewIP()
			ipS := htmlquery.FindOne(n, "./td[1]/text()").Data
			port := htmlquery.FindOne(n, "./td[2]/text()").Data
			ip.Data = util.GetIP(ipS) + ":" + port
			ip.Type = strings.ToLower(htmlquery.FindOne(n, "./td[4]/text()").Data)
			log.Println(ip.Data)
			result = append(result, ip)
		})

		time.Sleep(1 * time.Second)

	}

	log.Printf("Fetch IP3366 done. Get %d IPs", len(result))

	return result
}

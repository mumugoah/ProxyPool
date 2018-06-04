package getter

import (
	"log"

	"time"

	"strings"

	"fmt"

	"github.com/antchfx/htmlquery"
	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/util"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html"
)

// Xici get ip from xicidaili.com
func Xici() (result []*models.IP) {

	baseUrls := []string{
		"http://www.xicidaili.com/nn",
		"http://www.xicidaili.com/nt",
	}
	for _, baseUrl := range baseUrls {
		for i := 1; i < 5; i++ {
			resp, _, errs := gorequest.New().
				Get(fmt.Sprintf("%s/%d", baseUrl, i)).
				Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36").
				End()

			if errs != nil {
				log.Println(errs)
				return
			}

			doc, err := htmlquery.Parse(resp.Body)

			if err != nil {
				log.Println(err.Error())
				return
			}

			htmlquery.FindEach(doc, "//table/tbody/tr[position()>1]", func(i int, n *html.Node) {
				ip := models.NewIP()
				ipS := htmlquery.FindOne(n, "./td[2]/text()").Data
				port := htmlquery.FindOne(n, "./td[3]/text()").Data
				ip.Data = util.GetIP(ipS) + ":" + port
				ip.Type = strings.ToLower(htmlquery.FindOne(n, "./td[6]/text()").Data)
				result = append(result, ip)
			})

			time.Sleep(1 * time.Second)
		}
	}

	log.Printf("Fetch Xici done. Get %d IPs", len(result))
	return result
}

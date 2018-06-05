package getter

import (
	"log"

	"github.com/antchfx/htmlquery"
	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/util"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/net/html"
)

func Data5u() (result []*models.IP) {
	pollURLs := []string{
		"http://www.data5u.com/free/gngn/index.shtml",
		"http://www.data5u.com/free/gnpt/index.shtml",
		"http://www.data5u.com/free/gwgn/index.shtml",
		"http://www.data5u.com/free/gwpt/index.shtml",
	}

	for _, url := range pollURLs {
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
			log.Println(err.Error())
			return
		}

		htmlquery.FindEach(doc, "//div[@class='wlist']/ul/li/ul[position()>1]", func(i int, n *html.Node) {
			ip := models.NewIP()
			ipS := htmlquery.FindOne(n, "./span[1]/li/text()").Data
			port := htmlquery.FindOne(n, "./span[2]/li/text()").Data
			ip.Data = util.GetIP(ipS) + ":" + port
			ip.Type = htmlquery.FindOne(n, "./span[4]/li/a/text()").Data
			//log.Println(ip.Data)
			result = append(result, ip)
		})
	}

	log.Printf("Fetch Data5u done. Get %d IPs", len(result))

	return result
}

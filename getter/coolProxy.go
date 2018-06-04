package getter

import (
	"fmt"

	"log"

	"strings"

	"encoding/base64"

	"github.com/antchfx/htmlquery"
	"github.com/mumugoah/ProxyPool/models"
	"github.com/mumugoah/ProxyPool/util"
	"github.com/parnurzeal/gorequest"
	"github.com/robertkrimen/otto"
	"golang.org/x/net/html"
)

func CoolProxy() (result []*models.IP) {
	pollURL := "https://www.cool-proxy.net/proxies/http_proxy_list/country_code:/port:/anonymous:"

	vm := otto.New()

	vm.Run(`
		function str_rot13(str) {
			return (str + '').replace(/[a-z]/gi, function(s) {
				return String.fromCharCode(s.charCodeAt(0) + (s.toLowerCase() < 'n' ? 13 : -13));
			});
		}
	`)

	for i := 1; i < 5; i++ {
		url := fmt.Sprintf("%s/page:%d", pollURL, i)
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

		htmlquery.FindEach(doc, "//table/tbody/tr[position()>1]", func(i int, n *html.Node) {
			ip := models.NewIP()
			secretIp := htmlquery.FindOne(n, "./td[1]/script/text()")
			port := htmlquery.FindOne(n, "./td[2]/text()")
			if port == nil || secretIp == nil {
				return
			}
			//解密
			ipSe := strings.Split(secretIp.Data, "(str_rot13(\"")[1]
			ipSe = strings.Split(ipSe, "\")))")[0]

			ipB, err := vm.Run(fmt.Sprintf(`result = str_rot13("%s")`, ipSe))

			if err != nil {
				log.Printf("解析错误: %s", err)
			}
			ipS, err := base64.StdEncoding.DecodeString(ipB.String())
			if err != nil {
				log.Printf("解析错误: %s", err)
			}

			ip.Data = util.GetIP(string(ipS)) + ":" + port.Data
			ip.Type = "http"
			result = append(result, ip)
		})
	}

	log.Printf("Fetch CoolProxy done. Get %d IPs", len(result))

	return result
}

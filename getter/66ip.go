package getter

import (
	"log"
	"strings"

	"github.com/mumugoah/ProxyPool/models"
	"github.com/parnurzeal/gorequest"
)

// IP66 get ip from 66ip.cn
func IP66() (result []*models.IP) {
	url := "http://www.66ip.cn/mo.php?tqsl=100"
	_, body, errs := gorequest.New().
		Get(url).
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36").
		End()
	if errs != nil {
		log.Println(errs)
		return
	}
	body = strings.Split(body, "var mediav_ad_height = '60';")[1]
	body = strings.Split(body, "</script>")[1]
	body = strings.Split(body, "</div>")[0]
	body = strings.TrimSpace(body)
	body = strings.Replace(body, "	", "", -1)
	temp := strings.Split(body, "<br />")
	for index := 0; index < len(temp[:len(temp)-1]); index++ {
		ip := models.NewIP()
		ip.Data = strings.TrimSpace(temp[index])
		ip.Type = "http"
		//log.Println(ip.Data)
		result = append(result, ip)
	}
	log.Printf("Fetch IP66 done. Get %d IPs", len(result))

	return result
}

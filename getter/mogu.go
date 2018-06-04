package getter

import (
	"log"

	"net/http"

	"time"

	"io/ioutil"

	"github.com/mumugoah/ProxyPool/models"
	"github.com/tidwall/gjson"
)

func Mogu() (result []*models.IP) {
	url := "http://www.mogumiao.com/proxy/free/listFreeIp"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36")
	req.Header.Set("Referer", "http://www.mogumiao.com/web")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return
	}

	res := gjson.GetBytes(bytes, "msg").Array()

	for _, proxy := range res {
		ip := models.NewIP()
		ip.Data = proxy.Get("ip").Str + ":" + proxy.Get("port").Str
		ip.Type = "http"
		log.Println(ip.Data)
		result = append(result, ip)
	}

	return result
}

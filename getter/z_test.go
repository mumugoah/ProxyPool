package getter

import (
	"testing"
)

func TestIP66(t *testing.T) {
	proxies := IP66()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestAu2(t *testing.T) {
	proxies := Au2()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestCoolProxy(t *testing.T) {
	proxies := CoolProxy()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestData5u(t *testing.T) {
	proxies := Data5u()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestIP3366(t *testing.T) {
	proxies := IP3366()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestKuai(t *testing.T) {
	proxies := Kuai()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestMogu(t *testing.T) {
	proxies := Mogu()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

func TestXici(t *testing.T) {
	proxies := Xici()
	if len(proxies) < 1 {
		t.Errorf("爬取错误")
	}
}

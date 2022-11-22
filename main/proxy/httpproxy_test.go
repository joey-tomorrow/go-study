package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func Test_httpproxy(t *testing.T) {
	// 如果需要代理验证，那么如下进行设置
	// 否则直接设置为url.Parse("http://inproxy.sjtu.edu.cn:8000")
	//uri, err := url.Parse("http://34.120.112.65:443")
	//uri, err := url.Parse("http://34.120.112.65:443")


	//if err != nil{
	//	log.Fatal("parse url error: ", err)
	//}
	//log.Println(uri.User)
	//
	//client := http.Client{
	//	Transport: &http.Transport{
	//		// 设置代理
	//		Proxy: http.ProxyURL(uri),
	//	},
	//}

	proxyTransport, err := newTransport()
	if err != nil {
		panic(err)
	}

	//Create a custom client so we can make use of our RoundTripper
	//If you make use of http.Get(), the default http client located at http.DefaultClient is used instead
	//Since we have special needs, we have to make use of our own http.RoundTripper implementation
	client := &http.Client{
		Transport: proxyTransport,
		Timeout:   time.Second * 5,
	}

	//记录开始时间
	startTime := time.Now().UnixNano()
	resp, err := client.Get("http://external2.shopee.com.br/ext/brl_redesul/live/api/pullShopeeTest/api/services/shipment/tracking/")
	endTime := time.Now().UnixNano()
	//seconds:= float64((endTime - startTime) / 1e9)
	Milliseconds:= float64((endTime - startTime) / 1e6)
	//nanoSeconds:= float64(endTime - startTime)
	fmt.Println(Milliseconds)



	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//data, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(data))

}

type proxyTransport struct {
	originalTransport http.RoundTripper
	proxyTransport http.RoundTripper
}

func (p *proxyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Host == "api.pax.app.br" {
		return p.proxyTransport.RoundTrip(r)
	}
	return p.originalTransport.RoundTrip(r)
}

func newTransport() (*proxyTransport, error) {
	uri, err := url.Parse("http://34.120.112.65:443")
	if err != nil{
		log.Fatal("parse url error: ", err)
		return nil, err
	}

	return &proxyTransport{
		originalTransport: http.DefaultTransport,
		proxyTransport:    &http.Transport{Proxy: http.ProxyURL(uri)},
	}, nil
}
package GoQueryManager

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"log"
	"net/http"
	"time"
	"io/ioutil"
	//"golang.org/x/net/html"
	"strings"
)

const(
	UA_PHONE = "Mozilla/5.0 (Linux x86_64; Android 4.3; Nexus 10 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36";
	UA_PC = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36";
	UA_PAD = "Mozilla/5.0 (iPad; CPU OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1";

	TIME_OUT = 30 * time.Second
)

func Test(){
	//doc, err := goquery.NewDocument("http://metalsucks.net")
	doc,err := GetDocWithPC("http://metalsucks.net","")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}

func GetDocWithPhone(url,cookie string) (*goquery.Document,error) {
	headerMap := make(map[string]string)
	headerMap["Cookie"] = cookie
	headerMap["User-Agent"] = UA_PHONE
	return GetDoc(url,headerMap)
}

func GetDocWithPhoneAndHeader(url,cookie string, headerMap map[string]string) (*goquery.Document,error) {
	if headerMap == nil {
		headerMap = make(map[string]string)
	}
	headerMap["Cookie"] = cookie
	headerMap["User-Agent"] = UA_PHONE
	return GetDoc(url,headerMap)
}

func GetDocWithPC(url,cookie string)(*goquery.Document,error){
	headerMap := make(map[string]string)
	headerMap["Cookie"] = cookie
	headerMap["User-Agent"] = UA_PC
	return GetDoc(url,headerMap)
}

func GetDoc(url string,headerMap map[string]string)(*goquery.Document,error){
	//doc,err := goquery.NewDocumentFromResponse(GetUrlResponseByUrl(url,cookie,userAgent))
	resStr,err := GetUrlResponseByUrl(url,headerMap)
	if(nil != err){
		log.Println("get url response error:",err.Error())
		return nil,err
	}
	// Create and fill the document
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(resStr))
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return doc,nil
}
//Transport: &http.Transport{
//Dial: func(netw, addr string) (net.Conn, error) {
//deadline := time.Now().Add(TIME_OUT)
//c, err := net.DialTimeout(netw, addr, TIME_OUT)
//if err != nil {
//return nil, err
//}
//c.SetDeadline(deadline)
//return c, nil
//},ResponseHeaderTimeout: TIME_OUT,
//},
func GetUrlResponseByUrl(url string,headerMap map[string]string) (string,error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest error:", err.Error())
		return "",err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do error:", err.Error())
		return "",err
	}

	defer resp.Body.Close()

	var textByte []byte
	textByte, _ = ioutil.ReadAll(resp.Body)
	log.Println("result:", string(textByte[0:20]))
	return string(textByte),nil
}

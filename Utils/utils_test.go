package Utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWriteLog(t *testing.T) {
	err := WriteLog("清零成功","log.txt")
	if err != nil {
		fmt.Println("success")
	}
}
func TestAcquiredLock(t *testing.T) {
	var wg sync.WaitGroup
	fmt.Println(time.Now())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			err := WriteLog("test lock","log.txt")
			if err != nil {
				log.Fatal("something wrong")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now())    
}

func TestSend(t *testing.T) {
	client := &http.Client{}
	var data = strings.NewReader(`rwlx=2&xkly=0&bklx_id=0&sfkkjyxdxnxq=0&xqh_id=2&jg_id=04&njdm_id_1=2020&zyh_id_1=A737C8D6D90A009EE053C0A86D5D8E09&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09&zyfx_id=C1B92046DE0D0068E053C0A86D5C10B1&njdm_id=2020&bh_id=201048001&xbm=1&xslbdm=wlb&mzm=01&xz=4&ccdm=3&xsbj=4294967296&sfkknj=0&sfkkzy=0&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&kkbk=0&kkbkdj=0&sfkgbcx=0&sfrxtgkcxd=0&tykczgxdcs=0&xkxnm=2022&xkxqm=12&kklxdm=10&bbhzxjxb=0&rlkz=0&xkzgbj=0&kspage=1&jspage=10&jxbzb=`)
	req, err := http.NewRequest("POST", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?gnmkdm=N253512&su=05203744&sf_request_type=ajax", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Cookie", "JSESSIONID=48A392FB96E1868EB7A2AA580E270C40; BIGipServerp_new_hr_-_authserver.cumt.edu.cn=590601179.20480.0000; TWFID=0bc0c6ced2bce447")
	req.Header.Set("Origin", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118")
	req.Header.Set("Referer", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default&su=05203744")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

func TestGetDetail(t *testing.T) {
	client := &http.Client{}
	var data = strings.NewReader(`yl_list%5B0%5D=1&rwlx=2&xkly=0&bklx_id=0&sfkkjyxdxnxq=0&xqh_id=2&jg_id=04&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09&zyfx_id=C1B92046DE0D0068E053C0A86D5C10B1&njdm_id=2020&bh_id=201048001&xbm=1&xslbdm=wlb&mzm=01&xz=4&bbhzxjxb=0&ccdm=3&xsbj=4294967296&sfkknj=0&sfkkzy=0&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&kkbk=0&kkbkdj=0&xkxnm=2022&xkxqm=12&xkxskcgskg=1&rlkz=0&kklxdm=10&kch_id=ED27423051B59000E0531F70A8C0B2D2&jxbzcxskg=0&xkkz_id=EF9D03B5576FBC01E0531F70A8C07A1C&cxbj=0&fxbj=0`)
	req, err := http.NewRequest("POST", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?gnmkdm=N253512&su=05203744&sf_request_type=ajax", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Cookie", "JSESSIONID=48A392FB96E1868EB7A2AA580E270C40; BIGipServerp_new_hr_-_authserver.cumt.edu.cn=590601179.20480.0000; TWFID=0bc0c6ced2bce447")
	req.Header.Set("Origin", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118")
	req.Header.Set("Referer", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default&su=05203744")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}

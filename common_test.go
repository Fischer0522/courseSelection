package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	cookie := Cookie{
		JSESS: "",
		BIGipServer: "",
		TWFID: "",

	}
	result := searchClassCommon(cookie)
	for _,element := range result {
		infolist := getDetailCommon(element.kch_id,cookie)
		fmt.Println(infolist[0])
	}

}
func TestDetail(t *testing.T) {
	client := &http.Client{}
	var data = strings.NewReader(`yl_list%5B0%5D=1&rwlx=2&xkly=0&bklx_id=0&sfkkjyxdxnxq=0&xqh_id=2&jg_id=04&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09&zyfx_id=C1B92046DE0D0068E053C0A86D5C10B1&njdm_id=2020&bh_id=201048001&xbm=1&xslbdm=wlb&mzm=01&xz=4&bbhzxjxb=0&ccdm=3&xsbj=4294967296&sfkknj=0&sfkkzy=0&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&kkbk=0&kkbkdj=0&xkxnm=2022&xkxqm=12&xkxskcgskg=1&rlkz=0&kklxdm=10&kch_id=ED182959592F481BE0531F70A8C0B8A4&jxbzcxskg=0&xkkz_id=EF9D03B5576FBC01E0531F70A8C07A1C&cxbj=0&fxbj=0`)
	req, err := http.NewRequest("POST", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?gnmkdm=N253512&su=&sf_request_type=ajax", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Cookie", "JSESSIONID=FC2292E77D67C4DC0A09EBA2065AA056; BIGipServerp_new_hr_-_authserver.cumt.edu.cn=590601179.20480.0000; TWFID=2faba07b1454d880")
	req.Header.Set("Origin", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118")
	req.Header.Set("Referer", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default&su=")
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

func TestSelect(t *testing.T) {
	cookie := Cookie{
		JSESS: "",
		BIGipServer: "",
		TWFID: "",

	}
	result := searchClassCommon(cookie)
	for _,element := range result {
		fmt.Println(element)
	}
}
func TestDetailInfo(t *testing.T) {
	cookie := Cookie{
		JSESS: "",
		BIGipServer: "",
		TWFID: "",

	}
	doTaskWithGoRoutinesCommon(cookie)


}
func TestChoose(t *testing.T) {
	cookie := Cookie{
		JSESS: "",
		BIGipServer: "",
		TWFID: "",

	}
	jxs:="73e7e9842ca6a12672efa2d14f275563280e11c2b474aa492d49c11f8fad8f6e67d713a838236a8283f7214e83366fc12dcff9cb8ca286ee6955dcd0605d713af3a2792a5f011c5a579cc9de275cb7108a1aefb32c621367120cc0bff8a9fef3b408cb4bf2d1584511b5fe19d48b003a859740cf4a934ad6536c7f9b48cffc27"
	kch_id := "ED27423051B59000E0531F70A8C0B2D2"
	kcmc := "(Q30258)%E6%B2%B9%E6%B0%94%E4%B8%8E%E5%9C%B0%E7%83%AD%E5%BC%80%E9%87%87%E2%80%94%E2%80%94%E5%B2%A9%E7%9F%B3%E5%8A%9B%E5%AD%A6%E7%9A%84%E5%BA%94%E7%94%A8+-+2.0+%E5%AD%A6%E5%88%86"

	chooseClassCommon(jxs,kch_id,kcmc,cookie)
}


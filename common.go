package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func sendRequestCommon(newReader string,newRequestUrl string,cookie Cookie) []byte{
	client := &http.Client{}
	var data = strings.NewReader(newReader)
	req, err := http.NewRequest("POST", newRequestUrl, data)
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
	return bodyText

}


func searchClassCommon(cookie Cookie) []SearchResult {
	searchList := make([]SearchResult,0)
	newReader := "yl_list%5B0%5D=1&rwlx=2&xkly=0&bklx_id=0&sfkkjyxdxnxq=0&xqh_id=2&jg_id=04&njdm_id_1=2020&zyh_id_1=A737C8D6D90A009EE053C0A86D5D8E09&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09&zyfx_id=C1B92046DE0D0068E053C0A86D5C10B1&njdm_id=2020&bh_id=201048001&xbm=1&xslbdm=wlb&mzm=01&xz=4&ccdm=3&xsbj=4294967296&sfkknj=0&sfkkzy=0&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&kkbk=0&kkbkdj=0&sfkgbcx=0&sfrxtgkcxd=0&tykczgxdcs=0&xkxnm=2022&xkxqm=12&kklxdm=10&bbhzxjxb=0&rlkz=0&xkzgbj=0&kspage=1&jspage=10&jxbzb="

	url := "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?gnmkdm=N253512&su=&sf_request_type=ajax"

	bodyText := sendRequest(newReader,url,cookie)

	fmt.Printf("%s\n", bodyText)

	var classList ClassList
	json.Unmarshal([]byte(bodyText), &classList)

	infolist :=classList.TmpList
	for _,element := range infolist {
		se := SearchResult{
			kch: element.Kch,
			kch_id: element.KchID,
			jxb_id: element.JxbID,
			xf: element.Xf,
			kcmc: element.Kcmc,
		}
		if len(searchList) == 0 {
			searchList = append(searchList, se)
		}
		if len(searchList) > 0 && searchList[len(searchList)-1].kch != se.kch {
			searchList = append(searchList, se)
		}

	}


	return searchList

}
func getDetailCommon(kch_id string,cookie Cookie) []DetailResult {
	content := "yl_list%5B0%5D=1" +
		"&rwlx=2" +
		"&xkly=0" +
		"&bklx_id=0" +
		"&sfkkjyxdxnxq=0" +
		"&xqh_id=2" +
		"&jg_id=04" +
		"&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09" +
		"&zyfx_id=C1B92046DE0D0068E053C0A86D5C10B1" +
		"&njdm_id=2020&bh_id=201048001&xbm=1" +
		"&xslbdm=wlb" +
		"&mzm=01" +
		"&xz=4" +
		"&bbhzxjxb=0" +
		"&ccdm=3" +
		"&xsbj=4294967296" +
		"&sfkknj=0" +
		"&sfkkzy=0" +
		"&kzybkxy=0" +
		"&sfznkx=0" +
		"&zdkxms=0" +
		"&sfkxq=1" +
		"&sfkcfx=0" +
		"&kkbk=0" +
		"&kkbkdj=0" +
		"&xkxnm=2022" +
		"&xkxqm=12" +
		"&xkxskcgskg=1" +
		"&rlkz=0" +
		"&kklxdm=10" +
		"&kch_id="+kch_id +
		"&jxbzcxskg=0" +
		"&xkkz_id=EF9D03B5576FBC01E0531F70A8C07A1C" +
		"&cxbj=0" +
		"&fxbj=0"
	url := "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?gnmkdm=N253512&su=&sf_request_type=ajax"

	bodyText := sendRequest(content,url,cookie)

	var detailInfo []DetailInfoCommon
	detailResultList := make([]DetailResult,0)
	json.Unmarshal(bodyText, &detailInfo)
	for _,element := range detailInfo {
		detailResult := DetailResult{
			jxb_ids: element.DoJxbID,
		}

		detailResultList = append(detailResultList, detailResult)

	}

	return detailResultList

}

func chooseClassCommon(jxb_ids string,kch_id string,kcmc string,cookie Cookie) {

	content:= "jxb_ids="+jxb_ids +
		"&kch_id="+kch_id +
		"&kcmc="+kcmc +
		"&rwlx=2" +
		"&rlkz=0" +
		"&rlzlkz=1" +
		"&sxbj=1" +
		"&xxkbj=0" +
		"&qz=0" +
		"&cxbj=0" +
		"&xkkz_id=EF9D03B5576FBC01E0531F70A8C07A1C" +
		"&njdm_id=2020" +
		"&zyh_id=A737C8D6D90A009EE053C0A86D5D8E09" +
		"&kklxdm=10" +
		"&xklc=1" +
		"&xkxnm=2022" +
		"&xkxqm=12"

	url := "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_xkBcZyZzxkYzb.html?gnmkdm=N253512&su=&sf_request_type=ajax"


	bodyText := sendRequest(content,url,cookie)
	var chooseResult ChooseResult
	json.Unmarshal(bodyText,&chooseResult)
	if chooseResult.Flag == "-1" {
		fmt.Printf("%s:当前没课",kcmc)
	} else if chooseResult.Flag == "1" {
		fmt.Printf("%s:选课成功",kcmc)
	}
	fmt.Println()
	fmt.Printf("%s\n", bodyText)


}
var wg1 sync.WaitGroup

func doTaskWithGoRoutinesCommon(cookie Cookie){
	// set your cookie here

	for i := 0;i < 3;i++{
		wg1.Add(1)

		go func(cookie Cookie) {
			ls :=  searchClassCommon(cookie)
			if len(ls) != 0 {
				for _,element := range ls {

					nameToSend := "("+element.kch+")"+element.kcmc+" - "+element.xf+" 学分"
					fmt.Println(nameToSend)
					detailResult := getDetailCommon(element.kch_id,cookie)
					for _,detail := range detailResult {
						chooseClassCommon(detail.jxb_ids,element.kch_id,nameToSend,cookie)
						time.Sleep(time.Duration(1)*time.Second)
					}
				}
			} else {
				fmt.Println("暂无课程")
			}


			wg1.Done()
		}(cookie)


	}
	wg1.Wait()
}

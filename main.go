package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jwb/Utils"
	"jwb/model"
	"log"
	"net/http"
	 "net/url"
	"strings"
	"sync"
	"time"
)

var (
	// 设置学号
	studentId = ""
	urlToSearch = "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?gnmkdm=N253512&su="+studentId+"&sf_request_type=ajax"
	urlToGetDetail = "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?gnmkdm=N253512&su="+studentId+"&sf_request_type=ajax"
	urlToChoose = "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzbjk_xkBcZyZzxkYzb.html?gnmkdm=N253512&su="+studentId+"&sf_request_type=ajax"
	// 设置 cookie
	cookieJESS = ""
	cookieBigServer = ""
	cookieTWFID = ""
	// 选课的学年
	classYear  = "2022"
	// 设置自己的年级
	grade = "2020"
	majorId = "A737C8D6D90A009EE053C0A86D5D8E09"
	classType1 = "10" // 10为公选课 18为劳动
	classType2 = "1" // 1为公选 0 为劳动


	//公选课：EF9D03B5576FBC01E0531F70A8C07A1C 劳动课：EF9E3D0F99E56F21E0531F70A8C0CC18
	//不同人可能不一样，需要提前查询
	classTypeId = "EF9D03B5576FBC01E0531F70A8C07A1C"

	// 是否根据课程名称查询 注：需要提供准确名称，
	//如果提供模糊名称则会查出多个课最后去选择第一个，后续的可能和第一个冲突而选不上
	selectByName = true
	className = "风景背后的地质学"

	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54"

)
func sendRequest(newReader string,newRequestUrl string,cookie string) ([]byte,error){
	client := &http.Client{}
	var data = strings.NewReader(newReader)
	req, err := http.NewRequest("POST",newRequestUrl , data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Origin", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118")
	req.Header.Set("Referer", "https://jwxk1-cumt-edu-cn.webvpn.cumt.edu.cn:8118/jwglxt/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default&su="+studentId)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent",userAgent)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Microsoft Edge";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return make([]byte,0),err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)

	return bodyText,nil

}


func searchClass(urlStr string,cookie string) ([]model.SearchResult,error) {
	body := make(url.Values)
	searchList := make([]model.SearchResult,0)
	if selectByName {
		body.Set("filter_list[0]",className)
	}

	body.Set("yl_list[0]","1") // 是否只查看余量 1为查看 若选择全部则无需带该条参数
	body.Set("rwlx","2") // 任务流程
	body.Set("xkly","0")
	body.Set("bklx_id","0")
	body.Set("sfkkjyxdxnxq","0")
	body.Set("xqh_id","2")
	body.Set("jg_id","04") // 结果
	body.Set("njdm_id_1","2020") // 年级代码
	body.Set("zyh_id_1",majorId) // 专业号
	body.Set("zyh_id",majorId) // 专业号
	body.Set("zyfx_id","C1B92046DE0D0068E053C0A86D5C10B1") // 专业方向
	body.Set("njdm_id",grade) // 年级代码
	body.Set("bh_id","201048001") // 班号
	body.Set("xbm","1")
	body.Set("xslbdm","wlb") // 学生类别代码
	body.Set("mzm","01")
	body.Set("xz","4")
	body.Set("ccdm","3")
	body.Set("xsbj","4294967296") // 学生标记
	body.Set("sfkknj","0")
	body.Set("sfkkzy","0")
	body.Set("kzybkxy","0")
	body.Set("sfznkx","0")
	body.Set("zdkxms","0")
	body.Set("sfkxq","1") // 是否跨校区
	body.Set("sfkcfx","0") // 是否重复可选
	body.Set("kkbk","0")
	body.Set("kkbkdj","0")
	body.Set("sfkgbcx","0")
	body.Set("sfrxtgkcxd","0")
	body.Set("tykczgxdcs","0")
	body.Set("xkxnm",classYear) // 选课学年
	body.Set("xkxqm","12") // 选课月份
	body.Set("kklxdm",classType1) // 课程种类
	body.Set("bbhzxjxb","0")
	body.Set("rlkz","0")  // 人流控制
	body.Set("xkzgbj","0")
	body.Set("kspage","1") // 起始页
	body.Set("jspage","10") // 终止页
	body.Set("jxbzb","")


		bodyText,errRequest := sendRequest(body.Encode(),urlStr,cookie)
		if errRequest != nil {
			log.Fatal("sendRequest failed when searchClass")
			return nil, errRequest
		}
		fmt.Printf("%s\n", bodyText)


		var classList model.ClassList
	err := json.Unmarshal(bodyText, &classList)
	if err != nil {
		log.Fatal("列表信息解析失败")
	}

		infolist := classList.TmpList
		for _,element := range infolist {
			se := model.SearchResult{
				Kch:    element.Kch,
				Kch_id: element.KchID,
				Jxb_id: element.JxbID,
				Xf:     element.Xf,
				Kcmc:   element.Kcmc,
			}
			if len(searchList) == 0 {
				searchList = append(searchList, se)
			}
			if len(searchList) > 0 && searchList[len(searchList)-1].Kch != se.Kch {
				searchList = append(searchList, se)
			}

		}


	return searchList,err

}

func getDetail(kch_id string,urlStr string,cookie string) ([]model.DetailResult,error) {
	body := make(url.Values)
	body.Set("yl_list%5B0%5D", "1")
	body.Set("rwlx", "2")
	body.Set("xkly", "0")
	body.Set("bklx_id", "0")
	body.Set("sfkkjyxdxnxq", "0")
	body.Set("xqh_id", "2")
	body.Set("jg_id", "04")
	body.Set("zyh_id", majorId)
	body.Set("zyfx_id", "C1B92046DE0D0068E053C0A86D5C10B1")
	body.Set("njdm_id", grade)
	body.Set("bh_id", "201048001")
	body.Set("xbm", "1")
	body.Set("xslbdm", "wlb")
	body.Set("mzm", "01")
	body.Set("xz", "4")
	body.Set("bbhzxjxb", "0")
	body.Set("ccdm", "3")
	body.Set("xsbj", "4294967296") //学生标记
	body.Set("sfkknj", "0")
	body.Set("sfkkzy", "0")
	body.Set("kzybkxy", "0")
	body.Set("sfznkx", "0")
	body.Set("zdkxms", "0")
	body.Set("sfkxq", "1")
	body.Set("sfkcfx", "0")
	body.Set("kkbk", "0")
	body.Set("kkbkdj", "0")
	body.Set("xkxnm", classYear)
	body.Set("xkxqm", "12")

	body.Set("xkxskcgskg", classType2)

	body.Set("rlkz", "0")

	body.Set("kklxdm", classType1) // 10: 公选课  18: 劳动实践

	body.Set("kch_id", kch_id)
	body.Set("jxbzcxskg", "0")
	body.Set("xkkz_id", classTypeId) // 劳动实践课
	body.Set("cxbj", "0")
	body.Set("fxbj", "0")

	bodyText,requestErr := sendRequest(body.Encode(),urlStr,cookie)
	if requestErr != nil {
		log.Fatal("send request failed when getting detail")
		return nil, requestErr
	}
	fmt.Printf("%s\n", bodyText)
	if classType1 == "18" {
		var detailInfo []model.DetailInfo
		detailResultList := make([]model.DetailResult,0)
		err := json.Unmarshal(bodyText, &detailInfo)
		if err != nil {
			log.Fatal("详细信息解析失败！")
		}
		for _,element := range detailInfo {
			detailResult := model.DetailResult{
				Jxb_ids: element.DoJxbID,
			}

			detailResultList = append(detailResultList, detailResult)

		}

		return detailResultList,err
	} else {
		var detailInfo []model.DetailInfoCommon
		detailResultList := make([]model.DetailResult,0)
		err := json.Unmarshal(bodyText, &detailInfo)
		if err != nil {
			log.Fatal("公选课详细信息解析失败！")
		}
		for _,element := range detailInfo {
			detailResult := model.DetailResult{
				Jxb_ids: element.DoJxbID,
			}

			detailResultList = append(detailResultList, detailResult)

		}

		return detailResultList,err
	}


}

func selectClass(jxbIds string,kch_id string,kcmc string,urlStr string,cookie string) error {
	body := make(url.Values)
	body.Set("jxb_ids",jxbIds)
	body.Set("kch_id",kch_id) // 课程号
	body.Set("kcmc",kcmc) // 课程名称
	body.Set("rwlx","2") // 任务类型
	body.Set("rlkz","0") // 人数控制
	body.Set("rlzlkz","1") // 人数控制
	body.Set("sxbj","1") // 上下学期
	body.Set("xxkbj","0") // 选修课
	body.Set("qz","0") // 权重
	body.Set("cxbj","0") // 重修标记
	body.Set("xkkz_id",classTypeId) // 选课可组id 劳动课
	body.Set("njdm_id","2020") // 年级
	body.Set("zyh_id","A737C8D6D90A009EE053C0A86D5D8E09") // 专业号
	body.Set("kklxdm","18") // 课程类型 10为通选课 18为劳动课
	body.Set("xklc","1") // 选课流程
	body.Set("xkxnm",classYear) // 学年
	body.Set("xkxqm","12") // 学期


	bodyText,requestErr := sendRequest(body.Encode(),urlStr,cookie)
	if requestErr != nil {
		log.Fatal("send request failed when choosing class")
		return requestErr
	}
	var chooseResult model.ChooseResult
	err := json.Unmarshal(bodyText, &chooseResult)
	if err != nil {
		return err
	}

	if chooseResult.Flag == "-1" {
		record := fmt.Sprintf("%s %s  当前没课",time.Now(),kcmc)
		go func() {
			logError := Utils.WriteLog(record,"log/errorLog.txt")
			if logError != nil {
				fmt.Println("持久化失败")
			}
		}()
		fmt.Printf("%s %s:当前没课",time.Now(),kcmc)
	} else if chooseResult.Flag == "1" {
		record := fmt.Sprintf("%s %s 选课成功",time.Now(),kcmc)
		fmt.Printf("%s  %s:选课成功",time.Now(),kcmc)
		go func(record string) {
			logError := Utils.WriteLog(record,"log/successLog.txt")
			if logError != nil {
				fmt.Println("持久化失败")
			}
		}(record)
	} else {
		record := fmt.Sprintf("%s %s %s",time.Now(),chooseResult.Msg,kcmc)
		go func() {
			logError := Utils.WriteLog(record,"log/errorLog.txt")
			if logError != nil {
				fmt.Println("持久化失败")
			}
		}()
		fmt.Printf("%s %s:%s",time.Now(),chooseResult.Msg,kcmc)
	}
	fmt.Println()
	fmt.Printf("%s\n", bodyText)
	return nil
}
var wg sync.WaitGroup

func doTaskWithGoRoutines(cookie string){
	// set your cookie here

	for i := 0;i < 3;i++{
		wg.Add(1)

		go func(cookie string) {
			ls,err :=  searchClass(urlToSearch,cookie)
			if err != nil {
				log.Fatal("检查cookie是否过期或者url是否正确")
			}
			if len(ls) != 0 {
				for _,element := range ls {

					nameToSend := "("+element.Kch +")"+element.Kcmc +" - "+element.Xf +" 学分"
					detailResult,err2 := getDetail(element.Kch_id,urlToGetDetail,cookie)
					if err2 != nil {
						log.Fatal("详情获取失败")
					}
					for _,detail := range detailResult {
						err3 := selectClass(detail.Jxb_ids,element.Kch_id,nameToSend,urlToChoose,cookie)
						if err3 != nil {
							log.Fatal("选课失败")
						}
						time.Sleep(time.Duration(1)*time.Second)
					}
				}
			} else {
				fmt.Printf("%s 暂无课程\n",time.Now())
			}


			wg.Done()
		}(cookie)


	}
	wg.Wait()
}

func doTaskSequential(cookie string) {

		ls,err1 := searchClass(urlToSearch,cookie)
		if err1 != nil {
			log.Fatal("列表信息获取失败")
		}

		for _, element := range ls {

			nameToSend := "(" + element.Kch + ")" + element.Kcmc + " - " + element.Xf + " 学分"
			detailResult,err2 := getDetail(urlToGetDetail,element.Kch_id, cookie)
			if err2 != nil {
				log.Fatal("详情获取失败")
			}
			for _, detail := range detailResult {
				err3 := selectClass(detail.Jxb_ids, element.Kch_id, nameToSend,urlToChoose, cookie)
				if err3 != nil {
					log.Fatal("选课失败")
				}
				time.Sleep(time.Duration(1) * time.Second)
			}
		}


}

func main() {
	// set your cookie here
	cookie := model.Cookie{
		JSESS: cookieJESS,
		BIGipServer: cookieBigServer,
		TWFID: cookieTWFID,

	}
	cookieStr := Utils.GenerateCookie(cookie)

	for {
		doTaskWithGoRoutines(cookieStr)
	}

}


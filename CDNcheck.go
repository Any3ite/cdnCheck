package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func chkCDN(target string) (cdnResult string) {
	cname, err := net.LookupCNAME(target)
	if err != nil {
		fmt.Println(err)
	}
	cdnlist := map[string]string{
		"阿里云":                 "aliyun",
		"阿里云 - CS":            "aliyuncs",
		"阿里云 - kunlun":        "kunluncan",
		"aliyun - kunlun":"kunlunar",
		"360":                 "360",
		"上海云盾":                "yundun",
		"加速乐":                 "jiashule",
		"百度 - shifen":         "shifen",
		"百度 -  baidu":         "baidu",
		"Akamai":              "akamai",
		"亚马逊AWS":              "amazon",
		"亚马逊cloudfront":       "cloudfront",
		"azion":               "azion",
		"cachefly":            "cachefly",
		"cdn77":               "cdn77",
		"cdnetworks":          "cdnetworks",
		"cdnify":              "cdnify",
		"cdnsun":              "cdnsun",
		"cdnvideo":            "cdnvideo",
		"ChinaCache":          "ccgslb",
		"CHINACache":          "chinacache",
		"chinacache":          "lxdns",
		"ChinaNetCenter":      "wscloudcdn",
		"CloudFlare":          "cloudflare",
		"EdgeCast":            "edgecastcdn",
		"Fastly":              "fastly",
		"Highwinds":           "hwcdn",
		"Incapsula":           "incapdns",
		"Internap":            "internapcdn",
		"KeyCDN":              "kxcdn",
		"keyCDN":              "awsdns",
		"Leaseweb":            "lswcdn",
		"Level3":              "fpbns",
		"level3":              "footprint",
		"Limelight":           "llnwd",
		"MaxCDN":              "netdna",
		"NGENIX":              "ngenix",
		"QUANTIL":             "mwcloudcdn",
		"quantil":             "speedcdns",
		"SkyparkCDN":          "skyparkcdn",
		"Tata Communications": "bitgravity",
		"微软Azure":             "azureedge",
		"Ananke":              "anankecdn",
		"Presscdn":            "presscdn",
		"Telefonica":          "telefonica",
		"腾讯":                  "cdntip",
		"腾讯鹅厂":                "cdn.dnsv1",
	}
	/*接收 main函数传递来的 字符串，并在返回的cname内容中查找特定字符串，如果结果为true则可确认CDN归属
	  创建一个字典 map 然后遍历 map的 vaule ， 去做if判断， 如果结果为 true， 则 result = true 的条件 */
	for mapkey, mapvalue := range cdnlist {
		if strings.Contains(cname, mapvalue) == true {
			//fmt.Printf("CDN服务器提供商为 : %s", mapkey)
			cdnResult = mapkey
			break
		}
	}

	return cdnResult
}
func main() {

	cnametest := os.Args[1]
	rest := chkCDN(cnametest)
	if len(rest) < 2 {
		fmt.Println(" 未使用CDN或CDN提供商特征未被收录")
	} else {
		fmt.Printf(" CDN服务器提供商为 : %s", rest+"\r\n")
	}
}

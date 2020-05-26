package dns

import (
	"encoding/json"
	"github.com/miekg/dns"
	"log"
	"net"
	"strings"
)

// 获取记录
// 参数domain: 域名 lilu.red
// 参数t: 记录类型 https://pkg.go.dev/github.com/miekg/dns@v1.1.29?tab=doc#TypeNone
// 返回: 文本数组的JSON字符串
func DigShort(domain string, t int) string {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), uint16(t))
	m.RecursionDesired = true
	r, _, err := c.Exchange(m, net.JoinHostPort("8.8.8.8", "53"))
	if r == nil {
		log.Println("查询DNS记录失败:", err)
		return "[]"
	}
	if r.Rcode != dns.RcodeSuccess {
		log.Println("查询DNS记录失败:", r.Rcode)
		return "[]"
	}

	var ta []string
	for _, a := range r.Answer {
		txt := strings.ReplaceAll(a.String(), a.Header().String(), "")
		// TODO 中文不知如何转码
		// buy=\\229\\138\\159\\232\\131\\189\\233\\156\\128\\232\\166\\129\\230\\191\\128\\230\\180\\187\\229\\144\\142\\230\\137\\141\\232\\131\\189\\228\\189\\191\\231\\148\\168\\239\\188\\140\\232\\175\\183\\229\\164\\141\\229\\136\\182\\231\\148\\179\\232\\175\\183\\228\\191\\161\\230\\129\\175\\229\\138\\160\\229\\190\\174\\228\\191\\161957388815\\232\\180\\173\\228\\185\\176\\227\\128\\130
		ta = append(ta, txt)
	}
	jsonBytes, _ := json.Marshal(ta)
	result := string(jsonBytes)
	if result == "null" {
		result = "[]"
	}
	return result
}

// 注意: 与标准结果不同,没有引号. Android查询没有结果!
// http://networkbit.ch/golang-dns-lookup/#lookuptxt
func LookupTXT(domain string) string {
	as, err := net.LookupTXT(domain)
	if err != nil {
		log.Println(err)
		return "[]"
	}
	if len(as) == 0 {
		return "[]"
	}

	var ta []string
	for _, txt := range as {
		ta = append(ta, txt)
	}
	jsonBytes, _ := json.Marshal(ta)
	return string(jsonBytes)
}

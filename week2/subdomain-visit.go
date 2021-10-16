package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	if len(cpdomains) == 0 {
		return cpdomains
	}

	//1.域名拆分，hashmap，key是域名，value是出现的次数
	hashMap := make(map[string]int)
	resString := []string{}
	for _, str := range cpdomains {
		//1.以空格为第一次split
		parts := strings.Split(str, " ") //1.["900 google.mail.com"],parts[0]=900,parts[1]="google.mail.com"
		count, _ := strconv.Atoi(parts[0]) //2.string to int,ignore err
		domain := parts[len(parts) - 1]

		//2.以点为第二次split
		subDomain := strings.Split(domain, ".")

		//3.拼接子域名
		var joinStringArray []string
		for i := len(subDomain) - 1; i >= 0; i-- {
			joinStringArray = append([]string{subDomain[i]}, joinStringArray...) //插入数组的头部
			joinString := strings.Join(joinStringArray, ".")
			hashMap[joinString] += count
		}
	}

	//4.遍历map，重新拼接字符串数组
	for k, v := range hashMap {
		joinStringOutput := strings.Join([]string{strconv.Itoa(v), k} ," ")
		resString = append(resString, joinStringOutput)
	}
	return resString
}

func main() {
	domain := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	//["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
	res := subdomainVisits(domain)
	fmt.Printf("string %s\n",res)
}
